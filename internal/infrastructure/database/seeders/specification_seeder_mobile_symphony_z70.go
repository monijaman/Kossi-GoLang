package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileSymphonyZ70 seeds specifications/options for product 'symphony-z70'
type SpecificationSeederMobileSymphonyZ70 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSymphonyZ70 creates a new seeder instance
func NewSpecificationSeederMobileSymphonyZ70() *SpecificationSeederMobileSymphonyZ70 {
	return &SpecificationSeederMobileSymphonyZ70{BaseSeeder: BaseSeeder{name: "Specifications for Symphony Z70"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSymphonyZ70) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080p@30fps": "১০৮০p@৩০fps",
		"10W wired": "১০W তারযুক্ত",
		"164.27 x 76.02 x 8.45 mm": "১৬৪.২৭ x ৭৬.০২ x ৮.৪৫ মিমি",
		"193 g": "১৯৩ g",
		"4 GB": "৪ GB",
		"5000 mAh": "৫০০০ এমএএইচ",
		"52 MP + 2 MP + 2 MP": "৫২ MP + ২ MP + ২ MP",
		"6.56 inches": "৬.৫৬ ইঞ্চি",
		"64 GB": "৬৪ GB",
		"650 MHz GPU": "৬৫০ MHz GPU",
		"720 × 1612 px": "৭২০ × ১৬১২ px",
		"8 MP": "৮ MP",
		"Android 13": "Android ১৩",
		"November 2024": "November ২০২৪",
		"Octa-core 1.6 GHz": "Octa-core ১.৬ GHz",
		"Reflective Green, Electric Blue, Honey Dew Green, Fusion Gold": "Reflective সবুজ, Electric নীল, Honey Dew সবুজ, Fusion সোনালী",
		"Side fingerprint, accelerometer, proximity, compass": "Side ফিঙ্গারপ্রিন্ট, অ্যাক্সিলেরোমিটার, প্রক্সিমিটি, কম্পাস",
		"Unisoc T606": "Unisoc T৬০৬",
		"Unisoc T606 (12 nm)": "Unisoc T৬০৬ (১২ nm)",
		"Wi-Fi, Bluetooth 5.0, USB-C, OTG": "Wi-Fi, Bluetooth ৫.০, USB-C, OTG",
	}
}

// Seed inserts specification records for the product identified by slug 'symphony-z70'
func (s *SpecificationSeederMobileSymphonyZ70) Seed(db *gorm.DB) error {
	productSlug := "symphony-z70"

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

	// Override model-specific values for Symphony Z70
	specs["Display Size"] = "6.56 inches"
	specs["Display Type"] = "IPS In-cell"
	specs["Resolution"] = "720 × 1612 px"
	specs["Processor"] = "Unisoc T606"
	specs["Chipset"] = "Unisoc T606 (12 nm)"
	specs["Cpu Type"] = "Octa-core 1.6 GHz"
	specs["Gpu Type"] = "650 MHz GPU"
	specs["Ram"] = "4 GB"
	specs["Storage"] = "64 GB"
	specs["Operating System"] = "Android 13"
	specs["Rear Camera"] = "52 MP + 2 MP + 2 MP"
	specs["Camera Features"] = "AI, HDR, Panorama"
	specs["Camera Video Resolution"] = "1080p@30fps"
	specs["Front Camera"] = "8 MP"
	specs["Battery"] = "5000 mAh"
	specs["Charging"] = "10W wired"
	specs["Network Technology"] = "GSM / HSPA / LTE"
	specs["Sim Card Type"] = "Dual Nano-SIM"
	specs["Connectivity"] = "Wi-Fi, Bluetooth 5.0, USB-C, OTG"
	specs["Sensors"] = "Side fingerprint, accelerometer, proximity, compass"
	specs["Dimensions"] = "164.27 x 76.02 x 8.45 mm"
	specs["Weight"] = "193 g"
	specs["Build Material"] = "Plastic"
	specs["Available Colors"] = "Reflective Green, Electric Blue, Honey Dew Green, Fusion Gold"
	specs["Announcement Date"] = "November 2024"
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
