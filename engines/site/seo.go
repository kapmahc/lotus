package site

//GetBaidu baidu verify file
// @router /baidu_verify_:id([\w]+).html [get]
func (p *Controller) GetBaidu() {
	// TODO
	p.Data["code"] = p.Ctx.Input.Param(":id")
	p.TplName = "site/baidu.html"
}

//GetGoogle google verify file
// @router /google:id([\w]+).html [get]
func (p *Controller) GetGoogle() {
	// TODO
	p.Data["code"] = p.Ctx.Input.Param(":id")
	p.TplName = "site/google.html"
}

//GetRobots robots.txt
// @router /sitemap.xml.gz [get]
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
