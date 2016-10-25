package routers

import (
	"github.com/kapmahc/lotus/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
