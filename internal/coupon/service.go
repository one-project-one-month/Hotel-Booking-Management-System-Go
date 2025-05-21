package coupon

import (
	"math/rand"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/one-project-one-month/Hotel-Booking-Management-System-Go/pkg/models"
	"github.com/one-project-one-month/Hotel-Booking-Management-System-Go/pkg/response"
)

type Service struct {
	repo *Repository
}

func (s *Service) generateCode() string {
	var code strings.Builder
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	for i := 0; i < 6; i++ {
		code.WriteByte(charset[rand.Intn(len(charset))])
	}

	return code.String()
}

func (s *Service) create(coupon *CreateCouponDto) *response.ServiceResponse {
	var model models.Coupon
	model.Discount = coupon.Discounts
	model.Code = s.generateCode()
	if expiryDate, err := time.Parse(time.DateTime, coupon.ExpiryDate); err == nil {
		model.ExpiryDate = expiryDate
	}

	if err := s.repo.create(&model); err != nil {
		return &response.ServiceResponse{
			AppID: "CouponService",
			Data:  nil,
			Error: response.ErrInternalServer,
		}
	}

	return &response.ServiceResponse{
		AppID: "CouponService",
		Data:  nil,
		Error: nil,
	}
}

func (s *Service) findList(query *FindListCouponDto) *response.ServiceResponse {
	order := query.SortBy + " " + query.OrderBy
	offset := (query.Page - 1) * query.Limit
	coupons, err := s.repo.findList(order, offset, query.Limit)
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
		couponModel.IsClaimed = true
	}

	if coupon.Method == "activate" {
		couponModel.IsActive = true
	}

	if err := s.repo.update(id, couponModel); err != nil {
		return &response.ServiceResponse{
			AppID: "CouponService",
			Error: response.ErrInternalServer,
		}
	}

	return &response.ServiceResponse{
		AppID: "CouponService",
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
