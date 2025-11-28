package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileIphoneSe3rdGen seeds specifications/options for product 'iphone-se-3rd-gen'
type SpecificationSeederMobileIphoneSe3rdGen struct {
	BaseSeeder
}

// NewSpecificationSeederMobileIphoneSe3rdGen creates a new seeder instance
func NewSpecificationSeederMobileIphoneSe3rdGen() *SpecificationSeederMobileIphoneSe3rdGen {
	return &SpecificationSeederMobileIphoneSe3rdGen{BaseSeeder: BaseSeeder{name: "Specifications for iPhone SE (3rd Gen)"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileIphoneSe3rdGen) getBanglaTranslations() map[string]string {
	return map[string]string{
		"12 MP": "১২ MP",
		"138.4 x 67.3 x 7.3 mm": "১৩৮.৪ x ৬৭.৩ x ৭.৩ মিমি",
		"144 g": "১৪৪ g",
		"2,018 mAh": "২,০১৮ এমএএইচ",
		"4 GB": "৪ GB",
		"4-core Apple GPU": "৪-core Apple GPU",
		"4.7 inches": "৪.৭ ইঞ্চি",
		"5G": "৫G",
		"60Hz": "৬০Hz",
		"64 GB / 128 GB / 256 GB": "৬৪ GB / ১২৮ GB / ২৫৬ GB",
		"7 MP": "৭ MP",
		"750 x 1334 pixels": "৭৫০ x ১৩৩৪ pixels",
		"Aluminum frame, glass front/back": "অ্যালুমিনিয়াম ফ্রেম, গ্লাস সামনে/back",
		"Apple A15 Bionic": "Apple A১৫ Bionic",
		"Apple A15 Bionic (5 nm)": "Apple A১৫ Bionic (৫ nm)",
		"IP67": "IP৬৭",
		"March 2022": "March ২০২২",
		"মিডনাইট, Starlight, Red": "মিডনাইট, Starlight, লাল",
		"iOS 15.4, আপগ্রেডযোগ্য": "iOS ১৫.৪, আপগ্রেডযোগ্য",
	}
}

// Seed inserts specification_translations for existing specifications for product 'iphone-se-3rd-gen'
func (s *SpecificationSeederMobileIphoneSe3rdGen) Seed(db *gorm.DB) error {
	productSlug := "iphone-se-3rd-gen"

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
