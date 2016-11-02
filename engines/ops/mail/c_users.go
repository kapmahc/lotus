package mail

//IndexUser users
// @router /users [get]
func (p *Controller) IndexUser() {
	p.TplName = "ops/mail/users/index.html"
}
