package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobilePocoF7 seeds specifications/options for product 'poco-f7'
type SpecificationSeederMobilePocoF7 struct {
	BaseSeeder
}

// NewSpecificationSeederMobilePocoF7 creates a new seeder instance
func NewSpecificationSeederMobilePocoF7() *SpecificationSeederMobilePocoF7 {
	return &SpecificationSeederMobilePocoF7{BaseSeeder: BaseSeeder{name: "Specifications for POCO F7"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobilePocoF7) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080p @ 30fps": "১০৮০p @ ৩০fps",
		"120 Hz": "১২০ Hz",
		"128 / 256 GB": "১২৮ / ২৫৬ GB",
		"196 g": "১৯৬ g",
		"20 MP": "২০ মেগাপিক্সেল",
		"2023": "২০২৩",
		"2400 × 1080 px": "২৪০০ × ১০৮০ পিক্সেল",
		"33 W wired": "৩৩ W তারযুক্ত",
		"4500 mAh": "৪৫০০ এমএএইচ",
		"4K @ 30/60fps; 1080p @ 30/60/120fps": "৪K @ ৩০/৬০fps; ১০৮০p @ ৩০/৬০/১২০fps",
		"5.2": "৫.২",
		"6 / 8 GB": "৬ / ৮ GB",
		"6.67 inches": "৬.৬৭ ইঞ্চি",
		"64 MP + 8 MP + 2 MP": "৬৪ মেগাপিক্সেল + ৮ মেগাপিক্সেল + ২ মেগাপিক্সেল",
		"AMOLED, 120 Hz": "অ্যামোলেড, ১২০ Hz",
		"Android 12, MIUI 13": "অ্যান্ড্রয়েড ১২, এমআইইউআই ১৩",
		"Black, Silver, Blue": "কালো, রূপালী, নীল",
		"Dimensity 1200 (6 nm)": "ডাইমেনসিটি ১২০০ (৬ ন্যানোমিটার)",
		"Fingerprint (side), Accelerometer, Gyro, Compass": "ফিঙ্গারপ্রিন্ট (side), অ্যাক্সিলেরোমিটার, জাইরো, কম্পাস",
		"GSM / HSPA / LTE / 5G": "জিএসএম / এইচএসপিএ / এলটিই / ৫G",
		"Glass front/back, plastic frame": "গ্লাস সামনে/পেছনে, প্লাস্টিক ফ্রেম",
		"Li-Ion (non-removable)": "লি-আয়ন (অপসারণযোগ্য নয়)",
		"Mali-G77 MC9": "মালি-G৭৭ MC৯",
		"MediaTek Dimensity 1200": "মিডিয়াটেক ডাইমেনসিটি ১২০০",
		"Nano-SIM + Nano-SIM": "ন্যানো-সিম + ন্যানো-সিম",
		"USB-C 2.0": "ইউএসবি-সি ২.০",
		"Wi-Fi 802.11 a/b/g/n/ac/6": "ওয়াই-ফাই ৮০২.১১ a/b/g/n/ac/৬",
	}
}

// Seed inserts specification records for the product identified by slug 'poco-f7'
func (s *SpecificationSeederMobilePocoF7) Seed(db *gorm.DB) error {
	productSlug := "poco-f7"

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

	// Override model-specific values for POCO F7
	specs["Display Size"] = "6.67 inches"
	specs["Processor"] = "MediaTek Dimensity 1200"
	specs["Chipset"] = "Dimensity 1200 (6 nm)"
	specs["Cpu Type"] = "Octa-core"
	specs["Gpu Type"] = "Mali-G77 MC9"
	specs["Ram"] = "6 / 8 GB"
	specs["Storage"] = "128 / 256 GB"
	specs["Display Type"] = "AMOLED, 120 Hz"
	specs["Resolution"] = "2400 × 1080 px"
	specs["Refresh Rate"] = "120 Hz"
	specs["Build Material"] = "Glass front/back, plastic frame"
	specs["Weight"] = "196 g"
	specs["Network Technology"] = "GSM / HSPA / LTE / 5G"
	specs["Wifi Support"] = "Wi-Fi 802.11 a/b/g/n/ac/6"
	specs["Bluetooth Version"] = "5.2"
	specs["Nfc Support"] = "Yes"
	specs["Usb Type"] = "USB-C 2.0"
	specs["Rear Camera"] = "64 MP + 8 MP + 2 MP"
	specs["Camera Features"] = "OIS, LED flash"
	specs["Camera Video Resolution"] = "4K @ 30/60fps; 1080p @ 30/60/120fps"
	specs["Optical Zoom"] = "None"
	specs["Front Camera"] = "20 MP"
	specs["Front Camera Video Resolution"] = "1080p @ 30fps"
	specs["Operating System"] = "Android 12, MIUI 13"
	specs["Battery"] = "4500 mAh"
	specs["Battery Type"] = "Li-Ion (non-removable)"
	specs["Fast Charging"] = "33 W wired"
	specs["Wireless Charging"] = "No"
	specs["5G Support"] = "Yes"
	specs["Positioning System"] = "GPS, GLONASS, GALILEO, BDS"
	specs["Sensors"] = "Fingerprint (side), Accelerometer, Gyro, Compass"
	specs["Sim Card Type"] = "Nano-SIM + Nano-SIM"
	specs["Loudspeaker Quality"] = "Stereo"
	specs["Audio Jack"] = "No"
	specs["Available Colors"] = "Black, Silver, Blue"
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
