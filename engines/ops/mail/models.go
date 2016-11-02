package mail

import (
	"crypto/sha512"
	"encoding/base64"

	"github.com/astaxie/beego/orm"
	"github.com/kapmahc/lotus/engines/base"
)

// http://wiki2.dovecot.org/HowTo/DovecotPostgresql
// https://www.linode.com/docs/email/postfix/email-with-postfix-dovecot-and-mysql

//Domain domain
type Domain struct {
	base.Model
	Name string

	Users   []*User  `orm:"reverse(many)"`
	Aliases []*Alias `orm:"reverse(many)"`
}

//TableName table name
func (p *Domain) TableName() string {
	return "mail_domains"
}

//User user
type User struct {
	base.Model
	Email    string
	Password string
	Name     string

	Domain *Domain `orm:"rel(fk)"`
}

//SetPassword set password sha512 with salt
func (p *User) SetPassword(password string, num uint) error {
	salt, err := base.RandomBytes(num)
	if err != nil {
		return err
	}
	p.Password, err = p.sumSsha512([]byte(password), salt)
	return err
}

//ChkPassword check password
func (p *User) ChkPassword(password string) (bool, error) {
	buf, err := base64.StdEncoding.DecodeString(p.Password)
	if err != nil {
		return false, err
	}

	if len(buf) <= sha512.Size {
		return false, err
	}
	salt := buf[sha512.Size:]
	rst, err := p.sumSsha512([]byte(password), salt)
	if err != nil {
		return false, err
	}
	return rst == p.Password, nil
}

func (p *User) sumSsha512(plain, salt []byte) (string, error) {
	buf := append([]byte(plain), salt...)
	code := sha512.Sum512(buf)
	return base64.StdEncoding.EncodeToString(append(code[:], salt...)), nil
	// return base64.StdEncoding.EncodeToString(append(p.Hash.Sum(append(plain, salt...)), salt...)), nil
}

//TableName table name
func (p *User) TableName() string {
	return "mail_users"
}

//Alias alias
type Alias struct {
	base.Model
	Source      string
	Destination string

	Domain *Domain `orm:"rel(fk)"`
}

//TableName table name
func (p *Alias) TableName() string {
	return "mail_aliases"
}

func init() {
	orm.RegisterModel(
		new(User),
		new(Domain),
		new(Alias),
	)
}
