package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BankAccount struct {
	ID            uuid.UUID      `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"id"`
	AccountNumber string         `gorm:"unique;not null" json:"account_number"`
	Pin           string         `gorm:"not null" json:"pin"`
	Amount        float64        `gorm:"not null" json:"amount"`
	CreatedAt     time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt     time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
