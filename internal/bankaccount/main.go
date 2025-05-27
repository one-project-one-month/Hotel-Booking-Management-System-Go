package bankaccount

import (
	"github.com/labstack/echo/v4"
	"github.com/one-project-one-month/Hotel-Booking-Management-System-Go/pkg/mq"
	"gorm.io/gorm"
)

func Run(e *echo.Echo, db *gorm.DB, queue *mq.MQ) {
	if err := Seed(db); err != nil {
		e.Logger.Fatal(err)
	}

	handler := NewHandler(NewService(&repository{db: db}, queue))
	handler.RegisterRoutes(e)
}
