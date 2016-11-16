package shop

import (
	"github.com/gin-gonic/gin"
	"github.com/kapmahc/lotus/web"
)

//Mount web points
func (p *Engine) Mount(rt *gin.Engine) {
	sg := rt.Group("/shop")
	sg.GET("/countries", web.JSON(p.countriesIndex))
	sg.POST("/countries",
		p.Jwt.CurrentUserHandler(true), p.Jwt.MustAdminHandler(),
		web.JSON(p.countriesCreate),
	)
	sg.GET("/countries/:id", web.JSON(p.countriesShow))
	sg.POST("/countries/:id",
		p.Jwt.CurrentUserHandler(true), p.Jwt.MustAdminHandler(),
		web.JSON(p.countriesUpdate),
	)
	sg.DELETE("/countries/:id", web.JSON(p.countriesDestroy))
}
