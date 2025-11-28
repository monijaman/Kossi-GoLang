package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileSymphonyXplorerW92 seeds specifications/options for product 'symphony-xplorer-w92'
type SpecificationSeederMobileSymphonyXplorerW92 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSymphonyXplorerW92 creates a new seeder instance
func NewSpecificationSeederMobileSymphonyXplorerW92() *SpecificationSeederMobileSymphonyXplorerW92 {
	return &SpecificationSeederMobileSymphonyXplorerW92{BaseSeeder: BaseSeeder{name: "Specifications for Symphony Xplorer W92"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSymphonyXplorerW92) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1 GB": "১ GB",
		"145 g": "১৪৫ g",
		"145 x 73 x 9 mm": "১৪৫ x ৭৩ x ৯ মিমি",
		"2 MP": "২ মেগাপিক্সেল",
		"2,000 mAh": "২,০০০ এমএএইচ",
		"2013": "২০১৩",
		"3G": "৩G",
		"4 GB": "৪ GB",
		"480 x 854 pixels": "৪৮০ x ৮৫৪ পিক্সেল",
		"5 MP": "৫ মেগাপিক্সেল",
		"5.0 inches": "৫.০ ইঞ্চি",
		"60Hz": "৬০Hz",
		"Android 4.2 (Jelly Bean)": "অ্যান্ড্রয়েড ৪.২ (Jelly Bean)",
		"Black, White": "কালো, সাদা",
		"Mediatek MT6589": "মিডিয়াটেক MT৬৫৮৯",
		"PowerVR SGX544": "পাওয়ারভিআর SGX৫৪৪",
		"Quad-core 1.2GHz": "Quad-কোর ১.২GHz",
	}
}

// Seed inserts specification records for the product identified by slug 'symphony-xplorer-w92'
func (s *SpecificationSeederMobileSymphonyXplorerW92) Seed(db *gorm.DB) error {
	productSlug := "symphony-xplorer-w92"

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

	// Override model-specific values for Symphony Xplorer W92
	specs["Display Size"] = "5.0 inches"
	specs["Processor"] = "Quad-core 1.2GHz"
	specs["Chipset"] = "Mediatek MT6589"
	specs["Cpu Type"] = "Quad-core"
	specs["Gpu Type"] = "PowerVR SGX544"
	specs["Ram"] = "1 GB"
	specs["Storage"] = "4 GB"
	specs["Display Type"] = "IPS LCD"
	specs["Resolution"] = "480 x 854 pixels"
	specs["Screen Protection"] = "Glass front"
	specs["Refresh Rate"] = "60Hz"
	specs["Build Material"] = "Plastic body"
	specs["Weight"] = "145 g"
	specs["Dimensions"] = "145 x 73 x 9 mm"
	specs["Water Resistance"] = "No"
	specs["Network Technology"] = "3G"
	specs["Rear Camera"] = "5 MP"
	specs["Front Camera"] = "2 MP"
	specs["Battery"] = "2,000 mAh"
	specs["Operating System"] = "Android 4.2 (Jelly Bean)"
	specs["Available Colors"] = "Black, White"
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
