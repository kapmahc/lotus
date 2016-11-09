package auth

import (
	"crypto/x509/pkix"
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"os/user"
	"path"
	"path/filepath"
	"time"

	"bitbucket.org/liamstask/goose/lib/goose"
	"github.com/BurntSushi/toml"
	log "github.com/Sirupsen/logrus"
	"github.com/facebookgo/inject"
	"github.com/fvbock/endless"
	"github.com/gorilla/csrf"
	"github.com/gorilla/mux"
	"github.com/kapmahc/lotus/web"
	negronilogrus "github.com/meatballhat/negroni-logrus"
	"github.com/spf13/viper"
	"github.com/urfave/cli"
	"github.com/urfave/negroni"
	"golang.org/x/text/language"
)

//Shell command line
func (p *Engine) Shell() []cli.Command {
	return []cli.Command{
		{
			Name:    "init",
			Aliases: []string{"i"},
			Usage:   "init config file",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "environment, e",
					Value: "development",
					Usage: "environment, like: development, production, stage, test...",
				},
			},
			Action: func(c *cli.Context) error {
				const fn = "config.toml"
				if _, err := os.Stat(fn); err == nil {
					return fmt.Errorf("file %s already exists", fn)
				}
				fmt.Printf("generate file %s\n", fn)

				viper.Set("env", c.String("environment"))
				args := viper.AllSettings()
				fd, err := os.Create(fn)
				if err != nil {
					return err
				}
				defer fd.Close()
				end := toml.NewEncoder(fd)
				err = end.Encode(args)

				return err
			},
		},
		{
			Name:    "server",
			Aliases: []string{"s"},
			Usage:   "start the app server",
			Action: web.IocAction(func(*cli.Context, *inject.Graph) error {
				// mount
				web.Loop(func(en web.Engine) error {
					en.Mount(p.Router)
					return nil
				})

				// hnd := cors.New(cors.Options{
				// 	AllowCredentials: true,
				// 	AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
				// 	AllowedHeaders:   []string{"*"},
				// 	Debug:            !web.IsProduction(),
				// }).Handler(rt)
				hnd := csrf.Protect([]byte(viper.GetString("secrets.csrf")))(p.Router)
				adr := fmt.Sprintf(":%d", viper.GetInt("server.port"))

				ng := negroni.New()
				ng.Use(negronilogrus.NewMiddleware())
				ng.Use(negroni.HandlerFunc(web.Csrf))
				ng.Use(negroni.HandlerFunc(p.I18n.Handler))
				ng.Use(negroni.NewStatic(http.Dir(
					path.Join("themes", viper.GetString("server.theme"), "assets"),
				)))
				ng.UseHandler(hnd)

				if web.IsProduction() {
					return endless.ListenAndServe(
						adr,
						ng,
					)
				}

				return http.ListenAndServe(adr, ng)
			}),
		},
		{
			Name:    "routes",
			Aliases: []string{"rt"},
			Usage:   "list all routes",
			Action: web.CfgAction(func(*cli.Context) error {
				rt := mux.NewRouter()

				web.Loop(func(en web.Engine) error {
					en.Mount(rt)
					return nil
				})

				fmt.Println("NAME\tPATH")
				return rt.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
					pt, er := route.GetPathTemplate()
					if er != nil {
						return er
					}
					fmt.Printf("%s\t%s\n", route.GetName(), pt)
					return nil
				})
			}),
		},
		{
			Name:    "worker",
			Aliases: []string{"w"},
			Usage:   "start the worker progress",
			Action: web.IocAction(func(_ *cli.Context, inj *inject.Graph) error {
				hn, err := os.Hostname()
				if err != nil {
					return err
				}
				un, err := user.Current()
				if err != nil {
					return err
				}
				worker := p.Server.NewWorker(fmt.Sprintf("%s@%s", un.Name, hn))
				return worker.Launch()
			}),
		},
		{
			Name:    "openssl",
			Aliases: []string{"ssl"},
			Usage:   "generate ssl certificates",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "name, n",
					Usage: "name",
				},
				cli.StringFlag{
					Name:  "country, c",
					Value: "Earth",
					Usage: "country",
				},
				cli.StringFlag{
					Name:  "organization, o",
					Value: "Mother Nature",
					Usage: "organization",
				},
				cli.IntFlag{
					Name:  "years, y",
					Value: 1,
					Usage: "years",
				},
			},
			Action: web.CfgAction(func(c *cli.Context) error {
				name := c.String("name")
				if len(name) == 0 {
					cli.ShowCommandHelp(c, "openssl")
					return nil
				}
				root := path.Join("etc", "ssl", name)

				key, crt, err := CreateCertificate(
					true,
					pkix.Name{
						Country:      []string{c.String("country")},
						Organization: []string{c.String("organization")},
					},
					c.Int("years"),
				)
				if err != nil {
					return err
				}

				fnk := path.Join(root, "key.pem")
				fnc := path.Join(root, "crt.pem")

				fmt.Printf("generate pem file %s\n", fnk)
				err = WritePemFile(fnk, "RSA PRIVATE KEY", key)
				fmt.Printf("test: openssl rsa -noout -text -in %s\n", fnk)

				if err == nil {
					fmt.Printf("generate pem file %s\n", fnc)
					err = WritePemFile(fnc, "CERTIFICATE", crt)
					fmt.Printf("test: openssl x509 -noout -text -in %s\n", fnc)
				}
				if err == nil {
					fmt.Printf(
						"verify: diff <(openssl rsa -noout -modulus -in %s) <(openssl x509 -noout -modulus -in %s)",
						fnk,
						fnc,
					)
				}
				fmt.Println()
				return err
			}),
		},

		{
			Name:    "nginx",
			Aliases: []string{"ng"},
			Usage:   "init nginx config file",
			Action: web.CfgAction(func(*cli.Context) error {
				const tpl = `
server {
  listen 80;
  server_name {{.Name}};
  rewrite ^(.*) https://$host$1 permanent;
}

upstream {{.Name}}_prod {
  server localhost:{{.Port}} fail_timeout=0;
}

server {
  listen 443;

  ssl  on;
  ssl_certificate  /etc/ssl/certs/{{.Name}}.crt;
  ssl_certificate_key  /etc/ssl/private/{{.Name}}.key;
  ssl_session_timeout  5m;
  ssl_protocols  SSLv2 SSLv3 TLSv1;
  ssl_ciphers  RC4:HIGH:!aNULL:!MD5;
  ssl_prefer_server_ciphers  on;

  client_max_body_size 4G;
  keepalive_timeout 10;
  proxy_buffers 16 64k;
  proxy_buffer_size 128k;

  server_name {{.Name}};
  root {{.Root}}/public;
  index index.html;
  access_log /var/log/nginx/{{.Name}}.access.log;
  error_log /var/log/nginx/{{.Name}}.error.log;
  location / {
    try_files $uri $uri/ /index.html?/$request_uri;
  }
#  location ^~ /assets/ {
#    gzip_static on;
#    expires max;
#    access_log off;
#    add_header Cache-Control "public";
#  }
  location ~* \.(?:css|js)$ {
    gzip_static on;
    expires max;
    access_log off;
    add_header Cache-Control "public";
  }
  location ~* \.(?:jpg|jpeg|gif|png|ico|cur|gz|svg|svgz|mp4|ogg|ogv|webm|htc)$ {
    expires 1M;
    access_log off;
    add_header Cache-Control "public";
  }
  location ~* \.(?:rss|atom)$ {
    expires 12h;
    access_log off;
    add_header Cache-Control "public";
  }
  location ~ ^/api/{{.Version}}(/?)(.*) {
    proxy_set_header X-Forwarded-Proto https;
    proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
    proxy_set_header Host $http_host;
    proxy_set_header X-Real-IP $remote_addr;
    proxy_redirect off;
    proxy_pass http://{{.Name}}_prod/$2$is_args$args;
    # limit_req zone=one;
  }
}
`
				t, err := template.New("").Parse(tpl)
				if err != nil {
					return err
				}
				pwd, err := os.Getwd()
				if err != nil {
					return err
				}

				name := viper.GetString("server.name")
				fn := path.Join("etc", "nginx", "sites-enabled", name+".conf")
				if err = os.MkdirAll(path.Dir(fn), 0700); err != nil {
					return err
				}
				fmt.Printf("generate file %s\n", fn)
				fd, err := os.OpenFile(fn, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0600)
				if err != nil {
					return err
				}
				defer fd.Close()

				return t.Execute(fd, struct {
					Name    string
					Port    int
					Root    string
					Version string
				}{
					Name:    name,
					Port:    viper.GetInt("http.port"),
					Root:    pwd,
					Version: "v1",
				})
			}),
		},
		{
			Name:    "database",
			Aliases: []string{"db"},
			Usage:   "database operations",
			Subcommands: []cli.Command{
				{
					Name:    "example",
					Usage:   "scripts example for create database and user",
					Aliases: []string{"e"},
					Action: web.CfgAction(func(*cli.Context) error {
						drv := viper.GetString("database.driver")
						args := viper.GetStringMapString("database.args")
						var err error
						switch drv {
						case "postgres":
							fmt.Printf("CREATE USER %s WITH PASSWORD '%s';\n", args["user"], args["password"])
							fmt.Printf("CREATE DATABASE %s WITH ENCODING='UTF8';\n", args["dbname"])
							fmt.Printf("GRANT ALL PRIVILEGES ON DATABASE %s TO %s;\n", args["dbname"], args["user"])
						default:
							err = fmt.Errorf("unknown driver %s", drv)
						}
						return err
					}),
				},
				{
					Name:    "generate",
					Usage:   "generate database migration",
					Aliases: []string{"g"},
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "name, n",
							Usage: "name",
						},
					},
					Action: web.CfgAction(func(c *cli.Context) error {
						name := c.String("name")
						if name == "" {
							cli.ShowCommandHelp(c, "generate")
							return nil
						}
						conf := migrationConf()
						file, err := goose.CreateMigration(name, "sql", conf.MigrationsDir, time.Now())
						if err == nil {
							log.Println("generate file ", file)
						}
						return err
					}),
				},
				{
					Name:    "migrate",
					Usage:   "migrate the database",
					Aliases: []string{"m"},
					Action: web.CfgAction(func(*cli.Context) error {
						conf := migrationConf()
						ver, err := goose.GetMostRecentDBVersion(conf.MigrationsDir)
						if err != nil {
							return err
						}
						return goose.RunMigrations(conf, conf.MigrationsDir, ver)
					}),
				},
				{
					Name:    "rollback",
					Usage:   "rollback the database",
					Aliases: []string{"r"},
					Action: web.CfgAction(func(*cli.Context) error {
						conf := migrationConf()
						cur, err := goose.GetDBVersion(conf)
						if err != nil {
							return err
						}
						ver, err := goose.GetPreviousDBVersion(conf.MigrationsDir, cur)
						if err != nil {
							return err
						}
						return goose.RunMigrations(conf, conf.MigrationsDir, ver)
					}),
				},
				{
					Name:    "status",
					Usage:   "dump the migration status for the current DB",
					Aliases: []string{"st"},
					Action: web.CfgAction(func(*cli.Context) error {
						conf := migrationConf()
						min := int64(0)
						max := int64((1 << 63) - 1)
						migrations, err := goose.CollectMigrations(conf.MigrationsDir, min, max)
						if err != nil {
							return err
						}
						db, err := goose.OpenDBFromDBConf(conf)
						if err != nil {
							return err
						}
						defer db.Close()

						// must ensure that the version table exists if we're running on a pristine DB
						if _, err := goose.EnsureDBVersion(conf, db); err != nil {
							return err
						}

						fmt.Printf("goose: status for environment '%v'\n", conf.Env)
						fmt.Println("    Applied At                  Migration")
						for _, m := range migrations {
							var row goose.MigrationRecord
							q := fmt.Sprintf("SELECT tstamp, is_applied FROM goose_db_version WHERE version_id=%d ORDER BY tstamp DESC LIMIT 1", m.Version)
							e := db.QueryRow(q).Scan(&row.TStamp, &row.IsApplied)

							if e != nil && e != sql.ErrNoRows {
								return e
							}

							var appliedAt string
							if row.IsApplied {
								appliedAt = row.TStamp.Format(time.ANSIC)
							} else {
								appliedAt = "Pending"
							}
							fmt.Printf("    %-24s -- %v\n", appliedAt, filepath.Base(m.Source))
						}
						return nil
					}),
				},
				{
					Name:    "connect",
					Usage:   "connect database",
					Aliases: []string{"c"},
					Action: web.CfgAction(func(*cli.Context) error {
						drv := viper.GetString("database.driver")
						args := viper.GetStringMapString("database.args")
						var err error
						switch drv {
						case "postgres":
							err = web.Shell("psql",
								"-h", args["host"],
								"-p", args["port"],
								"-U", args["user"],
								args["dbname"],
							)
						default:
							err = fmt.Errorf("unknown driver %s", drv)
						}
						return err
					}),
				},
				{
					Name:    "create",
					Usage:   "create database",
					Aliases: []string{"n"},
					Action: web.CfgAction(func(*cli.Context) error {
						drv := viper.GetString("database.driver")
						args := viper.GetStringMapString("database.args")
						var err error
						switch drv {
						case "postgres":
							err = web.Shell("psql",
								"-h", args["host"],
								"-p", args["port"],
								"-U", args["user"],
								"-c", fmt.Sprintf("create database %s WITH ENCODING='UTF8'", args["dbname"]),
							)
						default:
							err = fmt.Errorf("unknown driver %s", drv)
						}
						return err
					}),
				},
				{
					Name:    "drop",
					Usage:   "drop database",
					Aliases: []string{"d"},
					Action: web.CfgAction(func(*cli.Context) error {
						drv := viper.GetString("database.driver")
						args := viper.GetStringMapString("database.args")
						var err error
						switch drv {
						case "postgres":
							err = web.Shell("psql",
								"-h", args["host"],
								"-p", args["port"],
								"-U", args["user"],
								"-c", fmt.Sprintf("drop database %s", args["dbname"]),
							)
						default:
							err = fmt.Errorf("unknown driver %s", drv)
						}
						return err
					}),
				},
			},
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
					Action: web.IocAction(func(*cli.Context, *inject.Graph) error {
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
					Action: web.IocAction(func(*cli.Context, *inject.Graph) error {
						return p.Cache.Flush()
					}),
				},
			},
		},
		{
			Name:    "locales",
			Aliases: []string{"l"},
			Usage:   "locales operations",
			Subcommands: []cli.Command{
				{
					Name:    "list",
					Usage:   "list all locales codes",
					Aliases: []string{"l"},
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "lang, l",
							Value: language.SimplifiedChinese.String(),
							Usage: "language",
						},
					},
					Action: web.IocAction(func(c *cli.Context, _ *inject.Graph) error {
						lang := c.String("lang")
						tag, err := language.Parse(lang)
						if err != nil {
							return err
						}
						keys, err := p.I18n.Codes(tag.String())
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
					Name:    "import",
					Usage:   "import locale items",
					Aliases: []string{"i"},
					Action: web.IocAction(func(c *cli.Context, _ *inject.Graph) error {
						return p.I18n.Load("locales")
					}),
				},
			},
		},
	}
}
