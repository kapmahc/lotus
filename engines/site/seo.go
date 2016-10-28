package site

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
// @router /robots.txt [get]
func (p *Controller) GetRobots() {
	// TODO
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
