package vpn

import (
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/kapmahc/lotus/engines/base"
)

//IndexUser index users
// @router /users [get]
func (p *Controller) IndexUser() {
	var users []User
	_, err := orm.NewOrm().
		QueryTable(new(User)).
		All(&users, "id", "email", "enable", "start_up", "shut_down")
	p.Check(err)
	p.Data["users"] = users
	p.Data["title"] = p.T("ops-vpn-pages.users")
	p.TplName = "ops/vpn/users/index.html"
}

//NewUser new user
// @router /users/new [get]
func (p *Controller) NewUser() {
	title := p.T("ops-vpn-pages.new-user")
	p.Data["title"] = title
	p.Data["form"] = p.NewForm(
		"fm-new-user",
		title,
		base.MethodPost,
		p.URLFor("vpn.Controller.CreateUser"),
		[]base.Field{
			&base.EmailField{
				ID:    "email",
				Label: p.T("attributes.email"),
			},
			&base.PasswordField{
				ID:     "password",
				Label:  p.T("attributes.password"),
				Helper: p.T("auth-pages.password-must-in-size"),
			},
			&base.PasswordField{
				ID:     "passwordConfirmation",
				Label:  p.T("attributes.passwordConfirmation"),
				Helper: p.T("auth-pages.passwords-must-match"),
			},
			&base.Radio{
				ID:    "enable",
				Label: p.T("attributes.enable"),
				Options: []base.Option{
					{
						Value:    true,
						Name:     p.T("attributes.enable"),
						Selected: true,
					},
					{
						Value: false,
						Name:  p.T("attributes.disable"),
					},
				},
			},
			&base.Textarea{
				ID:     "details",
				Label:  p.T("attributes.details"),
				Helper: p.T("site-pages.can-markdown"),
			},
			&base.DateField{
				ID:    "start_up",
				Label: p.T("attributes.startUp"),
				Value: time.Now(),
			},
			&base.DateField{
				ID:    "shut_down",
				Label: p.T("attributes.shutDown"),
				Value: time.Now(),
			},
		},
	)
	p.TplName = "auth/form.html"
}

const dateFormat = "2006-01-02"

//CreateUser create user
// @router /users [post]
func (p *Controller) CreateUser() {
	var fm fmAddUser
	fl, er := p.ParseForm(&fm)
	o := orm.NewOrm()
	if er == nil {
		if fm.Password != fm.PasswordConfirmation {
			er = p.Error("auth-logs.passwords-not-match")
		}
	}
	if er == nil {
		count, err := o.QueryTable(new(User)).Filter("email", fm.Email).Count()
		p.Check(err)
		if count > 0 {
			er = p.Error("auth-logs.email-already-exists")
		}
	}
	if er == nil {
		user := User{
			Email:    fm.Email,
			Password: fm.Password,
			Enable:   fm.Enable,
			Details:  fm.Details,
		}
		var err error
		user.StartUp, err = time.ParseInLocation(dateFormat, fm.StartUp, time.Local)
		p.Check(err)
		user.ShutDown, err = time.ParseInLocation(dateFormat, fm.ShutDown, time.Local)
		p.Check(err)
		_, err = o.Insert(&user)
		p.Check(err)
		fl.Notice(p.T("site-pages.success"))
		p.Redirect(fl, "vpn.Controller.IndexUser")
	} else {
		fl.Error(er.Error())
		p.Redirect(fl, "vpn.Controller.NewUser")
	}
}

//EditUserProfile edit user profile
// @router /users/:id/edit-profile [get]
func (p *Controller) EditUserProfile() {
	var user User
	err := orm.NewOrm().
		QueryTable(&user).
		Filter("id", p.Ctx.Input.Param(":id")).
		One(&user)
	p.Check(err)

	title := p.T("ops-vpn-pages.edit-user-profile")
	p.Data["title"] = title
	p.Data["form"] = p.NewForm(
		"fm-edit-user-profile",
		title,
		base.MethodPost,
		p.URLFor("vpn.Controller.UpdateUserProfile", ":id", user.ID),
		[]base.Field{
			&base.EmailField{
				ID:       "email",
				Label:    p.T("attributes.email"),
				Readonly: true,
				Value:    user.Email,
			},
			&base.Radio{
				ID:    "enable",
				Label: p.T("attributes.enable"),
				Options: []base.Option{
					{
						Value:    true,
						Name:     p.T("attributes.yes"),
						Selected: user.Enable,
					},
					{
						Value:    false,
						Name:     p.T("attributes.no"),
						Selected: !user.Enable,
					},
				},
			},
			&base.Textarea{
				ID:     "details",
				Label:  p.T("attributes.details"),
				Value:  user.Details,
				Helper: p.T("site-pages.can-markdown"),
			},
			&base.DateField{
				ID:    "start_up",
				Label: p.T("attributes.startUp"),
				Value: user.StartUp,
			},
			&base.DateField{
				ID:    "shut_down",
				Label: p.T("attributes.shutDown"),
				Value: user.ShutDown,
			},
		},
	)
	p.TplName = "auth/form.html"
}

