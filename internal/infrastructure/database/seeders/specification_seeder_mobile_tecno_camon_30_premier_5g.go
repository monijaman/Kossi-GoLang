package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileTecnoCamon30Premier5g seeds specifications/options for product 'tecno-camon-30-premier-5g'
type SpecificationSeederMobileTecnoCamon30Premier5g struct {
	BaseSeeder
}

// NewSpecificationSeederMobileTecnoCamon30Premier5g creates a new seeder instance
func NewSpecificationSeederMobileTecnoCamon30Premier5g() *SpecificationSeederMobileTecnoCamon30Premier5g {
	return &SpecificationSeederMobileTecnoCamon30Premier5g{BaseSeeder: BaseSeeder{name: "Specifications for Tecno CAMON 30 Premier 5G"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileTecnoCamon30Premier5g) getBanglaTranslations() map[string]string {
	return map[string]string{
		"12 GB": "১২ GB",
		"120Hz": "১২০Hz",
		"1264 x 2780 pixels": "১২৬৪ x ২৭৮০ pixels",
		"162.7 x 76.2 x 7.9 mm": "১৬২.৭ x ৭৬.২ x ৭.৯ মিমি",
		"210 g": "২১০ g",
		"5,000 mAh": "৫,০০০ এমএএইচ",
		"50 MP": "৫০ MP",
		"50 MP + 50 MP + 50 MP": "৫০ MP + ৫০ MP + ৫০ MP",
		"512 GB": "৫১২ GB",
		"5G": "৫G",
		"6.77 inches": "৬.৭৭ ইঞ্চি",
		"Alps Snowy Silver, Hawaii Lava Black": "Alps Snowy রূপালী, Hawaii Lava কালো",
		"Android 14, HIOS 14": "Android ১৪, HIOS ১৪",
		"Corning Gorilla Glass 5": "Corning Gorilla Glass ৫",
		"Dimensity 8200 Ultimate": "Dimensity ৮২০০ Ultimate",
		"February 2024": "February ২০২৪",
		"Glass front, glass back, aluminum frame": "গ্লাস সামনে, গ্লাস পেছনে, অ্যালুমিনিয়াম ফ্রেম",
		"IP54": "IP৫৪",
		"LTPO AMOLED, 120Hz": "LTPO AMOLED, ১২০Hz",
		"Mali-G610 MC6": "Mali-G৬১০ MC৬",
		"Mediatek Dimensity 8200 Ultimate (4 nm)": "Mediatek Dimensity ৮২০০ Ultimate (৪ nm)",
	}
}

// Seed inserts specification records for the product identified by slug 'tecno-camon-30-premier-5g'
func (s *SpecificationSeederMobileTecnoCamon30Premier5g) Seed(db *gorm.DB) error {
	productSlug := "tecno-camon-30-premier-5g"

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

	// Override model-specific values for Tecno CAMON 30 Premier 5G
	specs["Display Size"] = "6.77 inches"
	specs["Processor"] = "Dimensity 8200 Ultimate"
	specs["Chipset"] = "Mediatek Dimensity 8200 Ultimate (4 nm)"
	specs["Cpu Type"] = "Octa-core"
	specs["Gpu Type"] = "Mali-G610 MC6"
	specs["Ram"] = "12 GB"
	specs["Storage"] = "512 GB"
	specs["Display Type"] = "LTPO AMOLED, 120Hz"
	specs["Resolution"] = "1264 x 2780 pixels"
	specs["Screen Protection"] = "Corning Gorilla Glass 5"
	specs["Refresh Rate"] = "120Hz"
	specs["Build Material"] = "Glass front, glass back, aluminum frame"
	specs["Weight"] = "210 g"
	specs["Dimensions"] = "162.7 x 76.2 x 7.9 mm"
	specs["Water Resistance"] = "IP54"
	specs["Network Technology"] = "5G"
	specs["Rear Camera"] = "50 MP + 50 MP + 50 MP"
	specs["Front Camera"] = "50 MP"
	specs["Battery"] = "5,000 mAh"
	specs["Operating System"] = "Android 14, HIOS 14"
	specs["Available Colors"] = "Alps Snowy Silver, Hawaii Lava Black"
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
