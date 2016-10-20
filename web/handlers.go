package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//Redirect rediect response
func Redirect(fn func(*gin.Context) (string, error)) gin.HandlerFunc {
	return func(c *gin.Context) {
		if url, err := fn(c); err == nil {
			c.Redirect(http.StatusTemporaryRedirect, url)
		} else {
			if !c.IsAborted() {
				c.AbortWithStatus(http.StatusInternalServerError)
			}
			c.Writer.WriteString(err.Error())
		}
	}
}

//JSON json response
func JSON(fn func(*gin.Context) (interface{}, error)) gin.HandlerFunc {
	return func(c *gin.Context) {
		if val, err := fn(c); err == nil {
			c.JSON(http.StatusOK, val)
		} else {
			if !c.IsAborted() {
				c.AbortWithStatus(http.StatusInternalServerError)
			}
			c.Writer.WriteString(err.Error())
		}
	}
}
