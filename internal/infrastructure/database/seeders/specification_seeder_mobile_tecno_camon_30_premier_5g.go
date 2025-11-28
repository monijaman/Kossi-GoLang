package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileTecnoCamon30Premier5g seeds specifications/options for product 'tecno-camon-30-premier-5g'
type SpecificationSeederMobileTecnoCamon30Premier5g struct {
	BaseSeeder
}

// NewSpecificationSeederMobileTecnoCamon30Premier5g creates a new seeder instance
func NewSpecificationSeederMobileTecnoCamon30Premier5g() *SpecificationSeederMobileTecnoCamon30Premier5g {
	return &SpecificationSeederMobileTecnoCamon30Premier5g{BaseSeeder: BaseSeeder{name: "Specifications for Tecno CAMON 30 Premier 5G"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileTecnoCamon30Premier5g) getBanglaTranslations() map[string]string {
	return map[string]string{
		"12 GB": "১২ GB",
		"120Hz": "১২০Hz",
		"1264 x 2780 pixels": "১২৬৪ x ২৭৮০ pixels",
		"162.7 x 76.2 x 7.9 mm": "১৬২.৭ x ৭৬.২ x ৭.৯ মিমি",
		"210 g": "২১০ g",
		"5,000 mAh": "৫,০০০ এমএএইচ",
		"50 MP": "৫০ MP",
		"50 MP + 50 MP + 50 MP": "৫০ MP + ৫০ MP + ৫০ MP",
		"512 GB": "৫১২ GB",
		"5G": "৫G",
		"6.77 inches": "৬.৭৭ ইঞ্চি",
		"Alps Snowy Silver, Hawaii Lava Black": "Alps Snowy রূপালী, Hawaii Lava কালো",
		"Android 14, HIOS 14": "Android ১৪, HIOS ১৪",
		"Corning Gorilla Glass 5": "Corning Gorilla Glass ৫",
		"Dimensity 8200 Ultimate": "Dimensity ৮২০০ Ultimate",
		"February 2024": "February ২০২৪",
		"Glass front, glass back, aluminum frame": "গ্লাস সামনে, গ্লাস পেছনে, অ্যালুমিনিয়াম ফ্রেম",
		"IP54": "IP৫৪",
		"LTPO AMOLED, 120Hz": "LTPO AMOLED, ১২০Hz",
		"Mali-G610 MC6": "Mali-G৬১০ MC৬",
		"Mediatek Dimensity 8200 Ultimate (4 nm)": "Mediatek Dimensity ৮২০০ Ultimate (৪ nm)",
	}
}

// Seed inserts specification_translations for existing specifications for product 'tecno-camon-30-premier-5g'
func (s *SpecificationSeederMobileTecnoCamon30Premier5g) Seed(db *gorm.DB) error {
	productSlug := "tecno-camon-30-premier-5g"

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
