package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileSymphonyInnova30 seeds specifications/options for product 'symphony-innova-30'
type SpecificationSeederMobileSymphonyInnova30 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSymphonyInnova30 creates a new seeder instance
func NewSpecificationSeederMobileSymphonyInnova30() *SpecificationSeederMobileSymphonyInnova30 {
	return &SpecificationSeederMobileSymphonyInnova30{BaseSeeder: BaseSeeder{name: "Specifications for Symphony Innova 30"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSymphonyInnova30) getBanglaTranslations() map[string]string {
	return map[string]string{
		"108 MP + 2 MP + 2 MP": "১০৮ MP + ২ MP + ২ MP",
		"1080p@30fps": "১০৮০p@৩০fps",
		"128 GB + microSD": "১২৮ GB + microSD",
		"164.27 x 76.0 x 8.45 mm": "১৬৪.২৭ x ৭৬.০ x ৮.৪৫ মিমি",
		"18 W wired": "১৮ W তারযুক্ত",
		"193 g": "১৯৩ g",
		"5.0": "৫.০",
		"5000 mAh": "৫০০০ এমএএইচ",
		"6 / 8 GB": "৬ / ৮ GB",
		"6.56 inches": "৬.৫৬ ইঞ্চি",
		"720 × 1612 pixels": "৭২০ × ১৬১২ pixels",
		"8 MP": "৮ MP",
		"90Hz": "৯০Hz",
		"AI, UHD, Night mode, Portrait, Panorama": "AI, UHD, Night mode, Portrait, প্যানোরামা",
		"Android 13": "Android ১৩",
		"Li-Po (non-removable)": "লি-পো (অপসারণযোগ্য নয়)",
		"Mali-G57 @ 750 MHz": "Mali-G৫৭ @ ৭৫০ MHz",
		"March 2024": "March ২০২৪",
		"Mirror White, Reflective Green, Space Green": "Mirror সাদা, Reflective সবুজ, Space সবুজ",
		"Octa-core (2.0 GHz)": "Octa-core (২.০ GHz)",
		"Side fingerprint, face unlock, accelerometer, proximity, light, compass": "Side ফিঙ্গারপ্রিন্ট, face unlock, অ্যাক্সিলেরোমিটার, প্রক্সিমিটি, light, কম্পাস",
		"USB Type-C 2.0": "USB Type-C ২.০",
		"UniSOC T616": "UniSOC T৬১৬",
		"UniSOC T616 (12 nm)": "UniSOC T৬১৬ (১২ nm)",
		"Wi-Fi 802.11 b/g/n": "Wi-Fi ৮০২.১১ b/g/n",
		"Yes, 3.5 mm": "হ্যাঁ, ৩.৫ মিমি",
	}
}

// Seed inserts specification_translations for existing specifications for product 'symphony-innova-30'
func (s *SpecificationSeederMobileSymphonyInnova30) Seed(db *gorm.DB) error {
	productSlug := "symphony-innova-30"

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
