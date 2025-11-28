package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileRedmiNote15 seeds specifications/options for product 'redmi-note-15'
type SpecificationSeederMobileRedmiNote15 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileRedmiNote15 creates a new seeder instance
func NewSpecificationSeederMobileRedmiNote15() *SpecificationSeederMobileRedmiNote15 {
	return &SpecificationSeederMobileRedmiNote15{BaseSeeder: BaseSeeder{name: "Specifications for Redmi Note 15"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileRedmiNote15) getBanglaTranslations() map[string]string {
	return map[string]string{
		"108 MP + 8 MP + 2 MP": "১০৮ MP + ৮ MP + ২ MP",
		"1080 x 2400 pixels": "১০৮০ x ২৪০০ pixels",
		"120Hz": "১২০Hz",
		"128 GB / 256 GB": "১২৮ GB / ২৫৬ GB",
		"16 MP": "১৬ MP",
		"161.1 x 75 x 7.6 mm": "১৬১.১ x ৭৫ x ৭.৬ মিমি",
		"188 g": "১৮৮ g",
		"5,000 mAh": "৫,০০০ এমএএইচ",
		"5G": "৫G",
		"6 GB / 8 GB": "৬ GB / ৮ GB",
		"6.67 inches": "৬.৬৭ ইঞ্চি",
		"AMOLED, 120Hz, 1000 nits": "AMOLED, ১২০Hz, ১০০০ nits",
		"Android 15, HyperOS": "Android ১৫, HyperOS",
		"Black, White, Blue": "কালো, সাদা, নীল",
		"Corning Gorilla Glass 5": "Corning Gorilla Glass ৫",
		"Dimensity 6080": "Dimensity ৬০৮০",
		"Glass front, plastic frame, plastic back": "গ্লাস সামনে, plastic frame, প্লাস্টিক পেছনে",
		"IP54": "IP৫৪",
		"Mali-G57 MC2": "Mali-G৫৭ MC২",
		"Mediatek Dimensity 6080 (6 nm)": "Mediatek Dimensity ৬০৮০ (৬ nm)",
		"September 2025": "September ২০২৫",
	}
}

// Seed inserts specification_translations for existing specifications for product 'redmi-note-15'
func (s *SpecificationSeederMobileRedmiNote15) Seed(db *gorm.DB) error {
	productSlug := "redmi-note-15"

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
