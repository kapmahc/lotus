package site

import (
	"html/template"
	"path"

	"github.com/astaxie/beego"
	"github.com/kapmahc/lotus/engines/base"
)

//GetAdminCfgFiles config files
// @router /admin/cfg-files [get]
func (p *Controller) GetAdminCfgFiles() {
	p.Dashboard()
	p.MustAdmin()
	p.Data["links"] = []base.Link{
		{Href: p.URLFor("site.Controller.GetNginxConf"), Label: "nginx.conf"},
	}
	p.Data["title"] = p.T("site-pages.admin-cfg-files")
	p.TplName = "site/admin/cfg-files.html"
}

//GetNginxConf nginx.conf
// @router /admin/nginx.conf [get]
func (p *Controller) GetNginxConf() {
	tpl, err := template.ParseFiles(path.Join("views", "nginx.conf"))
	p.Check(err)
	port, err := beego.AppConfig.Int("httpport")
	p.Check(err)
	if err == nil {
		err = tpl.Execute(p.Ctx.ResponseWriter, struct {
			Name string
			Port int
			Env  string
		}{
			Name: beego.AppConfig.String("servername"),
			Port: port,
			Env:  beego.AppConfig.String("runmode"),
		})
	}
	p.Check(err)
}
