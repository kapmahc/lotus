package auth

import "github.com/gin-gonic/gin"

//Mount web points
func (p *Engine) Mount(rt *gin.Engine) {
	rt.GET("/locales/:lang", p.getLocales)
	rt.GET("/layout", p.getLayout)
	rt.GET("/dashboard", p.Handler.MustSignIn(), p.getDashboard)
}
