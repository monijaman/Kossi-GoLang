package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileRealme13 seeds specifications/options for product 'realme-13'
type SpecificationSeederMobileRealme13 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileRealme13 creates a new seeder instance
func NewSpecificationSeederMobileRealme13() *SpecificationSeederMobileRealme13 {
	return &SpecificationSeederMobileRealme13{BaseSeeder: BaseSeeder{name: "Specifications for Realme 13"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileRealme13) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080 x 2400 pixels": "১০৮০ x ২৪০০ pixels",
		"120Hz": "১২০Hz",
		"128 GB / 256 GB": "১২৮ GB / ২৫৬ GB",
		"16 MP": "১৬ MP",
		"190 g": "১৯০ g",
		"45W wired": "৪৫W তারযুক্ত",
		"50 MP + 2 MP": "৫০ MP + ২ MP",
		"5000 mAh": "৫০০০ এমএএইচ",
		"6.72 inches": "৬.৭২ ইঞ্চি",
		"8 GB": "৮ GB",
		"Android 14, Realme UI 5.0": "Android ১৪, Realme UI ৫.০",
		"August 2024": "August ২০২৪",
		"Dimensity 6300": "Dimensity ৬৩০০",
		"GSM / HSPA / LTE / 5G": "GSM / HSPA / LTE / ৫G",
		"Glass front, plastic back, plastic frame": "গ্লাস সামনে, প্লাস্টিক পেছনে, plastic frame",
		"IP64": "IP৬৪",
		"IPS LCD, 120Hz": "IPS LCD, ১২০Hz",
		"Mali-G57 MC2": "Mali-G৫৭ MC২",
		"Mediatek Dimensity 6300 (6 nm)": "Mediatek Dimensity ৬৩০০ (৬ nm)",
		"Speed Green, Dark Purple": "Speed সবুজ, Dark বেগুনি",
	}
}

// Seed inserts specification records for the product identified by slug 'realme-13'
func (s *SpecificationSeederMobileRealme13) Seed(db *gorm.DB) error {
	productSlug := "realme-13"

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

	// Override model-specific values for Realme 13
	specs["Display Size"] = "6.72 inches"
	specs["Processor"] = "Dimensity 6300"
	specs["Chipset"] = "Mediatek Dimensity 6300 (6 nm)"
	specs["Cpu Type"] = "Octa-core"
	specs["Gpu Type"] = "Mali-G57 MC2"
	specs["Ram"] = "8 GB"
	specs["Storage"] = "128 GB / 256 GB"
	specs["Display Type"] = "IPS LCD, 120Hz"
	specs["Resolution"] = "1080 x 2400 pixels"
	specs["Refresh Rate"] = "120Hz"
	specs["Build Material"] = "Glass front, plastic back, plastic frame"
	specs["Weight"] = "190 g"
	specs["Water Resistance"] = "IP64"
	specs["Network Technology"] = "GSM / HSPA / LTE / 5G"
	specs["Rear Camera"] = "50 MP + 2 MP"
	specs["Front Camera"] = "16 MP"
	specs["Battery"] = "5000 mAh"
	specs["Fast Charging"] = "45W wired"
	specs["Operating System"] = "Android 14, Realme UI 5.0"
	specs["Available Colors"] = "Speed Green, Dark Purple"
	specs["Announcement Date"] = "August 2024"
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
