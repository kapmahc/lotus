package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/astaxie/beego/session/redis"
	"github.com/benmanns/goworker"
	_ "github.com/kapmahc/lotus/routers"
	_ "github.com/lib/pq"
)

func main() {
	// orm
	orm.Debug = beego.AppConfig.String("runmode") != "prod"
	orm.RegisterDriver("postgres", orm.DRPostgres)
	orm.RegisterDataBase(
		"default",
		"postgres",
		beego.AppConfig.String("databaseurl"),
		beego.AppConfig.DefaultInt("databasemaxidle", 6),
		beego.AppConfig.DefaultInt("databasemaxconn", 180),
	)

	// workers
	goworker.SetSettings(goworker.WorkerSettings{
		URI:         beego.AppConfig.String("workerredis"),
		Connections: 10,
		UseNumber:   true,
		Queues:      beego.AppConfig.Strings("workerqueues"),
		Concurrency: beego.AppConfig.DefaultInt("workerconcurrency", 2),
		Namespace:   "resque:",
		Interval:    5.0,
	})

	go beego.Run()

	if err := goworker.Work(); err != nil {
		beego.Error(err)
	}
}
