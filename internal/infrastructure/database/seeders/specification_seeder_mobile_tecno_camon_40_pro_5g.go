package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileTecnoCamon40Pro5g seeds specifications/options for product 'tecno-camon-40-pro-5g'
type SpecificationSeederMobileTecnoCamon40Pro5g struct {
	BaseSeeder
}

// NewSpecificationSeederMobileTecnoCamon40Pro5g creates a new seeder instance
func NewSpecificationSeederMobileTecnoCamon40Pro5g() *SpecificationSeederMobileTecnoCamon40Pro5g {
	return &SpecificationSeederMobileTecnoCamon40Pro5g{BaseSeeder: BaseSeeder{name: "Specifications for Tecno CAMON 40 Pro 5G"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileTecnoCamon40Pro5g) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080 x 2436 pixels": "১০৮০ x ২৪৩৬ pixels",
		"12 GB": "১২ GB",
		"120Hz": "১২০Hz",
		"164 x 74.5 x 7.7 mm": "১৬৪ x ৭৪.৫ x ৭.৭ মিমি",
		"200 g": "২০০ g",
		"5,000 mAh": "৫,০০০ এমএএইচ",
		"50 MP": "৫০ MP",
		"50 MP + 50 MP + 2 MP": "৫০ MP + ৫০ MP + ২ MP",
		"512 GB": "৫১২ GB",
		"5G": "৫G",
		"6.8 inches": "৬.৮ ইঞ্চি",
		"AMOLED, 120Hz": "AMOLED, ১২০Hz",
		"Android 14, HIOS 14": "Android ১৪, HIOS ১৪",
		"Corning Gorilla Glass 5": "Corning Gorilla Glass ৫",
		"Dimensity 8200 Ultimate": "Dimensity ৮২০০ Ultimate",
		"February 2024": "February ২০২৪",
		"Glass front, glass back": "গ্লাস সামনে, গ্লাস পেছনে",
		"IP54": "IP৫৪",
		"Iceland Basalt, Alps Snowy Silver": "Iceland Basalt, Alps Snowy রূপালী",
		"Mali-G610 MC6": "Mali-G৬১০ MC৬",
		"Mediatek Dimensity 8200 Ultimate (4 nm)": "Mediatek Dimensity ৮২০০ Ultimate (৪ nm)",
	}
}

// Seed inserts specification_translations for existing specifications for product 'tecno-camon-40-pro-5g'
func (s *SpecificationSeederMobileTecnoCamon40Pro5g) Seed(db *gorm.DB) error {
	productSlug := "tecno-camon-40-pro-5g"

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
