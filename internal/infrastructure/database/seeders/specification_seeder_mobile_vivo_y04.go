package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileVivoY04 seeds specifications/options for product 'vivo-y04'
type SpecificationSeederMobileVivoY04 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileVivoY04 creates a new seeder instance
func NewSpecificationSeederMobileVivoY04() *SpecificationSeederMobileVivoY04 {
	return &SpecificationSeederMobileVivoY04{BaseSeeder: BaseSeeder{name: "Specifications for Vivo Y04"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileVivoY04) getBanglaTranslations() map[string]string {
	return map[string]string{
		"13 MP + 0.08 MP": "১৩ মেগাপিক্সেল + ০.০৮ মেগাপিক্সেল",
		"15W wired": "১৫W তারযুক্ত",
		"185 g": "১৮৫ g",
		"4 GB": "৪ GB",
		"5 MP": "৫ মেগাপিক্সেল",
		"5000 mAh": "৫০০০ এমএএইচ",
		"6.56 inches": "৬.৫৬ ইঞ্চি",
		"64 GB / 128 GB": "৬৪ জিবি / ১২৮ GB",
		"720 x 1612 pixels": "৭২০ x ১৬১২ পিক্সেল",
		"90Hz": "৯০Hz",
		"Android 14, Funtouch 14": "অ্যান্ড্রয়েড ১৪, Funথেকেuch ১৪",
		"Glass front, plastic back, plastic frame": "গ্লাস সামনে, প্লাস্টিক পেছনে, প্লাস্টিক ফ্রেম",
		"Helio G85": "হেলিও G৮৫",
		"IP54": "IP৫৪",
		"IPS LCD, 90Hz": "আইপিএস এলসিডি, ৯০Hz",
		"Mali-G52 MC2": "মালি-G৫২ MC২",
		"Mediatek Helio G85 (12nm)": "মিডিয়াটেক হেলিও G৮৫ (১২nm)",
		"September 2024": "সেপ্টেম্বর ২০২৪",
		"Space Black, Gem Green": "স্পেস কালো, Gem সবুজ",
	}
}

// Seed inserts specification records for the product identified by slug 'vivo-y04'
func (s *SpecificationSeederMobileVivoY04) Seed(db *gorm.DB) error {
	productSlug := "vivo-y04"

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

	// Override model-specific values for Vivo Y04
	specs["Display Size"] = "6.56 inches"
	specs["Processor"] = "Helio G85"
	specs["Chipset"] = "Mediatek Helio G85 (12nm)"
	specs["Cpu Type"] = "Octa-core"
	specs["Gpu Type"] = "Mali-G52 MC2"
	specs["Ram"] = "4 GB"
	specs["Storage"] = "64 GB / 128 GB"
	specs["Display Type"] = "IPS LCD, 90Hz"
	specs["Resolution"] = "720 x 1612 pixels"
	specs["Refresh Rate"] = "90Hz"
	specs["Build Material"] = "Glass front, plastic back, plastic frame"
	specs["Weight"] = "185 g"
	specs["Water Resistance"] = "IP54"
	specs["Network Technology"] = "GSM / HSPA / LTE"
	specs["Rear Camera"] = "13 MP + 0.08 MP"
	specs["Front Camera"] = "5 MP"
	specs["Battery"] = "5000 mAh"
	specs["Fast Charging"] = "15W wired"
	specs["Operating System"] = "Android 14, Funtouch 14"
	specs["Available Colors"] = "Space Black, Gem Green"
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
