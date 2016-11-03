package main

import (
	"flag"
	"log"

	_ "github.com/kapmahc/lotus/engines/auth"
	_ "github.com/kapmahc/lotus/engines/forum"
	_ "github.com/kapmahc/lotus/engines/ops/mail"
	_ "github.com/kapmahc/lotus/engines/ops/vpn"
	_ "github.com/kapmahc/lotus/engines/reading"
	_ "github.com/kapmahc/lotus/engines/shop"
	"github.com/kapmahc/lotus/web"
)

var version string

func main() {
	flag.Parse()
	if err := web.Main(version); err != nil {
		log.Fatal(err)
	}
}
