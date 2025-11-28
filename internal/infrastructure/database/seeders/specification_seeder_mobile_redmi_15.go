package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileRedmi15 seeds specifications/options for product 'redmi-15'
type SpecificationSeederMobileRedmi15 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileRedmi15 creates a new seeder instance
func NewSpecificationSeederMobileRedmi15() *SpecificationSeederMobileRedmi15 {
	return &SpecificationSeederMobileRedmi15{BaseSeeder: BaseSeeder{name: "Specifications for Redmi 15"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileRedmi15) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080p @ 30fps": "১০৮০p @ ৩০fps",
		"128 / 256 GB + microSD": "১২৮ / ২৫৬ GB + microSD",
		"169.48 × 80.45 × 8.40 mm": "১৬৯.৪৮ × ৮০.৪৫ × ৮.৪০ মিমি",
		"2025": "২০২৫",
		"214 g": "২১৪ g",
		"2340 × 1080 px": "২৩৪০ × ১০৮০ px",
		"33 W wired, 18 W reverse": "৩৩ W তারযুক্ত, ১৮ W reverse",
		"5.0": "৫.০",
		"50 MP + secondary lens": "৫০ MP + secondary lens",
		"6 / 8 GB (up to 16 GB with memory extension)": "৬ / ৮ GB (পর্যন্ত ১৬ GB with memory extension)",
		"6.9 inches": "৬.৯ ইঞ্চি",
		"7000 mAh (typ)": "৭০০০ এমএএইচ (typ)",
		"8 MP": "৮ MP",
		"802.11 a/b/g/n/ac": "৮০২.১১ a/b/g/n/ac",
		"GSM / HSPA / LTE (4G)": "GSM / HSPA / LTE (৪G)",
		"Glass front, plastic frame/back": "গ্লাস সামনে, plastic frame/back",
		"HyperOS 2": "HyperOS ২",
		"IP64": "IP৬৪",
		"Li-Ion (non-removable)": "লি-আয়ন (অপসারণযোগ্য নয়)",
		"Nano-SIM + Nano-SIM": "ন্যানো-সিম + ন্যানো-সিম",
		"Sandy Purple, Midnight Black, Titan Gray": "Sandy বেগুনি, Midnight কালো, Titan ধূসর",
		"Side fingerprint, Accelerometer, etc.": "Side ফিঙ্গারপ্রিন্ট, অ্যাক্সিলেরোমিটার, etc.",
		"Snapdragon 685": "Snapdragon ৬৮৫",
		"Snapdragon 685 (6 nm)": "Snapdragon ৬৮৫ (৬ nm)",
		"Up to 144 Hz": "Up to ১৪৪ Hz",
		"Yes, 3.5 mm": "হ্যাঁ, ৩.৫ মিমি",
	}
}

// Seed inserts specification_translations for existing specifications for product 'redmi-15'
func (s *SpecificationSeederMobileRedmi15) Seed(db *gorm.DB) error {
	productSlug := "redmi-15"

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
