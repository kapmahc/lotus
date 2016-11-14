package reading

import (
	"github.com/gin-gonic/gin"
	"github.com/kapmahc/lotus/engines/auth"
	"github.com/kapmahc/lotus/web"
)

func (p *Engine) notesIndex(c *gin.Context) (interface{}, error) {
	var items []Note
	err := p.Db.Order("updated_at DESC").Find(&items).Error
	return items, err
}

func (p *Engine) notesCreate(c *gin.Context) (interface{}, error) {
	user := c.MustGet(auth.CurrentUser).(*auth.User)
	var fm fmNote
	if err := c.Bind(&fm); err != nil {
		return nil, err
	}
	item := Note{BookID: fm.BookID, UserID: user.ID, Body: fm.Body}
	err := p.Db.Create(&item).Error
	return item, err
}

func (p *Engine) notesShow(c *gin.Context) (interface{}, error) {
	var item Note
	e := p.Db.Where("id = ?", c.Param("id")).Limit(1).Find(&item).Error
	return item, e
}

func (p *Engine) notesUpdate(c *gin.Context) (interface{}, error) {
	lang := c.MustGet(web.LOCALE).(string)
	user := c.MustGet(auth.CurrentUser).(*auth.User)

	var fm fmNote
	if err := c.Bind(&fm); err != nil {
		return nil, err
	}

	var item Note
	if err := p.Db.Where("id = ?", c.Param("id")).Limit(1).Find(&item).Error; err != nil {
		return nil, err
	}
	if !p._canNote(user, &item) {
		return nil, p.I18n.E(lang, "messages.bad-token")
	}

	err := p.Db.Model(&item).Update("body", fm.Body).Error
	return item, err
}

func (p *Engine) notesDestroy(c *gin.Context) (interface{}, error) {
	lang := c.MustGet(web.LOCALE).(string)
	var item Note
	if err := p.Db.Where("id = ?", c.Param("id")).Limit(1).Find(&item).Error; err != nil {
		return nil, err
	}
	user := c.MustGet(auth.CurrentUser).(*auth.User)
	if !p._canNote(user, &item) {
		return nil, p.I18n.E(lang, "messages.bad-token")
	}
	if err := p.Db.Delete(&item).Error; err != nil {
		return nil, err
	}

	return item, nil
}

func (p *Engine) _canNote(u *auth.User, n *Note) bool {
	return u.ID == n.UserID || p.Dao.Is(u.ID, auth.RoleAdmin)
}
