package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileOppoA745g seeds specifications/options for product 'oppo-a74-5g'
type SpecificationSeederMobileOppoA745g struct {
	BaseSeeder
}

// NewSpecificationSeederMobileOppoA745g creates a new seeder instance
func NewSpecificationSeederMobileOppoA745g() *SpecificationSeederMobileOppoA745g {
	return &SpecificationSeederMobileOppoA745g{BaseSeeder: BaseSeeder{name: "Specifications for Oppo A74 5G"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileOppoA745g) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080p @ 30fps": "১০৮০p @ ৩০fps",
		"128 GB + microSD": "১২৮ GB + microSD",
		"16 MP": "১৬ MP",
		"18 W wired": "১৮ W তারযুক্ত",
		"190 g": "১৯০ g",
		"2021": "২০২১",
		"2400 × 1080 px (~405 ppi)": "২৪০০ × ১০৮০ px (~৪০৫ ppi)",
		"3.5 mm headphone jack": "৩.৫ মিমি headphone jack",
		"4 / 6 GB": "৪ / ৬ GB",
		"48 MP + 2 MP + 2 MP": "৪৮ MP + ২ MP + ২ MP",
		"5.1": "৫.১",
		"5000 mAh": "৫০০০ এমএএইচ",
		"6.5 inches": "৬.৫ ইঞ্চি",
		"60 Hz": "৬০ Hz",
		"Adreno 619": "Adreno ৬১৯",
		"Android 11, ColorOS 11": "Android ১১, ColorOS ১১",
		"Dual SIM (Nano-SIM)": "ডুয়াল সিম (ন্যানো-সিম)",
		"Fingerprint (side), Accelerometer, Proximity, Compass": "ফিঙ্গারপ্রিন্ট (side), অ্যাক্সিলেরোমিটার, প্রক্সিমিটি, কম্পাস",
		"GSM / HSPA / LTE / 5G": "GSM / HSPA / LTE / ৫G",
		"Li-Ion (non-removable)": "লি-আয়ন (অপসারণযোগ্য নয়)",
		"Octa-core (2×2.0 GHz Cortex-A76 + 6×1.8 GHz Cortex-A55)": "Octa-core (২×২.০ GHz Cortex-A৭৬ + ৬×১.৮ GHz Cortex-A৫৫)",
		"Prism Black, Space Silver": "Prism কালো, Space রূপালী",
		"Snapdragon 480 5G": "Snapdragon ৪৮০ ৫G",
		"Snapdragon 480 5G (8 nm)": "Snapdragon ৪৮০ ৫G (৮ nm)",
		"Wi-Fi 802.11 a/b/g/n/ac": "Wi-Fi ৮০২.১১ a/b/g/n/ac",
	}
}

// Seed inserts specification_translations for existing specifications for product 'oppo-a74-5g'
func (s *SpecificationSeederMobileOppoA745g) Seed(db *gorm.DB) error {
	productSlug := "oppo-a74-5g"

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
