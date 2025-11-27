package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileInfinixHot60i4g seeds specifications/options for product 'infinix-hot-60i-4g'
type SpecificationSeederMobileInfinixHot60i4g struct {
	BaseSeeder
}

// NewSpecificationSeederMobileInfinixHot60i4g creates a new seeder instance
func NewSpecificationSeederMobileInfinixHot60i4g() *SpecificationSeederMobileInfinixHot60i4g {
	return &SpecificationSeederMobileInfinixHot60i4g{BaseSeeder: BaseSeeder{name: "Specifications for Infinix Hot 60i 4G"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileInfinixHot60i4g) getBanglaTranslations() map[string]string {
	return map[string]string{
		"128 GB": "১২৮ GB",
		"164 x 76 x 8.5 mm": "১৬৪ x ৭৬ x ৮.৫ মিমি",
		"190 g": "১৯০ g",
		"2024": "২০২৪",
		"32 MP": "৩২ MP",
		"4 GB / 8 GB": "৪ GB / ৮ GB",
		"4G": "৪G",
		"5,000 mAh": "৫,০০০ এমএএইচ",
		"50 MP + AI": "৫০ MP + AI",
		"6.6 inches": "৬.৬ ইঞ্চি",
		"720 x 1612 pixels": "৭২০ x ১৬১২ pixels",
		"90Hz": "৯০Hz",
		"Android 14": "Android ১৪",
		"Black, Blue, Gold": "কালো, নীল, সোনালী",
		"Helio G85": "Helio G৮৫",
		"IPS LCD, 90Hz": "IPS LCD, ৯০Hz",
		"Mali-G52": "Mali-G৫২",
		"Mediatek Helio G85": "Mediatek Helio G৮৫",
	}
}

// Seed inserts specification records for the product identified by slug 'infinix-hot-60i-4g'
func (s *SpecificationSeederMobileInfinixHot60i4g) Seed(db *gorm.DB) error {
	productSlug := "infinix-hot-60i-4g"

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

	// Override model-specific values for Infinix Hot 60i 4G
	specs["Display Size"] = "6.6 inches"
	specs["Processor"] = "Helio G85"
	specs["Chipset"] = "Mediatek Helio G85"
	specs["Cpu Type"] = "Octa-core"
	specs["Gpu Type"] = "Mali-G52"
	specs["Ram"] = "4 GB / 8 GB"
	specs["Storage"] = "128 GB"
	specs["Display Type"] = "IPS LCD, 90Hz"
	specs["Resolution"] = "720 x 1612 pixels"
	specs["Screen Protection"] = "Glass front"
	specs["Refresh Rate"] = "90Hz"
	specs["Build Material"] = "Plastic body"
	specs["Weight"] = "190 g"
	specs["Dimensions"] = "164 x 76 x 8.5 mm"
	specs["Water Resistance"] = "No"
	specs["Network Technology"] = "4G"
	specs["Rear Camera"] = "50 MP + AI"
	specs["Front Camera"] = "32 MP"
	specs["Battery"] = "5,000 mAh"
	specs["Operating System"] = "Android 14"
	specs["Available Colors"] = "Black, Blue, Gold"
	specs["Announcement Date"] = "2024"
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
