package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobilePocoX6Neo5g seeds specifications/options for product 'poco-x6-neo-5g'
type SpecificationSeederMobilePocoX6Neo5g struct {
	BaseSeeder
}

// NewSpecificationSeederMobilePocoX6Neo5g creates a new seeder instance
func NewSpecificationSeederMobilePocoX6Neo5g() *SpecificationSeederMobilePocoX6Neo5g {
	return &SpecificationSeederMobilePocoX6Neo5g{BaseSeeder: BaseSeeder{name: "Specifications for POCO X6 Neo 5G"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobilePocoX6Neo5g) getBanglaTranslations() map[string]string {
	return map[string]string{
		"108 MP + 2 MP": "১০৮ MP + ২ MP",
		"1080 x 2400 pixels": "১০৮০ x ২৪০০ pixels",
		"120Hz": "১২০Hz",
		"128 GB / 256 GB": "১২৮ GB / ২৫৬ GB",
		"16 MP": "১৬ MP",
		"175 g": "১৭৫ g",
		"33W wired": "৩৩W তারযুক্ত",
		"5000 mAh": "৫০০০ এমএএইচ",
		"6.67 inches": "৬.৬৭ ইঞ্চি",
		"8 GB / 12 GB": "৮ GB / ১২ GB",
		"AMOLED, 120Hz": "AMOLED, ১২০Hz",
		"Android 13, MIUI 14": "Android ১৩, MIUI ১৪",
		"Astral Black, Horizon Blue, Martian Orange": "Astral কালো, Horizon নীল, Martian Orange",
		"Dimensity 6080": "Dimensity ৬০৮০",
		"GSM / HSPA / LTE / 5G": "GSM / HSPA / LTE / ৫G",
		"Glass front (Gorilla Glass 5), plastic back, plastic frame": "গ্লাস সামনে (Gorilla Glass ৫), প্লাস্টিক পেছনে, plastic frame",
		"IP54": "IP৫৪",
		"Mali-G57 MC2": "Mali-G৫৭ MC২",
		"March 2024": "March ২০২৪",
		"Mediatek Dimensity 6080 (6 nm)": "Mediatek Dimensity ৬০৮০ (৬ nm)",
	}
}

// Seed inserts specification_translations for existing specifications for product 'poco-x6-neo-5g'
func (s *SpecificationSeederMobilePocoX6Neo5g) Seed(db *gorm.DB) error {
	productSlug := "poco-x6-neo-5g"

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
