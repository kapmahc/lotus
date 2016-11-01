package forum

import (
	"github.com/astaxie/beego/orm"
	"github.com/kapmahc/lotus/engines/auth"
	"github.com/kapmahc/lotus/engines/base"
)

//NewComment new comment
// @router /comments/new [get]
func (p *Controller) NewComment() {
	p.MustSignIn()
	title := p.T("forum-pages.new-comment")
	p.Data["title"] = title
	p.Data["form"] = p.NewForm(
		"fm-new-comment",
		title,
		base.MethodPost,
		p.URLFor("forum.Controller.CreateComment"),
		[]base.Field{
			&base.HiddenField{
				ID:    "article_id",
				Value: p.GetString("article_id"),
			},
			&base.Textarea{
				ID:     "body",
				Label:  p.T("attributes.body"),
				Helper: p.T("site-pages.can-markdown"),
			},
		},
	)
	p.TplName = "auth/form.html"

}

//CreateComment create comment
// @router /comments [post]
func (p *Controller) CreateComment() {
	p.MustSignIn()
	var fm fmComment
	fl, er := p.ParseForm(&fm)
	aid := p.GetString("article_id")

	if er == nil {
		var article Article
		o := orm.NewOrm()
		err := o.QueryTable(&article).Filter("id", aid).One(&article)
		p.Check(err)
		_, err = o.Insert(&Comment{
			Article: &article,
			Body:    fm.Body,
			User:    p.CurrentUser(),
		})
		p.Check(err)
		fl.Notice(p.T("site-pages.success"))
		p.Redirect(fl, "forum.Controller.ShowArticle", ":id", aid)
	} else {
		fl.Error(er.Error())
		p.Redirect(fl, "forum.Controller.NewComment", "article_id", aid)
	}
}

func (p *Controller) canComment() (Comment, bool) {
	var comment Comment
	err := orm.NewOrm().
		QueryTable(&comment).
		Filter("id", p.Ctx.Input.Param(":id")).
		One(&comment)
	p.Check(err)

	return comment, (p.IsSignIn() && (comment.User.ID == p.CurrentUser().ID || p.CurrentUser().Has(auth.AdminRole)))
}

//EditComment edit comment
// @router /comments/:id/edit [get]
func (p *Controller) EditComment() {
	comment, can := p.canComment()
	if !can {
		p.Abort("403")
	}

	title := p.T("forum-pages.edit-comment", comment.ID)
	p.Data["title"] = title
	p.Data["form"] = p.NewForm(
		"fm-edit-comment",
		title,
		base.MethodPost,
		p.URLFor("forum.Controller.UpdateComment", ":id", comment.ID),
		[]base.Field{
			&base.HiddenField{
				ID:    "article_id",
				Value: comment.Article.ID,
			},
			&base.Textarea{
				ID:     "body",
				Label:  p.T("attributes.body"),
				Value:  comment.Body,
				Helper: p.T("site-pages.can-markdown"),
			},
		},
	)
	p.TplName = "auth/form.html"
}

//UpdateComment update comment
// @router /comments/:id [post]
func (p *Controller) UpdateComment() {
	comment, can := p.canComment()
	if !can {
		p.Abort("403")
	}

	var fm fmComment
	fl, er := p.ParseForm(&fm)

	if er == nil {
		comment.Body = fm.Body
		_, err := orm.NewOrm().Update(&comment, "updated_at", "body")
		p.Check(err)
		fl.Notice(p.T("site-pages.success"))
		p.Redirect(fl, "forum.Controller.ShowArticle", ":id", comment.Article.ID)
	} else {
		fl.Error(er.Error())
		p.Redirect(fl, "forum.Controller.EditComment", ":id", comment.ID)
	}
}

//DestroyComment destroy comment
// @router /comments/:id [delete]
func (p *Controller) DestroyComment() {
	comment, can := p.canComment()
	if !can {
		p.Abort("403")
	}
	_, err := orm.NewOrm().Delete(&comment)
	p.Check(err)

	p.Data["json"] = map[string]string{
		"to": p.URLFor("forum.Controller.ShowArticle", ":id", comment.Article.ID),
	}
	p.ServeJSON()
}
