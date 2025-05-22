package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Coupon struct {
	ID         uuid.UUID      `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"id"`
	Code       string         `gorm:"not null;unique" json:"code"`
	Discount   float64        `gorm:"not null" json:"discount"`
	IsActive   bool           `gorm:"default:false" json:"is_active"`
	IsClaimed  bool           `gorm:"default:false" json:"is_claimed"`
	ExpiryDate time.Time      `json:"expiry_date"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	UserID     uuid.UUID      `json:"user_id" json:"-"`
}
