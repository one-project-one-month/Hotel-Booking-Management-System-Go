package user

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/one-project-one-month/Hotel-Booking-Management-System-Go/pkg/response"
)

// Handler manages HTTP requests for user operations.
type Handler struct {
	service *Service // Will be implemented in future releases
}

func newHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) findAllUsers(ctx echo.Context) error {
	data, err := h.service.findAllUsers()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, &response.HTTPErrorResponse{
			Message: "Fetch Users Failed!",
			Error:   err,
		})
	}

	return ctx.JSON(http.StatusOK, &response.HTTPSuccessResponse{
		Data:    data,
		Message: "Fetch Users Success!",
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
	room, err := h.service.getUserByID(roomUUID)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, &response.HTTPErrorResponse{
			Message: "User Not Found!",
		})
	}

	return ctx.JSON(http.StatusOK, &response.HTTPSuccessResponse{
		Message: "User Found!",
		Data:    room,
	})
}

func (h *Handler) createRoom(ctx echo.Context) error {
	var newUser CreateUserDto
	err := ctx.Bind(&newUser)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, &response.HTTPErrorResponse{
			Message: err.Error(),
			Error:   err,
		})
	}
	err = ctx.Validate(&newUser)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, &response.HTTPErrorResponse{
			Message: err.Error(),
			Error:   err,
		})
	}
	createdUser, err := h.service.createUser(&newUser)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, &response.HTTPErrorResponse{
			Message: "Create User Failed!",
			Error:   err,
		})
	}

	return ctx.JSON(http.StatusCreated, &response.HTTPSuccessResponse{
		Message: "Create User Success!",
		Data:    createdUser,
	})
}

func (h *Handler) updateUser(ctx echo.Context) error {
	id := ctx.Param("id")
	userUUID, err := uuid.Parse(id)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, &response.HTTPErrorResponse{
			Message: "Invalid ID!",
		})
	}
	var user UpdateUserDto
	err = ctx.Bind(&user)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, &response.HTTPErrorResponse{
			Message: "Invalid Request Body!",
			Error:   err,
		})
	}
	err = ctx.Validate(&user)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, &response.HTTPErrorResponse{
			Message: err.Error(),
			Error:   err,
		})
	}
	updatedUser, err := h.service.updateUser(&user, userUUID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, &response.HTTPErrorResponse{
			Message: "Update User Failed!",
			Error:   err,
		})
	}

	return ctx.JSON(http.StatusOK, &response.HTTPSuccessResponse{
		Message: "Update User Success!",
		Data:    updatedUser,
	})
}

func (h *Handler) deleteUser(ctx echo.Context) error {
	id := ctx.Param("id")
	userUUID, err := uuid.Parse(id)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, &response.HTTPErrorResponse{
			Message: "Invalid ID!",
		})
	}
	err = h.service.deleteUserByID(userUUID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, &response.HTTPErrorResponse{
			Message: "Delete User Failed!",
		})
	}

	return ctx.JSON(http.StatusNoContent, &response.HTTPSuccessResponse{})
}

func (h *Handler) findCouponsByUserID(ctx echo.Context) error {
	return nil
}
