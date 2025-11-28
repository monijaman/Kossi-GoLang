package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileOneplusNord4 seeds specifications/options for product 'oneplus-nord-4'
type SpecificationSeederMobileOneplusNord4 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileOneplusNord4 creates a new seeder instance
func NewSpecificationSeederMobileOneplusNord4() *SpecificationSeederMobileOneplusNord4 {
	return &SpecificationSeederMobileOneplusNord4{BaseSeeder: BaseSeeder{name: "Specifications for OnePlus Nord 4"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileOneplusNord4) getBanglaTranslations() map[string]string {
	return map[string]string{
		"120Hz": "১২০Hz",
		"1240 x 2772 pixels": "১২৪০ x ২৭৭২ pixels",
		"128 GB / 256 GB / 512 GB": "১২৮ GB / ২৫৬ GB / ৫১২ GB",
		"16 MP": "১৬ MP",
		"162.6 x 75 x 8 mm": "১৬২.৬ x ৭৫ x ৮ মিমি",
		"199.5 g": "১৯৯.৫ g",
		"5,500 mAh": "৫,৫০০ এমএএইচ",
		"50 MP + 8 MP": "৫০ MP + ৮ MP",
		"5G": "৫G",
		"6.74 inches": "৬.৭৪ ইঞ্চি",
		"8 GB / 12 GB / 16 GB": "৮ GB / ১২ GB / ১৬ GB",
		"Adreno 732": "Adreno ৭৩২",
		"Android 14, OxygenOS 14.1": "Android ১৪, OxygenOS ১৪.১",
		"Fluid AMOLED, 120Hz, HDR10+, 2150 nits": "Fluid AMOLED, ১২০Hz, HDR১০+, ২১৫০ nits",
		"Glass front, aluminum unibody": "গ্লাস সামনে, aluminum unibody",
		"IP65": "IP৬৫",
		"July 2024": "July ২০২৪",
		"Mercurial Silver, Oasis Green, Obsidian Midnight": "Mercurial রূপালী, Oasis সবুজ, Obsidian Midnight",
		"Qualcomm Snapdragon 7+ Gen 3 (4 nm)": "Qualcomm Snapdragon ৭+ Gen ৩ (৪ nm)",
		"Snapdragon 7+ Gen 3": "Snapdragon ৭+ Gen ৩",
	}
}

// Seed inserts specification_translations for existing specifications for product 'oneplus-nord-4'
func (s *SpecificationSeederMobileOneplusNord4) Seed(db *gorm.DB) error {
	productSlug := "oneplus-nord-4"

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
