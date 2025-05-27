package models

import (
	"time"

	"github.com/google/uuid"
)

type CheckInOut struct {
	ID          uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"id"`
	CheckIn     time.Time `gorm:"not null;index" json:"check_in" validate:"required"`
	CheckOut    time.Time `gorm:"null;index" json:"check_out" validate:"gtfield=CheckIn"`
	Status      string    `gorm:"not null" json:"status" validate:"required"`
	ExtraCharge float64   `gorm:"not null" json:"extra_charge" validate:"required"`
	BookingID   uuid.UUID `gorm:"not null" json:"booking_id" validate:"required"`
	Booking     Booking   `gorm:"foreignKey:BookingID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT" json:"booking"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt   time.Time `gorm:"index" json:"deleted_at"`
}
