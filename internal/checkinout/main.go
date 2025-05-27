package checkinout

import (
	"github.com/labstack/echo/v4"
	"github.com/one-project-one-month/Hotel-Booking-Management-System-Go/pkg/mq"
	"gorm.io/gorm"
)

func Run(e *echo.Echo, db *gorm.DB, queue *mq.MQ) {
	repo := NewRepository(db)
	service := NewService(repo, queue)
	handler := NewHandler(service)

	handler.RegisterRoutes(e)
}
