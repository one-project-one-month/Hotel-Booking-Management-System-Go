package booking

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/one-project-one-month/Hotel-Booking-Management-System-Go/pkg/models"
	"gorm.io/gorm"
)

func Seed(db *gorm.DB) error {
	// Check if bookings already exist
	var count int64
	if err := db.Model(&models.Booking{}).Count(&count).Error; err != nil {
		return fmt.Errorf("failed to count bookings: %w", err)
	}

	if count > 0 {
		fmt.Println("bookings table already has data, skipping seeding...")
		return nil
	}

	// Check if there are rooms available
	var roomCount int64
	if err := db.Model(&models.Room{}).Count(&roomCount).Error; err != nil {
		return fmt.Errorf("failed to count rooms: %w", err)
	}

	if roomCount == 0 {
		return fmt.Errorf("no rooms available to create bookings")
	}

	// Check if there are users available
	var userCount int64
	if err := db.Model(&models.User{}).Count(&userCount).Error; err != nil {
		return fmt.Errorf("failed to count users: %w", err)
	}

	if userCount == 0 {
		return fmt.Errorf("no users available to create bookings")
	}

	// Get existing room IDs
	var roomIDs []uuid.UUID
	if err := db.Model(&models.Room{}).Pluck("id", &roomIDs).Error; err != nil {
		return fmt.Errorf("failed to get room IDs: %w", err)
	}

	// Get existing user IDs
	var userIDs []uuid.UUID
	if err := db.Model(&models.User{}).Pluck("id", &userIDs).Error; err != nil {
		return fmt.Errorf("failed to get user IDs: %w", err)
	}

	const batchSize = 40
	bookings := make([]models.Booking, batchSize)
	now := time.Now()

	for i := 0; i < batchSize; i++ {
		seq := i + 1
		// Use existing room and user IDs with modulo to cycle through them
		roomID := roomIDs[i%len(roomIDs)]
		userID := userIDs[i%len(userIDs)]
		
		bookings[i] = models.Booking{
			UserID:        userID,
			RoomID:        roomID,
			CheckIn:       now.AddDate(0, 0, seq*7),
			GuestCount:    seq%4 + 1,
			DepositAmount: float64(seq * 10),
			TotalAmount:   float64(seq * 100),
			CreatedAt:     now,
			UpdatedAt:     now,
		}
	}

	result := db.Create(&bookings)
	if result.Error != nil {
		return fmt.Errorf("failed to seed bookings: %w", result.Error)
	}

	if result.RowsAffected != int64(batchSize) {
		return fmt.Errorf("expected to seed %d bookings, but only %d were inserted", batchSize, result.RowsAffected)
	}

	fmt.Printf("successfully seeded %d bookings\n", batchSize)
	return nil
}