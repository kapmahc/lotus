package forum

import (
	"github.com/astaxie/beego/orm"
	"github.com/kapmahc/lotus/engines/base"
)

//IndexTag index tags
// @router /tags [get]
func (p *Controller) IndexTag() {
	var tags []Tag
	_, err := orm.NewOrm().QueryTable(new(Tag)).All(&tags, "id", "name")
	p.Check(err)
	p.Data["tags"] = tags
	p.Data["title"] = p.T("forum-pages.tags")
	p.Data["can"] = p.IsAdmin()
	p.TplName = "forum/tags/index.html"
}

//NewTag new tag
// @router /tags/new [get]
func (p *Controller) NewTag() {
	p.MustAdmin()
	title := p.T("forum-pages.new-tag")
	p.Data["title"] = title
	p.Data["form"] = p.NewForm(
		"fm-new-tag",
		title,
		base.MethodPost,
		p.URLFor("forum.Controller.CreateTag"),
		[]base.Field{
			&base.TextField{
				ID:    "name",
				Label: p.T("attributes.name"),
			},
		},
	)
	p.TplName = "auth/form.html"

}

//CreateTag create tag
// @router /tags [post]
func (p *Controller) CreateTag() {
	p.MustAdmin()
	var fm fmTag
	fl, er := p.ParseForm(&fm)
	o := orm.NewOrm()
	if er == nil {
		count, err := o.QueryTable(new(Tag)).Filter("name", fm.Name).Count()
		p.Check(err)
		if count > 0 {
			er = p.Error("forum-logs.tag-name-already-exists")
		}
	}
	if er == nil {
		_, err := o.Insert(&Tag{Name: fm.Name})
		p.Check(err)
		fl.Notice(p.T("site-pages.success"))
		p.Redirect(fl, "forum.Controller.IndexTag")
	} else {
		fl.Error(er.Error())
		p.Redirect(fl, "forum.Controller.NewTag")
	}
}

//EditTag edit tag
// @router /tags/:id/edit [get]
func (p *Controller) EditTag() {
	p.MustAdmin()
	var tag Tag
	err := orm.NewOrm().
		QueryTable(&tag).
		Filter("id", p.Ctx.Input.Param(":id")).
		One(&tag)
	p.Check(err)

	title := p.T("forum-pages.edit-tag", tag.ID)
	p.Data["title"] = title
	p.Data["form"] = p.NewForm(
		"fm-edit-tag",
		title,
		base.MethodPost,
		p.URLFor("forum.Controller.UpdateTag", ":id", tag.ID),
		[]base.Field{
			&base.TextField{
				ID:    "name",
				Label: p.T("attributes.name"),
				Value: tag.Name,
			},
		},
	)
	p.TplName = "auth/form.html"
}

//UpdateTag update tag
// @router /tags/:id [post]
func (p *Controller) UpdateTag() {
	p.MustAdmin()
	var tag Tag
	o := orm.NewOrm()
	err := o.QueryTable(&tag).
		Filter("id", p.Ctx.Input.Param(":id")).
		One(&tag)
	p.Check(err)

	var fm fmTag
	fl, er := p.ParseForm(&fm)
	if er == nil {
		var count int64
		count, err = o.QueryTable(new(Tag)).Filter("name", fm.Name).Count()
		p.Check(err)
		if count > 0 {
			er = p.Error("forum-logs.tag-name-already-exists")
		}
	}
	if er == nil {
		tag.Name = fm.Name
		_, err = o.Update(&tag, "updated_at", "name")
		p.Check(err)
		fl.Notice(p.T("site-pages.success"))
		p.Redirect(fl, "forum.Controller.IndexTag")
	} else {
		fl.Error(er.Error())
		p.Redirect(fl, "forum.Controller.EditTag", ":id", tag.ID)
	}
}

//ShowTag show tag
// @router /tags/:id [get]
func (p *Controller) ShowTag() {
	var tag Tag
	err := orm.NewOrm().
		QueryTable(new(Tag)).
		Filter("id", p.Ctx.Input.Param(":id")).
		One(&tag, "id", "name")
	p.Check(err)

	var articles []Article
	_, err = orm.NewOrm().QueryTable(new(Article)).All(&articles, "id", "title", "summary", "type")
	p.Check(err)
	p.Data["tag"] = tag
	p.Data["title"] = p.T("forum-pages.show-tag")
	p.Data["articles"] = articles
	p.TplName = "forum/articles/index.html"
}

//DestroyTag destroy tag
// @router /tags/:id [delete]
func (p *Controller) DestroyTag() {
	p.MustAdmin()
	id := p.Ctx.Input.Param(":id")
	o := orm.NewOrm()
	_, err := o.
		QueryTable(new(ArticleTag)).
		Filter("tag_id", id).
		Delete()
	p.Check(err)
	_, err = o.
		QueryTable(new(Tag)).
		Filter("id", id).
		Delete()
	p.Check(err)

	p.Data["json"] = map[string]string{
		"to": p.URLFor("forum.Controller.IndexTag"),
	}
	p.ServeJSON()
}
