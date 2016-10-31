package site

import (
	"github.com/astaxie/beego/orm"
	"github.com/kapmahc/lotus/engines/base"
)

//IndexNotice notices
// @router /notices [get]
func (p *Controller) IndexNotice() {
	var items []Notice
	_, err := orm.NewOrm().
		QueryTable(new(Notice)).
		OrderBy("-id").
		Limit(60).
		All(&items)
	p.Check(err)
	p.Data["notices"] = items
	p.Data["title"] = p.T("site-pages.notices")
	p.TplName = "site/notices/index.html"
}

//NewNotice new notice
// @router /notices/new [get]
func (p *Controller) NewNotice() {
	p.Dashboard()
	p.MustAdmin()
	title := p.T("site-pages.new-notice")
	p.Data["title"] = title
	p.Data["form"] = p.NewForm(
		"fm-new-notice",
		title,
		base.MethodPost,
		p.URLFor("site.Controller.CreateNotice"),
		[]base.Field{
			&base.Textarea{
				ID:     "content",
				Label:  p.T("attributes.content"),
				Helper: p.T("site-pages.can-markdown"),
			},
		},
	)
	p.TplName = "auth/form.html"

}

//CreateNotice create notice
// @router /notices [post]
func (p *Controller) CreateNotice() {
	p.MustAdmin()
	var fm fmContent
	fl, er := p.ParseForm(&fm)
	if er == nil {
		_, err := orm.NewOrm().Insert(&Notice{Content: fm.Content})
		p.Check(err)
		fl.Notice(p.T("site-pages.success"))
		p.Redirect(fl, "site.Controller.IndexNotice")
	} else {
		fl.Error(er.Error())
		p.Redirect(fl, "site.Controller.NewNotice")
	}
}

//EditNotice edit notice
// @router /notices/:id/edit [get]
func (p *Controller) EditNotice() {
	p.Dashboard()
	p.MustAdmin()
	var notice Notice
	err := orm.NewOrm().
		QueryTable(&notice).
		Filter("id", p.Ctx.Input.Param(":id")).
		One(&notice)
	p.Check(err)

	title := p.T("site-pages.edit-notice", notice.ID)
	p.Data["title"] = title
	p.Data["form"] = p.NewForm(
		"fm-edit-notice",
		title,
		base.MethodPost,
		p.URLFor("site.Controller.UpdateNotice", ":id", notice.ID),
		[]base.Field{
			&base.Textarea{
				ID:     "content",
				Label:  p.T("attributes.content"),
				Helper: p.T("site-pages.can-markdown"),
				Value:  notice.Content,
			},
		},
	)
	p.TplName = "auth/form.html"
}

//UpdateNotice update notice
// @router /notices/:id [post]
func (p *Controller) UpdateNotice() {
	p.MustAdmin()
	var notice Notice
	o := orm.NewOrm()
	err := o.QueryTable(&notice).
		Filter("id", p.Ctx.Input.Param(":id")).
		One(&notice)
	p.Check(err)

	var fm fmContent
	fl, er := p.ParseForm(&fm)
	if er == nil {
		notice.Content = fm.Content
		_, err = o.Update(&notice, "updated_at", "content")
		p.Check(err)
		fl.Notice(p.T("site-pages.success"))
		p.Redirect(fl, "site.Controller.IndexNotice")
	} else {
		fl.Error(er.Error())
		p.Redirect(fl, "site.Controller.EditNotice", ":id", notice.ID)
	}
}

//DestroyNotice destroy notice
// @router /notices/:id [delete]
func (p *Controller) DestroyNotice() {
	p.MustAdmin()
	_, err := orm.NewOrm().
		QueryTable(new(Notice)).
		Filter("id", p.Ctx.Input.Param(":id")).
		Delete()
	p.Check(err)

	p.Data["json"] = map[string]string{
		"to": p.URLFor("site.Controller.IndexNotice"),
	}
	p.ServeJSON()
}
