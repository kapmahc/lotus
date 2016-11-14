package auth

import "github.com/gin-gonic/gin"

func (p *Engine) leavewordsIndex(c *gin.Context) (interface{}, error) {
	var items []Leaveword
	err := p.Db.Order("id DESC").Find(&items).Error
	return items, err
}

func (p *Engine) leavewordsCreate(c *gin.Context) (interface{}, error) {
	var fm fmContent
	if err := c.Bind(&fm); err != nil {
		return nil, err
	}
	lw := Leaveword{Content: fm.Content}
	err := p.Db.Create(&lw).Error
	return lw, err
}

func (p *Engine) leavewordDestroy(c *gin.Context) (interface{}, error) {
	e := p.Db.Where("id = ?", c.Param("id")).Delete(&Leaveword{}).Error
	return gin.H{}, e
}
