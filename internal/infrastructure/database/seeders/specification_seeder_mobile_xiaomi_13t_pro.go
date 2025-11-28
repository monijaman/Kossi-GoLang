package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileXiaomi13tPro seeds specifications/options for product 'xiaomi-13t-pro'
type SpecificationSeederMobileXiaomi13tPro struct {
	BaseSeeder
}

// NewSpecificationSeederMobileXiaomi13tPro creates a new seeder instance
func NewSpecificationSeederMobileXiaomi13tPro() *SpecificationSeederMobileXiaomi13tPro {
	return &SpecificationSeederMobileXiaomi13tPro{BaseSeeder: BaseSeeder{name: "Specifications for Xiaomi 13T Pro"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileXiaomi13tPro) getBanglaTranslations() map[string]string {
	return map[string]string{
		"12 GB / 16 GB": "১২ GB / ১৬ GB",
		"120W wired": "১২০W তারযুক্ত",
		"1220 x 2712 pixels": "১২২০ x ২৭১২ pixels",
		"144Hz": "১৪৪Hz",
		"20 MP": "২০ MP",
		"200 g": "২০০ g",
		"256 GB / 512 GB / 1 টিবি": "২৫৬ GB / ৫১২ GB / ১ টিবি",
		"50 MP + 50 MP + 12 MP": "৫০ MP + ৫০ MP + ১২ MP",
		"5000 mAh": "৫০০০ এমএএইচ",
		"6.67 inches": "৬.৬৭ ইঞ্চি",
		"AMOLED, 144Hz, Dolby Vision, HDR10+": "AMOLED, ১৪৪Hz, Dolby Vision, HDR১০+",
		"আলপাইন Blue, মিডো Green, Black": "আলপাইন নীল, মিডো সবুজ, কালো",
		"Android 13, আপগ্রেডযোগ্য Android 14, HyperOS": "Android ১৩, আপগ্রেডযোগ্য Android ১৪, HyperOS",
		"Dimensity 9200+": "Dimensity ৯২০০+",
		"GSM / HSPA / LTE / 5G": "GSM / HSPA / LTE / ৫G",
		"Glass front (Gorilla Glass 5), glass back বা সিলিকন পলিমার back, plastic frame": "গ্লাস সামনে (Gorilla Glass ৫), গ্লাস পেছনে বা সিলিকন পলিমার back, plastic frame",
		"IP68": "IP৬৮",
		"Immortalis-G715 MC11": "Immortalis-G৭১৫ MC১১",
		"Mediatek Dimensity 9200+ (4 nm)": "Mediatek Dimensity ৯২০০+ (৪ nm)",
		"September 2023": "September ২০২৩",
	}
}

// Seed inserts specification_translations for existing specifications for product 'xiaomi-13t-pro'
func (s *SpecificationSeederMobileXiaomi13tPro) Seed(db *gorm.DB) error {
	productSlug := "xiaomi-13t-pro"

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
