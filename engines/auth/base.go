package auth

import (
	"github.com/astaxie/beego"
	"github.com/kapmahc/lotus/engines/base"
)

//BaseController base controller
type BaseController struct {
	base.Controller
}

//Prepare prepare
func (p *BaseController) Prepare() {
	beego.ReadFromRequest(&p.Controller.Controller)
	p.SetLocale()
	p.CurrentUser()
}

//CurrentUser get current user
func (p *BaseController) CurrentUser() *User {
	uid := p.GetSession("uid")
	if uid == nil {
		return nil
	}
	user, err := GetUserByUID(uid.(string))
	if err == nil {
		p.Data["currentUser"] = struct {
			Name string
			Logo string
		}{
			Name: user.Name,
			Logo: user.Logo,
		}
		return user
	}
	beego.Error(err)
	p.DestroySession()
	return nil
}
