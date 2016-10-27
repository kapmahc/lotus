package auth

import (
	"github.com/astaxie/beego"
	"github.com/kapmahc/lotus/engines/base"
)

//Controller auth controller
type Controller struct {
	base.Controller
}

//Prepare prepare
func (p *Controller) Prepare() {
	beego.ReadFromRequest(&p.Controller.Controller)
	p.SetLocale()
}
