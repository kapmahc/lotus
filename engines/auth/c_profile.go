package auth

import (
	"github.com/astaxie/beego/orm"
	"github.com/kapmahc/lotus/engines/base"
)

//GetInfo info page
// @router /info [get]
func (p *Controller) GetInfo() {
	p.Dashboard()
	user := p.CurrentUser()
	title := p.T("auth-pages.info")
	p.Data["title"] = title
	p.Data["form"] = p.NewForm(
		"fm-info",
		title,
		base.MethodPost,
		p.URLFor("auth.Controller.PostInfo"),
		[]base.Field{
			&base.EmailField{
				ID:       "email",
				Label:    p.T("attributes.email"),
				Value:    user.Email,
				Readonly: true,
			},
			&base.TextField{
				ID:    "name",
				Label: p.T("auth-attributes.user-name"),
				Value: user.Name,
			},
			&base.TextField{
				ID:    "logo",
				Label: p.T("auth-attributes.user-logo"),
				Value: user.Logo,
			},
			&base.TextField{
				ID:    "home",
				Label: p.T("auth-attributes.user-home"),
				Value: user.Home,
			},
		},
	)
	p.TplName = "auth/form.html"
}

//PostInfo set user's info
// @router /info [post]
func (p *Controller) PostInfo() {
	p.MustSignIn()
	var fm fmInfo
	fl, er := p.ParseForm(&fm)

	if er == nil {
		user := p.CurrentUser()
		user.Home = fm.Home
		user.Logo = fm.Logo
		user.Name = fm.Name
		_, err := orm.NewOrm().Update(user, "updated_at", "name", "home", "logo")
		p.Check(err)
		user.Log(p.T("auth-logs.update-profile"))
		fl.Notice(p.T("site-pages.success"))
	} else {
		fl.Error(er.Error())
	}
	p.Redirect(fl, "auth.Controller.GetInfo")
}

//GetChangePassword change-password page
// @router /change-password [get]
func (p *Controller) GetChangePassword() {
	p.Dashboard()
	title := p.T("auth-pages.change-password")
	p.Data["title"] = title
	p.Data["form"] = p.NewForm(
		"fm-change-password",
		title,
		base.MethodPost,
		p.URLFor("auth.Controller.PostChangePassword"),
		[]base.Field{
			&base.PasswordField{
				ID:     "currentPassword",
				Label:  p.T("auth-attributes.user-currentPassword"),
				Helper: p.T("auth-pages.need-password-to-change"),
			},
			&base.PasswordField{
				ID:     "newPassword",
				Label:  p.T("auth-attributes.user-newPassword"),
				Helper: p.T("auth-pages.password-must-in-size"),
			},
			&base.PasswordField{
				ID:     "passwordConfirmation",
				Label:  p.T("attributes.passwordConfirmation"),
				Helper: p.T("auth-pages.passwords-must-match"),
			},
		},
	)
	p.TplName = "auth/form.html"
}

//PostChangePassword change-password
// @router /change-password [post]
func (p *Controller) PostChangePassword() {
	p.MustSignIn()
	var fm fmChangePassword
	fl, er := p.ParseForm(&fm)
	if er == nil {
		if fm.NewPassword != fm.PasswordConfirmation {
			er = p.Error("auth-logs.passwords-not-match")
		}
	}
	user := p.CurrentUser()
	if er == nil {
		if !user.IsPassword(fm.CurrentPassword) {
			er = p.Error("auth-logs.email-password-not-match")
		}
	}

	if er == nil {
		user.SetPassword(fm.NewPassword)
		_, err := orm.NewOrm().Update(user, "updated_at", "password")
		p.Check(err)
		user.Log(p.T("auth-logs.change-password"))
		fl.Notice(p.T("site-pages.success"))
	} else {
		fl.Error(er.Error())
	}
	p.Redirect(fl, "auth.Controller.GetChangePassword")
}

//GetLogs logs page
// @router /logs [get]
func (p *Controller) GetLogs() {
	p.Dashboard()
	p.Data["title"] = p.T("auth-pages.logs")
	var logs []Log
	_, err := orm.NewOrm().
		QueryTable(new(Log)).
		Filter("user_id", p.CurrentUser().ID).
		OrderBy("-id").Limit(120).All(&logs)
	// _, err := orm.NewOrm().
	// 	QueryTable(new(Log)).
	// 	Filter("User", p.CurrentUser().ID).
	// 	RelatedSel().
	// 	OrderBy("-id").Limit(120).
	// 	All(&logs)
	p.Check(err)
	p.Data["logs"] = logs
	p.TplName = "auth/logs.html"
}
