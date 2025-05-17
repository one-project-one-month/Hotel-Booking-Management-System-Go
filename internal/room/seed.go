package room

import (
	"fmt"

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
	query := "INSERT INTO rooms (\n    id, room_no, type, price, status,\n    is_featured, description, img_url, guest_limit,\n    created_at, updated_at\n)\nSELECT\n    uuid_generate_v4(),\n    CASE\n        WHEN seq % 2 = 0 THEN 'DEX-' || LPAD(seq::text, 3, '0')\n        ELSE 'STD-' || LPAD(seq::text, 3, '0')\n        END,\n    CASE WHEN seq % 2 = 0 THEN 'Deluxe' ELSE 'Standard' END,\n    (seq % 2) * 50 + 100, -- Deluxe = 150, Standard = 100\n    CASE\n        WHEN seq % 3 = 0 THEN 'Checked In'\n        WHEN seq % 3 = 1 THEN 'Available'\n        ELSE 'Check Out'\n        END,\n    (seq % 5 = 0),\n    'Room number ' || seq || ' with nice view and facilities.',\n    'https://example.com/images/room' || seq || '.jpg',\n    (seq % 4 + 1),\n    NOW(),\n    NOW()\nFROM generate_series(1, 40) AS seq;"
	err := db.Exec(query).Error
	if err != nil {
		return err
	}
	return nil
}
