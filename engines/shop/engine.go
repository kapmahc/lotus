package shop

import (
	"html/template"

	"github.com/facebookgo/inject"
	"github.com/gin-gonic/gin"
	"github.com/kapmahc/lotus/web"
	"github.com/urfave/cli"
)

//Engine shop engine
type Engine struct {
}

//Funcs html template funcs
func (p *Engine) Funcs() template.FuncMap {
	return template.FuncMap{}
}

//Home home
func (p *Engine) Home() gin.HandlerFunc {
	return func(c *gin.Context) {}
}

//Init init ioc objects
func (p *Engine) Init(*inject.Graph) error {
	return nil
}

//Mount web points
func (p *Engine) Mount(*gin.Engine) {

}

//Dashboard dashboard's nav-bar
func (p *Engine) Dashboard() []web.Dropdown {
	return []web.Dropdown{}
}

//Shell command line
func (p *Engine) Shell() []cli.Command {
	return []cli.Command{}
}

//Worker register worker
func (p *Engine) Worker() {

}

func init() {
	//	web.Register(&Engine{})
}
