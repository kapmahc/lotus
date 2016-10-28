package base

import (
	"time"

	"github.com/astaxie/beego"
	"github.com/gorilla/feeds"
)

//FeedsFunc feeds handler
type FeedsFunc func(string, *beego.Controller) []*feeds.Item

var feedHanders []FeedsFunc

//RegisterAtom registe feeds function
func RegisterAtom(args ...FeedsFunc) {
	feedHanders = append(feedHanders, args...)
}

//Atom get atom feeds
func Atom(locale string, crl *beego.Controller) *feeds.Feed {
	now := time.Now()
	home := beego.AppConfig.String("homeurl")
	feed := feeds.Feed{
		Title:       T(locale, "site.title"),
		Link:        &feeds.Link{Href: home},
		Subtitle:    T(locale, "site.subTitle"),
		Description: T(locale, "site.description"),
		Author: &feeds.Author{
			Name:  T(locale, "site.author-name"),
			Email: T(locale, "site.author-email"),
		},
		Items:   make([]*feeds.Item, 0),
		Created: now,
	}
	for _, fn := range feedHanders {
		args := fn(home, crl)
		feed.Items = append(feed.Items, args...)
	}
	return &feed
}
