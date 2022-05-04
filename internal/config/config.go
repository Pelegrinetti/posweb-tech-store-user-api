package config

import (
	"github.com/spf13/viper"
)

type ServerConfig struct {
	Port int `mapstructure:"PORT"`
}

type DatabaseConfig struct {
	Uri    string `mapstructure:"MONGODB_URI"`
	DBName string `mapstructure:"MONGODB_NAME"`
}

type AuthConfig struct {
	GoogleClientId string `mapstructure:"GOOGLE_CLIENT_ID"`
}

type Config struct {
	ServerConfig   ServerConfig     `mapstructure:",squash"`
	DatabaseConfig DatabaseConfig `mapstructure:",squash"`
	AuthConfig     AuthConfig              `mapstructure:",squash"`
}

func New() *Config {
	viper.SetDefault("PORT", 3001)
	viper.SetDefault("MONGODB_URI", "mongodb://localhost:27017/user-api")
	viper.SetDefault("MONGODB_NAME", "user-api")
	viper.SetDefault("GOOGLE_CLIENT_ID", "")
	viper.AutomaticEnv()

	config := Config{}

	err := viper.Unmarshal(&config)

	if err != nil {
		panic(err)
	}

	return &config
}
