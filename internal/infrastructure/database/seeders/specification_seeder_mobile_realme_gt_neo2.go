package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileRealmeGtNeo2 seeds specifications/options for product 'realme-gt-neo2'
type SpecificationSeederMobileRealmeGtNeo2 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileRealmeGtNeo2 creates a new seeder instance
func NewSpecificationSeederMobileRealmeGtNeo2() *SpecificationSeederMobileRealmeGtNeo2 {
	return &SpecificationSeederMobileRealmeGtNeo2{BaseSeeder: BaseSeeder{name: "Specifications for Realme GT Neo2"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileRealmeGtNeo2) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080p @ 30fps": "১০৮০p @ ৩০fps",
		"120 Hz": "১২০ Hz",
		"128 / 256 GB": "১২৮ / ২৫৬ GB",
		"16 MP": "১৬ MP",
		"200 g": "২০০ g",
		"2021": "২০২১",
		"2400 × 1080 px (~402 ppi)": "২৪০০ × ১০৮০ px (~৪০২ ppi)",
		"4K @ 30/60fps, 1080p @ 30/60/120fps": "৪K @ ৩০/৬০fps, ১০৮০p @ ৩০/৬০/১২০fps",
		"5.1": "৫.১",
		"5000 mAh": "৫০০০ এমএএইচ",
		"6.6 inches": "৬.৬ ইঞ্চি",
		"64 MP + 8 MP + 2 MP": "৬৪ MP + ৮ MP + ২ MP",
		"65 W wired": "৬৫ W তারযুক্ত",
		"8 / 12 GB": "৮ / ১২ GB",
		"AMOLED, 120 Hz": "AMOLED, ১২০ Hz",
		"Adreno 650": "Adreno ৬৫০",
		"Android 11, Realme UI 2.0": "Android ১১, Realme UI ২.০",
		"Dual SIM (Nano-SIM)": "ডুয়াল সিম (ন্যানো-সিম)",
		"Fingerprint (under display), Accelerometer, Gyro, Proximity, Compass": "ফিঙ্গারপ্রিন্ট (ডিসপ্লের নিচে), অ্যাক্সিলেরোমিটার, জাইরো, প্রক্সিমিটি, কম্পাস",
		"GSM / HSPA / LTE / 5G": "GSM / HSPA / LTE / ৫G",
		"Glass front, plastic frame/back": "গ্লাস সামনে, plastic frame/back",
		"Li-Ion (non-removable)": "লি-আয়ন (অপসারণযোগ্য নয়)",
		"Neo Green, Neo Blue, Neo Black": "Neo সবুজ, Neo নীল, Neo কালো",
		"Octa-core (1×3.2 GHz Cortex-A78 + 3×2.42 GHz Cortex-A78 + 4×1.8 GHz Cortex-A55)": "Octa-core (১×৩.২ GHz Cortex-A৭৮ + ৩×২.৪২ GHz Cortex-A৭৮ + ৪×১.৮ GHz Cortex-A৫৫)",
		"Qualcomm Snapdragon 870": "Qualcomm Snapdragon ৮৭০",
		"Snapdragon 870 (7 nm)": "Snapdragon ৮৭০ (৭ nm)",
		"USB-C 2.0": "USB-C ২.০",
		"Wi-Fi 802.11 a/b/g/n/ac/6": "Wi-Fi ৮০২.১১ a/b/g/n/ac/৬",
	}
}

// Seed inserts specification_translations for existing specifications for product 'realme-gt-neo2'
func (s *SpecificationSeederMobileRealmeGtNeo2) Seed(db *gorm.DB) error {
	productSlug := "realme-gt-neo2"

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
