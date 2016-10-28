package auth

import "github.com/kapmahc/lotus/engines/base"

//NavBarFunc nav-bar
type NavBarFunc func(*User) *base.Dropdown

var dashboard []NavBarFunc

//Register register dashboard
func Register(args ...NavBarFunc) {
	dashboard = append(dashboard, args...)
}

func init() {
	Register(func(user *User) *base.Dropdown {
		return &base.Dropdown{
			ID:    "auth",
			Label: "auth-pages.profile",
			Links: []base.Link{
				{
					Href:  "auth.Controller.GetInfo",
					Label: "auth-pages.info",
				},
				{
					Href:  "auth.Controller.GetChangePassword",
					Label: "auth-pages.change-password",
				},
				{
					Href:  "auth.Controller.GetLogs",
					Label: "auth-pages.logs",
				},
			},
		}
	})
}
