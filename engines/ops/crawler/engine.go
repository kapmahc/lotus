package crawler

import (
	"github.com/kapmahc/lotus/engines/auth"
	"github.com/kapmahc/lotus/web"
)

//Engine crawler engine
type Engine struct {
	I18n    *web.I18n     `inject:""`
	Logger  *web.Logger   `inject:""`
	Handler *auth.Handler `inject:""`
}

func init() {

}
