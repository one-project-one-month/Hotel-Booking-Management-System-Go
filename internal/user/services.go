package user

import (
	"github.com/google/uuid"
	"github.com/one-project-one-month/Hotel-Booking-Management-System-Go/pkg/events"
	"github.com/one-project-one-month/Hotel-Booking-Management-System-Go/pkg/models"
	"github.com/one-project-one-month/Hotel-Booking-Management-System-Go/pkg/mq"
	"github.com/one-project-one-month/Hotel-Booking-Management-System-Go/pkg/response"
	"github.com/one-project-one-month/Hotel-Booking-Management-System-Go/pkg/utils"
)

// Service
type Service struct {
	queue *mq.MQ
	repo  *Repository
}

func newService(repo *Repository, queue *mq.MQ) *Service {
	s := &Service{repo: repo, queue: queue}

	s.queue.Subscribe(events.USERFINDBYID, func(data any) any {
		dto := data.(*events.FindByIdDto)
		user, err := s.getUserByID(dto.ID)
		if err != nil {
			return &response.ServiceResponse{
				AppID: "UserService",
				Error: err,
			}
		}

		return &response.ServiceResponse{
			AppID: "UserService",
			Data:  user,
		}
	})

	return s
}

func (s *Service) findAllUsers() ([]ResponseUserDto, error) {
	users, err := s.repo.findAll()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *Service) getUserByID(id uuid.UUID) (*ResponseUserDto, error) {
	user, err := s.repo.findByID(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Service) createUser(userDto *CreateUserDto) (*ResponseUserDto, error) {
	newUser, err := utils.MapStruct(&models.User{}, userDto)
	if err != nil {
		return nil, err
	}
	createdUser, err := s.repo.create(newUser)
	if err != nil {
		return nil, err
	}

	return createdUser, nil
}

func (s *Service) updateUser(userDto *UpdateUserDto, id uuid.UUID) (*ResponseUserDto, error) {
	newUser, err := utils.MapStruct(&models.User{}, userDto)
	user, err := s.repo.update(newUser, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Service) deleteUserByID(id uuid.UUID) error {
	err := s.repo.delete(id)
	if err != nil {
		return err
	}

	return nil
}
