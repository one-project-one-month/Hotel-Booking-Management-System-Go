package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Status string

const (
	Pending  Status = "pending"
	Approved Status = "approved"
)

type Booking struct {
	ID            uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"id"`
	UserID        uuid.UUID `gorm:"type:uuid;not null;index" json:"userId" validate:"required"`
	RoomID        uuid.UUID `gorm:"type:uuid;not null;index" json:"roomId" validate:"required"`
	CheckIn       time.Time `gorm:"not null;index" json:"checkIn" validate:"required"`
	CheckOut      time.Time `gorm:"null;index" json:"checkOut" validate:"gtfield=CheckIn"`
	GuestCount    int       `gorm:"not null" json:"guestCount" validate:"gt=0"`
	DepositAmount float64   `gorm:"not null;default:0" json:"depositAmount" validate:"gte=0"`

	TotalAmount float64        `gorm:"not null" json:"totalAmount" validate:"gt=0"`
	Status      Status         `gorm:"type:varchar(20);default='pending';null" json:"status" validate:"oneof=pending approved"`
	CreatedAt   time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deletedAt"`
	// User          User           `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT" json:"user"`
	// Room          Room           `gorm:"foreignKey:RoomID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT" json:"room"`
	User User `json:"-" gorm:"foreignKey:UserID"`
	Room Room `json:"-" gorm:"foreignKey:RoomID"`
}
