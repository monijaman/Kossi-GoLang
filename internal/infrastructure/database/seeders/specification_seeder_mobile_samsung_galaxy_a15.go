package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileSamsungGalaxyA15 seeds specifications/options for product 'samsung-galaxy-a15'
type SpecificationSeederMobileSamsungGalaxyA15 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSamsungGalaxyA15 creates a new seeder instance
func NewSpecificationSeederMobileSamsungGalaxyA15() *SpecificationSeederMobileSamsungGalaxyA15 {
	return &SpecificationSeederMobileSamsungGalaxyA15{BaseSeeder: BaseSeeder{name: "Specifications for Samsung Galaxy A15"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSamsungGalaxyA15) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080 x 2340 pixels": "১০৮০ x ২৩৪০ pixels",
		"128 GB / 256 GB": "১২৮ GB / ২৫৬ GB",
		"13 MP": "১৩ MP",
		"160.1 x 76.8 x 8.4 mm": "১৬০.১ x ৭৬.৮ x ৮.৪ মিমি",
		"200 g": "২০০ g",
		"25W wired": "২৫W তারযুক্ত",
		"4 GB / 6 GB / 8 GB": "৪ GB / ৬ GB / ৮ GB",
		"50 MP + 5 MP + 2 MP": "৫০ MP + ৫ MP + ২ MP",
		"5000 mAh": "৫০০০ এমএএইচ",
		"6.5 inches": "৬.৫ ইঞ্চি",
		"90Hz": "৯০Hz",
		"Android 14, One UI 6": "Android ১৪, One UI ৬",
		"Brave Black, Optimistic Blue, Magical Blue, Personality Yellow": "Brave কালো, Optimistic নীল, Magical নীল, Personality হলুদ",
		"December 2023": "December ২০২৩",
		"Glass front, plastic back, plastic frame": "গ্লাস সামনে, প্লাস্টিক পেছনে, plastic frame",
		"Mali-G57 MC2": "Mali-G৫৭ MC২",
		"Mediatek Helio G99": "Mediatek Helio G৯৯",
		"Mediatek Helio G99 (6nm)": "Mediatek Helio G৯৯ (৬nm)",
		"Super AMOLED, 90Hz, 800 nits": "Super AMOLED, ৯০Hz, ৮০০ nits",
	}
}

// Seed inserts specification_translations for existing specifications for product 'samsung-galaxy-a15'
func (s *SpecificationSeederMobileSamsungGalaxyA15) Seed(db *gorm.DB) error {
	productSlug := "samsung-galaxy-a15"

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
