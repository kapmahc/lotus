package shop

import "github.com/astaxie/beego"

//Controller shop controller
type Controller struct {
	beego.Controller
}

//GetHome homepage
// @router / [get]
func (p *Controller) GetHome() {
	p.TplName = "shop/index.html"
}
