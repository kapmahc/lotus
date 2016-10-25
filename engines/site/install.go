package site

//GetInstall install
// @router /install [get]
func (p *Controller) GetInstall() {
	p.TplName = "site/install.html"
}

//PostInstall install
// @router /install [post]
func (p *Controller) PostInstall() {
}
