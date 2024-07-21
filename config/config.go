package config

import (
	"github.com/spf13/viper"
)

type AppConfig struct {
	Hotelbeds struct {
		API    string
		Secret string
	}
}

// Config load environment variables
func Config() *AppConfig {
	return loadConfig()
}

func loadConfig() *AppConfig {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	_ = viper.ReadInConfig()

	cfg := &AppConfig{}

	cfg.Hotelbeds.API = viper.GetString("HOTELBEDS_API_KEY")
	cfg.Hotelbeds.Secret = viper.GetString("HOTELBEDS_SECRET")

	return cfg
}
