package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileSymphonyXplorerZv seeds specifications/options for product 'symphony-xplorer-zv'
type SpecificationSeederMobileSymphonyXplorerZv struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSymphonyXplorerZv creates a new seeder instance
func NewSpecificationSeederMobileSymphonyXplorerZv() *SpecificationSeederMobileSymphonyXplorerZv {
	return &SpecificationSeederMobileSymphonyXplorerZv{BaseSeeder: BaseSeeder{name: "Specifications for Symphony Xplorer ZV"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSymphonyXplorerZv) getBanglaTranslations() map[string]string {
	return map[string]string{
		"13 MP": "১৩ মেগাপিক্সেল",
		"130 g": "১৩০ g",
		"142.5 x 70.8 x 6.8 mm": "১৪২.৫ x ৭০.৮ x ৬.৮ মিমি",
		"16 GB": "১৬ GB",
		"2 GB": "২ GB",
		"2 MP": "২ মেগাপিক্সেল",
		"2,050 mAh": "২,০৫০ এমএএইচ",
		"2014": "২০১৪",
		"3G": "৩G",
		"5.0 inches": "৫.০ ইঞ্চি",
		"60Hz": "৬০Hz",
		"720 x 1280 pixels": "৭২০ x ১২৮০ পিক্সেল",
		"Android 4.4.2 (KitKat)": "অ্যান্ড্রয়েড ৪.৪.২ (KitKat)",
		"Black, White": "কালো, সাদা",
		"Gorilla Glass 3": "গরিলা গ্লাস ৩",
		"Mali-450MP4": "মালি-৪৫০MP৪",
		"Mediatek MT6592": "মিডিয়াটেক MT৬৫৯২",
		"Octa-core 1.4GHz": "অক্টা-কোর ১.৪GHz",
	}
}

// Seed inserts specification records for the product identified by slug 'symphony-xplorer-zv'
func (s *SpecificationSeederMobileSymphonyXplorerZv) Seed(db *gorm.DB) error {
	productSlug := "symphony-xplorer-zv"

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

	// Override model-specific values for Symphony Xplorer ZV
	specs["Display Size"] = "5.0 inches"
	specs["Processor"] = "Octa-core 1.4GHz"
	specs["Chipset"] = "Mediatek MT6592"
	specs["Cpu Type"] = "Octa-core"
	specs["Gpu Type"] = "Mali-450MP4"
	specs["Ram"] = "2 GB"
	specs["Storage"] = "16 GB"
	specs["Display Type"] = "AMOLED"
	specs["Resolution"] = "720 x 1280 pixels"
	specs["Screen Protection"] = "Gorilla Glass 3"
	specs["Refresh Rate"] = "60Hz"
	specs["Build Material"] = "Plastic body"
	specs["Weight"] = "130 g"
	specs["Dimensions"] = "142.5 x 70.8 x 6.8 mm"
	specs["Water Resistance"] = "No"
	specs["Network Technology"] = "3G"
	specs["Rear Camera"] = "13 MP"
	specs["Front Camera"] = "2 MP"
	specs["Battery"] = "2,050 mAh"
	specs["Operating System"] = "Android 4.4.2 (KitKat)"
	specs["Available Colors"] = "Black, White"
	specs["Announcement Date"] = "2014"
	specs["Device Status"] = "Discontinued"

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
