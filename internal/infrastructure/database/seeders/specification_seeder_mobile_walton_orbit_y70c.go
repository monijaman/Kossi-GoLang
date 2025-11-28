package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileWaltonOrbitY70c seeds specifications/options for product 'walton-orbit-y70c'
type SpecificationSeederMobileWaltonOrbitY70c struct {
	BaseSeeder
}

// NewSpecificationSeederMobileWaltonOrbitY70c creates a new seeder instance
func NewSpecificationSeederMobileWaltonOrbitY70c() *SpecificationSeederMobileWaltonOrbitY70c {
	return &SpecificationSeederMobileWaltonOrbitY70c{BaseSeeder: BaseSeeder{name: "Specifications for Walton ORBIT Y70C"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileWaltonOrbitY70c) getBanglaTranslations() map[string]string {
	return map[string]string{
		"164.5 x 76 x 9 mm": "১৬৪.৫ x ৭৬ x ৯ মিমি",
		"188 g": "১৮৮ g",
		"3 GB": "৩ GB",
		"32 GB": "৩২ GB",
		"4,000 mAh": "৪,০০০ এমএএইচ",
		"4G": "৪G",
		"5 MP": "৫ মেগাপিক্সেল",
		"6.5 inches": "৬.৫ ইঞ্চি",
		"60Hz": "৬০Hz",
		"720 x 1600 pixels": "৭২০ x ১৬০০ পিক্সেল",
		"8 MP + 0.3 MP": "৮ মেগাপিক্সেল + ০.৩ মেগাপিক্সেল",
		"Android 12 Go Edition": "অ্যান্ড্রয়েড ১২ Go Edition",
		"Blue, Green": "নীল, সবুজ",
		"Helio A22": "হেলিও A২২",
		"June 2023": "জুন ২০২৩",
		"Mediatek Helio A22 (12 nm)": "মিডিয়াটেক হেলিও A২২ (১২ ন্যানোমিটার)",
		"PowerVR GE8320": "পাওয়ারভিআর GE৮৩২০",
	}
}

// Seed inserts specification records for the product identified by slug 'walton-orbit-y70c'
func (s *SpecificationSeederMobileWaltonOrbitY70c) Seed(db *gorm.DB) error {
	productSlug := "walton-orbit-y70c"

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

	// Override model-specific values for Walton ORBIT Y70C
	specs["Display Size"] = "6.5 inches"
	specs["Processor"] = "Helio A22"
	specs["Chipset"] = "Mediatek Helio A22 (12 nm)"
	specs["Cpu Type"] = "Quad-core"
	specs["Gpu Type"] = "PowerVR GE8320"
	specs["Ram"] = "3 GB"
	specs["Storage"] = "32 GB"
	specs["Display Type"] = "IPS LCD"
	specs["Resolution"] = "720 x 1600 pixels"
	specs["Screen Protection"] = "Glass front"
	specs["Refresh Rate"] = "60Hz"
	specs["Build Material"] = "Plastic body"
	specs["Weight"] = "188 g"
	specs["Dimensions"] = "164.5 x 76 x 9 mm"
	specs["Water Resistance"] = "No"
	specs["Network Technology"] = "4G"
	specs["Rear Camera"] = "8 MP + 0.3 MP"
	specs["Front Camera"] = "5 MP"
	specs["Battery"] = "4,000 mAh"
	specs["Operating System"] = "Android 12 Go Edition"
	specs["Available Colors"] = "Blue, Green"
	specs["Announcement Date"] = "June 2023"
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
