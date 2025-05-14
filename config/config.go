// Package config provides configuration loading and management functionality
package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Config represents the application configuration structure.
type Config struct {
	Environment string
	Server      Server
	Postgres    Postgres
}

// Server contains the HTTP server configuration.
type Server struct {
	Host string
	Port int
}

// Postgres contains the database connection configuration.
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
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yml")
	v.AddConfigPath(path)

	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("errors reading config file: %w", err)
	}

	return v, nil
}

// New creates a new Config instance loaded from the specified path.
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
