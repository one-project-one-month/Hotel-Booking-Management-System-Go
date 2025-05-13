package user

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func Run(e *echo.Echo) {
	e.GET("/api/v1/users", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Welcome to Hotel Booking System/Users APIs")
	})
}
