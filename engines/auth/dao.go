package auth

import (
	"github.com/jinzhu/gorm"
	"github.com/kapmahc/lotus/web"
)

//Dao dao
type Dao struct {
	Db     *gorm.DB    `inject:""`
	Logger *web.Logger `inject:""`
	Cache  *web.Cache  `inject:""`
}

//Set set key-val
func (p *Dao) Set(key string, val interface{}, flag bool) error {
	return nil
}

//Get get val
func (p *Dao) Get(key string, val interface{}) error {
	return nil
}
