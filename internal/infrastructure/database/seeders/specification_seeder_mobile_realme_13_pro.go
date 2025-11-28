package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileRealme13Pro seeds specifications/options for product 'realme-13-pro'
type SpecificationSeederMobileRealme13Pro struct {
	BaseSeeder
}

// NewSpecificationSeederMobileRealme13Pro creates a new seeder instance
func NewSpecificationSeederMobileRealme13Pro() *SpecificationSeederMobileRealme13Pro {
	return &SpecificationSeederMobileRealme13Pro{BaseSeeder: BaseSeeder{name: "Specifications for Realme 13 Pro"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileRealme13Pro) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080 x 2412 pixels": "১০৮০ x ২৪১২ pixels",
		"120Hz": "১২০Hz",
		"128 GB / 256 GB / 512 GB": "১২৮ GB / ২৫৬ GB / ৫১২ GB",
		"183.5 g": "১৮৩.৫ g",
		"32 MP": "৩২ MP",
		"45W wired": "৪৫W তারযুক্ত",
		"50 MP + 8 MP + 2 MP": "৫০ MP + ৮ MP + ২ MP",
		"5200 mAh": "৫২০০ এমএএইচ",
		"6.7 inches": "৬.৭ ইঞ্চি",
		"8 GB / 12 GB": "৮ GB / ১২ GB",
		"AMOLED, 120Hz, 1B colors": "AMOLED, ১২০Hz, ১B colors",
		"Adreno 710": "Adreno ৭১০",
		"Android 14, Realme UI 5.0": "Android ১৪, Realme UI ৫.০",
		"GSM / HSPA / LTE / 5G": "GSM / HSPA / LTE / ৫G",
		"Glass front, glass back, plastic frame": "গ্লাস সামনে, গ্লাস পেছনে, plastic frame",
		"IP65": "IP৬৫",
		"July 2024": "July ২০২৪",
		"Monet Gold, Monet Purple, Emerald Green": "Monet সোনালী, Monet বেগুনি, Emerald সবুজ",
		"Qualcomm Snapdragon 7s Gen 2 (4 nm)": "Qualcomm Snapdragon ৭s Gen ২ (৪ nm)",
		"Snapdragon 7s Gen 2": "Snapdragon ৭s Gen ২",
	}
}

// Seed inserts specification_translations for existing specifications for product 'realme-13-pro'
func (s *SpecificationSeederMobileRealme13Pro) Seed(db *gorm.DB) error {
	productSlug := "realme-13-pro"

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
