package reading

import (
	"path/filepath"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/benmanns/goworker"
	"github.com/kapmahc/lotus/engines/base"
)

//DestroyBook destroy book
// @router /books/:id [delete]
func (p *Controller) DestroyBook() {
	book := p.getBook()
	_, err := orm.NewOrm().Delete(&book)
	p.Check(err)
	p.Data["json"] = map[string]string{
		"to": p.URLFor("reading.Controller.ListBooks"),
	}
	p.ServeJSON()
}

//ListBooks list books
// @router /admin/books [get]
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
	goworker.Enqueue(&goworker.Job{
		Queue: base.QueueLow,
		Payload: goworker.Payload{
			Class: scanBookJob,
			Args:  []interface{}{filepath.Join("tmp", "books")},
		},
	})
	fl := beego.NewFlash()
	fl.Notice(p.T("reading-logs.scan-run-on-back"))
	p.Redirect(fl, "reading.Controller.ListBooks")
}
