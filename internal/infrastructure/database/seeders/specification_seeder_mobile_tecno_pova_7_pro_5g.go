package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileTecnoPova7Pro5g seeds specifications/options for product 'tecno-pova-7-pro-5g'
type SpecificationSeederMobileTecnoPova7Pro5g struct {
	BaseSeeder
}

// NewSpecificationSeederMobileTecnoPova7Pro5g creates a new seeder instance
func NewSpecificationSeederMobileTecnoPova7Pro5g() *SpecificationSeederMobileTecnoPova7Pro5g {
	return &SpecificationSeederMobileTecnoPova7Pro5g{BaseSeeder: BaseSeeder{name: "Specifications for Tecno POVA 7 Pro 5G"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileTecnoPova7Pro5g) getBanglaTranslations() map[string]string {
	return map[string]string{
		"108 MP + 2 MP": "১০৮ মেগাপিক্সেল + ২ মেগাপিক্সেল",
		"1080 x 2436 pixels": "১০৮০ x ২৪৩৬ পিক্সেল",
		"120Hz": "১২০Hz",
		"168.6 x 76.6 x 7.9 mm": "১৬৮.৬ x ৭৬.৬ x ৭.৯ মিমি",
		"198 g": "১৯৮ g",
		"256 GB": "২৫৬ GB",
		"32 MP": "৩২ মেগাপিক্সেল",
		"5G": "৫G",
		"6,000 mAh": "৬,০০০ এমএএইচ",
		"6.78 inches": "৬.৭৮ ইঞ্চি",
		"8 GB / 12 GB": "৮ জিবি / ১২ GB",
		"AMOLED, 120Hz": "অ্যামোলেড, ১২০Hz",
		"Android 14, HIOS 14": "অ্যান্ড্রয়েড ১৪, HIOS ১৪",
		"Dimensity 6080": "ডাইমেনসিটি ৬০৮০",
		"February 2024": "ফেব্রুয়ারি ২০২৪",
		"IP53": "IP৫৩",
		"Mali-G57 MC2": "মালি-G৫৭ MC২",
		"Mediatek Dimensity 6080 (6 nm)": "মিডিয়াটেক ডাইমেনসিটি ৬০৮০ (৬ ন্যানোমিটার)",
		"Meteorite Grey, Comet Green": "Meteবাite ধূসর, Comet সবুজ",
	}
}

// Seed inserts specification records for the product identified by slug 'tecno-pova-7-pro-5g'
func (s *SpecificationSeederMobileTecnoPova7Pro5g) Seed(db *gorm.DB) error {
	productSlug := "tecno-pova-7-pro-5g"

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

	// Override model-specific values for Tecno POVA 7 Pro 5G
	specs["Display Size"] = "6.78 inches"
	specs["Processor"] = "Dimensity 6080"
	specs["Chipset"] = "Mediatek Dimensity 6080 (6 nm)"
	specs["Cpu Type"] = "Octa-core"
	specs["Gpu Type"] = "Mali-G57 MC2"
	specs["Ram"] = "8 GB / 12 GB"
	specs["Storage"] = "256 GB"
	specs["Display Type"] = "AMOLED, 120Hz"
	specs["Resolution"] = "1080 x 2436 pixels"
	specs["Screen Protection"] = "Glass front"
	specs["Refresh Rate"] = "120Hz"
	specs["Build Material"] = "Plastic body"
	specs["Weight"] = "198 g"
	specs["Dimensions"] = "168.6 x 76.6 x 7.9 mm"
	specs["Water Resistance"] = "IP53"
	specs["Network Technology"] = "5G"
	specs["Rear Camera"] = "108 MP + 2 MP"
	specs["Front Camera"] = "32 MP"
	specs["Battery"] = "6,000 mAh"
	specs["Operating System"] = "Android 14, HIOS 14"
	specs["Available Colors"] = "Meteorite Grey, Comet Green"
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
