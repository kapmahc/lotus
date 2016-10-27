package reading

import "github.com/kapmahc/lotus/engines/base"

//Controller books controller
type Controller struct {
	base.Controller
}

//Prepare prepare
func (p *Controller) Prepare() {
	p.SetLocale()
}
