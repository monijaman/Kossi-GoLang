package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMotorcycleHondaCD70 seeds specifications/options for product 'honda-cd-70'
type SpecificationSeederMotorcycleHondaCD70 struct {
	BaseSeeder
}

// NewSpecificationSeederMotorcycleHondaCD70 creates a new seeder instance
func NewSpecificationSeederMotorcycleHondaCD70() *SpecificationSeederMotorcycleHondaCD70 {
	return &SpecificationSeederMotorcycleHondaCD70{BaseSeeder: BaseSeeder{name: "Specifications for Honda CD 70"}}
}

// Seed inserts specification records for the product identified by slug 'honda-cd-70'
func (s *SpecificationSeederMotorcycleHondaCD70) Seed(db *gorm.DB) error {
	productSlug := "honda-cd-70"

	var prod models.ProductModel
	if err := db.Where("slug = ?", productSlug).First(&prod).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			// product not present, skip
			return nil
		}
		return err
	}
	productID := prod.ID

	specs := DefaultMotorcycleSpecs()
	// Override model-specific values for Honda CD 70
	specs["ABS (Anti-lock Braking System)"] = "No (CBS)"
	specs["CC (Cubic Capacity)"] = "70 cc"
	specs["Clutch Type"] = "Wet multi-plate"
	specs["Cooling System"] = "Air-cooled"
	specs["Displacement"] = "70 cc"
	specs["Drive Type"] = "Chain"
	specs["Engine Type"] = "Single-cylinder, 4-stroke, SI"
	specs["Fuel Capacity"] = "8.5 L"
	specs["Fuel Tank Capacity"] = "8.5 L"
	specs["Fuel Type"] = "Petrol"
	specs["Gearbox"] = "Manual"
	specs["Ground Clearance"] = "165 mm"
	specs["Headlight Type"] = "Halogen"
	specs["Horsepower"] = "6.4 HP"
	specs["Ignition Type"] = "Digital CDI"
	specs["Kerb Weight"] = "82 kg (approx, varies by variant)"
	specs["Length"] = "1897 mm"
	specs["Max Power"] = "6.4 HP @ 7000 rpm"
	specs["Max Torque"] = "5.0 Nm @ 6000 rpm"
	specs["Mileage"] = "60 km/l (approx)"
	specs["Number of Gears"] = "4"
	specs["Number of Seats"] = "2"
	specs["Seating Capacity"] = "2"
	specs["Starting System"] = "Kick"
	specs["Suspension Type"] = "Telescopic front, Twin shock rear"
	specs["Top Speed"] = "80 km/h"
	specs["Torque"] = "5.0 Nm"
	specs["Transmission"] = "4-speed"
	specs["Tyre Size"] = "2.25-17 (front), 2.50-17 (rear)"
	specs["Tyre Type"] = "Tube"
	specs["Valve Configuration"] = "SOHC"
	specs["Wheelbase"] = "1200 mm"
	specs["Width"] = "751 mm"
	specs["Brake Type"] = "Drum"
	specs["Front Brake Type"] = "Drum"
	specs["Rear Brake Type"] = "Drum"
	specs["Rim Type"] = "Alloy"
	specs["Rim Size"] = "17 inch"
	specs["Seat Height"] = "771 mm"
	specs["Battery Capacity"] = "12V"
	specs["Fuel System"] = "Carburetor"
	specs["Emission Standard"] = "Euro 2"
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
