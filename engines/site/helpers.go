package site

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/kapmahc/lotus/engines/base"
)

//NavLinks nav-links
func NavLinks(locale, code string) []base.Link {
	var links []base.Link
	base.Get(
		fmt.Sprintf("%s://nav-links/%s", locale, code),
		&links)
	return links
}

func init() {
	beego.AddFuncMap("NL", NavLinks)
}
