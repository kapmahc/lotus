package forum

import (
	"github.com/gin-gonic/gin"
	"github.com/kapmahc/lotus/engines/auth"
)

func (p *Engine) articlesIndex(c *gin.Context) (interface{}, error) {
	var items []Article
	err := p.Db.Order("updated_at DESC").Find(&items).Error
	return items, err
}

func (p *Engine) articlesCreate(c *gin.Context) (interface{}, error) {
	user := c.MustGet(auth.CurrentUser).(*auth.User)
	var fm fmArticle
	if err := c.Bind(&fm); err != nil {
		return nil, err
	}
	item := Article{
		Title:   fm.Title,
		Summary: fm.Summary,
		Body:    fm.Body,
		UserID:  user.ID,
	}
	if err := p.Db.Create(&item).Error; err != nil {
		return nil, err
	}
	var tags []Tag
	for _, t := range fm.Tags {
		var tag Tag
		if err := p.Db.Find("id = ?", t).Limit(1).Find(&tag).Error; err != nil {
			return nil, err
		}
		tags = append(tags, tag)
	}
	err := p.Db.Model(&item).Association("Tags").Append(tags).Error
	return item, err
}

func (p *Engine) articlesShow(c *gin.Context) (interface{}, error) {
	var item Article
	if err := p.Db.Where("id = ?", c.Param("id")).Limit(1).Find(&item).Error; err != nil {
		return nil, err
	}
	if err := p.Db.Model(item).Related(&item.Tags, "Tags").Error; err != nil {
		return nil, err
	}
	if err := p.Db.Model(item).Related(&item.Comments, "Comments").Error; err != nil {
		return nil, err
	}
	if err := p.Db.Model(item).Related(&item.User, "User").Error; err != nil {
		return nil, err
	}
	return item, nil
}

func (p *Engine) articlesUpdate(c *gin.Context) (interface{}, error) {
	var fm fmArticle
	if err := c.Bind(&fm); err != nil {
		return nil, err
	}

	var item Article
	if err := p.Db.Where("id = ?", c.Param("id")).Limit(1).Find(&item).Error; err != nil {
		return nil, err
	}

	if err := p.Db.Model(&item).Updates(Article{
		Title:   fm.Title,
		Summary: fm.Summary,
		Body:    fm.Body,
	}).Error; err != nil {
		return nil, err
	}
	var tags []Tag
	for _, t := range fm.Tags {
		var tag Tag
		if err := p.Db.Find("id = ?", t).Limit(1).Find(&tag).Error; err != nil {
			return nil, err
		}
		tags = append(tags, tag)
	}
	err := p.Db.Model(&item).Association("Tags").Replace(tags).Error
	return item, err
}

func (p *Engine) articlesDestroy(c *gin.Context) (interface{}, error) {
	var item Article
	if err := p.Db.Where("id = ?", c.Param("id")).Limit(1).Find(&item).Error; err != nil {
		return nil, err
	}
	if err := p.Db.Model(&item).Association("Tags").Clear().Error; err != nil {
		return nil, err
	}
	if err := p.Db.Model(&item).Association("Comments").Clear().Error; err != nil {
		return nil, err
	}
	if err := p.Db.Delete(&item).Error; err != nil {
		return nil, err
	}

	return item, nil
}
