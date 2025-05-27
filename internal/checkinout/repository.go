package checkinout

import (
	"context"

	"github.com/google/uuid"
	"github.com/one-project-one-month/Hotel-Booking-Management-System-Go/pkg/models"
	"gorm.io/gorm"
)

type Repository interface {
	Create(ctx context.Context, checkInOut *models.CheckInOut) error
	FindByID(ctx context.Context, id uuid.UUID) (*models.CheckInOut, error)
	FindAll(ctx context.Context) ([]models.CheckInOut, error)
	Update(ctx context.Context, checkInOut *models.CheckInOut) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(ctx context.Context, checkInOut *models.CheckInOut) error {
	return r.db.WithContext(ctx).Create(checkInOut).Error
}

func (r *repository) FindByID(ctx context.Context, id uuid.UUID) (*models.CheckInOut, error) {
	var checkInOut models.CheckInOut
	err := r.db.WithContext(ctx).Preload("Booking").First(&checkInOut, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &checkInOut, nil
}

func (r *repository) FindAll(ctx context.Context) ([]models.CheckInOut, error) {
	var checkInOuts []models.CheckInOut
	err := r.db.WithContext(ctx).Preload("Booking").Find(&checkInOuts).Error
	if err != nil {
		return nil, err
	}
	return checkInOuts, nil
}

func (r *repository) Update(ctx context.Context, checkInOut *models.CheckInOut) error {
	return r.db.WithContext(ctx).Save(checkInOut).Error
}

func (r *repository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&models.CheckInOut{}, "id = ?", id).Error
}
