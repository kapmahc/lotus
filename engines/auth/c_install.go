package auth

import (
	"net/http"

	"github.com/kapmahc/lotus/web"
)

func (p *Engine) installNew(wrt http.ResponseWriter, req *http.Request) {
	if p._countUsers() > 0 {
		web.Redirect(wrt, req, web.URLFor(p.Router, "home", nil))
		return
	}
	ctx := req.Context()
	lang := ctx.Value(web.LOCALE).(string)
	title := p.I18n.T(lang, "auth.pages.install")

	p.Render.HTML(wrt, http.StatusOK, "users/non-sign-in", map[string]interface{}{
		"locale": lang,
		"title":  title,
		"form": web.NewForm(
			req,
			"install",
			title,
			web.URLFor(p.Router, "auth.install.create", nil),
			&web.TextField{
				ID:    "title",
				Label: p.I18n.T(lang, "attributes.title"),
			},
			&web.TextField{
				ID:    "subTitle",
				Label: p.I18n.T(lang, "auth.attributes.subTitle"),
			},
			&web.TextField{
				ID:    "username",
				Label: p.I18n.T(lang, "attributes.username"),
			},
			&web.EmailField{
				ID:    "email",
				Label: p.I18n.T(lang, "attributes.email"),
			},
			&web.PasswordField{
				ID:    "password",
				Label: p.I18n.T(lang, "attributes.password"),
			},
			&web.PasswordField{
				ID:    "passwordConfirmation",
				Label: p.I18n.T(lang, "attributes.passwordConfirmation"),
			},
		),
	})
}
func (p *Engine) installCreate(wrt http.ResponseWriter, req *http.Request) {
	//TODO
}
