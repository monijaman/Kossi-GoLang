package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMotorcycleSuzukiAccess125 seeds specifications/options for product 'suzuki-access-125'
type SpecificationSeederMotorcycleSuzukiAccess125 struct {
	BaseSeeder
}

// NewSpecificationSeederMotorcycleSuzukiAccess125 creates a new seeder instance
func NewSpecificationSeederMotorcycleSuzukiAccess125() *SpecificationSeederMotorcycleSuzukiAccess125 {
	return &SpecificationSeederMotorcycleSuzukiAccess125{BaseSeeder: BaseSeeder{name: "Specifications for Suzuki Access 125"}}
}

// Seed inserts specification records for the product identified by slug 'suzuki-access-125'
func (s *SpecificationSeederMotorcycleSuzukiAccess125) Seed(db *gorm.DB) error {
	productSlug := "suzuki-access-125"

	var prod models.ProductModel
	if err := db.Where("slug = ?", productSlug).First(&prod).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil
		}
		return err
	}
	productID := prod.ID

	specs := DefaultMotorcycleSpecs()
	// Override model-specific values for Suzuki Access 125
	specs["ABS (Anti-lock Braking System)"] = "No (CBS)"
	specs["CC (Cubic Capacity)"] = "125 cc"
	specs["Clutch Type"] = "Wet multi-plate"
	specs["Cooling System"] = "Air-cooled"
	specs["Displacement"] = "124 cc"
	specs["Drive Type"] = "Belt"
	specs["Engine Type"] = "Single-cylinder, 4-stroke, FI"
	specs["Fuel Capacity"] = "5.6 L"
	specs["Fuel Tank Capacity"] = "5.6 L"
	specs["Fuel Type"] = "Petrol"
	specs["Gearbox"] = "Automatic (CVT)"
	specs["Ground Clearance"] = "160 mm"
	specs["Headlight Type"] = "LED"
	specs["Horsepower"] = "8.7 HP"
	specs["Ignition Type"] = "Digital CDI"
	specs["Kerb Weight"] = "103 kg (approx, varies by variant)"
	specs["Length"] = "1870 mm"
	specs["Max Power"] = "8.7 HP @ 6750 rpm"
	specs["Max Torque"] = "10 Nm @ 5500 rpm"
	specs["Mileage"] = "45-55 km/l (approx)"
	specs["Number of Gears"] = "CVT"
	specs["Number of Seats"] = "2"
	specs["Seating Capacity"] = "2"
	specs["Starting System"] = "Electric"
	specs["Suspension Type"] = "Telescopic front, Single shock rear"
	specs["Top Speed"] = "105 km/h"
	specs["Torque"] = "10 Nm"
	specs["Transmission"] = "CVT"
	specs["Tyre Size"] = "90/90-12 (front), 90/100-10 (rear)"
	specs["Tyre Type"] = "Tubeless"
	specs["Valve Configuration"] = "SOHC"
	specs["Wheelbase"] = "1250 mm"
	specs["Width"] = "690 mm"
	specs["Brake Type"] = "Front Drum / Rear Drum (CBS)"
	specs["Front Brake Type"] = "Drum"
	specs["Rear Brake Type"] = "Drum"
	specs["Rim Type"] = "Alloy"
	specs["Rim Size"] = "12 inch (front), 10 inch (rear)"
	specs["Seat Height"] = "773 mm"
	specs["Battery Capacity"] = "12V"
	specs["Fuel System"] = "Fuel Injection"
	specs["Emission Standard"] = "BS6"
	specs["Compression Ratio"] = "10.3:1"
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
