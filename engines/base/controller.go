package base

import (
	"errors"
	"fmt"
	"html/template"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"github.com/beego/i18n"
)

//Controller base controller
type Controller struct {
	beego.Controller
	Locale string
}

//Redirect redirect with flash
func (p *Controller) Redirect(flash *beego.FlashData, action string, args ...interface{}) {
	if flash != nil {
		flash.Store(&p.Controller)
	}
	p.Controller.Redirect(p.URLFor(action, args...), 302)
}

//T translate
func (p *Controller) T(code string, args ...interface{}) string {
	return T(p.Locale, code, args)
}

//Error create error
func (p *Controller) Error(code string, args ...interface{}) error {
	return errors.New(T(p.Locale, code, args))
}

//Check check error
func (p *Controller) Check(err error) {
	if err != nil {
		beego.Error(err)
		p.Abort("500")
	}
}

//ParseForm parse form
func (p *Controller) ParseForm(form interface{}) (*beego.FlashData, error) {
	flash := beego.NewFlash()
	var valid validation.Validation
	err := p.Controller.ParseForm(form)
	if err == nil {
		if b, e := valid.Valid(form); e == nil {
			if !b {
				msg := "<ul>"
				for k, v := range valid.ErrorsMap {
					msg += fmt.Sprintf("<li>%s: %s</li>", k, v)
				}
				msg += "</ul>"
				err = errors.New(msg)
			}
		} else {
			err = e
		}
	}
	return flash, err
}

//NewForm new form model
func (p *Controller) NewForm(id, title, method, action string, fields []Field) *Form {
	return &Form{
		XSRF:   template.HTML(p.XSRFFormHTML()),
		ID:     id,
		Locale: p.Locale,
		Title:  title,
		Method: method,
		Action: action,
		Fields: fields,
	}
}

//SetLocale set locale
func (p *Controller) SetLocale() {
	const key = "locale"
	write := false

	// 1. Check URL arguments.
	lang := p.Input().Get(key)

	// 2. Get language information from cookies.
	if len(lang) == 0 {
		lang = p.Ctx.GetCookie(key)
	} else {
		write = true
	}

	// 3. Get language information from 'Accept-Language'.
	if len(lang) == 0 {
		al := p.Ctx.Request.Header.Get("Accept-Language")
		if len(al) > 4 {
			al = al[:5] // Only compare first 5 letters.
		}
		write = true
	}

	// 4. Default language is English.
	if !i18n.IsExist(lang) {
		lang = "en-US"
		write = true
	}

	// Save language information in cookies.
	if write {
		p.Ctx.SetCookie(key, lang, 1<<31-1, "/")
	}

	// Set language properties.
	p.Locale = lang
	p.Data[key] = lang
	p.Data["languages"] = i18n.ListLangs()
}
