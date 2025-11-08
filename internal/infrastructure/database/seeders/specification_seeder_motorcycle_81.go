package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMotorcycle81 seeds specifications/options for product ID 81
type SpecificationSeederMotorcycle81 struct {
	BaseSeeder
}

// NewSpecificationSeederMotorcycle81 creates a new seeder instance
func NewSpecificationSeederMotorcycle81() *SpecificationSeederMotorcycle81 {
	return &SpecificationSeederMotorcycle81{BaseSeeder: BaseSeeder{name: "Specifications for Product 81"}}
}

// Seed inserts specification records for product 81 using relevant keys
func (s *SpecificationSeederMotorcycle81) Seed(db *gorm.DB) error {
	productID := uint(81)

	// Only keys relevant to motorcycles (subset from specification_key_seeder.go)
	specs := map[string]string{
		"ABS (Anti-lock Braking System)": "Yes",
		"CC (Cubic Capacity)":            "150 cc",
		"Clutch Type":                    "Wet multi-plate",
		"Cooling System":                 "Air-cooled",
		"Displacement":                   "149.5 cc",
		"Drive Type":                     "Chain",
		"Engine Type":                    "Single-cylinder, 4-stroke",
		"Fuel Capacity":                  "12 L",
		"Fuel Tank Capacity":             "12 L",
		"Fuel Type":                      "Petrol",
		"Gearbox":                        "Manual",
		"Ground Clearance":               "170 mm",
		"Headlight Type":                 "LED",
		"Horsepower":                     "14.8 HP",
		"Ignition Type":                  "Digital",
		"Kerb Weight":                    "140 kg",
		"Length":                         "2000 mm",
		"Max Power":                      "14.8 HP @ 8500 rpm",
		"Max Torque":                     "13.9 Nm @ 6500 rpm",
		"Mileage":                        "45 km/l",
		"Number of Gears":                "5",
		"Number of Seats":                "2",
		"Seating Capacity":               "2",
		"Starting System":                "Electric & Kick",
		"Suspension Type":                "Telescopic front, Twin shock rear",
		"Top Speed":                      "120 km/h",
		"Torque":                         "13.9 Nm",
		"Transmission":                   "5-speed",
		"Tyre Size":                      "90/90-17 (front), 120/80-17 (rear)",
		"Tyre Type":                      "Tubeless",
		"Valve Configuration":            "SOHC",
		"Wheelbase":                      "1330 mm",
		"Width":                          "760 mm",
		// Additional motorcycle-specific keys
		"Brake Type":                  "Disc",
		"Front Brake Type":            "Disc",
		"Rear Brake Type":             "Disc",
		"Rim Type":                    "Alloy",
		"Rim Size":                    "17 inch",
		"Seat Height":                 "800 mm",
		"Battery Capacity":            "12V",
		"Fuel System":                 "Fuel Injection",
		"Emission Standard":           "Euro 3",
		"Compression Ratio":           "9.5:1",
		"Instrument Cluster":          "Digital",
		"DRL (Daytime Running Light)": "Yes",
		"Spark Plug Type":             "Standard",
		"Electrical System":           "12V",
	}

	for key, value := range specs {
		// ensure specification key exists (creates if missing)
		sk, err := CreateOrFindSpecificationKey(db, key)
		if err != nil {
			return err
		}

		// Check if the specification for this product+key already exists
		var existing models.SpecificationModel
		if err := db.Where("product_id = ? AND specification_key_id = ?", productID, sk.ID).First(&existing).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				// create new specification record
				sModel := &models.SpecificationModel{
					ProductID:          productID,
					SpecificationKeyID: sk.ID,
					Value:              value,
					Status:             1,
				}

				if err := db.Create(sModel).Error; err != nil {
					return err
				}
			} else {
				return err
			}
		}
	}

	return nil
}
