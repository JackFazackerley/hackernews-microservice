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
	GRPC
}

func New() Interface {
	v := viper.NewWithOptions(viper.EnvKeyReplacer(strings.NewReplacer(".", "_")))

	v.AutomaticEnv()
	v.SetEnvPrefix("hackernews")

	defaults := map[string]interface{}{
		"postgres": map[string]interface{}{
			"username": "postgres",
			"password": "",
			"address":  "127.0.0.1:5432",
			"database": "hackernews",
		},
		"grpc": map[string]interface{}{
			"address": ":50051",
		},
	}

	for key, value := range defaults {
		v.SetDefault(key, value)
	}

	return Config{v}
}
