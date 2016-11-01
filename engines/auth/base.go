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
	p.Data["xsrf_token"] = p.XSRFToken()
	p.Layout = "layout.html"
}

//Dashboard prepare dashboard
func (p *BaseController) Dashboard() {
	p.MustSignIn()
	user := p.CurrentUser()
	var links []*base.Dropdown
	for _, fn := range dashboard {
		lk := fn(user)
		if lk != nil {
			links = append(links, lk)
		}
	}
	p.Data["navBar"] = links
	p.Layout = "auth/dashboard.html"
}

//IsSignIn is sign in ?
func (p *BaseController) IsSignIn() bool {
	return p.Data[CurrentUser] != nil
}

//MustSignIn must sign in
func (p *BaseController) MustSignIn() {
	if !p.IsSignIn() {
		p.Abort("403")
	}
}

//IsAdmin is admin ?
func (p *BaseController) IsAdmin() bool {
	return p.IsSignIn() && p.CurrentUser().Has(AdminRole)
}

//MustAdmin must admin
func (p *BaseController) MustAdmin() {
	if !p.IsAdmin() {
		p.Abort("403")
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
	if err == nil {
		p.Data["is_admin"] = user.Has(AdminRole)
	} else {
		beego.Error(err)
		p.DestroySession()
		return
	}
	p.Data[CurrentUser] = user
}
