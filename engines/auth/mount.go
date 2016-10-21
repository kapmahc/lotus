package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/kapmahc/lotus/web"
)

//Mount mount web point
func (p *Engine) Mount(rt *gin.Engine) {
	rt.GET("/site-info", p.getSiteInfo)
	rt.GET("/locales/:lang", p.getLocales)

	ug := rt.Group("/users")
	ug.POST("sign-in", web.JSON(p.postUsersSignIn))
	ug.POST("sign-up", web.JSON(p.postUsersSignUp))
	ug.GET("confirm", web.Redirect(p.getUsersConfirm))
	ug.POST("confirm", web.JSON(p.postUsersConfirm))
	ug.GET("unlock", web.Redirect(p.getUsersUnlock))
	ug.POST("unlock", web.JSON(p.postUsersUnlock))
	ug.POST("forgot-password", web.JSON(p.postUsersForgotPassword))
	ug.POST("change-password", web.JSON(p.postUsersChangePassword))
	ug.DELETE("sign-out", p.Jwt.CurrentUserHandler(true), p.deleteSignOut)
	ug.GET("logs", p.Jwt.CurrentUserHandler(true), web.JSON(p.getUserLogs))
}
