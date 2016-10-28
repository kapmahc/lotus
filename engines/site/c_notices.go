package site

import "github.com/astaxie/beego/orm"

//IndexNotices notices
// @router /notices [get]
func (p *Controller) IndexNotices() {
	var items []Notice
	_, err := orm.NewOrm().
		QueryTable(new(Notice)).
		OrderBy("-id").
		Limit(60).
		All(&items)
	p.Check(err)
	p.Data["notices"] = items
	p.Layout = "layout.html"
}
