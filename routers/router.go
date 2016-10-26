package routers

import (
	"os"
	"path/filepath"

	"github.com/astaxie/beego"
	"github.com/beego/i18n"
	"github.com/kapmahc/lotus/engines/auth"
	"github.com/kapmahc/lotus/engines/forum"
	"github.com/kapmahc/lotus/engines/ops/mail"
	"github.com/kapmahc/lotus/engines/ops/vpn"
	"github.com/kapmahc/lotus/engines/reading"
	"github.com/kapmahc/lotus/engines/shop"
	"github.com/kapmahc/lotus/engines/site"
)

func init() {
	// load locales
	if err := filepath.Walk(filepath.Join("conf", "locales"), func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		const ext = ".ini"
		name := info.Name()
		if info.Mode().IsRegular() && filepath.Ext(name) == ext {
			lang := name[0 : len(name)-len(ext)]
			beego.Info("Loading language: ", lang)
			if err := i18n.SetMessage(lang, path); err != nil {
				beego.Error(err)
				return err
			}
		}
		return nil
	}); err != nil {
		beego.Error(err)
	}

	// register controllers
	beego.Include(&site.Controller{})
	beego.AddNamespace(
		beego.NewNamespace("/users", beego.NSInclude(&auth.Controller{})),

		beego.NewNamespace("/forum", beego.NSInclude(&forum.Controller{})),
		beego.NewNamespace("/reading", beego.NSInclude(&reading.Controller{})),
		beego.NewNamespace("/shop", beego.NSInclude(&shop.Controller{})),
		beego.NewNamespace(
			"/ops",
			beego.NSNamespace("mail", beego.NSInclude(&mail.Controller{})),
			beego.NSNamespace("vpn", beego.NSInclude(&vpn.Controller{})),
		),
	)
}
