package reading

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

//DestroyBook destroy book
// @router /books/:id [delete]
func (p *Controller) DestroyBook() {
	book := p.getBook()
	_, err := orm.NewOrm().Delete(&book)
	p.Check(err)
	fl := beego.NewFlash()
	fl.Notice(p.T("site-pages.success"))
	p.Redirect(fl, "reading.Controller.ListBooks")
}

//ListBooks list books
// @router /scan [get]
func (p *Controller) ListBooks() {
	p.Dashboard()
	p.MustAdmin()
	var books []Book
	_, err := orm.NewOrm().
		QueryTable(new(Book)).
		OrderBy("-file").
		All(&books, "id", "title", "file")
	p.Check(err)
	p.Data["title"] = p.T("reading-pages.books")
	p.Data["books"] = books
	p.TplName = "reading/books/manage.html"
}

//ScanBooks scan books
// @router /scan [get]
func (p *Controller) ScanBooks() {
	p.MustAdmin()
	// goworker.Enqueue(&goworker.Job{
	// 	Queue: base.QueueLow,
	// 	Payload: goworker.Payload{
	// 		Class: scanBookJob,
	// 		Args:  []interface{}{filepath.Join("tmp", "books")},
	// 	},
	// })
	fl := beego.NewFlash()
	fl.Notice(p.T("reading-pages.scan-run-on-back"))
	p.Redirect(fl, "reading.Controller.ListBooks")
}
