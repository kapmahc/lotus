package mail

import (
	"github.com/astaxie/beego/orm"
	"github.com/kapmahc/lotus/engines/base"
)

//IndexAlias index aliases
// @router /aliases [get]
func (p *Controller) IndexAlias() {
	var aliases []Alias
	_, err := orm.NewOrm().
		QueryTable(new(Alias)).
		OrderBy("-updated_at").
		RelatedSel().
		All(&aliases)
	p.Check(err)
	p.Data["aliases"] = aliases
	p.Data["title"] = p.T("ops-mail-pages.aliases")
	p.TplName = "ops/mail/aliases/index.html"
}

//NewAlias new alias
// @router /aliases/new [get]
func (p *Controller) NewAlias() {

	title := p.T("ops-mail-pages.new-alias")
	p.Data["title"] = title
	p.Data["form"] = p.NewForm(
		"fm-new-alias",
		title,
		base.MethodPost,
		p.URLFor("mail.Controller.CreateAlias"),
		[]base.Field{
			p.domainOptions(),
			&base.TextField{
				ID:    "source",
				Label: p.T("ops-mail-attributes.alias-source"),
			},
			&base.TextField{
				ID:    "destination",
				Label: p.T("ops-mail-attributes.alias-destination"),
			},
		},
	)
	p.TplName = "auth/form.html"
}

//CreateAlias create alias
// @router /aliases [post]
func (p *Controller) CreateAlias() {
	var fm fmAlias
	fl, er := p.ParseForm(&fm)
	o := orm.NewOrm()

	if er == nil {
		count1, err := o.QueryTable(new(User)).Filter("email", fm.Source).Count()
		p.Check(err)
		count2, err := o.QueryTable(new(Alias)).Filter("source", fm.Source).Count()
		if count1 > 0 || count2 > 0 {
			er = p.Error("auth-logs.email-already-exists")
		}
	}

	if er == nil {
		var domain Domain
		err := o.QueryTable(&domain).Filter("id", fm.DomainID).One(&domain, "id")
		p.Check(err)
		_, err = o.Insert(&Alias{
			Source:      fm.Source,
			Domain:      &domain,
			Destination: fm.Destination,
		})
		p.Check(err)
		fl.Notice(p.T("site-pages.success"))
		p.Redirect(fl, "mail.Controller.IndexAlias")
	} else {
		fl.Error(er.Error())
		p.Redirect(fl, "mail.Controller.NewAlias")
	}
}

//EditAlias edit alias
// @router /aliases/:id [get]
func (p *Controller) EditAlias() {
	var alias Alias
	err := orm.NewOrm().
		QueryTable(&alias).
		Filter("id", p.Ctx.Input.Param(":id")).
		One(&alias)
	p.Check(err)

	title := p.T("ops-mail-pages.edit-alias", alias.ID)
	p.Data["title"] = title
	p.Data["form"] = p.NewForm(
		"fm-edit-alias",
		title,
		base.MethodPost,
		p.URLFor("mail.Controller.UpdateAlias", ":id", alias.ID),
		[]base.Field{
			&base.HiddenField{
				ID:    "domain_id",
				Value: alias.Domain.ID,
			},
			&base.TextField{
				ID:    "source",
				Label: p.T("ops-mail-attributes.alias-source"),
				Value: alias.Source,
			},
			&base.TextField{
				ID:    "destination",
				Label: p.T("ops-mail-attributes.alias-destination"),
				Value: alias.Destination,
			},
		},
	)
	p.TplName = "auth/form.html"
}

//UpdateAlias update alias
// @router /aliases/:id [post]
func (p *Controller) UpdateAlias() {

	var alias Alias
	o := orm.NewOrm()
	err := o.QueryTable(&alias).
		Filter("id", p.Ctx.Input.Param(":id")).
		One(&alias)
	p.Check(err)

	var fm fmAlias
	fl, er := p.ParseForm(&fm)

	if er == nil {
		alias.Source = fm.Source
		alias.Destination = fm.Destination
		_, err = o.Update(&alias, "updated_at", "source", "destination")
		p.Check(err)
		fl.Notice(p.T("site-pages.success"))
		p.Redirect(fl, "mail.Controller.IndexAlias")
	} else {
		fl.Error(er.Error())
		p.Redirect(fl, "mail.Controller.EditAlias", ":id", alias.ID)
	}
}

//DestroyAlias destroy alias
// @router /aliases/:id [delete]
func (p *Controller) DestroyAlias() {
	_, err := orm.NewOrm().QueryTable(new(Alias)).
		Filter("id", p.Ctx.Input.Param(":id")).Delete()
	p.Check(err)

	p.Data["json"] = map[string]string{
		"to": p.URLFor("mail.Controller.IndexAlias"),
	}
	p.ServeJSON()
}
