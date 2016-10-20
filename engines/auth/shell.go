package auth

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"

	"github.com/facebookgo/inject"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"github.com/kapmahc/lotus/web"
	"github.com/kapmahc/lotus/web/i18n"
	"github.com/rs/cors"
	"github.com/spf13/viper"
	"github.com/urfave/cli"
)

//Shell command line
func (p *Engine) Shell() []cli.Command {
	return []cli.Command{

		{
			Name:    "server",
			Aliases: []string{"s"},
			Usage:   "start the app server",
			Action: IocAction(func(*cli.Context, *inject.Graph) error {
				if IsProduction() {
					gin.SetMode(gin.ReleaseMode)
				}
				rt := gin.Default()

				theme := viper.GetString("server.theme")

				tpl, err := template.
					New("").
					Funcs(template.FuncMap{
						"T": p.I18n.T,
					}).
					ParseGlob(
						fmt.Sprintf("templates/%s/**/*", theme),
					)
				if err != nil {
					return err
				}
				rt.SetHTMLTemplate(tpl)

				rt.Use(i18n.LocaleHandler(p.Logger))

				web.Loop(func(en web.Engine) error {
					en.Mount(rt)
					return nil
				})

				adr := fmt.Sprintf(":%d", viper.GetInt("server.port"))

				hnd := cors.New(cors.Options{
					AllowCredentials: true,
					AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
					AllowedHeaders:   []string{"*"},
					Debug:            !IsProduction(),
				}).Handler(rt)

				if IsProduction() {
					return endless.ListenAndServe(adr, hnd)
				}
				return http.ListenAndServe(adr, hnd)
			}),
		},

		{
			Name:    "worker",
			Aliases: []string{"w"},
			Usage:   "start the worker progress",
			Action: IocAction(func(_ *cli.Context, inj *inject.Graph) error {
				web.Loop(func(en web.Engine) error {
					en.Worker()
					return nil
				})

				return p.Jobber.Start()
			}),
		},

		{
			Name:    "redis",
			Aliases: []string{"r"},
			Usage:   "open redis connection",
			Action: Action(func(*cli.Context) error {
				return web.Shell(
					"redis-cli",
					"-h", viper.GetString("redis.host"),
					"-p", viper.GetString("redis.port"),
					"-n", viper.GetString("redis.db"),
				)
			}),
		},

		{
			Name:    "cache",
			Aliases: []string{"c"},
			Usage:   "cache operations",
			Subcommands: []cli.Command{
				{
					Name:    "list",
					Usage:   "list all cache keys",
					Aliases: []string{"l"},
					Action: IocAction(func(*cli.Context, *inject.Graph) error {
						keys, err := p.Cache.Keys()
						if err != nil {
							return err
						}
						for _, k := range keys {
							fmt.Println(k)
						}
						return nil
					}),
				},
				{
					Name:    "clear",
					Usage:   "clear cache items",
					Aliases: []string{"c"},
					Action: IocAction(func(*cli.Context, *inject.Graph) error {
						return p.Cache.Flush()
					}),
				},
			},
		},

		{
			Name:    "users",
			Aliases: []string{"u"},
			Usage:   "users operations",
			Subcommands: []cli.Command{
				{
					Name:    "list",
					Usage:   "list all users",
					Aliases: []string{"l"},
					Action: IocAction(func(*cli.Context, *inject.Graph) error {
						var users []User
						if err := p.Db.Select([]string{"uid", "email", "name"}).Find(&users).Error; err != nil {
							return err
						}
						fmt.Println("UID                                     INFO")
						for _, u := range users {
							fmt.Printf("%s\t%s<%s>\n", u.UID, u.Name, u.Email)
						}
						return nil
					}),
				},
				{
					Name:    "role",
					Usage:   "allow/deny role to user",
					Aliases: []string{"r"},
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "uid,u",
							Usage: "user's uid.",
						},
						cli.StringFlag{
							Name:  "name,n",
							Usage: "role's name.",
						},
						cli.BoolFlag{
							Name:  "deny,d",
							Usage: "deny?",
						},
					},
					Action: IocAction(func(c *cli.Context, _ *inject.Graph) error {
						uid := c.String("uid")
						if uid == "" {
							return errors.New("user's uid mustn't empty")
						}
						name := c.String("name")
						if name == "" {
							return errors.New("role's name mustn't empty")
						}
						deny := c.Bool("deny")
						user, err := p.Dao.GetUserByUID(uid)
						if err != nil {
							return err
						}
						role, err := p.Dao.Role(name, "-", 0)
						if err != nil {
							return err
						}
						if deny {
							err = p.Dao.Deny(role.ID, user.ID)
						} else {
							err = p.Dao.Allow(role.ID, user.ID, 20, 0, 0)
						}
						return err
					}),
				},
			},
		},
	}
}

func init() {
	viper.SetEnvPrefix("lotus")
	viper.BindEnv("env")
	viper.SetDefault("env", "development")

	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")

	viper.SetDefault("redis", map[string]interface{}{
		"host": "localhost",
		"port": 6379,
		"db":   8,
	})

	viper.SetDefault("home", map[string]interface{}{
		"backend":  "http://localhost:8080",
		"frontend": "http://localhost:4200",
	})

	viper.SetDefault("database", map[string]interface{}{
		"driver": "postgres",
		"args": map[string]interface{}{
			"host":    "localhost",
			"port":    5432,
			"user":    "postgres",
			"dbname":  "lotus_dev",
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
		"jwt":     web.RandomStr(32),
		"aes":     web.RandomStr(32),
		"hmac":    web.RandomStr(32),
		"session": web.RandomStr(32),
	})

	viper.SetDefault("workers", map[string]interface{}{
		"timeout": 30,
	})

	viper.SetDefault("elasticsearch", []string{"http://localhost:9200"})
}
