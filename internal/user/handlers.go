package user

import (
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/one-project-one-month/Hotel-Booking-Management-System-Go/pkg/events"
	"github.com/one-project-one-month/Hotel-Booking-Management-System-Go/pkg/mq"
	"github.com/one-project-one-month/Hotel-Booking-Management-System-Go/pkg/response"
)

// Handler manages HTTP requests for user operations.
type Handler struct {
	queue   *mq.MQ
	service *Service // Will be implemented in future releases
}

func newHandler(service *Service, queue *mq.MQ) *Handler {
	return &Handler{service: service, queue: queue}
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
	id := ctx.Param("id")
	userID, err := uuid.Parse(id)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, &response.HTTPErrorResponse{
			Message: "Invalid ID!",
		})
	}

	reply := h.queue.Publish(&mq.Message{
		AppID: "UserService",
		Topic: events.COUPONFETCHED,
		Data:  events.FindByUserIdDto{UserID: userID.String()},
	})

	select {
	case resp := <-reply:
		data := resp.(*response.ServiceResponse)
		if data.Error == nil {
			return ctx.JSON(http.StatusOK, &response.HTTPSuccessResponse{
				Message: "Fetch Coupons Success!",
				Data:    data.Data,
			})
		}
	case <-time.Tick(1 * time.Second):
		return ctx.JSON(http.StatusInternalServerError, &response.HTTPErrorResponse{
			Message: "Timeout!",
		})
	}

	return ctx.JSON(http.StatusInternalServerError, &response.HTTPErrorResponse{
		Message: "Fetch Coupons Failed!",
	})
}
