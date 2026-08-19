package migrations

import "gorm.io/gorm"

// AddProductIDToFeedback supports installations created before product
// feedback was separated from the productables relationship.
func AddProductIDToFeedback(db *gorm.DB) error {
	// The table may not exist yet on a fresh database. AutoMigrate will create
	// it from FeedbackModel, so this compatibility migration can safely defer.
	if !db.Migrator().HasTable("feedback") {
		return nil
	}
	if !db.Migrator().HasColumn("feedback", "product_id") {
		if err := db.Exec(`ALTER TABLE feedback ADD COLUMN product_id BIGINT NOT NULL DEFAULT 0`).Error; err != nil {
			return err
		}
	}
	if !db.Migrator().HasColumn("feedback", "rating") {
		if err := db.Exec(`ALTER TABLE feedback ADD COLUMN rating VARCHAR(10) NOT NULL DEFAULT ''`).Error; err != nil {
			return err
		}
	}
	if !db.Migrator().HasColumn("feedback", "source_url") {
		if err := db.Exec(`ALTER TABLE feedback ADD COLUMN source_url TEXT`).Error; err != nil {
			return err
		}
	}
	return nil
}
