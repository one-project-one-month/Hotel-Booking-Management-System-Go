package coupon

type CreateCouponDto struct {
	Discounts  float64 `json:"discounts" validate:"min=1"`
	ExpiryDate string  `json:"expiry_date" validate:"required,datetime=2006-01-02T15:04:05Z"`
}
