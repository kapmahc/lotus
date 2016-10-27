package auth

import "github.com/kapmahc/lotus/engines/base"

//Dashboard dashboard nav-links
type Dashboard func(*User) (string, []base.Link)

var dashboards []Dashboard

//Register register dashboard
func Register(args ...Dashboard) {
	dashboards = append(dashboards, args...)
}

func init() {
	Register(func(user *User) (string, []base.Link) {
		return "auth-pages.profile", []base.Link{
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
		}
	})
}
