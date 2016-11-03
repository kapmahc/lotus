package vpn

import (
	"github.com/kapmahc/lotus/engines/auth"
	"github.com/kapmahc/lotus/engines/base"
)

func init() {
	auth.Register(func(user *auth.User) *base.Dropdown {
		if user.Has(auth.AdminRole) {
			return &base.Dropdown{
				ID:    "ops-vpn",
				Label: "ops-vpn-pages.profile",
				Links: []base.Link{
					{
						Href:  "vpn.Controller.IndexUser",
						Label: "ops-vpn-pages.users",
					},
					{
						Href:  "vpn.Controller.IndexLog",
						Label: "ops-vpn-pages.logs",
					},
					{
						Href:  "vpn.Controller.ReadMe",
						Label: "ops-vpn-pages.read-me",
					},
				},
			}
		}
		return nil
	})
}
