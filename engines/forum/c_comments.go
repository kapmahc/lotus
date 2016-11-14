package forum

import (
	"github.com/gin-gonic/gin"
	"github.com/kapmahc/lotus/engines/auth"
	"github.com/kapmahc/lotus/web"
)

func (p *Engine) commentsIndex(c *gin.Context) (interface{}, error) {
	var items []Comment
	err := p.Db.Order("updated_at DESC").Find(&items).Error
	return items, err
}

func (p *Engine) commentsCreate(c *gin.Context) (interface{}, error) {
	user := c.MustGet(auth.CurrentUser).(*auth.User)
	var fm fmComment
	if err := c.Bind(&fm); err != nil {
		return nil, err
	}
	item := Comment{UserID: user.ID, ArticleID: fm.ArticleID, Body: fm.Body}
	err := p.Db.Create(&item).Error
	return item, err
}

func (p *Engine) commentsShow(c *gin.Context) (interface{}, error) {
	var item Comment
	err := p.Db.Where("id = ?", c.Param("id")).Limit(1).Find(&item).Error
	return item, err
}

func (p *Engine) commentsUpdate(c *gin.Context) (interface{}, error) {
	lang := c.MustGet(web.LOCALE).(string)
	user := c.MustGet(auth.CurrentUser).(*auth.User)

	var fm fmComment
	if err := c.Bind(&fm); err != nil {
		return nil, err
	}

	var item Comment
	if err := p.Db.Where("id = ?", c.Param("id")).Limit(1).Find(&item).Error; err != nil {
		return nil, err
	}
	if !p._canComment(user, &item) {
		return nil, p.I18n.E(lang, "messages.bad-token")
	}

	err := p.Db.Model(&item).Update("body", fm.Body).Error
	return item, err
}

func (p *Engine) commentsDestroy(c *gin.Context) (interface{}, error) {

	var item Comment
	if err := p.Db.Where("id = ?", c.Param("id")).Limit(1).Find(&item).Error; err != nil {
		return nil, err
	}

	user := c.MustGet(auth.CurrentUser).(*auth.User)
	lang := c.MustGet(web.LOCALE).(string)
	if !p._canComment(user, &item) {
		return nil, p.I18n.E(lang, "messages.bad-token")
	}

	if err := p.Db.Delete(&item).Error; err != nil {
		return nil, err
	}

	return item, nil
}

func (p *Engine) _canComment(u *auth.User, c *Comment) bool {
	return u.ID == c.UserID || p.Dao.Is(u.ID, auth.RoleAdmin)
}
