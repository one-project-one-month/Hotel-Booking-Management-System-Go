package room

import (
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/one-project-one-month/Hotel-Booking-Management-System-Go/pkg/response"
)

// Handler
type Handler struct {
	service *Service
}

func newHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) findAllRooms(ctx echo.Context) error {
	data, err := h.service.findAllRooms()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, &response.HTTPErrorResponse{
			Message: "Fetch Rooms Failed!",
			Error:   err,
		})
	}

	return ctx.JSON(http.StatusOK, &response.HTTPSuccessResponse{
		Data:    data,
		Message: "Fetch Rooms Success!",
	})
}

func (h *Handler) findRoomByID(ctx echo.Context) error {
	id := ctx.Param("id")
	roomUUID, err := uuid.Parse(id)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, &response.HTTPErrorResponse{
			Message: "Invalid ID!",
		})
	}
	room, err := h.service.getRoomByID(roomUUID)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, &response.HTTPErrorResponse{
			Message: "Room Not Found!",
		})
	}

	return ctx.JSON(http.StatusOK, &response.HTTPSuccessResponse{
		Message: "Room Found!",
		Data:    room,
	})
}

func (h *Handler) createRoom(ctx echo.Context) error {
	var newRoom RequestRoomDto

	err := ctx.Bind(&newRoom)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, &response.HTTPErrorResponse{
			Message: err.Error(),
			Error:   err,
		})
	}
	err = ctx.Validate(&newRoom)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, &response.HTTPErrorResponse{
			Message: err.Error(),
			Error:   err,
		})
	}
	newRoomID, err := h.service.createRoom(&newRoom)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, &response.HTTPErrorResponse{
			Message: "Create Room Failed!",
			Error:   err,
		})
	}

	return ctx.JSON(http.StatusCreated, &response.HTTPSuccessResponse{
		Message: "Create Room Success!",
		Data:    newRoomID,
	})
}

func (h *Handler) updateRoom(ctx echo.Context) error {
	id := ctx.Param("id")
	roomUUID, err := uuid.Parse(id)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, &response.HTTPErrorResponse{
			Message: "Invalid ID!",
		})
	}
	var room RequestRoomDto
	err = ctx.Bind(&room)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, &response.HTTPErrorResponse{
			Message: "Invalid Request Body!",
			Error:   err,
		})
	}
	err = ctx.Validate(&room)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, &response.HTTPErrorResponse{
			Message: err.Error(),
			Error:   err,
		})
	}
	updatedRoom, err := h.service.updateRoom(&room, roomUUID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, &response.HTTPErrorResponse{
			Message: "Update Room Failed!",
			Error:   err,
		})
	}

	return ctx.JSON(http.StatusOK, &response.HTTPSuccessResponse{
		Message: "Update Room Success!",
		Data:    updatedRoom,
	})
}

func (h *Handler) updateRoomStatus(ctx echo.Context) error {
	id := ctx.Param("id")
	roomUUID, err := uuid.Parse(id)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, &response.HTTPErrorResponse{
			Message: "Invalid ID!",
		})
	}
	var status RequestRoomStatusDto
	err = ctx.Bind(&status)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, &response.HTTPErrorResponse{
			Message: "Invalid Request Body!",
			Error:   err,
		})
	}
	err = ctx.Validate(&status)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, &response.HTTPErrorResponse{
			Message: err.Error(),
			Error:   err,
		})
	}
	updatedRoom, err := h.service.updateRoomStatus(status.Status, roomUUID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, &response.HTTPErrorResponse{
			Message: "Update Room Failed!",
			Error:   err,
		})
	}
	return ctx.JSON(http.StatusOK, &response.HTTPSuccessResponse{
		Message: "Update Room Success!",
		Data: map[string]string{
			"id":     updatedRoom.String(),
			"status": status.Status,
		},
	})
}

func (h *Handler) updateRoomIsFeatured(ctx echo.Context) error {
	id := ctx.Param("id")
	roomUUID, err := uuid.Parse(id)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, &response.HTTPErrorResponse{
			Message: "Invalid ID!",
		})
	}

	updatedRoom, err := h.service.updateRoomIsFeatured(roomUUID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, &response.HTTPErrorResponse{
			Message: "Update Room Failed!",
			Error:   err,
		})
	}
	return ctx.JSON(http.StatusOK, &response.HTTPSuccessResponse{
		Message: "Update Room Success!",
		Data:    updatedRoom,
	})
}

func (h *Handler) deleteRoom(ctx echo.Context) error {
	id := ctx.Param("id")
	roomUUID, err := uuid.Parse(id)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, &response.HTTPErrorResponse{
			Message: "Invalid ID!",
		})
	}
	err = h.service.deleteRoomByID(roomUUID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, &response.HTTPErrorResponse{
			Message: "Delete Room Failed!",
		})
	}

	return ctx.JSON(http.StatusNoContent, &response.HTTPSuccessResponse{})
}

func (h *Handler) getRoomByGuestLimit(ctx echo.Context) error {
	guestLimits := ctx.QueryParam("total_guests")
	guestLimitsInt, err := strconv.ParseInt(guestLimits, 10, 64)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, &response.HTTPErrorResponse{
			Message: "Invalid Guest Limit!",
		})
	}
	rooms, err := h.service.getRoomByGuestLimit(int(guestLimitsInt))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, &response.HTTPErrorResponse{
			Message: "Get Room By Guest Limit Failed!",
			Error:   err,
		})
	}
	return ctx.JSON(http.StatusOK, &response.HTTPSuccessResponse{
		Message: "Get Room By Guest Limit Success!",
		Data:    rooms,
	})
}
