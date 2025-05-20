// Package main is the entry point for the Hotel Booking Management System
package main

import (
	"github.com/one-project-one-month/Hotel-Booking-Management-System-Go/internal/coupon"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/one-project-one-month/Hotel-Booking-Management-System-Go/config"
	"github.com/one-project-one-month/Hotel-Booking-Management-System-Go/internal/room"
	"github.com/one-project-one-month/Hotel-Booking-Management-System-Go/internal/user"
	"github.com/one-project-one-month/Hotel-Booking-Management-System-Go/pkg/postgres"
	"github.com/one-project-one-month/Hotel-Booking-Management-System-Go/pkg/requestValidator"
)

func main() {
	app := echo.New()
	app.Validator = &requestValidator.CustomValidator{Validator: validator.New()}
	configData, err := config.New(".")
	if err != nil {
		app.Logger.Fatal(err)
	}
	db, err := postgres.New(&configData.Postgres)
	if err != nil {
		app.Logger.Fatal(err)
	}

	user.Run(app, db)
	room.Run(app, db)
	coupon.Run(app, db)

	app.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Welcome to Hotel Booking System APIs")
	})

	app.Logger.Fatal(app.Start(":8080"))
}
