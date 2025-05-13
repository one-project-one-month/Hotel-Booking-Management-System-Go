package postgres

import (
	"fmt"

	"github.com/one-project-one-month/Hotel-Booking-Management-System-Go/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New(cfg *config.Postgres) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
		cfg.Host, cfg.User, cfg.Password, cfg.DbName, cfg.Port, cfg.SslMode, cfg.TimeZone)
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
