package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/kapmahc/lotus/web"
)

func (p *Engine) leavewordsIndex(c *gin.Context) (interface{}, error) {
	var items []Leaveword
	err := p.Db.Order("id DESC").Find(&items).Error
	return items, err
}

func (p *Engine) leavewordsCreate(c *gin.Context) (interface{}, error) {
	lang := c.MustGet(web.LOCALE).(string)
	var fm fmContent
	if err := c.Bind(&fm); err != nil {
		return nil, err
	}
	err := p.Db.Create(&Leaveword{Content: fm.Content}).Error
	return gin.H{
		"message": p.I18n.T(lang, "messages.success"),
	}, err
}

func (p *Engine) leavewordDestroy(c *gin.Context) (interface{}, error) {
	lang := c.MustGet(web.LOCALE).(string)
	e := p.Db.Where("id = ?", c.Param("id")).Delete(&Leaveword{}).Error
	return gin.H{
		"message": p.I18n.T(lang, "messages.success"),
	}, e
}
