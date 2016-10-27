package auth

import (
	"github.com/astaxie/beego"
	"github.com/kapmahc/lotus/engines/base"
)

//CurrentUser current user
const CurrentUser = "currentUser"

//BaseController base controller
type BaseController struct {
	base.Controller
}

//Prepare prepare
func (p *BaseController) Prepare() {
	beego.ReadFromRequest(&p.Controller.Controller)
	p.SetLocale()
	p.SetCurrentUser()
}

//Dashboard prepare dashboard
func (p *BaseController) Dashboard() {
	p.MustSignIn()
	p.Layout = "auth/dashboard.html"
}

//MustSignIn must sign in
func (p *BaseController) MustSignIn() {
	if p.Data[CurrentUser] == nil {
		p.Abort("402")
	}
}

//CurrentUser get current user
func (p *BaseController) CurrentUser() *User {
	user := p.Data[CurrentUser]
	return user.(*User)
}

//SetCurrentUser set current user
func (p *BaseController) SetCurrentUser() {
	uid := p.GetSession("uid")
	if uid == nil {
		return
	}
	user, err := GetUserByUID(uid.(string))
	if err != nil {
		beego.Error(err)
		p.DestroySession()
		return
	}
	p.Data[CurrentUser] = user
}
