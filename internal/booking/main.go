// Package user handles user-related functionality including routes, handlers, and services
package booking

import (
	"github.com/labstack/echo/v4"
	"github.com/one-project-one-month/Hotel-Booking-Management-System-Go/pkg/mq"
	"gorm.io/gorm"
)

// Run configures and sets up user routes in the provided Echo instance.
func Run(e *echo.Echo, db *gorm.DB, queue *mq.MQ) {
	if err := Seed(db); err != nil {
		e.Logger.Fatal(err)
	}

	repo := newRepository(db)
	service := newService(repo)
	handler := newHandler(service, queue)
	g := e.Group("/api/v1/bookings")
	g.GET("", handler.findAllBookings)
	g.GET("/:id", handler.findBookingByID)
	g.POST("", handler.createBooking)
	g.PATCH("/:id", handler.updateBooking)
	g.DELETE("/:id", handler.deleteBooking)
}
