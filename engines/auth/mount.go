package auth

import "github.com/gorilla/mux"

//Mount web points
func (p *Engine) Mount(rt *mux.Router) {
	rt.HandleFunc("/admin/status", p.getAdminStatus).Methods("GET")
	rt.HandleFunc("/users/sign-in", p.getUsersSignIn).Methods("GET")
	// rt.GET("/locales/:lang", p.getLocales)
	// rt.GET("/layout", p.getLayout)
	// rt.GET("/dashboard", p.getDashboard)
	//
	// ug := rt.Group("/users")
	// ug.POST("/sign-in", web.JSON(p.postUsersSignIn))
	// ug.POST("/sign-up", web.JSON(p.postUsersSignUp))
	// ug.POST("/confirm", web.JSON(p.postUsersConfirm))
	// ug.POST("/unlock", web.JSON(p.postUsersUnlock))
	// ug.POST("/forgot-password", web.JSON(p.postUsersForgotPassword))
}
