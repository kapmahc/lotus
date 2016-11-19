package shop

import (
	"github.com/gin-gonic/gin"
	"github.com/kapmahc/lotus/web"
)

type fmPaymentMethod struct {
	Type        string `form:"type" binding:"required,max=16"`
	Name        string `form:"name" binding:"required,max=255"`
	Description string `form:"description"`
	Active      bool   `form:"active"`
}

func (p *Engine) paymentMethodsIndex(c *gin.Context) (interface{}, error) {
	var items []PaymentMethod
	if err := p.Db.Order("updated_at DESC").Find(&items).Error; err != nil {
		return nil, err
	}

	return items, nil
}

func (p *Engine) paymentMethodsCreate(c *gin.Context) (interface{}, error) {
	lang := c.MustGet(web.LOCALE).(string)
	var fm fmPaymentMethod
	if err := c.Bind(&fm); err != nil {
		return nil, err
	}
	var count int
	if err := p.Db.Model(&PaymentMethod{}).
		Where("name = ?", fm.Name).
		Count(&count).Error; err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, p.I18n.E(lang, "messages.name-already-exists")
	}

	item := PaymentMethod{
		Type:        fm.Type,
		Name:        fm.Name,
		Description: fm.Description,
		Active:      fm.Active,
	}
	err := p.Db.Create(&item).Error
	return item, err
}

func (p *Engine) paymentMethodsShow(c *gin.Context) (interface{}, error) {
	var item PaymentMethod
	e := p.Db.Where("id = ?", c.Param("id")).Limit(1).Find(&item).Error
	return item, e
}

func (p *Engine) paymentMethodsUpdate(c *gin.Context) (interface{}, error) {
	lang := c.MustGet(web.LOCALE).(string)

	var fm fmPaymentMethod
	if err := c.Bind(&fm); err != nil {
		return nil, err
	}

	var item PaymentMethod
	if err := p.Db.Where("id = ?", c.Param("id")).Limit(1).Find(&item).Error; err != nil {
		return nil, err
	}

	var count int
	if err := p.Db.Model(&PaymentMethod{}).
		Where("name = ? AND id != ?", fm.Name, item.ID).
		Count(&count).Error; err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, p.I18n.E(lang, "messages.name-already-exists")
	}

	err := p.Db.Model(&item).Updates(map[string]interface{}{
		"type":        fm.Type,
		"name":        fm.Name,
		"description": fm.Description,
		"active":      fm.Active,
	}).Error

	return item, err
}

func (p *Engine) paymentMethodsDestroy(c *gin.Context) (interface{}, error) {
	id := c.Param("id")

	var item PaymentMethod
	lang := c.MustGet(web.LOCALE).(string)

	if err := p.Db.Where("id = ?", id).Limit(1).Find(&item).Error; err != nil {
		return nil, err
	}
	for _, obj := range []interface{}{
		&Payment{},
	} {
		var count int
		if err := p.Db.Model(obj).
			Where("payment_method_id = ?", id).Count(&count).Error; err != nil {
			return nil, err
		}
		if count > 0 {
			return nil, p.I18n.E(lang, "messages.in-using")
		}
	}

	err := p.Db.Delete(&item).Error

	return item, err
}