package coupon

import "github.com/google/uuid"

type CreateCouponDto struct {
	Discounts  float64 `json:"discounts" validate:"min=1"`
	ExpiryDate string  `json:"expiry_date" validate:"required,datetime=2006-01-02T15:04:05Z"`
}

type FindListCouponDto struct {
	Limit   int    `json:"limit"`
	Page    int    `json:"page"`
	SortBy  string `json:"sort_by"`
	OrderBy string `json:"order_by"`
}

type UpdateCouponDataDto struct {
	UserID uuid.UUID `json:"user_id" validate:"required,uuid4"`
}

type UpdateCouponDto struct {
	Method string              `json:"method" validate:"required,oneof=claim activate"`
	Data   UpdateCouponDataDto `json:"data" validate:"required"`
}
