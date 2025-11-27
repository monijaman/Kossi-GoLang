package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileInfinixNote50Pro5g seeds specifications/options for product 'infinix-note-50-pro-5g'
type SpecificationSeederMobileInfinixNote50Pro5g struct {
	BaseSeeder
}

// NewSpecificationSeederMobileInfinixNote50Pro5g creates a new seeder instance
func NewSpecificationSeederMobileInfinixNote50Pro5g() *SpecificationSeederMobileInfinixNote50Pro5g {
	return &SpecificationSeederMobileInfinixNote50Pro5g{BaseSeeder: BaseSeeder{name: "Specifications for Infinix Note 50 Pro+ 5G"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileInfinixNote50Pro5g) getBanglaTranslations() map[string]string {
	return map[string]string{
		"108 MP + 2 MP + 2 MP": "১০৮ MP + ২ MP + ২ MP",
		"1080 x 2400 pixels": "১০৮০ x ২৪০০ pixels",
		"12 GB": "১২ GB",
		"120Hz": "১২০Hz",
		"162.7 x 75.9 x 8.1 mm": "১৬২.৭ x ৭৫.৯ x ৮.১ মিমি",
		"203 g": "২০৩ g",
		"256 GB": "২৫৬ GB",
		"32 MP": "৩২ MP",
		"5,000 mAh": "৫,০০০ এমএএইচ",
		"5G": "৫G",
		"6.78 inches": "৬.৭৮ ইঞ্চি",
		"AMOLED, 120Hz, 1B colors": "AMOLED, ১২০Hz, ১B colors",
		"Android 14, XOS 14": "Android ১৪, XOS ১৪",
		"Corning Gorilla Glass 5": "Corning Gorilla Glass ৫",
		"Dimensity 8050": "Dimensity ৮০৫০",
		"Glass front, plastic frame, glass back": "গ্লাস সামনে, plastic frame, গ্লাস পেছনে",
		"IP53": "IP৫৩",
		"Magic Black, Variable Gold": "Magic কালো, Variable সোনালী",
		"Mali-G77 MC9": "Mali-G৭৭ MC৯",
		"May 2024": "May ২০২৪",
		"Mediatek Dimensity 8050 (6 nm)": "Mediatek Dimensity ৮০৫০ (৬ nm)",
	}
}

// Seed inserts specification records for the product identified by slug 'infinix-note-50-pro-5g'
func (s *SpecificationSeederMobileInfinixNote50Pro5g) Seed(db *gorm.DB) error {
	productSlug := "infinix-note-50-pro-5g"

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

	// Override model-specific values for Infinix Note 50 Pro+ 5G
	specs["Display Size"] = "6.78 inches"
	specs["Processor"] = "Dimensity 8050"
	specs["Chipset"] = "Mediatek Dimensity 8050 (6 nm)"
	specs["Cpu Type"] = "Octa-core"
	specs["Gpu Type"] = "Mali-G77 MC9"
	specs["Ram"] = "12 GB"
	specs["Storage"] = "256 GB"
	specs["Display Type"] = "AMOLED, 120Hz, 1B colors"
	specs["Resolution"] = "1080 x 2400 pixels"
	specs["Screen Protection"] = "Corning Gorilla Glass 5"
	specs["Refresh Rate"] = "120Hz"
	specs["Build Material"] = "Glass front, plastic frame, glass back"
	specs["Weight"] = "203 g"
	specs["Dimensions"] = "162.7 x 75.9 x 8.1 mm"
	specs["Water Resistance"] = "IP53"
	specs["Network Technology"] = "5G"
	specs["Rear Camera"] = "108 MP + 2 MP + 2 MP"
	specs["Front Camera"] = "32 MP"
	specs["Battery"] = "5,000 mAh"
	specs["Operating System"] = "Android 14, XOS 14"
	specs["Available Colors"] = "Magic Black, Variable Gold"
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
