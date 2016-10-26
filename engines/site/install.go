package site

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/kapmahc/lotus/engines/auth"
)

//GetInstall install
// @router /install [get]
func (p *Controller) GetInstall() {
	p.mustEmptyDb()
	p.Data["form"] = &fmInstall{}
	p.TplName = "site/install.html"
}

//PostInstall install
// @router /install [post]
func (p *Controller) PostInstall() {
	p.mustEmptyDb()
}

func (p *Controller) mustEmptyDb() {
	o := orm.NewOrm()
	ct, er := o.QueryTable(new(auth.User)).Count()
	if er != nil {
		beego.Error(er)
		p.Abort("500")
	}
	if ct > 0 {
		p.Abort("404")
	}
}
