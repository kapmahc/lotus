package vpn

import (
	"time"

	"github.com/kapmahc/lotus/web"
)

// http://chagridsada.blogspot.com/2011/01/openvpn-system-based-on-userpass.html

//User user
type User struct {
	web.Model

	Email    string    `json:"email"`
	Password string    `json:"-"`
	Details  string    `json:"details"`
	Online   bool      `json:"online"`
	Enable   bool      `json:"enable"`
	StartUp  time.Time `json:"start_up"`
	ShutDown time.Time `json:"shut_down"`

	Logs []Log `json:"logs"`
}

//TableName table name
func (p *User) TableName() string {
	return "vpn_users"
}

//Log log
type Log struct {
	ID          uint      `json:"id"`
	TrustedIP   string    `json:"trusted_ip"`
	TrustedPort int       `json:"trusted_port"`
	RemoteIP    string    `json:"remote_ip"`
	RemotePort  int       `json:"remote_port"`
	StartUp     time.Time `json:"start_up"`
	ShutDown    time.Time `json:"shut_down"`
	Received    float32   `json:"received"`
	Send        float32   `json:"send"`

	UserID uint `json:"user_id"`
	User   User `json:"user"`
}

//TableName table name
func (p *Log) TableName() string {
	return "vpn_logs"
}
