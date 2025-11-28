package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileSamsungGalaxyF555g seeds specifications/options for product 'samsung-galaxy-f55-5g'
type SpecificationSeederMobileSamsungGalaxyF555g struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSamsungGalaxyF555g creates a new seeder instance
func NewSpecificationSeederMobileSamsungGalaxyF555g() *SpecificationSeederMobileSamsungGalaxyF555g {
	return &SpecificationSeederMobileSamsungGalaxyF555g{BaseSeeder: BaseSeeder{name: "Specifications for Samsung Galaxy F55 5G"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSamsungGalaxyF555g) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080 × 2408 px": "১০৮০ × ২৪০৮ px",
		"1080p @ 30/60fps": "১০৮০p @ ৩০/৬০fps",
		"1080p @ 30fps": "১০৮০p @ ৩০fps",
		"120 Hz": "১২০ Hz",
		"128 / 256 GB + microSD": "১২৮ / ২৫৬ GB + microSD",
		"13 MP": "১৩ MP",
		"164 × 76 × 8.3 mm": "১৬৪ × ৭৬ × ৮.৩ মিমি",
		"188 g": "১৮৮ g",
		"2024": "২০২৪",
		"25 W wired": "২৫ W তারযুক্ত",
		"5.2": "৫.২",
		"5000 mAh": "৫০০০ এমএএইচ",
		"6 / 8 GB": "৬ / ৮ GB",
		"6.6 inches": "৬.৬ ইঞ্চি",
		"64 MP + 5 MP + 2 MP": "৬৪ MP + ৫ MP + ২ MP",
		"Android 14, One UI 6": "Android ১৪, One UI ৬",
		"Black, Silver, Blue": "কালো, রূপালী, নীল",
		"Dimensity 6080": "Dimensity ৬০৮০",
		"GSM / HSPA / LTE / 5G": "GSM / HSPA / LTE / ৫G",
		"Glass front, plastic back": "গ্লাস সামনে, প্লাস্টিক পেছনে",
		"Li-Ion (non-removable)": "লি-আয়ন (অপসারণযোগ্য নয়)",
		"Mali-G57 MC3": "Mali-G৫৭ MC৩",
		"MediaTek Dimensity 6080": "MediaTek Dimensity ৬০৮০",
		"Nano-SIM + Nano-SIM": "ন্যানো-সিম + ন্যানো-সিম",
		"Side fingerprint, Accelerometer, Gyro, Compass": "Side ফিঙ্গারপ্রিন্ট, অ্যাক্সিলেরোমিটার, জাইরো, কম্পাস",
		"সুপার AMOLED, 120 Hz": "সুপার AMOLED, ১২০ Hz",
		"Wi-Fi 802.11 a/b/g/n/ac": "Wi-Fi ৮০২.১১ a/b/g/n/ac",
	}
}

// Seed inserts specification_translations for existing specifications for product 'samsung-galaxy-f55-5g'
func (s *SpecificationSeederMobileSamsungGalaxyF555g) Seed(db *gorm.DB) error {
	productSlug := "samsung-galaxy-f55-5g"

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
