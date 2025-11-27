package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileIphone16Pro seeds specifications/options for product 'iphone-16-pro'
type SpecificationSeederMobileIphone16Pro struct {
	BaseSeeder
}

// NewSpecificationSeederMobileIphone16Pro creates a new seeder instance
func NewSpecificationSeederMobileIphone16Pro() *SpecificationSeederMobileIphone16Pro {
	return &SpecificationSeederMobileIphone16Pro{BaseSeeder: BaseSeeder{name: "Specifications for iPhone 16 Pro"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileIphone16Pro) getBanglaTranslations() map[string]string {
	return map[string]string{
		"12 MP": "১২ MP",
		"1206 x 2622 pixels": "১২০৬ x ২৬২২ pixels",
		"120Hz": "১২০Hz",
		"128 GB / 256 GB / 512 GB / 1 TB": "১২৮ GB / ২৫৬ GB / ৫১২ GB / ১ TB",
		"149.6 x 71.5 x 8.25 mm": "১৪৯.৬ x ৭১.৫ x ৮.২৫ মিমি",
		"199 g": "১৯৯ g",
		"3,577 mAh": "৩,৫৭৭ এমএএইচ",
		"48 MP + 48 MP + 12 MP": "৪৮ MP + ৪৮ MP + ১২ MP",
		"5G": "৫G",
		"6-core Apple GPU": "৬-core Apple GPU",
		"6.3 inches": "৬.৩ ইঞ্চি",
		"8 GB": "৮ GB",
		"Apple A18 Pro": "Apple A১৮ Pro",
		"Apple A18 Pro (3 nm)": "Apple A১৮ Pro (৩ nm)",
		"Black Titanium, White Titanium, Natural Titanium, Desert Titanium": "কালো Titanium, সাদা Titanium, ন্যাচারাল Titanium, মরুভূমি Titanium",
		"IP68": "IP৬৮",
		"LTPO Super Retina XDR OLED, 120Hz": "LTPO Super Retina XDR OLED, ১২০Hz",
		"September 2024": "September ২০২৪",
		"Titanium frame, glass front/back": "টাইটানিয়াম ফ্রেম, গ্লাস সামনে/back",
		"iOS 18": "iOS ১৮",
	}
}

// Seed inserts specification records for the product identified by slug 'iphone-16-pro'
func (s *SpecificationSeederMobileIphone16Pro) Seed(db *gorm.DB) error {
	productSlug := "iphone-16-pro"

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

	// Override model-specific values for iPhone 16 Pro
	specs["Display Size"] = "6.3 inches"
	specs["Processor"] = "Apple A18 Pro"
	specs["Chipset"] = "Apple A18 Pro (3 nm)"
	specs["Cpu Type"] = "Hexa-core"
	specs["Gpu Type"] = "6-core Apple GPU"
	specs["Ram"] = "8 GB"
	specs["Storage"] = "128 GB / 256 GB / 512 GB / 1 TB"
	specs["Display Type"] = "LTPO Super Retina XDR OLED, 120Hz"
	specs["Resolution"] = "1206 x 2622 pixels"
	specs["Screen Protection"] = "Ceramic Shield"
	specs["Refresh Rate"] = "120Hz"
	specs["Build Material"] = "Titanium frame, glass front/back"
	specs["Weight"] = "199 g"
	specs["Dimensions"] = "149.6 x 71.5 x 8.25 mm"
	specs["Water Resistance"] = "IP68"
	specs["Network Technology"] = "5G"
	specs["Rear Camera"] = "48 MP + 48 MP + 12 MP"
	specs["Front Camera"] = "12 MP"
	specs["Battery"] = "3,577 mAh"
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
