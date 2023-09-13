package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Storage struct {
		StorageType      string `mapstructure:"storage_type"`
		ConnectionString string `mapstructure:"connection_string"`
	} `mapstructure:"storage"`
	Server struct {
		Host string `mapstructure:"host"`
		Port int    `mapstructure:"port"`
	} `mapstructure:"server"`
}

func setViperOpts() {
	viper.SetDefault("storage.storage_type", "memory")
	viper.SetDefault("storage.connection_string", "")

	viper.SetDefault("server.host", "127.0.0.1")
	viper.SetDefault("server.port", 8000)

	viper.SetConfigName("sheep")
	viper.SetConfigType("toml")
	viper.AddConfigPath("./conf/")
	viper.AddConfigPath("$HOME/.sheep/")
	viper.AddConfigPath("/etc/sheep/")
}

func NewConfig() (*Config, error) {
	setViperOpts()
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, fmt.Errorf("config file not found: %s", err)
		} else {
			return nil, fmt.Errorf("encountered error while parsing config: %s", err)
		}
	}

	var cfg Config

	err := viper.Unmarshal(&cfg)
	if err != nil {
		return nil, fmt.Errorf("unable to unmarshal config: %s", err)
	}

	return &cfg, nil
}
