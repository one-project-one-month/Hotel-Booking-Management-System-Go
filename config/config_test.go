package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	t.Run("Fail to load config with invalid path", func(t *testing.T) {
		path := "."
		_, err := load(path)
		assert.NotNil(t, err)
	})

	t.Run("Success to load config", func(t *testing.T) {
		path := "./testdata"
		v, err := load(path)
		assert.Nil(t, err)
		assert.NotNil(t, v)
		assert.Equal(t, v.Get("server.host"), "127.0.0.1")
		assert.Equal(t, v.Get("server.port"), 8080)
	})
}

func TestNewConfig(t *testing.T) {
	expected := &Config{
		Environment: "development",
		Server: Server{
			Host: "127.0.0.1",
			Port: 8080,
		},
		Postgres: Postgres{
			User:     "postgres",
			Password: "12345678",
			Host:     "127.0.0.1",
			Port:     5432,
			DbName:   "hotel_booking",
			SslMode:  "disable",
			TimeZone: "Asia/Yangon",
		},
	}
	path := "./testdata"
	config, err := New(path)

	assert.Nil(t, err)
	assert.NotNil(t, config)
	assert.Equal(t, expected, config)
}
