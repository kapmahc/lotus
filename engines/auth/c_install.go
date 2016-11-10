package auth

import (
	"net/http"
	"time"

	sessions "github.com/goincremental/negroni-sessions"
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
	session := sessions.GetSession(req)
	lang := req.Context().Value(web.LOCALE).(string)
	var fm fmInstall
	err := web.ParseForm(req, p.Validate, &fm)
	var user *User
	if err == nil {
		p.I18n.Set(lang, "site.title", fm.Title)
		p.I18n.Set(lang, "site.subTitle", fm.SubTitle)
		user, err = p.Dao.AddEmailUser(lang, fm.Email, fm.Username, fm.Password)

	}
	if err == nil {
		err = p.Db.Model(user).Update("confirmed_at", time.Now()).Error
	}
	if err == nil {
		session.AddFlash(web.Notice(p.I18n.T(lang, "messages.success")))
	} else {
		session.AddFlash(web.Error(err.Error()))
	}
	session.Save(wrt, req)
}
