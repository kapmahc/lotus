package forum

import "github.com/kapmahc/lotus/engines/auth"

//Controller forum controller
type Controller struct {
	auth.BaseController
}

//GetHome homepage
// @router / [get]
func (p *Controller) GetHome() {
	p.TplName = "forum/index.html"
}
