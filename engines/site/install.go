package site

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/beego/i18n"
	"github.com/kapmahc/lotus/engines/auth"
	"github.com/kapmahc/lotus/engines/base"
)

//GetInstall install
// @router /install [get]
func (p *Controller) GetInstall() {
	p.mustEmptyDb()
	p.Data["form"] = &base.Form{
		ID:     "fm-install",
		Method: "post",
		Title:  i18n.Tr(p.Locale, "site-forms.administrator"),
		Action: p.URLFor("site.Controller.GetInstall"),
		Fields: []base.Field{
			&base.TextField{
				ID:    "name",
				Label: i18n.Tr(p.Locale, "site-attributes.user-name"),
			},
			&base.EmailField{
				ID:    "email",
				Label: i18n.Tr(p.Locale, "attributes.email"),
			},
			&base.PasswordField{
				ID:    "password",
				Label: i18n.Tr(p.Locale, "attributes.password"),
			},
			&base.PasswordField{
				ID:    "passwordConfirmation",
				Label: i18n.Tr(p.Locale, "attributes.passwordConfirmation"),
			},
		},
	}
	p.TplName = "site/install.html"
}

//PostInstall install
// @router /install [post]
func (p *Controller) PostInstall() {
	p.mustEmptyDb()
}

func (p *Controller) mustEmptyDb() {
	o := orm.NewOrm()
	ct, er := o.QueryTable(new(auth.User)).Count()
	if er != nil {
		beego.Error(er)
		p.Abort("500")
	}
	if ct > 0 {
		p.Abort("404")
	}
}
