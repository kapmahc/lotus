package base

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

//SetLocale save locale
func SetLocale(locale, code, message string) {
	var l Locale
	o := orm.NewOrm()
	err := o.QueryTable(&l).
		Filter("lang", locale).
		Filter("code", code).One(&l)
	l.Message = message
	if err == nil {
		_, err = o.Update(&l, "message", "updated_at")
	} else if err == orm.ErrNoRows {
		l.Code = code
		l.Lang = locale
		_, err = o.Insert(&l)
	}
	if err != nil {
		beego.Error(err)
	}
}
