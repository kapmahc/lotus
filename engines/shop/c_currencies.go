package shop

import "github.com/gin-gonic/gin"

type fmCurrency struct {
	Active bool `form:"active"`
}

func (p *Engine) currenciesIndex(c *gin.Context) (interface{}, error) {
	var items []Currency
	if err := p.Db.Order("country ASC").Find(&items).Error; err != nil {
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
	var fm fmCountry
	if err := c.Bind(&fm); err != nil {
		return nil, err
	}

	var item Currency
	if err := p.Db.Where("id = ?", c.Param("id")).Limit(1).Find(&item).Error; err != nil {
		return nil, err
	}

	err := p.Db.Model(&item).Update("active", fm.Active).Error

	return item, err
}
