package migrations

import (
	"fmt"

	"gorm.io/gorm"
)

// ConvertProductReviewRatingToString converts the rating column in product_reviews from numeric to varchar
func ConvertProductReviewRatingToString(db *gorm.DB) error {
	fmt.Println("🔄 Converting product_reviews.rating column from numeric to varchar...")

	// Check if column exists
	if !db.Migrator().HasColumn("product_reviews", "rating") {
		fmt.Println("📋 Column 'rating' doesn't exist yet in product_reviews")
		return nil
	}

	// Try to alter the column to varchar
	alterSQL := `
		ALTER TABLE product_reviews 
		ALTER COLUMN rating TYPE varchar(50) USING rating::text;
	`

	result := db.Exec(alterSQL)
	if result.Error != nil {
		fmt.Printf("⚠️  Column alter returned error: %v\n", result.Error)
	} else {
		fmt.Println("✓ Successfully converted product_reviews.rating column to varchar")
	}

	return nil
}
