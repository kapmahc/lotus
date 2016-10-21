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

//LeaveWord leave word
type LeaveWord struct {
	ID        uint      `json:"id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

//Notice notice
type Notice struct {
	web.Model
	Lang    string `json:"lang"`
	Content string `json:"content"`
}

//Setting setting
type Setting struct {
	web.Model

	Key  string
	Val  []byte
	Flag bool
}

//User user model
type User struct {
	web.Model
	Email    string `json:"email"`
	UID      string `json:"uid"`
	Home     string `json:"home"`
	Logo     string `json:"logo"`
	Name     string `json:"name"`
	Password []byte `json:"-"`

	ProviderType string `json:"provider_type"`
	ProviderID   string `json:"provider_id"`

	LastSignInAt    *time.Time `json:"last_sign_in_at"`
	LastSignInIP    string     `json:"last_sign_in_ip"`
	CurrentSignInAt *time.Time `json:"current_sign_in_at"`
	CurrentSignInIP string     `json:"current_sign_in_ip"`
	SignInCount     uint       `json:"sign_in_count"`
	ConfirmedAt     *time.Time `json:"confirmed_at"`
	LockedAt        *time.Time `json:"locked_at"`

	Permissions []Permission `json:"permissions"`
	Logs        []Log        `json:"logs"`
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

//Log model
type Log struct {
	ID        uint      `json:"id"`
	UserID    uint      `json:"-"`
	User      User      `json:"-"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
}

//Role role model
type Role struct {
	web.Model

	Name         string
	ResourceType string
	ResourceID   uint
}

func (p Role) String() string {
	return fmt.Sprintf("%s@%s://%d", p.Name, p.ResourceType, p.ResourceID)
}

//Permission permission model
type Permission struct {
	web.Model
	User     User
	UserID   uint
	Role     Role
	RoleID   uint
	StartUp  time.Time
	ShutDown time.Time
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

//Attachment attachment
type Attachment struct {
	web.Model

	Title     string
	Name      string
	MediaType string
	Summary   string

	UserID uint
	User   User
}
