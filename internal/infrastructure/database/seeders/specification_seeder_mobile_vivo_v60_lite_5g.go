package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileVivoV60Lite5g seeds specifications/options for product 'vivo-v60-lite-5g'
type SpecificationSeederMobileVivoV60Lite5g struct {
	BaseSeeder
}

// NewSpecificationSeederMobileVivoV60Lite5g creates a new seeder instance
func NewSpecificationSeederMobileVivoV60Lite5g() *SpecificationSeederMobileVivoV60Lite5g {
	return &SpecificationSeederMobileVivoV60Lite5g{BaseSeeder: BaseSeeder{name: "Specifications for Vivo V60 Lite 5G"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileVivoV60Lite5g) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080p @ 30fps": "১০৮০p @ ৩০fps",
		"128 / 256 GB": "১২৮ / ২৫৬ GB",
		"16 MP": "১৬ মেগাপিক্সেল",
		"177 g": "১৭৭ g",
		"18 W wired": "১৮ W তারযুক্ত",
		"2400 × 1080 px": "২৪০০ × ১০৮০ পিক্সেল",
		"4100 mAh": "৪১০০ এমএএইচ",
		"5.2": "৫.২",
		"6 / 8 GB": "৬ / ৮ GB",
		"6.44 inches": "৬.৪৪ ইঞ্চি",
		"64 MP + 2 MP": "৬৪ মেগাপিক্সেল + ২ মেগাপিক্সেল",
		"90 Hz": "৯০ Hz",
		"AMOLED, 90 Hz": "অ্যামোলেড, ৯০ Hz",
		"Android 12, Funtouch OS 12": "অ্যান্ড্রয়েড ১২, Funথেকেuch OS ১২",
		"Black, Blue": "কালো, নীল",
		"Dimensity 900 (6 nm)": "ডাইমেনসিটি ৯০০ (৬ ন্যানোমিটার)",
		"Dual SIM (Nano-SIM)": "ডুয়াল সিম (ন্যানো-সিম)",
		"Fingerprint (under display), Accelerometer, Gyro, Proximity, Compass": "ফিঙ্গারপ্রিন্ট (ডিসপ্লের নিচে), অ্যাক্সিলেরোমিটার, জাইরো, প্রক্সিমিটি, কম্পাস",
		"GSM / HSPA / LTE / 5G": "জিএসএম / এইচএসপিএ / এলটিই / ৫G",
		"Glass front/back, plastic frame": "গ্লাস সামনে/পেছনে, প্লাস্টিক ফ্রেম",
		"June 2021": "জুন ২০২১",
		"Li-Ion (non-removable)": "লি-আয়ন (অপসারণযোগ্য নয়)",
		"Mali-G68 MC4": "মালি-G৬৮ MC৪",
		"MediaTek Dimensity 900": "মিডিয়াটেক ডাইমেনসিটি ৯০০",
		"USB-C 2.0": "ইউএসবি-সি ২.০",
		"Wi-Fi 802.11 a/b/g/n/ac/6": "ওয়াই-ফাই ৮০২.১১ a/b/g/n/ac/৬",
		"Yes, 3.5 mm": "হ্যাঁ, ৩.৫ মিমি",
	}
}

// Seed inserts specification records for the product identified by slug 'vivo-v60-lite-5g'
func (s *SpecificationSeederMobileVivoV60Lite5g) Seed(db *gorm.DB) error {
	productSlug := "vivo-v60-lite-5g"

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

	// Override model-specific values for Vivo V60 Lite 5G
	specs["Display Size"] = "6.44 inches"
	specs["Processor"] = "MediaTek Dimensity 900"
	specs["Chipset"] = "Dimensity 900 (6 nm)"
	specs["Cpu Type"] = "Octa-core"
	specs["Gpu Type"] = "Mali-G68 MC4"
	specs["Ram"] = "6 / 8 GB"
	specs["Storage"] = "128 / 256 GB"
	specs["Display Type"] = "AMOLED, 90 Hz"
	specs["Resolution"] = "2400 × 1080 px"
	specs["Refresh Rate"] = "90 Hz"
	specs["Build Material"] = "Glass front/back, plastic frame"
	specs["Weight"] = "177 g"
	specs["Network Technology"] = "GSM / HSPA / LTE / 5G"
	specs["Wifi Support"] = "Wi-Fi 802.11 a/b/g/n/ac/6"
	specs["Bluetooth Version"] = "5.2"
	specs["Nfc Support"] = "Yes"
	specs["Usb Type"] = "USB-C 2.0"
	specs["Rear Camera"] = "64 MP + 2 MP"
	specs["Camera Features"] = "LED flash"
	specs["Camera Video Resolution"] = "1080p @ 30fps"
	specs["Optical Zoom"] = "None"
	specs["Front Camera"] = "16 MP"
	specs["Front Camera Video Resolution"] = "1080p @ 30fps"
	specs["Operating System"] = "Android 12, Funtouch OS 12"
	specs["Battery"] = "4100 mAh"
	specs["Battery Type"] = "Li-Ion (non-removable)"
	specs["Fast Charging"] = "18 W wired"
	specs["Wireless Charging"] = "No"
	specs["5G Support"] = "Yes"
	specs["Positioning System"] = "GPS, GLONASS, GALILEO, BDS"
	specs["Sensors"] = "Fingerprint (under display), Accelerometer, Gyro, Proximity, Compass"
	specs["Special Features"] = "AMOLED display"
	specs["Sim Card Type"] = "Dual SIM (Nano-SIM)"
	specs["Loudspeaker Quality"] = "Stereo"
	specs["Audio Quality"] = "Hi-Res Audio"
	specs["Audio Jack"] = "Yes, 3.5 mm"
	specs["Available Colors"] = "Black, Blue"
	specs["Announcement Date"] = "June 2021"
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
