package site

import (
	"github.com/kapmahc/lotus/engines/auth"
	"github.com/kapmahc/lotus/engines/base"
)

func init() {
	auth.Register(func(user *auth.User) *base.Dropdown {
		if user.Has(auth.AdminRole) {

			return &base.Dropdown{
				ID:    "site",
				Label: "site-pages.profile",
				Links: []base.Link{
					{
						Href:  "site.Controller.GetAdminBase",
						Label: "site-pages.admin-base",
					},
					{
						Href:  "site.Controller.GetAdminAuthor",
						Label: "site-pages.admin-author",
					},
					{
						Href:  "site.Controller.GetAdminSeo",
						Label: "site-pages.admin-seo",
					},
					{
						Href:  "site.Controller.GetAdminStatus",
						Label: "site-pages.admin-status",
					},
					{
						Href:  "site.Controller.GetAdminUsers",
						Label: "site-pages.admin-users",
					},
					{
						Href:  "site.Controller.GetAdminNavBar",
						Label: "site-pages.admin-navBar",
					},
				},
			}
		}
		return nil
	})
}
