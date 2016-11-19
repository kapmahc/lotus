package shop

import "github.com/gin-gonic/gin"

func (p *Engine) postalCodesIndex(c *gin.Context) (interface{}, error) {
	var items []PostalCode
	if err := p.Db.Order("updated_at DESC").Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}
