package invoice

import (
	"github.com/google/uuid"
	"github.com/one-project-one-month/Hotel-Booking-Management-System-Go/pkg/models"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(invoice *models.Invoice) error {
	return r.db.Create(invoice).Error
}

func (r *Repository) FindByID(id uuid.UUID) (*models.Invoice, error) {
	var invoice models.Invoice
	if err := r.db.First(&invoice, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &invoice, nil
}

func (r *Repository) FindAll() ([]models.Invoice, error) {
	var invoices []models.Invoice
	if err := r.db.Find(&invoices).Error; err != nil {
		return nil, err
	}
	return invoices, nil
}

func (r *Repository) Update(invoice *models.Invoice) error {
	return r.db.Save(invoice).Error
}

func (r *Repository) Delete(id uuid.UUID) error {
	return r.db.Delete(&models.Invoice{}, id).Error
}
