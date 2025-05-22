package room

import "github.com/google/uuid"

type RequestRoomDto struct {
	RoomNo     int         `json:"roomNo"      validate:"required"`
	Type       string      `json:"type"        validate:"required,oneof=Deluxe Standard"`
	Price      float64     `json:"price"       validate:"required,numeric,gt=0"`
	Status     string      `json:"status"      validate:"required,oneof=Available CheckedIn CheckOut"`
	IsFeatured bool        `json:"isFeatured"  validate:"omitempty,boolean"`
	Details    interface{} `json:"details"     validate:"omitempty,json"`
	ImgURL     []string    `json:"imgUrl" 	   validate:"omitempty"`
	GuestLimit int         `json:"guestLimit"  validate:"required,numeric,gt=0"`
}

type ResponseRoomDto struct {
	ID         uuid.UUID `json:"id"`
	RoomNo     int       `json:"roomNo"`
	Type       string    `json:"type"`
	Price      float64   `json:"price"`
	Status     string    `json:"status"`
	IsFeatured bool      `json:"isFeatured"`
	Details    string    `json:"details"`
	ImgURL     string    `json:"imgUrl"`
	GuestLimit int       `json:"guestLimit"`
}
