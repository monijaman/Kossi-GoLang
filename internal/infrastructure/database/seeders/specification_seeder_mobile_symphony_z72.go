package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileSymphonyZ72 seeds specifications/options for product 'symphony-z72'
type SpecificationSeederMobileSymphonyZ72 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSymphonyZ72 creates a new seeder instance
func NewSpecificationSeederMobileSymphonyZ72() *SpecificationSeederMobileSymphonyZ72 {
	return &SpecificationSeederMobileSymphonyZ72{BaseSeeder: BaseSeeder{name: "Specifications for Symphony Z72"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSymphonyZ72) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080p@30fps": "১০৮০p@৩০fps",
		"10W wired": "১০W তারযুক্ত",
		"171.45 x 78.07 x 8.35 mm": "১৭১.৪৫ x ৭৮.০৭ x ৮.৩৫ মিমি",
		"195 g": "১৯৫ g",
		"4 GB": "৪ GB",
		"5000 mAh": "৫০০০ এমএএইচ",
		"52 MP + 2 MP": "৫২ MP + ২ MP",
		"6.88 inches": "৬.৮৮ ইঞ্চি",
		"64 GB / 128 GB": "৬৪ GB / ১২৮ GB",
		"650 MHz GPU": "৬৫০ MHz GPU",
		"720 × 1640 px": "৭২০ × ১৬৪০ px",
		"8 MP": "৮ MP",
		"Android 15": "Android ১৫",
		"April 2025": "April ২০২৫",
		"IP54 water & dust resistance": "IP৫৪ water & dust resistance",
		"Imperial Purple, Cosmic Gold, Titanium Gray, Graphite Black, Swamp Green": "Imperial বেগুনি, Cosmic সোনালী, Titanium ধূসর, Graphite কালো, Swamp সবুজ",
		"Octa-core 1.6 GHz": "Octa-core ১.৬ GHz",
		"Side fingerprint, accelerometer, proximity, compass, gyroscope": "Side ফিঙ্গারপ্রিন্ট, অ্যাক্সিলেরোমিটার, প্রক্সিমিটি, কম্পাস, gyroscope",
		"Unisoc T606": "Unisoc T৬০৬",
		"Unisoc T606 (12 nm)": "Unisoc T৬০৬ (১২ nm)",
	}
}

// Seed inserts specification records for the product identified by slug 'symphony-z72'
func (s *SpecificationSeederMobileSymphonyZ72) Seed(db *gorm.DB) error {
	productSlug := "symphony-z72"

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

	// Override model-specific values for Symphony Z72
	specs["Display Size"] = "6.88 inches"
	specs["Display Type"] = "IPS LCD"
	specs["Resolution"] = "720 × 1640 px"
	specs["Processor"] = "Unisoc T606"
	specs["Chipset"] = "Unisoc T606 (12 nm)"
	specs["Cpu Type"] = "Octa-core 1.6 GHz"
	specs["Gpu Type"] = "650 MHz GPU"
	specs["Ram"] = "4 GB"
	specs["Storage"] = "64 GB / 128 GB"
	specs["Operating System"] = "Android 15"
	specs["Rear Camera"] = "52 MP + 2 MP"
	specs["Camera Features"] = "HDR, LED flash"
	specs["Camera Video Resolution"] = "1080p@30fps"
	specs["Front Camera"] = "8 MP"
	specs["Battery"] = "5000 mAh"
	specs["Charging"] = "10W wired"
	specs["Network Technology"] = "GSM / HSPA / LTE"
	specs["Sim Card Type"] = "Dual Nano-SIM"
	specs["Connectivity"] = "Wi-Fi, Bluetooth, USB-C, OTG"
	specs["Sensors"] = "Side fingerprint, accelerometer, proximity, compass, gyroscope"
	specs["Special Features"] = "IP54 water & dust resistance"
	specs["Dimensions"] = "171.45 x 78.07 x 8.35 mm"
	specs["Weight"] = "195 g"
	specs["Build Material"] = "Plastic"
	specs["Available Colors"] = "Imperial Purple, Cosmic Gold, Titanium Gray, Graphite Black, Swamp Green"
	specs["Announcement Date"] = "April 2025"
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
