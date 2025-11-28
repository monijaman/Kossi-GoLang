package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileRealme13 seeds specifications/options for product 'realme-13'
type SpecificationSeederMobileRealme13 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileRealme13 creates a new seeder instance
func NewSpecificationSeederMobileRealme13() *SpecificationSeederMobileRealme13 {
	return &SpecificationSeederMobileRealme13{BaseSeeder: BaseSeeder{name: "Specifications for Realme 13"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileRealme13) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080 x 2400 pixels": "১০৮০ x ২৪০০ pixels",
		"120Hz": "১২০Hz",
		"128 GB / 256 GB": "১২৮ GB / ২৫৬ GB",
		"16 MP": "১৬ MP",
		"190 g": "১৯০ g",
		"45W wired": "৪৫W তারযুক্ত",
		"50 MP + 2 MP": "৫০ MP + ২ MP",
		"5000 mAh": "৫০০০ এমএএইচ",
		"6.72 inches": "৬.৭২ ইঞ্চি",
		"8 GB": "৮ GB",
		"Android 14, Realme UI 5.0": "Android ১৪, Realme UI ৫.০",
		"August 2024": "August ২০২৪",
		"Dimensity 6300": "Dimensity ৬৩০০",
		"GSM / HSPA / LTE / 5G": "GSM / HSPA / LTE / ৫G",
		"Glass front, plastic back, plastic frame": "গ্লাস সামনে, প্লাস্টিক পেছনে, plastic frame",
		"IP64": "IP৬৪",
		"IPS LCD, 120Hz": "IPS LCD, ১২০Hz",
		"Mali-G57 MC2": "Mali-G৫৭ MC২",
		"Mediatek Dimensity 6300 (6 nm)": "Mediatek Dimensity ৬৩০০ (৬ nm)",
		"Speed Green, Dark Purple": "Speed সবুজ, Dark বেগুনি",
	}
}

// Seed inserts specification_translations for existing specifications for product 'realme-13'
func (s *SpecificationSeederMobileRealme13) Seed(db *gorm.DB) error {
	productSlug := "realme-13"

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
