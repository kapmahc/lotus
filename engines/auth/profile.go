package auth

//GetProfile profile page
// @router /profile [get]
func (p *Controller) GetProfile() {
	p.Data["title"] = p.T("auth-pages.profile")
	p.Layout = "auth/dashboard.html"
	p.TplName = "auth/profile.html"
}
