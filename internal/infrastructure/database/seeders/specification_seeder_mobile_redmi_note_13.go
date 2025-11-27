package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileRedmiNote13 seeds specifications/options for product 'redmi-note-13'
type SpecificationSeederMobileRedmiNote13 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileRedmiNote13 creates a new seeder instance
func NewSpecificationSeederMobileRedmiNote13() *SpecificationSeederMobileRedmiNote13 {
	return &SpecificationSeederMobileRedmiNote13{BaseSeeder: BaseSeeder{name: "Specifications for Redmi Note 13"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileRedmiNote13) getBanglaTranslations() map[string]string {
	return map[string]string{
		"108 MP + 8 MP + 2 MP": "১০৮ MP + ৮ MP + ২ MP",
		"1080 x 2400 pixels": "১০৮০ x ২৪০০ pixels",
		"120Hz": "১২০Hz",
		"128 GB / 256 GB": "১২৮ GB / ২৫৬ GB",
		"16 MP": "১৬ MP",
		"161.1 x 75 x 7.6 mm": "১৬১.১ x ৭৫ x ৭.৬ মিমি",
		"173.5 g": "১৭৩.৫ g",
		"5,000 mAh": "৫,০০০ এমএএইচ",
		"5G": "৫G",
		"6 GB / 8 GB / 12 GB": "৬ GB / ৮ GB / ১২ GB",
		"6.67 inches": "৬.৬৭ ইঞ্চি",
		"AMOLED, 120Hz, 1000 nits": "AMOLED, ১২০Hz, ১০০০ nits",
		"Android 13, upgradable": "Android ১৩, upgradable",
		"Corning Gorilla Glass 5": "Corning Gorilla Glass ৫",
		"Dimensity 6080": "Dimensity ৬০৮০",
		"Glass front, plastic frame, plastic back": "গ্লাস সামনে, plastic frame, প্লাস্টিক পেছনে",
		"IP54": "IP৫৪",
		"January 2024": "January ২০২৪",
		"Mali-G57 MC2": "Mali-G৫৭ MC২",
		"Mediatek Dimensity 6080 (6 nm)": "Mediatek Dimensity ৬০৮০ (৬ nm)",
		"Stealth Black, Arctic White, Prism Gold": "Stealth কালো, Arctic সাদা, Prism সোনালী",
	}
}

// Seed inserts specification records for the product identified by slug 'redmi-note-13'
func (s *SpecificationSeederMobileRedmiNote13) Seed(db *gorm.DB) error {
	productSlug := "redmi-note-13"

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

	// Override model-specific values for Redmi Note 13
	specs["Display Size"] = "6.67 inches"
	specs["Processor"] = "Dimensity 6080"
	specs["Chipset"] = "Mediatek Dimensity 6080 (6 nm)"
	specs["Cpu Type"] = "Octa-core"
	specs["Gpu Type"] = "Mali-G57 MC2"
	specs["Ram"] = "6 GB / 8 GB / 12 GB"
	specs["Storage"] = "128 GB / 256 GB"
	specs["Display Type"] = "AMOLED, 120Hz, 1000 nits"
	specs["Resolution"] = "1080 x 2400 pixels"
	specs["Screen Protection"] = "Corning Gorilla Glass 5"
	specs["Refresh Rate"] = "120Hz"
	specs["Build Material"] = "Glass front, plastic frame, plastic back"
	specs["Weight"] = "173.5 g"
	specs["Dimensions"] = "161.1 x 75 x 7.6 mm"
	specs["Water Resistance"] = "IP54"
	specs["Network Technology"] = "5G"
	specs["Rear Camera"] = "108 MP + 8 MP + 2 MP"
	specs["Front Camera"] = "16 MP"
	specs["Battery"] = "5,000 mAh"
	specs["Operating System"] = "Android 13, upgradable"
	specs["Available Colors"] = "Stealth Black, Arctic White, Prism Gold"
	specs["Announcement Date"] = "January 2024"
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
