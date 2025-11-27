package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileInfinixNote50Pro seeds specifications/options for product 'infinix-note-50-pro'
type SpecificationSeederMobileInfinixNote50Pro struct {
	BaseSeeder
}

// NewSpecificationSeederMobileInfinixNote50Pro creates a new seeder instance
func NewSpecificationSeederMobileInfinixNote50Pro() *SpecificationSeederMobileInfinixNote50Pro {
	return &SpecificationSeederMobileInfinixNote50Pro{BaseSeeder: BaseSeeder{name: "Specifications for Infinix Note 50 Pro"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileInfinixNote50Pro) getBanglaTranslations() map[string]string {
	return map[string]string{
		"108 MP + 2 MP": "১০৮ MP + ২ MP",
		"1080 x 2460 pixels": "১০৮০ x ২৪৬০ pixels",
		"120Hz": "১২০Hz",
		"16 MP": "১৬ MP",
		"168.6 x 76.6 x 8.6 mm": "১৬৮.৬ x ৭৬.৬ x ৮.৬ মিমি",
		"203 g": "২০৩ g",
		"256 GB": "২৫৬ GB",
		"4G": "৪G",
		"5,000 mAh": "৫,০০০ এমএএইচ",
		"6.78 inches": "৬.৭৮ ইঞ্চি",
		"8 GB": "৮ GB",
		"AMOLED, 120Hz": "AMOLED, ১২০Hz",
		"Android 14, XOS 14": "Android ১৪, XOS ১৪",
		"Black, Gold": "কালো, সোনালী",
		"Glass front, plastic frame, plastic back": "গ্লাস সামনে, plastic frame, প্লাস্টিক পেছনে",
		"Helio G99": "Helio G৯৯",
		"IP53": "IP৫৩",
		"Mali-G57 MC2": "Mali-G৫৭ MC২",
		"May 2024": "May ২০২৪",
		"Mediatek Helio G99 (6 nm)": "Mediatek Helio G৯৯ (৬ nm)",
	}
}

// Seed inserts specification records for the product identified by slug 'infinix-note-50-pro'
func (s *SpecificationSeederMobileInfinixNote50Pro) Seed(db *gorm.DB) error {
	productSlug := "infinix-note-50-pro"

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

	// Override model-specific values for Infinix Note 50 Pro
	specs["Display Size"] = "6.78 inches"
	specs["Processor"] = "Helio G99"
	specs["Chipset"] = "Mediatek Helio G99 (6 nm)"
	specs["Cpu Type"] = "Octa-core"
	specs["Gpu Type"] = "Mali-G57 MC2"
	specs["Ram"] = "8 GB"
	specs["Storage"] = "256 GB"
	specs["Display Type"] = "AMOLED, 120Hz"
	specs["Resolution"] = "1080 x 2460 pixels"
	specs["Screen Protection"] = "Corning Gorilla Glass"
	specs["Refresh Rate"] = "120Hz"
	specs["Build Material"] = "Glass front, plastic frame, plastic back"
	specs["Weight"] = "203 g"
	specs["Dimensions"] = "168.6 x 76.6 x 8.6 mm"
	specs["Water Resistance"] = "IP53"
	specs["Network Technology"] = "4G"
	specs["Rear Camera"] = "108 MP + 2 MP"
	specs["Front Camera"] = "16 MP"
	specs["Battery"] = "5,000 mAh"
	specs["Operating System"] = "Android 14, XOS 14"
	specs["Available Colors"] = "Black, Gold"
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
