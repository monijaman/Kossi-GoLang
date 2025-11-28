package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileOneplusNordCe4Lite5g seeds specifications/options for product 'oneplus-nord-ce4-lite-5g'
type SpecificationSeederMobileOneplusNordCe4Lite5g struct {
	BaseSeeder
}

// NewSpecificationSeederMobileOneplusNordCe4Lite5g creates a new seeder instance
func NewSpecificationSeederMobileOneplusNordCe4Lite5g() *SpecificationSeederMobileOneplusNordCe4Lite5g {
	return &SpecificationSeederMobileOneplusNordCe4Lite5g{BaseSeeder: BaseSeeder{name: "Specifications for OnePlus Nord CE4 Lite 5G"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileOneplusNordCe4Lite5g) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080 x 2400 pixels": "১০৮০ x ২৪০০ pixels",
		"120Hz": "১২০Hz",
		"128 GB / 256 GB": "১২৮ GB / ২৫৬ GB",
		"16 MP": "১৬ MP",
		"162.9 x 75.6 x 8.1 mm": "১৬২.৯ x ৭৫.৬ x ৮.১ মিমি",
		"191 g": "১৯১ g",
		"5,110 mAh": "৫,১১০ এমএএইচ",
		"50 MP + 2 MP": "৫০ MP + ২ MP",
		"5G": "৫G",
		"6.67 inches": "৬.৬৭ ইঞ্চি",
		"8 GB": "৮ GB",
		"AMOLED, 120Hz, 2100 nits": "AMOLED, ১২০Hz, ২১০০ nits",
		"Adreno 619": "Adreno ৬১৯",
		"Android 14, OxygenOS 14": "Android ১৪, OxygenOS ১৪",
		"Glass front, plastic frame, plastic back": "গ্লাস সামনে, plastic frame, প্লাস্টিক পেছনে",
		"IP54": "IP৫৪",
		"June 2024": "June ২০২৪",
		"Qualcomm SM6375 Snapdragon 695 5G (6 nm)": "Qualcomm SM৬৩৭৫ Snapdragon ৬৯৫ ৫G (৬ nm)",
		"Snapdragon 695 5G": "Snapdragon ৬৯৫ ৫G",
		"Super Silver, Mega Blue, Ultra Orange": "Super রূপালী, Mega নীল, Ultra Orange",
	}
}

// Seed inserts specification_translations for existing specifications for product 'oneplus-nord-ce4-lite-5g'
func (s *SpecificationSeederMobileOneplusNordCe4Lite5g) Seed(db *gorm.DB) error {
	productSlug := "oneplus-nord-ce4-lite-5g"

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
