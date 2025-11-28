package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileOneplusOpen seeds specifications/options for product 'oneplus-open'
type SpecificationSeederMobileOneplusOpen struct {
	BaseSeeder
}

// NewSpecificationSeederMobileOneplusOpen creates a new seeder instance
func NewSpecificationSeederMobileOneplusOpen() *SpecificationSeederMobileOneplusOpen {
	return &SpecificationSeederMobileOneplusOpen{BaseSeeder: BaseSeeder{name: "Specifications for OnePlus Open"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileOneplusOpen) getBanglaTranslations() map[string]string {
	return map[string]string{
		"120Hz": "১২০Hz",
		"16 GB": "১৬ GB",
		"20 MP (Inner) + 32 MP (Cover)": "২০ MP (Inner) + ৩২ MP (Cover)",
		"2268 x 2440 pixels": "২২৬৮ x ২৪৪০ pixels",
		"239 g": "২৩৯ g",
		"4,805 mAh": "৪,৮০৫ এমএএইচ",
		"48 MP + 64 MP + 48 MP": "৪৮ MP + ৬৪ MP + ৪৮ MP",
		"512 GB": "৫১২ GB",
		"5G": "৫G",
		"7.82 inches (Main) / 6.31 inches (Cover)": "৭.৮২ ইঞ্চি (Main) / ৬.৩১ ইঞ্চি (Cover)",
		"Adreno 740": "Adreno ৭৪০",
		"Android 13, OxygenOS 13.2": "Android ১৩, OxygenOS ১৩.২",
		"Emerald Dusk, Voyager Black": "Emerald Dusk, Voyager কালো",
		"Foldable LTPO3 Flexi-fluid AMOLED, 120Hz, Dolby Vision, 2800 nits": "Foldable LTPO৩ Flexi-fluid AMOLED, ১২০Hz, Dolby Vision, ২৮০০ nits",
		"Glass front (Ceramic Guard), glass back, aluminum frame": "গ্লাস সামনে (Ceramic Guard), গ্লাস পেছনে, অ্যালুমিনিয়াম ফ্রেম",
		"IPX4": "IPX৪",
		"October 2023": "October ২০২৩",
		"Qualcomm SM8550-AB Snapdragon 8 Gen 2 (4 nm)": "Qualcomm SM৮৫৫০-AB Snapdragon ৮ Gen ২ (৪ nm)",
		"Snapdragon 8 Gen 2": "Snapdragon ৮ Gen ২",
		"Unfolded: 153.4 x 143.1 x 5.8 mm / Folded: 153.4 x 73.3 x 11.7 mm": "Unfolded: ১৫৩.৪ x ১৪৩.১ x ৫.৮ মিমি / Folded: ১৫৩.৪ x ৭৩.৩ x ১১.৭ মিমি",
	}
}

// Seed inserts specification_translations for existing specifications for product 'oneplus-open'
func (s *SpecificationSeederMobileOneplusOpen) Seed(db *gorm.DB) error {
	productSlug := "oneplus-open"

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
