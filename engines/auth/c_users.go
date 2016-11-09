package auth

import (
	"net/http"

	"github.com/kapmahc/lotus/web"
)

const (
	actConfirm       = "confirm"
	actResetPassword = "reset-password"
	actUnlock        = "unlock"
)

func (p *Engine) passwordsNew(wrt http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	lang := ctx.Value(web.LOCALE).(string)
	title := p.I18n.T(lang, "auth.pages.passwords.new")

	p.Render.HTML(wrt, http.StatusOK, "users/non-sign-in", web.H{
		"locale": lang,
		"title":  title,
		"form": web.NewForm(
			req,
			"forgot-password",
			title,
			web.URLFor(p.Router, "auth.passwords.create", nil),
			&web.EmailField{
				ID:    "email",
				Label: p.I18n.T(lang, "attributes.email"),
			},
		),
	})
}
func (p *Engine) passwordsCreate(wrt http.ResponseWriter, req *http.Request) {
	//TODO
}

func (p *Engine) passwordsEdit(wrt http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	lang := ctx.Value(web.LOCALE).(string)
	title := p.I18n.T(lang, "auth.pages.passwords.edit")

	p.Render.HTML(wrt, http.StatusOK, "users/non-sign-in", web.H{
		"locale": lang,
		"title":  title,
		"form": web.NewForm(
			req,
			"reset-password",
			title,
			web.URLFor(p.Router, "auth.passwords.update", nil),
			&web.HiddenField{
				ID:    "token",
				Value: req.URL.Query().Get("token"),
			},
			&web.PasswordField{
				ID:    "password",
				Label: p.I18n.T(lang, "attributes.password"),
			},
			&web.PasswordField{
				ID:    "passwordConfirmation",
				Label: p.I18n.T(lang, "attributes.passwordConfirmation"),
			},
		),
	})
}
func (p *Engine) passwordsUpdate(wrt http.ResponseWriter, req *http.Request) {
	//TODO
}

func (p *Engine) registrationsNew(wrt http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	lang := ctx.Value(web.LOCALE).(string)
	title := p.I18n.T(lang, "auth.pages.registrations.new")

	p.Render.HTML(wrt, http.StatusOK, "users/non-sign-in", web.H{
		"locale": lang,
		"title":  title,
		"form": web.NewForm(
			req,
			"sign-up",
			title,
			web.URLFor(p.Router, "auth.registrations.create", nil),
			&web.TextField{
				ID:    "name",
				Label: p.I18n.T(lang, "attributes.username"),
			},
			&web.EmailField{
				ID:    "email",
				Label: p.I18n.T(lang, "attributes.email"),
			},
			&web.PasswordField{
				ID:    "password",
				Label: p.I18n.T(lang, "attributes.password"),
			},
			&web.PasswordField{
				ID:    "passwordConfirmation",
				Label: p.I18n.T(lang, "attributes.passwordConfirmation"),
			},
		),
	})
}
func (p *Engine) registrationsCreate(wrt http.ResponseWriter, req *http.Request) {
	//TODO
}

func (p *Engine) registrationsEdit(wrt http.ResponseWriter, req *http.Request) {
	//TODO
}
func (p *Engine) registrationsUpdate(wrt http.ResponseWriter, req *http.Request) {
	//TODO
}

func (p *Engine) unlocksNew(wrt http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	lang := ctx.Value(web.LOCALE).(string)
	title := p.I18n.T(lang, "auth.pages.unlocks.new")

	p.Render.HTML(wrt, http.StatusOK, "users/non-sign-in", web.H{
		"locale": lang,
		"title":  title,
		"form": web.NewForm(
			req,
			"unlock",
			title,
			web.URLFor(p.Router, "auth.unlocks.create", nil),
			&web.EmailField{
				ID:    "email",
				Label: p.I18n.T(lang, "attributes.email"),
			},
		),
	})
}

func (p *Engine) unlocksCreate(wrt http.ResponseWriter, req *http.Request) {
	//TODO
}

func (p *Engine) confirmationsNew(wrt http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	lang := ctx.Value(web.LOCALE).(string)
	title := p.I18n.T(lang, "auth.pages.confirmations.new")

	p.Render.HTML(wrt, http.StatusOK, "users/non-sign-in", web.H{
		"locale": lang,
		"title":  title,
		"form": web.NewForm(
			req,
			"confirm",
			title,
			web.URLFor(p.Router, "auth.confirmations.create", nil),
			&web.EmailField{
				ID:    "email",
				Label: p.I18n.T(lang, "attributes.email"),
			},
		),
	})
}

func (p *Engine) confirmationsCreate(wrt http.ResponseWriter, req *http.Request) {
	//TODO
}

