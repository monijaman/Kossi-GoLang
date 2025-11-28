package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileRedmi12c seeds specifications/options for product 'redmi-12c'
type SpecificationSeederMobileRedmi12c struct {
	BaseSeeder
}

// NewSpecificationSeederMobileRedmi12c creates a new seeder instance
func NewSpecificationSeederMobileRedmi12c() *SpecificationSeederMobileRedmi12c {
	return &SpecificationSeederMobileRedmi12c{BaseSeeder: BaseSeeder{name: "Specifications for Redmi 12C"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileRedmi12c) getBanglaTranslations() map[string]string {
	return map[string]string{
		"10 W wired": "১০ W তারযুক্ত",
		"1080p @ 30fps": "১০৮০p @ ৩০fps",
		"13 MP + 2 MP": "১৩ MP + ২ MP",
		"1650 × 720 px": "১৬৫০ × ৭২০ px",
		"168.8 × 76.5 × 8.9 mm": "১৬৮.৮ × ৭৬.৫ × ৮.৯ মিমি",
		"196 g": "১৯৬ g",
		"2022": "২০২২",
		"3 / 4 / 6 GB": "৩ / ৪ / ৬ GB",
		"3.5 mm": "৩.৫ মিমি",
		"5 MP": "৫ MP",
		"5.1": "৫.১",
		"5000 mAh": "৫০০০ এমএএইচ",
		"6.71 inches": "৬.৭১ ইঞ্চি",
		"60 Hz": "৬০ Hz",
		"64 / 128 GB + microSD": "৬৪ / ১২৮ GB + microSD",
		"Accelerometer, Proximity": "অ্যাক্সিলেরোমিটার, প্রক্সিমিটি",
		"Android 13 (Go Edition) / MIUI": "Android ১৩ (Go Edition) / MIUI",
		"Graphite Gray, Coral Green, Sky Blue": "Graphite ধূসর, Coral সবুজ, Sky নীল",
		"Helio G85": "Helio G৮৫",
		"Li‑Ion (non-removable)": "Li‑Ion (অপসারণযোগ্য নয়)",
		"Mali-G52": "Mali-G৫২",
		"MediaTek Helio G85": "MediaTek Helio G৮৫",
		"Nano-SIM + Nano-SIM": "ন্যানো-সিম + ন্যানো-সিম",
		"Plastic back, glass front": "প্লাস্টিক পেছনে, গ্লাস সামনে",
		"Wi‑Fi 802.11 a/b/g/n/ac": "Wi‑Fi ৮০২.১১ a/b/g/n/ac",
	}
}

// Seed inserts specification_translations for existing specifications for product 'redmi-12c'
func (s *SpecificationSeederMobileRedmi12c) Seed(db *gorm.DB) error {
	productSlug := "redmi-12c"

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
