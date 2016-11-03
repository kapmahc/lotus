package auth

import (
	"crypto/aes"
	"crypto/sha512"

	"github.com/SermoDigital/jose/crypto"
	"github.com/facebookgo/inject"
	"github.com/gin-gonic/gin"
	"github.com/kapmahc/lotus/web"
	"github.com/spf13/viper"
)

//Engine auth engine
type Engine struct {
	I18n  *web.I18n  `inject:""`
	Cache *web.Cache `inject:""`
}

//Home home
func (p *Engine) Home() gin.HandlerFunc {
	return func(c *gin.Context) {}
}

//Init init ioc objects
func (p *Engine) Init(inj *inject.Graph) error {
	db, err := OpenDatabase()
	if err != nil {
		return err
	}

	rep := OpenRedis()

	cip, err := aes.NewCipher([]byte(viper.GetString("secrets.aes")))
	if err != nil {
		return err
	}

	return inj.Provide(
		&inject.Object{Value: db},
		&inject.Object{Value: rep},
		&inject.Object{Value: cip, Name: "aes.cip"},
		&inject.Object{Value: []byte(viper.GetString("secrets.hmac")), Name: "hmac.key"},
		&inject.Object{Value: sha512.New, Name: "hmac.hash"},
		&inject.Object{Value: []byte(viper.GetString("secrets.jwt")), Name: "jwt.key"},
		&inject.Object{Value: crypto.SigningMethodHS512, Name: "jwt.method"},
	)
}

//Mount web points
func (p *Engine) Mount(*gin.Engine) {

}

//Dashboard dashboard's nav-bar
func (p *Engine) Dashboard() []web.Dropdown {
	return []web.Dropdown{}
}

func init() {

	viper.SetDefault("redis", map[string]interface{}{
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
