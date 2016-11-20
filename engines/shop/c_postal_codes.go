package shop

import (
	"github.com/gin-gonic/gin"
	"github.com/kapmahc/lotus/web"
)

func (p *Engine) postalCodesIndex(c *gin.Context) (interface{}, error) {
	page, size, offset := web.ParsePager(c, 120)
	var total int64
	if err := p.Db.Model(&PostalCode{}).Count(&total).Error; err != nil {
		return nil, err
	}

	var items []PostalCode
	if err := p.Db.Order("cid ASC").
		Offset(offset).Limit(size).
		Find(&items).Error; err != nil {
		return nil, err
	}
	return gin.H{
		"items": items,
		"pager": web.Paginator(page, size, total),
	}, nil
}
