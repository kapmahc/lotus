package crawler

import "github.com/kapmahc/lotus/web"

//LineMessage line message
type LineMessage struct {
	web.Model
	UserID    string
	MessageID string
	Content   string
}

//
// func (p *Engine) lineCurrentUser(c *gin.Context) {
//
// }
//
// func (p *Engine) lineMustSignIn(c *gin.Context) {
//
// }
//
// func (p *Engine) postLineDownload(c *gin.Context) {
//
// }
//
// func (p *Engine) postLineSignIn(c *gin.Context) {
//
// }
//
// func (p *Engine) getLineSignOut(c *gin.Context) {
//
// }
//
// func (p *Engine) postLinePull(c *gin.Context) {
//
// }
//
// func pullLineMessage(userID, password, channelID string) ([]LineMessage, error) {
// 	return nil, nil
// }
