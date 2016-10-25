package forum

//GetArticles articles list
// @router /articles [get]
func (p *Controller) GetArticles() {
	p.TplName = "cms/articles/index.html"
}
