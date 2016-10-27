package site

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/kapmahc/lotus/engines/base"
)

//NavLinks nav-links
func NavLinks(locale, code string) []base.Link {
	var links []base.Link
	if err := Get(
		fmt.Sprintf("%s://nav-links/%s", locale, code),
		&links); err != nil {
		beego.Error(err)
	}
	return links
}

func init() {
	beego.AddFuncMap("NL", NavLinks)
}
