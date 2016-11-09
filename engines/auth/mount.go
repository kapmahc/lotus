package auth

import "github.com/gorilla/mux"

//Mount web points
func (p *Engine) Mount(rt *mux.Router) {
	rt.HandleFunc("/", p.getHome).Methods("GET").Name("home")
	rt.HandleFunc("/locales/{lang}", p.getLocales).Methods("GET").Name("locales")

	rt.HandleFunc("/install", p.installNew).Methods("GET").Name("auth.install.new")
	rt.HandleFunc("/install", p.installCreate).Methods("POST").Name("auth.install.create")

	rt.HandleFunc("/users/sign-in", p.sessionsNew).Methods("GET").Name("auth.sessions.new")
	rt.HandleFunc("/users/sign-in", p.sessionsCreate).Methods("POST").Name("auth.sessions.create")
	rt.HandleFunc("/users/sign-out", p.sessionsDestroy).Methods("DELETE").Name("auth.sessions.destroy")

	rt.HandleFunc("/users/forgot-password", p.passwordsNew).Methods("GET").Name("auth.passwords.new")
	rt.HandleFunc("/users/forgot-password", p.passwordsNew).Methods("POST").Name("auth.passwords.create")
	rt.HandleFunc("/users/change-password", p.passwordsEdit).Methods("GET").Name("auth.passwords.edit")
	rt.HandleFunc("/users/change-password", p.passwordsNew).Methods("POST").Name("auth.passwords.update")

	rt.HandleFunc("/users/sign-up", p.registrationsNew).Methods("GET").Name("auth.registrations.new")
	rt.HandleFunc("/users/sign-up", p.registrationsCreate).Methods("POST").Name("auth.registrations.create")

	rt.HandleFunc("/users/edit-profile", p.registrationsEdit).Methods("GET").Name("auth.registrations.edit")
	rt.HandleFunc("/users/edit-profile", p.registrationsUpdate).Methods("POST").Name("auth.registrations.update")

	rt.HandleFunc("/users/unlock", p.unlocksNew).Methods("GET").Name("auth.unlocks.new")
	rt.HandleFunc("/users/unlock", p.unlocksCreate).Methods("POST").Name("auth.unlocks.create")

	rt.HandleFunc("/users/confirm", p.confirmationsNew).Methods("GET").Name("auth.confirmations.new")
	rt.HandleFunc("/users/confirm", p.confirmationsCreate).Methods("POST").Name("auth.confirmations.create")

}
