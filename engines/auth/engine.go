package auth

import (
	"html/template"

	machinery "github.com/RichardKnop/machinery/v1"
	"github.com/kapmahc/lotus/web"
)

//Engine auth engine
type Engine struct {
	I18n      *web.I18n         `inject:""`
	Cache     *web.Cache        `inject:""`
	Logger    *web.Logger       `inject:""`
	JobLogger *web.JobLogger    `inject:""`
	Hmac      *web.Hmac         `inject:""`
	Server    *machinery.Server `inject:""`
	Dao       *Dao              `inject:""`
	Handler   *Handler          `inject:""`
}

//FuncMap html template funcs
func (p *Engine) FuncMap() template.FuncMap {
	return template.FuncMap{}
}

//Dashboard dashboard's nav-bar
func (p *Engine) Dashboard() []web.Dropdown {
	return []web.Dropdown{}
}
