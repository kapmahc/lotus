package base

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/beego/i18n"
)

//T translate
func T(locale, code string, args ...interface{}) string {
	o := orm.NewOrm()
	var l Locale
	err := o.QueryTable(&l).
		Filter("lang", locale).
		Filter("code", code).One(&l, "message")
	if err == nil {
		return fmt.Sprintf(l.Message, args...)
	}
	if err != orm.ErrNoRows {
		beego.Error(err)
	}
	return i18n.Tr(locale, code, args...)
}

func init() {
	beego.AddFuncMap("T", T)
}
