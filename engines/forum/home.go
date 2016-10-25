package forum

import "github.com/astaxie/beego"

//Controller forum controller
type Controller struct {
	beego.Controller
}

//GetHome homepage
// @router / [get]
func (p *Controller) GetHome() {
	p.TplName = "forum/index.html"
}
