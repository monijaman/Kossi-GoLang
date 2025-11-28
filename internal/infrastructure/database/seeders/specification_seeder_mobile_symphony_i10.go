package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileSymphonyI10 seeds specifications/options for product 'symphony-i10'
type SpecificationSeederMobileSymphonyI10 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSymphonyI10 creates a new seeder instance
func NewSpecificationSeederMobileSymphonyI10() *SpecificationSeederMobileSymphonyI10 {
	return &SpecificationSeederMobileSymphonyI10{BaseSeeder: BaseSeeder{name: "Specifications for Symphony i10"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSymphonyI10) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1 GB": "১ GB",
		"140 g": "১৪০ g",
		"144 x 72 x 9 mm": "১৪৪ x ৭২ x ৯ মিমি",
		"16 GB": "১৬ GB",
		"2,500 mAh": "২,৫০০ এমএএইচ",
		"2016": "২০১৬",
		"3G": "৩G",
		"5.0 inches": "৫.০ ইঞ্চি",
		"60Hz": "৬০Hz",
		"720 x 1280 pixels": "৭২০ x ১২৮০ পিক্সেল",
		"8 MP": "৮ মেগাপিক্সেল",
		"Android 6.0 (Marshmallow)": "অ্যান্ড্রয়েড ৬.০ (Marshmallow)",
		"Gold, Grey": "সোনালী, ধূসর",
		"Mali-400MP2": "মালি-৪০০MP২",
		"Mediatek MT6580": "মিডিয়াটেক MT৬৫৮০",
		"Quad-core 1.3GHz": "Quad-কোর ১.৩GHz",
	}
}

// Seed inserts specification records for the product identified by slug 'symphony-i10'
func (s *SpecificationSeederMobileSymphonyI10) Seed(db *gorm.DB) error {
	productSlug := "symphony-i10"

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

	// Override model-specific values for Symphony i10
	specs["Display Size"] = "5.0 inches"
	specs["Processor"] = "Quad-core 1.3GHz"
	specs["Chipset"] = "Mediatek MT6580"
	specs["Cpu Type"] = "Quad-core"
	specs["Gpu Type"] = "Mali-400MP2"
	specs["Ram"] = "1 GB"
	specs["Storage"] = "16 GB"
	specs["Display Type"] = "IPS LCD"
	specs["Resolution"] = "720 x 1280 pixels"
	specs["Screen Protection"] = "Glass front"
	specs["Refresh Rate"] = "60Hz"
	specs["Build Material"] = "Plastic body"
	specs["Weight"] = "140 g"
	specs["Dimensions"] = "144 x 72 x 9 mm"
	specs["Water Resistance"] = "No"
	specs["Network Technology"] = "3G"
	specs["Rear Camera"] = "8 MP"
	specs["Front Camera"] = "8 MP"
	specs["Battery"] = "2,500 mAh"
	specs["Operating System"] = "Android 6.0 (Marshmallow)"
	specs["Available Colors"] = "Gold, Grey"
	specs["Announcement Date"] = "2016"
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
