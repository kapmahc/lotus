package mail

import (
	"github.com/kapmahc/lotus/engines/auth"
	"github.com/kapmahc/lotus/engines/base"
)

func init() {
	auth.Register(func(user *auth.User) *base.Dropdown {
		if user.Has(auth.AdminRole) {
			return &base.Dropdown{
				ID:    "ops-mail",
				Label: "ops-mail-pages.profile",
				Links: []base.Link{
					{
						Href:  "mail.Controller.IndexDomain",
						Label: "ops-mail-pages.domains",
					},
					{
						Href:  "mail.Controller.IndexUser",
						Label: "ops-mail-pages.users",
					},
					{
						Href:  "mail.Controller.IndexAlias",
						Label: "ops-mail-pages.aliases",
					},
					{
						Href:  "mail.Controller.ReadMe",
						Label: "ops-mail-pages.read-me",
					},
				},
			}
		}
		return nil
	})
}