func (p *Engine) sessionsNew(wrt http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	lang := ctx.Value(web.LOCALE).(string)
	title := p.I18n.T(lang, "auth.pages.sessions.new")

	p.Render.HTML(wrt, http.StatusOK, "users/non-sign-in", web.H{
		"locale": lang,
		"title":  title,
		"form": web.NewForm(
			req,
			"sign-in",
			title,
			web.URLFor(p.Router, "auth.sessions.create", nil),
			&web.EmailField{
				ID:    "email",
				Label: p.I18n.T(lang, "attributes.email"),
			},
			&web.PasswordField{
				ID:    "password",
				Label: p.I18n.T(lang, "attributes.password"),
			},
		),
	})
}
func (p *Engine) sessionsDestroy(wrt http.ResponseWriter, req *http.Request) {
	//TODO
}
func (p *Engine) sessionsCreate(wrt http.ResponseWriter, req *http.Request) {
	//TODO
}

//
// func (p *Engine) postUsersSignIn(c *gin.Context) (interface{}, error) {
// 	lang := c.MustGet("locale").(string)
// 	var fm fmSignIn
// 	if err := c.Bind(&fm); err != nil {
// 		return nil, err
// 	}
// 	user, err := p.Dao.GetUserByEmail(fm.Email)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if !p.Hmac.Chk([]byte(fm.Password), user.Password) {
// 		return nil, errors.New(p.I18n.T(lang, "auth.messages.email-password-not-match"))
// 	}
// 	if !user.IsConfirmed() {
// 		return nil, errors.New(p.I18n.T(lang, "auth.messages.need-confirm"))
// 	}
// 	if user.IsLocked() {
// 		return nil, errors.New(p.I18n.T(lang, "auth.messages.need-unlock"))
// 	}
// 	p.Dao.Log(user.ID, p.I18n.T(lang, "auth.logs.sign-in"))
// 	return user, nil
// }
//
// func (p *Engine) postUsersSignUp(c *gin.Context) (interface{}, error) {
// 	lang := c.MustGet("locale").(string)
// 	var fm fmSignUp
// 	if err := c.Bind(&fm); err != nil {
// 		return nil, err
// 	}
// 	user, err := p.Dao.AddEmailUser(lang, fm.Email, fm.Name, fm.Password)
// 	if err != nil {
// 		return nil, err
// 	}
// 	p.Dao.Log(user.ID, p.I18n.T(lang, "auth.logs.sign-up"))
// 	p.sendMail(user, actConfirm)
// 	return user, nil
// }
//
// func (p *Engine) postUsersConfirm(c *gin.Context) (interface{}, error) {
// 	lang := c.MustGet("locale").(string)
// 	var fm fmEmail
// 	if err := c.Bind(&fm); err != nil {
// 		return nil, err
// 	}
// 	user, err := p.Dao.GetUserByEmail(fm.Email)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if user.IsConfirmed() {
// 		return nil, errors.New(p.I18n.T(lang, "auth.messages.user-already-confirmed"))
// 	}
// 	p.sendMail(user, actConfirm)
// 	return user, nil
// }
//
// func (p *Engine) postUsersUnlock(c *gin.Context) (interface{}, error) {
// 	lang := c.MustGet("locale").(string)
// 	var fm fmEmail
// 	if err := c.Bind(&fm); err != nil {
// 		return nil, err
// 	}
// 	user, err := p.Dao.GetUserByEmail(fm.Email)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if !user.IsLocked() {
// 		return nil, errors.New(p.I18n.T(lang, "auth.messages.user-not-locked"))
// 	}
// 	p.sendMail(user, actUnlock)
// 	return user, nil
// }
//
// func (p *Engine) postUsersForgotPassword(c *gin.Context) (interface{}, error) {
// 	var fm fmEmail
// 	if err := c.Bind(&fm); err != nil {
// 		return nil, err
// 	}
// 	user, err := p.Dao.GetUserByEmail(fm.Email)
// 	if err != nil {
// 		return nil, err
// 	}
// 	p.sendMail(user, actResetPassword)
// 	return user, nil
// }
//
// func (p *Engine) sendMail(user *User, act string) {
// 	var subject string
// 	var body string
// 	//TODO
// 	task := signatures.TaskSignature{
// 		Name: userEmailQueue,
// 		Args: []signatures.TaskArg{
// 			signatures.TaskArg{
// 				Type:  "string",
// 				Value: user.Email,
// 			},
// 			signatures.TaskArg{
// 				Type:  "string",
// 				Value: subject,
// 			},
// 			signatures.TaskArg{
// 				Type:  "string",
// 				Value: body,
// 			},
// 		},
// 	}
// 	if _, err := p.Server.SendTask(&task); err != nil {
// 		p.Logger.Error("add mail job: %v", err)
// 	}
// }
