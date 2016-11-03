package mail

import "github.com/kapmahc/lotus/engines/base"

//ReadMe read-me
// @router /read-me [get]
func (p *Controller) ReadMe() {
	p.Data["title"] = p.T("ops-mail-pages.read-me")
	p.Data["links"] = []base.Link{}
	p.TplName = "ops/mail/read-me.html"
}
