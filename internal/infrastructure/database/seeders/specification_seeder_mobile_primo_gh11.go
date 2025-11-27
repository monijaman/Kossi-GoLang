package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobilePrimoGh11 seeds specifications/options for product 'primo-gh11'
type SpecificationSeederMobilePrimoGh11 struct {
	BaseSeeder
}

// NewSpecificationSeederMobilePrimoGh11 creates a new seeder instance
func NewSpecificationSeederMobilePrimoGh11() *SpecificationSeederMobilePrimoGh11 {
	return &SpecificationSeederMobilePrimoGh11{BaseSeeder: BaseSeeder{name: "Specifications for Primo GH11"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobilePrimoGh11) getBanglaTranslations() map[string]string {
	return map[string]string{
		"13 MP + 2 MP": "১৩ MP + ২ MP",
		"165 x 76 x 9 mm": "১৬৫ x ৭৬ x ৯ মিমি",
		"190 g": "১৯০ g",
		"2 GB": "২ GB",
		"32 GB": "৩২ GB",
		"4,200 mAh": "৪,২০০ এমএএইচ",
		"4G": "৪G",
		"5 MP": "৫ MP",
		"6.52 inches": "৬.৫২ ইঞ্চি",
		"60Hz": "৬০Hz",
		"720 x 1600 pixels": "৭২০ x ১৬০০ pixels",
		"Android 11 Go Edition": "Android ১১ Go Edition",
		"Blue, Black": "নীল, কালো",
		"Helio A25": "Helio A২৫",
		"Mediatek Helio A25 (12 nm)": "Mediatek Helio A২৫ (১২ nm)",
		"November 2021": "November ২০২১",
		"PowerVR GE8320": "PowerVR GE৮৩২০",
	}
}

// Seed inserts specification records for the product identified by slug 'primo-gh11'
func (s *SpecificationSeederMobilePrimoGh11) Seed(db *gorm.DB) error {
	productSlug := "primo-gh11"

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

	// Override model-specific values for Primo GH11
	specs["Display Size"] = "6.52 inches"
	specs["Processor"] = "Helio A25"
	specs["Chipset"] = "Mediatek Helio A25 (12 nm)"
	specs["Cpu Type"] = "Octa-core"
	specs["Gpu Type"] = "PowerVR GE8320"
	specs["Ram"] = "2 GB"
	specs["Storage"] = "32 GB"
	specs["Display Type"] = "IPS LCD"
	specs["Resolution"] = "720 x 1600 pixels"
	specs["Screen Protection"] = "Glass front"
	specs["Refresh Rate"] = "60Hz"
	specs["Build Material"] = "Plastic body"
	specs["Weight"] = "190 g"
	specs["Dimensions"] = "165 x 76 x 9 mm"
	specs["Water Resistance"] = "No"
	specs["Network Technology"] = "4G"
	specs["Rear Camera"] = "13 MP + 2 MP"
	specs["Front Camera"] = "5 MP"
	specs["Battery"] = "4,200 mAh"
	specs["Operating System"] = "Android 11 Go Edition"
	specs["Available Colors"] = "Blue, Black"
	specs["Announcement Date"] = "November 2021"
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
