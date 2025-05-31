package auth

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/one-project-one-month/Hotel-Booking-Management-System-Go/pkg/response"
)

type Handler struct {
	service *Service
}

func (h *Handler) RegisterRoutes(e *echo.Echo) {
	e.POST("api/v1/auth/signin", h.Signin)
	e.POST("api/v1/auth/signup", h.Signup)
}

func (h *Handler) Signin(c echo.Context) error {
	var user SignInUserDto
	if err := c.Bind(&user); err != nil {
		return err
	}

	if err := c.Validate(user); err != nil {
		return c.JSON(http.StatusBadRequest, response.HTTPErrorResponse{
			Message: err.Error(),
		})
	}

	resp := h.service.Signin(&user)
	if errors.Is(resp.Error, response.ErrNotFound) {
		return c.JSON(http.StatusUnauthorized, &response.HTTPErrorResponse{
			Message: "Invalid email",
		})
	}

	if errors.Is(resp.Error, response.ErrUnauthorized) {
		return c.JSON(http.StatusUnauthorized, &response.HTTPErrorResponse{
			Message: "Invalid password",
		})
	}

	if resp.Error != nil {
		return c.JSON(http.StatusInternalServerError, &response.HTTPErrorResponse{
			Message: "Error while signing in",
		})
	}

	return c.JSON(http.StatusOK, &response.HTTPSuccessResponse{
		Message: "User signed in successfully",
		Data:    resp.Data,
	})
}

func (h *Handler) Signup(c echo.Context) error {
	var user SignUpUserDto
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, &response.HTTPErrorResponse{
			Message: err.Error(),
		})
	}

	if err := c.Validate(user); err != nil {
		return c.JSON(http.StatusBadRequest, &response.HTTPErrorResponse{
			Message: err.Error(),
		})
	}

	resp := h.service.Signup(&user)

	if errors.Is(resp.Error, response.ErrConflict) {
		return c.JSON(http.StatusConflict, &response.HTTPErrorResponse{
			Message: fmt.Sprintf("User with email %s already exists", user.Email),
		})
	}

	if resp.Error != nil {
		return c.JSON(http.StatusInternalServerError, &response.HTTPErrorResponse{
			Message: "Error while signing up",
		})
	}

	return c.JSON(http.StatusCreated, response.HTTPSuccessResponse{
		Message: "User created successfully",
		Data:    resp.Data,
	})
}

func newHandler(service *Service) *Handler {
	return &Handler{service: service}
}
