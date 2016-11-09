package auth

import "github.com/gorilla/mux"

//Mount web points
func (p *Engine) Mount(rt *mux.Router) {
	rt.HandleFunc("/", p.getHome).Methods("GET").Name("home")
	rt.HandleFunc("/locales/{lang}", p.getLocales).Methods("GET").Name("locales")

	rt.HandleFunc("/users/sign-in", p.sessionsNew).Methods("GET").Name("auth.sessions.new")
	rt.HandleFunc("/users/sign-in", p.sessionsCreate).Methods("POST").Name("auth.sessions.create")
	rt.HandleFunc("/users/sign-out", p.sessionsDestroy).Methods("DELETE").Name("auth.sessions.destroy")

	rt.HandleFunc("/users/forgot-password", p.passwordsNew).Methods("GET").Name("auth.passwords.new")
	rt.HandleFunc("/users/change-password", p.passwordsEdit).Methods("GET").Name("auth.passwords.edit")

	rt.HandleFunc("/users/sign-up", p.registrationsNew).Methods("GET").Name("auth.registrations.new")
	rt.HandleFunc("/users/edit-profile", p.registrationsEdit).Methods("GET").Name("auth.registrations.edit")

	rt.HandleFunc("/users/unlock", p.unlocksNew).Methods("GET").Name("auth.unlocks.new")
	rt.HandleFunc("/users/confirm", p.confirmationsNew).Methods("GET").Name("auth.confirmations.new")

}
