package seeders

import (
	"kossti/internal/domain/entities"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationKeySeeder handles seeding specification keys
type SpecificationKeySeeder struct {
	BaseSeeder
}

// NewSpecificationKeySeeder creates a new specification key seeder
func NewSpecificationKeySeeder() *SpecificationKeySeeder {
	return &SpecificationKeySeeder{
		BaseSeeder: BaseSeeder{name: "Specification Keys"},
	}
}

// Seed implements the Seeder interface
func (sks *SpecificationKeySeeder) Seed(db *gorm.DB) error {
	keys := []string{
		"Technology",
		"2G Bands",
		"3G Bands",
		"4G Bands",
		"Network Speed",
		"Announcement Date",
		"Device Status",
		"Dimensions",
		"Weight",
		"Build Material",
		"SIM Card Type",
		"Display Type",
		"Screen Size",
		"Resolution",
		"Operating System",
		"Chipset",
		"CPU Type",
		"GPU Type",
		"Card Slot Type",
		"Internal Memory Capacity",
		"Dual SIM",
		"Special Features",
		"Main Camera Video Resolution",
		"Front Camera Resolution",
		"Front Camera Video Resolution",
		"Loudspeaker Quality",
		"3.5mm Audio Jack",
		"WiFi",
		"Bluetooth Version",
		"Positioning System",
		"NFC Support",
		"Radio Support",
		"USB Type",
		"Sensors",
		"Battery Capacity",
		"Charging Speed",
		"Available Colors",
		"Model Variants",
		"SAR (Specific Absorption Rate) EU",
		"Price",
		"Triple Camera Setup",
		"SAR (Specific Absorption Rate)",
		"5G Bands",
		"Performance Metrics",
		"Display Characteristics",
		"New Battery Capacity",
		"Screen Protection",
		"Camera Features",
		"Old Battery Capacity",
		"Quad Camera Setup",
		"Audio Quality",
		"GPRS (General Packet Radio Service)",
		"EDGE (Enhanced Data Rates for GSM Evolution)",
		"Talk Time Duration",
		"Music Playback Duration",
		"Standby Time",
		"Infrared Port",
		"Device Type",
		"Phonebook Capacity",
		"Call Records",
		"Messaging Features",
		"Preinstalled Games",
		"Java Support",
		"Penta Camera Setup",
		"Dual or Triple Camera Setup",
		"Web Browser",
		"Alert Types",
		"Physical Keyboard",
		"Device Size",
		"Clock Feature",
		"Alarm Feature",
		"No Support",
		"Supported Languages",
	}

	for _, key := range keys {
		// Create SpecificationKey entity
		specKey := &entities.SpecificationKey{
			SpecificationKey: key,
		}

		// Check if specification key already exists
		var existingModel models.SpecificationKeyModel
		result := db.Where("specification_key = ?", key).First(&existingModel)

		if result.Error == gorm.ErrRecordNotFound {
			// Convert entity to model for database insertion
			model := &models.SpecificationKeyModel{
				SpecificationKey: specKey.SpecificationKey,
			}

			if err := db.Create(model).Error; err != nil {
				return err
			}
		}
	}

	return nil
}
