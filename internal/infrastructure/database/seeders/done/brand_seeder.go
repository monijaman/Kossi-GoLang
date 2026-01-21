package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// BrandSeeder handles seeding brand names into the database
type BrandSeeder struct {
	BaseSeeder
}

// NewBrandSeeder creates a new seeder for brands
func NewBrandSeeder() *BrandSeeder {
	return &BrandSeeder{
		BaseSeeder: BaseSeeder{name: "Brand Seeder"},
	}
}

// Seed implements the Seeder interface
func (bs *BrandSeeder) Seed(db *gorm.DB) error {
	// Define brand names
	brands := []string{
		"Toyota",
		"Honda",
		"Nissan",
		"Mitsubishi",
		"Mazda",
		"Subaru",
		"Suzuki",
		"Daihatsu",
		"Hyundai",
		"Kia",
		"Proton",
		"BMW",
		"Mercedes-Benz",
		"Audi",
		"Volkswagen",
		"Tata",
		"Ford",
		"Changan",
		"Chery",
		"Chevrolet",
		"Geely",
		"GMC",
		"Great Wall / Haval",
		"Isuzu",
		"Jaguar",
		"Jeep",
		"Lexus",
		"Mahindra",
		"MG",
		"MINI",
		"Peugeot",
		"Porsche",
		"Range Rover",
		"SsangYong",
		"Volvo",
	}

	// Insert brands into the database
	for _, brandName := range brands {
		var existingBrand models.BrandModel
		result := db.Where("name = ?", brandName).First(&existingBrand)

		if result.Error == gorm.ErrRecordNotFound {
			// Create brand if it doesn't exist
			newBrand := models.BrandModel{Name: brandName}
			if err := db.Create(&newBrand).Error; err != nil {
				continue
			}
		}
	}

	return nil
}
