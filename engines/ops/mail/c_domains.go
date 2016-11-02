package mail

import (
	"github.com/astaxie/beego/orm"
	"github.com/kapmahc/lotus/engines/base"
)

//IndexDomain index domains
// @router /domains [get]
func (p *Controller) IndexDomain() {
	var domains []Domain
	_, err := orm.NewOrm().
		QueryTable(new(Domain)).
		OrderBy("-updated_at").
		All(&domains, "id", "name")
	p.Check(err)
	p.Data["domains"] = domains
	p.Data["title"] = p.T("ops-mail-pages.domains")
	p.TplName = "ops/mail/domains/index.html"
}

//NewDomain new domain
// @router /domains/new [get]
func (p *Controller) NewDomain() {
	title := p.T("ops-mail-pages.new-domain")
	p.Data["title"] = title
	p.Data["form"] = p.NewForm(
		"fm-new-domain",
		title,
		base.MethodPost,
		p.URLFor("mail.Controller.CreateDomain"),
		[]base.Field{
			&base.TextField{
				ID:    "name",
				Label: p.T("attributes.name"),
			},
		},
	)
	p.TplName = "auth/form.html"
}

//CreateDomain create domain
// @router /domains [post]
func (p *Controller) CreateDomain() {
	var fm fmDomain
	fl, er := p.ParseForm(&fm)
	o := orm.NewOrm()

	if er == nil {
		count, err := o.QueryTable(new(Domain)).Filter("name", fm.Name).Count()
		p.Check(err)
		if count > 0 {
			er = p.Error("ops-mail-logs.domain-already-exists")
		}
	}

	if er == nil {
		_, err := o.Insert(&Domain{Name: fm.Name})
		p.Check(err)
		fl.Notice(p.T("site-pages.success"))
		p.Redirect(fl, "mail.Controller.IndexDomain")
	} else {
		fl.Error(er.Error())
		p.Redirect(fl, "mail.Controller.NewDomain")
	}
}

//EditDomain edit domain
// @router /domains/:id [get]
func (p *Controller) EditDomain() {
	var domain Domain
	err := orm.NewOrm().
		QueryTable(&domain).
		Filter("id", p.Ctx.Input.Param(":id")).
		One(&domain)
	p.Check(err)

	title := p.T("ops-mail-pages.edit-domain", domain.ID)
	p.Data["title"] = title
	p.Data["form"] = p.NewForm(
		"fm-edit-domain",
		title,
		base.MethodPost,
		p.URLFor("mail.Controller.UpdateDomain", ":id", domain.ID),
		[]base.Field{
			&base.TextField{
				ID:    "name",
				Label: p.T("attributes.name"),
				Value: domain.Name,
			},
		},
	)
	p.TplName = "auth/form.html"
}

//UpdateDomain update domain
// @router /domains/:id [post]
func (p *Controller) UpdateDomain() {

	var fm fmDomain
	fl, er := p.ParseForm(&fm)

	var domain Domain
	o := orm.NewOrm()
	err := o.QueryTable(&domain).
		Filter("id", p.Ctx.Input.Param(":id")).
		One(&domain)
	p.Check(err)
	if er == nil {
		var count int64
		count, err = o.QueryTable(new(Domain)).Filter("name", fm.Name).Count()
		p.Check(err)
		if count > 0 {
			er = p.Error("ops-mail-logs.domain-already-exists")
		}
	}

	if er == nil {
		domain.Name = fm.Name
		_, err = o.Update(&domain, "updated_at", "name")
		p.Check(err)
		fl.Notice(p.T("site-pages.success"))
		p.Redirect(fl, "mail.Controller.IndexDomain")
	} else {
		fl.Error(er.Error())
		p.Redirect(fl, "mail.Controller.EditDomain", ":id", domain.ID)
	}
}

//DestroyDomain destroy domain
// @router /domains/:id [delete]
func (p *Controller) DestroyDomain() {
	_, err := orm.NewOrm().QueryTable(new(Domain)).
		Filter("id", p.Ctx.Input.Param(":id")).Delete()
	p.Check(err)

	p.Data["json"] = map[string]string{
		"to": p.URLFor("mail.Controller.IndexDomain"),
	}
	p.ServeJSON()
}
