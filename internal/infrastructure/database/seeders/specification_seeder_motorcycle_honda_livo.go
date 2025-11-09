package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMotorcycleHondaLivo seeds specifications/options for product 'honda-livo'
type SpecificationSeederMotorcycleHondaLivo struct {
	BaseSeeder
}

// NewSpecificationSeederMotorcycleHondaLivo creates a new seeder instance
func NewSpecificationSeederMotorcycleHondaLivo() *SpecificationSeederMotorcycleHondaLivo {
	return &SpecificationSeederMotorcycleHondaLivo{BaseSeeder: BaseSeeder{name: "Specifications for Honda Livo"}}
}

// Seed inserts specification records for the product identified by slug 'honda-livo'
func (s *SpecificationSeederMotorcycleHondaLivo) Seed(db *gorm.DB) error {
	productSlug := "honda-livo"

	var prod models.ProductModel
	if err := db.Where("slug = ?", productSlug).First(&prod).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil
		}
		return err
	}
	productID := prod.ID

	specs := DefaultMotorcycleSpecs()
	// Override model-specific values for Honda Livo
	specs["ABS (Anti-lock Braking System)"] = "No (CBS)"
	specs["CC (Cubic Capacity)"] = "110 cc"
	specs["Clutch Type"] = "Wet multi-plate"
	specs["Cooling System"] = "Air-cooled"
	specs["Displacement"] = "109.1 cc"
	specs["Drive Type"] = "Chain"
	specs["Engine Type"] = "Single-cylinder, 4-stroke, SI"
	specs["Fuel Capacity"] = "9 L"
	specs["Fuel Tank Capacity"] = "9 L"
	specs["Fuel Type"] = "Petrol"
	specs["Gearbox"] = "Manual"
	specs["Ground Clearance"] = "180 mm"
	specs["Headlight Type"] = "Halogen"
	specs["Horsepower"] = "8.4 HP"
	specs["Ignition Type"] = "Digital CDI"
	specs["Kerb Weight"] = "115 kg (approx, varies by variant)"
	specs["Length"] = "2020 mm"
	specs["Max Power"] = "8.4 HP @ 7500 rpm"
	specs["Max Torque"] = "9.09 Nm @ 5000 rpm"
	specs["Mileage"] = "65 km/l (approx)"
	specs["Number of Gears"] = "4"
	specs["Number of Seats"] = "2"
	specs["Seating Capacity"] = "2"
	specs["Starting System"] = "Electric & Kick"
	specs["Suspension Type"] = "Telescopic front, Twin shock rear"
	specs["Top Speed"] = "86 km/h"
	specs["Torque"] = "9.09 Nm"
	specs["Transmission"] = "4-speed"
	specs["Tyre Size"] = "80/100-18 (front), 80/100-18 (rear)"
	specs["Tyre Type"] = "Tube"
	specs["Valve Configuration"] = "SOHC"
	specs["Wheelbase"] = "1285 mm"
	specs["Width"] = "746 mm"
	specs["Brake Type"] = "Front Drum / Rear Drum (CBS)"
	specs["Front Brake Type"] = "Drum"
	specs["Rear Brake Type"] = "Drum"
	specs["Rim Type"] = "Alloy"
	specs["Rim Size"] = "18 inch"
	specs["Seat Height"] = "800 mm"
	specs["Battery Capacity"] = "12V"
	specs["Fuel System"] = "Carburetor"
	specs["Emission Standard"] = "BS6"
	specs["Compression Ratio"] = "9.9:1"
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
