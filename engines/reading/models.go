package reading

import (
	"time"

	"github.com/kapmahc/lotus/engines/base"
)

//Book book
type Book struct {
	base.Model
	UID         string `orm:"column(uid)"`
	Author      string
	Publisher   string
	Title       string
	Type        string
	Lang        string
	File        string
	Vote        int
	PublishedAt time.Time
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
