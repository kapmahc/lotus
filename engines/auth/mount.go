package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/kapmahc/lotus/web"
)

//Mount web points
func (p *Engine) Mount(rt *gin.Engine) {
	rt.GET("/locales/:lang", web.JSON(p.getLocales))
	rt.GET("/layout", p.getLayout)

	ugn := rt.Group("/users")
	ugn.POST("/sign-in", web.JSON(p.postUsersSignIn))
	ugn.POST("/sign-up", web.JSON(p.postUsersSignUp))
	ugn.GET("/confirm", web.Redirect(p.getUsersConfirm))
	ugn.POST("/confirm", web.JSON(p.postUsersConfirm))
	ugn.GET("/unlock", web.Redirect(p.getUsersUnlock))
	ugn.POST("/unlock", web.JSON(p.postUsersUnlock))
	ugn.POST("/forgot-password", web.JSON(p.postUsersForgotPassword))
	ugn.POST("/reset-password", web.JSON(p.postUsersResetPassword))

	ugm := rt.Group("/users", p.Jwt.CurrentUserHandler(true))
	ugm.DELETE("/sign-out", web.JSON(p.deleteUsersSignOut))
	ugm.GET("/logs", web.JSON(p.getUsersLogs))
	ugm.GET("/info", web.JSON(p.getUsersInfo))
	ugm.POST("/info", web.JSON(p.postUsersInfo))
	ugm.POST("/change-password", web.JSON(p.postUsersChangePassword))

	ag := rt.Group("/admin", p.Jwt.CurrentUserHandler(true), p.Jwt.MustAdminHandler())
	ag.GET("/base", web.JSON(p.getAdminBase))
	ag.POST("/base", web.JSON(p.postAdminBase))
	ag.GET("/author", web.JSON(p.getAdminAuthor))
	ag.POST("/author", web.JSON(p.postAdminAuthor))
}
