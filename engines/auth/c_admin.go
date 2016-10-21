package auth

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/kapmahc/lotus/web"
)

func (p *Engine) getSiteStatus(c *gin.Context) (interface{}, error) {
	// TODO
	// ofo := os.Environ()
	return gin.H{
		"os": []string{
			fmt.Sprintf("%s %s", runtime.GOOS, runtime.GOARCH),
		},
	}, nil
}

func (p *Engine) getSiteBase(c *gin.Context) (interface{}, error) {
	lang := c.MustGet("locale").(string)
	info := make(map[string]string)
	for _, k := range []string{"title", "subTitle", "keywords", "description", "copyright"} {
		info[k] = p.I18n.T(lang, fmt.Sprintf("site.%s", k))
	}
	return info, nil
}

func (p *Engine) getSiteAuthor(c *gin.Context) (interface{}, error) {
	author := make(map[string]string)
	for _, k := range []string{"name", "email"} {
		var v string
		if err := p.Dao.Get(fmt.Sprintf("author.%s", k), &v); err != nil {
			p.Logger.Error(err)
		}
		author[k] = v
	}
	return author, nil
}

func (p *Engine) getSiteNav(c *gin.Context) (interface{}, error) {
	ret := make(map[string][]web.Link)
	for _, k := range []string{"top", "bottom"} {
		var links []web.Link
		if err := p.Dao.Get(fmt.Sprintf("links.%s", k), &links); err != nil {
			p.Logger.Error(err)
		}
		ret[k] = links
	}
	return ret, nil
}

func (p *Engine) postSiteBase(c *gin.Context) (interface{}, error) {
	lang := c.MustGet("locale").(string)
	for _, k := range []string{"title", "subTitle", "keywords", "description", "copyright"} {
		p.I18n.Store.Set(lang, fmt.Sprintf("site.%s", k), c.PostForm(k))
	}
	return gin.H{}, nil
}

func (p *Engine) postSiteAuthor(c *gin.Context) (interface{}, error) {
	for _, k := range []string{"name", "email"} {
		if err := p.Dao.Set(fmt.Sprintf("author.%s", k), c.PostForm(k), false); err != nil {
			return nil, err
		}
	}
	return gin.H{}, nil
}

func (p *Engine) postSiteNav(c *gin.Context) (interface{}, error) {
	for _, k := range []string{"top", "bottom"} {
		var links []web.Link
		for _, line := range strings.Split(c.PostForm(k), "\n") {
			lk := strings.Split(line, "=")
			if len(lk) != 2 {
				return nil, fmt.Errorf("bad links format: %s", line)
			}
			links = append(
				links,
				web.Link{Href: strings.TrimSpace(lk[0]), Label: strings.TrimSpace(lk[1])},
			)
		}
		if err := p.Dao.Set(fmt.Sprintf("links.%s", k), &links, false); err != nil {
			return nil, err
		}
	}
	return gin.H{}, nil
}
