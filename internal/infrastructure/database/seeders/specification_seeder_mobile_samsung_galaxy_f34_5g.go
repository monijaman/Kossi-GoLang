package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileSamsungGalaxyF345g seeds specifications/options for product 'samsung-galaxy-f34-5g'
type SpecificationSeederMobileSamsungGalaxyF345g struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSamsungGalaxyF345g creates a new seeder instance
func NewSpecificationSeederMobileSamsungGalaxyF345g() *SpecificationSeederMobileSamsungGalaxyF345g {
	return &SpecificationSeederMobileSamsungGalaxyF345g{BaseSeeder: BaseSeeder{name: "Specifications for Samsung Galaxy F34 5G"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSamsungGalaxyF345g) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080 × 2340 px": "১০৮০ × ২৩৪০ px",
		"1080p @ 30fps": "১০৮০p @ ৩০fps",
		"120 Hz": "১২০ Hz",
		"128 / 256 GB + microSD": "১২৮ / ২৫৬ GB + microSD",
		"13 MP": "১৩ MP",
		"161 × 78 × 8.2 mm": "১৬১ × ৭৮ × ৮.২ মিমি",
		"199 g": "১৯৯ g",
		"2024": "২০২৪",
		"25 W wired": "২৫ W তারযুক্ত",
		"48 MP + 8 MP + 5 MP": "৪৮ MP + ৮ MP + ৫ MP",
		"4K @ 30fps; 1080p @ 30/60fps": "৪K @ ৩০fps; ১০৮০p @ ৩০/৬০fps",
		"5.3": "৫.৩",
		"5000 mAh": "৫০০০ এমএএইচ",
		"6 / 8 GB": "৬ / ৮ GB",
		"6.6 inches": "৬.৬ ইঞ্চি",
		"Android 13, One UI 5.1": "Android ১৩, One UI ৫.১",
		"Dimensity 1080": "Dimensity ১০৮০",
		"GSM / HSPA / LTE / 5G": "GSM / HSPA / LTE / ৫G",
		"Glass front, plastic back": "গ্লাস সামনে, প্লাস্টিক পেছনে",
		"IP67": "IP৬৭",
		"Li-Ion (non-removable)": "লি-আয়ন (অপসারণযোগ্য নয়)",
		"Mali-G68 MC4": "Mali-G৬৮ MC৪",
		"MediaTek Dimensity 1080": "MediaTek Dimensity ১০৮০",
		"Nano-SIM + Nano-SIM": "ন্যানো-সিম + ন্যানো-সিম",
		"Side fingerprint, Accelerometer, Gyro, Compass": "Side ফিঙ্গারপ্রিন্ট, অ্যাক্সিলেরোমিটার, জাইরো, কম্পাস",
		"সুপার AMOLED, 120 Hz": "সুপার AMOLED, ১২০ Hz",
		"Wi-Fi 802.11 a/b/g/n/ac": "Wi-Fi ৮০২.১১ a/b/g/n/ac",
	}
}

// Seed inserts specification_translations for existing specifications for product 'samsung-galaxy-f34-5g'
func (s *SpecificationSeederMobileSamsungGalaxyF345g) Seed(db *gorm.DB) error {
	productSlug := "samsung-galaxy-f34-5g"

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
