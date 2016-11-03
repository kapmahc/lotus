package web

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"golang.org/x/text/language"

	"github.com/gin-gonic/gin"
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
	Db     *gorm.DB `inject:""`
	Logger *Logger  `inject:""`
}

//Handler locale handler
func (p *I18n) Handler(c *gin.Context) {

	write := false
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
		write = true
		al := c.Request.Header.Get("Accept-Language")
		if len(al) > 4 {
			lng = al[:5]
		}
	}

	tag, err := language.Parse(lng)
	if err != nil {
		p.Logger.Error("parse locale: %v", err)
		tag = language.AmericanEnglish
		write = true
	}

	// Write cookie
	if write {
		http.SetCookie(c.Writer, &http.Cookie{
			Name:    key,
			Value:   tag.String(),
			Expires: time.Now().AddDate(10, 1, 1),
			Path:    "/",
		})
	}

	c.Set(key, tag.String())

}

//T translate
func (p *I18n) T(lang string, code string, args ...interface{}) string {
	msg, err := p.Get(lang, code)
	if err == nil {
		return fmt.Sprintf(msg, args...)
	}
	p.Logger.Error("find locale: %s", err)
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

//Codes get codes
func (p *I18n) Codes(lang string) ([]string, error) {
	var keys []string
	err := p.Db.Model(&Locale{}).Where("lang = ?", lang).Order("code ASC").Pluck("code", &keys).Error
	return keys, err
}

//Languages supported languages
func (p *I18n) Languages() ([]string, error) {
	var keys []string
	err := p.Db.Model(&Locale{}).Pluck("DISTINCT lang", &keys).Error
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
			p.Logger.Info("Find locale file %s", path)
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