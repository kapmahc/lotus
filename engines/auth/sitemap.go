package auth

import (
	"github.com/astaxie/beego"
	"github.com/kapmahc/lotus/engines/base"
	"github.com/kapmahc/lotus/sitemap"
)

func init() {
	base.RegisterSitemap(func(home string, crl *beego.Controller) []sitemap.Item {
		return []sitemap.Item{
			{Link: home + crl.URLFor("auth.Controller.GetSignIn")},
			{Link: home + crl.URLFor("auth.Controller.GetSignUp")},
			{Link: home + crl.URLFor("auth.Controller.GetForgotPassword")},
			{Link: home + crl.URLFor("auth.Controller.GetConfirm")},
		}
	})
}
