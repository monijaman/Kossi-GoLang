package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileRedmiNote15Pro5g seeds specifications/options for product 'redmi-note-15-pro-5g'
type SpecificationSeederMobileRedmiNote15Pro5g struct {
	BaseSeeder
}

// NewSpecificationSeederMobileRedmiNote15Pro5g creates a new seeder instance
func NewSpecificationSeederMobileRedmiNote15Pro5g() *SpecificationSeederMobileRedmiNote15Pro5g {
	return &SpecificationSeederMobileRedmiNote15Pro5g{BaseSeeder: BaseSeeder{name: "Specifications for Redmi Note 15 Pro 5G"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileRedmiNote15Pro5g) getBanglaTranslations() map[string]string {
	return map[string]string{
		"120Hz": "১২০Hz",
		"1220 x 2712 pixels": "১২২০ x ২৭১২ pixels",
		"16 MP": "১৬ MP",
		"161.2 x 74.3 x 8 mm": "১৬১.২ x ৭৪.৩ x ৮ মিমি",
		"187 g": "১৮৭ g",
		"200 MP + 8 MP + 2 MP": "২০০ MP + ৮ MP + ২ MP",
		"256 GB / 512 GB": "২৫৬ GB / ৫১২ GB",
		"5,100 mAh": "৫,১০০ এমএএইচ",
		"5G": "৫G",
		"6.67 inches": "৬.৬৭ ইঞ্চি",
		"8 GB / 12 GB": "৮ GB / ১২ GB",
		"AMOLED, 68B colors, 120Hz, Dolby Vision, 1800 nits": "AMOLED, ৬৮B colors, ১২০Hz, Dolby Vision, ১৮০০ nits",
		"Adreno 710": "Adreno ৭১০",
		"Android 15, HyperOS": "Android ১৫, HyperOS",
		"Glass front/back, plastic frame": "গ্লাস সামনে/back, plastic frame",
		"IP68": "IP৬৮",
		"মিডনাইট Black, Aurora Purple, Ocean Teal": "মিডনাইট কালো, Aurora বেগুনি, Ocean টিল",
		"Qualcomm Snapdragon 7s Gen 3 (4 nm)": "Qualcomm Snapdragon ৭s Gen ৩ (৪ nm)",
		"September 2025": "September ২০২৫",
		"Snapdragon 7s Gen 3": "Snapdragon ৭s Gen ৩",
	}
}

// Seed inserts specification_translations for existing specifications for product 'redmi-note-15-pro-5g'
func (s *SpecificationSeederMobileRedmiNote15Pro5g) Seed(db *gorm.DB) error {
	productSlug := "redmi-note-15-pro-5g"

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
