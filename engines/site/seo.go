package site

import (
	"text/template"

	"github.com/astaxie/beego"
)

//GetBaidu baidu verify file
// @router /baidu_verify_:id([\w]+).html [get]
func (p *Controller) GetBaidu() {
	var code string
	Get("baidu.verify.code", &code)
	if code != p.Ctx.Input.Param(":id") {
		p.Abort("404")
	}
	p.Data["code"] = code
	p.TplName = "site/baidu.html"
}

//GetGoogle google verify file
// @router /google:id([\w]+).html [get]
func (p *Controller) GetGoogle() {
	var code string
	Get("google.verify.code", &code)
	if code != p.Ctx.Input.Param(":id") {
		p.Abort("404")
	}
	p.Data["code"] = code
	p.TplName = "site/google.html"
}

//GetRobots robots.txt
//See http://www.robotstxt.org/robotstxt.html for documentation on how to use the robots.txt file
// @router /robots.txt [get]
func (p *Controller) GetRobots() {
	txt := `
User-agent: *
Disallow:

SITEMAP: {{.Home}}/sitemap.xml.gz
`
	tpl, err := template.New("").Parse(txt)
	if err == nil {
		err = tpl.Execute(p.Ctx.ResponseWriter, struct {
			Home string
		}{
			Home: beego.AppConfig.String("homeurl"),
		})
	}
	if err != nil {
		beego.Error(err)
		p.Abort("500")
	}
}

//GetSitemap sitemap.xml.gz
// @router /sitemap.xml.gz [get]
func (p *Controller) GetSitemap() {
	// TODO
}

//GetRss rss.atom
// @router /rss.atom [get]
func (p *Controller) GetRss() {
	// TODO
}
