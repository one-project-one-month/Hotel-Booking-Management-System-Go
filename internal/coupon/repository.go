package coupon

import (
	"github.com/google/uuid"
	"github.com/one-project-one-month/Hotel-Booking-Management-System-Go/pkg/models"
	"gorm.io/gorm"
)

type Repository struct {
	database *gorm.DB
}

func (r *Repository) create(coupon *models.Coupon) error {
	return r.database.Create(coupon).Error
}

func (r *Repository) findList(order string, offset int, limit int) ([]models.Coupon, error) {
	var coupons []models.Coupon
	if err := r.database.Order(order).Limit(limit).Offset(offset).Find(&coupons).Error; err != nil {
		return nil, err
	}
	return coupons, nil
}

func (r *Repository) findByID(id string) (*models.Coupon, error) {
	var coupon models.Coupon

	if err := r.database.Where("id = ?", id).First(&coupon).Error; err != nil {
		return nil, err
	}

	return &coupon, nil
}

func (r *Repository) findByCode(code string) (*models.Coupon, error) {
	return nil, nil
}

func (r *Repository) findByUserID(userId string) ([]models.Coupon, error) {
	var coupons []models.Coupon
	if err := r.database.Where("user_id = ?", userId).Find(&coupons).Error; err != nil {
		return nil, err
	}

	return coupons, nil
}

func (r *Repository) update(id string, coupon *models.Coupon) error {
	return r.database.Model(coupon).Where("id = ?", id).Updates(coupon).Error
}

func (r *Repository) delete(id uuid.UUID) error {
	return r.database.Delete(&models.Coupon{}, id).Error
}
