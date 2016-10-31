package forum

import (
	"github.com/astaxie/beego/orm"
	"github.com/kapmahc/lotus/engines/auth"
)

//IndexTags tags
// @router /tags [get]
func (p *Controller) IndexTags() {
	var tags []Tag
	_, err := orm.NewOrm().QueryTable(new(Tag)).All(&tags, "id", "name")
	p.Check(err)
	p.Data["tags"] = tags
	p.Data["title"] = p.T("forum-pages.tags")
	p.Data["can"] = p.CurrentUser().Has(auth.AdminRole)
	p.TplName = "forum/tags/index.html"
}
