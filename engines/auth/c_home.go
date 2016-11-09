package auth

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (p *Engine) getHome(wrt http.ResponseWriter, req *http.Request) {
	//TODO
}

func (p *Engine) getLocales(wrt http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	p.Render.JSON(wrt, http.StatusOK, p.I18n.Locales(vars["lang"]))
}
