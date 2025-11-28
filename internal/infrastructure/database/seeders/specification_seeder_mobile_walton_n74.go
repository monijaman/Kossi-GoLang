package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileWaltonN74 seeds specifications/options for product 'walton-n74'
type SpecificationSeederMobileWaltonN74 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileWaltonN74 creates a new seeder instance
func NewSpecificationSeederMobileWaltonN74() *SpecificationSeederMobileWaltonN74 {
	return &SpecificationSeederMobileWaltonN74{BaseSeeder: BaseSeeder{name: "Specifications for Walton N74"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileWaltonN74) getBanglaTranslations() map[string]string {
	return map[string]string{
		"128 GB": "১২৮ GB",
		"164.2 x 76 x 8.5 mm": "১৬৪.২ x ৭৬ x ৮.৫ মিমি",
		"190 g": "১৯০ g",
		"4G": "৪G",
		"5,000 mAh": "৫,০০০ এমএএইচ",
		"50 MP + 2 MP": "৫০ মেগাপিক্সেল + ২ মেগাপিক্সেল",
		"6 GB": "৬ GB",
		"6.6 inches": "৬.৬ ইঞ্চি",
		"720 x 1612 pixels": "৭২০ x ১৬১২ পিক্সেল",
		"8 MP": "৮ মেগাপিক্সেল",
		"90Hz": "৯০Hz",
		"Android 13": "অ্যান্ড্রয়েড ১৩",
		"Green, Black": "সবুজ, কালো",
		"IPS LCD, 90Hz": "আইপিএস এলসিডি, ৯০Hz",
		"July 2023": "জুলাই ২০২৩",
		"Mali-G57 MP1": "মালি-G৫৭ মেগাপিক্সেল১",
		"Unisoc T606": "ইউনিসক T৬০৬",
		"Unisoc T606 (12 nm)": "ইউনিসক T৬০৬ (১২ ন্যানোমিটার)",
	}
}

// Seed inserts specification records for the product identified by slug 'walton-n74'
func (s *SpecificationSeederMobileWaltonN74) Seed(db *gorm.DB) error {
	productSlug := "walton-n74"

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

	// Override model-specific values for Walton N74
	specs["Display Size"] = "6.6 inches"
	specs["Processor"] = "Unisoc T606"
	specs["Chipset"] = "Unisoc T606 (12 nm)"
	specs["Cpu Type"] = "Octa-core"
	specs["Gpu Type"] = "Mali-G57 MP1"
	specs["Ram"] = "6 GB"
	specs["Storage"] = "128 GB"
	specs["Display Type"] = "IPS LCD, 90Hz"
	specs["Resolution"] = "720 x 1612 pixels"
	specs["Screen Protection"] = "Glass front"
	specs["Refresh Rate"] = "90Hz"
	specs["Build Material"] = "Plastic body"
	specs["Weight"] = "190 g"
	specs["Dimensions"] = "164.2 x 76 x 8.5 mm"
	specs["Water Resistance"] = "No"
	specs["Network Technology"] = "4G"
	specs["Rear Camera"] = "50 MP + 2 MP"
	specs["Front Camera"] = "8 MP"
	specs["Battery"] = "5,000 mAh"
	specs["Operating System"] = "Android 13"
	specs["Available Colors"] = "Green, Black"
	specs["Announcement Date"] = "July 2023"
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
