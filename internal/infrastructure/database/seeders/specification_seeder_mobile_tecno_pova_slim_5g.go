package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileTecnoPovaSlim5g seeds specifications/options for product 'tecno-pova-slim-5g'
type SpecificationSeederMobileTecnoPovaSlim5g struct {
	BaseSeeder
}

// NewSpecificationSeederMobileTecnoPovaSlim5g creates a new seeder instance
func NewSpecificationSeederMobileTecnoPovaSlim5g() *SpecificationSeederMobileTecnoPovaSlim5g {
	return &SpecificationSeederMobileTecnoPovaSlim5g{BaseSeeder: BaseSeeder{name: "Specifications for Tecno POVA Slim 5G"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileTecnoPovaSlim5g) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080 x 2460 pixels": "১০৮০ x ২৪৬০ pixels",
		"120Hz": "১২০Hz",
		"16 MP": "১৬ MP",
		"168.6 x 76.6 x 7.9 mm": "১৬৮.৬ x ৭৬.৬ x ৭.৯ মিমি",
		"198 g": "১৯৮ g",
		"256 GB": "২৫৬ GB",
		"50 MP + 2 MP": "৫০ MP + ২ MP",
		"5G": "৫G",
		"6,000 mAh": "৬,০০০ এমএএইচ",
		"6.78 inches": "৬.৭৮ ইঞ্চি",
		"8 GB": "৮ GB",
		"AMOLED, 120Hz": "AMOLED, ১২০Hz",
		"Android 14, HIOS 14": "Android ১৪, HIOS ১৪",
		"April 2024": "April ২০২৪",
		"Dimensity 6080": "Dimensity ৬০৮০",
		"IP53": "IP৫৩",
		"Mali-G57 MC2": "Mali-G৫৭ MC২",
		"Mecha Black, Amber Gold": "Mecha কালো, Amber সোনালী",
		"Mediatek Dimensity 6080 (6 nm)": "Mediatek Dimensity ৬০৮০ (৬ nm)",
	}
}

// Seed inserts specification records for the product identified by slug 'tecno-pova-slim-5g'
func (s *SpecificationSeederMobileTecnoPovaSlim5g) Seed(db *gorm.DB) error {
	productSlug := "tecno-pova-slim-5g"

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

	// Override model-specific values for Tecno POVA Slim 5G
	specs["Display Size"] = "6.78 inches"
	specs["Processor"] = "Dimensity 6080"
	specs["Chipset"] = "Mediatek Dimensity 6080 (6 nm)"
	specs["Cpu Type"] = "Octa-core"
	specs["Gpu Type"] = "Mali-G57 MC2"
	specs["Ram"] = "8 GB"
	specs["Storage"] = "256 GB"
	specs["Display Type"] = "AMOLED, 120Hz"
	specs["Resolution"] = "1080 x 2460 pixels"
	specs["Screen Protection"] = "Glass front"
	specs["Refresh Rate"] = "120Hz"
	specs["Build Material"] = "Plastic body"
	specs["Weight"] = "198 g"
	specs["Dimensions"] = "168.6 x 76.6 x 7.9 mm"
	specs["Water Resistance"] = "IP53"
	specs["Network Technology"] = "5G"
	specs["Rear Camera"] = "50 MP + 2 MP"
	specs["Front Camera"] = "16 MP"
	specs["Battery"] = "6,000 mAh"
	specs["Operating System"] = "Android 14, HIOS 14"
	specs["Available Colors"] = "Mecha Black, Amber Gold"
	specs["Announcement Date"] = "April 2024"
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
