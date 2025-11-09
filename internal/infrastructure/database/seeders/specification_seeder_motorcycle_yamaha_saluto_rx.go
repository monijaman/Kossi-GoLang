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

	specs := DefaultMotorcycleSpecs()
	// Override model-specific values for Yamaha Saluto RX
	specs["ABS (Anti-lock Braking System)"] = "No (CBS)"
	specs["CC (Cubic Capacity)"] = "125 cc"
	specs["Clutch Type"] = "Wet multi-plate"
	specs["Cooling System"] = "Air-cooled"
	specs["Displacement"] = "125 cc"
	specs["Drive Type"] = "Chain"
	specs["Engine Type"] = "Single-cylinder, 4-stroke, SOHC"
	specs["Fuel Capacity"] = "11 L"
	specs["Fuel Tank Capacity"] = "11 L"
	specs["Fuel Type"] = "Petrol"
	specs["Gearbox"] = "Manual"
	specs["Ground Clearance"] = "180 mm"
	specs["Headlight Type"] = "Halogen"
	specs["Horsepower"] = "8.2 HP"
	specs["Ignition Type"] = "Digital CDI"
	specs["Kerb Weight"] = "112 kg (approx, varies by variant)"
	specs["Length"] = "2020 mm"
	specs["Max Power"] = "8.2 HP @ 7000 rpm"
	specs["Max Torque"] = "10.1 Nm @ 4500 rpm"
	specs["Mileage"] = "60-65 km/l (approx)"
	specs["Number of Gears"] = "4"
	specs["Number of Seats"] = "2"
	specs["Seating Capacity"] = "2"
	specs["Starting System"] = "Electric & Kick"
	specs["Suspension Type"] = "Telescopic front, Twin shock rear"
	specs["Top Speed"] = "90 km/h"
	specs["Torque"] = "10.1 Nm"
	specs["Transmission"] = "4-speed"
	specs["Tyre Size"] = "80/100-17 (front), 100/90-17 (rear)"
	specs["Tyre Type"] = "Tube"
	specs["Valve Configuration"] = "SOHC"
	specs["Wheelbase"] = "1265 mm"
	specs["Width"] = "740 mm"
	specs["Brake Type"] = "Drum"
	specs["Front Brake Type"] = "Drum"
	specs["Rear Brake Type"] = "Drum"
	specs["Rim Type"] = "Alloy"
	specs["Rim Size"] = "17 inch"
	specs["Seat Height"] = "785 mm"
	specs["Battery Capacity"] = "12V"
	specs["Fuel System"] = "Carburetor"
	specs["Emission Standard"] = "BS4"
	specs["Compression Ratio"] = "10.0:1"
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
