package web

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

//Locale locale model
type Locale struct {
	Model

	Lang    string
	Code    string
	Message string
}

//I18n i18n
type I18n struct {
	Db *gorm.DB `inject:""`
}

//T translate
func (p *I18n) T(lang string, code string, args ...interface{}) string {
	if msg, err := p.Get(lang, code); err == nil {
		return fmt.Sprintf(msg, args...)
	}
	return code
}

//Set set message
func (p *I18n) Set(lang string, code, message string) error {
	var l Locale
	var err error
	if p.Db.Where("lang = ? AND code = ?", lang, code).First(&l).RecordNotFound() {
		l.Lang = lang
		l.Code = code
		l.Message = message
		err = p.Db.Create(&l).Error
	} else {
		l.Message = message
		err = p.Db.Save(&l).Error
	}
	return err
}

//Get get message
func (p *I18n) Get(lang string, code string) (string, error) {
	var l Locale
	err := p.Db.Where("lang = ? AND code = ?", lang, code).First(&l).Error
	return l.Message, err
}

//Del del locale
func (p *I18n) Del(lang string, code string) error {
	return p.Db.Where("lang = ? AND code = ?", lang, code).Delete(Locale{}).Error
}

//Keys locale keys
func (p *I18n) Keys(lang string) ([]string, error) {
	var keys []string
	err := p.Db.Model(&Locale{}).Where("lang = ?", lang).Pluck("code", &keys).Error
	return keys, err
}
