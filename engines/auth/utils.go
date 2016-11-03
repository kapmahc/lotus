package auth

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

func migrationRoot() string {
	root := filepath.Join("db", viper.GetString("database.driver"), "migrations")
	os.MkdirAll(root, 0700)
	return root
}

//IsProduction is production mode?
func IsProduction() bool {
	return viper.GetString("env") == "production"
}

//OpenDatabase open database
func OpenDatabase() (*gorm.DB, error) {
	//postgresql: "user=%s password=%s host=%s port=%d dbname=%s sslmode=%s"
	args := ""
	for k, v := range viper.GetStringMapString("database.args") {
		args += fmt.Sprintf(" %s=%s ", k, v)
	}
	db, err := gorm.Open(viper.GetString("database.driver"), args)
	if err != nil {
		return nil, err
	}
	if !IsProduction() {
		db.LogMode(true)
	}

	if err := db.DB().Ping(); err != nil {
		return nil, err
	}

	db.DB().SetMaxIdleConns(viper.GetInt("database.pool.max_idle"))
	db.DB().SetMaxOpenConns(viper.GetInt("database.pool.max_open"))
	return db, nil

}

//OpenRedis get redis connection pool
func OpenRedis() *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, e := redis.Dial(
				"tcp",
				fmt.Sprintf(
					"%s:%d",
					viper.GetString("redis.host"),
					viper.GetInt("redis.port"),
				),
			)
			if e != nil {
				return nil, e
			}
			if _, e = c.Do("SELECT", viper.GetInt("redis.db")); e != nil {
				c.Close()
				return nil, e
			}
			return c, nil
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}
