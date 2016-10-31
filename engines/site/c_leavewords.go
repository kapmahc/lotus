package site

import (
	"github.com/astaxie/beego/orm"
	"github.com/kapmahc/lotus/engines/base"
)

//IndexLeaveWord list leavewords
// @router /leavewords [get]
func (p *Controller) IndexLeaveWord() {
	p.Dashboard()
	p.MustAdmin()
	var items []LeaveWord
	_, err := orm.NewOrm().
		QueryTable(new(LeaveWord)).
		OrderBy("-id").
		Limit(60).
		All(&items)
	p.Check(err)
	p.Data["leavewords"] = items
	p.Data["title"] = p.T("site-pages.leavewords")
	p.TplName = "site/leavewords/index.html"
}

//DestroyLeaveWord remove leaveword
// @router /leavewords/:id [delete]
func (p *Controller) DestroyLeaveWord() {
	p.MustAdmin()
	_, err := orm.NewOrm().
		QueryTable(new(LeaveWord)).
		Filter("id", p.Ctx.Input.Param(":id")).
		Delete()
	p.Check(err)

	p.Data["json"] = map[string]string{
		"to": p.URLFor("site.Controller.IndexLeaveWord"),
	}
	p.ServeJSON()
}

//CreateLeaveWord create leaveword
// @router /leavewords [post]
func (p *Controller) CreateLeaveWord() {
	var fm fmContent
	fl, er := p.ParseForm(&fm)
	if er == nil {
		lw := LeaveWord{Content: fm.Content}
		_, err := orm.NewOrm().Insert(&lw)
		p.Check(err)
		fl.Notice(p.T("site-pages.success"))
	} else {
		fl.Error(er.Error())
	}
	p.Redirect(fl, "site.Controller.NewLeaveWord")
}

//NewLeaveWord new leaveword
// @router /leavewords/new [get]
func (p *Controller) NewLeaveWord() {
	p.Data["form"] = p.NewForm(
		"fm-new-leaveword",
		p.T("site-pages.new-leaveword"),
		base.MethodPost,
		p.URLFor("site.Controller.CreateLeaveWord"),
		[]base.Field{
			&base.Textarea{
				ID:     "content",
				Label:  p.T("attributes.content"),
				Helper: p.T("site-pages.can-markdown"),
			},
		},
	)
	p.Layout = ""
	p.TplName = "auth/non-sign-in.html"
}
