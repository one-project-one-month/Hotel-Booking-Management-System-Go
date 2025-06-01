package auth

import (
	"os"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/one-project-one-month/Hotel-Booking-Management-System-Go/pkg/mq"
)

func NewJWTConfig() *echojwt.Config {
	secret := os.Getenv("JWT_SECRET")
	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(JWTClaim)
		},
		SigningKey: []byte(secret),
	}

	return &config
}

func Run(e *echo.Echo, queue *mq.MQ) {
	service := newService(queue)
	handler := newHandler(service)

	handler.RegisterRoutes(e)
}
