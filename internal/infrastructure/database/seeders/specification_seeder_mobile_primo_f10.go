package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobilePrimoF10 seeds specifications/options for product 'primo-f10'
type SpecificationSeederMobilePrimoF10 struct {
	BaseSeeder
}

// NewSpecificationSeederMobilePrimoF10 creates a new seeder instance
func NewSpecificationSeederMobilePrimoF10() *SpecificationSeederMobilePrimoF10 {
	return &SpecificationSeederMobilePrimoF10{BaseSeeder: BaseSeeder{name: "Specifications for Primo F10"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobilePrimoF10) getBanglaTranslations() map[string]string {
	return map[string]string{
		"160 x 77 x 9.5 mm": "১৬০ x ৭৭ x ৯.৫ মিমি",
		"170 g": "১৭০ g",
		"2 GB": "২ GB",
		"3,000 mAh": "৩,০০০ এমএএইচ",
		"32 GB": "৩২ GB",
		"480 x 960 pixels": "৪৮০ x ৯৬০ pixels",
		"4G": "৪G",
		"5 MP": "৫ MP",
		"6.0 inches": "৬.০ ইঞ্চি",
		"60Hz": "৬০Hz",
		"8 MP": "৮ MP",
		"Android 11 Go Edition": "Android ১১ Go Edition",
		"Blue, Black": "নীল, কালো",
		"IMG8322": "IMG৮৩২২",
		"July 2021": "July ২০২১",
		"Unisoc SC9863A": "Unisoc SC৯৮৬৩A",
		"Unisoc SC9863A (28 nm)": "Unisoc SC৯৮৬৩A (২৮ nm)",
	}
}

// Seed inserts specification_translations for existing specifications for product 'primo-f10'
func (s *SpecificationSeederMobilePrimoF10) Seed(db *gorm.DB) error {
	productSlug := "primo-f10"

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
