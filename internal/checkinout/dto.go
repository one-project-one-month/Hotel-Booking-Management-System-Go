package checkinout

import (
	"time"

	"github.com/google/uuid"
)

type CreateCheckInOutDto struct {
	CheckIn     time.Time `json:"checkIn"`
	CheckOut    time.Time `json:"checkOut"`
	BookingID   uuid.UUID `json:"bookingId" validate:"required,uuid"`
	Status      string    `json:"status" validate:"required"`
	ExtraCharge float64   `json:"extraCharge"`
}

type UpdateCheckInOutDto struct {
	CheckIn     *time.Time `json:"checkIn,omitempty" validate:"omitempty"`
	CheckOut    *time.Time `json:"checkOut,omitempty" validate:"omitempty,gtfield=CheckIn"`
	Status      string     `json:"status,omitempty" validate:"omitempty"`
	ExtraCharge float64    `json:"extraCharge,omitempty" validate:"omitempty"`
}

type ResponseCheckInOutDto struct {
	ID          uuid.UUID  `json:"id"`
	CheckIn     time.Time  `json:"checkIn"`
	CheckOut    time.Time  `json:"checkOut"`
	Status      string     `json:"status"`
	ExtraCharge float64    `json:"extraCharge"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   time.Time  `json:"updatedAt"`
	DeletedAt   *time.Time `json:"deletedAt"`
}
