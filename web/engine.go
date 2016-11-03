package web

import (
	"github.com/facebookgo/inject"
	"github.com/gin-gonic/gin"
	"github.com/urfave/cli"
)

//Engine web engine
type Engine interface {
	//Home home
	Home() gin.HandlerFunc
	//Init init ioc objects
	Init(*inject.Graph) error
	//Mount web points
	Mount(*gin.Engine)
	//Dashboard dashboard's nav-bar
	Dashboard() []Dropdown
	//Shell command line
	Shell() []cli.Command
	//Worker register worker
	Worker()
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