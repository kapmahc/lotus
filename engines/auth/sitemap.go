package auth

import (
	"time"

	"github.com/astaxie/beego"
	"github.com/kapmahc/lotus/engines/base"
	"github.com/kapmahc/lotus/sitemap"
)

func init() {
	base.RegisterSitemap(func(home string, crl *beego.Controller) []sitemap.Item {
		var installed time.Time
		base.Get("site.install", &installed)

		return []sitemap.Item{
			{Link: home + crl.URLFor("auth.Controller.GetSignIn"), Updated: installed},
			{Link: home + crl.URLFor("auth.Controller.GetSignUp"), Updated: installed},
			{Link: home + crl.URLFor("auth.Controller.GetForgotPassword"), Updated: installed},
			{Link: home + crl.URLFor("auth.Controller.GetConfirm"), Updated: installed},
		}
	})
}
