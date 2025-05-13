package main

import (
	"github.com/one-project-one-month/Hotel-Booking-Management-System-Go/internal/user"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()

	user.Run(app)

	app.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Welcome to Hotel Booking System APIs")
	})

	app.Logger.Fatal(app.Start(":8080"))
}
