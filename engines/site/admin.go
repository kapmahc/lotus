package site

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/astaxie/beego/orm"
	"github.com/kapmahc/lotus/engines/auth"
	"github.com/kapmahc/lotus/engines/base"
)

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

	p.Data["google"] = google
	p.Data["baidu"] = baidu
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
	p.TplName = "site/admin/seo.html"
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
	p.Dashboard()
	p.MustAdmin()
	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)

	status := map[string]string{
		"OS":      fmt.Sprintf("%s %s", runtime.GOOS, runtime.GOARCH),
		"go-lang": fmt.Sprintf("%s %s", runtime.Version(), runtime.GOROOT()),
		"host": fmt.Sprintf(
			"cpus(%d) memory(%.1fG)",
			runtime.NumCPU(),
			float64(mem.HeapSys)/1024.0/1024.0,
		),
	}

	p.Data["status"] = status
	p.Data["title"] = p.T("site-pages.admin-status")
	p.TplName = "site/admin/status.html"
}

//GetAdminUsers users
// @router /admin/users [get]
func (p *Controller) GetAdminUsers() {
	p.Dashboard()
	p.MustAdmin()

	var users []auth.User
	_, err := orm.NewOrm().
		QueryTable(new(auth.User)).
		OrderBy("-last_sign_in_at").
		All(&users, "id", "name", "email", "last_sign_in_at")
	p.Check(err)

	p.Data["users"] = users
	p.Data["title"] = p.T("site-pages.admin-users")
	p.TplName = "site/admin/users.html"
}

//LinkSeperator link seperator
const LinkSeperator = " = "

//GetAdminNavBar nav-bar
// @router /admin/nav-bar [get]
func (p *Controller) GetAdminNavBar() {
	p.Dashboard()
	p.MustAdmin()
	title := p.T("site-pages.admin-nav-bar")
	p.Data["title"] = title

	link2str := func(links []base.Link) string {
		var s string
		for _, l := range links {
			s += fmt.Sprintf("%s%s%s\n", l.Href, LinkSeperator, l.Label)
		}
		return s
	}

	p.Data["form"] = p.NewForm(
		"fm-admin-nav-bar",
		title,
		base.MethodPost,
		p.URLFor("site.Controller.PostAdminNavBar"),
		[]base.Field{
			&base.Textarea{
				ID:     "header",
				Label:  p.T("site-attributes.header-links"),
				Value:  link2str(NavLinks(p.Locale, "header")),
				Helper: p.T("site-pages.one-link-per-line"),
			},
			&base.Textarea{
				ID:     "footer",
				Label:  p.T("site-attributes.footer-links"),
				Value:  link2str(NavLinks(p.Locale, "footer")),
				Helper: p.T("site-pages.one-link-per-line"),
			},
		},
	)
	p.TplName = "auth/form.html"
}

//PostAdminNavBar nav-bar
// @router /admin/nav-bar [post]
func (p *Controller) PostAdminNavBar() {
	p.MustAdmin()
	var fm fmNavBar
	fl, er := p.ParseForm(&fm)

	save := func(code, lines string) error {
		var links []base.Link
		for _, line := range strings.Split(lines, "\n") {
			ss := strings.Split(line, LinkSeperator)
			if len(ss) != 2 {
				return p.Error("site-logs.bad-format")
			}
			links = append(links, base.Link{Href: ss[0], Label: ss[1]})
		}
		Set(fmt.Sprintf("%s://nav-links/%s", p.Locale, code), links, false)
		return nil
	}
	if er == nil {
		er = save("header", fm.Header)
	}
	if er == nil {
		er = save("footer", fm.Footer)
	}

	if er == nil {
		user := p.CurrentUser()
		user.Log(p.T("site-logs.update-nav-bar"))

		fl.Notice(p.T("site-pages.success"))
	} else {
		fl.Error(er.Error())
	}
	p.Redirect(fl, "site.Controller.GetAdminNavBar")
}
