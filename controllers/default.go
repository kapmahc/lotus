package controllers

import (
	"github.com/astaxie/beego"
)

//Controller home controller
type Controller struct {
	beego.Controller
}

//Get /
func (c *Controller) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}
