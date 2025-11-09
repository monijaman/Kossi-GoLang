package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMotorcycleHondaCB125F seeds specifications/options for product 'honda-cb125f'
type SpecificationSeederMotorcycleHondaCB125F struct {
	BaseSeeder
}

// NewSpecificationSeederMotorcycleHondaCB125F creates a new seeder instance
func NewSpecificationSeederMotorcycleHondaCB125F() *SpecificationSeederMotorcycleHondaCB125F {
	return &SpecificationSeederMotorcycleHondaCB125F{BaseSeeder: BaseSeeder{name: "Specifications for Honda CB125F"}}
}

// Seed inserts specification records for the product identified by slug 'honda-cb125f'
func (s *SpecificationSeederMotorcycleHondaCB125F) Seed(db *gorm.DB) error {
	productSlug := "honda-cb125f"

	var prod models.ProductModel
	if err := db.Where("slug = ?", productSlug).First(&prod).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil
		}
		return err
	}
	productID := prod.ID

	specs := DefaultMotorcycleSpecs()
	// Override model-specific values for Honda CB125F
	specs["ABS (Anti-lock Braking System)"] = "No (CBS)"
	specs["CC (Cubic Capacity)"] = "125 cc"
	specs["Clutch Type"] = "Wet multi-plate"
	specs["Cooling System"] = "Air-cooled"
	specs["Displacement"] = "124.7 cc"
	specs["Drive Type"] = "Chain"
	specs["Engine Type"] = "Single-cylinder, 4-stroke, SI"
	specs["Fuel Capacity"] = "11 L"
	specs["Fuel Tank Capacity"] = "11 L"
	specs["Fuel Type"] = "Petrol"
	specs["Gearbox"] = "Manual"
	specs["Ground Clearance"] = "160 mm"
	specs["Headlight Type"] = "Halogen"
	specs["Horsepower"] = "10.7 HP"
	specs["Ignition Type"] = "Digital CDI"
	specs["Kerb Weight"] = "117 kg (approx, varies by variant)"
	specs["Length"] = "2046 mm"
	specs["Max Power"] = "10.7 HP @ 7500 rpm"
	specs["Max Torque"] = "11 Nm @ 6000 rpm"
	specs["Mileage"] = "55 km/l (approx)"
	specs["Number of Gears"] = "5"
	specs["Number of Seats"] = "2"
	specs["Seating Capacity"] = "2"
	specs["Starting System"] = "Electric & Kick"
	specs["Suspension Type"] = "Telescopic front, Twin shock rear"
	specs["Top Speed"] = "100 km/h"
	specs["Torque"] = "11 Nm"
	specs["Transmission"] = "5-speed"
	specs["Tyre Size"] = "80/100-18 (front), 80/100-18 (rear)"
	specs["Tyre Type"] = "Tubeless"
	specs["Valve Configuration"] = "SOHC"
	specs["Wheelbase"] = "1266 mm"
	specs["Width"] = "737 mm"
	specs["Brake Type"] = "Front Disc / Rear Drum (varies by variant)"
	specs["Front Brake Type"] = "Disc (or Drum in lower variant)"
	specs["Rear Brake Type"] = "Drum"
	specs["Rim Type"] = "Alloy"
	specs["Rim Size"] = "18 inch"
	specs["Seat Height"] = "791 mm"
	specs["Battery Capacity"] = "12V"
	specs["Fuel System"] = "Carburetor"
	specs["Emission Standard"] = "BS6"
	specs["Compression Ratio"] = "9.2:1"
	specs["Instrument Cluster"] = "Analog"
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
