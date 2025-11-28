package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileIphone16 seeds specifications/options for product 'iphone-16'
type SpecificationSeederMobileIphone16 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileIphone16 creates a new seeder instance
func NewSpecificationSeederMobileIphone16() *SpecificationSeederMobileIphone16 {
	return &SpecificationSeederMobileIphone16{BaseSeeder: BaseSeeder{name: "Specifications for iPhone 16"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileIphone16) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1179 x 2556 pixels": "১১৭৯ x ২৫৫৬ পিক্সেল",
		"12 MP": "১২ মেগাপিক্সেল",
		"128 GB / 256 GB / 512 GB": "১২৮ জিবি / ২৫৬ জিবি / ৫১২ GB",
		"147.6 x 71.6 x 7.8 mm": "১৪৭.৬ x ৭১.৬ x ৭.৮ মিমি",
		"170 g": "১৭০ g",
		"3,561 mAh": "৩,৫৬১ এমএএইচ",
		"48 MP + 12 MP": "৪৮ মেগাপিক্সেল + ১২ মেগাপিক্সেল",
		"5-core Apple GPU": "৫-কোর অ্যাপল জিপিইউ",
		"5G": "৫G",
		"6.1 inches": "৬.১ ইঞ্চি",
		"60Hz": "৬০Hz",
		"8 GB": "৮ GB",
		"Aluminum frame, glass front/back": "অ্যালুমিনিয়াম ফ্রেম, গ্লাস সামনে/পেছনে",
		"Apple A18": "অ্যাপল A১৮",
		"Apple A18 (3 nm)": "অ্যাপল A১৮ (৩ ন্যানোমিটার)",
		"Black, White, Pink, Teal, Ultramarine": "কালো, সাদা, গোলাপী, টিল, আল্ট্রামেরিন",
		"IP68": "IP৬৮",
		"September 2024": "সেপ্টেম্বর ২০২৪",
		"iOS 18": "আইওএস ১৮",
	}
}

// Seed inserts specification records for the product identified by slug 'iphone-16'
func (s *SpecificationSeederMobileIphone16) Seed(db *gorm.DB) error {
	productSlug := "iphone-16"

	var prod models.ProductModel
	if err := db.Where("slug = ?", productSlug).First(&prod).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil
		}
		return err
	}
	productID := prod.ID

	specs := DefaultMobileSpecs()
	banglaTranslations := s.getBanglaTranslations()

	// Override model-specific values for iPhone 16
	specs["Display Size"] = "6.1 inches"
	specs["Processor"] = "Apple A18"
	specs["Chipset"] = "Apple A18 (3 nm)"
	specs["Cpu Type"] = "Hexa-core"
	specs["Gpu Type"] = "5-core Apple GPU"
	specs["Ram"] = "8 GB"
	specs["Storage"] = "128 GB / 256 GB / 512 GB"
	specs["Display Type"] = "Super Retina XDR OLED"
	specs["Resolution"] = "1179 x 2556 pixels"
	specs["Screen Protection"] = "Ceramic Shield"
	specs["Refresh Rate"] = "60Hz"
	specs["Build Material"] = "Aluminum frame, glass front/back"
	specs["Weight"] = "170 g"
	specs["Dimensions"] = "147.6 x 71.6 x 7.8 mm"
	specs["Water Resistance"] = "IP68"
	specs["Network Technology"] = "5G"
	specs["Rear Camera"] = "48 MP + 12 MP"
	specs["Front Camera"] = "12 MP"
	specs["Battery"] = "3,561 mAh"
	specs["Operating System"] = "iOS 18"
	specs["Available Colors"] = "Black, White, Pink, Teal, Ultramarine"
	specs["Announcement Date"] = "September 2024"
	specs["Device Status"] = "Available"

	for key, value := range specs {
		sk, err := CreateOrFindSpecificationKey(db, key)
		if err != nil {
			return err
		}

		var existing models.SpecificationModel
		if err := db.Where("product_id = ? AND specification_key_id = ?", productID, sk.ID).First(&existing).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				sModel := &models.SpecificationModel{
					ProductID:          productID,
					SpecificationKeyID: sk.ID,
					Value:              value,
					Status:             1,
				}
				if err := db.Create(sModel).Error; err != nil {
					return err
				}

				// Create Bangla translation for the specification
				banglaValue, exists := banglaTranslations[value]
				if exists && banglaValue != "" {
					var existingTranslation models.SpecificationTranslationModel
					if err := db.Where("specification_id = ? AND locale = ?", sModel.ID, "bn").First(&existingTranslation).Error; err != nil {
						if err == gorm.ErrRecordNotFound {
							translation := &models.SpecificationTranslationModel{
								SpecificationID: sModel.ID,
								Locale:          "bn",
								Value:           banglaValue,
							}
							if err := db.Create(translation).Error; err != nil {
								return err
							}
						} else {
							return err
						}
					}
				}
			} else {
				return err
			}
		} else {
			// If specification already exists, check and create Bangla translation if missing
			banglaValue, exists := banglaTranslations[value]
			if exists && banglaValue != "" {
				var existingTranslation models.SpecificationTranslationModel
				if err := db.Where("specification_id = ? AND locale = ?", existing.ID, "bn").First(&existingTranslation).Error; err != nil {
					if err == gorm.ErrRecordNotFound {
						translation := &models.SpecificationTranslationModel{
							SpecificationID: existing.ID,
							Locale:          "bn",
							Value:           banglaValue,
						}
						if err := db.Create(translation).Error; err != nil {
							return err
						}
					} else {
						return err
					}
				}
			}
		}
	}

	return nil
}
