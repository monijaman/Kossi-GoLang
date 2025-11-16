package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// MobileOperatorBrandSeeder handles seeding mobile operator brands
type MobileOperatorBrandSeeder struct {
	BaseSeeder
}

// NewMobileOperatorBrandSeeder creates a new mobile operator brand seeder
func NewMobileOperatorBrandSeeder() *MobileOperatorBrandSeeder {
	return &MobileOperatorBrandSeeder{
		BaseSeeder: BaseSeeder{name: "Mobile Operator Brands"},
	}
}

// Seed implements the Seeder interface
func (mobs *MobileOperatorBrandSeeder) Seed(db *gorm.DB) error {
	mobileOperators := []struct {
		Name string
		Slug string
	}{
		{Name: "Banglalink", Slug: "banglalink"},
		{Name: "Grameenphone", Slug: "grameenphone"},
		{Name: "Robi", Slug: "robi"},
		{Name: "Teletalk", Slug: "teletalk"},
		{Name: "Airtel", Slug: "airtel"},
	}

	for _, operator := range mobileOperators {
		// Check if brand already exists
		var existingBrand models.BrandModel
		result := db.Where("slug = ?", operator.Slug).First(&existingBrand)

		if result.Error == gorm.ErrRecordNotFound {
			// Create new brand
			brand := models.BrandModel{
				Name: operator.Name,
				Slug: operator.Slug,
			}

			if err := db.Create(&brand).Error; err != nil {
				continue
			}
		}
	}

	return nil
}
