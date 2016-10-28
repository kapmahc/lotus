package auth

//Controller auth controller
type Controller struct {
	BaseController
}

//Prepare prepare
func (p *Controller) Prepare() {
	p.BaseController.Prepare()
	p.Layout = ""
}
