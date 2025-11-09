package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMotorcycleYamahaYBR125 seeds specifications/options for product 'yamaha-ybr-125'
type SpecificationSeederMotorcycleYamahaYBR125 struct {
	BaseSeeder
}

// NewSpecificationSeederMotorcycleYamahaYBR125 creates a new seeder instance
func NewSpecificationSeederMotorcycleYamahaYBR125() *SpecificationSeederMotorcycleYamahaYBR125 {
	return &SpecificationSeederMotorcycleYamahaYBR125{BaseSeeder: BaseSeeder{name: "Specifications for Yamaha YBR 125"}}
}

// Seed inserts specification records for the product identified by slug 'yamaha-ybr-125'
func (s *SpecificationSeederMotorcycleYamahaYBR125) Seed(db *gorm.DB) error {
	productSlug := "yamaha-ybr-125"

	var prod models.ProductModel
	if err := db.Where("slug = ?", productSlug).First(&prod).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil
		}
		return err
	}
	productID := prod.ID

	specs := DefaultMotorcycleSpecs()

	specs["ABS (Anti-lock Braking System)"] = "No"
	specs["CC (Cubic Capacity)"] = "125 cc"
	specs["Clutch Type"] = "Wet multi-plate"
	specs["Cooling System"] = "Air-cooled"
	specs["Displacement"] = "124 cc"
	specs["Drive Type"] = "Chain"
	specs["Engine Type"] = "Single-cylinder, 4-stroke, SOHC"
	specs["Fuel Capacity"] = "13 L"
	specs["Fuel Tank Capacity"] = "13 L"
	specs["Fuel Type"] = "Petrol"
	specs["Gearbox"] = "Manual"
	specs["Ground Clearance"] = "145 mm"
	specs["Headlight Type"] = "Halogen"
	specs["Horsepower"] = "10.7 HP"
	specs["Ignition Type"] = "CDI"
	specs["Kerb Weight"] = "118 kg"
	specs["Length"] = "1990 mm"
	specs["Max Power"] = "10.7 HP @ 7500 rpm"
	specs["Max Torque"] = "10.4 Nm @ 6500 rpm"
	specs["Mileage"] = "45–50 km/l"
	specs["Number of Gears"] = "5"
	specs["Number of Seats"] = "2"
	specs["Seating Capacity"] = "2"
	specs["Starting System"] = "Electric & Kick"
	specs["Suspension Type"] = "Telescopic front, Twin shock rear"
	specs["Top Speed"] = "100 km/h"
	specs["Torque"] = "10.4 Nm"
	specs["Transmission"] = "5-speed"
	specs["Tyre Size"] = "2.75-18 (front), 90/90-18 (rear)"
	specs["Tyre Type"] = "Tube"
	specs["Valve Configuration"] = "SOHC"
	specs["Wheelbase"] = "1295 mm"
	specs["Width"] = "745 mm"

	specs["Brake Type"] = "Front Disc / Rear Drum"
	specs["Front Brake Type"] = "Disc"
	specs["Rear Brake Type"] = "Drum"
	specs["Rim Type"] = "Alloy"
	specs["Rim Size"] = "18 inch"
	specs["Seat Height"] = "785 mm"
	specs["Battery Capacity"] = "12V"
	specs["Fuel System"] = "Fuel Injection (Carburetor in older models)"
	specs["Emission Standard"] = "Euro 3"
	specs["Compression Ratio"] = "10.0:1"
	specs["Instrument Cluster"] = "Analog"
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
