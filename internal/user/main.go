// Package user handles user-related functionality including routes, handlers, and services
package user

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Run configures and sets up user routes in the provided Echo instance.
func Run(e *echo.Echo) {
	e.GET("/api/v1/users", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Welcome to Hotel Booking System/Users APIs")
	})
}
