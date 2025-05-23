package room

import (
	"time"

	"github.com/google/uuid"
	"github.com/one-project-one-month/Hotel-Booking-Management-System-Go/pkg/models"
)

type RequestRoomDto struct {
	RoomNo      string  `json:"roomNo"      validate:"required,alphanum,max=20"`
	Type        string  `json:"type"        validate:"required,oneof=Deluxe Standard"`
	Price       float64 `json:"price"       validate:"required,numeric,gt=0"`
	Status      string  `json:"status"      validate:"required,oneof=Available CheckedIn CheckOut"`
	IsFeatured  string  `json:"isFeatured"  validate:"omitempty,boolean"`
	Description string  `json:"description" validate:"max=500"`
	ImgURL      string  `json:"imgUrl"      validate:"url"`
	GuestLimit  int     `json:"guestLimit"  validate:"required,numeric,gt=0"`
}

type ResponseRoomDto struct {
	ID          uuid.UUID `json:"id"`
	RoomNo      string    `json:"roomNo"`
	Type        string    `json:"type"`
	Price       float64   `json:"price"`
	Status      string    `json:"status"`
	IsFeatured  bool      `json:"isFeatured"`
	Description string    `json:"description"`
	ImgURL      string    `json:"imgUrl"`
	GuestLimit  int       `json:"guestLimit"`
	DeletedAt   *time.Time
}

func NewResponseDtoFromModel(room *models.Room) ResponseRoomDto {
	var deletedAt *time.Time
	if room.DeletedAt.Valid {
		deletedAt = &room.DeletedAt.Time
	}

	return ResponseRoomDto{
		ID:          room.ID,
		RoomNo:      room.RoomNo,
		Type:        room.Type,
		Price:       room.Price,
		Status:      room.Status,
		IsFeatured:  room.IsFeatured,
		Description: room.Description,
		ImgURL:      room.ImgURL,
		GuestLimit:  room.GuestLimit,
		DeletedAt:   deletedAt,
	}
}
