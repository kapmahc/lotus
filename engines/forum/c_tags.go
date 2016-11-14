package forum

import (
	"github.com/gin-gonic/gin"
	"github.com/kapmahc/lotus/web"
)

func (p *Engine) tagsIndex(c *gin.Context) (interface{}, error) {
	var items []Tag
	err := p.Db.Order("updated_at DESC").Find(&items).Error
	return items, err
}

func (p *Engine) tagsCreate(c *gin.Context) (interface{}, error) {
	lang := c.MustGet(web.LOCALE).(string)
	var fm fmTag
	if err := c.Bind(&fm); err != nil {
		return nil, err
	}
	var count int
	if err := p.Db.Where("name = ?", fm.Name).Count(&count).Error; err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, p.I18n.E(lang, "forum.messages.tag-name-already-exists")
	}
	item := Tag{Name: fm.Name}
	err := p.Db.Create(&item).Error
	return item, err
}

func (p *Engine) tagsShow(c *gin.Context) (interface{}, error) {
	var item Tag
	e := p.Db.Where("id = ?", c.Param("id")).Limit(1).Find(&item).Error
	return item, e
}

func (p *Engine) tagsUpdate(c *gin.Context) (interface{}, error) {
	lang := c.MustGet(web.LOCALE).(string)

	var fm fmTag
	if err := c.Bind(&fm); err != nil {
		return nil, err
	}

	var item Tag
	if err := p.Db.Where("id = ?", c.Param("id")).Limit(1).Find(&item).Error; err != nil {
		return nil, err
	}

	var count int
	if err := p.Db.Where("name = ?", fm.Name).Count(&count).Error; err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, p.I18n.E(lang, "forum.messages.tag-name-already-exists")
	}

	err := p.Db.Model(&item).Update("name", fm.Name).Error
	return item, err
}

func (p *Engine) tagsDestroy(c *gin.Context) (interface{}, error) {
	var item Tag
	if err := p.Db.Where("id = ?", c.Param("id")).Limit(1).Find(&item).Error; err != nil {
		return nil, err
	}
	if err := p.Db.Model(&item).Association("Articles").Clear().Error; err != nil {
		return nil, err
	}
	if err := p.Db.Delete(&item).Error; err != nil {
		return nil, err
	}

	return item, nil
}
