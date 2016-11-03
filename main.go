package main

import (
	"github.com/golang/glog"
	_ "github.com/kapmahc/lotus/engines/auth"
	_ "github.com/kapmahc/lotus/engines/forum"
	_ "github.com/kapmahc/lotus/engines/ops/mail"
	_ "github.com/kapmahc/lotus/engines/ops/vpn"
	_ "github.com/kapmahc/lotus/engines/reading"
	_ "github.com/kapmahc/lotus/engines/shop"
	"github.com/kapmahc/lotus/web"
	_ "github.com/lib/pq"
)

var version string

func main() {
	if err := web.Main(version); err != nil {
		glog.Fatal(err)
	}
}
