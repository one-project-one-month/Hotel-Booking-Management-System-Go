package invoice

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Run(e *echo.Echo, db *gorm.DB) {
	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)

	// Register routes
	g := e.Group("/api/v1/invoices")
	g.POST("", handler.CreateInvoice)
	g.GET("", handler.GetAllInvoices)
	g.GET("/:id", handler.GetInvoiceByID)
	g.PUT("/:id", handler.UpdateInvoice)
	g.DELETE("/:id", handler.DeleteInvoice)
}
