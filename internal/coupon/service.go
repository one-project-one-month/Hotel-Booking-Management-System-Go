package coupon

import (
	"math/rand"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/one-project-one-month/Hotel-Booking-Management-System-Go/pkg/events"
	"github.com/one-project-one-month/Hotel-Booking-Management-System-Go/pkg/models"
	"github.com/one-project-one-month/Hotel-Booking-Management-System-Go/pkg/mq"
	"github.com/one-project-one-month/Hotel-Booking-Management-System-Go/pkg/response"
)

type Service struct {
	queue *mq.MQ
	repo  *Repository
}

func (s *Service) generateCode() string {
	var code strings.Builder
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	for range 6 {
		code.WriteByte(charset[rand.Intn(len(charset))])
	}

	return code.String()
}

func (s *Service) create(coupon *CreateCouponDto) *response.ServiceResponse {
	userId, _ := uuid.Parse(coupon.UserID)
	reply := s.queue.Publish(&mq.Message{
		AppID: "CouponService",
		Topic: events.USERFINDBYID,
		Data: &events.FindByIdDto{
			ID: userId,
		},
	})

	select {
	case resp := <-reply:
		if data := resp.(*response.ServiceResponse); data.Error != nil {
			return &response.ServiceResponse{
				AppID:   "CouponService",
				Error:   response.ErrNotFound,
				Message: "User id is not found",
			}
		}
	case <-time.Tick(2 * time.Second):
		return &response.ServiceResponse{
			AppID:   "CouponService",
			Error:   response.ErrInternalServer,
			Message: "User service error",
		}
	}

	var model models.Coupon
	model.Discount = coupon.Discount
	model.UserID = userId
	model.Code = s.generateCode()
	if expiryDate, err := time.Parse(time.RFC3339, coupon.ExpiryDate); err == nil {
		model.ExpiryDate = expiryDate
	}

	if err := s.repo.create(&model); err != nil {
		return &response.ServiceResponse{
			AppID: "CouponService",
			Error: response.ErrInternalServer,
		}
	}

	return &response.ServiceResponse{
		AppID: "CouponService",
		Data:  model,
	}
}

func (s *Service) findList() *response.ServiceResponse {
	coupons, err := s.repo.findList()
	if err != nil {
		return &response.ServiceResponse{
			AppID: "CouponService",
			Data:  nil,
			Error: response.ErrInternalServer,
		}
	}

	return &response.ServiceResponse{
		AppID: "CouponService",
		Data:  coupons,
		Error: nil,
	}
}

func (s *Service) findByID(id string) *response.ServiceResponse {
	coupon, err := s.repo.findByID(id)
	if err != nil {
		return &response.ServiceResponse{
			AppID: "CouponService",
			Data:  nil,
			Error: response.ErrNotFound,
		}
	}

	return &response.ServiceResponse{
		AppID: "CouponService",
		Data:  coupon,
		Error: nil,
	}
}

func (s *Service) findByUserID(userId string) *response.ServiceResponse {
	coupons, err := s.repo.findByUserID(userId)
	if err != nil {
		return &response.ServiceResponse{
			AppID: "CouponService",
			Error: response.ErrNotFound,
		}
	}

	return &response.ServiceResponse{
		AppID: "CouponService",
		Data:  coupons,
	}
}

func (s *Service) update(id string, coupon *UpdateCouponDto) *response.ServiceResponse {
	couponModel, err := s.repo.findByID(id)
	if err != nil {
		return &response.ServiceResponse{
			AppID: "CouponService",
			Data:  nil,
			Error: response.ErrNotFound,
		}
	}

	if coupon.Method == "claim" {
		if couponModel.IsClaimed {
			return &response.ServiceResponse{
				AppID:   "CouponService",
				Error:   response.ErrBadRequest,
				Message: "Coupon is already claimed",
			}
		}
		couponModel.IsClaimed = true
	}

	if coupon.Method == "activate" {
		if couponModel.IsActive {
			return &response.ServiceResponse{
				AppID:   "CouponService",
				Error:   response.ErrBadRequest,
				Message: "Coupon is already activated",
			}
		}
		couponModel.IsActive = true
	}

	if err := s.repo.update(id, couponModel); err != nil {
		return &response.ServiceResponse{
			AppID: "CouponService",
			Error: response.ErrInternalServer,
		}
	}

	return &response.ServiceResponse{
		AppID:   "CouponService",
		Message: "Coupon updated successfully",
	}
}

func (s *Service) delete(id uuid.UUID) *response.ServiceResponse {
	if err := s.repo.delete(id); err != nil {
		return &response.ServiceResponse{
			AppID: "CouponService",
			Error: response.ErrInternalServer,
		}
	}

	return &response.ServiceResponse{
		AppID: "CouponService",
	}
}
