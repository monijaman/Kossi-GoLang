package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileTecnoSpark30c seeds specifications/options for product 'tecno-spark-30c'
type SpecificationSeederMobileTecnoSpark30c struct {
	BaseSeeder
}

// NewSpecificationSeederMobileTecnoSpark30c creates a new seeder instance
func NewSpecificationSeederMobileTecnoSpark30c() *SpecificationSeederMobileTecnoSpark30c {
	return &SpecificationSeederMobileTecnoSpark30c{BaseSeeder: BaseSeeder{name: "Specifications for Tecno SPARK 30C"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileTecnoSpark30c) getBanglaTranslations() map[string]string {
	return map[string]string{
		"120Hz": "১২০Hz",
		"128 GB / 256 GB": "১২৮ GB / ২৫৬ GB",
		"165 x 76 x 8 mm": "১৬৫ x ৭৬ x ৮ মিমি",
		"182 g": "১৮২ g",
		"4 GB / 6 GB / 8 GB": "৪ GB / ৬ GB / ৮ GB",
		"4G": "৪G",
		"5,000 mAh": "৫,০০০ এমএএইচ",
		"50 MP + AI": "৫০ MP + AI",
		"6.67 inches": "৬.৬৭ ইঞ্চি",
		"720 x 1600 pixels": "৭২০ x ১৬০০ pixels",
		"8 MP": "৮ MP",
		"Android 14": "Android ১৪",
		"Helio G81": "Helio G৮১",
		"IP54": "IP৫৪",
		"IPS LCD, 120Hz": "IPS LCD, ১২০Hz",
		"Mali-G52": "Mali-G৫২",
		"Mediatek Helio G81": "Mediatek Helio G৮১",
		"Orbit Black, Orbit White, Magic Skin": "Orbit কালো, Orbit সাদা, Magic Skin",
		"September 2024": "September ২০২৪",
	}
}

// Seed inserts specification records for the product identified by slug 'tecno-spark-30c'
func (s *SpecificationSeederMobileTecnoSpark30c) Seed(db *gorm.DB) error {
	productSlug := "tecno-spark-30c"

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

	// Override model-specific values for Tecno SPARK 30C
	specs["Display Size"] = "6.67 inches"
	specs["Processor"] = "Helio G81"
	specs["Chipset"] = "Mediatek Helio G81"
	specs["Cpu Type"] = "Octa-core"
	specs["Gpu Type"] = "Mali-G52"
	specs["Ram"] = "4 GB / 6 GB / 8 GB"
	specs["Storage"] = "128 GB / 256 GB"
	specs["Display Type"] = "IPS LCD, 120Hz"
	specs["Resolution"] = "720 x 1600 pixels"
	specs["Screen Protection"] = "Glass front"
	specs["Refresh Rate"] = "120Hz"
	specs["Build Material"] = "Plastic body"
	specs["Weight"] = "182 g"
	specs["Dimensions"] = "165 x 76 x 8 mm"
	specs["Water Resistance"] = "IP54"
	specs["Network Technology"] = "4G"
	specs["Rear Camera"] = "50 MP + AI"
	specs["Front Camera"] = "8 MP"
	specs["Battery"] = "5,000 mAh"
	specs["Operating System"] = "Android 14"
	specs["Available Colors"] = "Orbit Black, Orbit White, Magic Skin"
	specs["Announcement Date"] = "September 2024"
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
