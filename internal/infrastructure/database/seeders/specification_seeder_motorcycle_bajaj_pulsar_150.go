package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMotorcycleBajajPulsar150 seeds specifications/options for product 'bajaj-pulsar-150'
type SpecificationSeederMotorcycleBajajPulsar150 struct {
	BaseSeeder
}

// NewSpecificationSeederMotorcycleBajajPulsar150 creates a new seeder instance
func NewSpecificationSeederMotorcycleBajajPulsar150() *SpecificationSeederMotorcycleBajajPulsar150 {
	return &SpecificationSeederMotorcycleBajajPulsar150{BaseSeeder: BaseSeeder{name: "Specifications for Bajaj Pulsar 150"}}
}

// Seed inserts specification records for the product identified by slug 'bajaj-pulsar-150'
func (s *SpecificationSeederMotorcycleBajajPulsar150) Seed(db *gorm.DB) error {
	productSlug := "bajaj-pulsar-150"

	var prod models.ProductModel
	if err := db.Where("slug = ?", productSlug).First(&prod).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil
		}
		return err
	}
	productID := prod.ID

	specs := DefaultMotorcycleSpecs()
	// Override model-specific values for Bajaj Pulsar 150
	specs["ABS (Anti-lock Braking System)"] = "No (CBS)"
	specs["CC (Cubic Capacity)"] = "150 cc"
	specs["Clutch Type"] = "Wet multi-plate"
	specs["Cooling System"] = "Air-cooled"
	specs["Displacement"] = "149.5 cc"
	specs["Drive Type"] = "Chain"
	specs["Engine Type"] = "Single-cylinder, 4-stroke, DTS-i"
	specs["Fuel Capacity"] = "15 L"
	specs["Fuel Tank Capacity"] = "15 L"
	specs["Fuel Type"] = "Petrol"
	specs["Gearbox"] = "Manual"
	specs["Ground Clearance"] = "165 mm"
	specs["Headlight Type"] = "Halogen"
	specs["Horsepower"] = "14.0 HP"
	specs["Ignition Type"] = "Digital CDI"
	specs["Kerb Weight"] = "144 kg (approx, varies by variant)"
	specs["Length"] = "2055 mm"
	specs["Max Power"] = "14.0 HP @ 8500 rpm"
	specs["Max Torque"] = "13.6 Nm @ 6500 rpm"
	specs["Mileage"] = "45-50 km/l (approx)"
	specs["Number of Gears"] = "5"
	specs["Number of Seats"] = "2"
	specs["Seating Capacity"] = "2"
	specs["Starting System"] = "Electric & Kick"
	specs["Suspension Type"] = "Telescopic front, Twin shock rear"
	specs["Top Speed"] = "115 km/h"
	specs["Torque"] = "13.6 Nm"
	specs["Transmission"] = "5-speed"
	specs["Tyre Size"] = "80/100-17 (front), 100/90-17 (rear)"
	specs["Tyre Type"] = "Tubeless"
	specs["Valve Configuration"] = "SOHC"
	specs["Wheelbase"] = "1320 mm"
	specs["Width"] = "765 mm"
	specs["Brake Type"] = "Front Disc / Rear Drum (varies by variant)"
	specs["Front Brake Type"] = "Disc (or Drum in lower variant)"
	specs["Rear Brake Type"] = "Drum"
	specs["Rim Type"] = "Alloy"
	specs["Rim Size"] = "17 inch"
	specs["Seat Height"] = "790 mm"
	specs["Battery Capacity"] = "12V"
	specs["Fuel System"] = "Fuel Injection"
	specs["Emission Standard"] = "BS6"
	specs["Compression Ratio"] = "9.8:1"
	specs["Instrument Cluster"] = "Analog-Digital"
	specs["DRL (Daytime Running Light)"] = "No"
	specs["Spark Plug Type"] = "Standard"
	specs["Electrical System"] = "12V DC"

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
