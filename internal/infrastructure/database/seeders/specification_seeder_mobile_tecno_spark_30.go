package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileTecnoSpark30 seeds specifications/options for product 'tecno-spark-30'
type SpecificationSeederMobileTecnoSpark30 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileTecnoSpark30 creates a new seeder instance
func NewSpecificationSeederMobileTecnoSpark30() *SpecificationSeederMobileTecnoSpark30 {
	return &SpecificationSeederMobileTecnoSpark30{BaseSeeder: BaseSeeder{name: "Specifications for Tecno SPARK 30"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileTecnoSpark30) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080 x 2460 pixels": "১০৮০ x ২৪৬০ pixels",
		"128 GB / 256 GB": "১২৮ GB / ২৫৬ GB",
		"16 MP": "১৬ MP",
		"164 x 75 x 8 mm": "১৬৪ x ৭৫ x ৮ মিমি",
		"186 g": "১৮৬ g",
		"4G": "৪G",
		"5,000 mAh": "৫,০০০ এমএএইচ",
		"6.6 inches": "৬.৬ ইঞ্চি",
		"64 MP + AI": "৬৪ MP + AI",
		"8 GB": "৮ GB",
		"90Hz": "৯০Hz",
		"Android 14": "Android ১৪",
		"Helio G91": "Helio G৯১",
		"IP53": "IP৫৩",
		"IPS LCD, 90Hz": "IPS LCD, ৯০Hz",
		"Mali-G52": "Mali-G৫২",
		"Mediatek Helio G91": "Mediatek Helio G৯১",
		"Orbit Black, Orbit White": "Orbit কালো, Orbit সাদা",
		"September 2024": "September ২০২৪",
	}
}

// Seed inserts specification records for the product identified by slug 'tecno-spark-30'
func (s *SpecificationSeederMobileTecnoSpark30) Seed(db *gorm.DB) error {
	productSlug := "tecno-spark-30"

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

	// Override model-specific values for Tecno SPARK 30
	specs["Display Size"] = "6.6 inches"
	specs["Processor"] = "Helio G91"
	specs["Chipset"] = "Mediatek Helio G91"
	specs["Cpu Type"] = "Octa-core"
	specs["Gpu Type"] = "Mali-G52"
	specs["Ram"] = "8 GB"
	specs["Storage"] = "128 GB / 256 GB"
	specs["Display Type"] = "IPS LCD, 90Hz"
	specs["Resolution"] = "1080 x 2460 pixels"
	specs["Screen Protection"] = "Glass front"
	specs["Refresh Rate"] = "90Hz"
	specs["Build Material"] = "Plastic body"
	specs["Weight"] = "186 g"
	specs["Dimensions"] = "164 x 75 x 8 mm"
	specs["Water Resistance"] = "IP53"
	specs["Network Technology"] = "4G"
	specs["Rear Camera"] = "64 MP + AI"
	specs["Front Camera"] = "16 MP"
	specs["Battery"] = "5,000 mAh"
	specs["Operating System"] = "Android 14"
	specs["Available Colors"] = "Orbit Black, Orbit White"
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
