package shop

import (
	"github.com/gin-gonic/gin"
	"github.com/kapmahc/lotus/web"
)

type fmStateNew struct {
	Name      string `form:"name" binding:"required,max=255"`
	CountryID uint   `form:"country_id" binding:"required"`
}
type fmStateEdit struct {
	Name string `form:"name" binding:"required,max=255"`
}

func (p *Engine) statesIndex(c *gin.Context) (interface{}, error) {
	var items []State
	if err := p.Db.Order("updated_at DESC").Find(&items).Error; err != nil {
		return nil, err
	}
	for i, item := range items {
		if err := p.Db.
			Model(&item).
			Association("Country").
			Find(&items[i].Country).Error; err != nil {
			return nil, err
		}
	}

	return items, nil
}

func (p *Engine) statesCreate(c *gin.Context) (interface{}, error) {
	lang := c.MustGet(web.LOCALE).(string)
	var fm fmStateNew
	if err := c.Bind(&fm); err != nil {
		return nil, err
	}
	var count int
	if err := p.Db.Model(&State{}).
		Where("name = ? AND country_id = ?", fm.Name, fm.CountryID).
		Count(&count).Error; err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, p.I18n.E(lang, "messages.name-already-exists")
	}

	item := State{Name: fm.Name, CountryID: fm.CountryID}
	err := p.Db.Create(&item).Error
	return item, err
}

func (p *Engine) statesShow(c *gin.Context) (interface{}, error) {
	var item State
	e := p.Db.Where("id = ?", c.Param("id")).Limit(1).Find(&item).Error
	return item, e
}

func (p *Engine) statesUpdate(c *gin.Context) (interface{}, error) {
	lang := c.MustGet(web.LOCALE).(string)

	var fm fmStateEdit
	if err := c.Bind(&fm); err != nil {
		return nil, err
	}

	var item State
	if err := p.Db.Where("id = ?", c.Param("id")).Limit(1).Find(&item).Error; err != nil {
		return nil, err
	}

	var count int
	if err := p.Db.Model(&State{}).
		Where("name = ? AND country_id = ? AND id != ?", fm.Name, item.CountryID, item.ID).
		Count(&count).Error; err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, p.I18n.E(lang, "messages.name-already-exists")
	}

	err := p.Db.Model(&item).Update("name", fm.Name).Error

	return item, err
}

func (p *Engine) statesDestroy(c *gin.Context) (interface{}, error) {
	id := c.Param("id")
	lang := c.MustGet(web.LOCALE).(string)

	var item State
	if err := p.Db.Where("id = ?", id).Limit(1).Find(&item).Error; err != nil {
		return nil, err
	}

	for _, obj := range []interface{}{
		&Address{},
		&TaxRate{},
	} {
		var count int
		if err := p.Db.Model(obj).
			Where("state_id = ?", id).Count(&count).Error; err != nil {
			return nil, err
		}
		if count > 0 {
			return nil, p.I18n.E(lang, "messages.in-using")
		}
	}

	err := p.Db.Delete(&item).Error
	return item, err
}
