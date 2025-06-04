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
	now := time.Now()

	// Start a transaction to ensure all operations succeed or fail together
	err := db.Transaction(func(tx *gorm.DB) error {
		// First create the CheckInOut records
		checkInOuts := make([]models.CheckInOut, batchSize)
		for i := 0; i < batchSize; i++ {
			seq := i + 1
			checkInDate := now.AddDate(0, 0, seq*7)
			// Set a default check-out date 3 days after check-in
			checkOutDate := checkInDate.AddDate(0, 0, 3)
			
			checkInOuts[i] = models.CheckInOut{
				CheckIn:     checkInDate,
				CheckOut:    checkOutDate,
				Status:      "pending",
				ExtraCharge: 0, // No extra charges for new bookings
				CreatedAt:   now,
				UpdatedAt:   now,
			}
		}

		// Create all CheckInOut records
		if result := tx.Create(&checkInOuts); result.Error != nil {
			return fmt.Errorf("failed to seed check-in/check-out records: %w", result.Error)
		}

		// Now create the bookings with references to the CheckInOut records
		bookings := make([]models.Booking, batchSize)
		for i := 0; i < batchSize; i++ {
			seq := i + 1
			// Use existing room and user IDs with modulo to cycle through them
			roomID := roomIDs[i%len(roomIDs)]
			userID := userIDs[i%len(userIDs)]

			bookings[i] = models.Booking{
				UserID:        userID,
				RoomID:        roomID,
				CheckIn:       checkInOuts[i].CheckIn,
				CheckOut:      checkInOuts[i].CheckOut,
				GuestCount:    seq%4 + 1,
				DepositAmount: float64(seq * 10),
				TotalAmount:   float64(seq * 100),
				Status:        "pending",
				CreatedAt:     now,
				UpdatedAt:     now,
				CheckInOutID:  checkInOuts[i].ID, // Link to the corresponding CheckInOut record
			}
		}

		// Create all booking records
		if result := tx.Create(&bookings); result.Error != nil {
			return fmt.Errorf("failed to seed bookings: %w", result.Error)
		}

		if result := tx.Model(&models.Booking{}).Count(&count); result.Error != nil {
			return fmt.Errorf("failed to verify bookings count: %w", result.Error)
		}

		if count != int64(batchSize) {
			return fmt.Errorf("expected to seed %d bookings, but only %d were inserted", batchSize, count)
		}

		return nil
	})

	if err != nil {
		return err
	}

	fmt.Printf("successfully seeded %d bookings with check-in/check-out records\n", batchSize)
	return nil
}
