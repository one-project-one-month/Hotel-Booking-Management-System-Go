package user

import "gorm.io/gorm"

type Repository struct {
	db *gorm.DB
}
