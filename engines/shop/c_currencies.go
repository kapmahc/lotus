package shop

import "github.com/gin-gonic/gin"

type fmCurrency struct {
	Rate   float64 `form:"rate" binding:"required"`
	Active bool    `form:"active"`
}

func (p *Engine) currenciesIndex(c *gin.Context) (interface{}, error) {
	var items []Currency
	if err := p.Db.Order("updated_at DESC").Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

func (p *Engine) currenciesShow(c *gin.Context) (interface{}, error) {
	var item Currency
	e := p.Db.Where("id = ?", c.Param("id")).Limit(1).Find(&item).Error
	return item, e
}

func (p *Engine) currenciesUpdate(c *gin.Context) (interface{}, error) {
	var fm fmCurrency
	if err := c.Bind(&fm); err != nil {
		return nil, err
	}

	var item Currency
	if err := p.Db.Where("id = ?", c.Param("id")).Limit(1).Find(&item).Error; err != nil {
		return nil, err
	}

	err := p.Db.Model(&item).Updates(map[string]interface{}{
		"active": fm.Active,
		"rate":   fm.Rate,
	}).Error

	return item, err
}
