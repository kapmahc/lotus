package site

import "github.com/kapmahc/lotus/engines/base"

//GetAdminBase base
// @router /admin/base [get]
func (p *Controller) GetAdminBase() {
	p.Dashboard()
	p.MustAdmin()
	title := p.T("site-pages.admin-base")
	p.Data["title"] = title
	p.Data["form"] = p.NewForm(
		"fm-admin-base",
		title,
		base.MethodPost,
		p.URLFor("site.Controller.PostAdminBase"),
		[]base.Field{
			&base.TextField{
				ID:    "title",
				Label: p.T("site-attributes.title"),
				Value: p.T("site.title"),
			},
			&base.TextField{
				ID:    "subTitle",
				Label: p.T("site-attributes.subTitle"),
				Value: p.T("site.subTitle"),
			},
			&base.TextField{
				ID:    "keywords",
				Label: p.T("site-attributes.keywords"),
				Value: p.T("site.keywords"),
			},
			&base.Textarea{
				ID:    "description",
				Label: p.T("site-attributes.description"),
				Value: p.T("site.description"),
			},
			&base.TextField{
				ID:    "copyright",
				Label: p.T("site-attributes.copyright"),
				Value: p.T("site.copyright"),
			},
		},
	)
	p.TplName = "auth/form.html"
}

//PostAdminBase base
// @router /admin/base [post]
func (p *Controller) PostAdminBase() {
	p.MustAdmin()
	var fm fmBase
	fl, er := p.ParseForm(&fm)
	if er == nil {
		base.SetLocale(p.Locale, "site.title", fm.Title)
		base.SetLocale(p.Locale, "site.subTitle", fm.SubTitle)
		base.SetLocale(p.Locale, "site.keywords", fm.Keywords)
		base.SetLocale(p.Locale, "site.description", fm.Description)
		base.SetLocale(p.Locale, "site.copyright", fm.Copyright)

		user := p.CurrentUser()
		user.Log(p.T("site-logs.update-base"))

		fl.Notice(p.T("site-pages.success"))
	} else {
		fl.Error(er.Error())
	}
	p.Redirect(fl, "auth.Controller.GetInfo")

}

//GetAdminAuthor author
// @router /admin/author [get]
func (p *Controller) GetAdminAuthor() {
	p.MustAdmin()

}

//GetAdminSeo seo
// @router /admin/seo [get]
func (p *Controller) GetAdminSeo() {

}

//GetAdminStatus status
// @router /admin/status [get]
func (p *Controller) GetAdminStatus() {

}

//GetAdminUsers users
// @router /admin/users [get]
func (p *Controller) GetAdminUsers() {

}

//GetAdminNavBar nav-bar
// @router /admin/nav-bar [get]
func (p *Controller) GetAdminNavBar() {

}
