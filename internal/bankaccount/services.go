package bankaccount

import (
	"context"
	"crypto/rand"
	"fmt"

	"github.com/one-project-one-month/Hotel-Booking-Management-System-Go/pkg/events"
	"github.com/one-project-one-month/Hotel-Booking-Management-System-Go/pkg/mq"
	"github.com/one-project-one-month/Hotel-Booking-Management-System-Go/pkg/response"
)

type Service interface {
	GetAll(ctx context.Context) response.ServiceResponse
}

type service struct {
	queue *mq.MQ
	repo  Repository
}

func NewService(repo Repository, queue *mq.MQ) Service {
	s := &service{repo: repo, queue: queue}
	queue.Subscribe(events.BANKACCOUNTFETCHED, func(data any) any {
		accountNumber := data.(events.FindByAccountNumberDto).AccountNumber
		return s.GetByAccountNumber(context.Background(), accountNumber)
	})
	return s
}

func generateAccountNumber() string {
	b := make([]byte, 8)
	rand.Read(b)

	accountNumber := ""
	for i := range 8 {
		accountNumber += fmt.Sprintf("%02d", b[i]%100)
	}

	return accountNumber
}

func (s *service) GetAll(ctx context.Context) response.ServiceResponse {
	accounts, err := s.repo.FindAll(ctx)
	if err != nil {
		return response.ServiceResponse{
			AppID:   "BankAccountService",
			Error:   response.ErrInternalServer,
			Message: "Failed to fetch bank accounts!",
		}
	}

	responses := make([]ResponseBankAccountDto, len(accounts))
	for i, account := range accounts {
		responses[i] = NewResponseDtoFromModel(&account)
	}

	return response.ServiceResponse{
		AppID:   "BankAccountService",
		Message: "Success!",
		Data:    responses,
	}
}

func (s *service) GetByAccountNumber(ctx context.Context, accountNumber string) *response.ServiceResponse {
	account, err := s.repo.FindByAccountNumber(ctx, accountNumber)
	if err != nil {
		return &response.ServiceResponse{
			AppID:   "BankAccountService",
			Error:   response.ErrInternalServer,
			Message: "Failed to fetch bank account!",
		}
	}

	return &response.ServiceResponse{
		AppID:   "BankAccountService",
		Message: "Success!",
		Data:    NewResponseDtoFromModel(account),
	}
}
