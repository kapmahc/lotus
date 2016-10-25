package routers

import (
	"github.com/astaxie/beego"
	"github.com/kapmahc/lotus/engines/auth"
	"github.com/kapmahc/lotus/engines/books"
	"github.com/kapmahc/lotus/engines/forum"
	"github.com/kapmahc/lotus/engines/ops"
	"github.com/kapmahc/lotus/engines/shop"
	"github.com/kapmahc/lotus/engines/site"
)

func init() {
	beego.Include(&site.Controller{})
	beego.AddNamespace(
		beego.NewNamespace("/users", beego.NSInclude(&auth.Controller{})),
		beego.NewNamespace("/forum", beego.NSInclude(&forum.Controller{})),
		beego.NewNamespace("/books", beego.NSInclude(&books.Controller{})),
		beego.NewNamespace("/shop", beego.NSInclude(&shop.Controller{})),
		beego.NewNamespace("/ops", beego.NSInclude(&ops.Controller{})),
	)
}
