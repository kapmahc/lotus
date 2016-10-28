package reading

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"path"
	"path/filepath"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/benmanns/goworker"
	"github.com/kapmahc/epub"
	"github.com/kapmahc/lotus/engines/base"
)

//GetBooks list books
// @router /books [get]
func (p *Controller) GetBooks() {
	o := orm.NewOrm()

	const size = 60
	page, _ := p.GetInt64("page", 1)
	count, err := o.QueryTable(new(Book)).Count()
	p.Check(err)

	var books []Book
	_, err = o.QueryTable(new(Book)).
		OrderBy("-vote").
		Offset((page - 1) * size).
		Limit(size).
		All(&books)
	p.Check(err)

	p.Data["pager"] = base.NewPaginator(
		p.URLFor("reading.Controller.GetBooks"),
		count,
		page,
		size,
		books,
	)
	p.Data["title"] = p.T("reading-pages.books")
	p.TplName = "reading/books.html"
}

//GetScan scan books
// @router /scan [get]
func (p *Controller) GetScan() {
	p.MustAdmin()
	goworker.Enqueue(&goworker.Job{
		Queue: base.QueueLow,
		Payload: goworker.Payload{
			Class: scanBookJob,
			Args:  []interface{}{filepath.Join("tmp", "books")},
		},
	})
	p.Data["json"] = map[string]interface{}{"ok": true}
	p.ServeJSON()
}

func (p *Controller) getBook() *Book {
	var book Book
	err := orm.NewOrm().
		QueryTable(&book).
		Filter("id", p.Ctx.Input.Param(":id")).
		One(&book)
	p.Check(err)
	return &book
}

func (p *Controller) points2html(href string, points []epub.NavPoint) string {
	str := "<ol>"
	for _, pt := range points {
		str += fmt.Sprintf(
			`<li><a href="%s/%s" target="_blank">%s</a></li>`,
			href,
			pt.Content.Src,
			pt.Text,
		)
		str += p.points2html(href, pt.Points)
	}
	str += "</ol>"
	return str
}

//GetBookIndex show book index
// @router /books/:id [get]
func (p *Controller) GetBookIndex() {
	book := p.getBook()
	p.Data["title"] = book.Title
	bk, err := epub.Open(book.File)
	p.Check(err)
	defer bk.Close()

	p.Data["ncx"] = template.HTML(
		p.points2html(
			p.URLFor("reading.Controller.GetBookIndex", ":id", book.ID),
			bk.Ncx.Points,
		),
	)
	p.TplName = "reading/book.html"
}

//GetBook show book page
// @router /books/:id/* [get]
func (p *Controller) GetBook() {
	book := p.getBook()
	bk, err := epub.Open(book.File)
	p.Check(err)
	defer bk.Close()
	name := p.Ctx.Input.Param(":splat")

	// switch path.Ext(name) {
	// case ".css":
	// 	p.Ctx.Output.Header("Content-Type", "text/css; charset=utf-8")
	// 	p.Ctx.Output.Body(buf)
	// case ".xhtml":
	// 	p.Ctx.Output.Header("Content-Type", "application/xhtml+xml; charset=utf-8")
	// 	p.Ctx.Output.Body(buf)
	// default:
	// 	beego.Error("bad file", name)
	// 	p.Abort("404")
	// }

	fd, err := bk.Open(name)
	p.Check(err)
	defer fd.Close()
	if path.Ext(name) == ".xhtml" {
		p.Data["title"], p.Data["body"] = p.parseHTML(fd)
		p.TplName = "reading/page.html"
		return
	}
	for _, ext := range []string{".css", ".jpg"} {
		if path.Ext(name) == ext {
			for _, m := range bk.Opf.Manifest {
				if strings.HasPrefix(name, m.Href) {
					// p.Ctx.Output.ContentType(m.MediaType)
					buf, err := ioutil.ReadAll(fd)
					p.Check(err)

					p.Ctx.Output.Header("Content-Type", m.MediaType)
					p.Ctx.Output.Body(buf)
					return
				}
			}
		}
	}
	beego.Error("bad file", name)
	p.Abort("404")

}