package invoice

import (
	"github.com/google/uuid"
	"github.com/one-project-one-month/Hotel-Booking-Management-System-Go/pkg/models"
	"github.com/one-project-one-month/Hotel-Booking-Management-System-Go/pkg/response"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateInvoice(dto *CreateInvoiceDto) *response.ServiceResponse {
	invoice := &models.Invoice{
		ID:           uuid.New(),
		CheckInOutID: dto.CheckInOutID,
		TotalAmount:  dto.TotalAmount,
	}

	if err := s.repo.Create(invoice); err != nil {
		return &response.ServiceResponse{Error: err}
	}

	return &response.ServiceResponse{Data: invoice}
}

func (s *Service) GetInvoiceByID(id uuid.UUID) *response.ServiceResponse {
	invoice, err := s.repo.FindByID(id)
	if err != nil {
		return &response.ServiceResponse{Error: err}
	}

	return &response.ServiceResponse{Data: invoice}
}

func (s *Service) GetAllInvoices() *response.ServiceResponse {
	invoices, err := s.repo.FindAll()
	if err != nil {
		return &response.ServiceResponse{Error: err}
	}

	return &response.ServiceResponse{Data: invoices}
}

func (s *Service) UpdateInvoice(id uuid.UUID, dto *UpdateInvoiceDto) *response.ServiceResponse {
	invoice, err := s.repo.FindByID(id)
	if err != nil {
		return &response.ServiceResponse{Error: err}
	}

	invoice.TotalAmount = dto.TotalAmount

	if err := s.repo.Update(invoice); err != nil {
		return &response.ServiceResponse{Error: err}
	}

	return &response.ServiceResponse{Data: invoice}
}

func (s *Service) DeleteInvoice(id uuid.UUID) *response.ServiceResponse {
	if err := s.repo.Delete(id); err != nil {
		return &response.ServiceResponse{Error: err}
	}

	return &response.ServiceResponse{Data: true}
}
