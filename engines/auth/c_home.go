package auth

import (
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
	"github.com/kapmahc/lotus/web"
)

func (p *Engine) getHome(wrt http.ResponseWriter, req *http.Request) {
	if p._countUsers() == 0 {
		web.Redirect(wrt, req, web.URLFor(p.Router, "auth.install.new", nil))
		return
	}
	ctx := req.Context()
	lang := ctx.Value(web.LOCALE).(string)
	p.Render.HTML(wrt, http.StatusOK, "home", web.H{
		"locale": lang,
		"title":  p.I18n.T(lang, "auth.pages.home"),
	}) //TODO
}

func (p *Engine) getLocales(wrt http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	p.Render.JSON(wrt, http.StatusOK, p.I18n.Locales(vars["lang"]))
}

func (p *Engine) _countUsers() int {
	var count int
	if err := p.Db.Model(&User{}).Count(&count).Error; err != nil {
		log.Error(err)
	}
	return count
}
