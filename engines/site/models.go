package site

import (
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/kapmahc/lotus/engines/base"
)

//LeaveWord leave word
type LeaveWord struct {
	ID        uint      `json:"id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
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

//Setting setting
type Setting struct {
	base.Model

	Key  string
	Val  string
	Flag bool
}

//TableName table name
func (p *Setting) TableName() string {
	return "settings"
}

//Attachment attachment
type Attachment struct {
	base.Model

	Title     string
	Name      string
	MediaType string
	Summary   string

	UserID uint
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
		new(Setting),
		new(LeaveWord),
		new(Notice),
		new(Attachment),
		new(Page),
	)
}
