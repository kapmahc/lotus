package auth

import "github.com/gin-gonic/gin"

func (p *Engine) postUsersSignIn(c *gin.Context) {
	// TODO
	// lang := c.MustGet("locale").(string)
	// c.HTML(http.StatusOK, "users/non-sign-in", gin.H{
	// 	"locale": lang,
	// 	"form": gin.H{
	// 		"title": p.I18n.T(lang, "auth.users.sign-in"),
	// 		"fields": []gin.H{
	// 			gin.H{"type": "email", "id": "email"},
	// 			gin.H{"type": "password", "id": "password"},
	// 			gin.H{"type": "password", "id": "passwordConfirm"},
	// 		},
	// 	},
	// })
}

type fmSignUp struct {
	Name                 string `form:"name" binding:"max=32,min=2"`
	Email                string `form:"email" binding:"email"`
	Password             string `form:"password" binding:"max=50,min=8"`
	PasswordConfirmation string `form:"passwordConfirmation" binding:"eqfield=Password"`
}

func (p *Engine) postUsersSignUp(c *gin.Context) (interface{}, error) {
	var fm fmSignUp
	if err := c.Bind(&fm); err != nil {
		return nil, err
	}
	return p.Dao.AddEmailUser(fm.Email, fm.Name, fm.Password)
}
