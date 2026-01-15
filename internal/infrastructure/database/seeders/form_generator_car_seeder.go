package seeders

import (
	"encoding/json"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// FormGeneratorCarSeeder handles seeding form generator data for Car category
type FormGeneratorCarSeeder struct {
	BaseSeeder
}

// NewFormGeneratorCarSeeder creates a new car form generator seeder
func NewFormGeneratorCarSeeder() *FormGeneratorCarSeeder {
	return &FormGeneratorCarSeeder{
		BaseSeeder: BaseSeeder{name: "Form Generator - Car"},
	}
}

// Seed implements the Seeder interface
func (fgcs *FormGeneratorCarSeeder) Seed(db *gorm.DB) error {
	// Car category specification keys organized in 14 sections
	carSpecificationKeys := []string{
		// Basic Information (728-735)
		"Variant",
		"Generation",
		"Segment",
		"Launch Year",
		"Engine Configuration",
		"Valves Per Cylinder",
		"Engine Aspiration",
		"Differential Type",

		// Performance & Efficiency (736-739)
		"Power to Weight (HP/ton)",
		"Mileage City (km/L)",
		"Mileage Highway (km/L)",
		"Mileage Combined (km/L)",

		// Electric/Hybrid (740-743)
		"Battery Capacity (kWh)",
		"Motor Power (kW)",
		"Motor Torque (Nm)",
		"Charging Type",

		// Suspension & Steering (744-747)
		"Front Suspension",
		"Rear Suspension",
		"Steering Type",
		"Steering Adjustment",

		// Wheels & Tyres (748-749)
		"Wheel Size (inch)",
		"Spare Wheel Type",

		// Exterior Features (750-756)
		"DRL",
		"Fog Lamp Type",
		"Alloy Wheels",
		"Sunroof Type",
		"Roof Rails",
		"ORVM Type",
		"Wiper Type",

		// Interior & Comfort (757-764)
		"Driver Seat Adjustment",
		"Ventilated Seats",
		"Infotainment Screen (inch)",
		"Apple CarPlay",
		"Android Auto",
		"Sound System Brand",
		"Number of Speakers",
		"Ambient Lighting",

		// Safety Features (765-774)
		"EBD",
		"Traction Control",
		"ESC",
		"Hill Hold",
		"ISOFIX Mounts",
		"Camera Type",
		"Adaptive Cruise Control",
		"Lane Keep Assist",
		"Automatic Emergency Braking",
		"Blind Spot Monitor",

		// Technology & Convenience (775-781)
		"Keyless Entry",
		"Push Button Start",
		"Digital Instrument Cluster",
		"Heads Up Display",
		"Drive Modes",
		"Connected Car Features",
		"OTA Updates",

		// Warranty & Pricing (782-785)
		"Vehicle Warranty (Years)",
		"Engine Warranty (Years)",
		"Battery Warranty (Years)",
		"Ex-Showroom Price (USD)",
	}

	// Get all specification key IDs for cars
	var specKeyIDs []uint
	for _, keyName := range carSpecificationKeys {
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

	// Find Car category (ID: 18)
	var carCategory models.CategoryModel
	if err := db.Where("id = ?", 18).First(&carCategory).Error; err != nil {
		// If not found, try by slug
		if err := db.Where("slug = ?", "car").First(&carCategory).Error; err != nil {
			return err
		}
	}

	// Check if form generator entry already exists for car category
	var existingForm models.FormGeneratorModel
	result := db.Where("category_id = ?", carCategory.ID).First(&existingForm)

	if result.Error == gorm.ErrRecordNotFound {
		// Create form generator entry
		formGenerator := &models.FormGeneratorModel{
			CategoryID:      carCategory.ID,
			SpecificationID: string(specKeyIDsJSON),
			Status:          "active",
		}

		if err := db.Create(formGenerator).Error; err != nil {
			return err
		}
	}

	return nil
}
