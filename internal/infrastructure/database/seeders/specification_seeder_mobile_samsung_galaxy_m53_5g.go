package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileSamsungGalaxyM535g seeds specifications/options for product 'samsung-galaxy-m53-5g'
type SpecificationSeederMobileSamsungGalaxyM535g struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSamsungGalaxyM535g creates a new seeder instance
func NewSpecificationSeederMobileSamsungGalaxyM535g() *SpecificationSeederMobileSamsungGalaxyM535g {
	return &SpecificationSeederMobileSamsungGalaxyM535g{BaseSeeder: BaseSeeder{name: "Specifications for Samsung Galaxy M53 5G"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSamsungGalaxyM535g) getBanglaTranslations() map[string]string {
	return map[string]string{
		"108 MP + 8 MP + 2 MP + 2 MP": "১০৮ MP + ৮ MP + ২ MP + ২ MP",
		"120 Hz": "১২০ Hz",
		"128 / 256 GB + microSD": "১২৮ / ২৫৬ GB + microSD",
		"164.7 × 77.0 × 7.4 mm": "১৬৪.৭ × ৭৭.০ × ৭.৪ মিমি",
		"176 g": "১৭৬ g",
		"2408 × 1080 px": "২৪০৮ × ১০৮০ px",
		"25 W wired": "২৫ W তারযুক্ত",
		"32 MP": "৩২ MP",
		"4K @ 30fps; 1080p @ 30/60fps": "৪K @ ৩০fps; ১০৮০p @ ৩০/৬০fps",
		"4K @ 30fps; 1080p @ 30fps": "৪K @ ৩০fps; ১০৮০p @ ৩০fps",
		"5.2": "৫.২",
		"5000 mAh": "৫০০০ এমএএইচ",
		"6 / 8 GB": "৬ / ৮ GB",
		"6.7 inches": "৬.৭ ইঞ্চি",
		"Android 12, One UI 4.1 (or নতুন)": "Android ১২, One UI ৪.১ (or নতুন)",
		"April 2022": "April ২০২২",
		"Dimensity 900 (6 nm)": "Dimensity ৯০০ (৬ nm)",
		"Dual SIM (Nano + Nano)": "ডুয়াল সিম (Nano + Nano)",
		"GSM / HSPA / LTE / 5G": "GSM / HSPA / LTE / ৫G",
		"Glass front, plastic frame / back": "গ্লাস সামনে, plastic frame / back",
		"Green, Blue, Emerald Brown": "সবুজ, নীল, Emerald Brown",
		"Li-Ion (non-removable)": "লি-আয়ন (অপসারণযোগ্য নয়)",
		"Mali-G68 MC4": "Mali-G৬৮ MC৪",
		"MediaTek Dimensity 900": "MediaTek Dimensity ৯০০",
		"Octa-core (2×2.4 GHz Cortex-A78 + 6×2.0 GHz Cortex-A55)": "Octa-core (২×২.৪ GHz Cortex-A৭৮ + ৬×২.০ GHz Cortex-A৫৫)",
		"Side fingerprint, Accelerometer, Gyro, Proximity, Compass": "Side ফিঙ্গারপ্রিন্ট, অ্যাক্সিলেরোমিটার, জাইরো, প্রক্সিমিটি, কম্পাস",
		"সুপার AMOLED Plus, 120 Hz": "সুপার AMOLED Plus, ১২০ Hz",
		"USB-C 2.0": "USB-C ২.০",
		"Wi-Fi 802.11 a/b/g/n/ac, dual-band": "Wi-Fi ৮০২.১১ a/b/g/n/ac, dual-band",
	}
}

// Seed inserts specification_translations for existing specifications for product 'samsung-galaxy-m53-5g'
func (s *SpecificationSeederMobileSamsungGalaxyM535g) Seed(db *gorm.DB) error {
	productSlug := "samsung-galaxy-m53-5g"

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
