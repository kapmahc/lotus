package reading

import (
	"github.com/kapmahc/lotus/engines/auth"
	"github.com/kapmahc/lotus/engines/base"
)

func init() {
	auth.Register(func(user *auth.User) *base.Dropdown {

		nb := base.Dropdown{
			ID:    "reading",
			Label: "reading-pages.profile",
			Links: []base.Link{
				{
					Href:  "reading.Controller.GetMyNotes",
					Label: "reading-pages.my-notes",
				},
			},
		}
		if user.Has(auth.AdminRole) {
			nb.Links = append(
				nb.Links,
				base.Link{
					Href:  "reading.Controller.ScanBooks",
					Label: "reading-pages.scan-books",
				},
				base.Link{
					Href:  "reading.Controller.ListBooks",
					Label: "reading-pages.books",
				},
			)
		}
		return &nb
	})
}
