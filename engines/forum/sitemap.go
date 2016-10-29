package forum

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
			{Link: home + crl.URLFor("forum.Controller.IndexTags"), Updated: now},
			{Link: home + crl.URLFor("forum.Controller.IndexArticles"), Updated: now},
			{Link: home + crl.URLFor("forum.Controller.IndexComments"), Updated: now},
		}

		var tags []Tag
		var articles []Article
		o := orm.NewOrm()
		if _, err := o.QueryTable(new(Tag)).
			All(&tags, "id", "updated_at"); err != nil {
			beego.Error(err)
		}
		if _, err := o.QueryTable(new(Article)).
			All(&articles, "id", "updated_at"); err != nil {
			beego.Error(err)
		}
		for _, t := range tags {
			items = append(
				items,
				sitemap.Item{
					Link:    home + crl.URLFor("forum.Controller.ShowTag", ":id", t.ID),
					Updated: *t.UpdatedAt,
				},
			)
		}
		for _, a := range articles {
			items = append(
				items,
				sitemap.Item{
					Link:    home + crl.URLFor("forum.Controller.ShowArticle", ":id", a.ID),
					Updated: *a.UpdatedAt,
				},
			)
		}

		return items
	})
}
