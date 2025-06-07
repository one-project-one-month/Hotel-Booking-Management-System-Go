package booking

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/one-project-one-month/Hotel-Booking-Management-System-Go/pkg/mq"
	"github.com/one-project-one-month/Hotel-Booking-Management-System-Go/pkg/response"
)

// Handler manages HTTP requests for booking operations.
type Handler struct {
	queue   *mq.MQ
	service *Service // Will be implemented in future releases
}

func newHandler(service *Service, queue *mq.MQ) *Handler {
	return &Handler{service: service, queue: queue}
}

func (h *Handler) findAllBookings(ctx echo.Context) error {
	data, err := h.service.findAllBookings()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, &response.HTTPErrorResponse{
			Message: "Fetch Bookings Failed!",
			Error:   err,
		})
	}

	return ctx.JSON(http.StatusOK, &response.HTTPSuccessResponse{
		Data:    data,
		Message: "Fetch Bookings Success!",
	})
}

func (h *Handler) findBookingByID(ctx echo.Context) error {
	id := ctx.Param("id")
	bookingUUID, err := uuid.Parse(id)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, &response.HTTPErrorResponse{
			Message: "Invalid ID!",
		})
	}
	booking, err := h.service.getBookingByID(bookingUUID)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, &response.HTTPErrorResponse{
			Message: "Booking Not Found!",
		})
	}

	return ctx.JSON(http.StatusOK, &response.HTTPSuccessResponse{
		Message: "Booking Found!",
		Data:    booking,
	})
}

func (h *Handler) createBooking(ctx echo.Context) error {
	var newBooking CreateBookingDto
	err := ctx.Bind(&newBooking)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, &response.HTTPErrorResponse{
			Message: err.Error(),
			Error:   err,
		})
	}
	err = ctx.Validate(&newBooking)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, &response.HTTPErrorResponse{
			Message: err.Error(),
			Error:   err,
		})
	}
	createdBooking := h.service.createBooking(&newBooking)

	if createdBooking.Error != nil {
		if createdBooking.Error == response.ErrNotFound {
			return ctx.JSON(http.StatusNotFound, &response.HTTPErrorResponse{
				Message: createdBooking.Message,
				Error:   createdBooking.Error,
			})
		}

		return ctx.JSON(http.StatusInternalServerError, &response.HTTPErrorResponse{
			Message: "Create Booking Failed!",
		})

	}

	return ctx.JSON(http.StatusCreated, &response.HTTPSuccessResponse{
		Message: "Create Booking Success!",
		Data:    createdBooking,
	})
}

func (h *Handler) updateBooking(ctx echo.Context) error {
	id := ctx.Param("id")
	bookingUUID, err := uuid.Parse(id)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, &response.HTTPErrorResponse{
			Message: "Invalid ID!",
		})
	}
	var booking UpdateBookingDto
	err = ctx.Bind(&booking)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, &response.HTTPErrorResponse{
			Message: "Invalid Request Body!",
			Error:   err,
		})
	}
	err = ctx.Validate(&booking)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, &response.HTTPErrorResponse{
			Message: err.Error(),
			Error:   err,
		})
	}
	updatedBooking, err := h.service.updateBooking(&booking, bookingUUID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, &response.HTTPErrorResponse{
			Message: "Update Booking Failed!",
			Error:   err,
		})
	}

	return ctx.JSON(http.StatusOK, &response.HTTPSuccessResponse{
		Message: "Update Booking Success!",
		Data:    updatedBooking,
	})
}

func (h *Handler) deleteBooking(ctx echo.Context) error {
	id := ctx.Param("id")
	bookingUUID, err := uuid.Parse(id)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, &response.HTTPErrorResponse{
			Message: "Invalid ID!",
		})
	}
	err = h.service.deleteBookingByID(bookingUUID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, &response.HTTPErrorResponse{
			Message: "Delete Booking Failed!",
		})
	}

	return ctx.JSON(http.StatusNoContent, &response.HTTPSuccessResponse{})
}
