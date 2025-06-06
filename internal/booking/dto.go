package booking

import (
	"time"

	"github.com/google/uuid"
	"github.com/one-project-one-month/Hotel-Booking-Management-System-Go/internal/room"
	"github.com/one-project-one-month/Hotel-Booking-Management-System-Go/internal/user"
	"github.com/one-project-one-month/Hotel-Booking-Management-System-Go/pkg/models"
)

type CreateBookingDto struct {
	UserID        uuid.UUID `json:"userId" validate:"required,uuid"`
	RoomID        uuid.UUID `json:"roomId" validate:"required,uuid"`
	CheckInOutID  uuid.UUID `json:"checkInOutId"`
	CheckIn       time.Time `json:"checkIn" validate:"required"`
	CheckOut      time.Time `json:"checkOut" validate:"omitempty,gtfield=CheckIn"`
	GuestCount    int       `json:"guestCount" validate:"required,gt=0"`
	DepositAmount float64   `json:"depositAmount" validate:"gte=0"`
	TotalAmount   float64   `json:"totalAmount" validate:"required,gt=0"`
}

type UpdateBookingDto struct {
	UserID        *uuid.UUID `json:"userId,omitempty" validate:"omitempty,uuid"`
	RoomID        *uuid.UUID `json:"roomId,omitempty" validate:"omitempty,uuid"`
	CheckInOutID  uuid.UUID  `json:"checkInOutId,omitempty" validate:"omitempty,uuid"`
	CheckIn       *time.Time `json:"checkIn,omitempty" validate:"omitempty"`
	CheckOut      *time.Time `json:"checkOut,omitempty" validate:"omitempty,gtfield=CheckIn"`
	GuestCount    *int       `json:"guestCount,omitempty" validate:"omitempty,gt=0"`
	Status        *string    `json:"status,omitempty" validate:"omitempty,oneof=pending approved"`
	DepositAmount *float64   `json:"depositAmount,omitempty" validate:"omitempty,gte=0"`
	TotalAmount   *float64   `json:"totalAmount,omitempty" validate:"omitempty,gt=0"`
}

type ResponseBookingDto struct {
	ID            uuid.UUID `json:"id"`
	UserID        uuid.UUID `json:"userId"`
	RoomID        uuid.UUID `json:"roomId"`
	CheckInOutID  uuid.UUID `json:"checkInOutId"`
	CheckIn       time.Time `json:"checkIn"`
	CheckOut      time.Time `json:"checkOut"`
	GuestCount    int       `json:"guestCount"`
	DepositAmount float64   `json:"depositAmount"`
	TotalAmount   float64   `json:"totalAmount"`
	Status        string    `json:"status"`
	CreatedAt     time.Time `json:"createdAt"`

	UpdatedAt  time.Time             `json:"updatedAt"`
	DeletedAt  *time.Time            `json:"deletedAt"`
	User       *user.ResponseUserDto `json:"user"`
	Room       *room.ResponseRoomDto `json:"room"`
	CheckInOut *models.CheckInOut    `json:"checkInOut"`
}

func NewResponseDtoFromModel(booking *models.Booking) *ResponseBookingDto {
	var deletedAt *time.Time
	if booking.DeletedAt.Valid {
		deletedAt = &booking.DeletedAt.Time
	}

	var userDto *user.ResponseUserDto
	if booking.User.ID != uuid.Nil {
		userResp := user.NewResponseDtoFromModel(&booking.User)
		userDto = &userResp
	}

	var roomDto *room.ResponseRoomDto
	if booking.Room.ID != uuid.Nil {
		roomResp := room.NewResponseDtoFromModel(&booking.Room)
		roomDto = &roomResp
	}

	var checkInOutDto *models.CheckInOut
	if booking.CheckInOutID != uuid.Nil {
		checkInOutDto = &booking.CheckInOut
	}

	return &ResponseBookingDto{
		ID:            booking.ID,
		UserID:        booking.UserID,
		RoomID:        booking.RoomID,
		CheckInOutID:  booking.CheckInOutID,
		CheckIn:       booking.CheckIn,
		CheckOut:      booking.CheckOut,
		GuestCount:    booking.GuestCount,
		DepositAmount: booking.DepositAmount,
		TotalAmount:   booking.TotalAmount,
		Status:        string(booking.Status),
		CreatedAt:     booking.CreatedAt,
		UpdatedAt:     booking.UpdatedAt,
		DeletedAt:     deletedAt,
		User:          userDto,
		Room:          roomDto,
		CheckInOut:    checkInOutDto,
	}
}
