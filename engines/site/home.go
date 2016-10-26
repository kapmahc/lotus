package site

import "github.com/kapmahc/lotus/engines/base"

//Controller home controller
type Controller struct {
	base.Controller
}

//Prepare prepare
func (p *Controller) Prepare() {
	p.SetLocale()
}

//GetHome homepage
// @router / [get]
func (p *Controller) GetHome() {

	p.TplName = "site/index.html"
}
