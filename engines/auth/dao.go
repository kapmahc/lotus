package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/SermoDigital/jose/jws"
	"github.com/jinzhu/gorm"
	"github.com/kapmahc/lotus/web"
	"github.com/spf13/viper"
)

//Dao dao
type Dao struct {
	Db     *gorm.DB    `inject:""`
	Logger *web.Logger `inject:""`
	Cache  *web.Cache  `inject:""`
	Aes    *web.Aes    `inject:""`
	Hmac   *web.Hmac   `inject:""`
	I18n   *web.I18n   `inject:""`
}

//Set set key-val
func (p *Dao) Set(key string, val interface{}, flag bool) error {
	buf, err := web.ToBytes(val)
	if err != nil {
		return err
	}
	if flag {
		if buf, err = p.Aes.Encrypt(buf); err != nil {
			return err
		}
	}
	var m Setting
	if p.Db.Where("key = ?", key).First(&m).RecordNotFound() {
		m.Key = key
		m.Val = buf
		m.Flag = flag
		err = p.Db.Create(&m).Error
	} else {
		m.Val = buf
		m.Flag = flag
		err = p.Db.Update(&m).Error
	}
	return err
}

//Get get val
func (p *Dao) Get(key string, val interface{}) error {
	var m Setting
	err := p.Db.Where("key = ?", key).First(&m).Error
	if err != nil {
		return err
	}
	if m.Flag {
		if m.Val, err = p.Aes.Decrypt(m.Val); err != nil {
			return err
		}
	}
	return web.FromBytes(m.Val, val)
}

//-----------------------------------------------------------------------------

//Log add log
func (p *Dao) Log(user uint, msg string) {
	l := Log{UserID: user, Message: msg}
	if err := p.Db.Create(&l).Error; err != nil {
		p.Logger.Error(err.Error())
	}
}

//-----------------------------------------------------------------------------

//UserClaims generate user claims
func (p *Dao) UserClaims(u *User) jws.Claims {
	cm := jws.Claims{}
	cm.SetSubject(u.Name)
	cm.Set("uid", u.UID)

	cm.Set("roles", p.Authority(u.ID, DefaultResourceType, DefaultResourceID))
	return cm
}

//AddEmailUser add email user
func (p *Dao) AddEmailUser(lang, email, name, password string) (*User, error) {
	var u User
	var err error
	var count int
	if err = p.Db.Model(&u).Where(
		"(provider_id = ? AND provider_type = ?) OR email = ?",
		email,
		ProvideTypeEmail,
		email).
		Count(&count).Error; err != nil {
		return nil, err
	}
	if count != 0 {
		return nil, errors.New(p.I18n.T(lang, "auth.messages.email-already-exists"))
	}

	u.Email = email
	u.Name = name
	u.ProviderID = email
	u.ProviderType = ProvideTypeEmail
	u.LastSignInIP = "0.0.0.0"
	u.CurrentSignInIP = "0.0.0.0"
	u.Password = p.Hmac.Sum([]byte(password))
	u.SetGravatarLogo()
	u.SetUID()
	u.Home = fmt.Sprintf("%s/users/%s", viper.GetString("server.frontend"), u.UID)
	err = p.Db.Create(&u).Error

	return &u, err
}

//AddOpenIDUser add openid user
func (p *Dao) AddOpenIDUser(pid, pty, email, name, home, logo string) (*User, error) {
	var u User
	var err error
	now := time.Now()
	if p.Db.Where("provider_id = ? AND provider_type = ?", pid, pty).First(&u).RecordNotFound() {
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
		err = p.Db.Create(&u).Error
	} else {
		err = p.Db.Model(&u).Updates(map[string]interface{}{
			"email":         email,
			"name":          name,
			"logo":          logo,
			"home":          home,
			"sign_in_count": u.SignInCount + 1,
			"last_sign_in":  &now,
		}).Error
	}
	return &u, err
}

//GetUserByUID get user by uid
func (p *Dao) GetUserByUID(uid string) (*User, error) {
	var u User
	err := p.Db.Where("uid = ?", uid).First(&u).Error
	return &u, err
}

//GetUserByEmail get user by email
func (p *Dao) GetUserByEmail(email string) (*User, error) {
	var u User
	err := p.Db.Where("provider_type = ? AND provider_id = ?", ProvideTypeEmail, email).First(&u).Error
	return &u, err
}

//Authority get user's role names
func (p *Dao) Authority(user uint, rty string, rid uint) []string {
	var items []Role
	if err := p.Db.
		Where("resource_type = ? AND resource_id = ?", rty, rid).
		Find(&items).Error; err != nil {
		p.Logger.Error(err.Error())
	}
	var roles []string
	for _, r := range items {
		var pm Permission
		if err := p.Db.
			Where("role_id = ? AND user_id = ? ", r.ID, user).
			First(&pm).Error; err != nil {
			p.Logger.Error(err.Error())
			continue
		}
		if pm.Enable() {
			roles = append(roles, r.Name)
		}
	}

	return roles
}

//Is is role ?
func (p *Dao) Is(user uint, name string) bool {
	return p.Can(user, name, DefaultResourceType, DefaultResourceID)
}

//Can can?
func (p *Dao) Can(user uint, name string, rty string, rid uint) bool {
	var r Role
	if p.Db.
		Where("name = ? AND resource_type = ? AND resource_id = ?", name, rty, rid).
		First(&r).
		RecordNotFound() {
		return false
	}
	var pm Permission
	if p.Db.
		Where("user_id = ? AND role_id = ?", user, r.ID).
		First(&pm).
		RecordNotFound() {
		return false
	}

	return pm.Enable()
}

//Role check role exist
func (p *Dao) Role(name string, rty string, rid uint) (*Role, error) {
	var e error
	r := Role{}
	db := p.Db
	if db.
		Where("name = ? AND resource_type = ? AND resource_id = ?", name, rty, rid).
		First(&r).
		RecordNotFound() {
		r = Role{
			Name:         name,
			ResourceType: rty,
			ResourceID:   rid,
		}
		e = db.Create(&r).Error

	}
	return &r, e
}

//Deny deny permission
func (p *Dao) Deny(role uint, user uint) error {
	return p.Db.
		Where("role_id = ? AND user_id = ?", role, user).
		Delete(Permission{}).Error
}

//Allow allow permission
func (p *Dao) Allow(role uint, user uint, years, months, days int) error {
	begin := time.Now()
	end := begin.AddDate(years, months, days)
	var count int
	p.Db.
		Model(&Permission{}).
		Where("role_id = ? AND user_id = ?", role, user).
		Count(&count)
	if count == 0 {
		return p.Db.Create(&Permission{
			UserID:   user,
			RoleID:   role,
			StartUp:  begin,
			ShutDown: end,
		}).Error
	}
	return p.Db.
		Model(&Permission{}).
		Where("role_id = ? AND user_id = ?", role, user).
		UpdateColumns(map[string]interface{}{"begin": begin, "end": end}).Error

}
