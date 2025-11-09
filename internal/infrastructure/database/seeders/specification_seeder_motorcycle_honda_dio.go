package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMotorcycleHondaDio seeds specifications/options for product 'honda-dio'
type SpecificationSeederMotorcycleHondaDio struct {
	BaseSeeder
}

// NewSpecificationSeederMotorcycleHondaDio creates a new seeder instance
func NewSpecificationSeederMotorcycleHondaDio() *SpecificationSeederMotorcycleHondaDio {
	return &SpecificationSeederMotorcycleHondaDio{BaseSeeder: BaseSeeder{name: "Specifications for Honda Dio"}}
}

// Seed inserts specification records for the product identified by slug 'honda-dio'
func (s *SpecificationSeederMotorcycleHondaDio) Seed(db *gorm.DB) error {
	productSlug := "honda-dio"

	var prod models.ProductModel
	if err := db.Where("slug = ?", productSlug).First(&prod).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil
		}
		return err
	}
	productID := prod.ID

	specs := DefaultMotorcycleSpecs()
	// Override model-specific values for Honda Dio
	specs["ABS (Anti-lock Braking System)"] = "No (CBS)"
	specs["CC (Cubic Capacity)"] = "110 cc"
	specs["Clutch Type"] = "Wet multi-plate"
	specs["Cooling System"] = "Air-cooled"
	specs["Displacement"] = "109.51 cc"
	specs["Drive Type"] = "Belt"
	specs["Engine Type"] = "Single-cylinder, 4-stroke, SI"
	specs["Fuel Capacity"] = "5.3 L"
	specs["Fuel Tank Capacity"] = "5.3 L"
	specs["Fuel Type"] = "Petrol"
	specs["Gearbox"] = "Automatic (CVT)"
	specs["Ground Clearance"] = "160 mm"
	specs["Headlight Type"] = "LED"
	specs["Horsepower"] = "7.7 HP"
	specs["Ignition Type"] = "Digital CDI"
	specs["Kerb Weight"] = "105 kg (approx, varies by variant)"
	specs["Length"] = "1808 mm"
	specs["Max Power"] = "7.7 HP @ 8000 rpm"
	specs["Max Torque"] = "9 Nm @ 4750 rpm"
	specs["Mileage"] = "45 km/l (approx)"
	specs["Number of Gears"] = "CVT"
	specs["Number of Seats"] = "2"
	specs["Seating Capacity"] = "2"
	specs["Starting System"] = "Electric"
	specs["Suspension Type"] = "Telescopic front, Single shock rear"
	specs["Top Speed"] = "83 km/h"
	specs["Torque"] = "9 Nm"
	specs["Transmission"] = "CVT"
	specs["Tyre Size"] = "90/90-12 (front), 90/100-10 (rear)"
	specs["Tyre Type"] = "Tubeless"
	specs["Valve Configuration"] = "SOHC"
	specs["Wheelbase"] = "1260 mm"
	specs["Width"] = "710 mm"
	specs["Brake Type"] = "Front Drum / Rear Drum (CBS)"
	specs["Front Brake Type"] = "Drum"
	specs["Rear Brake Type"] = "Drum"
	specs["Rim Type"] = "Alloy"
	specs["Rim Size"] = "12 inch (front), 10 inch (rear)"
	specs["Seat Height"] = "650 mm"
	specs["Battery Capacity"] = "12V"
	specs["Fuel System"] = "Carburetor"
	specs["Emission Standard"] = "BS6"
	specs["Compression Ratio"] = "10.0:1"
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
