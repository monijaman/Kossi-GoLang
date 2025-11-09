package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMotorcycleYamahaFazer seeds specifications/options for product 'yamaha-fazer'
type SpecificationSeederMotorcycleYamahaFazer struct {
	BaseSeeder
}

// NewSpecificationSeederMotorcycleYamahaFazer creates a new seeder instance
func NewSpecificationSeederMotorcycleYamahaFazer() *SpecificationSeederMotorcycleYamahaFazer {
	return &SpecificationSeederMotorcycleYamahaFazer{BaseSeeder: BaseSeeder{name: "Specifications for Yamaha Fazer"}}
}

// Seed inserts specification records for the product identified by slug 'yamaha-fazer'
func (s *SpecificationSeederMotorcycleYamahaFazer) Seed(db *gorm.DB) error {
	productSlug := "yamaha-fazer"

	var prod models.ProductModel
	if err := db.Where("slug = ?", productSlug).First(&prod).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil
		}
		return err
	}
	productID := prod.ID

	specs := DefaultMotorcycleSpecs()
	// Override model-specific values for Yamaha Fazer
	specs["ABS (Anti-lock Braking System)"] = "No (CBS)"
	specs["CC (Cubic Capacity)"] = "150 cc"
	specs["Clutch Type"] = "Wet multi-plate"
	specs["Cooling System"] = "Air-cooled"
	specs["Displacement"] = "149.0 cc"
	specs["Drive Type"] = "Chain"
	specs["Engine Type"] = "Single-cylinder, 4-stroke, SOHC"
	specs["Fuel Capacity"] = "14 L"
	specs["Fuel Tank Capacity"] = "14 L"
	specs["Fuel Type"] = "Petrol"
	specs["Gearbox"] = "Manual"
	specs["Ground Clearance"] = "160 mm"
	specs["Headlight Type"] = "Halogen"
	specs["Horsepower"] = "13.6 HP"
	specs["Ignition Type"] = "Digital CDI"
	specs["Kerb Weight"] = "135 kg (approx, varies by variant)"
	specs["Length"] = "1975 mm"
	specs["Max Power"] = "13.6 HP @ 8000 rpm"
	specs["Max Torque"] = "12.8 Nm @ 6000 rpm"
	specs["Mileage"] = "40-45 km/l (approx)"
	specs["Number of Gears"] = "5"
	specs["Number of Seats"] = "2"
	specs["Seating Capacity"] = "2"
	specs["Starting System"] = "Electric & Kick"
	specs["Suspension Type"] = "Telescopic front, Twin shock rear"
	specs["Top Speed"] = "118 km/h"
	specs["Torque"] = "12.8 Nm"
	specs["Transmission"] = "5-speed"
	specs["Tyre Size"] = "100/80-17 (front), 140/60R-17 (rear)"
	specs["Tyre Type"] = "Tubeless"
	specs["Valve Configuration"] = "SOHC"
	specs["Wheelbase"] = "1334 mm"
	specs["Width"] = "760 mm"
	specs["Brake Type"] = "Front Disc / Rear Drum (varies by variant)"
	specs["Front Brake Type"] = "Disc"
	specs["Rear Brake Type"] = "Drum"
	specs["Rim Type"] = "Alloy"
	specs["Rim Size"] = "17 inch"
	specs["Seat Height"] = "790 mm"
	specs["Battery Capacity"] = "12V"
	specs["Fuel System"] = "Carburetor"
	specs["Emission Standard"] = "BS4"
	specs["Compression Ratio"] = "9.5:1"
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
