package migrations

import (
	"log"

	"gorm.io/gorm"
)

// AddFieldsToProductTranslations adds price and specifications fields to product_translations table
func AddFieldsToProductTranslations(db *gorm.DB) error {
	log.Println("Adding price and specifications fields to product_translations table...")

	// Add price column to product_translations table
	if err := db.Exec("ALTER TABLE product_translations ADD COLUMN IF NOT EXISTS price VARCHAR(255)").Error; err != nil {
		log.Printf("Error adding price column: %v", err)
		return err
	}

	// Add specifications column to product_translations table
	if err := db.Exec("ALTER TABLE product_translations ADD COLUMN IF NOT EXISTS specifications JSON").Error; err != nil {
		log.Printf("Error adding specifications column: %v", err)
		return err
	}

	log.Println("✅ Price and specifications fields added to product_translations table successfully")
	return nil
}
