package user

import (
	"fmt"
	"time"

	"github.com/one-project-one-month/Hotel-Booking-Management-System-Go/pkg/models"
	"golang.org/x/crypto/bcrypt"
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

	const batchSize = 4
	users := make([]models.User, batchSize)
	now := time.Now()

	// Define user data with plain text passwords that will be hashed
	userData := []struct {
		name     string
		email    string
		phone    string
		password string
		role     models.UserRole
		points   int
	}{
		{"Admin User", "admin@hotel.com", "+12001001001", "admin123", models.RoleAdmin, 100},
		{"John Doe", "john.doe@example.com", "+12001001002", "password123", models.RoleUser, 50},
		{"Jane Smith", "jane.smith@example.com", "+12001001003", "securepass", models.RoleUser, 75},
		{"Bob Johnson", "bob.johnson@example.com", "+12001001004", "mypassword", models.RoleUser, 25},
	}

	for i := 0; i < batchSize; i++ {
		// Hash the password
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userData[i].password), bcrypt.DefaultCost)
		if err != nil {
			return fmt.Errorf("failed to hash password for user %s: %w", userData[i].name, err)
		}

		users[i] = models.User{
			Name:        userData[i].name,
			Email:       userData[i].email,
			PhoneNumber: userData[i].phone,
			Password:    string(hashedPassword),
			Role:        userData[i].role,
			ImageURL:    fmt.Sprintf("https://example.com/avatars/user%d.jpg", i+1),
			Points:      userData[i].points,
			CreatedAt:   now,
			UpdatedAt:   now,
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
	fmt.Println("Seeded users:")
	for _, user := range userData {
		fmt.Printf("- %s (%s) - Password: %s\n", user.name, user.email, user.password)
	}

	return nil
}
