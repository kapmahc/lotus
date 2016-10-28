package auth

import (
	"fmt"
	"time"

	"github.com/SermoDigital/jose/jws"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

//GenerateUserClaims generate user claims
func GenerateUserClaims(u *User) jws.Claims {
	cm := jws.Claims{}
	cm.SetSubject(u.Name)
	cm.Set("uid", u.UID)
	cm.Set("id", u.ID)
	cm.Set("roles", u.GetRoleNames(DefaultResourceType, DefaultResourceID))
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
