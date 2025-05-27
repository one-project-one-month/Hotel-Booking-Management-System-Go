package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Gorm Model
type Room struct {
	ID         uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	RoomNo     int       `gorm:"not null;unique"`
	Type       string    `gorm:"not null"`
	Price      float64   `gorm:"not null"`
	Status     string    `gorm:"not null"`
	IsFeatured bool      `gorm:"default:false"`
	Details    string    `gorm:"type:text"`
	ImgURL     string    `gorm:"type:text"`
	GuestLimit int       `gorm:"not null"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt
	Bookings   []Booking `gorm:"foreignKey:RoomID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT" json:"bookings"`
}
