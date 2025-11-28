package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileSamsungGalaxyM145g seeds specifications/options for product 'samsung-galaxy-m14-5g'
type SpecificationSeederMobileSamsungGalaxyM145g struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSamsungGalaxyM145g creates a new seeder instance
func NewSpecificationSeederMobileSamsungGalaxyM145g() *SpecificationSeederMobileSamsungGalaxyM145g {
	return &SpecificationSeederMobileSamsungGalaxyM145g{BaseSeeder: BaseSeeder{name: "Specifications for Samsung Galaxy M14 5G"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSamsungGalaxyM145g) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080p @ 30fps": "১০৮০p @ ৩০fps",
		"13 MP": "১৩ MP",
		"15 W / 25 W (region dependent)": "১৫ W / ২৫ W (region dependent)",
		"166.8 × 77.2 × 9.4 mm": "১৬৬.৮ × ৭৭.২ × ৯.৪ মিমি",
		"206 g": "২০৬ g",
		"2408 × 1080 px": "২৪০৮ × ১০৮০ px",
		"3.5 mm": "৩.৫ মিমি",
		"4 / 6 GB": "৪ / ৬ GB",
		"50 MP + 2 MP + 2 MP": "৫০ MP + ২ MP + ২ MP",
		"6.6 inches": "৬.৬ ইঞ্চি",
		"6000 mAh": "৬০০০ এমএএইচ",
		"64 / 128 GB + microSD": "৬৪ / ১২৮ GB + microSD",
		"90 Hz": "৯০ Hz",
		"Android 13, One UI 5.1": "Android ১৩, One UI ৫.১",
		"Exynos 1330": "Exynos ১৩৩০",
		"Exynos 1330 (5 nm)": "Exynos ১৩৩০ (৫ nm)",
		"GSM / HSPA / LTE / 5G": "GSM / HSPA / LTE / ৫G",
		"Glass front, plastic back/frame": "গ্লাস সামনে, প্লাস্টিক পেছনে/frame",
		"Li-Ion (non-removable)": "লি-আয়ন (অপসারণযোগ্য নয়)",
		"Mali-G68 MP2": "Mali-G৬৮ MP২",
		"March 2023": "March ২০২৩",
		"Nano-SIM or Dual SIM (hybrid)": "ন্যানো-সিম or ডুয়াল সিম (hybrid)",
		"Navy Blue, Light Blue, Silver": "Navy নীল, Light নীল, রূপালী",
		"Octa-core (2×2.4 GHz Cortex-A78 + 6×2.0 GHz Cortex-A55)": "Octa-core (২×২.৪ GHz Cortex-A৭৮ + ৬×২.০ GHz Cortex-A৫৫)",
		"Side fingerprint, Accelerometer, Gyro, Proximity, Compass": "Side ফিঙ্গারপ্রিন্ট, অ্যাক্সিলেরোমিটার, জাইরো, প্রক্সিমিটি, কম্পাস",
		"USB-C 2.0": "USB-C ২.০",
		"Wi-Fi 802.11 a/b/g/n/ac": "Wi-Fi ৮০২.১১ a/b/g/n/ac",
	}
}

// Seed inserts specification_translations for existing specifications for product 'samsung-galaxy-m14-5g'
func (s *SpecificationSeederMobileSamsungGalaxyM145g) Seed(db *gorm.DB) error {
	productSlug := "samsung-galaxy-m14-5g"

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
