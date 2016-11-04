package mail

import "github.com/kapmahc/lotus/web"

// http://wiki2.dovecot.org/HowTo/DovecotPostgresql
// https://www.linode.com/docs/email/postfix/email-with-postfix-dovecot-and-mysql

//Domain domain
type Domain struct {
	web.Model

	Name string `json:"name"`

	Users   []User  `json:"users"`
	Aliases []Alias `json:"aliases"`
}

//TableName table name
func (p *Domain) TableName() string {
	return "mail_domains"
}

//User user
type User struct {
	web.Model

	Email    string `json:"email"`
	Password string `json:"-"`
	Name     string `json:"name"`

	DomainID uint   `json:"domain_id"`
	Domain   Domain `json:"domain"`
}

//TableName table name
func (p *User) TableName() string {
	return "mail_users"
}

//Alias alias
type Alias struct {
	web.Model

	Source      string `json:"source"`
	Destination string `json:"destination"`

	DomainID uint   `json:"domain_id"`
	Domain   Domain `json:"domain"`
}

//TableName table name
func (p *Alias) TableName() string {
	return "mail_aliases"
}
