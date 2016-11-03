package auth

import (
	"crypto/aes"
	"crypto/sha512"
	"fmt"
	"log/syslog"
	"time"

	"github.com/SermoDigital/jose/crypto"
	"github.com/facebookgo/inject"
	"github.com/garyburd/redigo/redis"
	"github.com/kapmahc/lotus/web"
	"github.com/spf13/viper"
)

//Init init ioc objects
func (p *Engine) Init(inj *inject.Graph) error {
	priority := syslog.LOG_DEBUG
	if web.IsProduction() {
		priority = syslog.LOG_INFO
	}
	app := viper.GetString("app.name")
	wlg, err := syslog.New(priority, app+"-web")
	if err != nil {
		return err
	}
	jlg, err := syslog.New(priority, app+"-jobs")
	if err != nil {
		return err
	}

	db, err := OpenDatabase()
	if err != nil {
		return err
	}

	cip, err := aes.NewCipher([]byte(viper.GetString("secrets.aes")))
	if err != nil {
		return err
	}

	return inj.Provide(
		&inject.Object{Value: wlg, Name: "logger.web"},
		&inject.Object{Value: jlg, Name: "logger.jobs"},
		&inject.Object{Value: db},
		&inject.Object{
			Name: "cache.redis",
			Value: &redis.Pool{
				MaxIdle:     3,
				IdleTimeout: 240 * time.Second,
				Dial: func() (redis.Conn, error) {
					c, e := redis.Dial(
						"tcp",
						fmt.Sprintf(
							"%s:%d",
							viper.GetString("cache.host"),
							viper.GetInt("cache.port"),
						),
					)
					if e != nil {
						return nil, e
					}
					if _, e = c.Do("SELECT", viper.GetInt("cache.db")); e != nil {
						c.Close()
						return nil, e
					}
					return c, nil
				},
				TestOnBorrow: func(c redis.Conn, t time.Time) error {
					_, err := c.Do("PING")
					return err
				},
			},
		},
		&inject.Object{Value: cip, Name: "aes.cip"},
		&inject.Object{Value: []byte(viper.GetString("secrets.hmac")), Name: "hmac.key"},
		&inject.Object{Value: sha512.New, Name: "hmac.hash"},
		&inject.Object{Value: []byte(viper.GetString("secrets.jwt")), Name: "jwt.key"},
		&inject.Object{Value: crypto.SigningMethodHS512, Name: "jwt.method"},
	)
}

func init() {
	viper.SetDefault("jobs", map[string]interface{}{
		"type": "redis",
		"host": "localhost",
		"port": 6379,
		"db":   8,
	})

	viper.SetDefault("cache", map[string]interface{}{
		"type": "redis",
		"host": "localhost",
		"port": 6379,
		"db":   8,
	})

	viper.SetDefault("database", map[string]interface{}{
		"driver": "postgres",
		"args": map[string]interface{}{
			"host":    "localhost",
			"port":    5432,
			"user":    "postgres",
			"dbname":  "lotus",
			"sslmode": "disable",
		},
		"pool": map[string]int{
			"max_open": 180,
			"max_idle": 6,
		},
	})

	viper.SetDefault("app", map[string]interface{}{
		"name": "lotus",
	})

	viper.SetDefault("server", map[string]interface{}{
		"port":  8080,
		"name":  "www.change-me.com",
		"theme": "bootstrap4",
	})

	viper.SetDefault("secrets", map[string]interface{}{
		"jwt":    web.RandomStr(32),
		"aes":    web.RandomStr(32),
		"hmac":   web.RandomStr(32),
		"cookie": web.RandomStr(32),
	})

	viper.SetDefault("elasticsearch", []string{"http://localhost:9200"})

	web.Register(&Engine{})
}
