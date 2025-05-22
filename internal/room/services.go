package room

import (
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"github.com/one-project-one-month/Hotel-Booking-Management-System-Go/pkg/models"
)

// Service
type Service struct {
	repo *Repository
}

func newService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) findAllRooms() ([]ResponseRoomDto, error) {
	rooms, err := s.repo.findAll()
	if err != nil {
		return nil, err
	}

	return rooms, nil
}

func (s *Service) getRoomByID(id uuid.UUID) (*models.Room, error) {
	room, err := s.repo.findByID(id)
	if err != nil {
		return nil, err
	}

	return room, nil
}

func (s *Service) createRoom(roomDto *RequestRoomDto) (uuid.UUID, error) {
	room := &models.Room{}
	err := MapRequestDtoToRoom(roomDto, room)
	if err != nil {
		return uuid.Nil, err
	}
	newRoomID, err := s.repo.create(room)
	if err != nil {
		return uuid.Nil, err
	}

	return newRoomID, nil
}

func (s *Service) deleteRoomByID(id uuid.UUID) error {
	err := s.repo.delete(id)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) updateRoom(roomDto *RequestRoomDto, id uuid.UUID) (*models.Room, error) {
	room, err := s.repo.update(roomDto, id)
	if err != nil {
		return nil, err
	}

	return room, nil
}

func mapStruct[T any](to *T, from any) (*T, error) {
	err := copier.Copy(&to, from)
	return to, err
}

func MapRequestDtoToRoom(dto *RequestRoomDto, room *models.Room) error {
	detailsJSON, err := json.Marshal(dto.Details)
	if err != nil {
		return fmt.Errorf("failed to marshal details: %w", err)
	}

	imgURLJSON, err := json.Marshal(dto.ImgURL)
	if err != nil {
		return fmt.Errorf("failed to marshal imgURL: %w", err)
	}

	room.RoomNo = dto.RoomNo
	room.Type = dto.Type
	room.Price = dto.Price
	room.Status = dto.Status
	room.IsFeatured = dto.IsFeatured
	room.Details = string(detailsJSON)
	room.ImgURL = string(imgURLJSON)
	room.GuestLimit = dto.GuestLimit
	return nil
}
