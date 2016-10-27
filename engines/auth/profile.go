package auth

//GetInfo info page
// @router /info [get]
func (p *Controller) GetInfo() {
	p.Dashboard()
	p.Data["title"] = p.T("auth-pages.profile")
	p.TplName = "auth/info.html"
}

//GetChangePassword change-password page
// @router /change-password [get]
func (p *Controller) GetChangePassword() {
	p.Dashboard()
	p.Data["title"] = p.T("auth-pages.change-password")
	p.TplName = "auth/change-password.html"
}

//GetLogs logs page
// @router /logs [get]
func (p *Controller) GetLogs() {
	p.Dashboard()
	p.Data["title"] = p.T("auth-pages.logs")
	p.TplName = "auth/logs.html"
}
