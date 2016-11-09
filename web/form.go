package web

import (
	"html/template"
	"net/http"
	"net/url"
	"time"

	"github.com/gorilla/csrf"
)

const (

	//TypeMarkdown markdown type
	TypeMarkdown = "markdown"
	//TypeHTML html type
	TypeHTML = "html"
)

//NewForm new form
func NewForm(req *http.Request, id, title string, action *url.URL, fields ...Field) *Form {
	return &Form{
		CSRF:   csrf.TemplateField(req),
		Locale: req.Context().Value(LOCALE).(string),
		ID:     id,
		Title:  title,
		Method: "POST",
		Action: action.String(),
		Fields: fields,
	}
}

//Form form model
type Form struct {
	CSRF   template.HTML
	Locale string
	ID     string
	Title  string
	Action string
	Method string
	Fields []Field
}

//Add add fields
func (p *Form) Add(fields ...Field) {
	p.Fields = append(p.Fields, fields...)
}

//Field input field
type Field interface {
	Type() string
}

//Option option
type Option struct {
	Value    interface{}
	Name     interface{}
	Selected bool
}

//Select select
type Select struct {
	ID       string
	Label    string
	Multi    bool
	Readonly bool
	Options  []Option
	Helper   string
}

//Type type
func (p *Select) Type() string {
	return "select"
}

//Radio radio
type Radio struct {
	ID       string
	Label    string
	Multi    bool
	Readonly bool
	Options  []Option
	Helper   string
}

//Type type
func (p *Radio) Type() string {
	return "radio"
}

//CheckBox checkbox
type CheckBox struct {
	ID       string
	Label    string
	Readonly bool
	Options  []Option
	Helper   string
}

//Type type
func (p *CheckBox) Type() string {
	return "checkbox"
}

//DateField date
type DateField struct {
	ID       string
	Label    string
	Value    time.Time
	Readonly bool
	Helper   string
}

//Type type
func (p *DateField) Type() string {
	return "date"
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
