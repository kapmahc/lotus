package auth

import (
	"fmt"
	"time"

	"github.com/SermoDigital/jose/jws"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

//AddLog add log
func AddLog(user uint, msg string) {

	if _, err := orm.NewOrm().Insert(&Log{UserID: user, Message: msg}); err != nil {
		beego.Error(err)
	}
}

//GenerateUserClaims generate user claims
func GenerateUserClaims(u *User) jws.Claims {
	cm := jws.Claims{}
	cm.SetSubject(u.Name)
	cm.Set("uid", u.UID)
	cm.Set("id", u.ID)
	cm.Set("roles", GetAuthority(u.ID, "-", 0))
	return cm
}

//ConfirmUser confirm user
func ConfirmUser(user *User) {
	now := time.Now()
	user.ConfirmedAt = &now

	if _, err := orm.NewOrm().Update(user, "confirmed_at", "updated_at"); err != nil {
		beego.Error(err)
	}
}

//AddEmailUser add email user
func AddEmailUser(email, name, password string) (*User, error) {
	var u User
	o := orm.NewOrm()
	count, err := o.
		QueryTable(&u).
		Filter("provider_type", ProvideByEmail).
		Filter("provider_id", email).Count()
	if err != nil {
		return nil, err
	}
	if count != 0 {
		return nil, fmt.Errorf("%s already exists", email)
	}

	u.Email = email
	u.Name = name
	u.ProviderID = email
	u.ProviderType = "email"
	u.SignInCount = 1
	u.LastSignInIP = "0.0.0.0"
	u.CurrentSignInIP = "0.0.0.0"
	u.SetPassword(password)
	u.SetGravatarLogo()
	u.SetUID()
	u.Home = fmt.Sprintf("%s/users/%s", beego.AppConfig.String("appurl"), u.UID)

	_, err = o.Insert(&u)
	return &u, err
}

//AddOpenIDUser add openid user
func AddOpenIDUser(pid, pty, email, name, home, logo string) (*User, error) {
	var u User
	now := time.Now()
	o := orm.NewOrm()
	err := o.QueryTable(&u).
		Filter("provider_type", pty).
		Filter("provider_id", pid).
		One(&u)
	if err == nil {
		u.Email = email
		u.Name = name
		u.Logo = logo
		u.Home = home
		u.SignInCount++
		u.LastSignInAt = &now
		_, err = o.Update(
			&u,
			"email",
			"name",
			"logo",
			"home",
			"sign_in_count",
			"last_sign_in_at",
			"updated_at",
		)
	}
	if err == orm.ErrNoRows {
		u.Email = email
		u.Name = name
		u.Logo = logo
		u.Home = home
		u.ProviderID = pid
		u.ProviderType = pty
		u.ConfirmedAt = &now
		u.SignInCount = 1
		u.LastSignInAt = &now

		u.SetUID()
		_, err = o.Insert(&u)
	}
	return &u, err
}

//GetUserByUID get user by uid
func GetUserByUID(uid string) (*User, error) {
	var u User
	err := orm.NewOrm().QueryTable(&u).Filter("uid", uid).One(&u)
	return &u, err
}

//GetUserByEmail get user by email
func GetUserByEmail(email string) (*User, error) {
	var u User
	err := orm.NewOrm().
		QueryTable(&u).
		Filter("provider_type", ProvideByEmail).
		Filter("provider_id", email).
		One(&u)
	return &u, err
}

//GetAuthority get user's role names
func GetAuthority(user uint, rty string, rid uint) []string {
	var items []Role
	o := orm.NewOrm()
	if _, err := o.QueryTable(new(Role)).
		Filter("resource_type", rty).
		Filter("resource_id", rid).
		All(&items, "name", "id"); err != nil {
		beego.Error(err)
	}
	var roles []string
	for _, r := range items {
		var pm Permission
		if err := o.QueryTable(&pm).
			Filter("role_id", r.ID).
			Filter("user_id", user).
			One(&pm); err != nil {
			beego.Error(err)
			continue
		}
		if pm.Enable() {
			roles = append(roles, r.Name)
		}
	}

	return roles
}

//Is is role ?
func Is(user uint, name string) bool {
	return Can(user, name, DefaultResourceType, DefaultResourceID)
}

//Can can?
func Can(user uint, name string, rty string, rid uint) bool {
	role, err := getRole(name, rty, rid)
	if err != nil {
		beego.Error(err)
		return false
	}
	o := orm.NewOrm()
	var pm Permission
	if err := o.QueryTable(&pm).
		Filter("user_id", user).
		Filter("role_id", role.ID).
		One(&pm); err != nil {
		return false
	}

	return pm.Enable()
}

func getRole(name string, rty string, rid uint) (*Role, error) {
	r := Role{}
	o := orm.NewOrm()
	err := o.QueryTable(&r).
		Filter("name", name).
		Filter("resource_type", rty).Filter("resource_id", rid).One(&r)
	if err == nil {
		return &r, nil
	}
	if err == orm.ErrNoRows {
		r.Name = name
		r.ResourceID = rid
		r.ResourceType = rty
		_, err = o.Insert(&r)
	}
	return &r, err
}

//Deny deny permission
func Deny(user uint, name, rty string, rid uint) {
	role, err := getRole(name, rty, rid)
	if err == nil {
		_, err = orm.NewOrm().
			QueryTable(new(Permission)).
			Filter("role_id", role.ID).
			Filter("user_id", user).
			Delete()
	}
	if err != nil {
		beego.Error(err)
	}
}

//Allow allow permission
func Allow(user uint, name, rty string, rid uint, years, months, days int) {
	role, err := getRole(name, rty, rid)
	if err == nil {

		begin := time.Now()
		end := begin.AddDate(years, months, days)
		var pm Permission
		o := orm.NewOrm()
		err = o.QueryTable(&pm).
			Filter("role_id", role.ID).
			Filter("user_id", user).
			One(&pm)
		if err == nil {
			pm.StartUp = begin
			pm.ShutDown = end
			_, err = o.Update(&pm, "start_up", "shut_down", "updated_at")

		} else if err == orm.ErrNoRows {
			pm.UserID = user
			pm.RoleID = role.ID
			pm.StartUp = begin
			pm.ShutDown = end
			_, err = o.Insert(&pm)
		}
	}

	if err != nil {
		beego.Error(err)
	}
}
