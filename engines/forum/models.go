package forum

import (
	"github.com/astaxie/beego/orm"
	"github.com/kapmahc/lotus/engines/base"
)

const (
	//TypeMarkdown markdown
	TypeMarkdown = "markdown"
	//TypeHTML html
	TypeHTML = "html"
)

//Article article
type Article struct {
	base.Model
	UserID  uint   `json:"user_id" orm:"column(user_id)"`
	Title   string `json:"title"`
	Summary string `json:"summary"`
	Body    string `json:"body"`
	Type    string `json:"type"`
	Vote    int    `json:"vote"`
}

//TableName table name
func (p *Article) TableName() string {
	return "forum_articles"
}

//Tag tag
type Tag struct {
	base.Model
	Name string `json:"name"`
	Vote int    `json:"vote"`
}

//TableName table name
func (p *Tag) TableName() string {
	return "forum_tags"
}

//ArticleTag article-tag-rel
type ArticleTag struct {
	ID        uint `json:"id" orm:"column(id)"`
	ArticleID uint `json:"article_id" orm:"column(article_id)"`
	TagID     uint `json:"tag_id" orm:"column(tag_id)"`
}

//TableName table name
func (p *ArticleTag) TableName() string {
	return "forum_articles_tags"
}

//Comment comment
type Comment struct {
	base.Model
	UserID    uint   `json:"user_id" orm:"column(user_id)"`
	ArticleID uint   `json:"user_id" orm:"column(article_id)"`
	Body      string `json:"body"`
	Type      string `json:"type"`
	Vote      int    `json:"vote"`
}

//TableName table name
func (p *Comment) TableName() string {
	return "forum_comments"
}

func init() {
	orm.RegisterModel(
		new(Article),
		new(Tag),
		new(ArticleTag),
		new(Comment),
	)

}
