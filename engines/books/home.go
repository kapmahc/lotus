package books

import (
	"os"
	"path/filepath"

	"github.com/astaxie/beego"
	"github.com/kapmahc/epub"
)

//Controller books controller
type Controller struct {
	beego.Controller
}

//GetHome homepage
// @router / [get]
func (p *Controller) GetHome() {
	p.TplName = "books/index.html"
}

//GetScan scan books
// @router /scan [get]
func (p *Controller) GetScan() {
	// TODO admin?
	const ext = ".epub"
	if err := filepath.Walk(filepath.Join("tmp", "books"), func(path string, info os.FileInfo, err error) error {
		if info.Mode().IsRegular() && filepath.Ext(info.Name()) == ext {
			beego.Info("find file ", path)
		}
		return nil
	}); err != nil {
		beego.Error(err)
		p.Abort("500")
	}
	p.ServeJSON()
}

//GetBook show book index
// @router /:id:int/:name:string [get]
func (p *Controller) GetBook() {
	book, err := epub.Open("tmp/books/")
	if err != nil {
		beego.Error(err)
		p.Abort("500")
	}
	defer book.Close()
	beego.Debug(book.Files())
}
