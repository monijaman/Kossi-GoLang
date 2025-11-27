package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileInfinixZero305g seeds specifications/options for product 'infinix-zero-30-5g'
type SpecificationSeederMobileInfinixZero305g struct {
	BaseSeeder
}

// NewSpecificationSeederMobileInfinixZero305g creates a new seeder instance
func NewSpecificationSeederMobileInfinixZero305g() *SpecificationSeederMobileInfinixZero305g {
	return &SpecificationSeederMobileInfinixZero305g{BaseSeeder: BaseSeeder{name: "Specifications for Infinix Zero 30 5G"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileInfinixZero305g) getBanglaTranslations() map[string]string {
	return map[string]string{
		"108 MP + 13 MP + 2 MP": "১০৮ MP + ১৩ MP + ২ MP",
		"1080 x 2400 pixels": "১০৮০ x ২৪০০ pixels",
		"144Hz": "১৪৪Hz",
		"164.5 x 75 x 7.9 mm": "১৬৪.৫ x ৭৫ x ৭.৯ মিমি",
		"185 g": "১৮৫ g",
		"256 GB": "২৫৬ GB",
		"5,000 mAh": "৫,০০০ এমএএইচ",
		"50 MP": "৫০ MP",
		"5G": "৫G",
		"6.78 inches": "৬.৭৮ ইঞ্চি",
		"8 GB / 12 GB": "৮ GB / ১২ GB",
		"AMOLED, 144Hz, Curved": "AMOLED, ১৪৪Hz, Curved",
		"Android 13, XOS 13": "Android ১৩, XOS ১৩",
		"Corning Gorilla Glass 5": "Corning Gorilla Glass ৫",
		"Dimensity 8020": "Dimensity ৮০২০",
		"Glass front, glass back, plastic frame": "গ্লাস সামনে, গ্লাস পেছনে, plastic frame",
		"IP53": "IP৫৩",
		"Mali-G77 MC9": "Mali-G৭৭ MC৯",
		"Mediatek Dimensity 8020 (6 nm)": "Mediatek Dimensity ৮০২০ (৬ nm)",
		"Rome Green, Golden Hour, Fantasy Purple": "Rome সবুজ, Golden Hour, Fantasy বেগুনি",
		"September 2023": "September ২০২৩",
	}
}

// Seed inserts specification records for the product identified by slug 'infinix-zero-30-5g'
func (s *SpecificationSeederMobileInfinixZero305g) Seed(db *gorm.DB) error {
	productSlug := "infinix-zero-30-5g"

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

	// Override model-specific values for Infinix Zero 30 5G
	specs["Display Size"] = "6.78 inches"
	specs["Processor"] = "Dimensity 8020"
	specs["Chipset"] = "Mediatek Dimensity 8020 (6 nm)"
	specs["Cpu Type"] = "Octa-core"
	specs["Gpu Type"] = "Mali-G77 MC9"
	specs["Ram"] = "8 GB / 12 GB"
	specs["Storage"] = "256 GB"
	specs["Display Type"] = "AMOLED, 144Hz, Curved"
	specs["Resolution"] = "1080 x 2400 pixels"
	specs["Screen Protection"] = "Corning Gorilla Glass 5"
	specs["Refresh Rate"] = "144Hz"
	specs["Build Material"] = "Glass front, glass back, plastic frame"
	specs["Weight"] = "185 g"
	specs["Dimensions"] = "164.5 x 75 x 7.9 mm"
	specs["Water Resistance"] = "IP53"
	specs["Network Technology"] = "5G"
	specs["Rear Camera"] = "108 MP + 13 MP + 2 MP"
	specs["Front Camera"] = "50 MP"
	specs["Battery"] = "5,000 mAh"
	specs["Operating System"] = "Android 13, XOS 13"
	specs["Available Colors"] = "Rome Green, Golden Hour, Fantasy Purple"
	specs["Announcement Date"] = "September 2023"
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
