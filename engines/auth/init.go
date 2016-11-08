package auth

import (
	"crypto/aes"
	"crypto/sha512"
	"html/template"
	"path"

	"github.com/SermoDigital/jose/crypto"
	"github.com/facebookgo/inject"
	"github.com/kapmahc/lotus/web"
	"github.com/spf13/viper"
	"github.com/unrolled/render"
)

//Init init ioc objects
func (p *Engine) Init(inj *inject.Graph) error {

	logger, err := web.OpenLogger(viper.GetString("app.name") + "-web")
	if err != nil {
		return err
	}

	db, err := openDatabase()
	if err != nil {
		return err
	}

	cip, err := aes.NewCipher([]byte(viper.GetString("secrets.aes")))
	if err != nil {
		return err
	}

	srv, err := openJobServer()
	if err != nil {
		return err
	}

	var fms []template.FuncMap
	web.Loop(func(en web.Engine) error {
		fms = append(fms, en.FuncMap())
		return nil
	})

	return inj.Provide(
		&inject.Object{Value: logger},
		&inject.Object{Value: db},
		&inject.Object{Value: srv},
		&inject.Object{
			Name:  "cache.redis",
			Value: openCacheRedis(),
		},
		&inject.Object{Value: cip, Name: "aes.cip"},
		&inject.Object{Value: []byte(viper.GetString("secrets.hmac")), Name: "hmac.key"},
		&inject.Object{Value: sha512.New, Name: "hmac.hash"},
		&inject.Object{Value: []byte(viper.GetString("secrets.jwt")), Name: "jwt.key"},
		&inject.Object{Value: crypto.SigningMethodHS512, Name: "jwt.method"},
		&inject.Object{Value: render.New(render.Options{
			Directory:     path.Join("themes", viper.GetString("server.theme"), "views"),
			Funcs:         fms,
			Layout:        "layout",
			Extensions:    []string{".html"},
			IsDevelopment: !web.IsProduction(),
		}),
		},
	)
}

func init() {
	viper.SetDefault("jobs", map[string]interface{}{
		"type": "redis",
		"host": "localhost",
		"port": 6379,
		"db":   6,
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
		"csrf":   web.RandomStr(32),
	})

	viper.SetDefault("elasticsearch", []string{"http://localhost:9200"})

	web.Register(&Engine{})
}
