package reading

import (
	"encoding/xml"
	"io"
)

// HTML html
type HTML struct {
	Title string `xml:"head>title"`
	Body  Body   `xml:"body"`
}

// Body html body
type Body struct {
	Content string `xml:",innerxml"`
}

func (p *Controller) parseHTML(rdr io.Reader) (string, string) {
	h := HTML{}
	err := xml.NewDecoder(rdr).Decode(&h)
	p.Check(err)
	return h.Title, h.Body.Content
}

// func (p *Controller) parseHTML(rdr io.Reader) (string, string) {
// 	doc, err := html.Parse(rdr)
// 	p.Check(err)
// 	var title bytes.Buffer
// 	var body bytes.Buffer
//
// 	var fn func(*html.Node)
// 	fn = func(n *html.Node) {
// 		if n.Type == html.ElementNode && n.Data == "title" {
// 			err := html.Render(&title, n)
// 			p.Check(err)
// 		}
// 		if n.Type == html.ElementNode && n.Data == "body" {
// 			err := html.Render(&body, n)
// 			p.Check(err)
// 		}
// 		for c := n.FirstChild; c != nil; c = c.NextSibling {
// 			fn(c)
// 		}
// 	}
// 	fn(doc)
// 	return title.String(), body.String()
// }
