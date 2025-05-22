package coupon

import (
	"errors"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/one-project-one-month/Hotel-Booking-Management-System-Go/pkg/events"
	"github.com/one-project-one-month/Hotel-Booking-Management-System-Go/pkg/mq"
	"github.com/one-project-one-month/Hotel-Booking-Management-System-Go/pkg/response"
)

type Handler struct {
	service *Service
	queue   *mq.MQ
}

func newHandler(service *Service, queue *mq.MQ) *Handler {
	handler := &Handler{service: service, queue: queue}

	queue.Subscribe(events.COUPONFETCHED, handler.findByUserID)

	return handler
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

	if resp := h.service.create(&coupon); resp.Error != nil {
		return c.JSON(http.StatusInternalServerError, response.HTTPErrorResponse{
			Message: "Failed to create coupon",
			Error:   resp.Error,
		})
	}

	return c.JSON(http.StatusCreated, response.HTTPSuccessResponse{
		Message: "Coupon created successfully",
		Data:    nil,
	})
}

func (h *Handler) findList(c echo.Context) error {
	query := bindQuery(c)

	coupons := h.service.findList(query)

	return c.JSON(http.StatusOK, response.HTTPSuccessResponse{
		Message: "Coupon list fetched successfully",
		Data:    coupons.Data,
	})
}

func (h *Handler) findByID(c echo.Context) error {
	paramId := c.Param("id")

	id, err := uuid.Parse(paramId)
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusBadRequest, response.HTTPErrorResponse{
			Message: "Invalid ID",
		})
	}

	resp := h.service.findByID(id.String())
	if resp.Data == nil {
		return c.JSON(http.StatusNotFound, response.HTTPSuccessResponse{
			Message: "Coupon not found",
			Data:    nil,
		})
	}

	if resp.Error != nil {
		return c.JSON(http.StatusInternalServerError, response.HTTPErrorResponse{
			Message: "Failed to fetch coupon",
		})
	}

	return c.JSON(http.StatusOK, response.HTTPSuccessResponse{
		Message: "Coupon fetched successfully",
		Data:    resp.Data,
	})
}

func (h *Handler) update(c echo.Context) error {
	paramId := c.Param("id")
	id, err := uuid.Parse(paramId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.HTTPErrorResponse{
			Message: "Invalid ID",
		})
	}

	var coupon UpdateCouponDto
	if err := c.Bind(&coupon); err != nil {
		return c.JSON(http.StatusBadRequest, response.HTTPErrorResponse{
			Message: "Invalid request body",
		})
	}

	if err := c.Validate(coupon); err != nil {
		return c.JSON(http.StatusBadRequest, response.HTTPErrorResponse{
			Message: err.Error(),
		})
	}

	if err := c.Validate(coupon.Data); err != nil {
		return c.JSON(http.StatusBadRequest, response.HTTPErrorResponse{
			Message: err.Error(),
		})
	}

	resp := h.service.update(id.String(), &coupon)
	if resp.Error != nil {
		if errors.Is(resp.Error, response.ErrNotFound) {
			return c.JSON(http.StatusBadRequest, &response.HTTPErrorResponse{
				Message: resp.Error.Error(),
			})
		}

		if errors.Is(resp.Error, response.ErrBadRequest) {
			return c.JSON(http.StatusBadRequest, &response.HTTPErrorResponse{
				Message: resp.Message,
			})
		}

		return c.JSON(http.StatusInternalServerError, &response.HTTPErrorResponse{
			Message: "Failed to update coupon",
		})
	}

	return c.JSON(http.StatusOK, response.HTTPSuccessResponse{
		Message: "Coupon updated successfully",
		Data:    resp.Data,
	})
}

func (h *Handler) delete(c echo.Context) error {
	paramId := c.Param("id")
	id, err := uuid.Parse(paramId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.HTTPErrorResponse{
			Message: "Invalid ID",
		})
	}

	if resp := h.service.delete(id); resp.Error != nil {
		return c.JSON(http.StatusInternalServerError, response.HTTPErrorResponse{
			Message: "Failed to delete coupon",
		})
	}

	return c.JSON(http.StatusOK, response.HTTPSuccessResponse{
		Message: "Coupon deleted successfully",
	})
}

func (h *Handler) findByUserID(data any) any {
	dto := data.(events.FindByUserIdDto)
	return h.service.findByUserID(dto.UserID)
}
