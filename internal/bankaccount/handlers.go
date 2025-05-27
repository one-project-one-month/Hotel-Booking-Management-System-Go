package bankaccount

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/one-project-one-month/Hotel-Booking-Management-System-Go/pkg/response"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) RegisterRoutes(e *echo.Echo) {
	bankAccount := e.Group("/api/v1/bank-accounts")
	bankAccount.GET("", h.GetAll)
}

func (h *Handler) GetAll(c echo.Context) error {
	return c.JSON(http.StatusOK, &response.HTTPSuccessResponse{
		Message: "Get All Bank Accounts Success!",
		Data:    h.service.GetAll(c.Request().Context()).Data,
	})
}
