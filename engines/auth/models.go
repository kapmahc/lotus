package auth

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/kapmahc/lotus/engines/base"
	uuid "github.com/satori/go.uuid"
)

const (
	//ProvideByEmail email provuder-type
	ProvideByEmail = "email"
)

//User user model
type User struct {
	base.Model

	Email    string `json:"email"`
	UID      string `json:"uid" orm:"column(uid)"`
	Home     string `json:"home"`
	Logo     string `json:"logo"`
	Name     string `json:"name"`
	Password string `json:"-"`

	ProviderType string `json:"provider_type"`
	ProviderID   string `json:"provider_id" orm:"column(provider_id)"`

	LastSignInAt    *time.Time `json:"last_sign_in_at"`
	LastSignInIP    string     `json:"last_sign_in_ip" orm:"column(last_sign_in_ip)"`
	CurrentSignInAt *time.Time `json:"current_sign_in_at"`
	CurrentSignInIP string     `json:"current_sign_in_ip" orm:"column(current_sign_in_ip)"`
	SignInCount     uint       `json:"sign_in_count"`
	ConfirmedAt     *time.Time `json:"confirmed_at"`
	LockedAt        *time.Time `json:"locked_at"`
}

//TableName table name
func (p *User) TableName() string {
	return "users"
}

//SetPassword sum password
func (p *User) SetPassword(password string) {
	p.Password = string(p.sumHmac(password))
}

//IsPassword check passwod
func (p *User) IsPassword(password string) bool {
	return hmac.Equal(p.sumHmac(password), []byte(p.Password))
}

func (p *User) sumHmac(plain string) []byte {
	mac := hmac.New(sha512.New, []byte(beego.AppConfig.String("hmackey")))
	return mac.Sum([]byte(plain))
}

//Log add log
func (p *User) Log(msg string) {
	if _, err := orm.NewOrm().
		Insert(&Log{UserID: p.ID, Message: msg}); err != nil {
		beego.Error(err)
	}
}

//IsConfirmed confirmed?
func (p *User) IsConfirmed() bool {
	return p.ConfirmedAt != nil
}

//IsLocked locked?
func (p *User) IsLocked() bool {
	return p.LockedAt != nil
}

//IsAvailable is valid?
func (p *User) IsAvailable() bool {
	return p.IsConfirmed() && !p.IsLocked()
}

//SetGravatarLogo set logo by gravatar
func (p *User) SetGravatarLogo() {
	buf := md5.Sum([]byte(strings.ToLower(p.Email)))
	p.Logo = fmt.Sprintf(
		"https://gravatar.com/avatar/%s.png",
		hex.EncodeToString(buf[:]),
	)
}

//SetUID generate uid
func (p *User) SetUID() {
	p.UID = uuid.NewV4().String()
}

func (p User) String() string {
	return fmt.Sprintf("%s<%s>", p.Name, p.Email)
}

//Has has role ?
func (p *User) Has(name string) bool {
	return p.Can(name, DefaultResourceType, DefaultResourceID)
}

//Can can?
func (p *User) Can(name string, rty string, rid uint) bool {
	role, err := p.getRole(name, rty, rid)
	if err != nil {
		beego.Error(err)
		return false
	}
	o := orm.NewOrm()
	var pm Permission
	if err := o.QueryTable(&pm).
		Filter("user_id", p.ID).
		Filter("role_id", role.ID).
		One(&pm); err != nil {
		return false
	}

	return pm.Enable()
}

//GetRoleNames get user's role names
func (p *User) GetRoleNames(rty string, rid uint) []string {
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
			Filter("user_id", p.ID).
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

//Deny deny permission
func (p *User) Deny(name, rty string, rid uint) {
	role, err := p.getRole(name, rty, rid)
	if err == nil {
		_, err = orm.NewOrm().
			QueryTable(new(Permission)).
			Filter("role_id", role.ID).
			Filter("user_id", p.ID).
			Delete()
	}
	if err != nil {
		beego.Error(err)
	}
}

//Allow allow permission
func (p *User) Allow(name, rty string, rid uint, years, months, days int) {
	role, err := p.getRole(name, rty, rid)
	if err == nil {

		begin := time.Now()
		end := begin.AddDate(years, months, days)
		var pm Permission
		o := orm.NewOrm()
		err = o.QueryTable(&pm).
			Filter("role_id", role.ID).
			Filter("user_id", p.ID).
			One(&pm)
		if err == nil {
			pm.StartUp = begin
			pm.ShutDown = end
			_, err = o.Update(&pm, "start_up", "shut_down", "updated_at")

		} else if err == orm.ErrNoRows {
			pm.UserID = p.ID
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

func (p *User) getRole(name string, rty string, rid uint) (*Role, error) {
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

//------------------------------------------------------------------------------

//Log model
type Log struct {
	ID        uint      `json:"id" orm:"column(id)"`
	UserID    uint      `json:"-" orm:"column(user_id)"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at" orm:"auto_now_add"`
}

//TableName table name
func (p *Log) TableName() string {
	return "logs"
}

//------------------------------------------------------------------------------

const (
	//DefaultResourceType default resource-type
	DefaultResourceType = "-"
	//DefaultResourceID default resource-id
	DefaultResourceID = 0

	//AdminRole admin
	AdminRole = "admin"
	//RootRole root
	RootRole = "root"
)

//Role role model
type Role struct {
	base.Model

	Name         string
	ResourceType string
	ResourceID   uint `orm:"column(resource_id)"`
}

//TableName table name
func (p *Role) TableName() string {
	return "roles"
}

func (p Role) String() string {
	return fmt.Sprintf("%s@%s://%d", p.Name, p.ResourceType, p.ResourceID)
}

//------------------------------------------------------------------------------

//Permission permission model
type Permission struct {
	base.Model

	UserID   uint `orm:"column(user_id)"`
	RoleID   uint `orm:"column(role_id)"`
	StartUp  time.Time
	ShutDown time.Time
}

//TableName table name
func (p *Permission) TableName() string {
	return "permissions"
}

//End end to string
func (p *Permission) End() string {
	return p.ShutDown.Format("2006-01-02")
}

//Begin begin to string
func (p *Permission) Begin() string {
	return p.StartUp.Format("2006-01-02")
}

//Enable is enable?
func (p *Permission) Enable() bool {
	now := time.Now()
	return now.After(p.StartUp) && now.Before(p.ShutDown)
}

//------------------------------------------------------------------------------

func init() {
	orm.RegisterModel(
		new(User),
		new(Log),
		new(Role),
		new(Permission),
	)
}
