package auth

import (
	"net/http"

	"github.com/kapmahc/lotus/web"
)

//Handler auth-handler
type Handler struct {
	I18n *web.I18n `inject:""`
	Dao  *Dao      `inject:""`
}

//CurrentUser current-user
func (p *Handler) CurrentUser(wrt http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
	//TODO
}

//MustSignIn must sign-in
func (p *Handler) MustSignIn(wrt http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
	//TODO
}

//MustAdmin must has admin role
func (p *Handler) MustAdmin(wrt http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
	//TODO
}
