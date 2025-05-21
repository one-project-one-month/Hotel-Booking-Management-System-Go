package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type UserRole int

const (
	RoleUser UserRole = iota + 1 // Starts at 1
	RoleAdmin
)

func (r UserRole) String() string {
	roles := [...]string{"", "user", "admin"} // Add empty string for index 0
	if r < RoleUser || r > RoleAdmin {
		return "unknown"
	}
	return roles[r]
}

type User struct {
	ID          uuid.UUID      `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Name        string         `gorm:"not null;type:varchar(100)"`
	Email       string         `gorm:"unique;not null;type:varchar(255)"`
	PhoneNumber string         `gorm:"unique;not null;type:varchar(20)"`
	Password    string         `gorm:"not null;type:varchar(255)"`
	Role        UserRole       `gorm:"not null;type:smallint"`
	ImgURL      string         `gorm:"type:text"`
	Points      int            `gorm:"not null;default:0"`
	Amount      int            `gorm:"not null;default:0"`
	CreatedAt   time.Time      `gorm:"autoCreateTime"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	Coupons     []Coupon       `gorm:"foreignKey:UserID"`
}
