package booking

import (
	"errors"
	"time"

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

	return s
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

func (s *Service) createBooking(createBookingDto *CreateBookingDto) (*models.Booking, error) {
	newBooking, err := utils.MapStruct(&models.Booking{}, createBookingDto)
	if err != nil {
		return nil, err
	}

	reply := s.queue.Publish(&mq.Message{
		Topic: events.CHECKINOUTCREATED,
		Data: &events.CreateCheckInOutDto{
			CheckIn:     newBooking.CheckIn,
			CheckOut:    newBooking.CheckOut,
			Status:      "",
			ExtraCharge: 0,
		},
	})

	select {
	case resp := <-reply:
		data := resp.(*response.ServiceResponse)
		if data.Error != nil {
			return nil, data.Error
		}

		newBooking.CheckInOutID = data.Data.(*models.CheckInOut).ID
		newBooking.CheckInOut = *data.Data.(*models.CheckInOut)
	case <-time.After(10 * time.Second):
		return nil, errors.New("timeout")
	}

	err = s.repo.create(newBooking)
	if err != nil {
		return nil, err
	}

	return newBooking, nil
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
