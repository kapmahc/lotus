package web

import (
	"net/http"

	"github.com/gorilla/csrf"
)

//Csrf write csrf header
func Csrf(wrt http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
	wrt.Header().Set("X-CSRF-Token", csrf.Token(req))
	next(wrt, req)
}
