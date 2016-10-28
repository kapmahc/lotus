package sitemap

import (
	"encoding/xml"
	"time"
)

//Item item
type Item struct {
	XMLName xml.Name  `xml:"url"`
	Link    string    `xml:"loc"`
	Updated time.Time `xml:"lastmod"`
}

//Sitemap sitemap
type Sitemap struct {
	XMLName xml.Name `xml:"urlset"`
	Xmlns   string   `xml:"xmlns,attr"`

	Items []Item
}

//New new sitemap
func New() *Sitemap {
	return &Sitemap{
		Xmlns: "http://www.sitemaps.org/schemas/sitemap/0.9",
		Items: make([]Item, 0),
	}
}
