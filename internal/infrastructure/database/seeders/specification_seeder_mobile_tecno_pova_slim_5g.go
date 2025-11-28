package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileTecnoPovaSlim5g seeds specifications/options for product 'tecno-pova-slim-5g'
type SpecificationSeederMobileTecnoPovaSlim5g struct {
	BaseSeeder
}

// NewSpecificationSeederMobileTecnoPovaSlim5g creates a new seeder instance
func NewSpecificationSeederMobileTecnoPovaSlim5g() *SpecificationSeederMobileTecnoPovaSlim5g {
	return &SpecificationSeederMobileTecnoPovaSlim5g{BaseSeeder: BaseSeeder{name: "Specifications for Tecno POVA Slim 5G"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileTecnoPovaSlim5g) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080 x 2460 pixels": "১০৮০ x ২৪৬০ pixels",
		"120Hz": "১২০Hz",
		"16 MP": "১৬ MP",
		"168.6 x 76.6 x 7.9 mm": "১৬৮.৬ x ৭৬.৬ x ৭.৯ মিমি",
		"198 g": "১৯৮ g",
		"256 GB": "২৫৬ GB",
		"50 MP + 2 MP": "৫০ MP + ২ MP",
		"5G": "৫G",
		"6,000 mAh": "৬,০০০ এমএএইচ",
		"6.78 inches": "৬.৭৮ ইঞ্চি",
		"8 GB": "৮ GB",
		"AMOLED, 120Hz": "AMOLED, ১২০Hz",
		"Android 14, HIOS 14": "Android ১৪, HIOS ১৪",
		"April 2024": "April ২০২৪",
		"Dimensity 6080": "Dimensity ৬০৮০",
		"IP53": "IP৫৩",
		"Mali-G57 MC2": "Mali-G৫৭ MC২",
		"Mecha Black, Amber Gold": "Mecha কালো, Amber সোনালী",
		"Mediatek Dimensity 6080 (6 nm)": "Mediatek Dimensity ৬০৮০ (৬ nm)",
	}
}

// Seed inserts specification_translations for existing specifications for product 'tecno-pova-slim-5g'
func (s *SpecificationSeederMobileTecnoPovaSlim5g) Seed(db *gorm.DB) error {
	productSlug := "tecno-pova-slim-5g"

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
