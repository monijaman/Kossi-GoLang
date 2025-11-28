package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobilePocoM6Plus5g seeds specifications/options for product 'poco-m6-plus-5g'
type SpecificationSeederMobilePocoM6Plus5g struct {
	BaseSeeder
}

// NewSpecificationSeederMobilePocoM6Plus5g creates a new seeder instance
func NewSpecificationSeederMobilePocoM6Plus5g() *SpecificationSeederMobilePocoM6Plus5g {
	return &SpecificationSeederMobilePocoM6Plus5g{BaseSeeder: BaseSeeder{name: "Specifications for POCO M6 Plus 5G"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobilePocoM6Plus5g) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080 × 2400 px": "১০৮০ × ২৪০০ পিক্সেল",
		"1080p @ 30fps": "১০৮০p @ ৩০fps",
		"120 Hz": "১২০ Hz",
		"200 g": "২০০ g",
		"2023": "২০২৩",
		"3.5 mm": "৩.৫ মিমি",
		"33 W wired": "৩৩ W তারযুক্ত",
		"4 / 6 GB": "৪ / ৬ GB",
		"5.1": "৫.১",
		"50 MP + 2 MP": "৫০ মেগাপিক্সেল + ২ মেগাপিক্সেল",
		"5000 mAh": "৫০০০ এমএএইচ",
		"6.7 inches": "৬.৭ ইঞ্চি",
		"64 / 128 GB": "৬৪ / ১২৮ GB",
		"8 MP": "৮ মেগাপিক্সেল",
		"Android 13, MIUI 14": "অ্যান্ড্রয়েড ১৩, এমআইইউআই ১৪",
		"Black, Green, Yellow": "কালো, সবুজ, হলুদ",
		"Dimensity 900 (6 nm)": "ডাইমেনসিটি ৯০০ (৬ ন্যানোমিটার)",
		"GSM / HSPA / LTE / 5G": "জিএসএম / এইচএসপিএ / এলটিই / ৫G",
		"IPS LCD, 120 Hz": "আইপিএস এলসিডি, ১২০ Hz",
		"Li-Ion (non-removable)": "লি-আয়ন (অপসারণযোগ্য নয়)",
		"Mali-G68": "মালি-G৬৮",
		"MediaTek Dimensity 900": "মিডিয়াটেক ডাইমেনসিটি ৯০০",
		"Nano-SIM + Nano-SIM": "ন্যানো-সিম + ন্যানো-সিম",
		"Plastic frame / back, Glass front": "Plastic ফ্রেম / পেছনে, গ্লাস সামনে",
		"Side fingerprint, Accelerometer, Gyro, Compass": "সাইড ফিঙ্গারপ্রিন্ট, অ্যাক্সিলেরোমিটার, জাইরো, কম্পাস",
		"Wi-Fi 802.11 a/b/g/n/ac": "ওয়াই-ফাই ৮০২.১১ a/b/g/n/ac",
	}
}

// Seed inserts specification records for the product identified by slug 'poco-m6-plus-5g'
func (s *SpecificationSeederMobilePocoM6Plus5g) Seed(db *gorm.DB) error {
	productSlug := "poco-m6-plus-5g"

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

	// Override model-specific values for POCO M6 Plus 5G
	specs["Display Size"] = "6.7 inches"
	specs["Processor"] = "MediaTek Dimensity 900"
	specs["Chipset"] = "Dimensity 900 (6 nm)"
	specs["Cpu Type"] = "Octa-core"
	specs["Gpu Type"] = "Mali-G68"
	specs["Ram"] = "4 / 6 GB"
	specs["Storage"] = "64 / 128 GB"
	specs["Display Type"] = "IPS LCD, 120 Hz"
	specs["Resolution"] = "1080 × 2400 px"
	specs["Refresh Rate"] = "120 Hz"
	specs["Build Material"] = "Plastic frame / back, Glass front"
	specs["Weight"] = "200 g"
	specs["Network Technology"] = "GSM / HSPA / LTE / 5G"
	specs["Wifi Support"] = "Wi-Fi 802.11 a/b/g/n/ac"
	specs["Bluetooth Version"] = "5.1"
	specs["Nfc Support"] = "Yes"
	specs["Usb Type"] = "USB-C"
	specs["Rear Camera"] = "50 MP + 2 MP"
	specs["Camera Features"] = "LED flash, HDR"
	specs["Camera Video Resolution"] = "1080p @ 30fps"
	specs["Optical Zoom"] = "None"
	specs["Front Camera"] = "8 MP"
	specs["Front Camera Video Resolution"] = "1080p @ 30fps"
	specs["Operating System"] = "Android 13, MIUI 14"
	specs["Battery"] = "5000 mAh"
	specs["Battery Type"] = "Li-Ion (non-removable)"
	specs["Fast Charging"] = "33 W wired"
	specs["Wireless Charging"] = "No"
	specs["5G Support"] = "Yes"
	specs["Positioning System"] = "GPS, GLONASS"
	specs["Sensors"] = "Side fingerprint, Accelerometer, Gyro, Compass"
	specs["Sim Card Type"] = "Nano-SIM + Nano-SIM"
	specs["Loudspeaker Quality"] = "Stereo"
	specs["Audio Jack"] = "3.5 mm"
	specs["Available Colors"] = "Black, Green, Yellow"
	specs["Announcement Date"] = "2023"
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
