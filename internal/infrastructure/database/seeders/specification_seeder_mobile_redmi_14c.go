package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileRedmi14c seeds specifications/options for product 'redmi-14c'
type SpecificationSeederMobileRedmi14c struct {
	BaseSeeder
}

// NewSpecificationSeederMobileRedmi14c creates a new seeder instance
func NewSpecificationSeederMobileRedmi14c() *SpecificationSeederMobileRedmi14c {
	return &SpecificationSeederMobileRedmi14c{BaseSeeder: BaseSeeder{name: "Specifications for Redmi 14C"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileRedmi14c) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080p @ 30fps": "১০৮০p @ ৩০fps",
		"120 Hz": "১২০ Hz",
		"128 / 256 GB + microSD": "১২৮ / ২৫৬ GB + microSD",
		"13 MP": "১৩ MP",
		"1640 × 720 px": "১৬৪০ × ৭২০ px",
		"171.88 × 77.8 × 8.22 mm": "১৭১.৮৮ × ৭৭.৮ × ৮.২২ মিমি",
		"18 W wired": "১৮ W তারযুক্ত",
		"2024": "২০২৪",
		"204–211 g": "২০৪–২১১ g",
		"3.5 mm": "৩.৫ মিমি",
		"4 / 6 GB": "৪ / ৬ GB",
		"5.4": "৫.৪",
		"50 MP + second lens": "৫০ MP + second lens",
		"5160 mAh": "৫১৬০ এমএএইচ",
		"6.88 inches": "৬.৮৮ ইঞ্চি",
		"Android 14 + HyperOS": "Android ১৪ + HyperOS",
		"Helio G81‑Ultra": "Helio G৮১‑Ultra",
		"Li‑Ion (non-removable)": "Li‑Ion (অপসারণযোগ্য নয়)",
		"Mali‑G52 MC2": "Mali‑G৫২ MC২",
		"MediaTek Helio G81‑Ultra": "MediaTek Helio G৮১‑Ultra",
		"Midnight Black, Sage Green, Dreamy Purple, Starry Blue": "Midnight কালো, Sage সবুজ, Dreamy বেগুনি, Starry নীল",
		"Plastic back, glass front": "প্লাস্টিক পেছনে, গ্লাস সামনে",
		"Side fingerprint, Ambient, Accelerometer, Compass": "Side ফিঙ্গারপ্রিন্ট, Ambient, অ্যাক্সিলেরোমিটার, কম্পাস",
		"Wi‑Fi 802.11 a/b/g/n/ac": "Wi‑Fi ৮০২.১১ a/b/g/n/ac",
	}
}

// Seed inserts specification_translations for existing specifications for product 'redmi-14c'
func (s *SpecificationSeederMobileRedmi14c) Seed(db *gorm.DB) error {
	productSlug := "redmi-14c"

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
