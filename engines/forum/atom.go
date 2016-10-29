package forum

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/gorilla/feeds"
	"github.com/kapmahc/lotus/engines/base"
)

func init() {
	base.RegisterAtom(func(home string, crl *beego.Controller) []*feeds.Item {
		var items []*feeds.Item
		var articles []Article
		o := orm.NewOrm()
		if _, err := o.QueryTable(new(Article)).
			OrderBy("-updated_at").
			Limit(12).
			All(&articles); err != nil {
			beego.Error(err)
		}
		for _, a := range articles {
			items = append(
				items, &feeds.Item{
					Title: a.Title,
					Link: &feeds.Link{
						Href: home + crl.URLFor("forum.Controller.ShowArticle", ":id", a.ID),
					},
					Description: a.Body,
					Created:     *a.UpdatedAt,
				},
			)
		}
		return items
	})
}
