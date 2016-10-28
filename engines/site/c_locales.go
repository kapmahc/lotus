package site

import (
	"github.com/astaxie/beego/orm"
	"github.com/kapmahc/lotus/engines/base"
)

//GetAdminI18n locales
// @router /admin/locales [get]
func (p *Controller) GetAdminI18n() {
	p.Dashboard()
	p.MustAdmin()
	title := p.T("site-pages.admin-i18n")
	p.Data["title"] = title
	var items []base.Locale
	_, err := orm.NewOrm().
		QueryTable(new(base.Locale)).
		Filter("lang", p.Locale).
		OrderBy("code").All(&items)
	p.Check(err)
	p.Data["locales"] = items

	p.Data["form"] = p.NewForm(
		"fm-admin-i18n",
		title,
		base.MethodPost,
		p.URLFor("site.Controller.PostAdminI18n"),
		[]base.Field{
			&base.TextField{
				ID:    "code",
				Label: p.T("site-attributes.locale-code"),
			},
			&base.Textarea{
				ID:    "message",
				Label: p.T("site-attributes.locale-message"),
			},
		},
	)

	p.TplName = "site/admin/locales.html"
}

//PostAdminI18n locales
// @router /admin/locales [post]
func (p *Controller) PostAdminI18n() {
	p.MustAdmin()
	var fm fmLocale
	fl, er := p.ParseForm(&fm)

	if er == nil {
		base.SetLocale(p.Locale, fm.Code, fm.Message)
		fl.Notice(p.T("site-pages.success"))
	} else {
		fl.Error(er.Error())
	}
	p.Redirect(fl, "site.Controller.GetAdminI18n")
}
