package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/kapmahc/lotus/web"
)

//Mount mount web point
func (p *Engine) Mount(rt *gin.Engine) {
	rt.GET("/site/info", p.getSiteInfo)
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

	sg := rt.Group("/self", p.Jwt.CurrentUserHandler(true))
	sg.GET("logs", web.JSON(p.getSelfLogs))
	sg.GET("profile", web.JSON(p.getSelfProfile))
	sg.POST("profile", web.JSON(p.postSelfProfile))
	sg.POST("password", web.JSON(p.postSelfPassword))

	ag := rt.Group("/site", p.Jwt.CurrentUserHandler(true), p.Jwt.MustAdminHandler())
	ag.GET("status", web.JSON(p.getSiteStatus))
	ag.GET("base", web.JSON(p.getSiteBase))
	ag.POST("base", web.JSON(p.postSiteBase))
	ag.GET("author", web.JSON(p.getSiteAuthor))
	ag.POST("author", web.JSON(p.postSiteAuthor))
	ag.GET("nav", web.JSON(p.getSiteNav))
	ag.POST("nav", web.JSON(p.postSiteNav))
	ag.GET("seo", web.JSON(p.getSiteSeo))
	ag.POST("seo", web.JSON(p.postSiteSeo))
	ag.GET("users", web.JSON(p.getSiteUsers))
}
