package vpn

import "github.com/kapmahc/lotus/engines/base"

//ReadMe read-me
// @router /read-me [get]
func (p *Controller) ReadMe() {
	p.Data["title"] = p.T("ops-vpn-pages.read-me")
	p.Data["links"] = []base.Link{}
	p.TplName = "ops/vpn/read-me.html"
}
