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
	title := p.T("site-pages.administrator")
	p.Data["title"] = title
	p.Data["form"] = p.NewForm(
		"fm-install",
		title,
		base.MethodPost,
		p.URLFor("site.Controller.PostInstall"),
		[]base.Field{
			&base.TextField{
				ID:    "name",
				Label: p.T("auth-attributes.user-name"),
			},
			&base.EmailField{
				ID:    "email",
				Label: p.T("attributes.email"),
			},
			&base.PasswordField{
				ID:    "password",
				Label: p.T("attributes.password"),
			},
			&base.PasswordField{
				ID:    "passwordConfirmation",
				Label: p.T("attributes.passwordConfirmation"),
			},
		},
	)
	p.TplName = "auth/form.html"
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
			user.Log(p.T("auth-logs.sign-up"))
			auth.ConfirmUser(user)
			user.Log(p.T("auth-logs.confirm"))
			for _, role := range []string{auth.AdminRole, auth.RootRole} {
				user.Allow(role, auth.DefaultResourceType, auth.DefaultResourceID, 120, 0, 0)
			}
			base.Set("site.install", time.Now(), false)
		}
		p.Check(err)
		p.Redirect(nil, "auth.Controller.GetSignIn")
		return
	}
	fl.Error(er.Error())
	fl.Store(&p.Controller.Controller)
	p.Redirect(fl, "site.Controller.GetInstall")
}

func (p *Controller) mustEmptyDb() {
	o := orm.NewOrm()
	ct, er := o.QueryTable(new(auth.User)).Count()
	p.Check(er)
	if ct > 0 {
		p.Abort("404")
	}
}
