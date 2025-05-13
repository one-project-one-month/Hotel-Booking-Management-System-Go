package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	Environment string
	Server      Server
	Postgres    Postgres
}

type Server struct {
	Host string
	Port int
}

type Postgres struct {
	User     string
	Password string
	Host     string
	Port     int
	DbName   string
	SslMode  string
	TimeZone string
}

func load(path string) (*viper.Viper, error) {
	var v = viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yml")
	v.AddConfigPath(path)

	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("errors reading config file, %v", err)
	}

	return v, nil
}

func New(configPath string) (*Config, error) {
	v, err := load(configPath)
	if err != nil {
		return nil, err
	}

	var config Config
	if err := v.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return &config, nil
}
