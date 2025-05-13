package postgres

import (
	"testing"

	"github.com/one-project-one-month/Hotel-Booking-Management-System-Go/config"
	"github.com/stretchr/testify/assert"
)

func TestNewPostgres(t *testing.T) {
	dbConfig := config.Postgres{
		User:     "postgres",
		Password: "12345678",
		Host:     "127.0.0.1",
		Port:     5432,
		DbName:   "hotel_booking",
		SslMode:  "disable",
		TimeZone: "Asia/Yangon",
	}
	postgres, err := New(&dbConfig)

	assert.Nil(t, err)
	assert.NotNil(t, postgres)
	assert.Equal(t, "postgres", postgres.Name())
}
