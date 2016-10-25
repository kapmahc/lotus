package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/astaxie/beego/session/redis"
	_ "github.com/kapmahc/lotus/routers"
	_ "github.com/lib/pq"
)

func main() {
	orm.Debug = beego.AppConfig.String("runmode") != "prod"
	orm.RegisterDriver("postgres", orm.DRPostgres)
	orm.RegisterDataBase(
		"default",
		"postgres",
		beego.AppConfig.String("databaseurl"),
		beego.AppConfig.DefaultInt("databasemaxidle", 6),
		beego.AppConfig.DefaultInt("databasemaxconn", 180),
	)
	beego.Run()
}
