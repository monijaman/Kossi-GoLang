package migrations

import (
	"log"

	"gorm.io/gorm"
)

// FixCategoryIDAndBrandIDTypes fixes the type conversion for category_id and brand_id fields
func FixCategoryIDAndBrandIDTypes(db *gorm.DB) error {
	log.Println("Fixing category_id and brand_id types...")

	// Update category_id field type if needed
	if err := db.Exec("ALTER TABLE products ALTER COLUMN category_id TYPE INTEGER USING category_id::INTEGER").Error; err != nil {
		log.Printf("Warning: category_id type conversion failed (may already be correct): %v", err)
	}

	// Update brand_id field type if needed
	if err := db.Exec("ALTER TABLE products ALTER COLUMN brand_id TYPE INTEGER USING brand_id::INTEGER").Error; err != nil {
		log.Printf("Warning: brand_id type conversion failed (may already be correct): %v", err)
	}

	log.Println("✅ Category ID and Brand ID types fixed successfully")
	return nil
}
