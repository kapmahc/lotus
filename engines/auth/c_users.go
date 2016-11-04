package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kapmahc/lotus/web"
)

func (p *Engine) getUsersSignIn(c *gin.Context) {
	lang := c.MustGet("data").(string)
	fm := c.MustGet("form").(*web.Form)
	fm.ID = "sign-in"
	fm.Title = p.I18n.T(lang, "auth.pages.sign-in")
	fm.Action = "/users/sign-in"
	fm.Add(
		&web.EmailField{
			ID:    "email",
			Label: p.I18n.T(lang, "attributes.email"),
		},
		&web.PasswordField{
			ID:    "password",
			Label: p.I18n.T(lang, "attributes.password"),
		},
	)
	c.HTML(
		http.StatusOK,
		"users/non-sign-in.html",
		gin.H{
			"title":  fm.Title,
			"locale": lang,
			"form":   fm,
		},
	)
}
