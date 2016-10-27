package auth

import "github.com/kapmahc/lotus/engines/base"

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

}

//GetConfirm confirm form
// @router /confirm [get]
func (p *Controller) GetConfirm() {
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

}
