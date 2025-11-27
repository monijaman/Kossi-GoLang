package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileIphone15 seeds specifications/options for product 'iphone-15'
type SpecificationSeederMobileIphone15 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileIphone15 creates a new seeder instance
func NewSpecificationSeederMobileIphone15() *SpecificationSeederMobileIphone15 {
	return &SpecificationSeederMobileIphone15{BaseSeeder: BaseSeeder{name: "Specifications for iPhone 15"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileIphone15) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1179 x 2556 pixels": "১১৭৯ x ২৫৫৬ pixels",
		"12 MP": "১২ MP",
		"128 GB / 256 GB / 512 GB": "১২৮ GB / ২৫৬ GB / ৫১২ GB",
		"147.6 x 71.6 x 7.8 mm": "১৪৭.৬ x ৭১.৬ x ৭.৮ মিমি",
		"171 g": "১৭১ g",
		"3,349 mAh": "৩,৩৪৯ এমএএইচ",
		"48 MP + 12 MP": "৪৮ MP + ১২ MP",
		"5-core Apple GPU": "৫-core Apple GPU",
		"5G": "৫G",
		"6 GB": "৬ GB",
		"6.1 inches": "৬.১ ইঞ্চি",
		"60Hz": "৬০Hz",
		"Aluminum frame, glass front/back": "অ্যালুমিনিয়াম ফ্রেম, গ্লাস সামনে/back",
		"Apple A16 Bionic": "Apple A১৬ Bionic",
		"Apple A16 Bionic (4 nm)": "Apple A১৬ Bionic (৪ nm)",
		"Black, Blue, Green, Yellow, Pink": "কালো, নীল, সবুজ, হলুদ, গোলাপী",
		"IP68": "IP৬৮",
		"September 2023": "September ২০২৩",
		"iOS 17, upgradable": "iOS ১৭, upgradable",
	}
}

// Seed inserts specification records for the product identified by slug 'iphone-15'
func (s *SpecificationSeederMobileIphone15) Seed(db *gorm.DB) error {
	productSlug := "iphone-15"

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

	// Override model-specific values for iPhone 15
	specs["Display Size"] = "6.1 inches"
	specs["Processor"] = "Apple A16 Bionic"
	specs["Chipset"] = "Apple A16 Bionic (4 nm)"
	specs["Cpu Type"] = "Hexa-core"
	specs["Gpu Type"] = "5-core Apple GPU"
	specs["Ram"] = "6 GB"
	specs["Storage"] = "128 GB / 256 GB / 512 GB"
	specs["Display Type"] = "Super Retina XDR OLED"
	specs["Resolution"] = "1179 x 2556 pixels"
	specs["Screen Protection"] = "Ceramic Shield"
	specs["Refresh Rate"] = "60Hz"
	specs["Build Material"] = "Aluminum frame, glass front/back"
	specs["Weight"] = "171 g"
	specs["Dimensions"] = "147.6 x 71.6 x 7.8 mm"
	specs["Water Resistance"] = "IP68"
	specs["Network Technology"] = "5G"
	specs["Rear Camera"] = "48 MP + 12 MP"
	specs["Front Camera"] = "12 MP"
	specs["Battery"] = "3,349 mAh"
	specs["Operating System"] = "iOS 17, upgradable"
	specs["Available Colors"] = "Black, Blue, Green, Yellow, Pink"
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
