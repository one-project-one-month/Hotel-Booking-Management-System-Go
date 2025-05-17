package user

import "gorm.io/gorm"

// Repository handles user data persistence operations.
type Repository struct {
	db *gorm.DB // Will be implemented in future releases
}
