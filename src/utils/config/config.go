package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	DB struct {
		Name string `mapstructure:"name"`
	} `mapstructure:"db"`

	Server struct {
		Port int `mapstructure:"port"`
	} `mapstructure:"server"`
}

func LoadConfig() *Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("/app")
	viper.AddConfigPath("./config")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("Unable to decode into struct: %v", err)
	}

	return &config
}
