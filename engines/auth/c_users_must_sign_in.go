package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/kapmahc/lotus/web"
)

func (p *Engine) deleteUsersSignOut(c *gin.Context) (interface{}, error) {
	user := c.MustGet(CurrentUser).(*User)
	lang := c.MustGet(web.LOCALE).(string)
	p.Dao.Logf(user.ID, lang, "auth.logs.sign-out")
	return user, nil
}

func (p *Engine) getUsersLogs(c *gin.Context) (interface{}, error) {
	user := c.MustGet(CurrentUser).(*User)
	var logs []Log
	err := p.Db.Model(&Log{}).
		Where("user_id = ?", user.ID).
		Order("id DESC").Limit(120).Find(&logs).Error
	return logs, err
}
func (p *Engine) getUsersInfo(c *gin.Context) (interface{}, error) {
	user := c.MustGet(CurrentUser).(*User)
	return user, nil
}

func (p *Engine) postUsersInfo(c *gin.Context) (interface{}, error) {
	user := c.MustGet(CurrentUser).(*User)

	var fm fmInfo
	if err := c.Bind(&fm); err != nil {
		return nil, err
	}
	if err := p.Db.Model(user).
		Updates(User{
			Name: fm.Name,
			Home: fm.Home,
			Logo: fm.Logo,
		}).Error; err != nil {
		return nil, err
	}

	return gin.H{
		"user": user,
	}, nil
}

func (p *Engine) postUsersChangePassword(c *gin.Context) (interface{}, error) {
	user := c.MustGet(CurrentUser).(*User)
	lang := c.MustGet(web.LOCALE).(string)

	var fm fmChangePassword
	if err := c.Bind(&fm); err != nil {
		return nil, err
	}
	if !p.Hmac.Chk([]byte(fm.CurrentPassword), user.Password) {
		return nil, p.I18n.E(lang, "auth.messages.email-password-not-match")
	}
	if err := p.Db.Model(user).
		Update("password", p.Hmac.Sum([]byte(fm.NewPassword))).Error; err != nil {
		return nil, err
	}

	return gin.H{
		"user": user,
	}, nil
}
