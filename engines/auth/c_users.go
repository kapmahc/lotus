package auth

import (
	"bytes"
	"fmt"
	"net/http"
	"text/template"
	"time"

	"github.com/SermoDigital/jose/jws"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type fmSignIn struct {
	Email    string `form:"email" binding:"email"`
	Password string `form:"password" binding:"required"`
}

func (p *Engine) postUsersSignIn(c *gin.Context) (interface{}, error) {
	var fm fmSignIn
	if err := c.Bind(&fm); err != nil {
		return nil, err
	}
	user, err := p.Dao.SignIn(fm.Email, fm.Password)
	if err != nil {
		return nil, err
	}
	if !user.IsAvailable() {
		return nil, fmt.Errorf("user %s isn't available", user.Email)
	}
	p.Dao.Log(user.ID, p.I18n.T(c.MustGet("locale").(string), "log.auth.sign-in"))

	cm := jws.Claims{}
	cm.Set("uid", user.UID)
	cm.Set("name", user.Name)
	cm.Set("roles", p.Dao.Authority(user.ID, "-", 0))
	tkn, err := p.Jwt.Sum(cm, 7)
	return gin.H{"token": string(tkn)}, err
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
	user, err := p.Dao.AddEmailUser(fm.Email, fm.Name, fm.Password)
	if err == nil {
		p.Dao.Log(user.ID, p.I18n.T(c.MustGet("locale").(string), "log.auth.sign-up"))
		err = p.sendMail(c.MustGet("locale").(string), "confirm", fm.Email, user.UID)
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
	err = p.sendMail(c.MustGet("locale").(string), "change-password", fm.Email, user.UID)
	return gin.H{}, err
}

type fmChangePassword struct {
	Token                string `form:"token" binding:"required"`
	Password             string `form:"password" binding:"max=50,min=8"`
	PasswordConfirmation string `form:"passwordConfirmation" binding:"eqfield=Password"`
}

func (p *Engine) postUsersChangePassword(c *gin.Context) (interface{}, error) {
	var fm fmChangePassword
	if err := c.Bind(&fm); err != nil {
		return nil, err
	}
	cm, err := p.Jwt.Validate([]byte(fm.Token))
	if err != nil {
		return nil, err
	}
	user, err := p.Dao.GetUserByUID(cm.Get("uid").(string))
	if err != nil {
		return nil, err
	}
	if !user.IsAvailable() {
		return nil, fmt.Errorf("user %s isn't available", user.Email)
	}
	err = p.Db.Model(&user).Updates(map[string]interface{}{
		"password": p.Encryptor.Sum([]byte(fm.Password)),
	}).Error
	if err == nil {
		p.Dao.Log(user.ID, p.I18n.T(c.MustGet("locale").(string), "log.auth.reset-password"))
	}
	return gin.H{}, err
}

func (p *Engine) getUsersConfirm(c *gin.Context) (string, error) {
	token := c.Query("token")
	cm, err := p.Jwt.Validate([]byte(token))
	if err != nil {
		return "", err
	}
	user, err := p.Dao.GetUserByUID(cm.Get("uid").(string))
	if err == nil {
		if user.IsConfirmed() {
			err = fmt.Errorf("user %s was confirmed", user.Email)
		}
	}
	if err == nil {
		now := time.Now()
		err = p.Db.Model(&user).Updates(map[string]interface{}{
			"confirmed_at": &now,
		}).Error
	}
	if err == nil {
		p.Dao.Log(user.ID, p.I18n.T(c.MustGet("locale").(string), "log.auth.confirm"))
	}
	return fmt.Sprintf("%s/users/sign-in", viper.GetString("home.frontend")), err
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
	err = p.sendMail(c.MustGet("locale").(string), "confirm", fm.Email, user.UID)
	return gin.H{}, err
}

func (p *Engine) getUsersUnlock(c *gin.Context) (string, error) {
	token := c.Query("token")
	cm, err := p.Jwt.Validate([]byte(token))
	if err != nil {
		return "", err
	}
	user, err := p.Dao.GetUserByUID(cm.Get("uid").(string))
	if err == nil {
		if !user.IsLocked() {
			err = fmt.Errorf("user %s wasn't confirmed", user.Email)
		}
	}
	if err == nil {
		err = p.Db.Model(&user).Updates(map[string]interface{}{
			"locked_at": nil,
		}).Error
	}
	if err == nil {
		p.Dao.Log(user.ID, p.I18n.T(c.MustGet("locale").(string), "log.auth.unlock"))
	}
	return fmt.Sprintf("%s/users/sign-in", viper.GetString("home.frontend")), err
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
		c.MustGet("locale").(string),
		"unlock",
		fm.Email,
		user.UID,
	)
	return gin.H{}, err
}

func (p *Engine) deleteSignOut(c *gin.Context) {
	user := c.MustGet("user").(*User)
	p.Dao.Log(user.ID, p.I18n.T(c.MustGet("locale").(string), "log.auth.sign-out"))
	c.JSON(http.StatusOK, gin.H{})
}

func (p *Engine) sendMail(lang, act, email, uid string) error {
	cm := jws.Claims{}
	cm.Set("action", act)
	cm.Set("uid", uid)
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
