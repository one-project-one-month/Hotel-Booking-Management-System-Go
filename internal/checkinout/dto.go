package checkinout

import (
	"time"

	"github.com/google/uuid"
	"github.com/one-project-one-month/Hotel-Booking-Management-System-Go/internal/booking"
	"github.com/one-project-one-month/Hotel-Booking-Management-System-Go/pkg/models"
)

type CreateCheckInOutDto struct {
	CheckIn     time.Time `json:"checkIn"`
	CheckOut    time.Time `json:"checkOut"`
	Status      string    `json:"status" validate:"required"`
	ExtraCharge float64   `json:"extraCharge"`
	BookingID   uuid.UUID `json:"bookingId" validate:"required,uuid"`
}

type UpdateCheckInOutDto struct {
	CheckIn     *time.Time `json:"checkIn,omitempty" validate:"omitempty"`
	CheckOut    *time.Time `json:"checkOut,omitempty" validate:"omitempty,gtfield=CheckIn"`
	Status      *string    `json:"status,omitempty" validate:"omitempty"`
	ExtraCharge *float64   `json:"extraCharge,omitempty" validate:"omitempty"`
	BookingID   *uuid.UUID `json:"bookingId,omitempty" validate:"omitempty,uuid"`
}

type ResponseCheckInOutDto struct {
	ID          uuid.UUID                   `json:"id"`
	CheckIn     time.Time                   `json:"checkIn"`
	CheckOut    time.Time                   `json:"checkOut"`
	Status      string                      `json:"status"`
	ExtraCharge float64                     `json:"extraCharge"`
	BookingID   uuid.UUID                   `json:"bookingId"`
	Booking     *booking.ResponseBookingDto `json:"booking"`
	CreatedAt   time.Time                   `json:"createdAt"`
	UpdatedAt   time.Time                   `json:"updatedAt"`
	DeletedAt   *time.Time                  `json:"deletedAt"`
}

func NewResponseDtoFromModel(checkInOut *models.CheckInOut) ResponseCheckInOutDto {
	var deletedAt *time.Time
	if !checkInOut.DeletedAt.IsZero() {
		deletedAt = &checkInOut.DeletedAt
	}

	var bookingDto *booking.ResponseBookingDto
	if checkInOut.Booking.ID != uuid.Nil {
		bookingResp := booking.NewResponseDtoFromModel(&checkInOut.Booking)
		bookingDto = &bookingResp
	}

	return ResponseCheckInOutDto{
		ID:          checkInOut.ID,
		CheckIn:     checkInOut.CheckIn,
		CheckOut:    checkInOut.CheckOut,
		Status:      checkInOut.Status,
		ExtraCharge: checkInOut.ExtraCharge,
		BookingID:   checkInOut.BookingID,
		Booking:     bookingDto,
		CreatedAt:   checkInOut.CreatedAt,
		UpdatedAt:   checkInOut.UpdatedAt,
		DeletedAt:   deletedAt,
	}
}
