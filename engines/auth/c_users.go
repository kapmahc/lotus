package auth

import (
	"time"

	"github.com/SermoDigital/jose/jwt"
	"github.com/astaxie/beego/orm"
	"github.com/kapmahc/lotus/engines/base"
)

//GetSignIn sign on form
// @router /sign-in [get]
func (p *Controller) GetSignIn() {
	p.Data["form"] = p.NewForm(
		"fm-sign-in",
		p.T("auth-pages.sign-in"),
		base.MethodPost,
		p.URLFor("auth.Controller.PostSignIn"),
		[]base.Field{
			&base.EmailField{
				ID:    "email",
				Label: p.T("attributes.email"),
			},
			&base.PasswordField{
				ID:    "password",
				Label: p.T("attributes.password"),
			},
		},
	)
	p.TplName = "auth/non-sign-in.html"
}

//PostSignIn sign in
// @router /sign-in [post]
func (p *Controller) PostSignIn() {
	var fm fmSignIn
	fl, er := p.ParseForm(&fm)

	if er == nil {
		var user User
		o := orm.NewOrm()
		err := o.
			QueryTable(&user).
			Filter("provider_type", ProvideByEmail).
			Filter("provider_id", fm.Email).One(&user)

		if err == nil && user.IsPassword(fm.Password) {
			user.SignInCount++
			now := time.Now()
			user.LastSignInAt = &now
			_, err = o.Update(&user, "updated_at", "last_sign_in_at", "sign_in_count")
			p.Check(err)

			user.Log(p.T("auth-logs.sign-in"))

			p.SetSession("uid", user.UID)
			p.SetSession("name", user.Name)
			p.Redirect(fl, "auth.Controller.GetLogs")
			return
		}
		er = p.Error("auth-logs.email-password-not-match")
	}
	fl.Error(er.Error())
	p.Redirect(fl, "auth.Controller.GetSignIn")
}

//GetSignOut sign out
// @router /sign-out [get]
func (p *Controller) GetSignOut() {
	p.MustSignIn()
	p.CurrentUser().Log(p.T("auth-logs.sign-out"))
	p.DestroySession()
	p.Redirect(nil, "auth.Controller.GetSignIn")
}

//GetSignUp sign up form
// @router /sign-up [get]
func (p *Controller) GetSignUp() {

	p.Data["form"] = p.NewForm(
		"fm-sign-up",
		p.T("auth-pages.sign-up"),
		base.MethodPost,
		p.URLFor("auth.Controller.PostSignUp"),
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
				ID:     "password",
				Label:  p.T("attributes.password"),
				Helper: p.T("auth-pages.password-must-in-size"),
			},
			&base.PasswordField{
				ID:     "passwordConfirmation",
				Label:  p.T("attributes.passwordConfirmation"),
				Helper: p.T("auth-pages.passwords-must-match"),
			},
		},
	)
	p.TplName = "auth/non-sign-in.html"
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
			user.Log(p.T("auth-logs.sign-up"))
		}
		p.Check(err)
		p.sendMail(actConfirm, user.Email, user.UID)
		fl.Notice(p.T("auth-pages.confirm-success"))
		p.Redirect(fl, "auth.Controller.GetSignIn")
		return
	}
	fl.Error(er.Error())
	p.Redirect(fl, "auth.Controller.GetSignUp")
}

//GetConfirm confirm form
// @router /confirm [get]
func (p *Controller) GetConfirm() {
	token := p.GetString("token")
	if token != "" {
		cm, er := base.ParseToken([]byte(token))
		if er == nil {
			if cm.Get("act").(string) != actConfirm {
				er = p.Error("auth-logs.bad-token")
			}
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
			user.Log(p.T("auth-logs.confirm"))
			p.Redirect(nil, "auth.Controller.GetSignIn")
			return
		}
	}

	p.Data["form"] = p.NewForm(
		"fm-confirm",
		p.T("auth-pages.confirm"),
		base.MethodPost,
		p.URLFor("auth.Controller.PostConfirm"),
		[]base.Field{
			&base.EmailField{
				ID:    "email",
				Label: p.T("attributes.email"),
			},
		},
	)
	p.TplName = "auth/non-sign-in.html"
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
		fl.Notice(p.T("auth-pages.confirm-success"))
		p.Redirect(fl, "auth.Controller.GetSignIn")
		return
	}
	fl.Error(er.Error())
	p.Redirect(fl, "auth.Controller.GetConfirm")
}

//GetForgotPassword forgot password form
// @router /forgot-password [get]
func (p *Controller) GetForgotPassword() {
	p.Data["form"] = p.NewForm(
		"fm-forgot-password",
		p.T("auth-pages.forgot-password"),
		base.MethodPost,
		p.URLFor("auth.Controller.PostForgotPassword"),
		[]base.Field{
			&base.EmailField{
				ID:    "email",
				Label: p.T("attributes.email"),
			},
		},
	)
	p.TplName = "auth/non-sign-in.html"
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
		fl.Notice(p.T("auth-pages.forgot-password-success"))
		p.Redirect(fl, "auth.Controller.GetSignIn")
		return
	}
	fl.Error(er.Error())
	p.Redirect(fl, "auth.Controller.GetResetPassword")
}

//GetResetPassword reset password form
// @router /reset-password [get]
func (p *Controller) GetResetPassword() {
	p.Data["form"] = p.NewForm(
		"fm-reset-password",
		p.T("auth-pages.reset-password"),
		base.MethodPost,
		p.URLFor("auth.Controller.PostResetPassword"),
		[]base.Field{
			&base.HiddenField{
				ID:    "token",
				Value: p.GetString("token"),
			},
			&base.PasswordField{
				ID:     "password",
				Label:  p.T("attributes.password"),
				Helper: p.T("auth-pages.password-must-in-size"),
			},
			&base.PasswordField{
				ID:     "passwordConfirmation",
				Label:  p.T("attributes.passwordConfirmation"),
				Helper: p.T("auth-pages.passwords-must-match"),
			},
		},
	)
	p.TplName = "auth/non-sign-in.html"
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
			user.Log(p.T("auth-logs.reset-password"))
		}
		p.Check(err)
		fl.Notice(p.T("auth-pages.reset-password-success"))
		p.Redirect(fl, "auth.Controller.GetSignIn")
		return
	}

	fl.Error(er.Error())
	p.Redirect(fl, "auth.Controller.GetResetPassword")
}
