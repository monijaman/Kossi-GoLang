package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileIphone17 seeds specifications/options for product 'iphone-17'
type SpecificationSeederMobileIphone17 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileIphone17 creates a new seeder instance
func NewSpecificationSeederMobileIphone17() *SpecificationSeederMobileIphone17 {
	return &SpecificationSeederMobileIphone17{BaseSeeder: BaseSeeder{name: "Specifications for iPhone 17"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileIphone17) getBanglaTranslations() map[string]string {
	return map[string]string{
		"12 MP": "১২ মেগাপিক্সেল",
		"128 GB / 256 GB / 512 GB": "১২৮ জিবি / ২৫৬ জিবি / ৫১২ GB",
		"147.6 x 71.6 x 7.8 mm": "১৪৭.৬ x ৭১.৬ x ৭.৮ মিমি",
		"171 g": "১৭১ g",
		"2556 x 1179 pixels (~461 ppi density)": "২৫৫৬ x ১১৭৯ পিক্সেল (~৪৬১ পিপিআই density)",
		"3,349 mAh": "৩,৩৪৯ এমএএইচ",
		"48 MP + 12 MP": "৪৮ মেগাপিক্সেল + ১২ মেগাপিক্সেল",
		"5-core Apple GPU": "৫-কোর অ্যাপল জিপিইউ",
		"5G": "৫G",
		"6.1 inches": "৬.১ ইঞ্চি",
		"60Hz": "৬০Hz",
		"8 GB": "৮ GB",
		"Apple A19": "অ্যাপল A১৯",
		"Apple A19 (3 nm)": "অ্যাপল A১৯ (৩ ন্যানোমিটার)",
		"Black, Blue, Green, Yellow, Pink": "কালো, নীল, সবুজ, হলুদ, গোলাপী",
		"Glass front/back, aluminum frame": "গ্লাস সামনে/পেছনে, অ্যালুমিনিয়াম ফ্রেম",
		"IP68": "IP৬৮",
		"September 2024": "সেপ্টেম্বর ২০২৪",
		"Super Retina XDR OLED, HDR10, Dolby Vision": "সুপার রেটিনা এক্সডিআর ওএলইডি, এইচডিআর১০, ডলবি ভিশন",
		"iOS 18": "আইওএস ১৮",
	}
}

// Seed inserts specification records for the product identified by slug 'iphone-17'
func (s *SpecificationSeederMobileIphone17) Seed(db *gorm.DB) error {
	productSlug := "iphone-17"

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

	// Override model-specific values for iPhone 17
	specs["Display Size"] = "6.1 inches"
	specs["Processor"] = "Apple A19"
	specs["Chipset"] = "Apple A19 (3 nm)"
	specs["Cpu Type"] = "Hexa-core"
	specs["Gpu Type"] = "5-core Apple GPU"
	specs["Ram"] = "8 GB"
	specs["Storage"] = "128 GB / 256 GB / 512 GB"
	specs["Display Type"] = "Super Retina XDR OLED, HDR10, Dolby Vision"
	specs["Resolution"] = "2556 x 1179 pixels (~461 ppi density)"
	specs["Screen Protection"] = "Ceramic Shield"
	specs["Refresh Rate"] = "60Hz"
	specs["Build Material"] = "Glass front/back, aluminum frame"
	specs["Weight"] = "171 g"
	specs["Dimensions"] = "147.6 x 71.6 x 7.8 mm"
	specs["Water Resistance"] = "IP68"
	specs["Network Technology"] = "5G"
	specs["Rear Camera"] = "48 MP + 12 MP"
	specs["Front Camera"] = "12 MP"
	specs["Battery"] = "3,349 mAh"
	specs["Operating System"] = "iOS 18"
	specs["Available Colors"] = "Black, Blue, Green, Yellow, Pink"
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
