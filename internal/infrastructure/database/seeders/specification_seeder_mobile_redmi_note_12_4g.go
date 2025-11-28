package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileRedmiNote124g seeds specifications/options for product 'redmi-note-12-4g'
type SpecificationSeederMobileRedmiNote124g struct {
	BaseSeeder
}

// NewSpecificationSeederMobileRedmiNote124g creates a new seeder instance
func NewSpecificationSeederMobileRedmiNote124g() *SpecificationSeederMobileRedmiNote124g {
	return &SpecificationSeederMobileRedmiNote124g{BaseSeeder: BaseSeeder{name: "Specifications for Redmi Note 12 4G"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileRedmiNote124g) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080 x 2400 pixels": "১০৮০ x ২৪০০ pixels",
		"120Hz": "১২০Hz",
		"13 MP": "১৩ MP",
		"183.5 g": "১৮৩.৫ g",
		"33W wired": "৩৩W তারযুক্ত",
		"4 GB / 6 GB / 8 GB": "৪ GB / ৬ GB / ৮ GB",
		"50 MP + 8 MP + 2 MP": "৫০ MP + ৮ MP + ২ MP",
		"5000 mAh": "৫০০০ এমএএইচ",
		"6.67 inches": "৬.৬৭ ইঞ্চি",
		"64 GB / 128 GB": "৬৪ GB / ১২৮ GB",
		"AMOLED, 120Hz": "AMOLED, ১২০Hz",
		"Adreno 610": "Adreno ৬১০",
		"Android 13, MIUI 14": "Android ১৩, MIUI ১৪",
		"Glass front (Gorilla Glass 3), plastic back, plastic frame": "গ্লাস সামনে (Gorilla Glass ৩), প্লাস্টিক পেছনে, plastic frame",
		"IP53": "IP৫৩",
		"March 2023": "March ২০২৩",
		"Onyx Gray, মিন্ট Green, Ice Blue": "Onyx ধূসর, মিন্ট সবুজ, Ice নীল",
		"Qualcomm Snapdragon 685 (6 nm)": "Qualcomm Snapdragon ৬৮৫ (৬ nm)",
		"Snapdragon 685": "Snapdragon ৬৮৫",
	}
}

// Seed inserts specification_translations for existing specifications for product 'redmi-note-12-4g'
func (s *SpecificationSeederMobileRedmiNote124g) Seed(db *gorm.DB) error {
	productSlug := "redmi-note-12-4g"

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
