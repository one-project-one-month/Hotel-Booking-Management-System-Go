package coupon

import (
	"github.com/labstack/echo/v4"
	"github.com/one-project-one-month/Hotel-Booking-Management-System-Go/pkg/mq"
	"gorm.io/gorm"
)

func Run(app *echo.Echo, db *gorm.DB, queue *mq.MQ) {
	repo := &Repository{database: db}
	service := &Service{repo: repo}
	handler := &Handler{service: service}

	app.POST("/api/v1/coupons", handler.create)
	app.GET("/api/v1/coupons", handler.findList)
	app.GET("/api/v1/coupons/:id", handler.findByID)
	app.PATCH("/api/v1/coupons/:id", handler.update)
	app.DELETE("/api/v1/coupons/:id", handler.delete)
}
