package site

import "github.com/kapmahc/lotus/engines/base"

//GetAdminSMTP smtp
// @router /admin/smtp [get]
func (p *Controller) GetAdminSMTP() {
	p.Dashboard()
	p.MustAdmin()
	var smtp SMTP
	base.Get("smtp", &smtp)
	title := p.T("site-pages.admin-smtp")
	p.Data["title"] = title
	p.Data["form"] = p.NewForm(
		"fm-admin-smtp",
		title,
		base.MethodPost,
		p.URLFor("site.Controller.PostAdminSMTP"),
		[]base.Field{
			&base.TextField{
				ID:    "host",
				Label: p.T("attributes.host"),
				Value: smtp.Host,
			},
			&base.Select{
				ID:    "port",
				Label: p.T("attributes.port"),
				Value: smtp.Port,
				Options: map[interface{}]interface{}{
					25:  25,
					465: 465,
					587: 587,
				},
			},
			&base.TextField{
				ID:    "username",
				Label: p.T("attributes.username"),
				Value: smtp.Username,
			},
			&base.PasswordField{
				ID:    "password",
				Label: p.T("attributes.password"),
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

//PostAdminSMTP smtp
// @router /admin/smtp [post]
func (p *Controller) PostAdminSMTP() {
	p.MustAdmin()
	var fm fmSMTP
	fl, er := p.ParseForm(&fm)
	if er == nil {
		if fm.Password != fm.PasswordConfirmation {
			er = p.Error("auth-logs.passwords-not-match")
		}
	}
	if er == nil {
		base.Set(
			"smtp",
			&SMTP{
				Host:     fm.Host,
				Port:     fm.Port,
				Username: fm.Username,
				Password: fm.Password,
			},
			true,
		)
		user := p.CurrentUser()
		user.Log(p.T("site-logs.update-smtp"))

		fl.Notice(p.T("site-pages.success"))
	} else {
		fl.Error(er.Error())
	}
	p.Redirect(fl, "site.Controller.GetAdminSMTP")
}
