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
	p.Redirect(fl, "site.Controller.GetAdminBase")
}

//GetAdminAuthor author
// @router /admin/author [get]
func (p *Controller) GetAdminAuthor() {
	p.Dashboard()
	p.MustAdmin()
	title := p.T("site-pages.admin-author")
	p.Data["title"] = title
	p.Data["form"] = p.NewForm(
		"fm-admin-author",
		title,
		base.MethodPost,
		p.URLFor("site.Controller.PostAdminAuthor"),
		[]base.Field{
			&base.TextField{
				ID:    "name",
				Label: p.T("site-attributes.author-name"),
				Value: p.T("site.author-name"),
			},
			&base.EmailField{
				ID:    "email",
				Label: p.T("site-attributes.author-email"),
				Value: p.T("site.author-email"),
			},
		},
	)
	p.TplName = "auth/form.html"

}

//PostAdminAuthor author
// @router /admin/author [post]
func (p *Controller) PostAdminAuthor() {
	p.MustAdmin()
	var fm fmAuthor
	fl, er := p.ParseForm(&fm)
	if er == nil {
		base.SetLocale(p.Locale, "site.author-name", fm.Name)
		base.SetLocale(p.Locale, "site.author-email", fm.Email)

		user := p.CurrentUser()
		user.Log(p.T("site-logs.update-author"))

		fl.Notice(p.T("site-pages.success"))
	} else {
		fl.Error(er.Error())
	}
	p.Redirect(fl, "site.Controller.GetAdminAuthor")
}

//GetAdminSeo seo
// @router /admin/seo [get]
func (p *Controller) GetAdminSeo() {
	p.Dashboard()
	p.MustAdmin()
	title := p.T("site-pages.admin-seo")
	p.Data["title"] = title
	var google string
	var baidu string
	Get("google.verify.code", &google)
	Get("baidu.verify.code", &baidu)
	p.Data["form"] = p.NewForm(
		"fm-admin-seo",
		title,
		base.MethodPost,
		p.URLFor("site.Controller.PostAdminSeo"),
		[]base.Field{
			&base.TextField{
				ID:    "google",
				Label: p.T("site-attributes.google-verify-code"),
				Value: google,
			},
			&base.TextField{
				ID:    "baidu",
				Label: p.T("site-attributes.baidu-verify-code"),
				Value: baidu,
			},
		},
	)
	p.TplName = "auth/form.html"
}

//PostAdminSeo seo
// @router /admin/seo [post]
func (p *Controller) PostAdminSeo() {
	p.MustAdmin()
	var fm fmSeo
	fl, er := p.ParseForm(&fm)
	if er == nil {
		Set("google.verify.code", fm.Google, false)
		Set("baidu.verify.code", fm.Baidu, false)

		user := p.CurrentUser()
		user.Log(p.T("site-logs.update-seo"))

		fl.Notice(p.T("site-pages.success"))
	} else {
		fl.Error(er.Error())
	}
	p.Redirect(fl, "site.Controller.GetAdminSeo")
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
