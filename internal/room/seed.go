package room

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/one-project-one-month/Hotel-Booking-Management-System-Go/pkg/models"
	"gorm.io/gorm"
)

func Seed(db *gorm.DB) error {
	var count int64
	if err := db.Table("rooms").Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		// Table already has data, skip seeding
		fmt.Println("skipping seeding...")
		return nil
	}

	file, err := os.Open("internal/room/mock.json")
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	var seeds []RequestRoomDto
	if err := json.NewDecoder(file).Decode(&seeds); err != nil {
		panic(err)
		return err
	}

	for _, s := range seeds {
		detailsJSON, _ := json.Marshal(s.Details)
		imgURLJSON, _ := json.Marshal(s.ImgURL)
		// isFeatured, _ := strconv.ParseBool(s.IsFeatured)

		room := models.Room{
			ID:         uuid.New(),
			RoomNo:     s.RoomNo,
			Type:       s.Type,
			Price:      s.Price,
			Status:     "available",
			IsFeatured: s.IsFeatured,
			Details:    string(detailsJSON),
			ImgURL:     string(imgURLJSON),
			GuestLimit: s.GuestLimit,
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		}

		if err := db.Create(&room).Error; err != nil {
			panic(err)
		}
	}

	return nil
}
