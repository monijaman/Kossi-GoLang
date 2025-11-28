package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileSymphonyHelio40 seeds specifications/options for product 'symphony-helio-40'
type SpecificationSeederMobileSymphonyHelio40 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSymphonyHelio40 creates a new seeder instance
func NewSpecificationSeederMobileSymphonyHelio40() *SpecificationSeederMobileSymphonyHelio40 {
	return &SpecificationSeederMobileSymphonyHelio40{BaseSeeder: BaseSeeder{name: "Specifications for Symphony Helio 40"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSymphonyHelio40) getBanglaTranslations() map[string]string {
	return map[string]string{
		"108 MP + 2 MP": "১০৮ MP + ২ MP",
		"1080p@30fps": "১০৮০p@৩০fps",
		"128 GB + microSD": "১২৮ GB + microSD",
		"16 MP": "১৬ MP",
		"167.4 x 77.6 x 8.54 mm": "১৬৭.৪ x ৭৭.৬ x ৮.৫৪ মিমি",
		"18 W wired": "১৮ W তারযুক্ত",
		"200 g": "২০০ g",
		"2K@30fps, 1080p@30fps": "২K@৩০fps, ১০৮০p@৩০fps",
		"5.0": "৫.০",
		"5000 mAh": "৫০০০ এমএএইচ",
		"6 / 8 GB": "৬ / ৮ GB",
		"6.67 inches": "৬.৬৭ ইঞ্চি",
		"720 × 1604 pixels": "৭২০ × ১৬০৪ pixels",
		"90Hz": "৯০Hz",
		"Android 14": "Android ১৪",
		"Dual Nano-SIM / Hybrid": "Dual ন্যানো-সিম / Hybrid",
		"Forest Green, Cosmic Gold, Sunshine Green, Twilight Gold, Sky Blue": "Forest সবুজ, Cosmic সোনালী, Sunshine সবুজ, Twilight সোনালী, Sky নীল",
		"Glass front, plastic back": "গ্লাস সামনে, প্লাস্টিক পেছনে",
		"June 2024": "June ২০২৪",
		"Li-Po (non-removable)": "লি-পো (অপসারণযোগ্য নয়)",
		"Mali-G52 MC2": "Mali-G৫২ MC২",
		"MediaTek Helio G85": "MediaTek Helio G৮৫",
		"MediaTek Helio G85 (12 nm)": "MediaTek Helio G৮৫ (১২ nm)",
		"Octa-core, up to 2.0 GHz": "Octa-core, up to ২.০ GHz",
		"Side fingerprint, proximity, light, gyroscope": "Side ফিঙ্গারপ্রিন্ট, প্রক্সিমিটি, light, gyroscope",
		"USB Type-C 2.0": "USB Type-C ২.০",
		"Wi-Fi 802.11 b/g/n": "Wi-Fi ৮০২.১১ b/g/n",
		"Yes, 3.5 mm": "হ্যাঁ, ৩.৫ মিমি",
	}
}

// Seed inserts specification_translations for existing specifications for product 'symphony-helio-40'
func (s *SpecificationSeederMobileSymphonyHelio40) Seed(db *gorm.DB) error {
	productSlug := "symphony-helio-40"

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
