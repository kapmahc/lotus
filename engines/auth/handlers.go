package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/justinas/nosurf"
	"github.com/kapmahc/lotus/web"
)

//Handler auth-handler
type Handler struct {
	I18n *web.I18n `inject:""`
	Dao  *Dao      `inject:""`
}

//Form create a form model
func (p *Handler) Form(c *gin.Context) {
	var fm web.Form
	fm.Locale = c.MustGet("locale").(string)
	fm.XSRF = nosurf.Token(c.Request)
	fm.Method = web.MethodPost
	c.Set("form", &fm)
}

//CurrentUser current-user
func (p *Handler) CurrentUser(c *gin.Context) {
	//TODO
}

//MustSignIn must sign-in
func (p *Handler) MustSignIn(c *gin.Context) {
	//TODO
}

//MustAdmin must has admin role
func (p *Handler) MustAdmin(c *gin.Context) {
	//TODO
}

//Page page model
type Page struct {
	Locale      string
	Languages   []string
	Title       string
	SubTitle    string
	Keywords    string
	Description string
	Copyright   string
	TopLinks    []web.Link
	BootomLinks []web.Link
	Author      struct {
		Name  string
		Email string
	}
}

//Layout layout
func (p *Handler) Layout(c *gin.Context) {
	lang := c.MustGet("locale").(string)
	var page Page
	page.Locale = lang
	page.Languages, _ = p.I18n.Languages()
	page.Title = p.I18n.T(lang, "site.title")
	page.SubTitle = p.I18n.T(lang, "site.subTitle")
	page.Keywords = p.I18n.T(lang, "site.keywords")
	page.Description = p.I18n.T(lang, "site.description")
	page.Copyright = p.I18n.T(lang, "site.copyright")
	p.Dao.Get("site.top-links", &page.TopLinks)
	p.Dao.Get("site.bootom-links", &page.BootomLinks)
	p.Dao.Get("site.author", &page.Author)

	c.Set("data", gin.H{"page": page, "locale": lang})
}

//Dashboard dashboard
func (p *Handler) Dashboard(c *gin.Context) {
	//TODO
}
