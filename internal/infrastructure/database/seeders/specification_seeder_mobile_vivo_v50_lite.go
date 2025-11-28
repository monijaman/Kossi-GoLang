package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileVivoV50Lite seeds specifications/options for product 'vivo-v50-lite'
type SpecificationSeederMobileVivoV50Lite struct {
	BaseSeeder
}

// NewSpecificationSeederMobileVivoV50Lite creates a new seeder instance
func NewSpecificationSeederMobileVivoV50Lite() *SpecificationSeederMobileVivoV50Lite {
	return &SpecificationSeederMobileVivoV50Lite{BaseSeeder: BaseSeeder{name: "Specifications for Vivo V50 Lite"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileVivoV50Lite) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080p @ 30fps": "১০৮০p @ ৩০fps",
		"128 GB": "১২৮ GB",
		"16 MP": "১৬ MP",
		"18 W wired": "১৮ W তারযুক্ত",
		"180 g": "১৮০ g",
		"2400 × 1080 px": "২৪০০ × ১০৮০ px",
		"4315 mAh": "৪৩১৫ এমএএইচ",
		"48 MP + 2 MP": "৪৮ MP + ২ MP",
		"5.1": "৫.১",
		"6 GB": "৬ GB",
		"6.44 inches": "৬.৪৪ ইঞ্চি",
		"60 Hz": "৬০ Hz",
		"Android 11, Funtouch OS 11": "Android ১১, Funtouch OS ১১",
		"Black, Blue": "কালো, নীল",
		"Dimensity 720 (7 nm)": "Dimensity ৭২০ (৭ nm)",
		"Dual SIM (Nano-SIM)": "ডুয়াল সিম (ন্যানো-সিম)",
		"Fingerprint (side-mounted), Accelerometer, Gyro, Proximity, Compass": "ফিঙ্গারপ্রিন্ট (পাশে মাউন্ট), অ্যাক্সিলেরোমিটার, জাইরো, প্রক্সিমিটি, কম্পাস",
		"Glass front, plastic back/frame": "গ্লাস সামনে, প্লাস্টিক পেছনে/frame",
		"IPS LCD, 60 Hz": "IPS LCD, ৬০ Hz",
		"January 2021": "January ২০২১",
		"Li-Ion (non-removable)": "লি-আয়ন (অপসারণযোগ্য নয়)",
		"Mali-G57 MC3": "Mali-G৫৭ MC৩",
		"MediaTek Dimensity 720": "MediaTek Dimensity ৭২০",
		"USB-C 2.0": "USB-C ২.০",
		"Wi-Fi 802.11 a/b/g/n/ac": "Wi-Fi ৮০২.১১ a/b/g/n/ac",
		"Yes, 3.5 mm": "হ্যাঁ, ৩.৫ মিমি",
	}
}

// Seed inserts specification_translations for existing specifications for product 'vivo-v50-lite'
func (s *SpecificationSeederMobileVivoV50Lite) Seed(db *gorm.DB) error {
	productSlug := "vivo-v50-lite"

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
