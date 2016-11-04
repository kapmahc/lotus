package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (p *Engine) getUsersSignIn(c *gin.Context) {
	lang := c.MustGet("locale").(string)
	c.HTML(http.StatusOK, "users/sign-in.html", gin.H{"locale": lang})
}
