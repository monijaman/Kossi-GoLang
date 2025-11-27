package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileInfinixNote40Pro seeds specifications/options for product 'infinix-note-40-pro'
type SpecificationSeederMobileInfinixNote40Pro struct {
	BaseSeeder
}

// NewSpecificationSeederMobileInfinixNote40Pro creates a new seeder instance
func NewSpecificationSeederMobileInfinixNote40Pro() *SpecificationSeederMobileInfinixNote40Pro {
	return &SpecificationSeederMobileInfinixNote40Pro{BaseSeeder: BaseSeeder{name: "Specifications for Infinix Note 40 Pro"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileInfinixNote40Pro) getBanglaTranslations() map[string]string {
	return map[string]string{
		"108 MP + 2 MP + 2 MP": "১০৮ MP + ২ MP + ২ MP",
		"1080 x 2436 pixels": "১০৮০ x ২৪৩৬ pixels",
		"120Hz": "১২০Hz",
		"164.4 x 74.6 x 7.8 mm": "১৬৪.৪ x ৭৪.৬ x ৭.৮ মিমি",
		"190 g": "১৯০ g",
		"256 GB": "২৫৬ GB",
		"32 MP": "৩২ MP",
		"4G": "৪G",
		"5,000 mAh": "৫,০০০ এমএএইচ",
		"6.78 inches": "৬.৭৮ ইঞ্চি",
		"8 GB / 12 GB": "৮ GB / ১২ GB",
		"AMOLED, 120Hz, 1300 nits": "AMOLED, ১২০Hz, ১৩০০ nits",
		"Android 14, XOS 14": "Android ১৪, XOS ১৪",
		"Glass front, plastic frame, plastic back": "গ্লাস সামনে, plastic frame, প্লাস্টিক পেছনে",
		"Helio G99 Ultimate": "Helio G৯৯ Ultimate",
		"IP54": "IP৫৪",
		"Mali-G57 MC2": "Mali-G৫৭ MC২",
		"March 2024": "March ২০২৪",
		"Mediatek Helio G99 Ultimate": "Mediatek Helio G৯৯ Ultimate",
		"Vintage Green, Titan Gold": "Vintage সবুজ, Titan সোনালী",
	}
}

// Seed inserts specification records for the product identified by slug 'infinix-note-40-pro'
func (s *SpecificationSeederMobileInfinixNote40Pro) Seed(db *gorm.DB) error {
	productSlug := "infinix-note-40-pro"

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

	// Override model-specific values for Infinix Note 40 Pro
	specs["Display Size"] = "6.78 inches"
	specs["Processor"] = "Helio G99 Ultimate"
	specs["Chipset"] = "Mediatek Helio G99 Ultimate"
	specs["Cpu Type"] = "Octa-core"
	specs["Gpu Type"] = "Mali-G57 MC2"
	specs["Ram"] = "8 GB / 12 GB"
	specs["Storage"] = "256 GB"
	specs["Display Type"] = "AMOLED, 120Hz, 1300 nits"
	specs["Resolution"] = "1080 x 2436 pixels"
	specs["Screen Protection"] = "Corning Gorilla Glass"
	specs["Refresh Rate"] = "120Hz"
	specs["Build Material"] = "Glass front, plastic frame, plastic back"
	specs["Weight"] = "190 g"
	specs["Dimensions"] = "164.4 x 74.6 x 7.8 mm"
	specs["Water Resistance"] = "IP54"
	specs["Network Technology"] = "4G"
	specs["Rear Camera"] = "108 MP + 2 MP + 2 MP"
	specs["Front Camera"] = "32 MP"
	specs["Battery"] = "5,000 mAh"
	specs["Operating System"] = "Android 14, XOS 14"
	specs["Available Colors"] = "Vintage Green, Titan Gold"
	specs["Announcement Date"] = "March 2024"
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
