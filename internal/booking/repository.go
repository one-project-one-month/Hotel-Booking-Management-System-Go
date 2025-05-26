package booking

import (
	"github.com/google/uuid"
	"github.com/one-project-one-month/Hotel-Booking-Management-System-Go/pkg/models"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func newRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) findAll() (*[]ResponseBookingDto, error) {
	var bookings []models.Booking
	if err := r.db.Preload("User").Preload("Room").Find(&bookings).Error; err != nil {
		return nil, err
	}

	response := make([]ResponseBookingDto, len(bookings))
	for i, booking := range bookings {
		response[i] = NewResponseDtoFromModel(&booking)
	}
	return &response, nil
}

func (r *Repository) findByID(id uuid.UUID) (*ResponseBookingDto, error) {
	var booking models.Booking
	if err := r.db.Preload("User").Preload("Room").First(&booking, "id = ?", id).Error; err != nil {
		return nil, err
	}

	response := NewResponseDtoFromModel(&booking)
	return &response, nil
}

func (r *Repository) create(booking *models.Booking) (*ResponseBookingDto, error) {
	result := r.db.Create(&booking)
	if err := result.Error; err != nil {
		return nil, err
	}

	response := NewResponseDtoFromModel(booking)

	return &response, nil
}

func (r *Repository) update(booking *models.Booking, id uuid.UUID) (*ResponseBookingDto, error) {
	var existingBooking models.Booking
	if err := r.db.First(&existingBooking, id).Error; err != nil {
		return nil, err
	}

	if err := r.db.Model(&existingBooking).Updates(booking).Error; err != nil {
		return nil, err
	}

	response := NewResponseDtoFromModel(&existingBooking)
	return &response, nil
}

func (r *Repository) delete(id uuid.UUID) error {
	result := r.db.Delete(&models.Booking{}, id)
	if err := result.Error; err != nil {
		return err
	}
	return nil
}
