package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileRedmiNote15Pro seeds specifications/options for product 'redmi-note-15-pro'
type SpecificationSeederMobileRedmiNote15Pro struct {
	BaseSeeder
}

// NewSpecificationSeederMobileRedmiNote15Pro creates a new seeder instance
func NewSpecificationSeederMobileRedmiNote15Pro() *SpecificationSeederMobileRedmiNote15Pro {
	return &SpecificationSeederMobileRedmiNote15Pro{BaseSeeder: BaseSeeder{name: "Specifications for Redmi Note 15 Pro+"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileRedmiNote15Pro) getBanglaTranslations() map[string]string {
	return map[string]string{
		"12 GB / 16 GB": "১২ GB / ১৬ GB",
		"120Hz": "১২০Hz",
		"120W wired": "১২০W তারযুক্ত",
		"1220 x 2712 pixels": "১২২০ x ২৭১২ pixels",
		"16 MP": "১৬ MP",
		"200 MP + 8 MP + 2 MP": "২০০ MP + ৮ MP + ২ MP",
		"205 g": "২০৫ g",
		"256 GB / 512 GB": "২৫৬ GB / ৫১২ GB",
		"5000 mAh": "৫০০০ এমএএইচ",
		"6.67 inches": "৬.৬৭ ইঞ্চি",
		"AMOLED, 120Hz, Dolby Vision, HDR10+": "AMOLED, ১২০Hz, Dolby Vision, HDR১০+",
		"Android 15, HyperOS": "Android ১৫, HyperOS",
		"Dimensity 8300 আল্ট্রা": "Dimensity ৮৩০০ আল্ট্রা",
		"GSM / HSPA / LTE / 5G": "GSM / HSPA / LTE / ৫G",
		"Glass front (Gorilla Glass Victus), glass back, aluminum frame": "গ্লাস সামনে (Gorilla Glass Victus), গ্লাস পেছনে, অ্যালুমিনিয়াম ফ্রেম",
		"IP68": "IP৬৮",
		"Mali-G615-MC6": "Mali-G৬১৫-MC৬",
		"Mediatek Dimensity 8300 আল্ট্রা (4 nm)": "Mediatek Dimensity ৮৩০০ আল্ট্রা (৪ nm)",
		"মিডনাইট Black, মুনলাইট White, Aurora Purple": "মিডনাইট কালো, মুনলাইট সাদা, Aurora বেগুনি",
		"September 2025": "September ২০২৫",
	}
}

// Seed inserts specification_translations for existing specifications for product 'redmi-note-15-pro'
func (s *SpecificationSeederMobileRedmiNote15Pro) Seed(db *gorm.DB) error {
	productSlug := "redmi-note-15-pro"

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
