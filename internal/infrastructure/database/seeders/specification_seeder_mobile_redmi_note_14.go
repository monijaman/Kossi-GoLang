package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileRedmiNote14 seeds specifications/options for product 'redmi-note-14'
type SpecificationSeederMobileRedmiNote14 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileRedmiNote14 creates a new seeder instance
func NewSpecificationSeederMobileRedmiNote14() *SpecificationSeederMobileRedmiNote14 {
	return &SpecificationSeederMobileRedmiNote14{BaseSeeder: BaseSeeder{name: "Specifications for Redmi Note 14"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileRedmiNote14) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080 x 2400 pixels": "১০৮০ x ২৪০০ pixels",
		"120Hz": "১২০Hz",
		"128 GB / 256 GB": "১২৮ GB / ২৫৬ GB",
		"16 MP": "১৬ MP",
		"162.4 x 75.7 x 8 mm": "১৬২.৪ x ৭৫.৭ x ৮ মিমি",
		"190 g": "১৯০ g",
		"5,110 mAh": "৫,১১০ এমএএইচ",
		"50 MP + 2 MP": "৫০ MP + ২ MP",
		"5G": "৫G",
		"6 GB / 8 GB / 12 GB": "৬ GB / ৮ GB / ১২ GB",
		"6.67 inches": "৬.৬৭ ইঞ্চি",
		"Android 14, HyperOS": "Android ১৪, HyperOS",
		"Corning Gorilla Glass 5": "Corning Gorilla Glass ৫",
		"Dimensity 7025 আল্ট্রা": "Dimensity ৭০২৫ আল্ট্রা",
		"Glass front, plastic frame, plastic back": "গ্লাস সামনে, plastic frame, প্লাস্টিক পেছনে",
		"IMG BXM-8-256": "IMG BXM-৮-২৫৬",
		"IP64": "IP৬৪",
		"Mediatek Dimensity 7025 আল্ট্রা (6 nm)": "Mediatek Dimensity ৭০২৫ আল্ট্রা (৬ nm)",
		"OLED, 120Hz, 2100 nits": "OLED, ১২০Hz, ২১০০ nits",
		"September 2024": "September ২০২৪",
		"Starry White, মিডনাইট Black, Phantom Blue": "Starry সাদা, মিডনাইট কালো, Phantom নীল",
	}
}

// Seed inserts specification_translations for existing specifications for product 'redmi-note-14'
func (s *SpecificationSeederMobileRedmiNote14) Seed(db *gorm.DB) error {
	productSlug := "redmi-note-14"

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
