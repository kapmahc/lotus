package auth

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"
	"time"

	"github.com/kapmahc/lotus/web"
	uuid "github.com/satori/go.uuid"
)

const (
	//ProvideTypeEmail email provuder-type
	ProvideTypeEmail = "email"
)

//User user model
type User struct {
	web.Model

	Email    string `json:"email"`
	UID      string `json:"uid"`
	Home     string `json:"home"`
	Logo     string `json:"logo"`
	Name     string `json:"name"`
	Password string `json:"-"`

	ProviderType string `json:"provider_type"`
	ProviderID   string `json:"provider_id"`

	LastSignInAt    *time.Time `json:"last_sign_in_at"`
	LastSignInIP    string     `json:"last_sign_in_ip"`
	CurrentSignInAt *time.Time `json:"current_sign_in_at"`
	CurrentSignInIP string     `json:"current_sign_in_ip"`
	SignInCount     uint       `json:"sign_in_count"`
	ConfirmedAt     *time.Time `json:"confirmed_at"`
	LockedAt        *time.Time `json:"locked_at"`

	Logs        []Log        `json:"logs"`
	Permissions []Permission `json:"permissions"`
}

//TableName table name
func (p *User) TableName() string {
	return "users"
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

//------------------------------------------------------------------------------

//Log model
type Log struct {
	ID        uint      `json:"id"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`

	UserID uint `json:"user_id"`
	User   User `json:"user"`
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

	//RoleAdmin admin
	RoleAdmin = "admin"
	//RoleRoot root
	RoleRoot = "root"
)

//Role role model
type Role struct {
	web.Model

	Name         string `json:"name"`
	ResourceType string `json:"resource_type"`
	ResourceID   uint   `json:"resource_id"`

	Permissions []Permission `json:"permissions"`
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
	web.Model

	StartUp  time.Time `json:"start_up"`
	ShutDown time.Time `json:"shut_down"`

	UserID uint `json:"user_id"`
	User   User `json:"user"`
	RoleID uint `json:"role_id"`
	Role   Role `json:"role"`
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
