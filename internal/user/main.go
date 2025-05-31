// Package user handles user-related functionality including routes, handlers, and services
package user

import (
	"github.com/labstack/echo/v4"
	"github.com/one-project-one-month/Hotel-Booking-Management-System-Go/config"
	"github.com/one-project-one-month/Hotel-Booking-Management-System-Go/pkg/mq"
	"gorm.io/gorm"
)

// Run configures and sets up user routes in the provided Echo instance.
func Run(e *echo.Echo, db *gorm.DB, queue *mq.MQ, cfg *config.Config) {
	// if cfg.Environment == "development" {
	if err := Seed(db); err != nil {
		e.Logger.Fatal(err)
	}
	//}

	repo := newRepository(db)
	service := newService(repo, queue)
	handler := newHandler(service, queue)

	g := e.Group("/api/v1/users")
	g.GET("", handler.findAllUsers)
	g.GET("/:id", handler.findRoomByID)
	g.POST("", handler.createRoom)
	g.PATCH("/:id", handler.updateUser)
	g.DELETE("/:id", handler.deleteUser)

	g.GET("/:id/coupons", handler.findCouponsByUserID)
}
