package coupon

import (
	echojwt "github.com/labstack/echo-jwt/v4"

	"github.com/labstack/echo/v4"
	"github.com/one-project-one-month/Hotel-Booking-Management-System-Go/internal/auth"
	"github.com/one-project-one-month/Hotel-Booking-Management-System-Go/pkg/mq"
	"gorm.io/gorm"
)

func Run(app *echo.Echo, db *gorm.DB, queue *mq.MQ) {
	repo := &Repository{database: db}
	service := &Service{repo: repo, queue: queue}
	handler := newHandler(service, queue)
	jwtConfig := auth.NewJWTConfig()

	g := app.Group("/api/v1/coupons")
	g.Use(echojwt.WithConfig(*jwtConfig))

	g.POST("", handler.create)
	g.GET("", handler.findList)
	g.GET("/:id", handler.findByID)
	g.PATCH("/:id", handler.update)
	g.DELETE("/:id", handler.delete)
}
