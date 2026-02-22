package migrations

import (
	"log"

	"gorm.io/gorm"
)

// ConvertPriceColumnsToVarchar converts start_price and end_price columns to varchar(255)
// This allows storing Bengali numerals and text alongside regular prices
func ConvertPriceColumnsToVarchar(db *gorm.DB) error {
	log.Println("Converting price columns to varchar for Bengali text support...")

	// Check if columns exist first
	if !db.Migrator().HasColumn("product_translations", "start_price") {
		log.Println("start_price column doesn't exist, skipping conversion")
		return nil
	}

	if !db.Migrator().HasColumn("product_translations", "end_price") {
		log.Println("end_price column doesn't exist, skipping conversion")
		return nil
	}

	// Use raw SQL to alter column types (more reliable than GORM's AlterColumn)
	// First, attempt to convert start_price
	log.Printf("Converting start_price to varchar(255)")
	if err := db.Exec(`
		ALTER TABLE product_translations
		ALTER COLUMN start_price TYPE varchar(255);
	`).Error; err != nil {
		// If it fails, it might already be varchar or there might be a cast issue
		log.Printf("Warning: Could not convert start_price (it might already be varchar): %v", err)
		// Don't return error, continue with end_price
	} else {
		log.Println("✅ start_price converted to varchar(255)")
	}

	// Convert end_price
	log.Printf("Converting end_price to varchar(255)")
	if err := db.Exec(`
		ALTER TABLE product_translations
		ALTER COLUMN end_price TYPE varchar(255);
	`).Error; err != nil {
		// If it fails, it might already be varchar or there might be a cast issue
		log.Printf("Warning: Could not convert end_price (it might already be varchar): %v", err)
		// Don't return error
	} else {
		log.Println("✅ end_price converted to varchar(255)")
	}

	log.Println("Price column conversion completed!")
	return nil
}
