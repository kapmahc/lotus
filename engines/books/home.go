package books

import "github.com/astaxie/beego"

//Controller books controller
type Controller struct {
	beego.Controller
}

//GetHome homepage
// @router / [get]
func (p *Controller) GetHome() {
	p.TplName = "books/index.html"
}
