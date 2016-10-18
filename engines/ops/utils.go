package ops

import (
	"errors"
	"path"
	"strings"

	"github.com/spf13/viper"
)

func dbMigrationsDir() string {
	return path.Join("db", viper.GetString("database.driver"), "migrations")
}

func errors2(ers []error) error {
	var msg []string
	for _, er := range ers {
		msg = append(msg, er.Error())
	}
	return errors.New(strings.Join(msg, "\n"))
}
