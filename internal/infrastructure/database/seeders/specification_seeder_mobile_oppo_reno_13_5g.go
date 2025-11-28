package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileOppoReno135g seeds specifications/options for product 'oppo-reno-13-5g'
type SpecificationSeederMobileOppoReno135g struct {
	BaseSeeder
}

// NewSpecificationSeederMobileOppoReno135g creates a new seeder instance
func NewSpecificationSeederMobileOppoReno135g() *SpecificationSeederMobileOppoReno135g {
	return &SpecificationSeederMobileOppoReno135g{BaseSeeder: BaseSeeder{name: "Specifications for Oppo Reno 13 5G"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileOppoReno135g) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080p @ 30fps": "১০৮০p @ ৩০fps",
		"120 Hz": "১২০ Hz",
		"128 / 256 GB": "১২৮ / ২৫৬ GB",
		"190 g": "১৯০ g",
		"2023": "২০২৩",
		"2412 × 1080 px": "২৪১২ × ১০৮০ পিক্সেল",
		"32 MP": "৩২ মেগাপিক্সেল",
		"33 W wired": "৩৩ W তারযুক্ত",
		"4500 mAh": "৪৫০০ এমএএইচ",
		"5.3": "৫.৩",
		"6.7 inches": "৬.৭ ইঞ্চি",
		"64 MP + 2 MP": "৬৪ মেগাপিক্সেল + ২ মেগাপিক্সেল",
		"8 GB": "৮ GB",
		"AMOLED, 120 Hz": "অ্যামোলেড, ১২০ Hz",
		"Android 14, ColorOS 14": "অ্যান্ড্রয়েড ১৪, কালারওএস ১৪",
		"Black, Silver": "কালো, রূপালী",
		"Dimensity 6080 (6 nm)": "ডাইমেনসিটি ৬০৮০ (৬ ন্যানোমিটার)",
		"Dual SIM (Nano-SIM)": "ডুয়াল সিম (ন্যানো-সিম)",
		"Fingerprint (under display), Accelerometer, Gyro, Proximity, Compass": "ফিঙ্গারপ্রিন্ট (ডিসপ্লের নিচে), অ্যাক্সিলেরোমিটার, জাইরো, প্রক্সিমিটি, কম্পাস",
		"GSM / HSPA / LTE / 5G": "জিএসএম / এইচএসপিএ / এলটিই / ৫G",
		"Glass front, plastic frame/back": "গ্লাস সামনে, প্লাস্টিক ফ্রেম/পেছনে",
		"HDR10+ display": "এইচডিআর১০+ display",
		"Li-Ion (non-removable)": "লি-আয়ন (অপসারণযোগ্য নয়)",
		"Mali-G57": "মালি-G৫৭",
		"MediaTek Dimensity 6080": "মিডিয়াটেক ডাইমেনসিটি ৬০৮০",
		"USB-C 2.0": "ইউএসবি-সি ২.০",
		"Wi-Fi 802.11 a/b/g/n/ac/6": "ওয়াই-ফাই ৮০২.১১ a/b/g/n/ac/৬",
	}
}

// Seed inserts specification records for the product identified by slug 'oppo-reno-13-5g'
func (s *SpecificationSeederMobileOppoReno135g) Seed(db *gorm.DB) error {
	productSlug := "oppo-reno-13-5g"

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

	// Override model-specific values for Oppo Reno 13 5G
	specs["Display Size"] = "6.7 inches"
	specs["Processor"] = "MediaTek Dimensity 6080"
	specs["Chipset"] = "Dimensity 6080 (6 nm)"
	specs["Cpu Type"] = "Octa-core"
	specs["Gpu Type"] = "Mali-G57"
	specs["Ram"] = "8 GB"
	specs["Storage"] = "128 / 256 GB"
	specs["Display Type"] = "AMOLED, 120 Hz"
	specs["Resolution"] = "2412 × 1080 px"
	specs["Refresh Rate"] = "120 Hz"
	specs["Build Material"] = "Glass front, plastic frame/back"
	specs["Weight"] = "190 g"
	specs["Network Technology"] = "GSM / HSPA / LTE / 5G"
	specs["Wifi Support"] = "Wi-Fi 802.11 a/b/g/n/ac/6"
	specs["Bluetooth Version"] = "5.3"
	specs["Nfc Support"] = "Yes"
	specs["Usb Type"] = "USB-C 2.0"
	specs["Rear Camera"] = "64 MP + 2 MP"
	specs["Camera Features"] = "LED flash"
	specs["Camera Video Resolution"] = "1080p @ 30fps"
	specs["Optical Zoom"] = "None"
	specs["Front Camera"] = "32 MP"
	specs["Front Camera Video Resolution"] = "1080p @ 30fps"
	specs["Operating System"] = "Android 14, ColorOS 14"
	specs["Battery"] = "4500 mAh"
	specs["Battery Type"] = "Li-Ion (non-removable)"
	specs["Fast Charging"] = "33 W wired"
	specs["Wireless Charging"] = "No"
	specs["5G Support"] = "Yes"
	specs["Positioning System"] = "GPS, GLONASS, GALILEO, BDS"
	specs["Sensors"] = "Fingerprint (under display), Accelerometer, Gyro, Proximity, Compass"
	specs["Special Features"] = "HDR10+ display"
	specs["Sim Card Type"] = "Dual SIM (Nano-SIM)"
	specs["Loudspeaker Quality"] = "Stereo"
	specs["Audio Jack"] = "No"
	specs["Available Colors"] = "Black, Silver"
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
