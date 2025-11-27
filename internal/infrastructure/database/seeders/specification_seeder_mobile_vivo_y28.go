package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileVivoY28 seeds specifications/options for product 'vivo-y28'
type SpecificationSeederMobileVivoY28 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileVivoY28 creates a new seeder instance
func NewSpecificationSeederMobileVivoY28() *SpecificationSeederMobileVivoY28 {
	return &SpecificationSeederMobileVivoY28{BaseSeeder: BaseSeeder{name: "Specifications for Vivo Y28"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileVivoY28) getBanglaTranslations() map[string]string {
	return map[string]string{
		"128 / 256 GB": "১২৮ / ২৫৬ GB",
		"199 g": "১৯৯ g",
		"44W wired": "৪৪W তারযুক্ত",
		"50 MP + 2 MP": "৫০ MP + ২ MP",
		"6.68 inches": "৬.৬৮ ইঞ্চি",
		"6000 mAh": "৬০০০ এমএএইচ",
		"720 x 1608 pixels": "৭২০ x ১৬০৮ pixels",
		"8 GB": "৮ GB",
		"8 MP": "৮ MP",
		"90Hz": "৯০Hz",
		"Agate Green, Gleaming Orange": "Agate সবুজ, Gleaming Orange",
		"Android 14, Funtouch 14": "Android ১৪, Funtouch ১৪",
		"Glass front, plastic back, plastic frame": "গ্লাস সামনে, প্লাস্টিক পেছনে, plastic frame",
		"Helio G85": "Helio G৮৫",
		"IP64": "IP৬৪",
		"IPS LCD, 90Hz": "IPS LCD, ৯০Hz",
		"January 2024": "January ২০২৪",
		"Mali-G52 MC2": "Mali-G৫২ MC২",
		"Mediatek Helio G85 (12nm)": "Mediatek Helio G৮৫ (১২nm)",
	}
}

// Seed inserts specification records for the product identified by slug 'vivo-y28'
func (s *SpecificationSeederMobileVivoY28) Seed(db *gorm.DB) error {
	productSlug := "vivo-y28"

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

	// Override model-specific values for Vivo Y28
	specs["Display Size"] = "6.68 inches"
	specs["Processor"] = "Helio G85"
	specs["Chipset"] = "Mediatek Helio G85 (12nm)"
	specs["Cpu Type"] = "Octa-core"
	specs["Gpu Type"] = "Mali-G52 MC2"
	specs["Ram"] = "8 GB"
	specs["Storage"] = "128 / 256 GB"
	specs["Display Type"] = "IPS LCD, 90Hz"
	specs["Resolution"] = "720 x 1608 pixels"
	specs["Refresh Rate"] = "90Hz"
	specs["Build Material"] = "Glass front, plastic back, plastic frame"
	specs["Weight"] = "199 g"
	specs["Water Resistance"] = "IP64"
	specs["Network Technology"] = "GSM / HSPA / LTE"
	specs["Rear Camera"] = "50 MP + 2 MP"
	specs["Front Camera"] = "8 MP"
	specs["Battery"] = "6000 mAh"
	specs["Fast Charging"] = "44W wired"
	specs["Operating System"] = "Android 14, Funtouch 14"
	specs["Available Colors"] = "Agate Green, Gleaming Orange"
	specs["Announcement Date"] = "January 2024"
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
