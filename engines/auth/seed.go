package auth

import (
	"time"

	"github.com/spf13/viper"
	"golang.org/x/text/language"
)

//Seed Insert seed data
func (p *Engine) Seed() error {
	var count int
	if err := p.Db.Model(&User{}).Count(&count).Error; err != nil {
		return err
	}
	if count == 0 {
		email := viper.GetString("server.manager")
		user, err := p.Dao.AddEmailUser(
			language.AmericanEnglish.String(),
			email,
			"administrator",
			"change-me",
		)
		if err != nil {
			return err
		}
		if err = p.Db.Model(user).Update("confirmed_at", time.Now).Error; err != nil {
			return err
		}
		for _, name := range []string{RoleAdmin, RoleRoot} {
			role, err := p.Dao.Role(name, DefaultResourceType, DefaultResourceID)
			if err != nil {
				return err
			}
			if err = p.Dao.Allow(role.ID, user.ID, 100, 0, 0); err != nil {
				return err
			}
		}
	}
	return nil
}
