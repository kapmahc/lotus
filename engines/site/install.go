package site

import (
	"time"

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
				Label: base.T(p.Locale, "auth.attributes.user-name"),
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
	fl, er := p.ParseForm(&fm)
	if er == nil {
		if fm.Password != fm.PasswordConfirmation {
			er = p.Error("auth-logs.passwords-not-match")
		}
	}
	if er == nil {
		user, err := auth.AddEmailUser(fm.Email, fm.Name, fm.Password)
		if err == nil {
			auth.AddLog(user.ID, p.T("auth-logs.sign-up"))
			auth.ConfirmUser(user)
			auth.AddLog(user.ID, p.T("auth-logs.confirm"))
			for _, role := range []string{"admin", "root"} {
				auth.Allow(user.ID, role, auth.DefaultResourceType, auth.DefaultResourceID, 120, 0, 0)
			}
			Set("site.install", time.Now(), false)
		}
		p.Check(err)
		p.Redirect(p.URLFor("auth.Controller.GetSignIn"), 302)
		return
	}
	fl.Error(er.Error())
	fl.Store(&p.Controller.Controller)
	p.Redirect(p.URLFor("site.Controller.GetInstall"), 302)
}

func (p *Controller) mustEmptyDb() {
	o := orm.NewOrm()
	ct, er := o.QueryTable(new(auth.User)).Count()
	p.Check(er)
	if ct > 0 {
		p.Abort("404")
	}
}
