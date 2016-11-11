package auth

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kapmahc/lotus/web"
	"golang.org/x/text/language"
)

func (p *Engine) getLocales(c *gin.Context) (interface{}, error) {
	lang, err := language.Parse(c.Param("lang"))
	if err != nil {
		return nil, err
	}
	return p.I18n.Locales(lang.String()), nil
}

func (p *Engine) getLayout(c *gin.Context) {
	lang := c.MustGet("locale").(string)
	ret := gin.H{}
	for _, k := range []string{"title", "subTitle", "keywords", "description", "copyright"} {
		ret[k] = p.I18n.T(lang, fmt.Sprintf("site.%s", k))
	}

	author := gin.H{}
	for _, k := range []string{"email", "name"} {
		var v string
		p.Dao.Get(fmt.Sprintf("site.author.%s", k), &v)
		author[k] = v
	}
	ret["author"] = author

	links := gin.H{}
	for _, k := range []string{"top", "bottom"} {
		var v []web.Link
		p.Dao.Get(fmt.Sprintf("site.links.%s", k), &v)
		links[k] = v
	}
	ret["links"] = links

	ret["languages"] = p.I18n.Languages()

	c.JSON(http.StatusOK, ret)
}

func (p *Engine) getDashboard(c *gin.Context) {
	//TODO
}
