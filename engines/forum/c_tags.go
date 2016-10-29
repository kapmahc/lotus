package forum

import "github.com/astaxie/beego/orm"

//GetAdminTags tags
// @router /admin/tags [get]
func (p *Controller) GetAdminTags() {
	p.Dashboard()
	p.MustAdmin()
	var tags []Tag
	_, err := orm.NewOrm().QueryTable(new(Tag)).All(&tags, "id", "name")
	p.Check(err)
	p.Data["tags"] = tags
	p.Data["title"] = p.T("forum-pages.tags")
	p.TplName = "forum/tags/index.html"
}
