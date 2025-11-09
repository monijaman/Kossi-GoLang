package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMotorcycleSuzukiGixxerSF seeds specifications/options for product 'suzuki-gixxer-sf'
type SpecificationSeederMotorcycleSuzukiGixxerSF struct {
	BaseSeeder
}

// NewSpecificationSeederMotorcycleSuzukiGixxerSF creates a new seeder instance
func NewSpecificationSeederMotorcycleSuzukiGixxerSF() *SpecificationSeederMotorcycleSuzukiGixxerSF {
	return &SpecificationSeederMotorcycleSuzukiGixxerSF{BaseSeeder: BaseSeeder{name: "Specifications for Suzuki Gixxer SF"}}
}

// Seed inserts specification records for the product identified by slug 'suzuki-gixxer-sf'
func (s *SpecificationSeederMotorcycleSuzukiGixxerSF) Seed(db *gorm.DB) error {
	productSlug := "suzuki-gixxer-sf"

	var prod models.ProductModel
	if err := db.Where("slug = ?", productSlug).First(&prod).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil
		}
		return err
	}
	productID := prod.ID

	specs := DefaultMotorcycleSpecs()
	// Override model-specific values for Suzuki Gixxer SF
	specs["ABS (Anti-lock Braking System)"] = "Single Channel"
	specs["CC (Cubic Capacity)"] = "155 cc"
	specs["Clutch Type"] = "Wet multi-plate"
	specs["Cooling System"] = "Air-cooled"
	specs["Displacement"] = "155 cc"
	specs["Drive Type"] = "Chain"
	specs["Engine Type"] = "Single-cylinder, 4-stroke, SOHC"
	specs["Fuel Capacity"] = "12 L"
	specs["Fuel Tank Capacity"] = "12 L"
	specs["Fuel Type"] = "Petrol"
	specs["Gearbox"] = "Manual"
	specs["Ground Clearance"] = "165 mm"
	specs["Headlight Type"] = "LED"
	specs["Horsepower"] = "13.6 HP"
	specs["Ignition Type"] = "Digital CDI"
	specs["Kerb Weight"] = "146 kg (approx, varies by variant)"
	specs["Length"] = "2025 mm"
	specs["Max Power"] = "13.6 HP @ 8000 rpm"
	specs["Max Torque"] = "13.8 Nm @ 6000 rpm"
	specs["Mileage"] = "38-48 km/l (approx)"
	specs["Number of Gears"] = "5"
	specs["Number of Seats"] = "2"
	specs["Seating Capacity"] = "2"
	specs["Starting System"] = "Electric & Kick"
	specs["Suspension Type"] = "Telescopic front, Twin shock rear"
	specs["Top Speed"] = "120 km/h"
	specs["Torque"] = "13.8 Nm"
	specs["Transmission"] = "5-speed"
	specs["Tyre Size"] = "100/80-17 (front), 140/60R-17 (rear)"
	specs["Tyre Type"] = "Tubeless"
	specs["Valve Configuration"] = "SOHC"
	specs["Wheelbase"] = "1330 mm"
	specs["Width"] = "795 mm"
	specs["Brake Type"] = "Front Disc / Rear Drum (varies by variant)"
	specs["Front Brake Type"] = "Disc"
	specs["Rear Brake Type"] = "Drum"
	specs["Rim Type"] = "Alloy"
	specs["Rim Size"] = "17 inch"
	specs["Seat Height"] = "795 mm"
	specs["Battery Capacity"] = "12V"
	specs["Fuel System"] = "Fuel Injection"
	specs["Emission Standard"] = "BS6"
	specs["Compression Ratio"] = "9.8:1"
	specs["Instrument Cluster"] = "Digital"
	specs["DRL (Daytime Running Light)"] = "Yes"
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
