package booking

import (
	"fmt"
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

func (s *Service) findAllBookings() ([]*ResponseBookingDto, error) {
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

func (s *Service) createBooking(createBookingDto *CreateBookingDto) *response.ServiceResponse {
	newBooking, _ := utils.MapStruct(&models.Booking{}, createBookingDto)
	userReply := s.queue.Publish(&mq.Message{
		Topic: events.USERFINDBYID,
		Data: &events.FindByIdDto{
			ID: newBooking.UserID,
		},
	})

	select {
	case resp := <-userReply:
		data := resp.(*response.ServiceResponse)
		if data.Data == nil {
			return &response.ServiceResponse{
				AppID:   "BookingService",
				Error:   response.ErrNotFound,
				Message: fmt.Sprintf("user with id %s not found", newBooking.UserID.String()),
			}
		}
	case <-time.Tick(2 * time.Second):
		return &response.ServiceResponse{
			AppID:   "BookingService",
			Error:   response.ErrTimeout,
			Message: "timeout",
		}
	}

	roomReply := s.queue.Publish(&mq.Message{
		Topic: events.ROOMFINDBYID,
		Data: &events.FindByIdDto{
			ID: newBooking.RoomID,
		},
	})

	select {
	case resp := <-roomReply:
		data := resp.(*response.ServiceResponse)
		if data.Data == nil {
			return &response.ServiceResponse{
				AppID:   "BookingService",
				Error:   response.ErrNotFound,
				Message: fmt.Sprintf("room with id %s not found", newBooking.RoomID.String()),
			}
		}
	case <-time.Tick(2 * time.Second):
		return &response.ServiceResponse{
			AppID:   "BookingService",
			Error:   response.ErrTimeout,
			Message: "timeout",
		}
	}

	if newBooking.Status == "" {
		newBooking.Status = "pending"
	}

	reply := s.queue.Publish(&mq.Message{
		Topic: events.CHECKINOUTCREATED,
		Data: &events.CreateCheckInOutDto{
			CheckIn:     newBooking.CheckIn,
			CheckOut:    newBooking.CheckOut,
			Status:      string(newBooking.Status),
			ExtraCharge: 0,
		},
	})

	select {
	case resp := <-reply:
		data := resp.(*response.ServiceResponse)
		if data.Error != nil {
			return &response.ServiceResponse{
				AppID:   "BookingService",
				Error:   data.Error,
				Message: data.Message,
			}
		}

		newBooking.CheckInOutID = data.Data.(*models.CheckInOut).ID
		newBooking.CheckInOut = *data.Data.(*models.CheckInOut)
	case <-time.After(2 * time.Second):
		return &response.ServiceResponse{
			AppID:   "BookingService",
			Error:   response.ErrTimeout,
			Message: "timeout",
		}
	}

	err := s.repo.create(newBooking)
	if err != nil {
		return &response.ServiceResponse{
			AppID:   "BookingService",
			Error:   err,
			Message: "failed to create booking",
		}
	}

	return &response.ServiceResponse{
		AppID:   "BookingService",
		Data:    newBooking,
		Message: "booking created successfully",
	}
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
