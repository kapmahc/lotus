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
	p.Logo = fmt.Sprintf("https://gravatar.com/avatar/%s.png", hex.EncodeToString(buf[:]))
}

//SetUID generate uid
func (p *User) SetUID() {
	p.UID = uuid.NewV4().String()
}

func (p User) String() string {
	return fmt.Sprintf("%s<%s>", p.Name, p.Email)
}

//------------------------------------------------------------------------------

//Log model
type Log struct {
	ID        uint      `json:"id"`
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
	DefaultResourceID = "0"
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
