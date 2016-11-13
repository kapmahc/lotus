package auth

import (
	"bytes"
	"fmt"
	"html/template"
	"time"

	"github.com/RichardKnop/machinery/v1/signatures"
	"github.com/SermoDigital/jose/jws"
	"github.com/gin-gonic/gin"
	"github.com/kapmahc/lotus/web"
	"github.com/spf13/viper"
)

const (
	actConfirm       = "confirm"
	actResetPassword = "reset-password"
	actUnlock        = "unlock"
)

func (p *Engine) postUsersSignIn(c *gin.Context) (interface{}, error) {
	lang := c.MustGet(web.LOCALE).(string)
	var fm fmSignIn
	if err := c.Bind(&fm); err != nil {
		return nil, err
	}
	user, err := p.Dao.GetUserByEmail(fm.Email)
	if err != nil {
		return nil, err
	}
	if !p.Hmac.Chk([]byte(fm.Password), user.Password) {
		return nil, p.I18n.E(lang, "auth.messages.email-password-not-match")
	}
	if !user.IsConfirmed() {
		return nil, p.I18n.E(lang, "auth.messages.need-confirm")
	}
	if user.IsLocked() {
		return nil, p.I18n.E(lang, "auth.messages.need-unlock")
	}
	now := time.Now()
	p.Db.Model(user).Updates(User{
		SignInCount:  user.SignInCount + 1,
		LastSignInAt: &now,
	})
	p.Dao.Logf(user.ID, lang, "auth.logs.sign-in")

	tkn, err := p.Jwt.Sum(p.Dao.UserClaims(user), 7)
	if err != nil {
		return nil, err
	}
	return gin.H{
		"user":  user,
		"token": string(tkn),
	}, err
}

func (p *Engine) postUsersSignUp(c *gin.Context) (interface{}, error) {
	lang := c.MustGet(web.LOCALE).(string)
	var fm fmSignUp
	if err := c.Bind(&fm); err != nil {
		return nil, err
	}
	user, err := p.Dao.AddEmailUser(lang, fm.Email, fm.Name, fm.Password)
	if err != nil {
		return nil, err
	}
	p.Dao.Logf(user.ID, lang, "auth.logs.sign-up")
	err = p._sendMail(lang, user, actConfirm)
	return gin.H{
		"user":    user,
		"message": p.I18n.T(lang, "auth.pages.confirm-success"),
	}, err
}

func (p *Engine) getUsersConfirm(c *gin.Context) (string, error) {
	lang := c.MustGet(web.LOCALE).(string)
	cm, err := p.Jwt.Validate([]byte(c.Query("token")))
	if err == nil {
		if cm.Get("act").(string) != actConfirm {
			err = p.I18n.E(lang, "messages.bad-token")
		}
	}
	var user *User
	if err == nil {
		user, err = p.Dao.GetUserByEmail(cm.Get("email").(string))
	}
	if err == nil {
		if user.IsConfirmed() {
			err = p.I18n.E(lang, "auth.logs.user-already-comfirmed")
		}
	}
	if err == nil {
		err = p.Db.Model(user).Update("confirmed_at", time.Now()).Error
	}
	if err == nil {
		p.Dao.Logf(user.ID, lang, "auth.logs.confirm")
	}
	return p._signInURL(), err
}

func (p *Engine) _signInURL() string {
	return fmt.Sprintf("%s/users/sign-in", viper.GetString("server.frontend"))
}

func (p *Engine) postUsersConfirm(c *gin.Context) (interface{}, error) {
	lang := c.MustGet(web.LOCALE).(string)
	var fm fmEmail
	if err := c.Bind(&fm); err != nil {
		return nil, err
	}
	user, err := p.Dao.GetUserByEmail(fm.Email)
	if err != nil {
		return nil, err
	}
	if user.IsConfirmed() {
		return nil, p.I18n.E(lang, "auth.messages.user-already-confirmed")
	}
	err = p._sendMail(lang, user, actConfirm)
	return gin.H{
		"user":    user,
		"message": p.I18n.T(lang, "auth.pages.confirm-success"),
	}, err
}

