package site

import (
	"time"

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
	flash, succ := p.ParseForm(&fm)
	if succ {
		if fm.Password == fm.PasswordConfirmation {
			user, err := auth.AddEmailUser(fm.Email, fm.Name, fm.Password)
			if err == nil {
				auth.AddLog(user.ID, "sign up")
				auth.ConfirmUser(user)
				auth.AddLog(user.ID, "confirm")
				for _, role := range []string{"admin", "root"} {
					auth.Allow(user.ID, role, auth.DefaultResourceType, auth.DefaultResourceID, 120, 0, 0)
				}
				Set("site.install", time.Now(), false)
			}
			if err != nil {
				beego.Error(err)
				p.Abort("500")
			}
			p.Redirect(p.URLFor("auth.Controller.GetSignIn"), 302)
			return
		}
		flash.Error("passwords not match")
	}
	flash.Store(&p.Controller.Controller)
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
