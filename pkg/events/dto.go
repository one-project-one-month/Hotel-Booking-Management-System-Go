package events

import (
	"time"

	"github.com/google/uuid"
)

type FindByUserIdDto struct {
	UserID string
}

type FindByIdDto struct {
	ID uuid.UUID
}

type FindByEmailDto struct {
	Email string
}

type FindByAccountNumberDto struct {
	AccountNumber string
}

type FindByPhoneNumberDto struct {
	PhoneNumber string
}

type UpdateBookingDto struct {
	ID            uuid.UUID
	CheckInOutID  uuid.UUID
	CheckIn       *time.Time
	CheckOut      *time.Time
	GuestCount    int
	Status        string
	DepositAmount float64
	TotalAmount   float64
}

type CreateCheckInOutDto struct {
	CheckIn     time.Time
	CheckOut    time.Time
	Status      string
	ExtraCharge float64
}
