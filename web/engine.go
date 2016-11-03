package web

import (
	"github.com/facebookgo/inject"
	"github.com/gin-gonic/gin"
	"github.com/urfave/cli"
)

//Engine web engine
type Engine interface {
	//Init set default config
	Init()
	//Map ioc objects
	Map(*inject.Graph) error
	//Mount web points
	Mount(*gin.Engine)
	//Dashboard dashboard's nav-bar
	Dashboard() []Dropdown
	//Shell command line
	Shell() []cli.Command
}

var engines []Engine

//Register register engines
func Register(en ...Engine) {
	engines = append(engines, en...)
}

//Loop loop engines
func Loop(fn func(en Engine) error) error {
	for _, en := range engines {
		if err := fn(en); err != nil {
			return err
		}
	}
	return nil
}
