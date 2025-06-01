package user

import (
	"github.com/google/uuid"
	"github.com/one-project-one-month/Hotel-Booking-Management-System-Go/pkg/models"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func newRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) findAll() ([]ResponseUserDto, error) {
	var users []models.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}

	response := make([]ResponseUserDto, len(users))
	for i, user := range users {
		response[i] = NewResponseDtoFromModel(&user)
	}
	return response, nil
}

func (r *Repository) findByID(id uuid.UUID) (*ResponseUserDto, error) {
	var user models.User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}

	response := NewResponseDtoFromModel(&user)
	return &response, nil
}

func (r *Repository) findByEmail(email string) (*models.User, error) {
	var user models.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *Repository) findByPhoneNumber(phoneNumber string) (*models.User, error) {
	var user models.User
	if err := r.db.Where("phone_number = ?", phoneNumber).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *Repository) create(user *models.User) error {
	result := r.db.Create(&user)
	if err := result.Error; err != nil {
		return err
	}

	return nil
}

func (r *Repository) update(user *models.User, id uuid.UUID) (*ResponseUserDto, error) {
	var existingUser models.User
	if err := r.db.First(&existingUser, id).Error; err != nil {
		return nil, err
	}

	if err := r.db.Model(&existingUser).Updates(user).Error; err != nil {
		return nil, err
	}

	response := NewResponseDtoFromModel(&existingUser)
	return &response, nil
}

func (r *Repository) delete(id uuid.UUID) error {
	result := r.db.Delete(&models.User{}, id)
	if err := result.Error; err != nil {
		return err
	}
	return nil
}
