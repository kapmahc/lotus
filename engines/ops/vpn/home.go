package vpn

import "github.com/kapmahc/lotus/engines/auth"

//Controller ops controller
type Controller struct {
	auth.BaseController
}

//Prepare prepare
func (p *Controller) Prepare() {
	p.BaseController.Prepare()
	p.Dashboard()
	p.MustAdmin()
}
