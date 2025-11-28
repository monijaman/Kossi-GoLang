package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileSamsungGalaxyA23 seeds specifications/options for product 'samsung-galaxy-a23'
type SpecificationSeederMobileSamsungGalaxyA23 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSamsungGalaxyA23 creates a new seeder instance
func NewSpecificationSeederMobileSamsungGalaxyA23() *SpecificationSeederMobileSamsungGalaxyA23 {
	return &SpecificationSeederMobileSamsungGalaxyA23{BaseSeeder: BaseSeeder{name: "Specifications for Samsung Galaxy A23"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSamsungGalaxyA23) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080p @ 30fps": "১০৮০p @ ৩০fps",
		"165.4 × 76.9 × 8.4 mm": "১৬৫.৪ × ৭৬.৯ × ৮.৪ মিমি",
		"195 g": "১৯৫ g",
		"2408 × 1080 px (~400 ppi)": "২৪০৮ × ১০৮০ px (~৪০০ ppi)",
		"25 W wired": "২৫ W তারযুক্ত",
		"3.5 mm headphone jack": "৩.৫ মিমি headphone jack",
		"4 / 6 / 8 GB": "৪ / ৬ / ৮ GB",
		"5.0": "৫.০",
		"50 MP + 5 MP + 2 MP + 2 MP": "৫০ MP + ৫ MP + ২ MP + ২ MP",
		"5000 mAh": "৫০০০ এমএএইচ",
		"6.6 inches": "৬.৬ ইঞ্চি",
		"64 / 128 GB + microSD": "৬৪ / ১২৮ GB + microSD",
		"8 MP": "৮ MP",
		"90 Hz": "৯০ Hz",
		"Adreno 610": "Adreno ৬১০",
		"Android 12, One UI 4.1 (আপগ্রেডযোগ্য)": "Android ১২, One UI ৪.১ (আপগ্রেডযোগ্য)",
		"Black, White, Peach, Blue": "কালো, সাদা, Peach, নীল",
		"GSM / HSPA / LTE (4G)": "GSM / HSPA / LTE (৪G)",
		"Glass front (Gorilla Glass 5), plastic frame/back": "গ্লাস সামনে (Gorilla Glass ৫), plastic frame/back",
		"Li-Ion (non-removable)": "লি-আয়ন (অপসারণযোগ্য নয়)",
		"March 2022": "March ২০২২",
		"Octa-core (4×2.4 GHz + 4×1.9 GHz)": "Octa-core (৪×২.৪ GHz + ৪×১.৯ GHz)",
		"PLS LCD, 90 Hz": "PLS LCD, ৯০ Hz",
		"Side fingerprint, Accelerometer, Gyro, Proximity, Compass": "Side ফিঙ্গারপ্রিন্ট, অ্যাক্সিলেরোমিটার, জাইরো, প্রক্সিমিটি, কম্পাস",
		"Snapdragon 680 (4G)": "Snapdragon ৬৮০ (৪G)",
		"Snapdragon 680 (6 nm)": "Snapdragon ৬৮০ (৬ nm)",
		"Some variants (A23 5G), base A23 is 4G": "Some variants (A২৩ ৫G), base A২৩ is ৪G",
		"USB-C 2.0": "USB-C ২.০",
		"Wi-Fi 802.11 a/b/g/n/ac": "Wi-Fi ৮০২.১১ a/b/g/n/ac",
		"Yes (বাজার নির্ভরশীল)": "হ্যাঁ (বাজার নির্ভরশীল)",
	}
}

// Seed inserts specification_translations for existing specifications for product 'samsung-galaxy-a23'
func (s *SpecificationSeederMobileSamsungGalaxyA23) Seed(db *gorm.DB) error {
	productSlug := "samsung-galaxy-a23"

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
