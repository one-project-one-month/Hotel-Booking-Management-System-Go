package room

import (
	"github.com/google/uuid"
	"github.com/one-project-one-month/Hotel-Booking-Management-System-Go/pkg/models"
	"gorm.io/gorm"
)

// Repository handles user data persistence operations.
type Repository struct {
	db *gorm.DB // Will be implemented in future releases
}

func newRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) findAll() ([]ResponseRoomDto, error) {
	var rooms []ResponseRoomDto
	if err := r.db.Model(&models.Room{}).Find(&rooms, &ResponseRoomDto{}).Error; err != nil {
		return nil, err
	}

	return rooms, nil
}

func (r *Repository) findByID(id uuid.UUID) (*models.Room, error) {
	var room models.Room
	if err := r.db.First(&room, id).Error; err != nil {
		return nil, err
	}

	return &room, nil
}

func (r *Repository) create(room *models.Room) (uuid.UUID, error) {
	result := r.db.Create(&room)
	if err := result.Error; err != nil {
		return uuid.Nil, err
	}

	return room.ID, nil
}

func (r *Repository) update(updatedRoom *RequestRoomDto, id uuid.UUID) (*models.Room, error) {
	var room models.Room
	err := r.db.First(&room, id).Error
	if err != nil {
		return nil, err
	}
	err = r.db.Model(&room).Updates(updatedRoom).Error
	if err != nil {
		return nil, err
	}

	return &room, nil
}

func (r *Repository) delete(id uuid.UUID) error {
	result := r.db.Delete(&models.Room{}, id)
	if err := result.Error; err != nil {
		return err
	}

	return nil
}
