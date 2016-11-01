package site

import (
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/kapmahc/lotus/engines/auth"
	"github.com/kapmahc/lotus/engines/base"
)

//LeaveWord leave word
type LeaveWord struct {
	ID        uint      `json:"id" orm:"column(id)"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at" orm:"auto_now_add"`
}

//TableName table name
func (p *LeaveWord) TableName() string {
	return "leave_words"
}

//Notice notice
type Notice struct {
	base.Model

	Lang    string `json:"lang"`
	Content string `json:"content"`
}

//TableName table name
func (p *Notice) TableName() string {
	return "notices"
}

//Attachment attachment
type Attachment struct {
	base.Model

	Title     string
	Name      string
	MediaType string
	Summary   string

	User *auth.User `orm:"rel(fk)"`
}

//TableName table name
func (p *Attachment) TableName() string {
	return "attachments"
}

//Page page model
type Page struct {
	base.Model

	Loc     string
	Picture string
	Title   string
	Summary string
	Href    string
	Lang    string
}

//TableName table name
func (p *Page) TableName() string {
	return "pages"
}

func init() {
	orm.RegisterModel(
		new(LeaveWord),
		new(Notice),
		new(Attachment),
		new(Page),
	)
}
