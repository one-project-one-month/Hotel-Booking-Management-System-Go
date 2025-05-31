package user

import (
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/one-project-one-month/Hotel-Booking-Management-System-Go/pkg/models"
)

type CreateUserDto struct {
	Name        string `json:"name" validate:"required,min=2,max=100"`
	Email       string `json:"email" validate:"required,email,max=255"`
	PhoneNumber string `json:"phoneNumber" validate:"required,e164"` // Added e164 validation
	Password    string `json:"password" validate:"required,min=8"`   // Added min length
	Role        string `json:"role" validate:"required,oneof=user admin"`
	ImageURL    string `json:"imageUrl,omitempty" validate:"omitempty,url"`
}

type UpdateUserDto struct {
	Name        *string `json:"name,omitempty" validate:"omitempty,min=2,max=100"`
	Email       *string `json:"email,omitempty" validate:"omitempty,email,max=255"`
	PhoneNumber *string `json:"phoneNumber,omitempty" validate:"omitempty,e164"`
	Password    *string `json:"password,omitempty" validate:"omitempty,min=8"`
	Role        *string `json:"role,omitempty" validate:"omitempty,oneof=user admin"`
	ImageURL    *string `json:"imageUrl,omitempty" validate:"omitempty,url"`
}

type ResponseUserDto struct {
	ID          uuid.UUID        `json:"id"`
	Name        string           `json:"name"`
	Email       string           `json:"email"`
	PhoneNumber string           `json:"phoneNumber"`
	Role        string           `json:"role"`
	ImageUrl    string           `json:"imageUrl"`
	Points      int              `json:"points"`
	CreatedAt   time.Time        `json:"createdAt"`
	UpdatedAt   time.Time        `json:"updatedAt"`
	DeletedAt   *time.Time       `json:"deletedAt"`
	Bookings    []models.Booking `json:"bookings,omitempty"`
}
func NewResponseDtoFromModel(user *models.User) ResponseUserDto {
	var deletedAt *time.Time
	if user.DeletedAt.Valid {
		deletedAt = &user.DeletedAt.Time
	}

	// Clear nested relationships from bookings
	bookings := make([]models.Booking, len(user.Bookings))
	for i, booking := range user.Bookings {
		bookings[i] = booking
		// Clear nested relationships to avoid circular references
		bookings[i].User = models.User{}
		bookings[i].Room = models.Room{}
	}

	return ResponseUserDto{
		ID:          user.ID,
		Name:        user.Name,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		Role:        strings.ToLower(user.Role.String()), // Ensure consistent case
		ImageUrl:    user.ImageURL,
		Points:      user.Points,
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
		DeletedAt:   deletedAt,
		Bookings:    bookings,
	}
}