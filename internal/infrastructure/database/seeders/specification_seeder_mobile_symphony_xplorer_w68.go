package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileSymphonyXplorerW68 seeds specifications/options for product 'symphony-xplorer-w68'
type SpecificationSeederMobileSymphonyXplorerW68 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSymphonyXplorerW68 creates a new seeder instance
func NewSpecificationSeederMobileSymphonyXplorerW68() *SpecificationSeederMobileSymphonyXplorerW68 {
	return &SpecificationSeederMobileSymphonyXplorerW68{BaseSeeder: BaseSeeder{name: "Specifications for Symphony Xplorer W68"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSymphonyXplorerW68) getBanglaTranslations() map[string]string {
	return map[string]string{
		"0.3 MP": "০.৩ MP",
		"1,500 mAh": "১,৫০০ এমএএইচ",
		"130 g": "১৩০ g",
		"135 x 67 x 9.5 mm": "১৩৫ x ৬৭ x ৯.৫ মিমি",
		"2013": "২০১৩",
		"3G": "৩G",
		"4 GB": "৪ GB",
		"4.5 inches": "৪.৫ ইঞ্চি",
		"480 x 854 pixels": "৪৮০ x ৮৫৪ pixels",
		"5 MP": "৫ MP",
		"512 MB": "৫১২ MB",
		"60Hz": "৬০Hz",
		"Android 4.2 (Jelly Bean)": "Android ৪.২ (Jelly Bean)",
		"Dual-core 1.2GHz": "Dual-core ১.২GHz",
		"Mali-400": "Mali-৪০০",
		"Mediatek MT6572": "Mediatek MT৬৫৭২",
	}
}

// Seed inserts specification_translations for existing specifications for product 'symphony-xplorer-w68'
func (s *SpecificationSeederMobileSymphonyXplorerW68) Seed(db *gorm.DB) error {
	productSlug := "symphony-xplorer-w68"

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
