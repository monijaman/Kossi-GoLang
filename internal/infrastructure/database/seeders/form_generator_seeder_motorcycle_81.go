package seeders

import (
	"encoding/json"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// FormGeneratorSeederMotorcycle81 seeds form generator config for motorcycles (product id 81 context)
// It stores a JSON array of specification keys relevant to the Motorcycles category.
type FormGeneratorSeederMotorcycle81 struct {
	BaseSeeder
}

// NewFormGeneratorSeederMotorcycle81 creates a new seeder instance
func NewFormGeneratorSeederMotorcycle81() *FormGeneratorSeederMotorcycle81 {
	return &FormGeneratorSeederMotorcycle81{BaseSeeder: BaseSeeder{name: "Form Generator (Motorcycles)"}}
}

// Seed inserts a form generator config for the Motorcycles category
func (s *FormGeneratorSeederMotorcycle81) Seed(db *gorm.DB) error {
	// Ensure category exists
	cat, err := CreateOrFindCategory(db, "Motorcycles", "motorcycles")
	if err != nil {
		return err
	}

	// Specification keys (use the same keys used for motorcycle specs)
	keyNames := []string{
		"ABS (Anti-lock Braking System)",
		"CC (Cubic Capacity)",
		"Clutch Type",
		"Cooling System",
		"Displacement",
		"Drive Type",
		"Engine Type",
		"Fuel Capacity",
		"Fuel Tank Capacity",
		"Fuel Type",
		"Gearbox",
		"Ground Clearance",
		"Headlight Type",
		"Horsepower",
		"Ignition Type",
		"Kerb Weight",
		"Length",
		"Max Power",
		"Max Torque",
		"Mileage",
		"Number of Gears",
		"Number of Seats",
		"Seating Capacity",
		"Starting System",
		"Suspension Type",
		"Top Speed",
		"Torque",
		"Transmission",
		"Tyre Size",
		"Tyre Type",
		"Valve Configuration",
		"Wheelbase",
		"Width",
		"Brake Type",
		"Front Brake Type",
		"Rear Brake Type",
		"Rim Type",
		"Rim Size",
		"Seat Height",
		"Battery Capacity",
		"Fuel System",
		"Emission Standard",
		"Compression Ratio",
		"Instrument Cluster",
		"DRL (Daytime Running Light)",
		"Spark Plug Type",
		"Electrical System",
	}

	// Resolve keys to their numeric IDs and marshal as JSON array of IDs
	var keyIDs []uint
	for _, name := range keyNames {
		sk, err := CreateOrFindSpecificationKey(db, name)
		if err != nil {
			return err
		}
		keyIDs = append(keyIDs, sk.ID)
	}

	specJSON, err := json.Marshal(keyIDs)
	if err != nil {
		return err
	}

	// Check existing form generator for this category
	var existing models.FormGeneratorModel
	if err := db.Where("category_id = ?", cat.ID).First(&existing).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			fg := &models.FormGeneratorModel{
				CategoryID:      cat.ID,
				SpecificationID: string(specJSON),
				Status:          "active",
			}

			if err := db.Create(fg).Error; err != nil {
				return err
			}
		} else {
			return err
		}
	}

	return nil
}
