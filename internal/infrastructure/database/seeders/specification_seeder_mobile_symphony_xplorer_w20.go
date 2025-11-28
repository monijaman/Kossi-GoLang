package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileSymphonyXplorerW20 seeds specifications/options for product 'symphony-xplorer-w20'
type SpecificationSeederMobileSymphonyXplorerW20 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSymphonyXplorerW20 creates a new seeder instance
func NewSpecificationSeederMobileSymphonyXplorerW20() *SpecificationSeederMobileSymphonyXplorerW20 {
	return &SpecificationSeederMobileSymphonyXplorerW20{BaseSeeder: BaseSeeder{name: "Specifications for Symphony Xplorer W20"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSymphonyXplorerW20) getBanglaTranslations() map[string]string {
	return map[string]string{
		"0.3 MP": "০.৩ MP",
		"1,400 mAh": "১,৪০০ এমএএইচ",
		"120 g": "১২০ g",
		"125 x 64 x 10 mm": "১২৫ x ৬৪ x ১০ মিমি",
		"2 MP": "২ MP",
		"2013": "২০১৩",
		"3G": "৩G",
		"4 GB": "৪ GB",
		"4.0 inches": "৪.০ ইঞ্চি",
		"480 x 800 pixels": "৪৮০ x ৮০০ pixels",
		"512 MB": "৫১২ MB",
		"60Hz": "৬০Hz",
		"Android 4.2 (Jelly Bean)": "Android ৪.২ (Jelly Bean)",
		"Dual-core 1.0GHz": "Dual-core ১.০GHz",
		"Mali-400": "Mali-৪০০",
		"Mediatek MT6572": "Mediatek MT৬৫৭২",
	}
}

// Seed inserts specification_translations for existing specifications for product 'symphony-xplorer-w20'
func (s *SpecificationSeederMobileSymphonyXplorerW20) Seed(db *gorm.DB) error {
	productSlug := "symphony-xplorer-w20"

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
