package site

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/kapmahc/lotus/engines/auth"
	"github.com/kapmahc/lotus/engines/base"
)

//GetInstall install
// @router /install [get]
func (p *Controller) GetInstall() {
	p.mustEmptyDb()
	p.Data["form"] = p.NewForm(
		"fm-install",
		base.T(p.Locale, "site-forms.administrator"),
		base.MethodPost,
		p.URLFor("site.Controller.PostInstall"),
		[]base.Field{
			&base.TextField{
				ID:    "name",
				Label: base.T(p.Locale, "site-attributes.user-name"),
			},
			&base.EmailField{
				ID:    "email",
				Label: base.T(p.Locale, "attributes.email"),
			},
			&base.PasswordField{
				ID:    "password",
				Label: base.T(p.Locale, "attributes.password"),
			},
			&base.PasswordField{
				ID:    "passwordConfirmation",
				Label: base.T(p.Locale, "attributes.passwordConfirmation"),
			},
		},
	)
	p.TplName = "site/install.html"
}

//PostInstall install
// @router /install [post]
func (p *Controller) PostInstall() {
	p.mustEmptyDb()
	var fm fmInstall
	if p.ParseForm(&fm) {

	}
	p.Redirect(p.URLFor("site.Controller.GetInstall"), 302)
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
