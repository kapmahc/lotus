package crawler

import "github.com/gorilla/mux"

//Mount web points
func (p *Engine) Mount(rt *mux.Router) {
	// rgl := rt.Group("/ops/crawler/line", p.lineCurrentUser)
	// rgl.POST("/sign-in", p.postLineSignIn)
	// rgl.POST("/sign-out", p.lineMustSignIn, p.getLineSignOut)
	// rgl.POST("/pull", p.lineMustSignIn, p.postLinePull)
	// rgl.GET("/messages", p.lineMustSignIn, p.postLinePull)
	// rgl.GET("/download", p.lineMustSignIn, p.postLineDownload)
}
