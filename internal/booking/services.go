package booking

import (
	"github.com/google/uuid"
	"github.com/one-project-one-month/Hotel-Booking-Management-System-Go/pkg/models"
	"github.com/one-project-one-month/Hotel-Booking-Management-System-Go/pkg/utils"
)

// Service
type Service struct {
	repo *Repository
}

func newService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) findAllBookings() (*[]ResponseBookingDto, error) {
	bookings, err := s.repo.findAll()
	if err != nil {
		return nil, err
	}

	return bookings, nil
}

func (s *Service) getBookingByID(id uuid.UUID) (*ResponseBookingDto, error) {
	booking, err := s.repo.findByID(id)
	if err != nil {
		return nil, err
	}

	return booking, nil
}

func (s *Service) createBooking(createBookingDto *CreateBookingDto) (*ResponseBookingDto, error) {
	newUser, err := utils.MapStruct(&models.Booking{}, createBookingDto)
	if err != nil {
		return nil, err
	}
	createdBooking, err := s.repo.create(newUser)
	if err != nil {
		return nil, err
	}

	return createdBooking, nil
}

func (s *Service) updateBooking(updateBookingDto *UpdateBookingDto, id uuid.UUID) (*ResponseBookingDto, error) {
	newBooking, err := utils.MapStruct(&models.Booking{}, updateBookingDto)
	booking, err := s.repo.update(newBooking, id)
	if err != nil {
		return nil, err
	}

	return booking, nil
}

func (s *Service) deleteBookingByID(id uuid.UUID) error {
	err := s.repo.delete(id)
	if err != nil {
		return err
	}

	return nil
}
