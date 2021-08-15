package config

import (
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	*viper.Viper
}

type Interface interface {
	Postgres
}

func New() Interface {
	v := viper.NewWithOptions(viper.EnvKeyReplacer(strings.NewReplacer(".", "_")))

	v.AutomaticEnv()
	v.SetEnvPrefix("hackernews")

	defaults := map[string]interface{}{
		"postgres": map[string]interface{}{
			"username": "postgres",
			"password": "",
			"address":  "localhost:5432",
			"database": "postgres",
		},
	}

	for key, value := range defaults {
		v.SetDefault(key, value)
	}

	return Config{v}
}
