package auth

import (
	"bytes"
	"fmt"
	"text/template"
	"time"

	"github.com/SermoDigital/jose/jws"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

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
	_, err := p.Dao.AddEmailUser(fm.Email, fm.Name, fm.Password)
	if err == nil {
		err = p.sendMail(c.MustGet("locale").(string), "confirm", fm.Email)
	}
	return gin.H{}, err
}

type fmEmail struct {
	Email string `form:"email" binding:"email"`
}

func (p *Engine) postUsersForgotPassword(c *gin.Context) (interface{}, error) {
	var fm fmEmail
	if err := c.Bind(&fm); err != nil {
		return nil, err
	}
	user, err := p.Dao.GetUserByEmail(fm.Email)
	if err != nil {
		return nil, err
	}
	if !user.IsAvailable() {
		return nil, fmt.Errorf("email account %s isn't available", fm.Email)
	}
	err = p.sendMail(c.MustGet("locale").(string), "change-password", fm.Email)
	return gin.H{}, err
}

func (p *Engine) getUsersConfirm(c *gin.Context) (string, error) {
	token := c.Query("token")
	cm, err := p.Jwt.Validate([]byte(token))
	if err != nil {
		return "", err
	}
	email := cm.Get("email").(string)
	user, err := p.Dao.GetUserByEmail(email)
	if err == nil {
		if user.IsConfirmed() {
			err = fmt.Errorf("user %s was confirmed", email)
		}
	}
	if err == nil {
		now := time.Now()
		err = p.Db.Model(&user).Updates(map[string]interface{}{
			"confirmed_at": &now,
		}).Error
	}
	return viper.GetString("home.frontend"), err
}

func (p *Engine) postUsersConfirm(c *gin.Context) (interface{}, error) {
	var fm fmEmail
	if err := c.Bind(&fm); err != nil {
		return nil, err
	}
	user, err := p.Dao.GetUserByEmail(fm.Email)
	if err != nil {
		return nil, err
	}
	if user.IsConfirmed() {
		return nil, fmt.Errorf("email account %s was confirmed", fm.Email)
	}
	err = p.sendMail(c.MustGet("locale").(string), "confirm", fm.Email)
	return gin.H{}, err
}

func (p *Engine) getUsersUnlock(c *gin.Context) (string, error) {
	token := c.Query("token")
	cm, err := p.Jwt.Validate([]byte(token))
	if err != nil {
		return "", err
	}
	email := cm.Get("email").(string)
	user, err := p.Dao.GetUserByEmail(email)
	if err == nil {
		if !user.IsLocked() {
			err = fmt.Errorf("user %s wasn't confirmed", email)
		}
	}
	if err == nil {
		err = p.Db.Model(&user).Updates(map[string]interface{}{
			"locked_at": nil,
		}).Error
	}
	return viper.GetString("home.frontend"), err
}

func (p *Engine) postUsersUnlock(c *gin.Context) (interface{}, error) {
	var fm fmEmail
	if err := c.Bind(&fm); err != nil {
		return nil, err
	}
	user, err := p.Dao.GetUserByEmail(fm.Email)
	if err != nil {
		return nil, err
	}
	if !user.IsLocked() {
		return nil, fmt.Errorf("email account %s is not locked", fm.Email)
	}
	err = p.sendMail(
		c.MustGet("locale").(string), "unlock", fm.Email)
	return gin.H{}, err
}

func (p *Engine) sendMail(lang, act, email string) error {
	cm := jws.Claims{}
	cm.Set("action", act)
	cm.Set("email", email)
	tkn, err := p.Jwt.Sum(cm, 1)
	if err != nil {
		return err
	}

	st, err := template.New("").Parse(p.I18n.T(lang, fmt.Sprintf("email.auth.%s.subject", act)))
	if err != nil {
		return err
	}
	bt, err := template.New("").Parse(p.I18n.T(lang, fmt.Sprintf("email.auth.%s.body", act)))
	if err != nil {
		return err
	}
	model := struct {
		Frontend string
		Backend  string
		Token    string
	}{
		Frontend: viper.GetString("home.frontend"),
		Backend:  viper.GetString("home.backend"),
		Token:    string(tkn),
	}

	var subject bytes.Buffer
	var body bytes.Buffer

	if err = st.Execute(&subject, model); err != nil {
		return err
	}
	if err = bt.Execute(&body, model); err != nil {
		return err
	}

	return p.MailSender.Send(email, subject.String(), body.String())
}
