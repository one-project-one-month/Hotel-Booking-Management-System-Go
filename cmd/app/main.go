// Package main is the entry point for the Hotel Booking Management System
package main

import (
	"log"
	"net/http"
	"sync"

	"github.com/labstack/echo/v4"
	"github.com/one-project-one-month/Hotel-Booking-Management-System-Go/config"
	"github.com/one-project-one-month/Hotel-Booking-Management-System-Go/internal/auth"
	"github.com/one-project-one-month/Hotel-Booking-Management-System-Go/internal/bankaccount"
	"github.com/one-project-one-month/Hotel-Booking-Management-System-Go/internal/booking"
	"github.com/one-project-one-month/Hotel-Booking-Management-System-Go/internal/checkinout"
	"github.com/one-project-one-month/Hotel-Booking-Management-System-Go/internal/coupon"
	"github.com/one-project-one-month/Hotel-Booking-Management-System-Go/internal/invoice"
	"github.com/one-project-one-month/Hotel-Booking-Management-System-Go/internal/room"
	"github.com/one-project-one-month/Hotel-Booking-Management-System-Go/internal/user"
	"github.com/one-project-one-month/Hotel-Booking-Management-System-Go/pkg/mq"
	"github.com/one-project-one-month/Hotel-Booking-Management-System-Go/pkg/postgres"
)

func main() {
	cfg, err := config.New(".")
	if err != nil {
		log.Fatal(err)
	}

	db, err := postgres.New(&cfg.Postgres)
	if err != nil {
		log.Fatal(err)
	}

	var wg sync.WaitGroup
	app := NewApp(&wg, cfg)
	queue := mq.New(&wg, 100)

	auth.Run(app.echo, queue)
	user.Run(app.echo, db, queue, cfg)
	room.Run(app.echo, db, cfg)
	coupon.Run(app.echo, db, queue)
	booking.Run(app.echo, db, queue)
	checkinout.Run(app.echo, db, queue)
	bankaccount.Run(app.echo, db, queue)
	invoice.Run(app.echo, db)

	app.echo.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Welcome to Hotel Booking System APIs")
	})

	app.start()

	wg.Wait()
}
