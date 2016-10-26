package site

import (
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/kapmahc/lotus/models"
)

//LeaveWord leave word
type LeaveWord struct {
	ID        uint      `json:"id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

//Notice notice
type Notice struct {
	models.Base

	Lang    string `json:"lang"`
	Content string `json:"content"`
}

//Setting setting
type Setting struct {
	models.Base

	Key  string
	Val  string
	Flag bool
}

//Attachment attachment
type Attachment struct {
	models.Base

	Title     string
	Name      string
	MediaType string
	Summary   string

	UserID uint
}

//Locale locale mdoe
type Locale struct {
	models.Base

	Code    string
	Lang    string
	Message string
}

//Page page model
type Page struct {
	models.Base

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
