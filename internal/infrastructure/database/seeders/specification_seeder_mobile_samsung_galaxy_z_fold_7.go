package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileSamsungGalaxyZFold7 seeds specifications/options for product 'samsung-galaxy-z-fold-7'
type SpecificationSeederMobileSamsungGalaxyZFold7 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSamsungGalaxyZFold7 creates a new seeder instance
func NewSpecificationSeederMobileSamsungGalaxyZFold7() *SpecificationSeederMobileSamsungGalaxyZFold7 {
	return &SpecificationSeederMobileSamsungGalaxyZFold7{BaseSeeder: BaseSeeder{name: "Specifications for Samsung Galaxy Z Fold 7"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSamsungGalaxyZFold7) getBanglaTranslations() map[string]string {
	return map[string]string{
		"12 GB / 16 GB": "১২ GB / ১৬ GB",
		"120Hz": "১২০Hz",
		"1856 x 2160 pixels (Main)": "১৮৫৬ x ২১৬০ pixels (Main)",
		"235 g": "২৩৫ g",
		"256 GB / 512 GB / 1 TB": "২৫৬ GB / ৫১২ GB / ১ TB",
		"4 MP (Under Display) + 10 MP (Cover)": "৪ MP (ডিসপ্লের নিচে) + ১০ MP (Cover)",
		"4,600 mAh": "৪,৬০০ এমএএইচ",
		"50 MP + 12 MP + 10 MP": "৫০ MP + ১২ MP + ১০ MP",
		"5G": "৫G",
		"7.6 inches (Main) / 6.3 inches (Cover)": "৭.৬ ইঞ্চি (Main) / ৬.৩ ইঞ্চি (Cover)",
		"Adreno 830": "Adreno ৮৩০",
		"Android 15, One UI 7.1.1": "Android ১৫, One UI ৭.১.১",
		"Corning Gorilla Armor 2": "Corning Gorilla Armor ২",
		"Foldable Dynamic LTPO AMOLED 2X, 120Hz, HDR10+": "Foldable Dynamic LTPO AMOLED ২X, ১২০Hz, HDR১০+",
		"Glass front (Gorilla Armor 2), glass back (Victus 2), titanium frame": "গ্লাস সামনে (Gorilla Armor ২), গ্লাস পেছনে (Victus ২), টাইটানিয়াম ফ্রেম",
		"IPX8": "IPX৮",
		"July 2025": "July ২০২৫",
		"Phantom Black, Icy Blue, Cream": "Phantom কালো, Icy নীল, Cream",
		"Qualcomm SM8750-AC Snapdragon 8 Elite (3 nm)": "Qualcomm SM৮৭৫০-AC Snapdragon ৮ Elite (৩ nm)",
		"Snapdragon 8 Elite": "Snapdragon ৮ Elite",
		"Unfolded: 153.5 x 132.6 x 5.6 mm / Folded: 153.5 x 68.1 x 11.2 mm": "Unfolded: ১৫৩.৫ x ১৩২.৬ x ৫.৬ মিমি / Folded: ১৫৩.৫ x ৬৮.১ x ১১.২ মিমি",
	}
}

// Seed inserts specification_translations for existing specifications for product 'samsung-galaxy-z-fold-7'
func (s *SpecificationSeederMobileSamsungGalaxyZFold7) Seed(db *gorm.DB) error {
	productSlug := "samsung-galaxy-z-fold-7"

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
