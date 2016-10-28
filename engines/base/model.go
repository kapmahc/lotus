package base

import (
	"time"

	"github.com/astaxie/beego/orm"
)

//Model base model
type Model struct {
	ID        uint       `json:"id" orm:"column(id)"`
	UpdatedAt *time.Time `json:"updated_at" orm:"auto_now"`
	CreatedAt time.Time  `json:"created_at" orm:"auto_now_add"`
}

//Locale locale mdoe
type Locale struct {
	Model

	Code    string
	Lang    string
	Message string
}

//TableName table name
func (p *Locale) TableName() string {
	return "locales"
}

//Setting setting
type Setting struct {
	Model

	Key  string
	Val  string
	Flag bool
}

//TableName table name
func (p *Setting) TableName() string {
	return "settings"
}

func init() {
	orm.RegisterModel(
		new(Locale),
		new(Setting),
	)
}