func (p *Engine) getUsersUnlock(c *gin.Context) (string, error) {
	lang := c.MustGet(web.LOCALE).(string)
	cm, err := p.Jwt.Validate([]byte(c.Query("token")))
	if err == nil {
		if cm.Get("act").(string) != actUnlock {
			err = p.I18n.E(lang, "messages.bad-token")
		}
	}
	var user *User
	if err == nil {
		user, err = p.Dao.GetUserByEmail(cm.Get("email").(string))
	}
	if err == nil {
		if !user.IsLocked() {
			err = p.I18n.E(lang, "auth.logs.user-not-locked")
		}
	}
	if err == nil {
		err = p.Db.Model(user).Update("locked_at", nil).Error
	}
	if err == nil {
		p.Dao.Logf(user.ID, lang, "auth.logs.unlock")
	}
	return p._signInURL(), err
}

func (p *Engine) postUsersUnlock(c *gin.Context) (interface{}, error) {
	lang := c.MustGet(web.LOCALE).(string)
	var fm fmEmail
	if err := c.Bind(&fm); err != nil {
		return nil, err
	}
	user, err := p.Dao.GetUserByEmail(fm.Email)
	if err != nil {
		return nil, err
	}
	if !user.IsLocked() {
		return nil, p.I18n.E(lang, "auth.messages.user-not-locked")
	}
	err = p._sendMail(lang, user, actUnlock)
	return gin.H{
		"user":    user,
		"message": p.I18n.T(lang, "auth.pages.unlock-success"),
	}, err
}

func (p *Engine) postUsersForgotPassword(c *gin.Context) (interface{}, error) {
	lang := c.MustGet(web.LOCALE).(string)
	var fm fmEmail
	if err := c.Bind(&fm); err != nil {
		return nil, err
	}
	user, err := p.Dao.GetUserByEmail(fm.Email)
	if err != nil {
		return nil, err
	}
	err = p._sendMail(lang, user, actResetPassword)
	return gin.H{
		"user":    user,
		"message": p.I18n.T(lang, "auth.pages.forgot-password-success"),
	}, err
}

func (p *Engine) postUsersResetPassword(c *gin.Context) (interface{}, error) {
	lang := c.MustGet(web.LOCALE).(string)
	var fm fmResetPassword
	if err := c.Bind(&fm); err != nil {
		return nil, err
	}
	cm, err := p.Jwt.Validate([]byte(fm.Token))
	if err == nil {
		if cm.Get("act").(string) != actResetPassword {
			err = p.I18n.E(lang, "messages.bad-token")
		}
	}
	var user *User
	if err == nil {
		user, err = p.Dao.GetUserByEmail(cm.Get("email").(string))
	}

	if err == nil {
		err = p.Db.Model(user).Update("password", p.Hmac.Sum([]byte(fm.Password))).Error
	}
	if err == nil {
		p.Dao.Logf(user.ID, lang, "auth.logs.reset-password")
	}

	return gin.H{
		"user":    user,
		"message": p.I18n.T(lang, "auth.pages.reset-password-success"),
	}, err
}

func (p *Engine) _sendMail(lang string, user *User, act string) error {
	cm := jws.Claims{}
	cm.Set("act", act)
	cm.Set("email", user.Email)
	tkn, err := p.Jwt.Sum(cm, 1)
	if err != nil {
		return err
	}

	st, err := template.New("").
		Parse(p.I18n.T(lang, fmt.Sprintf("auth.emails.%s-subject", act)))
	if err != nil {
		return err
	}
	bt, err := template.New("").
		Parse(p.I18n.T(lang, fmt.Sprintf("auth.emails.%s-body", act)))
	if err != nil {
		return err
	}

	model := struct {
		Frontend string
		Backend  string
		Action   string
		Token    string
	}{
		Frontend: viper.GetString("server.frontend"),
		Backend:  viper.GetString("server.backend"),
		Action:   act,
		Token:    string(tkn),
	}

	var subject bytes.Buffer
	var body bytes.Buffer

	err = st.Execute(&subject, model)
	if err != nil {
		return err
	}
	bt.Execute(&body, model)
	if err != nil {
		return err
	}

	task := signatures.TaskSignature{
		Name: userEmailQueue,
		Args: []signatures.TaskArg{
			signatures.TaskArg{
				Type:  "string",
				Value: user.Email,
			},
			signatures.TaskArg{
				Type:  "string",
				Value: subject.String(),
			},
			signatures.TaskArg{
				Type:  "string",
				Value: body.String(),
			},
		},
	}
	_, err = p.Server.SendTask(&task)
	return err
}
