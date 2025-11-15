package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMotorcycleYamahaSalutoRX seeds specifications/options for product 'yamaha-saluto-rx'
type SpecificationSeederMotorcycleYamahaSalutoRX struct {
	BaseSeeder
}

// NewSpecificationSeederMotorcycleYamahaSalutoRX creates a new seeder instance
func NewSpecificationSeederMotorcycleYamahaSalutoRX() *SpecificationSeederMotorcycleYamahaSalutoRX {
	return &SpecificationSeederMotorcycleYamahaSalutoRX{BaseSeeder: BaseSeeder{name: "Specifications for Yamaha Saluto RX"}}
}

// Seed inserts specification records for the product identified by slug 'yamaha-saluto-rx'
func (s *SpecificationSeederMotorcycleYamahaSalutoRX) Seed(db *gorm.DB) error {
	productSlug := "yamaha-saluto-rx"

	var prod models.ProductModel
	if err := db.Where("slug = ?", productSlug).First(&prod).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil
		}
		return err
	}
	productID := prod.ID

	specs := map[string]string{
		"ABS (Anti-lock Braking System)": "No (CBS)",
		"CC (Cubic Capacity)":            "125 cc",
		"Clutch Type":                    "Wet multi-plate",
		"Cooling System":                 "Air-cooled",
		"Displacement":                   "125 cc",
		"Drive Type":                     "Chain",
		"Engine Type":                    "Single-cylinder, 4-stroke, SOHC",
		"Fuel Capacity":                  "11 L",
		"Fuel Tank Capacity":             "11 L",
		"Fuel Type":                      "Petrol",
		"Gearbox":                        "Manual",
		"Ground Clearance":               "180 mm",
		"Headlight Type":                 "Halogen",
		"Horsepower":                     "8.2 HP",
		"Ignition Type":                  "Digital CDI",
		"Kerb Weight":                    "112 kg (approx, varies by variant)",
		"Length":                         "2020 mm",
		"Max Power":                      "8.2 HP @ 7000 rpm",
		"Max Torque":                     "10.1 Nm @ 4500 rpm",
		"Mileage":                        "60-65 km/l (approx)",
		"Number of Gears":                "4",
		"Number of Seats":                "2",
		"Seating Capacity":               "2",
		"Starting System":                "Electric & Kick",
		"Suspension Type":                "Telescopic front, Twin shock rear",
		"Top Speed":                      "90 km/h",
		"Torque":                         "10.1 Nm",
		"Transmission":                   "4-speed",
		"Tyre Size":                      "80/100-17 (front), 100/90-17 (rear)",
		"Tyre Type":                      "Tube",
		"Valve Configuration":            "SOHC",
		"Wheelbase":                      "1265 mm",
		"Width":                          "740 mm",
		"Brake Type":                     "Drum",
		"Front Brake Type":               "Drum",
		"Rear Brake Type":                "Drum",
		"Rim Type":                       "Alloy",
		"Rim Size":                       "17 inch",
		"Seat Height":                    "785 mm",
		"Battery Capacity":               "12V",
		"Fuel System":                    "Carburetor",
		"Emission Standard":              "BS4",
		"Compression Ratio":              "10.0:1",
		"Instrument Cluster":             "Analog",
		"DRL (Daytime Running Light)":    "No",
		"Spark Plug Type":                "Standard",
		"Electrical System":              "12V DC",
		"Bore":                           "",
		"Stroke":                         "",
		"Lubrication System":             "",
		"Oil Pump Type":                  "",
		"Front Suspension Travel":        "",
		"Rear Suspension Travel":         "",
		"Oil Capacity":                   "",
		"Frame Type":                     "",
		"Reserve Fuel Capacity":          "",
		"Side Stand Engine Cutoff":       "",
		"Front Tyre Size":                "",
		"Rear Tyre Size":                 "",
		"Color Options":                  "",
		"Speedometer Type":               "",
		"Battery Type":                   "",
		"Brake Diameter":                 "",
		"Suspension Material":            "",
		"Wheel Type":                     "",
		"Cooling Type":                   "",
	}

	for key, value := range specs {
		sk, err := CreateOrFindSpecificationKey(db, key)
		if err != nil {
			return err
		}

		var existing models.SpecificationModel
		if err := db.Where("product_id = ? AND specification_key_id = ?", productID, sk.ID).First(&existing).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
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
