package auth

import (
	"fmt"
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
		defEmail := viper.GetString("server.manager")
		defPassword := "change-me"
		var email, password string
		fmt.Printf("Your email address[%s]:\n", defEmail)
		fmt.Scanf("%s", &email)

		fmt.Printf("Password[%s]:\n", defPassword)
		fmt.Scanf("%s", &password)
		if email == "" {
			email = defEmail
		}
		if password == "" {
			password = defPassword
		}

		user, err := p.Dao.AddEmailUser(
			language.AmericanEnglish.String(),
			email,
			"administrator",
			password,
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
