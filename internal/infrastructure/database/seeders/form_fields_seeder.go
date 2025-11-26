package seeders

import (
	"fmt"

	"gorm.io/gorm"
)

type SpecificationFormField struct {
	ID    uint   `gorm:"primaryKey"`
	Label string `gorm:"not null"`
	Type  string `gorm:"not null"`
}

func SeedFormFields(db *gorm.DB) error {
	fields := []SpecificationFormField{
		{Label: "Fuel Efficiency (km/l)", Type: "text"},
		{Label: "Engine Displacement (cc)", Type: "text"},
		{Label: "Max Power (hp)", Type: "text"},
		{Label: "Max Torque (Nm)", Type: "text"},
		{Label: "City Fuel Economy (km/l)", Type: "text"},
		{Label: "Highway Fuel Economy (km/l)", Type: "text"},
		{Label: "Combined Fuel Economy (km/l)", Type: "text"},
		{Label: "Fuel Tank Capacity (l)", Type: "text"},
		{Label: "Length (mm)", Type: "text"},
		{Label: "Width (mm)", Type: "text"},
		{Label: "Height (mm)", Type: "text"},
		{Label: "Wheelbase (mm)", Type: "text"},
		{Label: "AC Type", Type: "text"},
		{Label: "Power Windows", Type: "checkbox"},
		{Label: "Seat Material", Type: "text"},
		{Label: "Driver Airbag", Type: "checkbox"},
		{Label: "Passenger Airbag", Type: "checkbox"},
		{Label: "Side Airbags", Type: "checkbox"},
		{Label: "Curtain Airbags", Type: "checkbox"},
		{Label: "Camera", Type: "checkbox"},
		{Label: "Child Safety Lock", Type: "checkbox"},
		{Label: "Stability Control", Type: "checkbox"},
		{Label: "Service Interval (km)", Type: "text"},
		{Label: "Spare Parts Availability", Type: "checkbox"},
		{Label: "Estimated Annual Maintenance Cost (BDT)", Type: "text"},
		{Label: "Registration Cost (BDT)", Type: "text"},
		{Label: "Insurance Cost (BDT)", Type: "text"},
	}

	for _, field := range fields {
		var existingField SpecificationFormField
		result := db.Where("label = ?", field.Label).First(&existingField)
		if result.Error == gorm.ErrRecordNotFound {
			if err := db.Create(&field).Error; err != nil {
				fmt.Printf("Error seeding field %s: %v\n", field.Label, err)
				continue
			}
			fmt.Printf("Seeded field: %s\n", field.Label)
		} else if result.Error != nil {
			fmt.Printf("Error checking field %s: %v\n", field.Label, result.Error)
		}
	}

	return nil
}
