package forum

import (
	"github.com/gin-gonic/gin"
	"github.com/kapmahc/lotus/web"
)

//Mount web points
func (p *Engine) Mount(rt *gin.Engine) {
	fg := rt.Group("/forum")

	fg.GET("/tags", web.JSON(p.tagsIndex))
	fg.POST("/tags",
		p.Jwt.CurrentUserHandler(true), p.Jwt.MustAdminHandler(),
		web.JSON(p.tagsCreate),
	)
	fg.GET("/tags/:id", web.JSON(p.tagsShow))
	fg.POST("/tags/:id",
		p.Jwt.CurrentUserHandler(true), p.Jwt.MustAdminHandler(),
		web.JSON(p.tagsUpdate),
	)
	fg.DELETE("/tags/:id",
		p.Jwt.CurrentUserHandler(true), p.Jwt.MustAdminHandler(),
		web.JSON(p.tagsDestroy))

	fg.GET("/articles", web.JSON(p.articlesIndex))
	fg.POST("/articles", p.Jwt.CurrentUserHandler(true), web.JSON(p.articlesCreate))
	fg.GET("/articles/:id", web.JSON(p.articlesShow))
	fg.POST("/articles/:id", p.Jwt.CurrentUserHandler(true), web.JSON(p.articlesUpdate))
	fg.DELETE("/articles/:id", p.Jwt.CurrentUserHandler(true), web.JSON(p.articlesDestroy))

	fg.GET("/comments", web.JSON(p.commentsIndex))
	fg.POST("/comments", p.Jwt.CurrentUserHandler(true), web.JSON(p.commentsCreate))
	fg.GET("/comments/:id", web.JSON(p.commentsShow))
	fg.POST("/comments/:id", p.Jwt.CurrentUserHandler(true), web.JSON(p.commentsUpdate))
	fg.DELETE("/comments/:id", p.Jwt.CurrentUserHandler(true), web.JSON(p.commentsDestroy))

}
