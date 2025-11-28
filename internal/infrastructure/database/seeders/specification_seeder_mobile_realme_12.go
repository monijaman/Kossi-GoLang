package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileRealme12 seeds specifications/options for product 'realme-12'
type SpecificationSeederMobileRealme12 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileRealme12 creates a new seeder instance
func NewSpecificationSeederMobileRealme12() *SpecificationSeederMobileRealme12 {
	return &SpecificationSeederMobileRealme12{BaseSeeder: BaseSeeder{name: "Specifications for Realme 12"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileRealme12) getBanglaTranslations() map[string]string {
	return map[string]string{
		"108 MP + 2 MP": "১০৮ MP + ২ MP",
		"1080 x 2400 pixels": "১০৮০ x ২৪০০ pixels",
		"120Hz": "১২০Hz",
		"128 GB": "১২৮ GB",
		"188 g": "১৮৮ g",
		"45W wired": "৪৫W তারযুক্ত",
		"5000 mAh": "৫০০০ এমএএইচ",
		"6 GB / 8 GB": "৬ GB / ৮ GB",
		"6.72 inches": "৬.৭২ ইঞ্চি",
		"8 MP": "৮ MP",
		"Android 14, Realme UI 5.0": "Android ১৪, Realme UI ৫.০",
		"Dimensity 6100+": "Dimensity ৬১০০+",
		"GSM / HSPA / LTE / 5G": "GSM / HSPA / LTE / ৫G",
		"Glass front, plastic back, plastic frame": "গ্লাস সামনে, প্লাস্টিক পেছনে, plastic frame",
		"IP54": "IP৫৪",
		"IPS LCD, 120Hz, 950 nits": "IPS LCD, ১২০Hz, ৯৫০ nits",
		"Mali-G57 MC2": "Mali-G৫৭ MC২",
		"March 2024": "March ২০২৪",
		"Mediatek Dimensity 6100+ (6 nm)": "Mediatek Dimensity ৬১০০+ (৬ nm)",
		"টোয়াইলাইট Purple, Woodland Green": "টোয়াইলাইট বেগুনি, Woodland সবুজ",
	}
}

// Seed inserts specification_translations for existing specifications for product 'realme-12'
func (s *SpecificationSeederMobileRealme12) Seed(db *gorm.DB) error {
	productSlug := "realme-12"

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
