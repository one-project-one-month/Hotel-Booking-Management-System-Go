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
	UserID        uuid.UUID `gorm:"type:uuid;not null;index" json:"user_id" validate:"required"`
	RoomID        uuid.UUID `gorm:"type:uuid;not null;index" json:"room_id" validate:"required"`
	CheckIn       time.Time `gorm:"not null;index" json:"check_in" validate:"required"`
	CheckOut      time.Time `gorm:"null;index" json:"check_out" validate:"gtfield=CheckIn"`
	GuestCount    int       `gorm:"not null" json:"guests" validate:"gt=0"`
	DepositAmount float64   `gorm:"not null;default:0" json:"deposit_amount" validate:"gte=0"`

	TotalAmount float64        `gorm:"not null" json:"total_amount" validate:"gt=0"`
	Status      Status         `gorm:"type:varchar(20);default:pending;null" json:"status" validate:"oneof=pending approved"`
	CreatedAt   time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	// User          User           `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT" json:"user"`
	// Room          Room           `gorm:"foreignKey:RoomID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT" json:"room"`
	User User `json:"-" gorm:"foreignKey:UserID"`
	Room Room `json:"-" gorm:"foreignKey:RoomID"`

	CheckInOutID uuid.UUID  `json:"checkinout_id"`
	CheckInOut   CheckInOut `gorm:"ForeignKey:CheckInOutID;constraint:OnUpdate=CASCADE,OnDelete=SET NULL"`
}
