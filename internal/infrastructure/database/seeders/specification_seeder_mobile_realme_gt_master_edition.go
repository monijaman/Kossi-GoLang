package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileRealmeGtMasterEdition seeds specifications/options for product 'realme-gt-master-edition'
type SpecificationSeederMobileRealmeGtMasterEdition struct {
	BaseSeeder
}

// NewSpecificationSeederMobileRealmeGtMasterEdition creates a new seeder instance
func NewSpecificationSeederMobileRealmeGtMasterEdition() *SpecificationSeederMobileRealmeGtMasterEdition {
	return &SpecificationSeederMobileRealmeGtMasterEdition{BaseSeeder: BaseSeeder{name: "Specifications for Realme GT Master Edition"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileRealmeGtMasterEdition) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080p @ 30fps": "১০৮০p @ ৩০fps",
		"120 Hz": "১২০ Hz",
		"128 / 256 GB": "১২৮ / ২৫৬ GB",
		"174 g": "১৭৪ g",
		"2021": "২০২১",
		"2400 × 1080 px (~409 ppi)": "২৪০০ × ১০৮০ px (~৪০৯ ppi)",
		"32 MP": "৩২ MP",
		"4300 mAh": "৪৩০০ এমএএইচ",
		"4K @ 30fps, 1080p @ 30/60/120fps": "৪K @ ৩০fps, ১০৮০p @ ৩০/৬০/১২০fps",
		"5.2": "৫.২",
		"6 / 8 GB": "৬ / ৮ GB",
		"6.43 inches": "৬.৪৩ ইঞ্চি",
		"64 MP + 8 MP + 2 MP": "৬৪ MP + ৮ MP + ২ MP",
		"65 W wired": "৬৫ W তারযুক্ত",
		"Adreno 642L": "Adreno ৬৪২L",
		"Android 11, Realme UI 2.0": "Android ১১, Realme UI ২.০",
		"Dual SIM (Nano-SIM)": "ডুয়াল সিম (ন্যানো-সিম)",
		"Fingerprint (under display), Accelerometer, Gyro, Proximity, Compass": "ফিঙ্গারপ্রিন্ট (ডিসপ্লের নিচে), অ্যাক্সিলেরোমিটার, জাইরো, প্রক্সিমিটি, কম্পাস",
		"GSM / HSPA / LTE / 5G": "GSM / HSPA / LTE / ৫G",
		"Glass front, plastic frame/back": "গ্লাস সামনে, plastic frame/back",
		"Li-Ion (non-removable)": "লি-আয়ন (অপসারণযোগ্য নয়)",
		"Octa-core (4×2.4 GHz Cortex-A78 + 4×1.8 GHz Cortex-A55)": "Octa-core (৪×২.৪ GHz Cortex-A৭৮ + ৪×১.৮ GHz Cortex-A৫৫)",
		"Qualcomm Snapdragon 778G": "Qualcomm Snapdragon ৭৭৮G",
		"Snapdragon 778G (6 nm)": "Snapdragon ৭৭৮G (৬ nm)",
		"সুপার AMOLED, 120 Hz": "সুপার AMOLED, ১২০ Hz",
		"USB-C 2.0": "USB-C ২.০",
		"Voyager Grey, Daybreak Blue": "Voyager ধূসর, Daybreak নীল",
		"Wi-Fi 802.11 a/b/g/n/ac/6": "Wi-Fi ৮০২.১১ a/b/g/n/ac/৬",
	}
}

// Seed inserts specification_translations for existing specifications for product 'realme-gt-master-edition'
func (s *SpecificationSeederMobileRealmeGtMasterEdition) Seed(db *gorm.DB) error {
	productSlug := "realme-gt-master-edition"

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
