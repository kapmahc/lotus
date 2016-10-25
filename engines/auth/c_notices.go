package auth

import "github.com/gin-gonic/gin"

func (p *Engine) getNotices(c *gin.Context) (interface{}, error) {
	var items []Notice
	err := p.Db.Order("updated_at DESC").Find(&items).Error
	return items, err
}
