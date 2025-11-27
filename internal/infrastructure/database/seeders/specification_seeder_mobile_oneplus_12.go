package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileOneplus12 seeds specifications/options for product 'oneplus-12'
type SpecificationSeederMobileOneplus12 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileOneplus12 creates a new seeder instance
func NewSpecificationSeederMobileOneplus12() *SpecificationSeederMobileOneplus12 {
	return &SpecificationSeederMobileOneplus12{BaseSeeder: BaseSeeder{name: "Specifications for OnePlus 12"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileOneplus12) getBanglaTranslations() map[string]string {
	return map[string]string{
		"12 GB / 16 GB / 24 GB": "১২ GB / ১৬ GB / ২৪ GB",
		"120Hz": "১২০Hz",
		"1440 x 3168 pixels": "১৪৪০ x ৩১৬৮ pixels",
		"164.3 x 75.8 x 9.2 mm": "১৬৪.৩ x ৭৫.৮ x ৯.২ মিমি",
		"220 g": "২২০ g",
		"256 GB / 512 GB / 1 TB": "২৫৬ GB / ৫১২ GB / ১ TB",
		"32 MP": "৩২ MP",
		"5,400 mAh": "৫,৪০০ এমএএইচ",
		"50 MP + 64 MP + 48 MP": "৫০ MP + ৬৪ MP + ৪৮ MP",
		"5G": "৫G",
		"6.82 inches": "৬.৮২ ইঞ্চি",
		"Adreno 750": "Adreno ৭৫০",
		"Android 14, OxygenOS 14": "Android ১৪, OxygenOS ১৪",
		"Corning Gorilla Glass Victus 2": "Corning Gorilla Glass Victus ২",
		"December 2023": "December ২০২৩",
		"Flowy Emerald, Silky Black, Silver": "Flowy Emerald, Silky কালো, রূপালী",
		"Glass front/back, aluminum frame": "গ্লাস সামনে/back, অ্যালুমিনিয়াম ফ্রেম",
		"IP65": "IP৬৫",
		"LTPO AMOLED, 120Hz, Dolby Vision, HDR10+, 4500 nits": "LTPO AMOLED, ১২০Hz, Dolby Vision, HDR১০+, ৪৫০০ nits",
		"Qualcomm SM8650-AB Snapdragon 8 Gen 3 (4 nm)": "Qualcomm SM৮৬৫০-AB Snapdragon ৮ Gen ৩ (৪ nm)",
		"Snapdragon 8 Gen 3": "Snapdragon ৮ Gen ৩",
	}
}

// Seed inserts specification records for the product identified by slug 'oneplus-12'
func (s *SpecificationSeederMobileOneplus12) Seed(db *gorm.DB) error {
	productSlug := "oneplus-12"

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

	// Override model-specific values for OnePlus 12
	specs["Display Size"] = "6.82 inches"
	specs["Processor"] = "Snapdragon 8 Gen 3"
	specs["Chipset"] = "Qualcomm SM8650-AB Snapdragon 8 Gen 3 (4 nm)"
	specs["Cpu Type"] = "Octa-core"
	specs["Gpu Type"] = "Adreno 750"
	specs["Ram"] = "12 GB / 16 GB / 24 GB"
	specs["Storage"] = "256 GB / 512 GB / 1 TB"
	specs["Display Type"] = "LTPO AMOLED, 120Hz, Dolby Vision, HDR10+, 4500 nits"
	specs["Resolution"] = "1440 x 3168 pixels"
	specs["Screen Protection"] = "Corning Gorilla Glass Victus 2"
	specs["Refresh Rate"] = "120Hz"
	specs["Build Material"] = "Glass front/back, aluminum frame"
	specs["Weight"] = "220 g"
	specs["Dimensions"] = "164.3 x 75.8 x 9.2 mm"
	specs["Water Resistance"] = "IP65"
	specs["Network Technology"] = "5G"
	specs["Rear Camera"] = "50 MP + 64 MP + 48 MP"
	specs["Front Camera"] = "32 MP"
	specs["Battery"] = "5,400 mAh"
	specs["Operating System"] = "Android 14, OxygenOS 14"
	specs["Available Colors"] = "Flowy Emerald, Silky Black, Silver"
	specs["Announcement Date"] = "December 2023"
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
