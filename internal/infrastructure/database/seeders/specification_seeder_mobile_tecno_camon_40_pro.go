package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileTecnoCamon40Pro seeds specifications/options for product 'tecno-camon-40-pro'
type SpecificationSeederMobileTecnoCamon40Pro struct {
	BaseSeeder
}

// NewSpecificationSeederMobileTecnoCamon40Pro creates a new seeder instance
func NewSpecificationSeederMobileTecnoCamon40Pro() *SpecificationSeederMobileTecnoCamon40Pro {
	return &SpecificationSeederMobileTecnoCamon40Pro{BaseSeeder: BaseSeeder{name: "Specifications for Tecno CAMON 40 Pro"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileTecnoCamon40Pro) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080 x 2436 pixels": "১০৮০ x ২৪৩৬ পিক্সেল",
		"120Hz": "১২০Hz",
		"163.4 x 75.9 x 7.8 mm": "১৬৩.৪ x ৭৫.৯ x ৭.৮ মিমি",
		"188 g": "১৮৮ g",
		"256 GB": "২৫৬ GB",
		"32 MP": "৩২ মেগাপিক্সেল",
		"4G": "৪G",
		"5,000 mAh": "৫,০০০ এমএএইচ",
		"6.67 inches": "৬.৬৭ ইঞ্চি",
		"64 MP + 2 MP + 2 MP": "৬৪ মেগাপিক্সেল + ২ মেগাপিক্সেল + ২ মেগাপিক্সেল",
		"8 GB / 12 GB": "৮ জিবি / ১২ GB",
		"AMOLED, 120Hz": "অ্যামোলেড, ১২০Hz",
		"Android 14, HIOS 14": "অ্যান্ড্রয়েড ১৪, HIOS ১৪",
		"Black, Blue": "কালো, নীল",
		"February 2024": "ফেব্রুয়ারি ২০২৪",
		"Glass front, glass back": "গ্লাস সামনে, গ্লাস পেছনে",
		"Helio G99 Ultimate": "হেলিও G৯৯ আল্টিমেট",
		"IP53": "IP৫৩",
		"Mali-G57 MC2": "মালি-G৫৭ MC২",
		"Mediatek Helio G99 Ultimate": "মিডিয়াটেক হেলিও G৯৯ আল্টিমেট",
	}
}

// Seed inserts specification records for the product identified by slug 'tecno-camon-40-pro'
func (s *SpecificationSeederMobileTecnoCamon40Pro) Seed(db *gorm.DB) error {
	productSlug := "tecno-camon-40-pro"

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

	// Override model-specific values for Tecno CAMON 40 Pro
	specs["Display Size"] = "6.67 inches"
	specs["Processor"] = "Helio G99 Ultimate"
	specs["Chipset"] = "Mediatek Helio G99 Ultimate"
	specs["Cpu Type"] = "Octa-core"
	specs["Gpu Type"] = "Mali-G57 MC2"
	specs["Ram"] = "8 GB / 12 GB"
	specs["Storage"] = "256 GB"
	specs["Display Type"] = "AMOLED, 120Hz"
	specs["Resolution"] = "1080 x 2436 pixels"
	specs["Screen Protection"] = "Corning Gorilla Glass"
	specs["Refresh Rate"] = "120Hz"
	specs["Build Material"] = "Glass front, glass back"
	specs["Weight"] = "188 g"
	specs["Dimensions"] = "163.4 x 75.9 x 7.8 mm"
	specs["Water Resistance"] = "IP53"
	specs["Network Technology"] = "4G"
	specs["Rear Camera"] = "64 MP + 2 MP + 2 MP"
	specs["Front Camera"] = "32 MP"
	specs["Battery"] = "5,000 mAh"
	specs["Operating System"] = "Android 14, HIOS 14"
	specs["Available Colors"] = "Black, Blue"
	specs["Announcement Date"] = "February 2024"
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
