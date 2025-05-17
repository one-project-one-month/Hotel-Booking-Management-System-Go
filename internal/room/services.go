package room

import (
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
	newRoom, err := mapStruct(&models.Room{}, roomDto)
	if err != nil {
		return uuid.Nil, err
	}
	newRoomID, err := s.repo.create(newRoom)
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
