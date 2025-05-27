// Package postgres provides database connection and management functionality.
package postgres

import (
	"fmt"
	"log"

	"github.com/one-project-one-month/Hotel-Booking-Management-System-Go/pkg/models"

	"github.com/one-project-one-month/Hotel-Booking-Management-System-Go/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// New creates a new PostgreSQL database connection using the provided configuration.
func New(cfg *config.Postgres) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
		cfg.Host, cfg.User, cfg.Password, cfg.DbName, cfg.Port, cfg.SslMode, cfg.TimeZone)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	//if cfg.Host == "127.0.0.1" {

		err = db.AutoMigrate(&models.User{}, &models.Room{}, &models.Coupon{}, &models.Booking{}, &models.CheckInOut{})
		if err != nil {
			log.Fatal(err)
		}

	//}

	return db, nil
}
