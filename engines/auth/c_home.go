package auth

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (p *Engine) getLocales(wrt http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	p.Render.JSON(wrt, http.StatusOK, p.I18n.Locales(vars["lang"]))
}

//
// func (p *Engine) getLocales(c *gin.Context) {
// 	c.JSON(http.StatusOK, p.I18n.Locales(c.Param("lang")))
// }
//
// func (p *Engine) getLayout(c *gin.Context) {
// 	lang := c.MustGet("locale").(string)
// 	ret := gin.H{}
// 	for _, k := range []string{"title", "subTitle", "keywords", "description", "copyright"} {
// 		ret[k] = p.I18n.T(lang, fmt.Sprintf("site.%s", k))
// 	}
//
// 	author := gin.H{}
// 	for _, k := range []string{"email", "name"} {
// 		var v string
// 		p.Dao.Get(fmt.Sprintf("site.author.%s", k), &v)
// 		author[k] = v
// 	}
// 	ret["author"] = author
//
// 	links := gin.H{}
// 	for _, k := range []string{"top", "bottom"} {
// 		var v []web.Link
// 		p.Dao.Get(fmt.Sprintf("site.links.%s", k), &v)
// 		links[k] = v
// 	}
// 	ret["links"] = links
//
// 	c.JSON(http.StatusOK, ret)
// }
//
// func (p *Engine) getDashboard(c *gin.Context) {
// 	//TODO
// }
