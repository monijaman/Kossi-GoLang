package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileXiaomi14 seeds specifications/options for product 'xiaomi-14'
type SpecificationSeederMobileXiaomi14 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileXiaomi14 creates a new seeder instance
func NewSpecificationSeederMobileXiaomi14() *SpecificationSeederMobileXiaomi14 {
	return &SpecificationSeederMobileXiaomi14{BaseSeeder: BaseSeeder{name: "Specifications for Xiaomi 14"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileXiaomi14) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1200 x 2670 pixels": "১২০০ x ২৬৭০ pixels",
		"120Hz": "১২০Hz",
		"152.8 x 71.5 x 8.2 mm": "১৫২.৮ x ৭১.৫ x ৮.২ মিমি",
		"188 g": "১৮৮ g",
		"256 GB / 512 GB / 1 TB": "২৫৬ GB / ৫১২ GB / ১ TB",
		"32 MP": "৩২ MP",
		"4,610 mAh": "৪,৬১০ এমএএইচ",
		"50 MP + 50 MP + 50 MP": "৫০ MP + ৫০ MP + ৫০ MP",
		"5G": "৫G",
		"6.36 inches": "৬.৩৬ ইঞ্চি",
		"8 GB / 12 GB / 16 GB": "৮ GB / ১২ GB / ১৬ GB",
		"Adreno 750": "Adreno ৭৫০",
		"Android 14, HyperOS": "Android ১৪, HyperOS",
		"Black, White, Jade Green, Pink": "কালো, সাদা, Jade সবুজ, গোলাপী",
		"Glass front, glass/silicone polymer back, aluminum frame": "গ্লাস সামনে, glass/silicone polymer back, অ্যালুমিনিয়াম ফ্রেম",
		"IP68": "IP৬৮",
		"LTPO OLED, 120Hz, Dolby Vision, HDR10+, 3000 nits": "LTPO OLED, ১২০Hz, Dolby Vision, HDR১০+, ৩০০০ nits",
		"October 2023": "October ২০২৩",
		"Qualcomm SM8650-AB Snapdragon 8 Gen 3 (4 nm)": "Qualcomm SM৮৬৫০-AB Snapdragon ৮ Gen ৩ (৪ nm)",
		"Snapdragon 8 Gen 3": "Snapdragon ৮ Gen ৩",
	}
}

// Seed inserts specification records for the product identified by slug 'xiaomi-14'
func (s *SpecificationSeederMobileXiaomi14) Seed(db *gorm.DB) error {
	productSlug := "xiaomi-14"

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

	// Override model-specific values for Xiaomi 14
	specs["Display Size"] = "6.36 inches"
	specs["Processor"] = "Snapdragon 8 Gen 3"
	specs["Chipset"] = "Qualcomm SM8650-AB Snapdragon 8 Gen 3 (4 nm)"
	specs["Cpu Type"] = "Octa-core"
	specs["Gpu Type"] = "Adreno 750"
	specs["Ram"] = "8 GB / 12 GB / 16 GB"
	specs["Storage"] = "256 GB / 512 GB / 1 TB"
	specs["Display Type"] = "LTPO OLED, 120Hz, Dolby Vision, HDR10+, 3000 nits"
	specs["Resolution"] = "1200 x 2670 pixels"
	specs["Screen Protection"] = "Corning Gorilla Glass Victus"
	specs["Refresh Rate"] = "120Hz"
	specs["Build Material"] = "Glass front, glass/silicone polymer back, aluminum frame"
	specs["Weight"] = "188 g"
	specs["Dimensions"] = "152.8 x 71.5 x 8.2 mm"
	specs["Water Resistance"] = "IP68"
	specs["Network Technology"] = "5G"
	specs["Rear Camera"] = "50 MP + 50 MP + 50 MP"
	specs["Front Camera"] = "32 MP"
	specs["Battery"] = "4,610 mAh"
	specs["Operating System"] = "Android 14, HyperOS"
	specs["Available Colors"] = "Black, White, Jade Green, Pink"
	specs["Announcement Date"] = "October 2023"
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
