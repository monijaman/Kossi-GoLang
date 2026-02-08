package migrations

import (
	"gorm.io/gorm"
)

// AddDefaultPhotoToImages adds default_photo column to images table
func AddDefaultPhotoToImages(db *gorm.DB) error {
	// Check if default_photo column already exists
	if db.Migrator().HasColumn("images", "default_photo") {
		return nil
	}

	// Add default_photo column with default value of 0
	if err := db.Migrator().AddColumn("images", "default_photo integer default 0"); err != nil {
		return err
	}

	return nil
}
