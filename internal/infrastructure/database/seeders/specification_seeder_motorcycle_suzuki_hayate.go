package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMotorcycleSuzukiHayate seeds specifications/options for product 'suzuki-hayate'
type SpecificationSeederMotorcycleSuzukiHayate struct {
	BaseSeeder
}

// NewSpecificationSeederMotorcycleSuzukiHayate creates a new seeder instance
func NewSpecificationSeederMotorcycleSuzukiHayate() *SpecificationSeederMotorcycleSuzukiHayate {
	return &SpecificationSeederMotorcycleSuzukiHayate{BaseSeeder: BaseSeeder{name: "Specifications for Suzuki Hayate"}}
}

// Seed inserts specification records for the product identified by slug 'suzuki-hayate'
func (s *SpecificationSeederMotorcycleSuzukiHayate) Seed(db *gorm.DB) error {
	productSlug := "suzuki-hayate"

	var prod models.ProductModel
	if err := db.Where("slug = ?", productSlug).First(&prod).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil
		}
		return err
	}
	productID := prod.ID

	specs := DefaultMotorcycleSpecs()
	// Override model-specific values for Suzuki Hayate
	specs["ABS (Anti-lock Braking System)"] = "No (CBS)"
	specs["CC (Cubic Capacity)"] = "110 cc"
	specs["Clutch Type"] = "Wet multi-plate"
	specs["Cooling System"] = "Air-cooled"
	specs["Displacement"] = "113 cc"
	specs["Drive Type"] = "Chain"
	specs["Engine Type"] = "Single-cylinder, 4-stroke, SOHC"
	specs["Fuel Capacity"] = "11 L"
	specs["Fuel Tank Capacity"] = "11 L"
	specs["Fuel Type"] = "Petrol"
	specs["Gearbox"] = "Manual"
	specs["Ground Clearance"] = "165 mm"
	specs["Headlight Type"] = "Halogen"
	specs["Horsepower"] = "8.7 HP"
	specs["Ignition Type"] = "Digital CDI"
	specs["Kerb Weight"] = "114 kg (approx, varies by variant)"
	specs["Length"] = "2025 mm"
	specs["Max Power"] = "8.7 HP @ 7500 rpm"
	specs["Max Torque"] = "9.3 Nm @ 5000 rpm"
	specs["Mileage"] = "50-60 km/l (approx)"
	specs["Number of Gears"] = "4"
	specs["Number of Seats"] = "2"
	specs["Seating Capacity"] = "2"
	specs["Starting System"] = "Kick & Electric"
	specs["Suspension Type"] = "Telescopic front, Twin shock rear"
	specs["Top Speed"] = "90 km/h"
	specs["Torque"] = "9.3 Nm"
	specs["Transmission"] = "4-speed"
	specs["Tyre Size"] = "2.75-17 (front), 3.00-17 (rear)"
	specs["Tyre Type"] = "Tube"
	specs["Valve Configuration"] = "SOHC"
	specs["Wheelbase"] = "1300 mm"
	specs["Width"] = "740 mm"
	specs["Brake Type"] = "Drum"
	specs["Front Brake Type"] = "Drum"
	specs["Rear Brake Type"] = "Drum"
	specs["Rim Type"] = "Alloy"
	specs["Rim Size"] = "17 inch"
	specs["Seat Height"] = "790 mm"
	specs["Battery Capacity"] = "12V"
	specs["Fuel System"] = "Carburetor"
	specs["Emission Standard"] = "BS4"
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
