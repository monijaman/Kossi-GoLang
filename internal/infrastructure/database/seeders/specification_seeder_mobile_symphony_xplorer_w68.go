package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileSymphonyXplorerW68 seeds specifications/options for product 'symphony-xplorer-w68'
type SpecificationSeederMobileSymphonyXplorerW68 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSymphonyXplorerW68 creates a new seeder instance
func NewSpecificationSeederMobileSymphonyXplorerW68() *SpecificationSeederMobileSymphonyXplorerW68 {
	return &SpecificationSeederMobileSymphonyXplorerW68{BaseSeeder: BaseSeeder{name: "Specifications for Symphony Xplorer W68"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSymphonyXplorerW68) getBanglaTranslations() map[string]string {
	return map[string]string{
		"0.3 MP": "০.৩ MP",
		"1,500 mAh": "১,৫০০ এমএএইচ",
		"130 g": "১৩০ g",
		"135 x 67 x 9.5 mm": "১৩৫ x ৬৭ x ৯.৫ মিমি",
		"2013": "২০১৩",
		"3G": "৩G",
		"4 GB": "৪ GB",
		"4.5 inches": "৪.৫ ইঞ্চি",
		"480 x 854 pixels": "৪৮০ x ৮৫৪ pixels",
		"5 MP": "৫ MP",
		"512 MB": "৫১২ MB",
		"60Hz": "৬০Hz",
		"Android 4.2 (Jelly Bean)": "Android ৪.২ (Jelly Bean)",
		"Dual-core 1.2GHz": "Dual-core ১.২GHz",
		"Mali-400": "Mali-৪০০",
		"Mediatek MT6572": "Mediatek MT৬৫৭২",
	}
}

// Seed inserts specification records for the product identified by slug 'symphony-xplorer-w68'
func (s *SpecificationSeederMobileSymphonyXplorerW68) Seed(db *gorm.DB) error {
	productSlug := "symphony-xplorer-w68"

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

	// Override model-specific values for Symphony Xplorer W68
	specs["Display Size"] = "4.5 inches"
	specs["Processor"] = "Dual-core 1.2GHz"
	specs["Chipset"] = "Mediatek MT6572"
	specs["Cpu Type"] = "Dual-core"
	specs["Gpu Type"] = "Mali-400"
	specs["Ram"] = "512 MB"
	specs["Storage"] = "4 GB"
	specs["Display Type"] = "TFT"
	specs["Resolution"] = "480 x 854 pixels"
	specs["Screen Protection"] = "Plastic"
	specs["Refresh Rate"] = "60Hz"
	specs["Build Material"] = "Plastic body"
	specs["Weight"] = "130 g"
	specs["Dimensions"] = "135 x 67 x 9.5 mm"
	specs["Water Resistance"] = "No"
	specs["Network Technology"] = "3G"
	specs["Rear Camera"] = "5 MP"
	specs["Front Camera"] = "0.3 MP"
	specs["Battery"] = "1,500 mAh"
	specs["Operating System"] = "Android 4.2 (Jelly Bean)"
	specs["Available Colors"] = "Black"
	specs["Announcement Date"] = "2013"
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
