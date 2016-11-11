package web

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/text/language"
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

//Handler locale handler
func (p *I18n) Handler(c *gin.Context) {

	const key = "locale"
	// 1. Check URL arguments.
	lng := c.Request.URL.Query().Get(key)

	// 2. Get language information from cookies.
	if len(lng) == 0 {
		if ck, er := c.Request.Cookie(key); er == nil {
			lng = ck.Value
		}
	}

	// 3. Get language information from 'Accept-Language'.
	if len(lng) == 0 {
		al := c.Request.Header.Get("Accept-Language")
		if len(al) > 4 {
			lng = al[:5]
		}
	}

	tag, err := language.Parse(lng)
	if err != nil {
		log.Error(err)
		tag = language.AmericanEnglish
	}

	// Write cookie
	// if write {
	// 	http.SetCookie(c.Writer, &http.Cookie{
	// 		Name:    key,
	// 		Value:   tag.String(),
	// 		Expires: time.Now().AddDate(10, 1, 1),
	// 		Path:    "/",
	// 	})
	// }

	c.Set(key, tag.String())

}

//T translate
func (p *I18n) T(lang string, code string, args ...interface{}) string {
	msg, err := p.Get(lang, code)
	if err == nil {
		return fmt.Sprintf(msg, args...)
	}
	log.Error(err)
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
	err := p.Db.Select([]string{"message"}).Where("lang = ? AND code = ?", lang, code).First(&l).Error
	return l.Message, err
}

//Del del locale
func (p *I18n) Del(lang string, code string) error {
	return p.Db.Where("lang = ? AND code = ?", lang, code).Delete(Locale{}).Error
}

//Codes get codes
func (p *I18n) Codes(lang string) ([]string, error) {
	var keys []string
	err := p.Db.Model(&Locale{}).Where("lang = ?", lang).Order("code ASC").Pluck("code", &keys).Error
	return keys, err
}

//Languages supported languages
func (p *I18n) Languages() []string {
	var keys []string
	if err := p.Db.Model(&Locale{}).Pluck("DISTINCT lang", &keys).Error; err != nil {
		log.Error(err)
	}
	return keys
}

//Locales list locales by lang
func (p *I18n) Locales(lang string) map[string]interface{} {
	items := make(map[string]interface{})
	var locales []Locale
	if err := p.Db.
		Select([]string{"code", "message"}).
		Where("lang = ?", lang).
		Order("code ASC").
		Find(&locales).Error; err != nil {
		log.Error(err)
	}

	for _, l := range locales {
		codes := strings.Split(l.Code, ".")
		tmp := items
		for i, c := range codes {
			if i+1 == len(codes) {
				tmp[c] = l.Message
			} else {
				if tmp[c] == nil {
					tmp[c] = make(map[string]interface{})
				}
				tmp = tmp[c].(map[string]interface{})
			}
		}
	}

	return items
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
			log.Infof("Find locale file %s", path)
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
