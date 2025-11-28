package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileOppoFindN4Fold seeds specifications/options for product 'oppo-find-n4-fold'
type SpecificationSeederMobileOppoFindN4Fold struct {
	BaseSeeder
}

// NewSpecificationSeederMobileOppoFindN4Fold creates a new seeder instance
func NewSpecificationSeederMobileOppoFindN4Fold() *SpecificationSeederMobileOppoFindN4Fold {
	return &SpecificationSeederMobileOppoFindN4Fold{BaseSeeder: BaseSeeder{name: "Specifications for Oppo Find N4 Fold"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileOppoFindN4Fold) getBanglaTranslations() map[string]string {
	return map[string]string{
		"100 W wired": "১০০ W তারযুক্ত",
		"1080p @ 30fps": "১০৮০p @ ৩০fps",
		"12 / 16 GB": "১২ / ১৬ GB",
		"120 Hz": "১২০ Hz",
		"2024": "২০২৪",
		"256 / 512 GB": "২৫৬ / ৫১২ GB",
		"265 g": "২৬৫ g",
		"2716 × 1212 px (main)": "২৭১৬ × ১২১২ px (main)",
		"2× optical": "২× optical",
		"32 MP": "৩২ MP",
		"4800 mAh": "৪৮০০ এমএএইচ",
		"4K @ 30/60fps; 1080p @ 30/60/120fps": "৪K @ ৩০/৬০fps; ১০৮০p @ ৩০/৬০/১২০fps",
		"5.3": "৫.৩",
		"50 MP + 13 MP + 12 MP": "৫০ MP + ১৩ MP + ১২ MP",
		"50 W wireless": "৫০ W বেতার",
		"7.1 inches (main) / 6.3 inches (cover)": "৭.১ ইঞ্চি (main) / ৬.৩ ইঞ্চি (cover)",
		"AMOLED, 120 Hz": "AMOLED, ১২০ Hz",
		"Adreno 740": "Adreno ৭৪০",
		"Android 14, ColorOS 14": "Android ১৪, ColorOS ১৪",
		"Black, Blue": "কালো, নীল",
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

// Seed inserts specification_translations for existing specifications for product 'oppo-find-n4-fold'
func (s *SpecificationSeederMobileOppoFindN4Fold) Seed(db *gorm.DB) error {
	productSlug := "oppo-find-n4-fold"

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
