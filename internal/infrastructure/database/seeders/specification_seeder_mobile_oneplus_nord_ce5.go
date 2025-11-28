package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileOneplusNordCe5 seeds specifications/options for product 'oneplus-nord-ce5'
type SpecificationSeederMobileOneplusNordCe5 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileOneplusNordCe5 creates a new seeder instance
func NewSpecificationSeederMobileOneplusNordCe5() *SpecificationSeederMobileOneplusNordCe5 {
	return &SpecificationSeederMobileOneplusNordCe5{BaseSeeder: BaseSeeder{name: "Specifications for OnePlus Nord CE5"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileOneplusNordCe5) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080 x 2412 pixels": "১০৮০ x ২৪১২ pixels",
		"120Hz": "১২০Hz",
		"128 GB / 256 GB": "১২৮ GB / ২৫৬ GB",
		"16 MP": "১৬ MP",
		"162.7 x 75.5 x 8.2 mm": "১৬২.৭ x ৭৫.৫ x ৮.২ মিমি",
		"186 g": "১৮৬ g",
		"5,500 mAh": "৫,৫০০ এমএএইচ",
		"50 MP + 8 MP": "৫০ MP + ৮ MP",
		"5G": "৫G",
		"6.7 inches": "৬.৭ ইঞ্চি",
		"8 GB / 12 GB": "৮ GB / ১২ GB",
		"AMOLED, 120Hz, HDR10+": "AMOLED, ১২০Hz, HDR১০+",
		"Adreno 720": "Adreno ৭২০",
		"Android 15, OxygenOS 15": "Android ১৫, OxygenOS ১৫",
		"April 2025": "April ২০২৫",
		"Glass front, plastic frame, plastic back": "গ্লাস সামনে, plastic frame, প্লাস্টিক পেছনে",
		"IP54": "IP৫৪",
		"Qualcomm Snapdragon 7 Gen 3 (4 nm)": "Qualcomm Snapdragon ৭ Gen ৩ (৪ nm)",
		"Snapdragon 7 Gen 3": "Snapdragon ৭ Gen ৩",
	}
}

// Seed inserts specification_translations for existing specifications for product 'oneplus-nord-ce5'
func (s *SpecificationSeederMobileOneplusNordCe5) Seed(db *gorm.DB) error {
	productSlug := "oneplus-nord-ce5"

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
