package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (p *Engine) getLocales(c *gin.Context) {
	c.JSON(http.StatusOK, p.I18n.Locales(c.Param("lang")))
}
