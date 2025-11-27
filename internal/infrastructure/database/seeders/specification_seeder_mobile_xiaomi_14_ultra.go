package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileXiaomi14Ultra seeds specifications/options for product 'xiaomi-14-ultra'
type SpecificationSeederMobileXiaomi14Ultra struct {
	BaseSeeder
}

// NewSpecificationSeederMobileXiaomi14Ultra creates a new seeder instance
func NewSpecificationSeederMobileXiaomi14Ultra() *SpecificationSeederMobileXiaomi14Ultra {
	return &SpecificationSeederMobileXiaomi14Ultra{BaseSeeder: BaseSeeder{name: "Specifications for Xiaomi 14 Ultra"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileXiaomi14Ultra) getBanglaTranslations() map[string]string {
	return map[string]string{
		"12 GB / 16 GB": "১২ GB / ১৬ GB",
		"120Hz": "১২০Hz",
		"1440 x 3200 pixels": "১৪৪০ x ৩২০০ pixels",
		"161.4 x 75.3 x 9.2 mm": "১৬১.৪ x ৭৫.৩ x ৯.২ মিমি",
		"219 g": "২১৯ g",
		"256 GB / 512 GB / 1 TB": "২৫৬ GB / ৫১২ GB / ১ TB",
		"32 MP": "৩২ MP",
		"5,000 mAh": "৫,০০০ এমএএইচ",
		"50 MP + 50 MP + 50 MP + 50 MP": "৫০ MP + ৫০ MP + ৫০ MP + ৫০ MP",
		"5G": "৫G",
		"6.73 inches": "৬.৭৩ ইঞ্চি",
		"Adreno 750": "Adreno ৭৫০",
		"Android 14, HyperOS": "Android ১৪, HyperOS",
		"Black, White, Blue, Titanium Gray": "কালো, সাদা, নীল, Titanium ধূসর",
		"February 2024": "February ২০২৪",
		"Glass front, eco leather back, titanium frame": "গ্লাস সামনে, eco leather back, টাইটানিয়াম ফ্রেম",
		"IP68": "IP৬৮",
		"LTPO AMOLED, 120Hz, Dolby Vision, HDR10+, 3000 nits": "LTPO AMOLED, ১২০Hz, Dolby Vision, HDR১০+, ৩০০০ nits",
		"Qualcomm SM8650-AB Snapdragon 8 Gen 3 (4 nm)": "Qualcomm SM৮৬৫০-AB Snapdragon ৮ Gen ৩ (৪ nm)",
		"Snapdragon 8 Gen 3": "Snapdragon ৮ Gen ৩",
	}
}

// Seed inserts specification records for the product identified by slug 'xiaomi-14-ultra'
func (s *SpecificationSeederMobileXiaomi14Ultra) Seed(db *gorm.DB) error {
	productSlug := "xiaomi-14-ultra"

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

	// Override model-specific values for Xiaomi 14 Ultra
	specs["Display Size"] = "6.73 inches"
	specs["Processor"] = "Snapdragon 8 Gen 3"
	specs["Chipset"] = "Qualcomm SM8650-AB Snapdragon 8 Gen 3 (4 nm)"
	specs["Cpu Type"] = "Octa-core"
	specs["Gpu Type"] = "Adreno 750"
	specs["Ram"] = "12 GB / 16 GB"
	specs["Storage"] = "256 GB / 512 GB / 1 TB"
	specs["Display Type"] = "LTPO AMOLED, 120Hz, Dolby Vision, HDR10+, 3000 nits"
	specs["Resolution"] = "1440 x 3200 pixels"
	specs["Screen Protection"] = "Xiaomi Shield Glass"
	specs["Refresh Rate"] = "120Hz"
	specs["Build Material"] = "Glass front, eco leather back, titanium frame"
	specs["Weight"] = "219 g"
	specs["Dimensions"] = "161.4 x 75.3 x 9.2 mm"
	specs["Water Resistance"] = "IP68"
	specs["Network Technology"] = "5G"
	specs["Rear Camera"] = "50 MP + 50 MP + 50 MP + 50 MP"
	specs["Front Camera"] = "32 MP"
	specs["Battery"] = "5,000 mAh"
	specs["Operating System"] = "Android 14, HyperOS"
	specs["Available Colors"] = "Black, White, Blue, Titanium Gray"
	specs["Announcement Date"] = "February 2024"
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
