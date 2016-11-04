package auth

import "github.com/gin-gonic/gin"

//Mount web points
func (p *Engine) Mount(rt *gin.Engine) {
	rt.GET("/locales/:lang", p.getLocales)

	utp := rt.Group("/users", p.Handler.Layout, p.Handler.Form)
	utp.GET("/sign-in", p.getUsersSignIn)
}
