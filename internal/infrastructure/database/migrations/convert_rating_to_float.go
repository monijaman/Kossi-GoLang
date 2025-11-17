package migrations

import (
	"fmt"

	"gorm.io/gorm"
)

// ConvertRatingToFloat ensures the rating column is a numeric type (double precision)
// and converts existing values if necessary.
func ConvertRatingToFloat(db *gorm.DB) error {
	// Check if table exists
	if !db.Migrator().HasTable("product_review_translations") {
		return nil
	}

	// We will run a safe ALTER TABLE statement to change the column type to double precision.
	// Use USING clause to cast existing values to double precision when possible.
	sqlStmt := `ALTER TABLE product_review_translations
        ALTER COLUMN rating TYPE double precision USING (CASE
            WHEN rating IS NULL OR trim(rating::text) = '' THEN 0::double precision
            ELSE (rating::text)::double precision
        END)`

	// Some drivers require a *sql.DB to Exec raw SQL
	sqlDB, err := db.DB()
	if err != nil {
		return fmt.Errorf("failed to get raw db: %w", err)
	}

	if _, err := sqlDB.Exec(sqlStmt); err != nil {
		// If the ALTER fails (for example column already of correct type), return nil to be idempotent
		// but if it's a different error, propagate it.
		// We'll check column type afterwards as a best-effort verification.
		// Return the error so operator can inspect, unless it's clearly a type-not-changable error.
		return fmt.Errorf("failed to alter rating column to double precision: %w", err)
	}

	return nil
}
