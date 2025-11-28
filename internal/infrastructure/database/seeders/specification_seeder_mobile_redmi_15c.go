package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileRedmi15c seeds specifications/options for product 'redmi-15c'
type SpecificationSeederMobileRedmi15c struct {
	BaseSeeder
}

// NewSpecificationSeederMobileRedmi15c creates a new seeder instance
func NewSpecificationSeederMobileRedmi15c() *SpecificationSeederMobileRedmi15c {
	return &SpecificationSeederMobileRedmi15c{BaseSeeder: BaseSeeder{name: "Specifications for Redmi 15C"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileRedmi15c) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080p @ 30fps": "১০৮০p @ ৩০fps",
		"120 Hz": "১২০ Hz",
		"128 / 256 GB + microSD": "১২৮ / ২৫৬ GB + microSD",
		"1600 × 720 px": "১৬০০ × ৭২০ px",
		"171.6 × 79.5 × ~8 mm": "১৭১.৬ × ৭৯.৫ × ~৮ মিমি",
		"2025": "২০২৫",
		"205 g": "২০৫ g",
		"33 W wired": "৩৩ W তারযুক্ত",
		"4 / 6 / 8 GB": "৪ / ৬ / ৮ GB",
		"5.4": "৫.৪",
		"50 MP + secondary lens": "৫০ MP + secondary lens",
		"6.9 inches": "৬.৯ ইঞ্চি",
		"6000 mAh (typ)": "৬০০০ এমএএইচ (typ)",
		"8 MP": "৮ MP",
		"802.11 a/b/g/n": "৮০২.১১ a/b/g/n",
		"GSM / HSPA / LTE (4G)": "GSM / HSPA / LTE (৪G)",
		"Glass front, plastic back": "গ্লাস সামনে, প্লাস্টিক পেছনে",
		"Helio G81 Ultra": "Helio G৮১ Ultra",
		"Li-Ion (non-removable)": "লি-আয়ন (অপসারণযোগ্য নয়)",
		"Mali‑G52 MC2": "Mali‑G৫২ MC২",
		"MediaTek Helio G81 Ultra": "MediaTek Helio G৮১ Ultra",
		"Moonlight Blue, Mint Green, Midnight Gray, Twilight Orange": "Moonlight নীল, Mint সবুজ, Midnight ধূসর, Twilight Orange",
		"Nano-SIM + Nano-SIM": "ন্যানো-সিম + ন্যানো-সিম",
		"Side fingerprint, Accelerometer": "Side ফিঙ্গারপ্রিন্ট, অ্যাক্সিলেরোমিটার",
		"Yes, 3.5 mm": "হ্যাঁ, ৩.৫ মিমি",
	}
}

// Seed inserts specification_translations for existing specifications for product 'redmi-15c'
func (s *SpecificationSeederMobileRedmi15c) Seed(db *gorm.DB) error {
	productSlug := "redmi-15c"

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
