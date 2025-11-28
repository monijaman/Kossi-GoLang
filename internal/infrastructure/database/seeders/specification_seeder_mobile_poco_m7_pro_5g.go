package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobilePocoM7Pro5g seeds specifications/options for product 'poco-m7-pro-5g'
type SpecificationSeederMobilePocoM7Pro5g struct {
	BaseSeeder
}

// NewSpecificationSeederMobilePocoM7Pro5g creates a new seeder instance
func NewSpecificationSeederMobilePocoM7Pro5g() *SpecificationSeederMobilePocoM7Pro5g {
	return &SpecificationSeederMobilePocoM7Pro5g{BaseSeeder: BaseSeeder{name: "Specifications for POCO M7 Pro 5G"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobilePocoM7Pro5g) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080 × 2400 px": "১০৮০ × ২৪০০ পিক্সেল",
		"1080p @ 30/60fps": "১০৮০p @ ৩০/৬০fps",
		"1080p @ 30fps": "১০৮০p @ ৩০fps",
		"120 Hz": "১২০ Hz",
		"190 g": "১৯০ g",
		"2024": "২০২৪",
		"3.5 mm": "৩.৫ মিমি",
		"33 W wired": "৩৩ W তারযুক্ত",
		"4 / 6 GB": "৪ / ৬ GB",
		"5.1": "৫.১",
		"50 MP + 2 MP": "৫০ মেগাপিক্সেল + ২ মেগাপিক্সেল",
		"5000 mAh": "৫০০০ এমএএইচ",
		"6.6 inches": "৬.৬ ইঞ্চি",
		"64 / 128 GB": "৬৪ / ১২৮ GB",
		"8 MP": "৮ মেগাপিক্সেল",
		"Android 13, MIUI 14": "অ্যান্ড্রয়েড ১৩, এমআইইউআই ১৪",
		"Black, Blue, Green": "কালো, নীল, সবুজ",
		"Dimensity 930 (6 nm)": "ডাইমেনসিটি ৯৩০ (৬ ন্যানোমিটার)",
		"GSM / HSPA / LTE / 5G": "জিএসএম / এইচএসপিএ / এলটিই / ৫G",
		"IPS LCD, 120 Hz": "আইপিএস এলসিডি, ১২০ Hz",
		"Li-Ion (non-removable)": "লি-আয়ন (অপসারণযোগ্য নয়)",
		"Mali-G68": "মালি-G৬৮",
		"MediaTek Dimensity 930": "মিডিয়াটেক ডাইমেনসিটি ৯৩০",
		"Nano-SIM + Nano-SIM": "ন্যানো-সিম + ন্যানো-সিম",
		"Octa-core (2×2.2 GHz Cortex-A78 + 6×2.0 GHz Cortex-A55)": "অক্টা-কোর (২×২.২ গিগাহার্টজ Cবাtex-A৭৮ + ৬×২.০ গিগাহার্টজ Cবাtex-A৫৫)",
		"Plastic frame / back, Glass front": "Plastic ফ্রেম / পেছনে, গ্লাস সামনে",
		"Side fingerprint, Accelerometer, Gyro, Compass": "সাইড ফিঙ্গারপ্রিন্ট, অ্যাক্সিলেরোমিটার, জাইরো, কম্পাস",
		"Wi-Fi 802.11 a/b/g/n/ac": "ওয়াই-ফাই ৮০২.১১ a/b/g/n/ac",
	}
}

// Seed inserts specification records for the product identified by slug 'poco-m7-pro-5g'
func (s *SpecificationSeederMobilePocoM7Pro5g) Seed(db *gorm.DB) error {
	productSlug := "poco-m7-pro-5g"

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

	// Override model-specific values for POCO M7 Pro 5G
	specs["Display Size"] = "6.6 inches"
	specs["Processor"] = "MediaTek Dimensity 930"
	specs["Chipset"] = "Dimensity 930 (6 nm)"
	specs["Cpu Type"] = "Octa-core (2×2.2 GHz Cortex-A78 + 6×2.0 GHz Cortex-A55)"
	specs["Gpu Type"] = "Mali-G68"
	specs["Ram"] = "4 / 6 GB"
	specs["Storage"] = "64 / 128 GB"
	specs["Display Type"] = "IPS LCD, 120 Hz"
	specs["Resolution"] = "1080 × 2400 px"
	specs["Refresh Rate"] = "120 Hz"
	specs["Build Material"] = "Plastic frame / back, Glass front"
	specs["Weight"] = "190 g"
	specs["Network Technology"] = "GSM / HSPA / LTE / 5G"
	specs["Wifi Support"] = "Wi-Fi 802.11 a/b/g/n/ac"
	specs["Bluetooth Version"] = "5.1"
	specs["Nfc Support"] = "Yes"
	specs["Usb Type"] = "USB-C"
	specs["Rear Camera"] = "50 MP + 2 MP"
	specs["Camera Features"] = "LED flash, HDR"
	specs["Camera Video Resolution"] = "1080p @ 30/60fps"
	specs["Optical Zoom"] = "None"
	specs["Front Camera"] = "8 MP"
	specs["Front Camera Video Resolution"] = "1080p @ 30fps"
	specs["Operating System"] = "Android 13, MIUI 14"
	specs["Battery"] = "5000 mAh"
	specs["Battery Type"] = "Li-Ion (non-removable)"
	specs["Fast Charging"] = "33 W wired"
	specs["Wireless Charging"] = "No"
	specs["5G Support"] = "Yes"
	specs["Positioning System"] = "GPS, GLONASS, Galileo, BDS"
	specs["Sensors"] = "Side fingerprint, Accelerometer, Gyro, Compass"
	specs["Sim Card Type"] = "Nano-SIM + Nano-SIM"
	specs["Loudspeaker Quality"] = "Stereo"
	specs["Audio Jack"] = "3.5 mm"
	specs["Available Colors"] = "Black, Blue, Green"
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
