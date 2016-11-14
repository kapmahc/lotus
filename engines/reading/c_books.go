package reading

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"path"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/kapmahc/epub"
)

func (p *Engine) booksIndex(c *gin.Context) (interface{}, error) {
	var items []Book
	err := p.Db.Order("updated_at DESC").Find(&items).Error
	return items, err
}

func (p *Engine) _points2html(id uint, points []epub.NavPoint) string {
	str := "<ol>"
	for _, pt := range points {
		str += fmt.Sprintf(
			`<li><a href="/reading/books/%d/%s" target="_blank">%s</a></li>`,
			id,
			pt.Content.Src,
			pt.Text,
		)
		str += p._points2html(id, pt.Points)
	}
	str += "</ol>"
	return str
}

func (p *Engine) booksShow(c *gin.Context) {
	var item Book
	if err := p.Db.Where("id = ?", c.Param("id")).Limit(1).Find(&item).Error; err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	bk, err := epub.Open(item.File)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	defer bk.Close()
	page := c.Param("page")
	if page == "" {
		c.Data(http.StatusOK, "", []byte(p._points2html(item.ID, bk.Ncx.Points)))
		return
	}

	fd, err := bk.Open(page)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	defer fd.Close()

	for _, ext := range []string{".xhtml", ".css", ".jpg"} {
		if path.Ext(page) == ext {
			for _, m := range bk.Opf.Manifest {
				if strings.HasPrefix(page, m.Href) {
					// p.Ctx.Output.ContentType(m.MediaType)
					if buf, err := ioutil.ReadAll(fd); err == nil {
						c.Data(http.StatusOK, m.MediaType, buf)
					} else {
						c.AbortWithError(http.StatusInternalServerError, err)
					}
					return
				}
			}
		}
	}

}

func (p *Engine) showBookPage() {
	// book := p.getBook()
	// bk, err := epub.Open(book.File)
	// p.Check(err)
	// defer bk.Close()
	// name := p.Ctx.Input.Param(":splat")
	//
	// fd, err := bk.Open(name)
	// p.Check(err)
	// defer fd.Close()

	// beego.Error("bad file", name)
	// p.Abort("404")

}

func (p *Engine) booksDestroy(c *gin.Context) (interface{}, error) {
	var item Book
	if err := p.Db.Where("id = ?", c.Param("id")).Limit(1).Find(&item).Error; err != nil {
		return nil, err
	}
	if err := p.Db.Model(&item).Association("Notes").Clear().Error; err != nil {
		return nil, err
	}
	if err := p.Db.Delete(&item).Error; err != nil {
		return nil, err
	}
	return item, nil
}
