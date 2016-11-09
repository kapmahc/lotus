package auth

import (
	"fmt"
	"html/template"
	"net/url"
	"time"

	machinery "github.com/RichardKnop/machinery/v1"
	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
	"github.com/kapmahc/lotus/web"
	"github.com/unrolled/render"
)

//Engine auth engine
type Engine struct {
	Router  *mux.Router       `inject:""`
	I18n    *web.I18n         `inject:""`
	Cache   *web.Cache        `inject:""`
	Hmac    *web.Hmac         `inject:""`
	Server  *machinery.Server `inject:""`
	Dao     *Dao              `inject:""`
	Handler *Handler          `inject:""`
	Render  *render.Render    `inject:""`
}

//FuncMap html template funcs
func (p *Engine) FuncMap() template.FuncMap {
	return template.FuncMap{
		"t": func(locale string, format string, args ...interface{}) string {
			return p.I18n.T(locale, format, args...)
		},
		"urlfor": func(locale, name string, pairs ...string) *url.URL {
			url, err := p.Router.Get(name).URL(pairs...)
			if err != nil {
				log.Error(err)
			}
			qry := url.Query()
			qry.Set("locale", locale)
			url.RawQuery = qry.Encode()
			return url
		},
		"languages": func() []string {
			items, err := p.I18n.Languages()
			if err != nil {
				log.Error(err)
			}
			return items
		},
		"date_fmt": func(t time.Time, l string) string {
			return t.Format(l)
		},
		"nav_links": func(name string) []web.Link {
			var links []web.Link
			if err := p.Dao.Get(fmt.Sprintf("nav-links.%s", name), &links); err != nil {
				log.Error(err)
			}
			return links
		},
		"assets_css": func(name string) template.HTML {
			return template.HTML(fmt.Sprintf(`<link rel="stylesheet" href="/css/%s.css">`, name))
		},
		"assets_js": func(name string) template.HTML {
			return template.HTML(fmt.Sprintf(`<script src="/js/%s.js"></script>`, name))
		},
	}
}

//Dashboard dashboard's nav-bar
func (p *Engine) Dashboard() []web.Dropdown {

	return []web.Dropdown{}
}
