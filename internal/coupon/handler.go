package coupon

import (
	"github.com/labstack/echo/v4"
	"github.com/one-project-one-month/Hotel-Booking-Management-System-Go/pkg/response"
	"net/http"
)

type Handler struct {
	service *Service
}

func (h *Handler) create(c echo.Context) error {
	var coupon CreateCouponDto

	if err := c.Bind(&coupon); err != nil {
		return c.JSON(http.StatusInternalServerError, response.HTTPErrorResponse{
			Message: "Invalid request body",
			Error:   err,
		})
	}

	if err := c.Validate(coupon); err != nil {
		return c.JSON(http.StatusBadRequest, response.HTTPErrorResponse{
			Message: err.Error(),
			Error:   err,
		})
	}

	return c.JSON(http.StatusCreated, response.HTTPSuccessResponse{
		Message: "Coupon created successfully",
		Data:    h.service.create(&coupon),
	})
}
