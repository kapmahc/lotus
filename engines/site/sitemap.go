package site

import (
	"time"

	"github.com/astaxie/beego"
	"github.com/kapmahc/lotus/engines/base"
	"github.com/kapmahc/lotus/sitemap"
)

func init() {
	base.RegisterSitemap(func(home string, crl *beego.Controller) []sitemap.Item {
		var installed time.Time
		base.Get("site.install", &installed)
		return []sitemap.Item{
			{Link: home, Updated: time.Now()},
			{Link: home + crl.URLFor("site.Controller.IndexNotice"), Updated: time.Now()},
			{Link: home + crl.URLFor("site.Controller.NewLeaveword"), Updated: installed},
		}
	})
}