//UpdateUserProfile update user
// @router /users/:id/profile [post]
func (p *Controller) UpdateUserProfile() {

	var user User
	o := orm.NewOrm()
	err := o.QueryTable(&user).
		Filter("id", p.Ctx.Input.Param(":id")).
		One(&user)
	p.Check(err)

	var fm fmEditUserProfile
	fl, er := p.ParseForm(&fm)

	if er == nil {
		user.Details = fm.Details
		user.Enable = fm.Enable
		var err error
		user.StartUp, err = time.ParseInLocation(dateFormat, fm.StartUp, time.Local)
		p.Check(err)
		user.ShutDown, err = time.ParseInLocation(dateFormat, fm.ShutDown, time.Local)
		p.Check(err)
		_, err = o.Update(&user, "updated_at", "details", "start_up", "shut_down", "enable")
		p.Check(err)
		fl.Notice(p.T("site-pages.success"))
		p.Redirect(fl, "vpn.Controller.IndexUser")
	} else {
		fl.Error(er.Error())
		p.Redirect(fl, "vpn.Controller.EditUserProfile", ":id", user.ID)
	}
}

//EditUserPassword edit user password
// @router /users/:id/password [get]
func (p *Controller) EditUserPassword() {
	var user User
	err := orm.NewOrm().
		QueryTable(&user).
		Filter("id", p.Ctx.Input.Param(":id")).
		One(&user)
	p.Check(err)

	title := p.T("ops-vpn-pages.reset-user-password")
	p.Data["title"] = title
	p.Data["form"] = p.NewForm(
		"fm-edit-user-profile",
		title,
		base.MethodPost,
		p.URLFor("vpn.Controller.UpdateUserPassword", ":id", user.ID),
		[]base.Field{
			&base.EmailField{
				ID:       "email",
				Label:    p.T("attributes.email"),
				Readonly: true,
				Value:    user.Email,
			},
			&base.PasswordField{
				ID:     "password",
				Label:  p.T("attributes.password"),
				Helper: p.T("auth-pages.password-must-in-size"),
			},
			&base.PasswordField{
				ID:     "passwordConfirmation",
				Label:  p.T("attributes.passwordConfirmation"),
				Helper: p.T("auth-pages.passwords-must-match"),
			},
		},
	)
	p.TplName = "auth/form.html"
}

//UpdateUserPassword change user password
// @router /users/:id/password [post]
func (p *Controller) UpdateUserPassword() {

	var user User
	o := orm.NewOrm()
	err := o.QueryTable(&user).
		Filter("id", p.Ctx.Input.Param(":id")).
		One(&user)
	p.Check(err)

	var fm fmResetUserPassword
	fl, er := p.ParseForm(&fm)
	if er == nil {
		if fm.Password != fm.PasswordConfirmation {
			er = p.Error("auth-logs.passwords-not-match")
		}
	}
	if er == nil {
		user.Password = fm.Password
		_, err := o.Update(&user, "updated_at", "password")
		p.Check(err)
		fl.Notice(p.T("site-pages.success"))
		p.Redirect(fl, "vpn.Controller.IndexUser")
	} else {
		fl.Error(er.Error())
		p.Redirect(fl, "vpn.Controller.EditUserPassword", ":id", user.ID)
	}
}

//ShowUser show user
// @router /users/:id [get]
func (p *Controller) ShowUser() {
	id := p.Ctx.Input.Param(":id")

	var user User
	o := orm.NewOrm()
	err := o.QueryTable(&user).
		Filter("id", id).One(&user)
	p.Check(err)
	var logs []Log
	_, err = o.QueryTable(new(Log)).Filter("user_id", id).OrderBy("-id").All(&logs)
	p.Check(err)

	p.Data["title"] = p.T("ops-vpn-pages.show-user", user.ID)
	p.Data["user"] = user
	p.Data["logs"] = logs
	p.TplName = "ops/vpn/users/show.html"
}

//DestroyUser destroy user
// @router /users/:id [delete]
func (p *Controller) DestroyUser() {
	_, err := orm.NewOrm().QueryTable(new(User)).
		Filter("id", p.Ctx.Input.Param(":id")).Delete()
	p.Check(err)

	p.Data["json"] = map[string]string{
		"to": p.URLFor("vpn.Controller.IndexUser"),
	}
	p.ServeJSON()
}
