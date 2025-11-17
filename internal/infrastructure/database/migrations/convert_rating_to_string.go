package migrations

import (
	"fmt"

	"gorm.io/gorm"
)

// ConvertRatingToString converts the rating column from numeric to text
// and migrates existing numeric values to strings (e.g., 4.15 -> "4.15")
func ConvertRatingToString(db *gorm.DB) error {
	// Check if column exists
	if !db.Migrator().HasColumn("product_review_translations", "rating") {
		// Column doesn't exist yet, AutoMigrate will create it as varchar (from the model)
		fmt.Println("📋 Column 'rating' doesn't exist yet - will be created by AutoMigrate")
		return nil
	}

	fmt.Println("🔄 Attempting to convert rating column to varchar...")

	// Try to alter the column to varchar
	// PostgreSQL syntax: ALTER TABLE table_name ALTER COLUMN column_name TYPE varchar(50) USING column_name::varchar(50);
	alterSQL := `
		ALTER TABLE product_review_translations 
		ALTER COLUMN rating TYPE varchar(50) USING rating::text;
	`

	result := db.Exec(alterSQL)
	if result.Error != nil {
		fmt.Printf("⚠️  Column alter returned error: %v\n", result.Error)
	} else {
		fmt.Println("✓ Successfully converted rating column to varchar")
	}

	return nil
}
