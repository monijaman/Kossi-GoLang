package migrations

import (
	"log"

	"gorm.io/gorm"
)

// AddDescriptionToProducts adds description field to products table
func AddDescriptionToProducts(db *gorm.DB) error {
	log.Println("Adding description field to products table...")

	// Add description column to products table
	if err := db.Exec("ALTER TABLE products ADD COLUMN IF NOT EXISTS description TEXT").Error; err != nil {
		log.Printf("Error adding description column: %v", err)
		return err
	}

	log.Println("✅ Description field added to products table successfully")
	return nil
}
