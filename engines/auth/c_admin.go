package auth

import (
	"fmt"
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/kapmahc/lotus/web"
)

func (p *Engine) getAdminAuthor(c *gin.Context) (interface{}, error) {
	author := gin.H{}
	for _, k := range []string{"email", "name"} {
		var v string
		p.Dao.Get(fmt.Sprintf("site.author.%s", k), &v)
		author[k] = v
	}
	return author, nil
}

func (p *Engine) postAdminAuthor(c *gin.Context) (interface{}, error) {
	var fm fmSiteAuthor
	if err := c.Bind(&fm); err != nil {
		return nil, err
	}
	p.Dao.Set("site.author.name", fm.Name, false)
	p.Dao.Set("site.author.email", fm.Email, false)

	return gin.H{}, nil
}

func (p *Engine) getAdminBase(c *gin.Context) (interface{}, error) {
	lang := c.MustGet(web.LOCALE).(string)
	ret := gin.H{}
	for _, k := range []string{"title", "subTitle", "keywords", "description", "copyright"} {
		ret[k] = p.I18n.T(lang, fmt.Sprintf("site.%s", k))
	}
	return ret, nil
}

func (p *Engine) postAdminBase(c *gin.Context) (interface{}, error) {
	lang := c.MustGet(web.LOCALE).(string)

	var fm fmSiteBase
	if err := c.Bind(&fm); err != nil {
		return nil, err
	}

	p.I18n.Set(lang, "site.title", fm.Title)
	p.I18n.Set(lang, "site.subTitle", fm.SubTitle)
	p.I18n.Set(lang, "site.keywords", fm.Keywords)
	p.I18n.Set(lang, "site.description", fm.Description)
	p.I18n.Set(lang, "site.copyright", fm.Copyright)

	return gin.H{}, nil
}

func (p *Engine) getAdminI18n(c *gin.Context) (interface{}, error) {
	lang := c.MustGet(web.LOCALE).(string)
	var items []web.Locale
	err := p.Db.Select([]string{"code", "message"}).
		Where("lang = ?", lang).
		Order("code ASC").Find(&items).Error
	return items, err
}

func (p *Engine) postAdminI18n(c *gin.Context) (interface{}, error) {
	lang := c.MustGet(web.LOCALE).(string)

	var fm fmLocale
	if err := c.Bind(&fm); err != nil {
		return nil, err
	}

	p.I18n.Set(lang, fm.Code, fm.Message)
	return gin.H{}, nil
}

func (p *Engine) getAdminSeo(c *gin.Context) (interface{}, error) {
	ret := gin.H{}
	for _, k := range []string{"google", "baidu"} {
		var code string
		p.Dao.Get(fmt.Sprintf("%s.verify.code", k), &code)
		ret[k] = code
	}
	return ret, nil
}

func (p *Engine) postAdminSeo(c *gin.Context) (interface{}, error) {
	var fm fmSeo
	if err := c.Bind(&fm); err != nil {
		return nil, err
	}

	p.Dao.Set("google.verify.code", fm.Google, false)
	p.Dao.Set("baidu.verify.code", fm.Baidu, false)
	return gin.H{}, nil
}

func (p *Engine) getAdminSMTP(c *gin.Context) (interface{}, error) {
	var m SMTP
	p.Dao.Get("site.smtp", &m)
	return m, nil
}

func (p *Engine) postAdminSMTP(c *gin.Context) (interface{}, error) {
	var fm fmSMTP
	if err := c.Bind(&fm); err != nil {
		return nil, err
	}

	p.Dao.Set("site.smtp", SMTP{
		Host:     fm.Host,
		Port:     fm.Port,
		Username: fm.Username,
		Password: fm.Password,
	}, true)
	return gin.H{}, nil
}

func (p *Engine) getAdminStatus(c *gin.Context) (interface{}, error) {
	status := gin.H{}
	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)

	status["os"] = map[string]string{
		"OS":      fmt.Sprintf("%s %s", runtime.GOOS, runtime.GOARCH),
		"go-lang": fmt.Sprintf("%s %s", runtime.Version(), runtime.GOROOT()),
		"host": fmt.Sprintf(
			"cpus(%d) memory(%.1fG)",
			runtime.NumCPU(),
			float64(mem.HeapSys)/1024.0/1024.0,
		),
	}
	status["jobs"] = gin.H{
		"tasks": p.Server.GetRegisteredTaskNames(),
	}

	return status, nil
}

func (p *Engine) getAdminUsers(c *gin.Context) (interface{}, error) {
	var users []User
	err := p.Db.
		Select([]string{"email", "name", "last_sign_in_at"}).
		Order("sign_in_count DESC").Find(&users).Error

	return users, err
}
