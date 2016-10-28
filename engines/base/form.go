package base

import "html/template"

const (
	//MethodPost post method
	MethodPost = "post"
	//MethodPatch patch method
	MethodPatch = "patch"

	//TypeMarkdown markdown type
	TypeMarkdown = "markdown"
	//TypeHTML html type
	TypeHTML = "html"
)

//Form form model
type Form struct {
	XSRF   template.HTML
	Locale string
	ID     string
	Title  string
	Action string
	Method string
	Fields []Field
}

//Field input field
type Field interface {
	Type() string
}

//Select select
type Select struct {
	ID       string
	Label    string
	Value    interface{}
	Readonly bool
	Options  map[interface{}]interface{}
	Helper   string
}

//Type type
func (p *Select) Type() string {
	return "select"
}

//TextField text
type TextField struct {
	ID       string
	Label    string
	Value    string
	Readonly bool
	Helper   string
}

//Type type
func (p *TextField) Type() string {
	return "text"
}

//Textarea textarea
type Textarea struct {
	ID       string
	Label    string
	Value    string
	Readonly bool
	Helper   string
}

//Type type
func (p *Textarea) Type() string {
	return "textarea"
}

//EmailField text
type EmailField struct {
	ID       string
	Label    string
	Value    string
	Readonly bool
	Helper   string
}

//Type type
func (p *EmailField) Type() string {
	return "email"
}

//PasswordField text
type PasswordField struct {
	ID     string
	Label  string
	Helper string
}

//Type type
func (p *PasswordField) Type() string {
	return "password"
}

//HiddenField hidden
type HiddenField struct {
	ID    string
	Value interface{}
}

//Type type
func (p *HiddenField) Type() string {
	return "hidden"
}
