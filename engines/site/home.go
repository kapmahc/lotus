package site

import (
	"github.com/astaxie/beego"
)

//Controller home controller
type Controller struct {
	beego.Controller
}

//GetHome homepage
// @router / [get]
func (p *Controller) GetHome() {
	p.TplName = "site/index.html"
}
