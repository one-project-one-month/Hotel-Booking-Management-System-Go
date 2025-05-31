package auth

import (
	"github.com/labstack/echo/v4"
	"github.com/one-project-one-month/Hotel-Booking-Management-System-Go/pkg/mq"
)

func Run(e *echo.Echo, queue *mq.MQ) {
	service := newService(queue)
	handler := newHandler(service)

	handler.RegisterRoutes(e)
}
