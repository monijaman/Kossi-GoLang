package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileXiaomi13t seeds specifications/options for product 'xiaomi-13t'
type SpecificationSeederMobileXiaomi13t struct {
	BaseSeeder
}

// NewSpecificationSeederMobileXiaomi13t creates a new seeder instance
func NewSpecificationSeederMobileXiaomi13t() *SpecificationSeederMobileXiaomi13t {
	return &SpecificationSeederMobileXiaomi13t{BaseSeeder: BaseSeeder{name: "Specifications for Xiaomi 13T"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileXiaomi13t) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1220 x 2712 pixels": "১২২০ x ২৭১২ pixels",
		"144Hz": "১৪৪Hz",
		"193 g": "১৯৩ g",
		"20 MP": "২০ MP",
		"256 GB": "২৫৬ GB",
		"50 MP + 50 MP + 12 MP": "৫০ MP + ৫০ MP + ১২ MP",
		"5000 mAh": "৫০০০ এমএএইচ",
		"6.67 inches": "৬.৬৭ ইঞ্চি",
		"67W wired": "৬৭W তারযুক্ত",
		"8 GB / 12 GB": "৮ GB / ১২ GB",
		"AMOLED, 144Hz, Dolby Vision, HDR10+": "AMOLED, ১৪৪Hz, Dolby Vision, HDR১০+",
		"Alpine Blue, Meadow Green, Black": "Alpine নীল, Meadow সবুজ, কালো",
		"Android 13, upgradable to Android 14, HyperOS": "Android ১৩, upgradable to Android ১৪, HyperOS",
		"Dimensity 8200 Ultra": "Dimensity ৮২০০ Ultra",
		"GSM / HSPA / LTE / 5G": "GSM / HSPA / LTE / ৫G",
		"Glass front (Gorilla Glass 5), glass back or silicone polymer back, plastic frame": "গ্লাস সামনে (Gorilla Glass ৫), গ্লাস পেছনে or silicone polymer back, plastic frame",
		"IP68": "IP৬৮",
		"Mali-G610 MC6": "Mali-G৬১০ MC৬",
		"Mediatek Dimensity 8200 Ultra (4 nm)": "Mediatek Dimensity ৮২০০ Ultra (৪ nm)",
		"September 2023": "September ২০২৩",
	}
}

// Seed inserts specification_translations for existing specifications for product 'xiaomi-13t'
func (s *SpecificationSeederMobileXiaomi13t) Seed(db *gorm.DB) error {
	productSlug := "xiaomi-13t"

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
