package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileSamsungGalaxyA15 seeds specifications/options for product 'samsung-galaxy-a15'
type SpecificationSeederMobileSamsungGalaxyA15 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSamsungGalaxyA15 creates a new seeder instance
func NewSpecificationSeederMobileSamsungGalaxyA15() *SpecificationSeederMobileSamsungGalaxyA15 {
	return &SpecificationSeederMobileSamsungGalaxyA15{BaseSeeder: BaseSeeder{name: "Specifications for Samsung Galaxy A15"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSamsungGalaxyA15) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080 x 2340 pixels": "১০৮০ x ২৩৪০ pixels",
		"128 GB / 256 GB": "১২৮ GB / ২৫৬ GB",
		"13 MP": "১৩ MP",
		"160.1 x 76.8 x 8.4 mm": "১৬০.১ x ৭৬.৮ x ৮.৪ মিমি",
		"200 g": "২০০ g",
		"25W wired": "২৫W তারযুক্ত",
		"4 GB / 6 GB / 8 GB": "৪ GB / ৬ GB / ৮ GB",
		"50 MP + 5 MP + 2 MP": "৫০ MP + ৫ MP + ২ MP",
		"5000 mAh": "৫০০০ এমএএইচ",
		"6.5 inches": "৬.৫ ইঞ্চি",
		"90Hz": "৯০Hz",
		"Android 14, One UI 6": "Android ১৪, One UI ৬",
		"Brave Black, Optimistic Blue, Magical Blue, Personality Yellow": "Brave কালো, Optimistic নীল, Magical নীল, Personality হলুদ",
		"December 2023": "December ২০২৩",
		"Glass front, plastic back, plastic frame": "গ্লাস সামনে, প্লাস্টিক পেছনে, plastic frame",
		"Mali-G57 MC2": "Mali-G৫৭ MC২",
		"Mediatek Helio G99": "Mediatek Helio G৯৯",
		"Mediatek Helio G99 (6nm)": "Mediatek Helio G৯৯ (৬nm)",
		"Super AMOLED, 90Hz, 800 nits": "Super AMOLED, ৯০Hz, ৮০০ nits",
	}
}

// Seed inserts specification records for the product identified by slug 'samsung-galaxy-a15'
func (s *SpecificationSeederMobileSamsungGalaxyA15) Seed(db *gorm.DB) error {
	productSlug := "samsung-galaxy-a15"

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

	// Override model-specific values for Samsung Galaxy A15
	specs["Display Size"] = "6.5 inches"
	specs["Processor"] = "Mediatek Helio G99"
	specs["Chipset"] = "Mediatek Helio G99 (6nm)"
	specs["Cpu Type"] = "Octa-core"
	specs["Gpu Type"] = "Mali-G57 MC2"
	specs["Ram"] = "4 GB / 6 GB / 8 GB"
	specs["Storage"] = "128 GB / 256 GB"
	specs["Display Type"] = "Super AMOLED, 90Hz, 800 nits"
	specs["Resolution"] = "1080 x 2340 pixels"
	specs["Refresh Rate"] = "90Hz"
	specs["Build Material"] = "Glass front, plastic back, plastic frame"
	specs["Weight"] = "200 g"
	specs["Dimensions"] = "160.1 x 76.8 x 8.4 mm"
	specs["Water Resistance"] = "No"
	specs["Network Technology"] = "GSM / HSPA / LTE"
	specs["Rear Camera"] = "50 MP + 5 MP + 2 MP"
	specs["Front Camera"] = "13 MP"
	specs["Battery"] = "5000 mAh"
	specs["Fast Charging"] = "25W wired"
	specs["Operating System"] = "Android 14, One UI 6"
	specs["Available Colors"] = "Brave Black, Optimistic Blue, Magical Blue, Personality Yellow"
	specs["Announcement Date"] = "December 2023"
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
