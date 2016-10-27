package reading

import (
	"github.com/astaxie/beego/orm"
	"github.com/kapmahc/lotus/engines/base"
)

//Book book
type Book struct {
	base.Model
	Author      string
	Publisher   string
	Title       string
	Type        string
	Lang        string
	File        string
	Subject     string
	Description string
	Vote        int
	PublishedAt string
}

//TableName table name
func (p *Book) TableName() string {
	return "reading_books"
}

//Note note
type Note struct {
	base.Model
	UserID uint `orm:"column(user_id)"`
	BookID uint `orm:"column(book_id)"`
	Body   string
	Type   string
	Vote   int
}

//TableName table name
func (p *Note) TableName() string {
	return "reading_notes"
}

//------------------------------------------------------------------------------

func init() {
	orm.RegisterModel(
		new(Book),
		new(Note),
	)
}
