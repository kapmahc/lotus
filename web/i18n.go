package web

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/text/language"

	"github.com/golang/glog"
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
	msg, err := p.Get(lang, code)
	if err == nil {
		return fmt.Sprintf(msg, args...)
	}
	glog.Error(err)
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

//Load load locales from file
func (p *I18n) Load(dir string) error {
	return filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		const ext = ".txt"
		name := info.Name()
		if info.Mode().IsRegular() && filepath.Ext(name) == ext {
			glog.Infof("Find locale file %s", path)
			lang := name[0 : len(name)-len(ext)]
			if _, err := language.Parse(lang); err != nil {
				return err
			}
			fd, err := os.Open(path)
			if err != nil {
				return err
			}
			defer fd.Close()

			scanner := bufio.NewScanner(fd)
			for scanner.Scan() {
				line := scanner.Text()
				idx := strings.Index(line, " = ")
				if idx > 0 {
					code := strings.TrimSpace(line[0:idx])
					msg := strings.TrimSpace(line[idx+3:])
					var count int
					err = p.Db.Model(&Locale{}).Where("lang = ? AND code = ?", lang, code).Count(&count).Error
					if err == nil && count == 0 {
						err = p.Db.Create(&Locale{Lang: lang, Code: code, Message: msg}).Error
					}
					if err != nil {
						return err
					}
				}
			}

			return scanner.Err()
		}
		return nil
	})

}
