package auth

import "github.com/gin-gonic/gin"

//Mount mount web point
func (p *Engine) Mount(rt *gin.Engine) {
	rt.GET("/siteInfo", p.getSiteInfo)
	rt.GET("/locales/:lang", p.getLocales)
	ug := rt.Group("/users")
	ug.POST("signIn", p.postUsersSignIn)
}
