package base

import (
	"github.com/astaxie/beego"
	"github.com/beego/i18n"
)

//Controller base controller
type Controller struct {
	beego.Controller
	Locale string
}

//SetLocale set locale
func (p *Controller) SetLocale() {
	const key = "locale"

	// 1. Check URL arguments.
	lang := p.Input().Get(key)

	// 2. Get language information from cookies.
	if len(lang) == 0 {
		lang = p.Ctx.GetCookie(key)
	}

	// 3. Get language information from 'Accept-Language'.
	if len(lang) == 0 {
		al := p.Ctx.Request.Header.Get("Accept-Language")
		if len(al) > 4 {
			al = al[:5] // Only compare first 5 letters.
		}
	}

	// 4: Check lang
	if !i18n.IsExist(lang) {
		lang = ""
	}

	// 5. Default language is English.
	if len(lang) == 0 {
		lang = "en-US"
	}

	// Save language information in cookies.
	p.Ctx.SetCookie(key, lang, 1<<31-1, "/")

	// Set language properties.
	p.Locale = lang
	p.Data[key] = lang
	p.Data["languages"] = i18n.ListLangs()
}
