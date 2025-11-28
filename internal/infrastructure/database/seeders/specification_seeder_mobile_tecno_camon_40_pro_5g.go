package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileTecnoCamon40Pro5g seeds specifications/options for product 'tecno-camon-40-pro-5g'
type SpecificationSeederMobileTecnoCamon40Pro5g struct {
	BaseSeeder
}

// NewSpecificationSeederMobileTecnoCamon40Pro5g creates a new seeder instance
func NewSpecificationSeederMobileTecnoCamon40Pro5g() *SpecificationSeederMobileTecnoCamon40Pro5g {
	return &SpecificationSeederMobileTecnoCamon40Pro5g{BaseSeeder: BaseSeeder{name: "Specifications for Tecno CAMON 40 Pro 5G"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileTecnoCamon40Pro5g) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080 x 2436 pixels": "১০৮০ x ২৪৩৬ পিক্সেল",
		"12 GB": "১২ GB",
		"120Hz": "১২০Hz",
		"164 x 74.5 x 7.7 mm": "১৬৪ x ৭৪.৫ x ৭.৭ মিমি",
		"200 g": "২০০ g",
		"5,000 mAh": "৫,০০০ এমএএইচ",
		"50 MP": "৫০ মেগাপিক্সেল",
		"50 MP + 50 MP + 2 MP": "৫০ মেগাপিক্সেল + ৫০ মেগাপিক্সেল + ২ মেগাপিক্সেল",
		"512 GB": "৫১২ GB",
		"5G": "৫G",
		"6.8 inches": "৬.৮ ইঞ্চি",
		"AMOLED, 120Hz": "অ্যামোলেড, ১২০Hz",
		"Android 14, HIOS 14": "অ্যান্ড্রয়েড ১৪, HIOS ১৪",
		"Corning Gorilla Glass 5": "কর্নিং গরিলা গ্লাস ৫",
		"Dimensity 8200 Ultimate": "ডাইমেনসিটি ৮২০০ আল্টিমেট",
		"February 2024": "ফেব্রুয়ারি ২০২৪",
		"Glass front, glass back": "গ্লাস সামনে, গ্লাস পেছনে",
		"IP54": "IP৫৪",
		"Iceland Basalt, Alps Snowy Silver": "আইসlএবং ব্যাসল্ট, Alps Snowy রূপালী",
		"Mali-G610 MC6": "মালি-G৬১০ MC৬",
		"Mediatek Dimensity 8200 Ultimate (4 nm)": "মিডিয়াটেক ডাইমেনসিটি ৮২০০ আল্টিমেট (৪ ন্যানোমিটার)",
	}
}

// Seed inserts specification records for the product identified by slug 'tecno-camon-40-pro-5g'
func (s *SpecificationSeederMobileTecnoCamon40Pro5g) Seed(db *gorm.DB) error {
	productSlug := "tecno-camon-40-pro-5g"

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

	// Override model-specific values for Tecno CAMON 40 Pro 5G
	specs["Display Size"] = "6.8 inches"
	specs["Processor"] = "Dimensity 8200 Ultimate"
	specs["Chipset"] = "Mediatek Dimensity 8200 Ultimate (4 nm)"
	specs["Cpu Type"] = "Octa-core"
	specs["Gpu Type"] = "Mali-G610 MC6"
	specs["Ram"] = "12 GB"
	specs["Storage"] = "512 GB"
	specs["Display Type"] = "AMOLED, 120Hz"
	specs["Resolution"] = "1080 x 2436 pixels"
	specs["Screen Protection"] = "Corning Gorilla Glass 5"
	specs["Refresh Rate"] = "120Hz"
	specs["Build Material"] = "Glass front, glass back"
	specs["Weight"] = "200 g"
	specs["Dimensions"] = "164 x 74.5 x 7.7 mm"
	specs["Water Resistance"] = "IP54"
	specs["Network Technology"] = "5G"
	specs["Rear Camera"] = "50 MP + 50 MP + 2 MP"
	specs["Front Camera"] = "50 MP"
	specs["Battery"] = "5,000 mAh"
	specs["Operating System"] = "Android 14, HIOS 14"
	specs["Available Colors"] = "Iceland Basalt, Alps Snowy Silver"
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
