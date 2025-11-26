package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// CarSpecificationSeeder handles seeding car specifications and specification keys
type CarSpecificationSeeder struct {
	BaseSeeder
}

// NewCarSpecificationSeeder creates a new car specification seeder
func NewCarSpecificationSeeder() *CarSpecificationSeeder {
	return &CarSpecificationSeeder{
		BaseSeeder: BaseSeeder{name: "Car Specifications"},
	}
}

// Seed implements the Seeder interface
func (cs *CarSpecificationSeeder) Seed(db *gorm.DB) error {
	// Car-specific specification keys (excluding already existing specs)
	specificationKeys := []string{
		"Engine Type",
		"Fuel Efficiency (km/l)",
		"Transmission Type",
		"Seating Capacity",
		"Airbags",
		"ABS (Anti-lock Braking System)",
		"Infotainment System",
		"Cruise Control",
		"Parking Sensors",
		"Rear Camera",
		"Sunroof",
		"Ground Clearance (mm)",
		"Boot Space (liters)",
		"Warranty Period (Years)",
		"Price (BDT)",
		"Availability in Bangladesh",
		"After-Sales Service",
		"Return / Exchange Policy",
		"Brand",
		"Model",
		"Year",
		"Body Type",
		"Engine Displacement (cc)",
		"Cylinders",
		"Max Power (hp)",
		"Max Torque (Nm)",
		"City Fuel Economy (km/l)",
		"Highway Fuel Economy (km/l)",
		"Combined Fuel Economy (km/l)",
		"Fuel Tank Capacity (l)",
		"Length (mm)",
		"Width (mm)",
		"Height (mm)",
		"Wheelbase (mm)",
		"Drive Type",
		"AC Type",
		"Power Windows",
		"Power Steering",
		"Infotainment System",
		"Seat Material",
		"Driver Airbag",
		"Passenger Airbag",
		"Side Airbags",
		"Curtain Airbags",
		"Camera",
		"Parking Sensors",
		"Child Safety Lock",
		"Stability Control",
		"Service Interval (km)",
		"Spare Parts Availability",
		"Estimated Annual Maintenance Cost (BDT)",
		"Registration Cost (BDT)",
		"Insurance Cost (BDT)",
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
