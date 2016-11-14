package reading

import (
	"github.com/gin-gonic/gin"
	"github.com/kapmahc/lotus/web"
)

//Mount web points
func (p *Engine) Mount(rt *gin.Engine) {
	rg := rt.Group("/reading")

	rg.GET("/books", web.JSON(p.booksIndex))
	rg.GET("/books/:id/*page", p.booksShow)
	rg.DELETE("/books/:id",
		p.Jwt.CurrentUserHandler(true), p.Jwt.MustAdminHandler(),
		web.JSON(p.booksDestroy),
	)

	rg.GET("/notes", web.JSON(p.notesIndex))
	rg.POST("/notes", p.Jwt.CurrentUserHandler(true), web.JSON(p.notesCreate))
	rg.GET("/notes/:id", web.JSON(p.notesShow))
	rg.POST("/notes/:id", p.Jwt.CurrentUserHandler(true), web.JSON(p.notesUpdate))
	rg.DELETE("/notes/:id", p.Jwt.CurrentUserHandler(true), web.JSON(p.notesDestroy))
}
