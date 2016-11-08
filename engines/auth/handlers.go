package auth

import "github.com/kapmahc/lotus/web"

//Handler auth-handler
type Handler struct {
	I18n *web.I18n `inject:""`
	Dao  *Dao      `inject:""`
}

// //CurrentUser current-user
// func (p *Handler) CurrentUser(c *gin.Context) {
// 	//TODO
// }
//
// //MustSignIn must sign-in
// func (p *Handler) MustSignIn(c *gin.Context) {
// 	//TODO
// }
//
// //MustAdmin must has admin role
// func (p *Handler) MustAdmin(c *gin.Context) {
// 	//TODO
// }
