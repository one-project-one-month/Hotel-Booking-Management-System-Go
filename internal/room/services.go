package room

import (
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"github.com/one-project-one-month/Hotel-Booking-Management-System-Go/pkg/events"
	"github.com/one-project-one-month/Hotel-Booking-Management-System-Go/pkg/models"
	"github.com/one-project-one-month/Hotel-Booking-Management-System-Go/pkg/mq"
	"github.com/one-project-one-month/Hotel-Booking-Management-System-Go/pkg/response"
)

// Service
type Service struct {
	queue *mq.MQ
	repo  *Repository
}

func newService(repo *Repository, queue *mq.MQ) *Service {
	s := &Service{repo: repo, queue: queue}

	s.queue.Subscribe(events.ROOMFINDBYID, func(data any) any {
		id := data.(*events.FindByIdDto).ID
		room, err := s.repo.findByID(id)
		if err != nil {
			return &response.ServiceResponse{
				AppID: "RoomService",
				Error: err,
			}
		}
		return &response.ServiceResponse{
			AppID: "RoomService",
			Data:  room,
		}
	})

	return s
}

func (s *Service) findAllRooms() ([]ResponseRoomDto, error) {
	rooms, err := s.repo.findAll()
	if err != nil {
		return nil, err
	}

	return rooms, nil
}

func (s *Service) getRoomByID(id uuid.UUID) (*ResponseRoomDto, error) {
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

func (s *Service) updateRoom(roomDto *RequestRoomDto, id uuid.UUID) (*ResponseRoomDto, error) {
	room, err := s.repo.update(roomDto, id)
	resRoom := &ResponseRoomDto{}
	if err != nil {
		return nil, err
	}
	_ = copier.Copy(&resRoom, room)
	return resRoom, nil
}

func (s *Service) updateRoomStatus(status string, id uuid.UUID) (uuid.UUID, error) {
	updatedRoomID, err := s.repo.updateRoomStatus(status, id)
	if err != nil {
		return uuid.Nil, err
	}
	return updatedRoomID, nil
}

func (s *Service) updateRoomIsFeatured(id uuid.UUID) (UpdateRoomIsFeaturedDto, error) {
	updateRoomIsFeaturedDto, err := s.repo.updateRoomIsFeatured(id)
	if err != nil {
		return UpdateRoomIsFeaturedDto{}, err
	}
	return updateRoomIsFeaturedDto, nil
}

func (s *Service) getRoomByGuestLimit(guestLimit int) ([]ResponseRoomDto, error) {
	rooms, err := s.repo.getRoomByGuestLimit(guestLimit)
	if err != nil {
		return nil, err
	}
	return rooms, nil
}

// Helper Functions
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
