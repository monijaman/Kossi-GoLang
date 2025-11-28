package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileOppoFindN4Flip seeds specifications/options for product 'oppo-find-n4-flip'
type SpecificationSeederMobileOppoFindN4Flip struct {
	BaseSeeder
}

// NewSpecificationSeederMobileOppoFindN4Flip creates a new seeder instance
func NewSpecificationSeederMobileOppoFindN4Flip() *SpecificationSeederMobileOppoFindN4Flip {
	return &SpecificationSeederMobileOppoFindN4Flip{BaseSeeder: BaseSeeder{name: "Specifications for Oppo Find N4 Flip"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileOppoFindN4Flip) getBanglaTranslations() map[string]string {
	return map[string]string{
		"100 W wired": "১০০ W তারযুক্ত",
		"1080p @ 30fps": "১০৮০p @ ৩০fps",
		"12 / 16 GB": "১২ / ১৬ GB",
		"120 Hz": "১২০ Hz",
		"2024": "২০২৪",
		"210 g": "২১০ g",
		"2412 × 1080 px": "২৪১২ × ১০৮০ px",
		"256 / 512 GB": "২৫৬ / ৫১২ GB",
		"2× optical": "২× optical",
		"32 MP": "৩২ MP",
		"4500 mAh": "৪৫০০ এমএএইচ",
		"4K @ 30/60fps; 1080p @ 30/60/120fps": "৪K @ ৩০/৬০fps; ১০৮০p @ ৩০/৬০/১২০fps",
		"5.3": "৫.৩",
		"50 MP + 13 MP + 12 MP": "৫০ MP + ১৩ MP + ১২ MP",
		"50 W wireless": "৫০ W বেতার",
		"6.8 inches": "৬.৮ ইঞ্চি",
		"AMOLED, 120 Hz": "AMOLED, ১২০ Hz",
		"Adreno 740": "Adreno ৭৪০",
		"Android 14, ColorOS 14": "Android ১৪, ColorOS ১৪",
		"Black, Silver": "কালো, রূপালী",
		"Dual SIM (Nano-SIM)": "ডুয়াল সিম (ন্যানো-সিম)",
		"Fingerprint (under display), Accelerometer, Gyro, Proximity, Compass": "ফিঙ্গারপ্রিন্ট (ডিসপ্লের নিচে), অ্যাক্সিলেরোমিটার, জাইরো, প্রক্সিমিটি, কম্পাস",
		"Foldable design, HDR10+": "Foldable design, HDR১০+",
		"GSM / HSPA / LTE / 5G": "GSM / HSPA / LTE / ৫G",
		"Glass front/back, aluminum frame": "গ্লাস সামনে/back, অ্যালুমিনিয়াম ফ্রেম",
		"Li-Ion (non-removable)": "লি-আয়ন (অপসারণযোগ্য নয়)",
		"Snapdragon 8+ Gen 2": "Snapdragon ৮+ Gen ২",
		"Snapdragon 8+ Gen 2 (4 nm)": "Snapdragon ৮+ Gen ২ (৪ nm)",
		"USB-C 3.1": "USB-C ৩.১",
		"Wi-Fi 802.11 a/b/g/n/ac/6/6E": "Wi-Fi ৮০২.১১ a/b/g/n/ac/৬/৬E",
	}
}

// Seed inserts specification_translations for existing specifications for product 'oppo-find-n4-flip'
func (s *SpecificationSeederMobileOppoFindN4Flip) Seed(db *gorm.DB) error {
	productSlug := "oppo-find-n4-flip"

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
