package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/kapmahc/lotus/web"
)

//Mount web points
func (p *Engine) Mount(rt *gin.Engine) {
	rt.GET("/locales/:lang", web.JSON(p.getLocales))
	rt.GET("/layout", p.getLayout)
	rt.GET("/dashboard", p.getDashboard)

	ug := rt.Group("/users")
	ug.POST("/sign-in", web.JSON(p.postUsersSignIn))
	ug.POST("/sign-up", web.JSON(p.postUsersSignUp))
	ug.GET("/confirm", web.Redirect(p.getUsersConfirm))
	ug.POST("/confirm", web.JSON(p.postUsersConfirm))
	ug.GET("/unlock", web.Redirect(p.getUsersUnlock))
	ug.POST("/unlock", web.JSON(p.postUsersUnlock))
	ug.POST("/forgot-password", web.JSON(p.postUsersForgotPassword))
	ug.POST("/reset-password", web.JSON(p.postUsersResetPassword))
	ug.DELETE("/sign-out", p.Jwt.CurrentUserHandler(true), web.JSON(p.deleteSignOut))
}
