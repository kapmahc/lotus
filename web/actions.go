package web

import (
	"github.com/facebookgo/inject"
	"github.com/spf13/viper"
	"github.com/urfave/cli"
)

//IocAction ioc action
func IocAction(fn func(*cli.Context, *inject.Graph) error) cli.ActionFunc {
	return CfgAction(func(ctx *cli.Context) error {
		var inj inject.Graph
		for _, en := range engines {
			if err := en.Map(&inj); err != nil {
				return err
			}
		}
		if err := inj.Populate(); err != nil {
			return err
		}
		return fn(ctx, &inj)
	})
}

//CfgAction config action
func CfgAction(f cli.ActionFunc) cli.ActionFunc {
	return func(c *cli.Context) error {
		viper.SetEnvPrefix("lotus")
		viper.BindEnv("env")
		viper.SetDefault("env", "development")

		viper.SetConfigName("config")
		viper.SetConfigType("toml")
		viper.AddConfigPath(".")

		for _, en := range engines {
			en.Init()
		}

		if err := viper.ReadInConfig(); err != nil {
			return err
		}
		return f(c)
	}
}
