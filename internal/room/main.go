package room

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// Run Entry Point For Room Feature
func Run(e *echo.Echo, db *gorm.DB) {
	if err := Seed(db); err != nil {
		e.Logger.Fatal(err)
	}

	repo := newRepository(db)
	service := newService(repo)
	handler := newHandler(service)

	g := e.Group("/api/v1/room")
	g.GET("", handler.findAllRooms)
	g.GET("/:id", handler.findRoomByID)
	g.POST("", handler.createRoom)
	g.PATCH("/:id", handler.updateRoom)
	g.DELETE("/:id", handler.deleteRoom)
}
