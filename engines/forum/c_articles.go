package forum

import (
	"strconv"

	"github.com/astaxie/beego/orm"
	"github.com/kapmahc/lotus/engines/auth"
	"github.com/kapmahc/lotus/engines/base"
)

//IndexArticle index articles
// @router /articles [get]
func (p *Controller) IndexArticle() {
	o := orm.NewOrm()

	const size = 60
	page, _ := p.GetInt64("page", 1)
	count, err := o.QueryTable(new(Article)).Count()
	p.Check(err)

	var articles []Article
	_, err = o.QueryTable(new(Article)).
		OrderBy("-vote").
		Offset((page-1)*size).
		Limit(size).
		All(&articles, "id", "title", "summary")
	p.Check(err)

	p.Data["pager"] = base.NewPaginator(
		p.URLFor("forum.Controller.IndexArticle"),
		count,
		page,
		size,
		articles,
	)
	p.Data["articles"] = articles
	p.Data["title"] = p.T("forum-pages.articles")
	p.Data["can"] = p.IsAdmin()
	p.TplName = "forum/articles/index.html"
}

//NewArticle new article
// @router /articles/new [get]
func (p *Controller) NewArticle() {
	p.MustSignIn()
	title := p.T("forum-pages.new-article")

	var tags []Tag
	_, err := orm.NewOrm().QueryTable(new(Tag)).All(&tags, "id", "name")
	p.Check(err)

	var options []base.Option
	for _, t := range tags {
		options = append(options, base.Option{
			Value: t.ID,
			Name:  t.Name,
		})
	}

	p.Data["title"] = title
	p.Data["form"] = p.NewForm(
		"fm-new-article",
		title,
		base.MethodPost,
		p.URLFor("forum.Controller.CreateArticle"),
		[]base.Field{
			&base.TextField{
				ID:    "title",
				Label: p.T("attributes.title"),
			},
			&base.Textarea{
				ID:    "summary",
				Label: p.T("attributes.summary"),
			},
			&base.CheckBox{
				ID:      "tags",
				Label:   p.T("forum-attributes.article-tags"),
				Options: options,
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

//CreateArticle create article
// @router /articles [post]
func (p *Controller) CreateArticle() {
	p.MustSignIn()
	var fm fmArticle
	fl, er := p.ParseForm(&fm)

	if er == nil {
		user := p.CurrentUser()
		article := Article{
			Title:   fm.Title,
			Summary: fm.Summary,
			Body:    fm.Body,
			User:    user,
		}
		o := orm.NewOrm()
		_, err := o.Insert(&article)
		p.Check(err)

		var tags []interface{}
		for _, t := range fm.Tags {
			id, err := strconv.Atoi(t)
			p.Check(err)
			var tag Tag
			tag.ID = uint(id)
			tags = append(tags, &tag)
		}
		if len(tags) > 0 {
			_, err := o.QueryM2M(&article, "Tags").Add(tags...)
			p.Check(err)
		}
		fl.Notice(p.T("site-pages.success"))
		p.Redirect(fl, "forum.Controller.ShowArticle", ":id", article.ID)
	} else {
		fl.Error(er.Error())
		p.Redirect(fl, "forum.Controller.NewArticle")
	}
}

func (p *Controller) canArticle() (Article, bool) {
	p.MustSignIn()
	var article Article
	err := orm.NewOrm().
		QueryTable(&article).
		Filter("id", p.Ctx.Input.Param(":id")).
		One(&article)
	p.Check(err)
	user := p.CurrentUser()
	return article, (article.User.ID == user.ID || user.Has(auth.AdminRole))

}

//EditArticle edit article
// @router /articles/:id/edit [get]
func (p *Controller) EditArticle() {
	article, can := p.canArticle()
	if !can {
		p.Abort("403")
	}

	o := orm.NewOrm()
	var tags []Tag
	_, err := o.QueryTable(new(Tag)).All(&tags, "id", "name")
	p.Check(err)
	var options []base.Option
	for _, t := range tags {
		options = append(options, base.Option{
			Value:    t.ID,
			Name:     t.Name,
			Selected: o.QueryM2M(&article, "Tags").Exist(&t),
		})
	}

	title := p.T("forum-pages.edit-article", article.ID)
	p.Data["title"] = title
	p.Data["form"] = p.NewForm(
		"fm-edit-article",
		title,
		base.MethodPost,
		p.URLFor("forum.Controller.UpdateArticle", ":id", article.ID),
		[]base.Field{
			&base.TextField{
				ID:    "title",
				Label: p.T("attributes.title"),
				Value: article.Title,
			},
			&base.Textarea{
				ID:    "summary",
				Label: p.T("attributes.summary"),
				Value: article.Summary,
			},
			&base.CheckBox{
				ID:      "tags",
				Label:   p.T("forum-attributes.article-tags"),
				Options: options,
			},
			&base.Textarea{
				ID:     "body",
				Label:  p.T("attributes.body"),
				Value:  article.Body,
				Helper: p.T("site-pages.can-markdown"),
			},
		},
	)
	p.TplName = "auth/form.html"
}

//UpdateArticle update article
// @router /articles/:id [post]
func (p *Controller) UpdateArticle() {
	article, can := p.canArticle()
	if !can {
		p.Abort("403")
	}

	var fm fmArticle
	fl, er := p.ParseForm(&fm)

	if er == nil {
		article.Title = fm.Title
		article.Summary = fm.Summary
		article.Body = fm.Body

		o := orm.NewOrm()
		_, err := o.Update(&article, "updated_at", "title", "summary", "body")
		p.Check(err)
		m2m := o.QueryM2M(&article, "Tags")
		_, err = m2m.Clear()
		p.Check(err)
		var tags []interface{}
		for _, t := range fm.Tags {
			id, err := strconv.Atoi(t)
			p.Check(err)
			var tag Tag
			tag.ID = uint(id)
			tags = append(tags, &tag)
		}
		if len(tags) > 0 {
			m2m.Add(tags...)
		}

		fl.Notice(p.T("site-pages.success"))
		p.Redirect(fl, "forum.Controller.ShowArticle", ":id", article.ID)
	} else {
		fl.Error(er.Error())
		p.Redirect(fl, "forum.Controller.EditArticle", ":id", article.ID)
	}
}

//ShowArticle show article
// @router /articles/:id [get]
func (p *Controller) ShowArticle() {
	article, can := p.canArticle()

	o := orm.NewOrm()
	_, err := o.LoadRelated(&article, "Tags", "id", "name")
	p.Check(err)
	_, err = o.LoadRelated(&article, "Comments")
	// _, err := o.QueryTable(new(Comment)).
	// 	Filter("Article", article.ID).
	// 	OrderBy("-updated_at").
	// 	All(&article.Comments)
	p.Check(err)

	p.Data["article"] = article
	p.Data["title"] = article.Title
	p.Data["can"] = can

	p.Data["form"] = p.NewForm(
		"fm-new-comment",
		p.T("forum-pages.new-comment"),
		base.MethodPost,
		p.URLFor("forum.Controller.CreateComment", ":id", article.ID),
		[]base.Field{
			&base.HiddenField{
				ID:    "article_id",
				Value: article.ID,
			},
			&base.Textarea{
				ID:     "body",
				Label:  p.T("attributes.body"),
				Helper: p.T("site-pages.can-markdown"),
			},
		},
	)
	p.TplName = "forum/articles/show.html"
}

//DestroyArticle destroy article
// @router /articles/:id [delete]
func (p *Controller) DestroyArticle() {
	article, can := p.canArticle()
	if !can {
		p.Abort("403")
	}
	_, err := orm.NewOrm().Delete(&article)
	p.Check(err)

	p.Data["json"] = map[string]string{
		"to": p.URLFor("forum.Controller.IndexArticle"),
	}
	p.ServeJSON()
}
