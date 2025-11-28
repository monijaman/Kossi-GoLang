package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileSamsungGalaxyF545g seeds specifications/options for product 'samsung-galaxy-f54-5g'
type SpecificationSeederMobileSamsungGalaxyF545g struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSamsungGalaxyF545g creates a new seeder instance
func NewSpecificationSeederMobileSamsungGalaxyF545g() *SpecificationSeederMobileSamsungGalaxyF545g {
	return &SpecificationSeederMobileSamsungGalaxyF545g{BaseSeeder: BaseSeeder{name: "Specifications for Samsung Galaxy F54 5G"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSamsungGalaxyF545g) getBanglaTranslations() map[string]string {
	return map[string]string{
		"108 MP + 12 MP + 5 MP": "১০৮ MP + ১২ MP + ৫ MP",
		"1080 × 2400 px": "১০৮০ × ২৪০০ px",
		"1080p @ 30fps": "১০৮০p @ ৩০fps",
		"120 Hz": "১২০ Hz",
		"128 / 256 GB": "১২৮ / ২৫৬ GB",
		"165 × 77 × 8.2 mm": "১৬৫ × ৭৭ × ৮.২ মিমি",
		"202 g": "২০২ g",
		"2025": "২০২৫",
		"25 W wired": "২৫ W তারযুক্ত",
		"2x": "২x",
		"32 MP": "৩২ MP",
		"4K @ 30fps; 1080p @ 30/60fps": "৪K @ ৩০fps; ১০৮০p @ ৩০/৬০fps",
		"5.3": "৫.৩",
		"5000 mAh": "৫০০০ এমএএইচ",
		"6 / 8 GB": "৬ / ৮ GB",
		"6.7 inches": "৬.৭ ইঞ্চি",
		"Android 14, One UI 6": "Android ১৪, One UI ৬",
		"Exynos 1380": "Exynos ১৩৮০",
		"GSM / HSPA / LTE / 5G": "GSM / HSPA / LTE / ৫G",
		"Glass front & back": "গ্লাস সামনে & back",
		"Graphite, Blue, Violet": "Graphite, নীল, Violet",
		"IP67": "IP৬৭",
		"Li-Ion (non-removable)": "লি-আয়ন (অপসারণযোগ্য নয়)",
		"Mali-G68 MP5": "Mali-G৬৮ MP৫",
		"Nano-SIM + Nano-SIM": "ন্যানো-সিম + ন্যানো-সিম",
		"Side fingerprint, Accelerometer, Gyro, Compass": "Side ফিঙ্গারপ্রিন্ট, অ্যাক্সিলেরোমিটার, জাইরো, কম্পাস",
		"সুপার AMOLED+, 120 Hz": "সুপার AMOLED+, ১২০ Hz",
		"Wi-Fi 802.11 a/b/g/n/ac": "Wi-Fi ৮০২.১১ a/b/g/n/ac",
	}
}

// Seed inserts specification_translations for existing specifications for product 'samsung-galaxy-f54-5g'
func (s *SpecificationSeederMobileSamsungGalaxyF545g) Seed(db *gorm.DB) error {
	productSlug := "samsung-galaxy-f54-5g"

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
