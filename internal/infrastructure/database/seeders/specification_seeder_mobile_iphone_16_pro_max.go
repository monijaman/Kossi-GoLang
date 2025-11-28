package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileIphone16ProMax seeds specifications/options for product 'iphone-16-pro-max'
type SpecificationSeederMobileIphone16ProMax struct {
	BaseSeeder
}

// NewSpecificationSeederMobileIphone16ProMax creates a new seeder instance
func NewSpecificationSeederMobileIphone16ProMax() *SpecificationSeederMobileIphone16ProMax {
	return &SpecificationSeederMobileIphone16ProMax{BaseSeeder: BaseSeeder{name: "Specifications for iPhone 16 Pro Max"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileIphone16ProMax) getBanglaTranslations() map[string]string {
	return map[string]string{
		"12 MP": "১২ মেগাপিক্সেল",
		"120Hz": "১২০Hz",
		"1320 x 2868 pixels": "১৩২০ x ২৮৬৮ পিক্সেল",
		"163.0 x 77.6 x 8.25 mm": "১৬৩.০ x ৭৭.৬ x ৮.২৫ মিমি",
		"225 g": "২২৫ g",
		"256 GB / 512 GB / 1 TB": "২৫৬ জিবি / ৫১২ জিবি / ১ টিবি",
		"4,676 mAh": "৪,৬৭৬ এমএএইচ",
		"48 MP + 48 MP + 12 MP": "৪৮ মেগাপিক্সেল + ৪৮ মেগাপিক্সেল + ১২ মেগাপিক্সেল",
		"5G": "৫G",
		"6-core Apple GPU": "৬-কোর অ্যাপল জিপিইউ",
		"6.9 inches": "৬.৯ ইঞ্চি",
		"8 GB": "৮ GB",
		"Apple A18 Pro": "অ্যাপল A১৮ প্রো",
		"Apple A18 Pro (3 nm)": "অ্যাপল A১৮ প্রো (৩ ন্যানোমিটার)",
		"Black Titanium, White Titanium, Natural Titanium, Desert Titanium": "কালো টাইটানিয়াম, সাদা টাইটানিয়াম, ন্যাচারাল টাইটানিয়াম, মরুভূমি টাইটানিয়াম",
		"IP68": "IP৬৮",
		"LTPO Super Retina XDR OLED, 120Hz": "এলটিপিও সুপার রেটিনা এক্সডিআর ওএলইডি, ১২০Hz",
		"September 2024": "সেপ্টেম্বর ২০২৪",
		"Titanium frame, glass front/back": "টাইটানিয়াম ফ্রেম, গ্লাস সামনে/পেছনে",
		"iOS 18": "আইওএস ১৮",
	}
}

// Seed inserts specification records for the product identified by slug 'iphone-16-pro-max'
func (s *SpecificationSeederMobileIphone16ProMax) Seed(db *gorm.DB) error {
	productSlug := "iphone-16-pro-max"

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

	// Override model-specific values for iPhone 16 Pro Max
	specs["Display Size"] = "6.9 inches"
	specs["Processor"] = "Apple A18 Pro"
	specs["Chipset"] = "Apple A18 Pro (3 nm)"
	specs["Cpu Type"] = "Hexa-core"
	specs["Gpu Type"] = "6-core Apple GPU"
	specs["Ram"] = "8 GB"
	specs["Storage"] = "256 GB / 512 GB / 1 TB"
	specs["Display Type"] = "LTPO Super Retina XDR OLED, 120Hz"
	specs["Resolution"] = "1320 x 2868 pixels"
	specs["Screen Protection"] = "Ceramic Shield"
	specs["Refresh Rate"] = "120Hz"
	specs["Build Material"] = "Titanium frame, glass front/back"
	specs["Weight"] = "225 g"
	specs["Dimensions"] = "163.0 x 77.6 x 8.25 mm"
	specs["Water Resistance"] = "IP68"
	specs["Network Technology"] = "5G"
	specs["Rear Camera"] = "48 MP + 48 MP + 12 MP"
	specs["Front Camera"] = "12 MP"
	specs["Battery"] = "4,676 mAh"
	specs["Operating System"] = "iOS 18"
	specs["Available Colors"] = "Black Titanium, White Titanium, Natural Titanium, Desert Titanium"
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
