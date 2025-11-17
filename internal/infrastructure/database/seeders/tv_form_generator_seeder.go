package seeders

import (
	"encoding/json"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// TVFormGeneratorSeeder handles seeding form generator data for TV category
type TVFormGeneratorSeeder struct {
	BaseSeeder
}

// NewTVFormGeneratorSeeder creates a new TV form generator seeder
func NewTVFormGeneratorSeeder() *TVFormGeneratorSeeder {
	return &TVFormGeneratorSeeder{
		BaseSeeder: BaseSeeder{name: "TV Form Generator"},
	}
}

// Seed implements the Seeder interface
func (tvfgs *TVFormGeneratorSeeder) Seed(db *gorm.DB) error {
	// TV specification keys for form generation
	tvSpecificationKeys := []string{
		"Brand",
		"Model Name",
		"Screen Size",
		"Resolution",
		"Panel Type",
		"Refresh Rate",
		"Brightness (Nits)",
		"Contrast Ratio",
		"Color Accuracy",
		"Viewing Angle",
		"Response Time",
		"HDR Support",
		"Dolby Vision",
		"Local Dimming",
		"Backlight Type",
		"Smart TV OS",
		"Built-in Apps",
		"WiFi Connectivity",
		"Bluetooth Connectivity",
		"HDMI Ports",
		"USB Ports",
		"Audio Output (Watts)",
		"Speaker Configuration",
		"Dolby Atmos Support",
		"Built-in Tuner",
		"Supported Video Formats",
		"Supported Audio Formats",
		"Frame Rate Support",
		"ALLM Support",
		"VRR Support",
		"Game Mode",
		"Motion Smoothing",
		"Picture Quality Modes",
		"Energy Consumption (Watts)",
		"Power Consumption Standby (Watts)",
		"Dimensions (W x H x D)",
		"Weight Without Stand (kg)",
		"Weight With Stand (kg)",
		"VESA Mount Pattern",
		"Warranty Period (Years)",
		"Warranty Coverage",
		"Lifespan (Hours)",
		"Thickness",
		"Bezel Width",
		"Remote Type",
		"Voice Control Support",
		"Gaming Features",
		"Sports Mode",
		"Movie Mode",
		"Eye Comfort",
		"Flicker-Free",
		"Blue Light Filter",
		"Eco Mode",
		"Auto Brightness",
		"Motion Interpolation",
		"Upscaling Technology",
		"Color Enhancement",
		"Price (BDT)",
		"Availability in Bangladesh",
		"After-Sales Service",
		"Return / Exchange Policy",
	}

	// Get all specification key IDs for TVs
	var specKeyIDs []uint
	for _, keyName := range tvSpecificationKeys {
		var specKey models.SpecificationKeyModel
		if err := db.Where("specification_key = ?", keyName).First(&specKey).Error; err == nil {
			specKeyIDs = append(specKeyIDs, specKey.ID)
		}
	}

	// Convert IDs to JSON
	specKeyIDsJSON, err := json.Marshal(specKeyIDs)
	if err != nil {
		return err
	}

	// Find TV category (ID: 124)
	var tvCategory models.CategoryModel
	if err := db.Where("id = ?", 124).First(&tvCategory).Error; err != nil {
		// If not found by ID, try by slug
		if err := db.Where("slug = ?", "tv").First(&tvCategory).Error; err != nil {
			return nil // Skip if category not found
		}
	}

	// Check if form generator entry already exists for TV category
	var existingForm models.FormGeneratorModel
	result := db.Where("category_id = ?", tvCategory.ID).First(&existingForm)

	if result.Error == gorm.ErrRecordNotFound {
		// Create form generator entry
		formGenerator := &models.FormGeneratorModel{
			CategoryID:      tvCategory.ID,
			SpecificationID: string(specKeyIDsJSON),
			Status:          "active",
		}

		if err := db.Create(formGenerator).Error; err != nil {
			return err
		}
	}

	return nil
}
