package bankaccount

import (
	"fmt"
	"log"
	"time"

	"github.com/one-project-one-month/Hotel-Booking-Management-System-Go/pkg/models"
	"gorm.io/gorm"
)

func Seed(db *gorm.DB) error {
	var count int64
	if err := db.Model(&models.BankAccount{}).Count(&count).Error; err != nil {
		return fmt.Errorf("failed to count bank accounts: %w", err)
	}

	if count > 0 {
		fmt.Println("bank accounts table already has data, skipping seeding...")
		return nil
	}

	const batchSize = 30
	accounts := make([]models.BankAccount, batchSize)
	now := time.Now()

	pins := []string{"123456", "234567", "345678", "456789", "567890"}

	for i := range batchSize {
		seq := i + 1
		accountNumber := generateAccountNumber()

		accounts[i] = models.BankAccount{
			AccountNumber: accountNumber,
			Pin:           pins[i%len(pins)],
			Amount:        float64(seq * 3000),
			CreatedAt:     now,
			UpdatedAt:     now,
		}
	}

	result := db.Create(&accounts)
	if result.Error != nil {
		return fmt.Errorf("failed to seed bank accounts: %w", result.Error)
	}

	if result.RowsAffected != int64(batchSize) {
		return fmt.Errorf("expected to seed %d bank accounts, but only %d were inserted", batchSize, result.RowsAffected)
	}

	log.Printf("successfully seeded %d bank accounts\n", batchSize)
	return nil
}
