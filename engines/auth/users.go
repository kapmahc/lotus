package auth

//GetSignIn sign on form
// @router /sign-in [get]
func (p *Controller) GetSignIn() {
	p.TplName = "auth/sign-in.html"
}

//GetSignUp sign up form
// @router /sign-up [get]
func (p *Controller) GetSignUp() {
	p.TplName = "auth/sign-up.html"
}
