package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// TVSpecificationSeeder handles seeding TV specifications and specification keys
type TVSpecificationSeeder struct {
	BaseSeeder
}

// NewTVSpecificationSeeder creates a new TV specification seeder
func NewTVSpecificationSeeder() *TVSpecificationSeeder {
	return &TVSpecificationSeeder{
		BaseSeeder: BaseSeeder{name: "TV Specifications"},
	}
}

// Seed implements the Seeder interface
func (tvs *TVSpecificationSeeder) Seed(db *gorm.DB) error {
	// TV-specific specification keys (excluding already existing specs)
	specificationKeys := []string{
		"Smart TV OS",
		"Built-in Apps",
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

	// Create or find specification keys
	for _, keyName := range specificationKeys {
		var existingKey models.SpecificationKeyModel
		result := db.Where("specification_key = ?", keyName).First(&existingKey)

		if result.Error == gorm.ErrRecordNotFound {
			// Create new specification key
			specKey := models.SpecificationKeyModel{
				SpecificationKey: keyName,
			}
			if err := db.Create(&specKey).Error; err != nil {
				continue
			}
		}
	}

	return nil
}
