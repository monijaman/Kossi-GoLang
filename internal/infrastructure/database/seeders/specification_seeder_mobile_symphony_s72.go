package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileSymphonyS72 seeds specifications/options for product 'symphony-s72'
type SpecificationSeederMobileSymphonyS72 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSymphonyS72 creates a new seeder instance
func NewSpecificationSeederMobileSymphonyS72() *SpecificationSeederMobileSymphonyS72 {
	return &SpecificationSeederMobileSymphonyS72{BaseSeeder: BaseSeeder{name: "Specifications for Symphony S72"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSymphonyS72) getBanglaTranslations() map[string]string {
	return map[string]string{
		"0.08 MP": "০.০৮ MP",
		"1000 contacts": "১০০০ contacts",
		"128 × 160 px": "১২৮ × ১৬০ px",
		"1500 mAh": "১৫০০ এমএএইচ",
		"2.4 inches": "২.৪ ইঞ্চি",
		"2G GSM": "২G GSM",
		"Bluetooth, microUSB, 3.5mm jack": "Bluetooth, microUSB, ৩.৫mm jack",
		"December 2024": "December ২০২৪",
		"Li-Ion (removable)": "লি-আয়ন (removable)",
		"Straw Green, Dew Green, Ocean Blue, Ink Black, Caramel Gold": "Straw সবুজ, Dew সবুজ, Ocean নীল, Ink কালো, Caramel সোনালী",
		"microSD up to 16 GB": "microSD up to ১৬ GB",
	}
}

// Seed inserts specification_translations for existing specifications for product 'symphony-s72'
func (s *SpecificationSeederMobileSymphonyS72) Seed(db *gorm.DB) error {
	productSlug := "symphony-s72"

	var prod models.ProductModel
	if err := db.Where("slug = ?", productSlug).First(&prod).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil
		}
		return err
	}

	productID := prod.ID
	banglaTranslations := s.getBanglaTranslations()

	// Get all existing specifications for this product
	var existingSpecs []models.SpecificationModel
	if err := db.Where("product_id = ?", productID).Find(&existingSpecs).Error; err != nil {
		return err
	}

	// Insert translations for all existing specifications
	for _, spec := range existingSpecs {
		banglaValue, exists := banglaTranslations[spec.Value]
		if exists && banglaValue != "" {
			translation := &models.SpecificationTranslationModel{
				SpecificationID: spec.ID,
				Locale:          "bn",
				Value:           banglaValue,
			}
			// Use OnConflict to ignore if translation already exists
			if err := db.Clauses(clause.OnConflict{DoNothing: true}).Create(translation).Error; err != nil {
				return err
			}
		}
	}

	return nil
}
