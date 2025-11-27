package main

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Specification struct {
	ID    int
	Name  string
	Value string
}

type SpecificationTranslation struct {
	SpecificationID int
	Language        string
	Translation     string
}

func runInsertTranslations() {
	// Database connection setup
	dsn := "host=localhost user=your_user password=your_password dbname=your_db port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	// Data to translate and insert
	specifications := []struct {
		Name        string
		Value       string
		Translation string
	}{
		{"display_size", "6.9 inches", "৬.৯ ইঞ্চি"},
		{"processor", "Apple A19 Pro", "অ্যাপল এ১৯ প্রো"},
		{"ram", "12 GB", "১২ জিবি"},
		// Add more translations here
	}

	for _, spec := range specifications {
		var existingSpec Specification
		// Check if the specification exists
		result := db.Where("name = ? AND value = ?", spec.Name, spec.Value).First(&existingSpec)
		if result.Error != nil {
			if result.Error == gorm.ErrRecordNotFound {
				// Insert the specification if it doesn't exist
				existingSpec = Specification{Name: spec.Name, Value: spec.Value}
				if err := db.Create(&existingSpec).Error; err != nil {
					log.Printf("Failed to insert specification: %v", err)
					continue
				}
			} else {
				log.Printf("Error querying specification: %v", result.Error)
				continue
			}
		}

		// Insert the translation
		translation := SpecificationTranslation{
			SpecificationID: existingSpec.ID,
			Language:        "bn",
			Translation:     spec.Translation,
		}
		if err := db.Create(&translation).Error; err != nil {
			log.Printf("Failed to insert translation: %v", err)
		} else {
			fmt.Printf("Inserted translation for specification ID %d\n", existingSpec.ID)
		}
	}
}
