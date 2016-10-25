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
	beego.Router("/", &controllers.MainController{})
	beego.Router("/users", &auth.Controller{})
	beego.Router("/forum", &forum.Controller{})
	beego.Router("/shop", &shop.Controller{})
	beego.Router("/ops", &ops.Controller{})
}
