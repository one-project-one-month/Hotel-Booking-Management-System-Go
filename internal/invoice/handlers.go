package invoice

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/one-project-one-month/Hotel-Booking-Management-System-Go/pkg/response"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) CreateInvoice(c echo.Context) error {
	var dto CreateInvoiceDto
	if err := c.Bind(&dto); err != nil {
		return c.JSON(http.StatusBadRequest, &response.HTTPErrorResponse{
			Message: "Invalid request body",
			Error:   err,
		})
	}

	if err := c.Validate(&dto); err != nil {
		return c.JSON(http.StatusBadRequest, &response.HTTPErrorResponse{
			Message: "Validation failed",
			Error:   err,
		})
	}

	result := h.service.CreateInvoice(&dto)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, &response.HTTPErrorResponse{
			Message: "Failed to create invoice",
			Error:   result.Error,
		})
	}

	return c.JSON(http.StatusCreated, &response.HTTPSuccessResponse{
		Message: "Invoice created successfully",
		Data:    result.Data,
	})
}

func (h *Handler) GetInvoiceByID(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, &response.HTTPErrorResponse{
			Message: "Invalid invoice ID",
		})
	}

	result := h.service.GetInvoiceByID(id)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, &response.HTTPErrorResponse{
			Message: "Invoice not found",
		})
	}

	return c.JSON(http.StatusOK, &response.HTTPSuccessResponse{
		Message: "Invoice retrieved successfully",
		Data:    result.Data,
	})
}

func (h *Handler) GetAllInvoices(c echo.Context) error {
	result := h.service.GetAllInvoices()
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, &response.HTTPErrorResponse{
			Message: "Failed to retrieve invoices",
			Error:   result.Error,
		})
	}

	return c.JSON(http.StatusOK, &response.HTTPSuccessResponse{
		Message: "Invoices retrieved successfully",
		Data:    result.Data,
	})
}

func (h *Handler) UpdateInvoice(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, &response.HTTPErrorResponse{
			Message: "Invalid invoice ID",
		})
	}

	var dto UpdateInvoiceDto
	if err := c.Bind(&dto); err != nil {
		return c.JSON(http.StatusBadRequest, &response.HTTPErrorResponse{
			Message: "Invalid request body",
			Error:   err,
		})
	}

	if err := c.Validate(&dto); err != nil {
		return c.JSON(http.StatusBadRequest, &response.HTTPErrorResponse{
			Message: "Validation failed",
			Error:   err,
		})
	}

	result := h.service.UpdateInvoice(id, &dto)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, &response.HTTPErrorResponse{
			Message: "Failed to update invoice",
			Error:   result.Error,
		})
	}

	return c.JSON(http.StatusOK, &response.HTTPSuccessResponse{
		Message: "Invoice updated successfully",
		Data:    result.Data,
	})
}

func (h *Handler) DeleteInvoice(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, &response.HTTPErrorResponse{
			Message: "Invalid invoice ID",
		})
	}

	result := h.service.DeleteInvoice(id)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, &response.HTTPErrorResponse{
			Message: "Failed to delete invoice",
			Error:   result.Error,
		})
	}

	return c.JSON(http.StatusNoContent, nil)
}
