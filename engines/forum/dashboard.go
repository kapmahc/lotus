package forum

import (
	"github.com/astaxie/beego/orm"
	"github.com/kapmahc/lotus/engines/auth"
	"github.com/kapmahc/lotus/engines/base"
)

//GetMyArticles my articles
// @router /my-articles [get]
func (p *Controller) GetMyArticles() {
	p.Dashboard()
	var articles []Article
	o := orm.NewOrm().QueryTable(new(Article))
	if !p.IsAdmin() {
		o = o.Filter("user_id", p.CurrentUser().ID)
	}
	_, err := o.OrderBy("-updated_at").All(&articles, "id", "title", "updated_at")
	p.Check(err)
	p.Data["title"] = p.T("forum-pages.my-articles")
	p.Data["articles"] = articles
	p.TplName = "forum/articles/my.html"
}

//GetMyComments my comments
// @router /my-comments [get]
func (p *Controller) GetMyComments() {
	p.Dashboard()
	var comments []Comment
	o := orm.NewOrm().QueryTable(new(Comment))
	if !p.IsAdmin() {
		o = o.Filter("user_id", p.CurrentUser().ID)
	}
	_, err := o.OrderBy("-updated_at").All(&comments, "id", "body", "updated_at")
	p.Check(err)
	p.Data["title"] = p.T("forum-pages.my-comments")
	p.Data["comments"] = comments
	p.TplName = "forum/comments/my.html"
}

func init() {
	auth.Register(func(user *auth.User) *base.Dropdown {
		nb := base.Dropdown{
			ID:    "forum",
			Label: "forum-pages.profile",
			Links: []base.Link{
				{
					Href:  "forum.Controller.GetMyArticles",
					Label: "forum-pages.my-articles",
				},
				{
					Href:  "forum.Controller.GetMyComments",
					Label: "forum-pages.my-comments",
				},
			},
		}
		return &nb
	})
}
