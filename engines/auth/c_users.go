package auth

import (
	"errors"

	"github.com/RichardKnop/machinery/v1/signatures"
	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
)

const (
	actConfirm       = "confirm"
	actResetPassword = "reset-password"
	actUnlock        = "unlock"
)

func (p *Engine) postUsersSignIn(c *gin.Context) (interface{}, error) {
	lang := c.MustGet("locale").(string)
	var fm fmSignIn
	if err := c.Bind(&fm); err != nil {
		return nil, err
	}
	user, err := p.Dao.GetUserByEmail(fm.Email)
	if err != nil {
		return nil, err
	}
	if !p.Hmac.Chk([]byte(fm.Password), user.Password) {
		return nil, errors.New(p.I18n.T(lang, "auth.messages.email-password-not-match"))
	}
	if !user.IsConfirmed() {
		return nil, errors.New(p.I18n.T(lang, "auth.messages.need-confirm"))
	}
	if user.IsLocked() {
		return nil, errors.New(p.I18n.T(lang, "auth.messages.need-unlock"))
	}
	p.Dao.Log(user.ID, p.I18n.T(lang, "auth.logs.sign-in"))
	return user, nil
}

func (p *Engine) postUsersSignUp(c *gin.Context) (interface{}, error) {
	lang := c.MustGet("locale").(string)
	var fm fmSignUp
	if err := c.Bind(&fm); err != nil {
		return nil, err
	}
	user, err := p.Dao.AddEmailUser(lang, fm.Email, fm.Name, fm.Password)
	if err != nil {
		return nil, err
	}
	p.Dao.Log(user.ID, p.I18n.T(lang, "auth.logs.sign-up"))
	p.sendMail(user, actConfirm)
	return user, nil
}

func (p *Engine) postUsersConfirm(c *gin.Context) (interface{}, error) {
	lang := c.MustGet("locale").(string)
	var fm fmEmail
	if err := c.Bind(&fm); err != nil {
		return nil, err
	}
	user, err := p.Dao.GetUserByEmail(fm.Email)
	if err != nil {
		return nil, err
	}
	if user.IsConfirmed() {
		return nil, errors.New(p.I18n.T(lang, "auth.messages.user-already-confirmed"))
	}
	p.sendMail(user, actConfirm)
	return user, nil
}

func (p *Engine) postUsersUnlock(c *gin.Context) (interface{}, error) {
	lang := c.MustGet("locale").(string)
	var fm fmEmail
	if err := c.Bind(&fm); err != nil {
		return nil, err
	}
	user, err := p.Dao.GetUserByEmail(fm.Email)
	if err != nil {
		return nil, err
	}
	if !user.IsLocked() {
		return nil, errors.New(p.I18n.T(lang, "auth.messages.user-not-locked"))
	}
	p.sendMail(user, actUnlock)
	return user, nil
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
	p.sendMail(user, actResetPassword)
	return user, nil
}

func (p *Engine) sendMail(user *User, act string) {
	var subject string
	var body string
	//TODO
	task := signatures.TaskSignature{
		Name: userEmailQueue,
		Args: []signatures.TaskArg{
			signatures.TaskArg{
				Type:  "string",
				Value: user.Email,
			},
			signatures.TaskArg{
				Type:  "string",
				Value: subject,
			},
			signatures.TaskArg{
				Type:  "string",
				Value: body,
			},
		},
	}
	if _, err := p.Server.SendTask(&task); err != nil {
		log.Error(err)
	}
}
