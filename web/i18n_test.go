package web_test

import (
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/kapmahc/lotus/web"
	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/text/language"
)

var lang = language.SimplifiedChinese.String()

func TestDatabase(t *testing.T) {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		t.Fatal(err)
	}
	db.LogMode(true)
	db.AutoMigrate(&web.Locale{})
	db.Model(&web.Locale{}).AddUniqueIndex("idx_locales_lang_code", "lang", "code")

	p := web.I18n{Db: db}
	key := "hello"
	val := "你好"
	p.Set(lang, key, val)
	p.Set(lang, key+".1", val)
	if val1, _ := p.Get(lang, key); val != val1 {
		t.Errorf("want %s, get %s", val, val1)
	}
	ks, err := p.Codes(lang)
	if err != nil {
		t.Fatal(err)
	}
	if len(ks) == 0 {
		t.Errorf("empty keys")
	} else {
		t.Log(ks)
	}
	p.Del(lang, key)

	if err = p.Load("../locales"); err != nil {
		t.Fatal(err)
	}
	t.Log(p.T("zh-Hans", "buttons.submit"))

	t.Log(p.Languages())
}
