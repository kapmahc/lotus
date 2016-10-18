package main

import (
	"log"

	_ "github.com/kapmahc/lotus/engines/auth"
	_ "github.com/kapmahc/lotus/engines/forum"
	_ "github.com/kapmahc/lotus/engines/ops"
	_ "github.com/kapmahc/lotus/engines/reading"
	_ "github.com/kapmahc/lotus/engines/shop"
	"github.com/kapmahc/lotus/web"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/mattes/migrate/driver/postgres"
)

var version string

func main() {
	if err := web.Run(version); err != nil {
		log.Fatal(err)
	}
}
