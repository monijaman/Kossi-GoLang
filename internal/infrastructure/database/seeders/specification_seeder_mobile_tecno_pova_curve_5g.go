package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileTecnoPovaCurve5g seeds specifications/options for product 'tecno-pova-curve-5g'
type SpecificationSeederMobileTecnoPovaCurve5g struct {
	BaseSeeder
}

// NewSpecificationSeederMobileTecnoPovaCurve5g creates a new seeder instance
func NewSpecificationSeederMobileTecnoPovaCurve5g() *SpecificationSeederMobileTecnoPovaCurve5g {
	return &SpecificationSeederMobileTecnoPovaCurve5g{BaseSeeder: BaseSeeder{name: "Specifications for Tecno POVA Curve 5G"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileTecnoPovaCurve5g) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080 x 2400 pixels": "১০৮০ x ২৪০০ pixels",
		"120Hz": "১২০Hz",
		"16 MP": "১৬ MP",
		"162 x 75 x 7.9 mm": "১৬২ x ৭৫ x ৭.৯ মিমি",
		"185 g": "১৮৫ g",
		"2024": "২০২৪",
		"256 GB": "২৫৬ GB",
		"5,000 mAh": "৫,০০০ এমএএইচ",
		"50 MP + 2 MP": "৫০ MP + ২ MP",
		"5G": "৫G",
		"6.67 inches": "৬.৬৭ ইঞ্চি",
		"8 GB": "৮ GB",
		"AMOLED, 120Hz, Curved": "AMOLED, ১২০Hz, Curved",
		"Android 14": "Android ১৪",
		"Black, Blue": "কালো, নীল",
		"Dimensity 7050": "Dimensity ৭০৫০",
		"Mali-G68 MC4": "Mali-G৬৮ MC৪",
		"Mediatek Dimensity 7050 (6 nm)": "Mediatek Dimensity ৭০৫০ (৬ nm)",
	}
}

// Seed inserts specification_translations for existing specifications for product 'tecno-pova-curve-5g'
func (s *SpecificationSeederMobileTecnoPovaCurve5g) Seed(db *gorm.DB) error {
	productSlug := "tecno-pova-curve-5g"

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
