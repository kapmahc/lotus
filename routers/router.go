package routers

import (
	"github.com/astaxie/beego"
	"github.com/kapmahc/lotus/controllers"
	"github.com/kapmahc/lotus/controllers/auth"
	"github.com/kapmahc/lotus/controllers/forum"
	"github.com/kapmahc/lotus/controllers/ops"
	"github.com/kapmahc/lotus/controllers/shop"
)

func init() {
	beego.Include(&controllers.Controller{})
	beego.AddNamespace(
		beego.NewNamespace("/users", beego.NSInclude(&auth.Controller{})),
		beego.NewNamespace("/forum", beego.NSInclude(&forum.Controller{})),
		beego.NewNamespace("/shop", beego.NSInclude(&shop.Controller{})),
		beego.NewNamespace("/ops", beego.NSInclude(&ops.Controller{})),
	)
}
