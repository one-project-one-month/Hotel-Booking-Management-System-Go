package checkinout

import (
	"context"

	"github.com/google/uuid"
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
	s := &service{repo: repo, queue: queue}

	s.queue.Subscribe(events.CHECKINOUTCREATED, func(data any) any {
		dto := data.(*events.CreateCheckInOutDto)

		result := s.Create(context.Background(), CreateCheckInOutDto{
			CheckIn:     dto.CheckIn,
			CheckOut:    dto.CheckOut,
			Status:      dto.Status,
			ExtraCharge: dto.ExtraCharge,
		})

		return &result
	})

	return s
}

func (s *service) Create(ctx context.Context, dto CreateCheckInOutDto) response.ServiceResponse {
	checkInOut := &models.CheckInOut{
		CheckIn:     dto.CheckIn,
		CheckOut:    dto.CheckOut,
		Status:      dto.Status,
		ExtraCharge: dto.ExtraCharge,
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
		Data:    checkInOut,
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
		Data:    checkInOut,
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

	responses := make([]*models.CheckInOut, len(checkInOuts))
	for i, checkInOut := range checkInOuts {
		responses[i] = &checkInOut
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

	if dto.Status != "" {
		checkInOut.Status = dto.Status
	}
	if dto.ExtraCharge != 0 {
		checkInOut.ExtraCharge = dto.ExtraCharge
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
		Data:    checkInOut,
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
