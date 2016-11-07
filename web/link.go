package web

//Link link
type Link struct {
	Href  string
	Label string
}

//Dropdown dropdown
type Dropdown struct {
	ID    string
	Label string
	Href  string
	Links []Link
}
