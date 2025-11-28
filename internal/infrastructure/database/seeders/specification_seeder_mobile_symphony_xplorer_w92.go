package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileSymphonyXplorerW92 seeds specifications/options for product 'symphony-xplorer-w92'
type SpecificationSeederMobileSymphonyXplorerW92 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSymphonyXplorerW92 creates a new seeder instance
func NewSpecificationSeederMobileSymphonyXplorerW92() *SpecificationSeederMobileSymphonyXplorerW92 {
	return &SpecificationSeederMobileSymphonyXplorerW92{BaseSeeder: BaseSeeder{name: "Specifications for Symphony Xplorer W92"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSymphonyXplorerW92) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1 GB": "১ GB",
		"145 g": "১৪৫ g",
		"145 x 73 x 9 mm": "১৪৫ x ৭৩ x ৯ মিমি",
		"2 MP": "২ MP",
		"2,000 mAh": "২,০০০ এমএএইচ",
		"2013": "২০১৩",
		"3G": "৩G",
		"4 GB": "৪ GB",
		"480 x 854 pixels": "৪৮০ x ৮৫৪ pixels",
		"5 MP": "৫ MP",
		"5.0 inches": "৫.০ ইঞ্চি",
		"60Hz": "৬০Hz",
		"Android 4.2 (Jelly Bean)": "Android ৪.২ (Jelly Bean)",
		"Black, White": "কালো, সাদা",
		"Mediatek MT6589": "Mediatek MT৬৫৮৯",
		"PowerVR SGX544": "PowerVR SGX৫৪৪",
		"Quad-core 1.2GHz": "Quad-core ১.২GHz",
	}
}

// Seed inserts specification_translations for existing specifications for product 'symphony-xplorer-w92'
func (s *SpecificationSeederMobileSymphonyXplorerW92) Seed(db *gorm.DB) error {
	productSlug := "symphony-xplorer-w92"

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
