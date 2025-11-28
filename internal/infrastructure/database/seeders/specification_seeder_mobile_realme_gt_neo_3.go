package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileRealmeGtNeo3 seeds specifications/options for product 'realme-gt-neo-3'
type SpecificationSeederMobileRealmeGtNeo3 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileRealmeGtNeo3 creates a new seeder instance
func NewSpecificationSeederMobileRealmeGtNeo3() *SpecificationSeederMobileRealmeGtNeo3 {
	return &SpecificationSeederMobileRealmeGtNeo3{BaseSeeder: BaseSeeder{name: "Specifications for Realme GT Neo 3"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileRealmeGtNeo3) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080p @ 30fps": "১০৮০p @ ৩০fps",
		"120 Hz": "১২০ Hz",
		"128 / 256 GB": "১২৮ / ২৫৬ GB",
		"150 W wired": "১৫০ W তারযুক্ত",
		"16 MP": "১৬ MP",
		"188 g": "১৮৮ g",
		"2022": "২০২২",
		"2412 × 1080 px (~394 ppi)": "২৪১২ × ১০৮০ px (~৩৯৪ ppi)",
		"4500 mAh": "৪৫০০ এমএএইচ",
		"4K @ 30/60fps, 1080p @ 30/60/120fps": "৪K @ ৩০/৬০fps, ১০৮০p @ ৩০/৬০/১২০fps",
		"5.2": "৫.২",
		"50 MP + 8 MP + 2 MP": "৫০ MP + ৮ MP + ২ MP",
		"6.7 inches": "৬.৭ ইঞ্চি",
		"8 / 12 GB": "৮ / ১২ GB",
		"AMOLED, 120 Hz": "AMOLED, ১২০ Hz",
		"Android 12, Realme UI 3.0": "Android ১২, Realme UI ৩.০",
		"Dimensity 8100-Max (5 nm)": "Dimensity ৮১০০-Max (৫ nm)",
		"Dual SIM (Nano-SIM)": "ডুয়াল সিম (ন্যানো-সিম)",
		"Fingerprint (under display), Accelerometer, Gyro, Proximity, Compass": "ফিঙ্গারপ্রিন্ট (ডিসপ্লের নিচে), অ্যাক্সিলেরোমিটার, জাইরো, প্রক্সিমিটি, কম্পাস",
		"GSM / HSPA / LTE / 5G": "GSM / HSPA / LTE / ৫G",
		"Glass front, plastic frame/back": "গ্লাস সামনে, plastic frame/back",
		"Li-Ion (non-removable)": "লি-আয়ন (অপসারণযোগ্য নয়)",
		"Mali-G610 MC6": "Mali-G৬১০ MC৬",
		"MediaTek Dimensity 8100-Max": "MediaTek Dimensity ৮১০০-Max",
		"Octa-core (4×2.85 GHz Cortex-A78 + 4×2.0 GHz Cortex-A55)": "Octa-core (৪×২.৮৫ GHz Cortex-A৭৮ + ৪×২.০ GHz Cortex-A৫৫)",
		"Sprint Blue, Nitro Blue, Shade Black": "Sprint নীল, Nitro নীল, Shade কালো",
		"USB-C 2.0": "USB-C ২.০",
		"Wi-Fi 802.11 a/b/g/n/ac/6": "Wi-Fi ৮০২.১১ a/b/g/n/ac/৬",
	}
}

// Seed inserts specification_translations for existing specifications for product 'realme-gt-neo-3'
func (s *SpecificationSeederMobileRealmeGtNeo3) Seed(db *gorm.DB) error {
	productSlug := "realme-gt-neo-3"

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
