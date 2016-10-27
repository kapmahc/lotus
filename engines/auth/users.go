package auth

import (
	"github.com/SermoDigital/jose/jwt"
	"github.com/astaxie/beego/orm"
	"github.com/kapmahc/lotus/engines/base"
)

//GetSignIn sign on form
// @router /sign-in [get]
func (p *Controller) GetSignIn() {
	p.Data["form"] = p.NewForm(
		"fm-sign-in",
		base.T(p.Locale, "auth-pages.sign-in"),
		base.MethodPost,
		p.URLFor("auth.Controller.PostSignIn"),
		[]base.Field{
			&base.EmailField{
				ID:    "email",
				Label: base.T(p.Locale, "attributes.email"),
			},
			&base.PasswordField{
				ID:    "password",
				Label: base.T(p.Locale, "attributes.password"),
			},
		},
	)
	p.TplName = "auth/form.html"
}

//PostSignIn sign in
// @router /sign-in [post]
func (p *Controller) PostSignIn() {

}

//GetSignUp sign up form
// @router /sign-up [get]
func (p *Controller) GetSignUp() {

	p.Data["form"] = p.NewForm(
		"fm-sign-up",
		base.T(p.Locale, "auth-pages.sign-up"),
		base.MethodPost,
		p.URLFor("auth.Controller.PostSignUp"),
		[]base.Field{
			&base.TextField{
				ID:    "name",
				Label: base.T(p.Locale, "auth-attributes.user-name"),
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
	p.TplName = "auth/form.html"
}

//PostSignUp sign up form
// @router /sign-up [post]
func (p *Controller) PostSignUp() {
	var fm fmSignUp
	fl, er := p.ParseForm(&fm)
	if er == nil {
		if fm.Password != fm.PasswordConfirmation {
			er = p.Error("auth-logs.passwords-not-match")
		}
	}
	if er == nil {
		if _, err := GetUserByEmail(fm.Email); err != orm.ErrNoRows {
			er = p.Error("auth-logs.email-already-exists")
		}
	}

	if er == nil {
		user, err := AddEmailUser(fm.Email, fm.Name, fm.Password)
		if err == nil {
			AddLog(user.ID, p.T("auth-logs.sign-up"))
		}
		p.Check(err)
		p.sendMail(actConfirm, user.Email, user.UID)
		p.Redirect(p.URLFor("auth.Controller.GetSignIn"), 302)
		return
	}
	fl.Error(er.Error())
	fl.Store(&p.Controller.Controller)
	p.Redirect(p.URLFor("auth.Controller.GetSignIn"), 302)
}

//GetConfirm confirm form
// @router /confirm [get]
func (p *Controller) GetConfirm() {
	token := p.GetString("token")
	if token != "" {
		cm, er := base.ParseToken([]byte(token))
		if cm.Get("act").(string) != actConfirm {
			er = p.Error("auth-logs.bad-token")
		}
		var user *User
		if er == nil {
			user, er = GetUserByUID(cm.Get("uid").(string))
		}
		if er == nil {
			if user.IsConfirmed() {
				er = p.Error("auth-logs.user-already-comfirmed")
			}
		}
		if er == nil {
			ConfirmUser(user)
			AddLog(user.ID, p.T("auth-logs.confirm"))
			p.Redirect(p.URLFor("auth.Controller.GetSignIn"), 302)
			return
		}
	}

	p.Data["form"] = p.NewForm(
		"fm-confirm",
		base.T(p.Locale, "auth-pages.confirm"),
		base.MethodPost,
		p.URLFor("auth.Controller.PostConfirm"),
		[]base.Field{
			&base.EmailField{
				ID:    "email",
				Label: base.T(p.Locale, "attributes.email"),
			},
		},
	)
	p.TplName = "auth/form.html"
}

//PostConfirm confirm
// @router /confirm [post]
func (p *Controller) PostConfirm() {
	var fm fmEmail
	fl, er := p.ParseForm(&fm)

	var user *User
	if er == nil {
		user, er = GetUserByEmail(fm.Email)
	}
	if er == nil {
		if user.IsConfirmed() {
			er = p.Error("auth-logs.user-already-comfirmed")
		}
	}

	if er == nil {
		p.sendMail(actConfirm, user.Email, user.UID)
		p.Redirect(p.URLFor("auth.Controller.GetSignIn"), 302)
		return
	}
	fl.Error(er.Error())
	fl.Store(&p.Controller.Controller)
	p.Redirect(p.URLFor("auth.Controller.GetConfirm"), 302)
}

//GetForgotPassword forgot password form
// @router /forgot-password [get]
func (p *Controller) GetForgotPassword() {
	p.Data["form"] = p.NewForm(
		"fm-forgot-password",
		base.T(p.Locale, "auth-pages.forgot-password"),
		base.MethodPost,
		p.URLFor("auth.Controller.PostForgotPassword"),
		[]base.Field{
			&base.EmailField{
				ID:    "email",
				Label: base.T(p.Locale, "attributes.email"),
			},
		},
	)
	p.TplName = "auth/form.html"
}

//PostForgotPassword forgot password
// @router /forgot-password [post]
func (p *Controller) PostForgotPassword() {
	var fm fmEmail
	fl, er := p.ParseForm(&fm)

	var user *User
	if er == nil {
		user, er = GetUserByEmail(fm.Email)
	}

	if er == nil {
		p.sendMail(actResetPassword, user.Email, user.UID)
		p.Redirect(p.URLFor("auth.Controller.GetSignIn"), 302)
		return
	}
	fl.Error(er.Error())
	fl.Store(&p.Controller.Controller)
	p.Redirect(p.URLFor("auth.Controller.GetResetPassword"), 302)
}

//GetResetPassword reset password form
// @router /reset-password [get]
func (p *Controller) GetResetPassword() {
	p.Data["form"] = p.NewForm(
		"fm-reset-password",
		base.T(p.Locale, "auth-pages.reset-password"),
		base.MethodPost,
		p.URLFor("auth.Controller.PostResetPassword"),
		[]base.Field{
			&base.HiddenField{
				ID:    "token",
				Value: p.GetString("token"),
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
	p.TplName = "auth/form.html"
}

//PostResetPassword reset password form
// @router /reset-password [post]
func (p *Controller) PostResetPassword() {
	var fm fmResetPassword
	fl, er := p.ParseForm(&fm)
	if er == nil {
		if fm.Password != fm.PasswordConfirmation {
			er = p.Error("auth-logs.passwords-not-match")
		}
	}
	var cm jwt.Claims
	if er == nil {
		cm, er = base.ParseToken([]byte(fm.Token))
	}
	if cm.Get("act").(string) != actResetPassword {
		er = p.Error("auth-logs.bad-token")
	}
	var user *User
	if er == nil {
		user, er = GetUserByUID(cm.Get("uid").(string))
	}
	if er == nil {
		user.SetPassword(fm.Password)
		_, err := orm.NewOrm().Update(user, "password", "updated_at")
		if err == nil {
			AddLog(user.ID, p.T("auth-logs.reset-password"))
		}
		p.Check(err)
		p.Redirect(p.URLFor("auth.Controller.GetSignIn"), 302)
		return
	}

	fl.Error(er.Error())
	fl.Store(&p.Controller.Controller)
	p.Redirect(p.URLFor("auth.Controller.GetSignIn"), 302)
}
