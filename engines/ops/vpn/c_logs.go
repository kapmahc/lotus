package vpn

//IndexLog logs
// @router /logs [get]
func (p *Controller) IndexLog() {
	p.TplName = "ops/vpn/logs/index.html"
}
