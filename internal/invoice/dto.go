package invoice

import (
	"time"

	"github.com/google/uuid"
)

type CreateInvoiceDto struct {
	CheckInOutID uuid.UUID `json:"check_in_out_id" validate:"required,uuid"`
	TotalAmount  float64   `json:"total_amount" validate:"required,gt=0"`
}

type UpdateInvoiceDto struct {
	TotalAmount float64 `json:"total_amount" validate:"required,gt=0"`
}

type ResponseInvoiceDto struct {
	ID           uuid.UUID `json:"id"`
	CheckInOutID uuid.UUID `json:"check_in_out_id"`
	TotalAmount  float64   `json:"total_amount"`
	CreatedAt    time.Time `json:"created_at"`
}
