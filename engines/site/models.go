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

//Notice notice
type Notice struct {
	base.Model

	Lang    string `json:"lang"`
	Content string `json:"content"`
}

//Setting setting
type Setting struct {
	base.Model

	Key  string
	Val  string
	Flag bool
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

//Locale locale mdoe
type Locale struct {
	base.Model

	Code    string
	Lang    string
	Message string
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

func init() {
	orm.RegisterModel(
		new(Setting),
		new(Locale),
		new(LeaveWord),
		new(Notice),
		new(Attachment),
		new(Page),
	)
}
