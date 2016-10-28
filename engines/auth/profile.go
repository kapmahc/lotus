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
	p.Data["title"] = p.T("auth-pages.change-password")
	p.TplName = "auth/change-password.html"
}

//GetLogs logs page
// @router /logs [get]
func (p *Controller) GetLogs() {
	p.Dashboard()
	p.Data["title"] = p.T("auth-pages.logs")
	var logs []Log
	_, err := orm.NewOrm().QueryTable(new(Log)).OrderBy("-id").Limit(120).All(&logs)
	p.Check(err)
	p.Data["logs"] = logs
	p.TplName = "auth/logs.html"
}
