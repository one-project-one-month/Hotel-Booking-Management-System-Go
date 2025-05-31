package checkinout

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/one-project-one-month/Hotel-Booking-Management-System-Go/internal/booking"
	"github.com/one-project-one-month/Hotel-Booking-Management-System-Go/pkg/events"
	"github.com/one-project-one-month/Hotel-Booking-Management-System-Go/pkg/models"
	"github.com/one-project-one-month/Hotel-Booking-Management-System-Go/pkg/mq"
	"github.com/one-project-one-month/Hotel-Booking-Management-System-Go/pkg/response"
)

type Service interface {
	Create(ctx context.Context, dto CreateCheckInOutDto) response.ServiceResponse
	GetByID(ctx context.Context, id uuid.UUID) response.ServiceResponse
	GetAll(ctx context.Context) response.ServiceResponse
	Update(ctx context.Context, id uuid.UUID, dto UpdateCheckInOutDto) response.ServiceResponse
	Delete(ctx context.Context, id uuid.UUID) response.ServiceResponse
}

type service struct {
	queue *mq.MQ
	repo  Repository
}

func NewService(repo Repository, queue *mq.MQ) Service {
	return &service{repo: repo, queue: queue}
}

func (s *service) Create(ctx context.Context, dto CreateCheckInOutDto) response.ServiceResponse {
	reply := s.queue.Publish(&mq.Message{
		AppID: "CheckInOutService",
		Topic: events.BOOKINGFETCHED,
		Data:  events.FindByIdDto{ID: dto.BookingID},
	})

	var data any
	select {
	case data = <-reply:
		if data == nil {
			return response.ServiceResponse{
				AppID:   "CheckInOutService",
				Error:   response.ErrBadRequest,
				Message: "Invalid Booking ID!",
			}
		}
	case <-time.After(1 * time.Second):
		return response.ServiceResponse{
			AppID:   "CheckInOutService",
			Error:   response.ErrInternalServer,
			Message: "Timeout!",
		}
	}

	booking := data.(*booking.ResponseBookingDto)
	checkInOut := &models.CheckInOut{
		CheckIn:     booking.CheckIn,
		CheckOut:    booking.CheckOut,
		Status:      dto.Status,
		ExtraCharge: dto.ExtraCharge,
		BookingID:   dto.BookingID,
	}

	if err := s.repo.Create(ctx, checkInOut); err != nil {
		return response.ServiceResponse{
			AppID:   "CheckInOutService",
			Error:   response.ErrInternalServer,
			Message: "Create Check-in/out Failed!",
		}
	}

	return response.ServiceResponse{
		AppID:   "CheckInOutService",
		Message: "Success!",
		Data:    NewResponseDtoFromModel(checkInOut),
	}
}

func (s *service) GetByID(ctx context.Context, id uuid.UUID) response.ServiceResponse {
	checkInOut, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return response.ServiceResponse{
			AppID:   "CheckInOutService",
			Error:   response.ErrNotFound,
			Message: "Check-in/out record not found!",
		}
	}

	return response.ServiceResponse{
		AppID:   "CheckInOutService",
		Message: "Success!",
		Data:    NewResponseDtoFromModel(checkInOut),
	}
}

func (s *service) GetAll(ctx context.Context) response.ServiceResponse {
	checkInOuts, err := s.repo.FindAll(ctx)
	if err != nil {
		return response.ServiceResponse{
			AppID:   "CheckInOutService",
			Error:   response.ErrInternalServer,
			Message: "Failed to fetch check-in/out records!",
		}
	}

	responses := make([]ResponseCheckInOutDto, len(checkInOuts))
	for i, checkInOut := range checkInOuts {
		responses[i] = NewResponseDtoFromModel(&checkInOut)
	}

	return response.ServiceResponse{
		AppID:   "CheckInOutService",
		Message: "Success!",
		Data:    responses,
	}
}

func (s *service) Update(ctx context.Context, id uuid.UUID, dto UpdateCheckInOutDto) response.ServiceResponse {
	checkInOut, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return response.ServiceResponse{
			AppID:   "CheckInOutService",
			Error:   response.ErrNotFound,
			Message: "Check-in/out record not found!",
		}
	}

	if dto.CheckIn != nil {
		checkInOut.CheckIn = *dto.CheckIn
	}
	if dto.CheckOut != nil {
		checkInOut.CheckOut = *dto.CheckOut
	}
	if dto.Status != nil {
		checkInOut.Status = *dto.Status
	}
	if dto.ExtraCharge != nil {
		checkInOut.ExtraCharge = *dto.ExtraCharge
	}
	if dto.BookingID != nil {
		checkInOut.BookingID = *dto.BookingID
	}

	if err := s.repo.Update(ctx, checkInOut); err != nil {
		return response.ServiceResponse{
			AppID:   "CheckInOutService",
			Error:   response.ErrInternalServer,
			Message: "Update Check-in/out Failed!",
		}
	}

	return response.ServiceResponse{
		AppID:   "CheckInOutService",
		Message: "Success!",
		Data:    NewResponseDtoFromModel(checkInOut),
	}
}

func (s *service) Delete(ctx context.Context, id uuid.UUID) response.ServiceResponse {
	if err := s.repo.Delete(ctx, id); err != nil {
		return response.ServiceResponse{
			AppID:   "CheckInOutService",
			Error:   response.ErrInternalServer,
			Message: "Delete Check-in/out Failed!",
		}
	}

	return response.ServiceResponse{
		AppID:   "CheckInOutService",
		Message: "Success!",
	}
}
