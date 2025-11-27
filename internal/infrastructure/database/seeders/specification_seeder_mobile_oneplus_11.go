package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileOneplus11 seeds specifications/options for product 'oneplus-11'
type SpecificationSeederMobileOneplus11 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileOneplus11 creates a new seeder instance
func NewSpecificationSeederMobileOneplus11() *SpecificationSeederMobileOneplus11 {
	return &SpecificationSeederMobileOneplus11{BaseSeeder: BaseSeeder{name: "Specifications for OnePlus 11"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileOneplus11) getBanglaTranslations() map[string]string {
	return map[string]string{
		"120Hz": "১২০Hz",
		"128 GB / 256 GB / 512 GB": "১২৮ GB / ২৫৬ GB / ৫১২ GB",
		"1440 x 3216 pixels": "১৪৪০ x ৩২১৬ pixels",
		"16 MP": "১৬ MP",
		"163.1 x 74.1 x 8.5 mm": "১৬৩.১ x ৭৪.১ x ৮.৫ মিমি",
		"205 g": "২০৫ g",
		"5,000 mAh": "৫,০০০ এমএএইচ",
		"50 MP + 32 MP + 48 MP": "৫০ MP + ৩২ MP + ৪৮ MP",
		"5G": "৫G",
		"6.7 inches": "৬.৭ ইঞ্চি",
		"8 GB / 12 GB / 16 GB": "৮ GB / ১২ GB / ১৬ GB",
		"Adreno 740": "Adreno ৭৪০",
		"Android 13, OxygenOS 13": "Android ১৩, OxygenOS ১৩",
		"Glass front/back, aluminum frame": "গ্লাস সামনে/back, অ্যালুমিনিয়াম ফ্রেম",
		"IP64": "IP৬৪",
		"January 2023": "January ২০২৩",
		"LTPO3 Fluid AMOLED, 120Hz, Dolby Vision, HDR10+": "LTPO৩ Fluid AMOLED, ১২০Hz, Dolby Vision, HDR১০+",
		"Qualcomm SM8550-AB Snapdragon 8 Gen 2 (4 nm)": "Qualcomm SM৮৫৫০-AB Snapdragon ৮ Gen ২ (৪ nm)",
		"Snapdragon 8 Gen 2": "Snapdragon ৮ Gen ২",
		"Titan Black, Eternal Green": "Titan কালো, Eternal সবুজ",
	}
}

// Seed inserts specification records for the product identified by slug 'oneplus-11'
func (s *SpecificationSeederMobileOneplus11) Seed(db *gorm.DB) error {
	productSlug := "oneplus-11"

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

	// Override model-specific values for OnePlus 11
	specs["Display Size"] = "6.7 inches"
	specs["Processor"] = "Snapdragon 8 Gen 2"
	specs["Chipset"] = "Qualcomm SM8550-AB Snapdragon 8 Gen 2 (4 nm)"
	specs["Cpu Type"] = "Octa-core"
	specs["Gpu Type"] = "Adreno 740"
	specs["Ram"] = "8 GB / 12 GB / 16 GB"
	specs["Storage"] = "128 GB / 256 GB / 512 GB"
	specs["Display Type"] = "LTPO3 Fluid AMOLED, 120Hz, Dolby Vision, HDR10+"
	specs["Resolution"] = "1440 x 3216 pixels"
	specs["Screen Protection"] = "Corning Gorilla Glass Victus"
	specs["Refresh Rate"] = "120Hz"
	specs["Build Material"] = "Glass front/back, aluminum frame"
	specs["Weight"] = "205 g"
	specs["Dimensions"] = "163.1 x 74.1 x 8.5 mm"
	specs["Water Resistance"] = "IP64"
	specs["Network Technology"] = "5G"
	specs["Rear Camera"] = "50 MP + 32 MP + 48 MP"
	specs["Front Camera"] = "16 MP"
	specs["Battery"] = "5,000 mAh"
	specs["Operating System"] = "Android 13, OxygenOS 13"
	specs["Available Colors"] = "Titan Black, Eternal Green"
	specs["Announcement Date"] = "January 2023"
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
