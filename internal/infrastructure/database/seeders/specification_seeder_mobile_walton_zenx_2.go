package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileWaltonZenx2 seeds specifications/options for product 'walton-zenx-2'
type SpecificationSeederMobileWaltonZenx2 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileWaltonZenx2 creates a new seeder instance
func NewSpecificationSeederMobileWaltonZenx2() *SpecificationSeederMobileWaltonZenx2 {
	return &SpecificationSeederMobileWaltonZenx2{BaseSeeder: BaseSeeder{name: "Specifications for Walton ZENX 2"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileWaltonZenx2) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080 x 2460 pixels": "১০৮০ x ২৪৬০ pixels",
		"120Hz": "১২০Hz",
		"128 GB": "১২৮ GB",
		"16 MP": "১৬ MP",
		"168.5 x 76.5 x 9.2 mm": "১৬৮.৫ x ৭৬.৫ x ৯.২ মিমি",
		"205 g": "২০৫ g",
		"5,000 mAh": "৫,০০০ এমএএইচ",
		"5G": "৫G",
		"6.8 inches": "৬.৮ ইঞ্চি",
		"64 MP + 8 MP + 2 MP": "৬৪ MP + ৮ MP + ২ MP",
		"8 GB": "৮ GB",
		"Adreno 619": "Adreno ৬১৯",
		"Android 13": "Android ১৩",
		"Corning Gorilla Glass 3": "Corning Gorilla Glass ৩",
		"Glass front, plastic back": "গ্লাস সামনে, প্লাস্টিক পেছনে",
		"IPS LCD, 120Hz": "IPS LCD, ১২০Hz",
		"November 2023": "November ২০২৩",
		"Qualcomm Snapdragon 695 5G (6 nm)": "Qualcomm Snapdragon ৬৯৫ ৫G (৬ nm)",
		"Silver, Black": "রূপালী, কালো",
		"Snapdragon 695": "Snapdragon ৬৯৫",
	}
}

// Seed inserts specification records for the product identified by slug 'walton-zenx-2'
func (s *SpecificationSeederMobileWaltonZenx2) Seed(db *gorm.DB) error {
	productSlug := "walton-zenx-2"

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

	// Override model-specific values for Walton ZENX 2
	specs["Display Size"] = "6.8 inches"
	specs["Processor"] = "Snapdragon 695"
	specs["Chipset"] = "Qualcomm Snapdragon 695 5G (6 nm)"
	specs["Cpu Type"] = "Octa-core"
	specs["Gpu Type"] = "Adreno 619"
	specs["Ram"] = "8 GB"
	specs["Storage"] = "128 GB"
	specs["Display Type"] = "IPS LCD, 120Hz"
	specs["Resolution"] = "1080 x 2460 pixels"
	specs["Screen Protection"] = "Corning Gorilla Glass 3"
	specs["Refresh Rate"] = "120Hz"
	specs["Build Material"] = "Glass front, plastic back"
	specs["Weight"] = "205 g"
	specs["Dimensions"] = "168.5 x 76.5 x 9.2 mm"
	specs["Water Resistance"] = "No"
	specs["Network Technology"] = "5G"
	specs["Rear Camera"] = "64 MP + 8 MP + 2 MP"
	specs["Front Camera"] = "16 MP"
	specs["Battery"] = "5,000 mAh"
	specs["Operating System"] = "Android 13"
	specs["Available Colors"] = "Silver, Black"
	specs["Announcement Date"] = "November 2023"
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
