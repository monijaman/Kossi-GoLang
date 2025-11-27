package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileIphone15Pro seeds specifications/options for product 'iphone-15-pro'
type SpecificationSeederMobileIphone15Pro struct {
	BaseSeeder
}

// NewSpecificationSeederMobileIphone15Pro creates a new seeder instance
func NewSpecificationSeederMobileIphone15Pro() *SpecificationSeederMobileIphone15Pro {
	return &SpecificationSeederMobileIphone15Pro{BaseSeeder: BaseSeeder{name: "Specifications for iPhone 15 Pro"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileIphone15Pro) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1179 x 2556 pixels": "১১৭৯ x ২৫৫৬ pixels",
		"12 MP": "১২ MP",
		"120Hz": "১২০Hz",
		"128 GB / 256 GB / 512 GB / 1 TB": "১২৮ GB / ২৫৬ GB / ৫১২ GB / ১ TB",
		"146.6 x 70.6 x 8.3 mm": "১৪৬.৬ x ৭০.৬ x ৮.৩ মিমি",
		"187 g": "১৮৭ g",
		"3,274 mAh": "৩,২৭৪ এমএএইচ",
		"48 MP + 12 MP + 12 MP": "৪৮ MP + ১২ MP + ১২ MP",
		"5G": "৫G",
		"6-core Apple GPU": "৬-core Apple GPU",
		"6.1 inches": "৬.১ ইঞ্চি",
		"8 GB": "৮ GB",
		"Apple A17 Pro": "Apple A১৭ Pro",
		"Apple A17 Pro (3 nm)": "Apple A১৭ Pro (৩ nm)",
		"Black Titanium, White Titanium, Blue Titanium, Natural Titanium": "কালো Titanium, সাদা Titanium, নীল Titanium, ন্যাচারাল Titanium",
		"IP68": "IP৬৮",
		"LTPO Super Retina XDR OLED, 120Hz": "LTPO Super Retina XDR OLED, ১২০Hz",
		"September 2023": "September ২০২৩",
		"Titanium frame, glass front/back": "টাইটানিয়াম ফ্রেম, গ্লাস সামনে/back",
		"iOS 17, upgradable": "iOS ১৭, upgradable",
	}
}

// Seed inserts specification records for the product identified by slug 'iphone-15-pro'
func (s *SpecificationSeederMobileIphone15Pro) Seed(db *gorm.DB) error {
	productSlug := "iphone-15-pro"

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

	// Override model-specific values for iPhone 15 Pro
	specs["Display Size"] = "6.1 inches"
	specs["Processor"] = "Apple A17 Pro"
	specs["Chipset"] = "Apple A17 Pro (3 nm)"
	specs["Cpu Type"] = "Hexa-core"
	specs["Gpu Type"] = "6-core Apple GPU"
	specs["Ram"] = "8 GB"
	specs["Storage"] = "128 GB / 256 GB / 512 GB / 1 TB"
	specs["Display Type"] = "LTPO Super Retina XDR OLED, 120Hz"
	specs["Resolution"] = "1179 x 2556 pixels"
	specs["Screen Protection"] = "Ceramic Shield"
	specs["Refresh Rate"] = "120Hz"
	specs["Build Material"] = "Titanium frame, glass front/back"
	specs["Weight"] = "187 g"
	specs["Dimensions"] = "146.6 x 70.6 x 8.3 mm"
	specs["Water Resistance"] = "IP68"
	specs["Network Technology"] = "5G"
	specs["Rear Camera"] = "48 MP + 12 MP + 12 MP"
	specs["Front Camera"] = "12 MP"
	specs["Battery"] = "3,274 mAh"
	specs["Operating System"] = "iOS 17, upgradable"
	specs["Available Colors"] = "Black Titanium, White Titanium, Blue Titanium, Natural Titanium"
	specs["Announcement Date"] = "September 2023"
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
