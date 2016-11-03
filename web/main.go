package web

import (
	"os"

	"github.com/spf13/viper"
	"github.com/urfave/cli"
)

//Main entry
func Main(version string) error {

	app := cli.NewApp()
	app.Name = os.Args[0]
	app.Version = version
	app.Usage = "LOTUS web application(by go-lang)."
	app.EnableBashCompletion = true
	app.Commands = []cli.Command{}

	for _, en := range engines {
		cmd := en.Shell()
		app.Commands = append(app.Commands, cmd...)
	}

	return app.Run(os.Args)
}

func init() {
	viper.SetDefault("env", "development")
}
