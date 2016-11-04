package auth

import "github.com/gin-gonic/gin"

//Handler auth-handler
type Handler struct {
}

//CurrentUser current-user
func (p *Handler) CurrentUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		//TODO
	}
}

//MustSignIn must sign-in
func (p *Handler) MustSignIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		//TODO check user status
	}
}

//MustAdmin must has admin role
func (p *Handler) MustAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		//TODO chekc user role
	}
}
