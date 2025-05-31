package room

import (
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"github.com/one-project-one-month/Hotel-Booking-Management-System-Go/pkg/models"
	"gorm.io/gorm"
)

type UpdateRoomIsFeaturedDto struct {
	ID         uuid.UUID
	IsFeatured bool
}

// Repository handles user data persistence operations.
type Repository struct {
	db *gorm.DB // Will be implemented in future releases
}

func newRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) findAll() ([]ResponseRoomDto, error) {
	var rooms []models.Room
	if err := r.db.Find(&rooms).Error; err != nil {
		return nil, err
	}

	response := make([]ResponseRoomDto, len(rooms))
	for i, room := range rooms {
		resRoom := ResponseRoomDto{}
		_ = copier.Copy(&resRoom, &room)
		response[i] = resRoom
	}
	return response, nil
}

func (r *Repository) findByID(id uuid.UUID) (*ResponseRoomDto, error) {
	var room models.Room
	if err := r.db.First(&room, id).Error; err != nil {
		return nil, err
	}
	resRoom := &ResponseRoomDto{}
	_ = copier.Copy(&resRoom, &room)

	return resRoom, nil
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
	err = MapRequestDtoToRoom(updatedRoom, &room)
	err = r.db.Save(&room).Error
	if err != nil {
		return nil, err
	}

	return &room, nil
}

func (r *Repository) updateRoomStatus(status string, id uuid.UUID) (uuid.UUID, error) {
	var room *models.Room
	err := r.db.First(&room, id).Error
	room.Status = status
	err = r.db.Save(&room).Error
	if err != nil {
		return uuid.Nil, err
	}

	return room.ID, nil
}

func (r *Repository) updateRoomIsFeatured(id uuid.UUID) (UpdateRoomIsFeaturedDto, error) {
	var room *models.Room
	err := r.db.First(&room, id).Error
	room.IsFeatured = !room.IsFeatured
	err = r.db.Save(&room).Error
	if err != nil {
		return UpdateRoomIsFeaturedDto{}, err
	}
	return UpdateRoomIsFeaturedDto{
		ID:         room.ID,
		IsFeatured: room.IsFeatured,
	}, nil
}

func (r *Repository) delete(id uuid.UUID) error {
	result := r.db.Delete(&models.Room{}, id)
	if err := result.Error; err != nil {
		return err
	}

	return nil
}

func (r *Repository) getRoomByGuestLimit(guests int) ([]ResponseRoomDto, error) {
	var rooms []models.Room
	if err := r.db.Where("guest_limit >= ?", guests).Find(&rooms).Error; err != nil {
		return nil, err
	}
	response := make([]ResponseRoomDto, len(rooms))
	for i, room := range rooms {
		resRoom := ResponseRoomDto{}
		_ = copier.Copy(&resRoom, &room)
		response[i] = resRoom
	}
	return response, nil

}
