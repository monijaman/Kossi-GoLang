package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobilePrimoS8 seeds specifications/options for product 'primo-s8'
type SpecificationSeederMobilePrimoS8 struct {
	BaseSeeder
}

// NewSpecificationSeederMobilePrimoS8 creates a new seeder instance
func NewSpecificationSeederMobilePrimoS8() *SpecificationSeederMobilePrimoS8 {
	return &SpecificationSeederMobilePrimoS8{BaseSeeder: BaseSeeder{name: "Specifications for Primo S8"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobilePrimoS8) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080 x 2460 pixels": "১০৮০ x ২৪৬০ pixels",
		"128 GB": "১২৮ GB",
		"169 x 76.5 x 8.9 mm": "১৬৯ x ৭৬.৫ x ৮.৯ মিমি",
		"205 g": "২০৫ g",
		"48 MP + 5 MP + 2 MP + 2 MP": "৪৮ MP + ৫ MP + ২ MP + ২ MP",
		"4G": "৪G",
		"5,000 mAh": "৫,০০০ এমএএইচ",
		"6 GB": "৬ GB",
		"6.78 inches": "৬.৭৮ ইঞ্চি",
		"8 MP": "৮ MP",
		"90Hz": "৯০Hz",
		"Android 11": "Android ১১",
		"Black, Blue": "কালো, নীল",
		"December 2021": "December ২০২১",
		"Glass front, plastic frame, plastic back": "গ্লাস সামনে, plastic frame, প্লাস্টিক পেছনে",
		"Helio G88": "Helio G৮৮",
		"IPS LCD, 90Hz": "IPS LCD, ৯০Hz",
		"Mali-G52 MC2": "Mali-G৫২ MC২",
		"Mediatek Helio G88 (12 nm)": "Mediatek Helio G৮৮ (১২ nm)",
	}
}

// Seed inserts specification records for the product identified by slug 'primo-s8'
func (s *SpecificationSeederMobilePrimoS8) Seed(db *gorm.DB) error {
	productSlug := "primo-s8"

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

	// Override model-specific values for Primo S8
	specs["Display Size"] = "6.78 inches"
	specs["Processor"] = "Helio G88"
	specs["Chipset"] = "Mediatek Helio G88 (12 nm)"
	specs["Cpu Type"] = "Octa-core"
	specs["Gpu Type"] = "Mali-G52 MC2"
	specs["Ram"] = "6 GB"
	specs["Storage"] = "128 GB"
	specs["Display Type"] = "IPS LCD, 90Hz"
	specs["Resolution"] = "1080 x 2460 pixels"
	specs["Screen Protection"] = "Glass front"
	specs["Refresh Rate"] = "90Hz"
	specs["Build Material"] = "Glass front, plastic frame, plastic back"
	specs["Weight"] = "205 g"
	specs["Dimensions"] = "169 x 76.5 x 8.9 mm"
	specs["Water Resistance"] = "No"
	specs["Network Technology"] = "4G"
	specs["Rear Camera"] = "48 MP + 5 MP + 2 MP + 2 MP"
	specs["Front Camera"] = "8 MP"
	specs["Battery"] = "5,000 mAh"
	specs["Operating System"] = "Android 11"
	specs["Available Colors"] = "Black, Blue"
	specs["Announcement Date"] = "December 2021"
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
