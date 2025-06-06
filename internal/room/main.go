package room

import (
	"github.com/labstack/echo/v4"
	"github.com/one-project-one-month/Hotel-Booking-Management-System-Go/config"
	"github.com/one-project-one-month/Hotel-Booking-Management-System-Go/pkg/mq"
	"gorm.io/gorm"
)

// Run Entry Point For Room Feature
func Run(e *echo.Echo, db *gorm.DB, cfg *config.Config, queue *mq.MQ) {
	if err := Seed(db); err != nil {
		e.Logger.Fatal(err)
	}

	repo := newRepository(db)
	service := newService(repo, queue)
	handler := newHandler(service)

	g := e.Group("/api/v1/room")
	g.GET("", handler.findAllRooms)
	g.GET("/:id", handler.findRoomByID)
	g.POST("", handler.createRoom)
	g.PATCH("/:id", handler.updateRoom)
	g.DELETE("/:id", handler.deleteRoom)
	g.PATCH("/:id/status", handler.updateRoomStatus)
	g.PATCH("/:id/is_featured", handler.updateRoomIsFeatured)
	g.GET("/search", handler.getRoomByGuestLimit)
}
