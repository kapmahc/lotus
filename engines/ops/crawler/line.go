package crawler

import (
	"github.com/gin-gonic/gin"
	"github.com/kapmahc/lotus/web"
)

//LineMessage line message
type LineMessage struct {
	web.Model
	UserID    string
	MessageID string
	Content   string
}

func (p *Engine) lineCurrentUser(c *gin.Context) {

}

func (p *Engine) lineMustSignIn(c *gin.Context) {

}

func (p *Engine) postLineDownload(c *gin.Context) {

}

func (p *Engine) postLineSignIn(c *gin.Context) {

}

func (p *Engine) getLineSignOut(c *gin.Context) {

}

func (p *Engine) postLinePull(c *gin.Context) {

}
