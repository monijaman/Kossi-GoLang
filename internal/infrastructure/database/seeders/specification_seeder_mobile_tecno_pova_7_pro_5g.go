package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileTecnoPova7Pro5g seeds specifications/options for product 'tecno-pova-7-pro-5g'
type SpecificationSeederMobileTecnoPova7Pro5g struct {
	BaseSeeder
}

// NewSpecificationSeederMobileTecnoPova7Pro5g creates a new seeder instance
func NewSpecificationSeederMobileTecnoPova7Pro5g() *SpecificationSeederMobileTecnoPova7Pro5g {
	return &SpecificationSeederMobileTecnoPova7Pro5g{BaseSeeder: BaseSeeder{name: "Specifications for Tecno POVA 7 Pro 5G"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileTecnoPova7Pro5g) getBanglaTranslations() map[string]string {
	return map[string]string{
		"108 MP + 2 MP": "১০৮ MP + ২ MP",
		"1080 x 2436 pixels": "১০৮০ x ২৪৩৬ pixels",
		"120Hz": "১২০Hz",
		"168.6 x 76.6 x 7.9 mm": "১৬৮.৬ x ৭৬.৬ x ৭.৯ মিমি",
		"198 g": "১৯৮ g",
		"256 GB": "২৫৬ GB",
		"32 MP": "৩২ MP",
		"5G": "৫G",
		"6,000 mAh": "৬,০০০ এমএএইচ",
		"6.78 inches": "৬.৭৮ ইঞ্চি",
		"8 GB / 12 GB": "৮ GB / ১২ GB",
		"AMOLED, 120Hz": "AMOLED, ১২০Hz",
		"Android 14, HIOS 14": "Android ১৪, HIOS ১৪",
		"Dimensity 6080": "Dimensity ৬০৮০",
		"February 2024": "February ২০২৪",
		"IP53": "IP৫৩",
		"Mali-G57 MC2": "Mali-G৫৭ MC২",
		"Mediatek Dimensity 6080 (6 nm)": "Mediatek Dimensity ৬০৮০ (৬ nm)",
		"Meteorite Grey, Comet Green": "Meteorite ধূসর, Comet সবুজ",
	}
}

// Seed inserts specification_translations for existing specifications for product 'tecno-pova-7-pro-5g'
func (s *SpecificationSeederMobileTecnoPova7Pro5g) Seed(db *gorm.DB) error {
	productSlug := "tecno-pova-7-pro-5g"

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
