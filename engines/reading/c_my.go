package reading

import "github.com/astaxie/beego/orm"

//GetMyNotes my notes
// @router /my-notes [get]
func (p *Controller) GetMyNotes() {
	p.Dashboard()
	var notes []Note
	o := orm.NewOrm().QueryTable(new(Note))
	if !p.IsAdmin() {
		o = o.Filter("user_id", p.CurrentUser().ID)
	}
	_, err := o.OrderBy("-updated_at").All(&notes, "id", "body", "updated_at")
	p.Check(err)
	p.Data["title"] = p.T("reading-pages.my-notes")
	p.Data["articles"] = notes
	p.TplName = "reading/notes/my.html"
}
