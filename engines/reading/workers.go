package reading

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/benmanns/goworker"
	"github.com/kapmahc/epub"
)

func scanBooksWorker(queue string, args ...interface{}) error {
	root := args[0].(string)
	beego.Info("scan books from", root)
	const ext = ".epub"
	count := 0
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
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
	})
	if err != nil {
		beego.Error(err)
	}
	return err
}

const scanBookJob = "ScanBooks"

func init() {
	goworker.Register(scanBookJob, scanBooksWorker)
}
