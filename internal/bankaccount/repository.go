package bankaccount

import (
	"context"

	"github.com/one-project-one-month/Hotel-Booking-Management-System-Go/pkg/models"
	"gorm.io/gorm"
)

type Repository interface {
	FindAll(ctx context.Context) ([]models.BankAccount, error)
	FindByAccountNumber(ctx context.Context, accountNumber string) (*models.BankAccount, error)
}

type repository struct {
	db *gorm.DB
}

func (r *repository) FindAll(ctx context.Context) ([]models.BankAccount, error) {
	var accounts []models.BankAccount
	err := r.db.WithContext(ctx).Find(&accounts).Error
	if err != nil {
		return nil, err
	}
	return accounts, nil
}

func (r *repository) FindByAccountNumber(ctx context.Context, accountNumber string) (*models.BankAccount, error) {
	var account models.BankAccount
	err := r.db.WithContext(ctx).First(&account, "account_number = ?", accountNumber).Error
	if err != nil {
		return nil, err
	}
	return &account, nil
}
