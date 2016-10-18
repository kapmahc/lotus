package i18n

import (
	"time"

	"github.com/jinzhu/gorm"
	logging "github.com/op/go-logging"
)

//Migrate migrate database
func Migrate(db *gorm.DB) {
	db.AutoMigrate(&Locale{})
	db.Model(&Locale{}).AddUniqueIndex("idx_locales_lang_code", "lang", "code")
}

//Locale locale model
type Locale struct {
	ID        uint   `gorm:"primary_key"`
	Lang      string `gorm:"not null;type:varchar(8);index"`
	Code      string `gorm:"not null;index;type:VARCHAR(255)"`
	Message   string `gorm:"not null;type:varchar(800)"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

//GormStore db provider
type GormStore struct {
	Db     *gorm.DB        `inject:""`
	Logger *logging.Logger `inject:""`
}

//Set set locale
func (p *GormStore) Set(lng string, code, message string) {
	var l Locale
	var err error
	if p.Db.Where("lang = ? AND code = ?", lng, code).First(&l).RecordNotFound() {
		l.Lang = lng
		l.Code = code
		l.Message = message
		err = p.Db.Create(&l).Error
	} else {
		l.Message = message
		err = p.Db.Save(&l).Error
	}
	if err != nil {
		p.Logger.Error(err)
	}
}

//Get get locale
func (p *GormStore) Get(lng string, code string) string {
	var l Locale
	if err := p.Db.Where("lang = ? AND code = ?", lng, code).First(&l).Error; err != nil {
		p.Logger.Error(err)
	}
	return l.Message

}

//Del del locale
func (p *GormStore) Del(lng string, code string) {
	if err := p.Db.Where("lang = ? AND code = ?", lng, code).Delete(Locale{}).Error; err != nil {
		p.Logger.Error(err)
	}
}

//Keys list locale keys
func (p *GormStore) Keys(lng string) ([]string, error) {
	var keys []string
	err := p.Db.Model(&Locale{}).Where("lang = ?", lng).Pluck("code", &keys).Error

	return keys, err
}
