package vpn

import (
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/kapmahc/lotus/engines/base"
)

// http://chagridsada.blogspot.com/2011/01/openvpn-system-based-on-userpass.html

//User user
type User struct {
	base.Model

	Email    string
	Password string
	Details  string
	Online   bool
	Enable   bool
	StartUp  time.Time
	ShutDown time.Time

	Logs []*Log `orm:"reverse(many)"`
}

//TableName table name
func (p *User) TableName() string {
	return "vpn_users"
}

//Log log
type Log struct {
	ID          uint   `json:"id" orm:"column(id)"`
	TrustedIP   string `json:"trusted_ip" orm:"column(trusted_ip)"`
	TrustedPort int    `json:"trusted_port"`
	RemoteIP    string `json:"remote_ip" orm:"column(remote_ip)"`
	RemotePort  int    `json:"remote_port"`
	StartUp     time.Time
	ShutDown    *time.Time
	Received    float32
	Send        float32

	User *User `orm:"rel(fk)"`
}

//TableName table name
func (p *Log) TableName() string {
	return "vpn_logs"
}

func init() {
	orm.RegisterModel(
		new(User),
		new(Log),
	)
}
