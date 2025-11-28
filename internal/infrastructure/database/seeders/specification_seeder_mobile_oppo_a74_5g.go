package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileOppoA745g seeds specifications/options for product 'oppo-a74-5g'
type SpecificationSeederMobileOppoA745g struct {
	BaseSeeder
}

// NewSpecificationSeederMobileOppoA745g creates a new seeder instance
func NewSpecificationSeederMobileOppoA745g() *SpecificationSeederMobileOppoA745g {
	return &SpecificationSeederMobileOppoA745g{BaseSeeder: BaseSeeder{name: "Specifications for Oppo A74 5G"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileOppoA745g) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080p @ 30fps": "১০৮০p @ ৩০fps",
		"128 GB + microSD": "১২৮ জিবি + মাইক্রোএসডি",
		"16 MP": "১৬ মেগাপিক্সেল",
		"18 W wired": "১৮ W তারযুক্ত",
		"190 g": "১৯০ g",
		"2021": "২০২১",
		"2400 × 1080 px (~405 ppi)": "২৪০০ × ১০৮০ পিক্সেল (~৪০৫ পিপিআই)",
		"3.5 mm headphone jack": "৩.৫ মিমি হেডফোন জ্যাক",
		"4 / 6 GB": "৪ / ৬ GB",
		"48 MP + 2 MP + 2 MP": "৪৮ মেগাপিক্সেল + ২ মেগাপিক্সেল + ২ মেগাপিক্সেল",
		"5.1": "৫.১",
		"5000 mAh": "৫০০০ এমএএইচ",
		"6.5 inches": "৬.৫ ইঞ্চি",
		"60 Hz": "৬০ Hz",
		"Adreno 619": "অ্যাড্রেনো ৬১৯",
		"Android 11, ColorOS 11": "অ্যান্ড্রয়েড ১১, কালারওএস ১১",
		"Dual SIM (Nano-SIM)": "ডুয়াল সিম (ন্যানো-সিম)",
		"Fingerprint (side), Accelerometer, Proximity, Compass": "ফিঙ্গারপ্রিন্ট (side), অ্যাক্সিলেরোমিটার, প্রক্সিমিটি, কম্পাস",
		"GSM / HSPA / LTE / 5G": "জিএসএম / এইচএসপিএ / এলটিই / ৫G",
		"Li-Ion (non-removable)": "লি-আয়ন (অপসারণযোগ্য নয়)",
		"Octa-core (2×2.0 GHz Cortex-A76 + 6×1.8 GHz Cortex-A55)": "অক্টা-কোর (২×২.০ গিগাহার্টজ Cবাtex-A৭৬ + ৬×১.৮ গিগাহার্টজ Cবাtex-A৫৫)",
		"Prism Black, Space Silver": "প্রিজম কালো, স্পেস রূপালী",
		"Snapdragon 480 5G": "স্ন্যাপড্রাগন ৪৮০ ৫G",
		"Snapdragon 480 5G (8 nm)": "স্ন্যাপড্রাগন ৪৮০ ৫G (৮ ন্যানোমিটার)",
		"Wi-Fi 802.11 a/b/g/n/ac": "ওয়াই-ফাই ৮০২.১১ a/b/g/n/ac",
	}
}

// Seed inserts specification records for the product identified by slug 'oppo-a74-5g'
func (s *SpecificationSeederMobileOppoA745g) Seed(db *gorm.DB) error {
	productSlug := "oppo-a74-5g"

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

	// Override model-specific values for Oppo A74 5G
	specs["Display Size"] = "6.5 inches"
	specs["Processor"] = "Snapdragon 480 5G"
	specs["Chipset"] = "Snapdragon 480 5G (8 nm)"
	specs["Cpu Type"] = "Octa-core (2×2.0 GHz Cortex-A76 + 6×1.8 GHz Cortex-A55)"
	specs["Gpu Type"] = "Adreno 619"
	specs["Ram"] = "4 / 6 GB"
	specs["Storage"] = "128 GB + microSD"
	specs["Display Type"] = "IPS LCD"
	specs["Resolution"] = "2400 × 1080 px (~405 ppi)"
	specs["Refresh Rate"] = "60 Hz"
	specs["Build Material"] = "Plastic frame/back"
	specs["Weight"] = "190 g"
	specs["Network Technology"] = "GSM / HSPA / LTE / 5G"
	specs["Wifi Support"] = "Wi-Fi 802.11 a/b/g/n/ac"
	specs["Bluetooth Version"] = "5.1"
	specs["Nfc Support"] = "No"
	specs["Usb Type"] = "USB-C"
	specs["Rear Camera"] = "48 MP + 2 MP + 2 MP"
	specs["Camera Features"] = "LED flash, HDR"
	specs["Camera Video Resolution"] = "1080p @ 30fps"
	specs["Optical Zoom"] = "None"
	specs["Front Camera"] = "16 MP"
	specs["Front Camera Video Resolution"] = "1080p @ 30fps"
	specs["Operating System"] = "Android 11, ColorOS 11"
	specs["Battery"] = "5000 mAh"
	specs["Battery Type"] = "Li-Ion (non-removable)"
	specs["Fast Charging"] = "18 W wired"
	specs["Wireless Charging"] = "No"
	specs["5G Support"] = "Yes"
	specs["Positioning System"] = "GPS, GLONASS, GALILEO, BDS"
	specs["Sensors"] = "Fingerprint (side), Accelerometer, Proximity, Compass"
	specs["Sim Card Type"] = "Dual SIM (Nano-SIM)"
	specs["Loudspeaker Quality"] = "Mono"
	specs["Audio Jack"] = "3.5 mm headphone jack"
	specs["Available Colors"] = "Prism Black, Space Silver"
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
