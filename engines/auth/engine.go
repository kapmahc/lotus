package auth

import (
	machinery "github.com/RichardKnop/machinery/v1"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/kapmahc/lotus/web"
)

//Engine auth engine
type Engine struct {
	I18n    *web.I18n         `inject:""`
	Cache   *web.Cache        `inject:""`
	Hmac    *web.Hmac         `inject:""`
	Server  *machinery.Server `inject:""`
	Dao     *Dao              `inject:""`
	Db      *gorm.DB          `inject:""`
	Handler *Handler          `inject:""`
}

//Home home
func (p *Engine) Home() gin.HandlerFunc {
	return func(c *gin.Context) {}
}

//Dashboard dashboard's nav-bar
func (p *Engine) Dashboard() []web.Dropdown {
	return []web.Dropdown{}
}
