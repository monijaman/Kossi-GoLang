package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileWaltonN27 seeds specifications/options for product 'walton-n27'
type SpecificationSeederMobileWaltonN27 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileWaltonN27 creates a new seeder instance
func NewSpecificationSeederMobileWaltonN27() *SpecificationSeederMobileWaltonN27 {
	return &SpecificationSeederMobileWaltonN27{BaseSeeder: BaseSeeder{name: "Specifications for Walton N27"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileWaltonN27) getBanglaTranslations() map[string]string {
	return map[string]string{
		"13 MP + 2 MP": "১৩ মেগাপিক্সেল + ২ মেগাপিক্সেল",
		"164 x 75.8 x 8.9 mm": "১৬৪ x ৭৫.৮ x ৮.৯ মিমি",
		"192 g": "১৯২ g",
		"4 GB": "৪ GB",
		"4G": "৪G",
		"5 MP": "৫ মেগাপিক্সেল",
		"5,000 mAh": "৫,০০০ এমএএইচ",
		"6.5 inches": "৬.৫ ইঞ্চি",
		"60Hz": "৬০Hz",
		"64 GB": "৬৪ GB",
		"720 x 1600 pixels": "৭২০ x ১৬০০ পিক্সেল",
		"Android 13 Go Edition": "অ্যান্ড্রয়েড ১৩ Go Edition",
		"Blue, Black": "নীল, কালো",
		"Helio G36": "হেলিও G৩৬",
		"January 2024": "জানুয়ারি ২০২৪",
		"Mediatek Helio G36 (12 nm)": "মিডিয়াটেক হেলিও G৩৬ (১২ ন্যানোমিটার)",
		"PowerVR GE8320": "পাওয়ারভিআর GE৮৩২০",
	}
}

// Seed inserts specification records for the product identified by slug 'walton-n27'
func (s *SpecificationSeederMobileWaltonN27) Seed(db *gorm.DB) error {
	productSlug := "walton-n27"

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

	// Override model-specific values for Walton N27
	specs["Display Size"] = "6.5 inches"
	specs["Processor"] = "Helio G36"
	specs["Chipset"] = "Mediatek Helio G36 (12 nm)"
	specs["Cpu Type"] = "Octa-core"
	specs["Gpu Type"] = "PowerVR GE8320"
	specs["Ram"] = "4 GB"
	specs["Storage"] = "64 GB"
	specs["Display Type"] = "IPS LCD"
	specs["Resolution"] = "720 x 1600 pixels"
	specs["Screen Protection"] = "Glass front"
	specs["Refresh Rate"] = "60Hz"
	specs["Build Material"] = "Plastic body"
	specs["Weight"] = "192 g"
	specs["Dimensions"] = "164 x 75.8 x 8.9 mm"
	specs["Water Resistance"] = "No"
	specs["Network Technology"] = "4G"
	specs["Rear Camera"] = "13 MP + 2 MP"
	specs["Front Camera"] = "5 MP"
	specs["Battery"] = "5,000 mAh"
	specs["Operating System"] = "Android 13 Go Edition"
	specs["Available Colors"] = "Blue, Black"
	specs["Announcement Date"] = "January 2024"
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
