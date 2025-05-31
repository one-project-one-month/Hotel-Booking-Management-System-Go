package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRole int

const (
	RoleUser UserRole = iota + 1 // Starts at 1
	RoleAdmin
)

func (r UserRole) String() string {
	roles := [...]string{"", "user", "admin"}
	if r < RoleUser || r > RoleAdmin {
		return "unknown"
	}
	return roles[r]
}

type User struct {
	ID          uuid.UUID      `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"id"`
	Name        string         `gorm:"type:varchar(100);not null" json:"name" validate:"required,max=100"`
	Email       string         `gorm:"type:varchar(255);unique;not null;index" json:"email" validate:"required,email,max=255"`
	PhoneNumber string         `gorm:"type:varchar(20);unique;not null;index" json:"phone_number" validate:"required,max=20"`
	Password    string         `gorm:"type:varchar(255);not null" json:"-" validate:"required,min=8"`
	Role        UserRole       `gorm:"type:smallint;not null;default:1" json:"role" validate:"gte=0,lte=3"`
	ImageURL    string         `gorm:"type:text" json:"image_url"`
	Points      int            `gorm:"not null;default:0" json:"points" validate:"gte=0"`
	CreatedAt   time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	Bookings    []Booking      `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT" json:"bookings"`
	Coupons     []Coupon       `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT" json:"coupons"`
}
