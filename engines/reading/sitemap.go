package reading

import (
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/kapmahc/lotus/engines/base"
	"github.com/kapmahc/lotus/sitemap"
)

func init() {
	base.RegisterSitemap(func(home string, crl *beego.Controller) []sitemap.Item {
		now := time.Now()
		items := []sitemap.Item{
			{Link: home + crl.URLFor("reading.Controller.GetBooks"), Updated: now},
			{Link: home + crl.URLFor("reading.Controller.IndexNotes"), Updated: now},
		}

		var books []Book
		o := orm.NewOrm()
		if _, err := o.QueryTable(new(Book)).
			All(&books, "id", "updated_at"); err != nil {
			beego.Error(err)
		}

		for _, b := range books {
			items = append(
				items,
				sitemap.Item{
					Link:    home + crl.URLFor("reading.Controller.GetBookIndex", ":id", b.ID),
					Updated: *b.UpdatedAt,
				},
			)
		}
		return items
	})
}
