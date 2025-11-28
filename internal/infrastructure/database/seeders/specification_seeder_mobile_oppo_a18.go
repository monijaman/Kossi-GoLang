package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileOppoA18 seeds specifications/options for product 'oppo-a18'
type SpecificationSeederMobileOppoA18 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileOppoA18 creates a new seeder instance
func NewSpecificationSeederMobileOppoA18() *SpecificationSeederMobileOppoA18 {
	return &SpecificationSeederMobileOppoA18{BaseSeeder: BaseSeeder{name: "Specifications for Oppo A18"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileOppoA18) getBanglaTranslations() map[string]string {
	return map[string]string{
		"10 W wired": "১০ W তারযুক্ত",
		"1080p @ 30fps": "১০৮০p @ ৩০fps",
		"13 MP + 2 MP": "১৩ মেগাপিক্সেল + ২ মেগাপিক্সেল",
		"1612 × 720 px": "১৬১২ × ৭২০ পিক্সেল",
		"176 g": "১৭৬ g",
		"2021": "২০২১",
		"3.5 mm headphone jack": "৩.৫ মিমি হেডফোন জ্যাক",
		"4 GB": "৪ GB",
		"4230 mAh": "৪২৩০ এমএএইচ",
		"5.0": "৫.০",
		"6.56 inches": "৬.৫৬ ইঞ্চি",
		"60 Hz": "৬০ Hz",
		"64 GB + microSD": "৬৪ জিবি + মাইক্রোএসডি",
		"8 MP": "৮ মেগাপিক্সেল",
		"Accelerometer, Proximity": "অ্যাক্সিলেরোমিটার, প্রক্সিমিটি",
		"Android 11, ColorOS 11": "অ্যান্ড্রয়েড ১১, কালারওএস ১১",
		"Black, Blue": "কালো, নীল",
		"Dual SIM (Nano-SIM)": "ডুয়াল সিম (ন্যানো-সিম)",
		"Helio G35 (12 nm)": "হেলিও G৩৫ (১২ ন্যানোমিটার)",
		"Li-Ion (non-removable)": "লি-আয়ন (অপসারণযোগ্য নয়)",
		"MediaTek Helio G35": "মিডিয়াটেক হেলিও G৩৫",
		"PowerVR GE8320": "পাওয়ারভিআর GE৮৩২০",
		"Wi-Fi 802.11 a/b/g/n": "ওয়াই-ফাই ৮০২.১১ a/b/g/n",
	}
}

// Seed inserts specification records for the product identified by slug 'oppo-a18'
func (s *SpecificationSeederMobileOppoA18) Seed(db *gorm.DB) error {
	productSlug := "oppo-a18"

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

	// Override model-specific values for Oppo A18
	specs["Display Size"] = "6.56 inches"
	specs["Processor"] = "MediaTek Helio G35"
	specs["Chipset"] = "Helio G35 (12 nm)"
	specs["Cpu Type"] = "Octa-core"
	specs["Gpu Type"] = "PowerVR GE8320"
	specs["Ram"] = "4 GB"
	specs["Storage"] = "64 GB + microSD"
	specs["Display Type"] = "IPS LCD"
	specs["Resolution"] = "1612 × 720 px"
	specs["Refresh Rate"] = "60 Hz"
	specs["Build Material"] = "Plastic frame/back"
	specs["Weight"] = "176 g"
	specs["Network Technology"] = "GSM / HSPA / LTE"
	specs["Wifi Support"] = "Wi-Fi 802.11 a/b/g/n"
	specs["Bluetooth Version"] = "5.0"
	specs["Nfc Support"] = "No"
	specs["Usb Type"] = "Micro-USB"
	specs["Rear Camera"] = "13 MP + 2 MP"
	specs["Camera Features"] = "LED flash"
	specs["Camera Video Resolution"] = "1080p @ 30fps"
	specs["Optical Zoom"] = "None"
	specs["Front Camera"] = "8 MP"
	specs["Front Camera Video Resolution"] = "1080p @ 30fps"
	specs["Operating System"] = "Android 11, ColorOS 11"
	specs["Battery"] = "4230 mAh"
	specs["Battery Type"] = "Li-Ion (non-removable)"
	specs["Fast Charging"] = "10 W wired"
	specs["Wireless Charging"] = "No"
	specs["5G Support"] = "No"
	specs["Positioning System"] = "GPS, GLONASS"
	specs["Sensors"] = "Accelerometer, Proximity"
	specs["Sim Card Type"] = "Dual SIM (Nano-SIM)"
	specs["Loudspeaker Quality"] = "Mono"
	specs["Audio Jack"] = "3.5 mm headphone jack"
	specs["Available Colors"] = "Black, Blue"
	specs["Announcement Date"] = "2021"
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
