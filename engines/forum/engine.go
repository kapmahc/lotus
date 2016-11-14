package forum

import (
	"github.com/facebookgo/inject"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/kapmahc/lotus/engines/auth"
	"github.com/kapmahc/lotus/web"
	"github.com/urfave/cli"
)

//Engine shop engine
type Engine struct {
	I18n  *web.I18n  `inject:""`
	Cache *web.Cache `inject:""`
	Db    *gorm.DB   `inject:""`
	Jwt   *auth.Jwt  `inject:""`
	Dao   *auth.Dao  `inject:""`
}

//Home home
func (p *Engine) Home() gin.HandlerFunc {
	return func(c *gin.Context) {}
}

//Init init ioc objects
func (p *Engine) Init(*inject.Graph) error {
	return nil
}

//Shell command line
func (p *Engine) Shell() []cli.Command {
	return []cli.Command{}
}

//Worker register worker
func (p *Engine) Worker() {

}

func init() {
	web.Register(&Engine{})
}
