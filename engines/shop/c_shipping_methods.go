package shop

import (
	"github.com/gin-gonic/gin"
	"github.com/kapmahc/lotus/web"
)

type fmShippingMethod struct {
	Name     string `form:"name" binding:"required,max=255"`
	Tracking string `form:"tracking" binding:"required,max=255"`
	Active   bool   `form:"active"`
}

func (p *Engine) shippingMethodsIndex(c *gin.Context) (interface{}, error) {
	var items []ShippingMethod
	if err := p.Db.Order("updated_at DESC").Find(&items).Error; err != nil {
		return nil, err
	}

	return items, nil
}

func (p *Engine) shippingMethodsCreate(c *gin.Context) (interface{}, error) {
	lang := c.MustGet(web.LOCALE).(string)
	var fm fmShippingMethod
	if err := c.Bind(&fm); err != nil {
		return nil, err
	}
	var count int
	if err := p.Db.Model(&ShippingMethod{}).
		Where("name = ?", fm.Name).
		Count(&count).Error; err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, p.I18n.E(lang, "messages.name-already-exists")
	}

	item := ShippingMethod{
		Name:     fm.Name,
		Tracking: fm.Tracking,
		Active:   fm.Active,
	}
	err := p.Db.Create(&item).Error
	return item, err
}

func (p *Engine) shippingMethodsShow(c *gin.Context) (interface{}, error) {
	var item ShippingMethod
	e := p.Db.Where("id = ?", c.Param("id")).Limit(1).Find(&item).Error
	return item, e
}

func (p *Engine) shippingMethodsUpdate(c *gin.Context) (interface{}, error) {
	lang := c.MustGet(web.LOCALE).(string)

	var fm fmShippingMethod
	if err := c.Bind(&fm); err != nil {
		return nil, err
	}

	var item ShippingMethod
	if err := p.Db.Where("id = ?", c.Param("id")).Limit(1).Find(&item).Error; err != nil {
		return nil, err
	}

	var count int
	if err := p.Db.Model(&ShippingMethod{}).
		Where("name = ? AND id != ?", fm.Name, item.ID).
		Count(&count).Error; err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, p.I18n.E(lang, "messages.name-already-exists")
	}

	err := p.Db.Model(&item).Updates(map[string]interface{}{
		"name":     fm.Name,
		"tracking": fm.Tracking,
		"active":   fm.Active,
	}).Error

	return item, err
}

func (p *Engine) shippingMethodsDestroy(c *gin.Context) (interface{}, error) {
	id := c.Param("id")

	var item ShippingMethod
	lang := c.MustGet(web.LOCALE).(string)

	if err := p.Db.Where("id = ?", id).Limit(1).Find(&item).Error; err != nil {
		return nil, err
	}
	for _, obj := range []interface{}{
		&Shipment{},
	} {
		var count int
		if err := p.Db.Model(obj).
			Where("shipping_method_id = ?", id).Count(&count).Error; err != nil {
			return nil, err
		}
		if count > 0 {
			return nil, p.I18n.E(lang, "messages.in-using")
		}
	}

	err := p.Db.Delete(&item).Error

	return item, err
}
