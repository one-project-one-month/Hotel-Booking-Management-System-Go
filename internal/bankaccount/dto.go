package bankaccount

import (
	"time"

	"github.com/google/uuid"
	"github.com/one-project-one-month/Hotel-Booking-Management-System-Go/pkg/models"
)

type CreateBankAccountDto struct {
	AccountNumber string  `json:"accountNumber" validate:"required,len=16"`
	Pin           string  `json:"pin" validate:"required,len=6"`
	Amount        float64 `json:"amount" validate:"required,gte=0"`
}

type UpdateBankAccountDto struct {
	AccountNumber *string  `json:"accountNumber,omitempty" validate:"omitempty,len=16"`
	Pin           *string  `json:"pin,omitempty" validate:"omitempty,len=6"`
	Amount        *float64 `json:"amount,omitempty" validate:"omitempty,gte=0"`
}

type ResponseBankAccountDto struct {
	ID            uuid.UUID  `json:"id"`
	AccountNumber string     `json:"accountNumber"`
	Pin           string     `json:"pin"`
	Amount        float64    `json:"amount"`
	CreatedAt     time.Time  `json:"createdAt"`
	UpdatedAt     time.Time  `json:"updatedAt"`
	DeletedAt     *time.Time `json:"deletedAt,omitempty"`
}

func NewResponseDtoFromModel(account *models.BankAccount) ResponseBankAccountDto {
	var deletedAt *time.Time
	if account.DeletedAt.Valid {
		deletedAt = &account.DeletedAt.Time
	}

	return ResponseBankAccountDto{
		ID:            account.ID,
		AccountNumber: account.AccountNumber,
		Amount:        account.Amount,
		Pin:           account.Pin,
		CreatedAt:     account.CreatedAt,
		UpdatedAt:     account.UpdatedAt,
		DeletedAt:     deletedAt,
	}
}
