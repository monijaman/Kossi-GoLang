package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileOppoA78 seeds specifications/options for product 'oppo-a78'
type SpecificationSeederMobileOppoA78 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileOppoA78 creates a new seeder instance
func NewSpecificationSeederMobileOppoA78() *SpecificationSeederMobileOppoA78 {
	return &SpecificationSeederMobileOppoA78{BaseSeeder: BaseSeeder{name: "Specifications for Oppo A78"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileOppoA78) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080p @ 30fps": "১০৮০p @ ৩০fps",
		"120 Hz": "১২০ Hz",
		"128 / 256 GB + microSD": "১২৮ / ২৫৬ GB + microSD",
		"16 MP": "১৬ MP",
		"186 g": "১৮৬ g",
		"2023": "২০২৩",
		"2412 × 1080 px (~401 ppi)": "২৪১২ × ১০৮০ px (~৪০১ ppi)",
		"3.5 mm headphone jack": "৩.৫ মিমি headphone jack",
		"33 W wired": "৩৩ W তারযুক্ত",
		"4K @ 30fps": "৪K @ ৩০fps",
		"5.2": "৫.২",
		"5000 mAh": "৫০০০ এমএএইচ",
		"6 / 8 GB": "৬ / ৮ GB",
		"6.56 inches": "৬.৫৬ ইঞ্চি",
		"64 MP + 2 MP": "৬৪ MP + ২ MP",
		"Android 13, ColorOS 13": "Android ১৩, ColorOS ১৩",
		"Black, Blue, Green": "কালো, নীল, সবুজ",
		"Dimensity 930 (6 nm)": "Dimensity ৯৩০ (৬ nm)",
		"Dual SIM (Nano-SIM)": "ডুয়াল সিম (ন্যানো-সিম)",
		"Fingerprint (side), Accelerometer, Gyro, Proximity, Compass": "ফিঙ্গারপ্রিন্ট (side), অ্যাক্সিলেরোমিটার, জাইরো, প্রক্সিমিটি, কম্পাস",
		"GSM / HSPA / LTE / 5G": "GSM / HSPA / LTE / ৫G",
		"Glass front, plastic frame/back": "গ্লাস সামনে, plastic frame/back",
		"Li-Ion (non-removable)": "লি-আয়ন (অপসারণযোগ্য নয়)",
		"Mali-G68 MC4": "Mali-G৬৮ MC৪",
		"MediaTek Dimensity 930": "MediaTek Dimensity ৯৩০",
		"Octa-core (2×2.2 GHz Cortex-A78 + 6×2.0 GHz Cortex-A55)": "Octa-core (২×২.২ GHz Cortex-A৭৮ + ৬×২.০ GHz Cortex-A৫৫)",
		"USB-C 2.0": "USB-C ২.০",
		"Wi-Fi 802.11 a/b/g/n/ac": "Wi-Fi ৮০২.১১ a/b/g/n/ac",
	}
}

// Seed inserts specification_translations for existing specifications for product 'oppo-a78'
func (s *SpecificationSeederMobileOppoA78) Seed(db *gorm.DB) error {
	productSlug := "oppo-a78"

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
