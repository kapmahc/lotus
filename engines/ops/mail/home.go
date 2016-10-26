package mail

import "github.com/astaxie/beego"

//Controller ops controller
type Controller struct {
	beego.Controller
}

//GetHome homepage
// @router / [get]
func (p *Controller) GetHome() {
	p.TplName = "ops/mail/index.html"
}