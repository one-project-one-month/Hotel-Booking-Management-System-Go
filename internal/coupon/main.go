package coupon

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Run(app *echo.Echo, db *gorm.DB) {
	repo := &Repository{database: db}
	service := &Service{repo: repo}
	handler := &Handler{service: service}

	app.POST("/api/v1/coupons", handler.create)
}
