package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileSamsungGalaxyA065g seeds specifications/options for product 'samsung-galaxy-a06-5g'
type SpecificationSeederMobileSamsungGalaxyA065g struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSamsungGalaxyA065g creates a new seeder instance
func NewSpecificationSeederMobileSamsungGalaxyA065g() *SpecificationSeederMobileSamsungGalaxyA065g {
	return &SpecificationSeederMobileSamsungGalaxyA065g{BaseSeeder: BaseSeeder{name: "Specifications for Samsung Galaxy A06 5G"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSamsungGalaxyA065g) getBanglaTranslations() map[string]string {
	return map[string]string{
		"167.3 x 77.3 x 8.0 mm": "১৬৭.৩ x ৭৭.৩ x ৮.০ মিমি",
		"195 g": "১৯৫ g",
		"25W wired": "২৫W তারযুক্ত",
		"4 GB / 6 GB": "৪ GB / ৬ GB",
		"50 MP + 2 MP": "৫০ MP + ২ MP",
		"5000 mAh": "৫০০০ এমএএইচ",
		"6.7 inches": "৬.৭ ইঞ্চি",
		"64 GB / 128 GB": "৬৪ GB / ১২৮ GB",
		"720 x 1600 pixels": "৭২০ x ১৬০০ pixels",
		"8 MP": "৮ MP",
		"90Hz": "৯০Hz",
		"Android 14": "Android ১৪",
		"Black, Light Blue, Gold": "কালো, Light নীল, সোনালী",
		"GSM / HSPA / LTE / 5G": "GSM / HSPA / LTE / ৫G",
		"Glass front, plastic back, plastic frame": "গ্লাস সামনে, প্লাস্টিক পেছনে, plastic frame",
		"Mali-G57 MC2": "Mali-G৫৭ MC২",
		"Mediatek Dimensity 6300": "Mediatek Dimensity ৬৩০০",
		"Mediatek Dimensity 6300 (6 nm)": "Mediatek Dimensity ৬৩০০ (৬ nm)",
		"October 2024": "October ২০২৪",
		"PLS LCD, 90Hz": "PLS LCD, ৯০Hz",
	}
}

// Seed inserts specification_translations for existing specifications for product 'samsung-galaxy-a06-5g'
func (s *SpecificationSeederMobileSamsungGalaxyA065g) Seed(db *gorm.DB) error {
	productSlug := "samsung-galaxy-a06-5g"

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
