package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileSamsungGalaxyA345g seeds specifications/options for product 'samsung-galaxy-a34-5g'
type SpecificationSeederMobileSamsungGalaxyA345g struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSamsungGalaxyA345g creates a new seeder instance
func NewSpecificationSeederMobileSamsungGalaxyA345g() *SpecificationSeederMobileSamsungGalaxyA345g {
	return &SpecificationSeederMobileSamsungGalaxyA345g{BaseSeeder: BaseSeeder{name: "Specifications for Samsung Galaxy A34 5G"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSamsungGalaxyA345g) getBanglaTranslations() map[string]string {
	return map[string]string{
		"120 Hz": "১২০ Hz",
		"128 / 256 GB + microSD (expandable)": "১২৮ / ২৫৬ GB + microSD (expandable)",
		"13 MP": "১৩ MP",
		"161.3 × 78.1 × 8.2 mm": "১৬১.৩ × ৭৮.১ × ৮.২ মিমি",
		"199 g": "১৯৯ g",
		"2340 × 1080 pixels (~390 ppi)": "২৩৪০ × ১০৮০ pixels (~৩৯০ ppi)",
		"25 W wired": "২৫ W তারযুক্ত",
		"4 / 6 / 8 GB": "৪ / ৬ / ৮ GB",
		"4 OS updates, 5 years security updates": "৪ OS updates, ৫ years security updates",
		"48 MP + 8 MP + 5 MP": "৪৮ MP + ৮ MP + ৫ MP",
		"4K @ 30fps; 1080p @ 30fps": "৪K @ ৩০fps; ১০৮০p @ ৩০fps",
		"4K @ 30fps; 1080p @30/60fps": "৪K @ ৩০fps; ১০৮০p @৩০/৬০fps",
		"5.3": "৫.৩",
		"5000 mAh (typical)": "৫০০০ এমএএইচ (সাধারণ)",
		"6.6 inches": "৬.৬ ইঞ্চি",
		"Android 13, One UI 5.1": "Android ১৩, One UI ৫.১",
		"Awesome Lime, Awesome Graphite, Awesome Silver, Awesome Violet": "Awesome Lime, Awesome Graphite, Awesome রূপালী, Awesome Violet",
		"Dimensity 1080": "Dimensity ১০৮০",
		"Fingerprint (side), Accelerometer, Gyro, Proximity, Compass": "ফিঙ্গারপ্রিন্ট (side), অ্যাক্সিলেরোমিটার, জাইরো, প্রক্সিমিটি, কম্পাস",
		"GSM / HSPA / LTE / 5G": "GSM / HSPA / LTE / ৫G",
		"Glass front (Gorilla Glass 5), plastic frame & back": "গ্লাস সামনে (Gorilla Glass ৫), plastic frame & back",
		"IP67 (dust / water up to 1m for 30 min)": "IP৬৭ (dust / water up to ১m for ৩০ min)",
		"Li-Ion (non-removable)": "লি-আয়ন (অপসারণযোগ্য নয়)",
		"Mali-G68 MC4": "Mali-G৬৮ MC৪",
		"March 2023": "March ২০২৩",
		"MediaTek Dimensity 1080": "MediaTek Dimensity ১০৮০",
		"Octa-core (2×2.6 GHz Cortex-A78 + 6×2.0 GHz Cortex-A55)": "Octa-core (২×২.৬ GHz Cortex-A৭৮ + ৬×২.০ GHz Cortex-A৫৫)",
		"Super AMOLED, 120 Hz": "Super AMOLED, ১২০ Hz",
		"Wi-Fi 802.11 a/b/g/n/ac": "Wi-Fi ৮০২.১১ a/b/g/n/ac",
		"Yes (market dependent)": "হ্যাঁ (market dependent)",
	}
}

// Seed inserts specification_translations for existing specifications for product 'samsung-galaxy-a34-5g'
func (s *SpecificationSeederMobileSamsungGalaxyA345g) Seed(db *gorm.DB) error {
	productSlug := "samsung-galaxy-a34-5g"

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
