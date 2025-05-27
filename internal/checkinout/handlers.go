package checkinout

import (
	"net/http"

	"github.com/google/uuid"
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
	checkInOut := e.Group("/api/v1/check-in-out")
	{
		checkInOut.POST("", h.Create)
		checkInOut.GET("", h.GetAll)
		checkInOut.GET("/:id", h.GetByID)
		checkInOut.PATCH("/:id", h.Update)
		checkInOut.DELETE("/:id", h.Delete)
	}
}

func (h *Handler) Create(c echo.Context) error {
	var dto CreateCheckInOutDto
	if err := c.Bind(&dto); err != nil {
		return c.JSON(http.StatusBadRequest, &response.HTTPErrorResponse{
			Message: "Invalid Request Body!",
		})
	}

	if err := c.Validate(&dto); err != nil {
		return c.JSON(http.StatusBadRequest, &response.HTTPErrorResponse{
			Message: err.Error(),
		})
	}

	checkInOut := h.service.Create(c.Request().Context(), dto)
	if checkInOut.Error != nil {
		return c.JSON(http.StatusInternalServerError, &response.HTTPErrorResponse{
			Message: checkInOut.Message,
		})
	}

	return c.JSON(http.StatusCreated, &response.HTTPSuccessResponse{
		Message: "Create Check-in/out Success!",
		Data:    checkInOut,
	})
}

func (h *Handler) GetByID(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, &response.HTTPErrorResponse{
			Message: "Invalid ID!",
		})
	}

	checkInOut := h.service.GetByID(c.Request().Context(), id)
	if checkInOut.Error != nil {
		return c.JSON(http.StatusInternalServerError, &response.HTTPErrorResponse{
			Message: "Get Check-in/out Failed!",
		})
	}

	return c.JSON(http.StatusOK, &response.HTTPSuccessResponse{
		Message: "Get Check-in/out Success!",
		Data:    checkInOut,
	})
}

func (h *Handler) GetAll(c echo.Context) error {
	checkInOuts := h.service.GetAll(c.Request().Context())
	if checkInOuts.Error != nil {
		return c.JSON(http.StatusInternalServerError, &response.HTTPErrorResponse{
			Message: "Get All Check-in/out Failed!",
		})
	}

	return c.JSON(http.StatusOK, &response.HTTPSuccessResponse{
		Message: "Get All Check-in/out Success!",
		Data:    checkInOuts,
	})
}

func (h *Handler) Update(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, &response.HTTPErrorResponse{
			Message: "Invalid ID!",
		})
	}

	var dto UpdateCheckInOutDto
	if err := c.Bind(&dto); err != nil {
		return c.JSON(http.StatusBadRequest, &response.HTTPErrorResponse{
			Message: "Invalid Request Body!",
		})
	}

	if err := c.Validate(&dto); err != nil {
		return c.JSON(http.StatusBadRequest, &response.HTTPErrorResponse{
			Message: err.Error(),
		})
	}

	checkInOut := h.service.Update(c.Request().Context(), id, dto)
	if checkInOut.Error != nil {
		return c.JSON(http.StatusInternalServerError, &response.HTTPErrorResponse{
			Message: "Update Check-in/out Failed!",
		})
	}

	return c.JSON(http.StatusOK, &response.HTTPSuccessResponse{
		Message: "Update Check-in/out Success!",
		Data:    checkInOut,
	})
}

func (h *Handler) Delete(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, &response.HTTPErrorResponse{
			Message: "Invalid ID!",
		})
	}

	checkInOut := h.service.Delete(c.Request().Context(), id)
	if checkInOut.Error != nil {
		return c.JSON(http.StatusInternalServerError, &response.HTTPErrorResponse{
			Message: "Delete Check-in/out Failed!",
		})
	}

	return c.NoContent(http.StatusNoContent)
}
