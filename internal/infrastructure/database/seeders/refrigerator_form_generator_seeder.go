package seeders

import (
	"encoding/json"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// RefrigeratorFormGeneratorSeeder handles seeding form generator data for Refrigerators category
type RefrigeratorFormGeneratorSeeder struct {
	BaseSeeder
}

// NewRefrigeratorFormGeneratorSeeder creates a new refrigerator form generator seeder
func NewRefrigeratorFormGeneratorSeeder() *RefrigeratorFormGeneratorSeeder {
	return &RefrigeratorFormGeneratorSeeder{
		BaseSeeder: BaseSeeder{name: "Refrigerator Form Generator"},
	}
}

// Seed implements the Seeder interface
func (rfgs *RefrigeratorFormGeneratorSeeder) Seed(db *gorm.DB) error {
	// Refrigerator specification keys for form generation
	refrigeratorSpecificationKeys := []string{
		"Brand",
		"Model Name",
		"Door Type",                 // Top / Bottom / Side-by-Side
		"Capacity",                  // Total capacity (L)
		"Refrigerator Capacity",     // Fresh food capacity (L)
		"Freezer Capacity",          // Freezer capacity (L)
		"Energy Efficiency Rating",  // e.g. 4 Star
		"Energy Star Rating",        // optional star label
		"Annual Energy Consumption", // kWh/year
		"Product Dimensions",        // W x D x H
		"Product Weight",            // kg
		"Color",
		"Compressor Type",
		"Cooling Technology",
		"Defrost Type",
		"Temperature Control",
		"Shelf Material",
		"Number of Shelves",
		"Door Bins",
		"Crisper Drawers",
		"Ice Maker",
		"Water Dispenser",
		"Noise Level",
		"Voltage",
		"Frequency (Hz)",
		"App Control",
		"Voice Assistant Support",
		"Warranty Period",             // Product warranty (years)
		"Compressor Warranty (Years)", // Compressor warranty (years)
		"Special Features",
	}

	// Ensure all specification keys exist and collect their IDs
	var specKeyIDs []uint
	for _, keyName := range refrigeratorSpecificationKeys {
		var specKey models.SpecificationKeyModel
		if err := db.Where("specification_key = ?", keyName).First(&specKey).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				// Create the specification key if it does not exist
				newKey := &models.SpecificationKeyModel{SpecificationKey: keyName}
				if createErr := db.Create(newKey).Error; createErr != nil {
					return createErr
				}
				specKey = *newKey
			} else {
				return err
			}
		}

		specKeyIDs = append(specKeyIDs, specKey.ID)
	}

	// Convert IDs to JSON
	specKeyIDsJSON, err := json.Marshal(specKeyIDs)
	if err != nil {
		return err
	}

	// Find Refrigerator category by slug
	var refrigeratorCategory models.CategoryModel
	if err := db.Where("slug = ?", "refrigerators").First(&refrigeratorCategory).Error; err != nil {
		// If not found, try by name as a fallback
		if err := db.Where("name = ?", "Refrigerators").First(&refrigeratorCategory).Error; err != nil {
			return nil // Skip if category not found
		}
	}

	// Check if form generator entry already exists for refrigerator category
	var existingForm models.FormGeneratorModel
	result := db.Where("category_id = ?", refrigeratorCategory.ID).First(&existingForm)

	if result.Error == gorm.ErrRecordNotFound {
		// Create form generator entry
		formGenerator := &models.FormGeneratorModel{
			CategoryID:      refrigeratorCategory.ID,
			SpecificationID: string(specKeyIDsJSON),
			Status:          "active",
		}

		if err := db.Create(formGenerator).Error; err != nil {
			return err
		}
	} else if result.Error == nil {
		// Update existing form generator with the refrigerator spec keys
		if err := db.Model(&existingForm).Updates(map[string]interface{}{
			"specification_id": string(specKeyIDsJSON),
			"status":           "active",
		}).Error; err != nil {
			return err
		}
	}

	return nil
}
