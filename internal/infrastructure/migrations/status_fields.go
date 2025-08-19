package migrations

import (
	"log"

	"gorm.io/gorm"
)

// MigrateStatusFields handles migration of status fields from boolean to integer
func MigrateStatusFields(db *gorm.DB) error {
	log.Println("Starting status fields migration...")

	// List of tables that need status field migration from boolean to integer
	tables := []struct {
		tableName string
		hasStatus bool
	}{
		{"products", true},
		{"feedback", true},
		{"categories", false},     // Will be added as new field
		{"images", false},         // Will be added as new field
		{"specifications", false}, // Will be added as new field
	}

	for _, table := range tables {
		// Check if table exists first
		var exists bool
		err := db.Raw("SELECT EXISTS (SELECT FROM information_schema.tables WHERE table_schema = 'public' AND table_name = ?)", table.tableName).Scan(&exists).Error
		if err != nil {
			log.Printf("Error checking if table %s exists: %v", table.tableName, err)
			continue
		}

		if !exists {
			log.Printf("Table %s does not exist, skipping migration", table.tableName)
			continue
		}

		if table.hasStatus {
			// Check if status field is already integer type
			var dataType string
			err := db.Raw("SELECT data_type FROM information_schema.columns WHERE table_name = ? AND column_name = 'status'", table.tableName).Scan(&dataType).Error
			if err != nil {
				log.Printf("Error checking status field type for %s: %v", table.tableName, err)
				continue
			}

			if dataType == "integer" || dataType == "bigint" {
				log.Printf("Status field for table %s is already integer type (%s), skipping conversion", table.tableName, dataType)
				continue
			}

			// For tables that already have boolean status, convert to integer
			log.Printf("Converting status field from boolean to integer for table: %s", table.tableName)

			// Step 1: Add a temporary integer status column
			err = db.Exec("ALTER TABLE " + table.tableName + " ADD COLUMN status_int INTEGER DEFAULT 1").Error
			if err != nil {
				log.Printf("Warning: Could not add temp column to %s: %v", table.tableName, err)
			}

			// Step 2: Copy boolean values to integer (true=1, false=0)
			err = db.Exec("UPDATE " + table.tableName + " SET status_int = CASE WHEN status::boolean = true THEN 1 ELSE 0 END").Error
			if err != nil {
				return err
			}

			// Step 3: Drop the old boolean column
			err = db.Exec("ALTER TABLE " + table.tableName + " DROP COLUMN status").Error
			if err != nil {
				return err
			}

			// Step 4: Rename the temporary column to status
			err = db.Exec("ALTER TABLE " + table.tableName + " RENAME COLUMN status_int TO status").Error
			if err != nil {
				return err
			}

			log.Printf("Successfully converted status field for table: %s", table.tableName)
		} else {
			// For tables that don't have status field, add it as integer
			log.Printf("Adding new integer status field to table: %s", table.tableName)

			err := db.Exec("ALTER TABLE " + table.tableName + " ADD COLUMN status INTEGER DEFAULT 1").Error
			if err != nil {
				log.Printf("Warning: Could not add status column to %s (might already exist): %v", table.tableName, err)
			}
		}
	}

	log.Println("Status fields migration completed!")
	return nil
}
