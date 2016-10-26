package reading

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/kapmahc/epub"
	uuid "github.com/satori/go.uuid"
)

//GetBooks list books
// @router /books [get]
func (p *Controller) GetBooks() {
	p.TplName = "reading/books.html"
}

//GetScan scan books
// @router /scan [get]
func (p *Controller) GetScan() {
	// TODO admin?
	const ext = ".epub"
	count := 0
	if err := filepath.Walk(filepath.Join("tmp", "books"), func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.Mode().IsRegular() && filepath.Ext(info.Name()) == ext {
			beego.Info("find file ", path)
			bk, err := epub.Open(path)
			if err != nil {
				return err
			}
			defer bk.Close()

			var ids []string
			for _, b := range bk.Opf.Metadata.Identifier {
				ids = append(ids, b.Data)
			}
			var authors []string
			for _, a := range bk.Opf.Metadata.Creator {
				authors = append(authors, a.Data)
			}
			var dates []string
			for _, d := range bk.Opf.Metadata.Date {
				dates = append(dates, d.Data)
			}

			var book Book
			o := orm.NewOrm()
			err = o.QueryTable(&book).Filter("file", path).One(&book)
			book.Type = bk.Mimetype
			book.Title = strings.Join(bk.Opf.Metadata.Title, "-")
			book.Publisher = strings.Join(bk.Opf.Metadata.Publisher, ",")
			book.Author = strings.Join(authors, "-")
			book.UID = uuid.NewV4().String()
			book.Lang = strings.Join(bk.Opf.Metadata.Language, ",")
			book.Subject = strings.Join(bk.Opf.Metadata.Subject, ",")
			book.Description = strings.Join(bk.Opf.Metadata.Description, ",")
			book.PublishedAt = strings.Join(dates, ",")

			if err == nil {
				_, err = o.Update(&book)
			} else if err == orm.ErrNoRows {
				book.File = path
				_, err = o.Insert(&book)
			}
			if err != nil {
				return err
			}
			count++
		}
		return nil
	}); err != nil {
		beego.Error(err)
		p.Abort("500")
	}
	p.Data["json"] = map[string]interface{}{"count": count}
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
