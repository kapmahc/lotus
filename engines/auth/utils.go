package auth

import (
	"fmt"
	"log/syslog"
	"os"
	"path/filepath"
	"time"

	"bitbucket.org/liamstask/goose/lib/goose"
	machinery "github.com/RichardKnop/machinery/v1"
	"github.com/RichardKnop/machinery/v1/config"
	"github.com/RichardKnop/machinery/v1/logger"
	"github.com/garyburd/redigo/redis"
	"github.com/jinzhu/gorm"
	"github.com/kapmahc/lotus/web"
	"github.com/spf13/viper"
)

func migrationConf() *goose.DBConf {
	root := filepath.Join("db", viper.GetString("database.driver"), "migrations")
	os.MkdirAll(root, 0700)
	var drv goose.DBDriver
	drv.Name, drv.OpenStr = databaseURL()
	switch drv.Name {
	case "postgres":
		drv.Import = "github.com/lib/pq"
		drv.Dialect = &goose.PostgresDialect{}
	case "mysql":
		drv.Import = "github.com/go-sql-driver/mysql"
		drv.Dialect = &goose.MySqlDialect{}
	case "sqlite3":
		drv.Import = "github.com/mattn/go-sqlite3"
		drv.Dialect = &goose.Sqlite3Dialect{}
	}
	return &goose.DBConf{
		MigrationsDir: root,
		Env:           viper.GetString("env"),
		Driver:        drv,
	}
}

func databaseURL() (string, string) {
	//postgresql: "user=%s password=%s host=%s port=%d dbname=%s sslmode=%s"
	args := ""
	for k, v := range viper.GetStringMapString("database.args") {
		args += fmt.Sprintf(" %s=%s ", k, v)
	}
	return viper.GetString("database.driver"), args
}

func openCacheRedis() *redis.Pool {
	return &redis.Pool{
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
	}
}

func openDatabase() (*gorm.DB, error) {
	drv, url := databaseURL()
	db, err := gorm.Open(drv, url)
	if err != nil {
		return nil, err
	}
	if !web.IsProduction() {
		db.LogMode(true)
	}

	if err := db.DB().Ping(); err != nil {
		return nil, err
	}

	db.DB().SetMaxIdleConns(viper.GetInt("database.pool.max_idle"))
	db.DB().SetMaxOpenConns(viper.GetInt("database.pool.max_open"))
	return db, nil
}

func openLogger(tag string) (*syslog.Writer, error) {
	priority := syslog.LOG_DEBUG
	if web.IsProduction() {
		priority = syslog.LOG_INFO
	}
	return syslog.New(priority, fmt.Sprintf("%s-%s", viper.GetString("app.name"), tag))
}

func openJobServer() (*machinery.Server, error) {
	lg, err := openLogger("jobs")
	if err != nil {
		return nil, err
	}
	logger.Set(&web.JobLogger{Writer: lg})

	url := fmt.Sprintf(
		"redis://%s:%d/%d",
		viper.GetString("jobs.host"),
		viper.GetInt("jobs.port"),
		viper.GetInt("jobs.db"),
	)
	var cnf = config.Config{
		Broker:          url,
		ResultBackend:   url,
		ResultsExpireIn: 60 * 60 * 24 * 7 * 10,
		DefaultQueue:    viper.GetString("app.name") + "-tasks",
	}

	return machinery.NewServer(&cnf)
}
