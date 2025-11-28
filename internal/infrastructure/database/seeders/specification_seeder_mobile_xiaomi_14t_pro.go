package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileXiaomi14tPro seeds specifications/options for product 'xiaomi-14t-pro'
type SpecificationSeederMobileXiaomi14tPro struct {
	BaseSeeder
}

// NewSpecificationSeederMobileXiaomi14tPro creates a new seeder instance
func NewSpecificationSeederMobileXiaomi14tPro() *SpecificationSeederMobileXiaomi14tPro {
	return &SpecificationSeederMobileXiaomi14tPro{BaseSeeder: BaseSeeder{name: "Specifications for Xiaomi 14T Pro"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileXiaomi14tPro) getBanglaTranslations() map[string]string {
	return map[string]string{
		"12 GB / 16 GB": "১২ GB / ১৬ GB",
		"120W wired, 50W wireless": "১২০W তারযুক্ত, ৫০W বেতার",
		"1220 x 2712 pixels": "১২২০ x ২৭১২ pixels",
		"144Hz": "১৪৪Hz",
		"209 g": "২০৯ g",
		"256 GB / 512 GB / 1 টিবি": "২৫৬ GB / ৫১২ GB / ১ টিবি",
		"32 MP": "৩২ MP",
		"50 MP + 50 MP + 12 MP": "৫০ MP + ৫০ MP + ১২ MP",
		"5000 mAh": "৫০০০ এমএএইচ",
		"6.67 inches": "৬.৬৭ ইঞ্চি",
		"AMOLED, 144Hz, Dolby Vision, HDR10+": "AMOLED, ১৪৪Hz, Dolby Vision, HDR১০+",
		"Android 14, HyperOS": "Android ১৪, HyperOS",
		"Dimensity 9300+": "Dimensity ৯৩০০+",
		"GSM / HSPA / LTE / 5G": "GSM / HSPA / LTE / ৫G",
		"Glass front, aluminum frame, glass back": "গ্লাস সামনে, অ্যালুমিনিয়াম ফ্রেম, গ্লাস পেছনে",
		"IP68": "IP৬৮",
		"Immortalis-G720 MC12": "Immortalis-G৭২০ MC১২",
		"Mediatek Dimensity 9300+ (4 nm)": "Mediatek Dimensity ৯৩০০+ (৪ nm)",
		"September 2024": "September ২০২৪",
		"Titan Gray, Titan Blue, Titan Black": "Titan ধূসর, Titan নীল, Titan কালো",
	}
}

// Seed inserts specification_translations for existing specifications for product 'xiaomi-14t-pro'
func (s *SpecificationSeederMobileXiaomi14tPro) Seed(db *gorm.DB) error {
	productSlug := "xiaomi-14t-pro"

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
