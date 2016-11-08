package auth

import (
	"fmt"
	"html/template"
	"time"

	machinery "github.com/RichardKnop/machinery/v1"
	log "github.com/Sirupsen/logrus"
	"github.com/kapmahc/lotus/web"
	"github.com/unrolled/render"
)

//Engine auth engine
type Engine struct {
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
		"df": func(t time.Time, l string) string {
			return t.Format(l)
		},
		"nl": func(name string) []web.Link {
			var links []web.Link
			if err := p.Dao.Get(fmt.Sprintf("nav-links.%s", name), &links); err != nil {
				log.Error(err)
			}
			return links
		},
	}
}

//Dashboard dashboard's nav-bar
func (p *Engine) Dashboard() []web.Dropdown {

	return []web.Dropdown{}
}
