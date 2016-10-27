package site

import "github.com/kapmahc/lotus/engines/auth"

//Controller home controller
type Controller struct {
	auth.BaseController
}

//GetHome homepage
// @router / [get]
func (p *Controller) GetHome() {

	p.TplName = "site/index.html"
}
