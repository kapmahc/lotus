package forum

import (
	"github.com/kapmahc/lotus/engines/auth"
	"github.com/kapmahc/lotus/web"
)

const (
	//TypeMarkdown markdown
	TypeMarkdown = "markdown"
	//TypeHTML html
	TypeHTML = "html"
)

//Article article
type Article struct {
	web.Model

	Title   string `json:"title"`
	Summary string `json:"summary"`
	Body    string `json:"body"`
	Type    string `json:"type"`
	Vote    int    `json:"vote"`

	UserID   uint      `json:"user_id"`
	User     auth.User `json:"user"`
	Tags     []Tag     `json:"tags" gorm:"many2many:forum_articles_tags;"`
	Comments []Comment `json:"comments"`
}

//TableName table name
func (p *Article) TableName() string {
	return "forum_articles"
}

//Tag tag
type Tag struct {
	web.Model

	Name string `json:"name"`
	Vote int    `json:"vote"`

	Articles []Article `json:"articles" gorm:"many2many:forum_articles_tags;"`
}

//TableName table name
func (p *Tag) TableName() string {
	return "forum_tags"
}

//Comment comment
type Comment struct {
	web.Model

	Body string `json:"body"`
	Type string `json:"type"`
	Vote int    `json:"vote"`

	UserID    uint      `json:"user_id"`
	User      auth.User `json:"user"`
	ArticleID uint      `json:"article_id"`
	Article   Article   `json:"article"`
}

//TableName table name
func (p *Comment) TableName() string {
	return "forum_comments"
}
