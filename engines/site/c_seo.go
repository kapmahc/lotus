package site

import (
	"encoding/xml"
	"text/template"

	"github.com/astaxie/beego"
	"github.com/beego/i18n"
	"github.com/kapmahc/lotus/engines/base"
)

//GetBaidu baidu verify file
// @router /baidu_verify_:id([\w]+).html [get]
func (p *Controller) GetBaidu() {
	var code string
	base.Get("baidu.verify.code", &code)
	if code != p.Ctx.Input.Param(":id") {
		p.Abort("404")
	}
	p.Data["code"] = code
	p.Layout = ""
	p.TplName = "site/baidu.html"
}

//GetGoogle google verify file
// @router /google:id([\w]+).html [get]
func (p *Controller) GetGoogle() {
	var code string
	base.Get("google.verify.code", &code)
	if code != p.Ctx.Input.Param(":id") {
		p.Abort("404")
	}
	p.Data["code"] = code
	p.Layout = ""
	p.TplName = "site/google.html"
}

//GetRobots robots.txt
//See http://www.robotstxt.org/robotstxt.html for documentation on how to use the robots.txt file
// @router /robots.txt [get]
func (p *Controller) GetRobots() {
	txt := `
User-agent: *
Disallow:

SITEMAP: {{.Home}}{{.Href}}
`
	tpl, err := template.New("").Parse(txt)
	if err == nil {
		err = tpl.Execute(p.Ctx.ResponseWriter, struct {
			Home string
			Href string
		}{
			Home: beego.AppConfig.String("homeurl"),
			Href: p.URLFor("site.Controller.GetSitemap"),
		})
	}
	p.Check(err)
}

//GetSitemap sitemap.xml
// @router /sitemap.xml [get]
func (p *Controller) GetSitemap() {
	sm := base.SitemapXML(&p.Controller.Controller)

	wrt := p.Ctx.ResponseWriter
	wrt.Write([]byte(xml.Header))
	en := xml.NewEncoder(wrt)
	//en.Indent("", "  ")
	err := en.Encode(sm)
	p.Check(err)
}

//GetRss rss.atom
// @router /rss/:lang([-\w]+).atom [get]
func (p *Controller) GetRss() {
	lang := p.Ctx.Input.Param(":lang")
	if !i18n.IsExist(lang) {
		p.Abort("404")
	}
	err := base.
		Atom(lang, &p.Controller.Controller).
		WriteAtom(p.Ctx.ResponseWriter)
	p.Check(err)
}
