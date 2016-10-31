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
	Tags    []*Tag `json:"tags" orm:"rel(m2m);rel_table(forum_articles_tags)"`
}

//TableName table name
func (p *Article) TableName() string {
	return "forum_articles"
}

//Tag tag
type Tag struct {
	base.Model
	Name     string     `json:"name"`
	Vote     int        `json:"vote"`
	Articles []*Article `json:"articles" orm:"reverse(many)"`
}

//TableName table name
func (p *Tag) TableName() string {
	return "forum_tags"
}

//Comment comment
type Comment struct {
	base.Model
	UserID uint   `json:"user_id" orm:"column(user_id)"`
	Body   string `json:"body"`
	Type   string `json:"type"`
	Vote   int    `json:"vote"`
}

//TableName table name
func (p *Comment) TableName() string {
	return "forum_comments"
}

func init() {
	orm.RegisterModel(
		new(Article),
		new(Tag),
		new(Comment),
	)

}
