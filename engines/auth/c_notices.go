package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/kapmahc/lotus/web"
)

func (p *Engine) noticesIndex(c *gin.Context) (interface{}, error) {
	var items []Notice
	err := p.Db.Order("updated_at DESC").Find(&items).Error
	return items, err
}

func (p *Engine) noticesCreate(c *gin.Context) (interface{}, error) {
	lang := c.MustGet(web.LOCALE).(string)
	var fm fmContent
	if err := c.Bind(&fm); err != nil {
		return nil, err
	}
	n := Notice{Lang: lang, Content: fm.Content}
	err := p.Db.Create(&n).Error
	return n, err
}

func (p *Engine) noticesShow(c *gin.Context) (interface{}, error) {
	var n Notice
	e := p.Db.Where("id = ?", c.Param("id")).Limit(1).Find(&n).Error
	return n, e
}

func (p *Engine) noticesUpdate(c *gin.Context) (interface{}, error) {
	var n Notice
	var fm fmContent
	if err := c.Bind(&fm); err != nil {
		return nil, err
	}
	e := p.Db.Where("id = ?", c.Param("id")).Limit(1).Find(&n).Error
	if e == nil {
		e = p.Db.Model(&n).Update("content", fm.Content).Error
	}
	return n, e
}

func (p *Engine) noticeDestroy(c *gin.Context) (interface{}, error) {
	e := p.Db.Where("id = ?", c.Param("id")).Delete(&Notice{}).Error
	return gin.H{}, e
}
