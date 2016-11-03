package web

import (
	"github.com/facebookgo/inject"
	"github.com/spf13/viper"
	"github.com/urfave/cli"
)

//IocAction ioc action
func IocAction(fn func(*cli.Context, *inject.Graph) error) func(c *cli.Context) error {
	return CfgAction(func(ctx *cli.Context) error {
		var inj inject.Graph
		for _, en := range engines {
			if err := en.Init(&inj); err != nil {
				return err
			}
			if err := inj.Provide(&inject.Object{Value: en}); err != nil {
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
func CfgAction(f cli.ActionFunc) func(c *cli.Context) error {
	return func(c *cli.Context) error {
		viper.SetEnvPrefix("lotus")
		viper.BindEnv("env")
		viper.SetDefault("env", "development")

		viper.SetConfigName("config")
		viper.SetConfigType("toml")
		viper.AddConfigPath(".")

		if err := viper.ReadInConfig(); err != nil {
			return err
		}
		return f(c)
	}
}
