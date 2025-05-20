// Package user handles user-related functionality including routes, handlers, and services
package user

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// Run configures and sets up user routes in the provided Echo instance.
func Run(e *echo.Echo, db *gorm.DB) {
	if err := Seed(db); err != nil {
		e.Logger.Fatal(err)
	}

	repo := newRepository(db)
	service := newService(repo)
	handler := newHandler(service)

	g := e.Group("/api/v1/user")
	g.GET("", handler.findAllUsers)
	g.GET("/:id", handler.findRoomByID)
	g.POST("", handler.createRoom)
	g.PATCH("/:id", handler.updateUser)
	g.DELETE("/:id", handler.deleteUser)
}
