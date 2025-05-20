package coupon

import (
	"github.com/one-project-one-month/Hotel-Booking-Management-System-Go/pkg/models"
	"gorm.io/gorm"
)

type Repository struct {
	database *gorm.DB
}

func (r *Repository) create(coupon *models.Coupon) error {
	return r.database.Create(coupon).Error
}

func (r *Repository) findByCode(code string) (*models.Coupon, error) {
	return nil, nil
}
