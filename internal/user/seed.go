package user

import (
	"fmt"
	"time"

	"github.com/one-project-one-month/Hotel-Booking-Management-System-Go/pkg/models"
	"gorm.io/gorm"
)

func Seed(db *gorm.DB) error {
	var count int64
	if err := db.Model(&models.User{}).Count(&count).Error; err != nil {
		return fmt.Errorf("failed to count users: %w", err)
	}

	if count > 0 {
		fmt.Println("users table already has data, skipping seeding...")
		return nil
	}

	const batchSize = 40
	users := make([]models.User, batchSize)
	now := time.Now()

	for i := 0; i < batchSize; i++ {
		seq := i + 1
		users[i] = models.User{
			Name:        fmt.Sprintf("User %d", seq),
			Email:       fmt.Sprintf("user%d@example.com", seq),
			PhoneNumber: fmt.Sprintf("+1%010d", 2000000000+seq),      // Generate unique US phone numbers
			Password:    "$2a$10$prehashedpassword1234567890abcdefg", // Should be properly hashed in production
			Role:        models.RoleUser,
			ImageURL:      fmt.Sprintf("https://example.com/avatars/user%d.jpg", seq),
			Points:      seq * 10,
			Amount:      seq * 100,
			CreatedAt:   now,
			UpdatedAt:   now,
		}
		if seq == 1 {
			users[i].Role = models.RoleAdmin
		}
	}

	result := db.Create(&users)
	if result.Error != nil {
		return fmt.Errorf("failed to seed users: %w", result.Error)
	}

	if result.RowsAffected != int64(batchSize) {
		return fmt.Errorf("expected to seed %d users, but only %d were inserted", batchSize, result.RowsAffected)
	}

	fmt.Printf("successfully seeded %d users\n", batchSize)
	return nil
}
