package mail

import (
	"github.com/astaxie/beego/orm"
	"github.com/kapmahc/lotus/engines/base"
)

//IndexUser index users
// @router /users [get]
func (p *Controller) IndexUser() {
	var users []User
	_, err := orm.NewOrm().
		QueryTable(new(User)).
		OrderBy("-updated_at").
		RelatedSel().
		All(&users)
	p.Check(err)
	p.Data["users"] = users
	p.Data["title"] = p.T("ops-mail-pages.users")
	p.TplName = "ops/mail/users/index.html"
}

func (p *Controller) domainOptions() *base.Select {

	var options []base.Option
	var domains []Domain
	_, err := orm.NewOrm().QueryTable(new(Domain)).OrderBy("name").All(&domains, "id", "name")
	p.Check(err)
	for _, d := range domains {
		options = append(options, base.Option{
			Name:  d.Name,
			Value: d.ID,
		})
	}

	return &base.Select{
		ID:      "domain_id",
		Label:   p.T("ops-mail-attributes.domain"),
		Options: options,
	}
}

//NewUser new user
// @router /users/new [get]
func (p *Controller) NewUser() {
	title := p.T("ops-mail-pages.new-user")
	p.Data["title"] = title

	p.Data["form"] = p.NewForm(
		"fm-new-user",
		title,
		base.MethodPost,
		p.URLFor("mail.Controller.CreateUser"),
		[]base.Field{
			p.domainOptions(),
			&base.TextField{
				ID:    "name",
				Label: p.T("attributes.username"),
			},
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
		},
	)
	p.TplName = "auth/form.html"
}

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
		var domain Domain
		err := o.QueryTable(&domain).Filter("id", fm.DomainID).One(&domain, "id")
		p.Check(err)
		user := User{
			Name:   fm.Name,
			Domain: &domain,
			Email:  fm.Email,
		}
		p.Check(user.SetPassword(fm.Password, 8))
		_, err = o.Insert(&user)
		p.Check(err)
		fl.Notice(p.T("site-pages.success"))
		p.Redirect(fl, "mail.Controller.IndexUser")
	} else {
		fl.Error(er.Error())
		p.Redirect(fl, "mail.Controller.NewUser")
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

	title := p.T("ops-mail-pages.edit-user-profile")
	p.Data["title"] = title
	p.Data["form"] = p.NewForm(
		"fm-edit-user-profile",
		title,
		base.MethodPost,
		p.URLFor("mail.Controller.UpdateUserProfile", ":id", user.ID),
		[]base.Field{
			&base.EmailField{
				ID:       "email",
				Label:    p.T("attributes.email"),
				Readonly: true,
				Value:    user.Email,
			},
			&base.TextField{
				ID:    "name",
				Label: p.T("attributes.name"),
				Value: user.Name,
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
		user.Name = fm.Name
		_, err = o.Update(&user, "updated_at", "name")
		p.Check(err)
		fl.Notice(p.T("site-pages.success"))
		p.Redirect(fl, "mail.Controller.IndexUser")
	} else {
		fl.Error(er.Error())
		p.Redirect(fl, "mail.Controller.EditUserProfile", ":id", user.ID)
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

	title := p.T("ops-mail-pages.reset-user-password")
	p.Data["title"] = title
	p.Data["form"] = p.NewForm(
		"fm-edit-user-profile",
		title,
		base.MethodPost,
		p.URLFor("mail.Controller.UpdateUserPassword", ":id", user.ID),
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
		p.Redirect(fl, "mail.Controller.IndexUser")
	} else {
		fl.Error(er.Error())
		p.Redirect(fl, "mail.Controller.EditUserPassword", ":id", user.ID)
	}
}

//DestroyUser destroy user
// @router /users/:id [delete]
func (p *Controller) DestroyUser() {
	_, err := orm.NewOrm().QueryTable(new(User)).
		Filter("id", p.Ctx.Input.Param(":id")).Delete()
	p.Check(err)

	p.Data["json"] = map[string]string{
		"to": p.URLFor("mail.Controller.IndexUser"),
	}
	p.ServeJSON()
}
