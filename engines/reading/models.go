package reading

import (
	"github.com/kapmahc/lotus/engines/auth"
	"github.com/kapmahc/lotus/web"
)

//Book book
type Book struct {
	web.Model

	Author      string `json:"author"`
	Publisher   string `json:"publisher"`
	Title       string `json:"title"`
	Type        string `json:"type"`
	Lang        string `json:"lang"`
	File        string `json:"file"`
	Subject     string `json:"subject"`
	Description string `json:"description"`
	Vote        int    `json:"vote"`
	PublishedAt string `json:"publish_at"`

	Notes []Note `json:"notes"`
}

//TableName table name
func (p *Book) TableName() string {
	return "reading_books"
}

//Note note
type Note struct {
	web.Model

	Body string `json:"body"`
	Type string `json:"type"`
	Vote int    `json:"vote"`

	UserID uint      `json:"user_id"`
	User   auth.User `json:"user"`
	BookID uint      `json:"book_id"`
	Book   Book      `json:"book"`
}

//TableName table name
func (p *Note) TableName() string {
	return "reading_notes"
}
