package migrations

import (
	"gorm.io/gorm"
)

// AddRatingToProductReviewTranslations adds rating column to product_review_translations table
func AddRatingToProductReviewTranslations(db *gorm.DB) error {
	// Check if rating column already exists
	if db.Migrator().HasColumn("product_review_translations", "rating") {
		return nil
	}

	// Add rating column with default value of 0
	if err := db.Migrator().AddColumn("product_review_translations", "rating float32 default 0"); err != nil {
		return err
	}

	return nil
}
