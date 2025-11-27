package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileInfinixNote50 seeds specifications/options for product 'infinix-note-50'
type SpecificationSeederMobileInfinixNote50 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileInfinixNote50 creates a new seeder instance
func NewSpecificationSeederMobileInfinixNote50() *SpecificationSeederMobileInfinixNote50 {
	return &SpecificationSeederMobileInfinixNote50{BaseSeeder: BaseSeeder{name: "Specifications for Infinix Note 50"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileInfinixNote50) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080 x 2460 pixels": "১০৮০ x ২৪৬০ pixels",
		"120Hz": "১২০Hz",
		"128 GB / 256 GB": "১২৮ GB / ২৫৬ GB",
		"16 MP": "১৬ MP",
		"168 x 76.6 x 8.5 mm": "১৬৮ x ৭৬.৬ x ৮.৫ মিমি",
		"198 g": "১৯৮ g",
		"4G": "৪G",
		"5,000 mAh": "৫,০০০ এমএএইচ",
		"6.78 inches": "৬.৭৮ ইঞ্চি",
		"64 MP + 2 MP": "৬৪ MP + ২ MP",
		"8 GB": "৮ GB",
		"Android 14": "Android ১৪",
		"Black, Blue": "কালো, নীল",
		"Helio G99": "Helio G৯৯",
		"IPS LCD, 120Hz": "IPS LCD, ১২০Hz",
		"Mali-G57 MC2": "Mali-G৫৭ MC২",
		"May 2024": "May ২০২৪",
		"Mediatek Helio G99 (6 nm)": "Mediatek Helio G৯৯ (৬ nm)",
	}
}

// Seed inserts specification records for the product identified by slug 'infinix-note-50'
func (s *SpecificationSeederMobileInfinixNote50) Seed(db *gorm.DB) error {
	productSlug := "infinix-note-50"

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

	// Override model-specific values for Infinix Note 50
	specs["Display Size"] = "6.78 inches"
	specs["Processor"] = "Helio G99"
	specs["Chipset"] = "Mediatek Helio G99 (6 nm)"
	specs["Cpu Type"] = "Octa-core"
	specs["Gpu Type"] = "Mali-G57 MC2"
	specs["Ram"] = "8 GB"
	specs["Storage"] = "128 GB / 256 GB"
	specs["Display Type"] = "IPS LCD, 120Hz"
	specs["Resolution"] = "1080 x 2460 pixels"
	specs["Screen Protection"] = "Glass front"
	specs["Refresh Rate"] = "120Hz"
	specs["Build Material"] = "Plastic body"
	specs["Weight"] = "198 g"
	specs["Dimensions"] = "168 x 76.6 x 8.5 mm"
	specs["Water Resistance"] = "No"
	specs["Network Technology"] = "4G"
	specs["Rear Camera"] = "64 MP + 2 MP"
	specs["Front Camera"] = "16 MP"
	specs["Battery"] = "5,000 mAh"
	specs["Operating System"] = "Android 14"
	specs["Available Colors"] = "Black, Blue"
	specs["Announcement Date"] = "May 2024"
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
