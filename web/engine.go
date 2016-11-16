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
	//Shell command line
	Shell() []cli.Command
	//Insert seed data
	Seed() error
	//Worker register worker
	Worker()
}

var engines []Engine

//Register register engines
func Register(en ...Engine) {
	// for _, e := range en {
	// 	viper.SetDefault(
	// 		fmt.Sprintf("engines.%s", reflect.TypeOf(e).Elem().PkgPath()),
	// 		true,
	// 	)
	// }
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
