package base

import (
	"github.com/astaxie/beego"
	"github.com/kapmahc/lotus/sitemap"
)

//SitemapFunc sitemap items
type SitemapFunc func(string, *beego.Controller) []sitemap.Item

var sitemapHandlers []SitemapFunc

//RegisterSitemap registe sitemap handler
func RegisterSitemap(args ...SitemapFunc) {
	sitemapHandlers = append(sitemapHandlers, args...)
}

//SitemapXML sitemap.xml
func SitemapXML(crl *beego.Controller) *sitemap.Sitemap {
	home := beego.AppConfig.String("homeurl")
	sm := sitemap.New()
	for _, fn := range sitemapHandlers {
		items := fn(home, crl)
		sm.Items = append(sm.Items, items...)
	}
	return sm
}
