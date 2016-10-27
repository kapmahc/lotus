package auth

import "github.com/kapmahc/lotus/engines/base"

//GetSignIn sign on form
// @router /sign-in [get]
func (p *Controller) GetSignIn() {
	p.Data["form"] = p.NewForm(
		"fm-install",
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
	p.TplName = "auth/sign-up.html"
}

//PostSignUp sign up form
// @router /sign-up [post]
func (p *Controller) PostSignUp() {

}

//GetConfirm confirm form
// @router /confirm [get]
func (p *Controller) GetConfirm() {
	p.TplName = "auth/confirm.html"
}

//PostConfirm confirm
// @router /confirm [post]
func (p *Controller) PostConfirm() {

}

//GetForgotPassword forgot password form
// @router /forgot-password [get]
func (p *Controller) GetForgotPassword() {
	p.TplName = "auth/forgot-password.html"
}

//PostForgotPassword forgot password
// @router /forgot-password [post]
func (p *Controller) PostForgotPassword() {

}

//GetResetPassword reset password form
// @router /reset-password [get]
func (p *Controller) GetResetPassword() {
	p.TplName = "auth/reset-password.html"
}

//PostResetPassword reset password form
// @router /reset-password [post]
func (p *Controller) PostResetPassword() {

}
