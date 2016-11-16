package shop

import (
	"github.com/gin-gonic/gin"
	"github.com/kapmahc/lotus/web"
)

type fmCountry struct {
	Name   string `form:"name" binding:"required,max=255"`
	Active bool   `form:"active"`
}

func (p *Engine) countriesIndex(c *gin.Context) (interface{}, error) {
	var items []Country
	if err := p.Db.Order("updated_at DESC").Find(&items).Error; err != nil {
		return nil, err
	}
	for i, item := range items {
		if err := p.Db.
			Model(&item).Order("name ASC").
			Association("States").
			Find(&items[i].States).Error; err != nil {
			return nil, err
		}
	}

	return items, nil
}

func (p *Engine) countriesCreate(c *gin.Context) (interface{}, error) {
	lang := c.MustGet(web.LOCALE).(string)
	var fm fmCountry
	if err := c.Bind(&fm); err != nil {
		return nil, err
	}
	var count int
	if err := p.Db.Model(&Country{}).
		Where("name = ?", fm.Name).Count(&count).Error; err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, p.I18n.E(lang, "shop.messages.country-name-already-exists")
	}
	item := Country{Name: fm.Name, Active: fm.Active}
	err := p.Db.Create(&item).Error
	return item, err
}

func (p *Engine) countriesShow(c *gin.Context) (interface{}, error) {
	var item Country
	e := p.Db.Where("id = ?", c.Param("id")).Limit(1).Find(&item).Error
	return item, e
}

func (p *Engine) countriesUpdate(c *gin.Context) (interface{}, error) {
	lang := c.MustGet(web.LOCALE).(string)

	var fm fmCountry
	if err := c.Bind(&fm); err != nil {
		return nil, err
	}

	var item Country
	if err := p.Db.Where("id = ?", c.Param("id")).Limit(1).Find(&item).Error; err != nil {
		return nil, err
	}

	var count int
	if err := p.Db.Model(&Country{}).
		Where("name = ? AND id != ?", fm.Name, item.ID).Count(&count).Error; err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, p.I18n.E(lang, "shop.messages.country-name-already-exists")
	}

	err := p.Db.Model(&item).
		Updates(map[string]interface{}{"name": fm.Name, "active": fm.Active}).Error

	return item, err
}

func (p *Engine) countriesDestroy(c *gin.Context) (interface{}, error) {
	lang := c.MustGet(web.LOCALE).(string)
	var item Country
	if err := p.Db.Where("id = ?", c.Param("id")).Limit(1).Find(&item).Error; err != nil {
		return nil, err
	}
	if p.Db.Model(&item).Association("States").Count() > 0 {
		return nil, p.I18n.E(lang, "messages.in-using")
	}
	err := p.Db.Delete(&item).Error
	return item, err
}
