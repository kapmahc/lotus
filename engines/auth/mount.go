package auth

import "github.com/gorilla/mux"

//Mount web points
func (p *Engine) Mount(rt *mux.Router) {
	rt.HandleFunc("/", p.getHome).Methods("GET").Name("home")
	rt.HandleFunc("/users/sign-in", p.getUsersSignIn).Methods("GET").Name("users.sign-in")
	rt.HandleFunc("/locales/{lang}", p.getLocales).Methods("GET").Name("locales")
}
