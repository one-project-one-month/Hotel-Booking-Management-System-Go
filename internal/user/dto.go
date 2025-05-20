package user

import (
	"github.com/google/uuid"
	"github.com/one-project-one-month/Hotel-Booking-Management-System-Go/pkg/models"
	"strings"
	"time"
)

type CreateUserDto struct {
	Name        string `json:"name" validate:"required,min=2,max=100"`
	Email       string `json:"email" validate:"required,email,max=255"`
	PhoneNumber string `json:"phoneNumber" validate:"required,e164"` // Added e164 validation
	Password    string `json:"password" validate:"required,min=8"`   // Added min length
	Role        string `json:"role" validate:"required,oneof=user admin"`
	ImgURL      string `json:"imgUrl,omitempty" validate:"omitempty,url"`
}

type UpdateUserDto struct {
	Name        *string `json:"name,omitempty" validate:"omitempty,min=2,max=100"`
	Email       *string `json:"email,omitempty" validate:"omitempty,email,max=255"`
	PhoneNumber *string `json:"phoneNumber,omitempty" validate:"omitempty,e164"`
	Password    *string `json:"password,omitempty" validate:"omitempty,min=8"`
	Role        *string `json:"role,omitempty" validate:"omitempty,oneof=user admin"`
	ImgURL      *string `json:"imgUrl,omitempty" validate:"omitempty,url"`
}

type ResponseUserDto struct {
	ID          uuid.UUID  `json:"id"`
	Name        string     `json:"name"`
	Email       string     `json:"email"`
	PhoneNumber string     `json:"phoneNumber"`
	Role        string     `json:"role"`
	ImgURL      string     `json:"imgUrl,omitempty"`
	Points      int        `json:"points"`
	Amount      int        `json:"amount"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   time.Time  `json:"updatedAt"`
	DeletedAt   *time.Time `json:"deletedAt,omitempty"`
}

func NewResponseDtoFromModel(user *models.User) ResponseUserDto {
	var deletedAt *time.Time
	if user.DeletedAt.Valid {
		deletedAt = &user.DeletedAt.Time
	}

	return ResponseUserDto{
		ID:          user.ID,
		Name:        user.Name,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		Role:        strings.ToLower(user.Role.String()), // Ensure consistent case
		ImgURL:      user.ImgURL,
		Points:      user.Points,
		Amount:      user.Amount,
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
		DeletedAt:   deletedAt,
	}
}
