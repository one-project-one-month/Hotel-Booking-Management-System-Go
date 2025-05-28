package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Invoice struct {
	ID           uuid.UUID      `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"id"`
	CheckInOutID uuid.UUID      `gorm:"type:uuid;not null;index" json:"check_in_out_id" validate:"required"`
	CheckInOut   CheckInOut     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT" json:"checkInOut"`
	TotalAmount  float64        `gorm:"not null" json:"total_amount" validate:"gt=0"`
	CreatedAt    time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
