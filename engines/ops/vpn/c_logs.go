package vpn

import "github.com/astaxie/beego/orm"

//IndexLog logs
// @router /logs [get]
func (p *Controller) IndexLog() {
	var logs []Log
	_, err := orm.NewOrm().
		QueryTable(new(Log)).
		RelatedSel().
		OrderBy("-id").
		Limit(120).
		All(&logs)
	p.Check(err)
	p.Data["logs"] = logs
	p.Data["title"] = p.T("ops-vpn-pages.logs")
	p.TplName = "ops/vpn/logs/index.html"
}
