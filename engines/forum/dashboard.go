package forum

import (
	"github.com/kapmahc/lotus/engines/auth"
	"github.com/kapmahc/lotus/engines/base"
)

func init() {
	auth.Register(func(user *auth.User) *base.Dropdown {
		nb := base.Dropdown{
			ID:    "forum",
			Label: "forum-pages.profile",
			Links: []base.Link{
				{
					Href:  "forum.Controller.GetMyArticles",
					Label: "forum-pages.my-articles",
				},
				{
					Href:  "forum.Controller.GetMyComments",
					Label: "forum-pages.my-comments",
				},
			},
		}
		if user.Has(auth.AdminRole) {
			nb.Links = append(
				nb.Links,
				base.Link{
					Href:  "forum.Controller.GetAdminTags",
					Label: "forum-pages.tags",
				},
			)
		}
		return &nb
	})
}
