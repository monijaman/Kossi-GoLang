package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileSymphonyI10 seeds specifications/options for product 'symphony-i10'
type SpecificationSeederMobileSymphonyI10 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSymphonyI10 creates a new seeder instance
func NewSpecificationSeederMobileSymphonyI10() *SpecificationSeederMobileSymphonyI10 {
	return &SpecificationSeederMobileSymphonyI10{BaseSeeder: BaseSeeder{name: "Specifications for Symphony i10"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSymphonyI10) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1 GB": "১ GB",
		"140 g": "১৪০ g",
		"144 x 72 x 9 mm": "১৪৪ x ৭২ x ৯ মিমি",
		"16 GB": "১৬ GB",
		"2,500 mAh": "২,৫০০ এমএএইচ",
		"2016": "২০১৬",
		"3G": "৩G",
		"5.0 inches": "৫.০ ইঞ্চি",
		"60Hz": "৬০Hz",
		"720 x 1280 pixels": "৭২০ x ১২৮০ pixels",
		"8 MP": "৮ MP",
		"Android 6.0 (Marshmallow)": "Android ৬.০ (Marshmallow)",
		"Gold, Grey": "সোনালী, ধূসর",
		"Mali-400MP2": "Mali-৪০০MP২",
		"Mediatek MT6580": "Mediatek MT৬৫৮০",
		"Quad-core 1.3GHz": "Quad-core ১.৩GHz",
	}
}

// Seed inserts specification_translations for existing specifications for product 'symphony-i10'
func (s *SpecificationSeederMobileSymphonyI10) Seed(db *gorm.DB) error {
	productSlug := "symphony-i10"

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
