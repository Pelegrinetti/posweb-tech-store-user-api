package config

import (
	"github.com/Pelegrinetti/posweb-user-api/pkg/server"
	"github.com/spf13/viper"
)

type Config struct {
	ServerConfig server.ServerConfig `mapstructure:",squash"`
}

func New() Config {
	viper.SetDefault("PORT", 3001)
	viper.AutomaticEnv()

	config := Config{}

	err := viper.Unmarshal(&config)

	if err != nil {
		panic(err)
	}

	return config
}
