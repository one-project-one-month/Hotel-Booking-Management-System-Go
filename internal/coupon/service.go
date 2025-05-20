package coupon

import (
	"github.com/one-project-one-month/Hotel-Booking-Management-System-Go/pkg/models"
	"github.com/one-project-one-month/Hotel-Booking-Management-System-Go/pkg/response"
	"math/rand"
	"strings"
	"time"
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
			Error: err,
		}
	}

	return &response.ServiceResponse{
		AppID: "CouponService",
		Data:  nil,
		Error: nil,
	}
}
