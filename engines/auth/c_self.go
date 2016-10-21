package auth

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func (p *Engine) getSelfLogs(c *gin.Context) (interface{}, error) {
	user := c.MustGet("user").(*User)
	var logs []Log
	err := p.Db.Where("user_id = ?", user.ID).Order("id DESC").Find(&logs).Limit(64).Error
	return logs, err
}

func (p *Engine) getSelfProfile(c *gin.Context) (interface{}, error) {
	return c.MustGet("user").(*User), nil
}

type fmUserPassword struct {
	CurrentPassword      string `form:"currentPassword" binding:"required"`
	Password             string `form:"password" binding:"max=50,min=8"`
	PasswordConfirmation string `form:"passwordConfirmation" binding:"eqfield=Password"`
}

func (p *Engine) postSelfPassword(c *gin.Context) (interface{}, error) {
	user := c.MustGet("user").(*User)
	var fm fmUserPassword
	if err := c.Bind(&fm); err != nil {
		return nil, err
	}
	if !p.Encryptor.Chk([]byte(fm.Password), user.Password) {
		return nil, errors.New("wrong current password")
	}

	err := p.Db.Updates(map[string]interface{}{
		"password": p.Encryptor.Sum([]byte(fm.Password)),
	}).Error
	if err == nil {
		p.Dao.Log(user.ID, p.I18n.T(c.MustGet("locale").(string), "log.auth.change-password"))
	}
	return gin.H{}, err
}

type fmUserProfile struct {
	Name string `form:"name" binding:"max=32,min=2"`
	Logo string `form:"logo" binding:"max=255"`
	Home string `form:"home" binding:"max=255"`
}

func (p *Engine) postSelfProfile(c *gin.Context) (interface{}, error) {
	user := c.MustGet("user").(*User)
	var fm fmUserProfile
	if err := c.Bind(&fm); err != nil {
		return nil, err
	}

	err := p.Db.Updates(map[string]interface{}{
		"name": fm.Name,
		"logo": fm.Logo,
		"home": fm.Home,
	}).Error
	if err == nil {
		p.Dao.Log(user.ID, p.I18n.T(c.MustGet("locale").(string), "log.auth.update-profile"))
	}
	return gin.H{}, err
}
