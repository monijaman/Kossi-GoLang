package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileSymphonyXplorerV55 seeds specifications/options for product 'symphony-xplorer-v55'
type SpecificationSeederMobileSymphonyXplorerV55 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSymphonyXplorerV55 creates a new seeder instance
func NewSpecificationSeederMobileSymphonyXplorerV55() *SpecificationSeederMobileSymphonyXplorerV55 {
	return &SpecificationSeederMobileSymphonyXplorerV55{BaseSeeder: BaseSeeder{name: "Specifications for Symphony Xplorer V55"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSymphonyXplorerV55) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1 GB": "১ GB",
		"140 g": "১৪০ g",
		"145 x 72 x 9 mm": "১৪৫ x ৭২ x ৯ মিমি",
		"2 MP": "২ MP",
		"2,200 mAh": "২,২০০ এমএএইচ",
		"2015": "২০১৫",
		"3G": "৩G",
		"480 x 854 pixels": "৪৮০ x ৮৫৪ pixels",
		"5.0 inches": "৫.০ ইঞ্চি",
		"60Hz": "৬০Hz",
		"8 GB": "৮ GB",
		"8 MP": "৮ MP",
		"Android 5.1 (Lollipop)": "Android ৫.১ (Lollipop)",
		"Mali-400MP2": "Mali-৪০০MP২",
		"Mediatek MT6582": "Mediatek MT৬৫৮২",
		"Quad-core 1.3GHz": "Quad-core ১.৩GHz",
	}
}

// Seed inserts specification_translations for existing specifications for product 'symphony-xplorer-v55'
func (s *SpecificationSeederMobileSymphonyXplorerV55) Seed(db *gorm.DB) error {
	productSlug := "symphony-xplorer-v55"

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
