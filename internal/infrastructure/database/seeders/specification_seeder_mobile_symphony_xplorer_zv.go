package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileSymphonyXplorerZv seeds specifications/options for product 'symphony-xplorer-zv'
type SpecificationSeederMobileSymphonyXplorerZv struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSymphonyXplorerZv creates a new seeder instance
func NewSpecificationSeederMobileSymphonyXplorerZv() *SpecificationSeederMobileSymphonyXplorerZv {
	return &SpecificationSeederMobileSymphonyXplorerZv{BaseSeeder: BaseSeeder{name: "Specifications for Symphony Xplorer ZV"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSymphonyXplorerZv) getBanglaTranslations() map[string]string {
	return map[string]string{
		"13 MP": "১৩ MP",
		"130 g": "১৩০ g",
		"142.5 x 70.8 x 6.8 mm": "১৪২.৫ x ৭০.৮ x ৬.৮ মিমি",
		"16 GB": "১৬ GB",
		"2 GB": "২ GB",
		"2 MP": "২ MP",
		"2,050 mAh": "২,০৫০ এমএএইচ",
		"2014": "২০১৪",
		"3G": "৩G",
		"5.0 inches": "৫.০ ইঞ্চি",
		"60Hz": "৬০Hz",
		"720 x 1280 pixels": "৭২০ x ১২৮০ pixels",
		"Android 4.4.2 (KitKat)": "Android ৪.৪.২ (KitKat)",
		"Black, White": "কালো, সাদা",
		"Gorilla Glass 3": "Gorilla Glass ৩",
		"Mali-450MP4": "Mali-৪৫০MP৪",
		"Mediatek MT6592": "Mediatek MT৬৫৯২",
		"Octa-core 1.4GHz": "Octa-core ১.৪GHz",
	}
}

// Seed inserts specification_translations for existing specifications for product 'symphony-xplorer-zv'
func (s *SpecificationSeederMobileSymphonyXplorerZv) Seed(db *gorm.DB) error {
	productSlug := "symphony-xplorer-zv"

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
