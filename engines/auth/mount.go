package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/kapmahc/lotus/web"
)

//Mount mount web point
func (p *Engine) Mount(rt *gin.Engine) {
	rt.GET("/siteInfo", p.getSiteInfo)
	rt.GET("/locales/:lang", p.getLocales)

	ug := rt.Group("/users")
	ug.POST("signIn", p.postUsersSignIn)
	ug.POST("signUp", web.JSON(p.postUsersSignUp))
}
