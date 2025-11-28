package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileVivoV40 seeds specifications/options for product 'vivo-v40'
type SpecificationSeederMobileVivoV40 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileVivoV40 creates a new seeder instance
func NewSpecificationSeederMobileVivoV40() *SpecificationSeederMobileVivoV40 {
	return &SpecificationSeederMobileVivoV40{BaseSeeder: BaseSeeder{name: "Specifications for Vivo V40"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileVivoV40) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080p @ 30fps": "১০৮০p @ ৩০fps",
		"120 Hz": "১২০ Hz",
		"128 / 256 GB": "১২৮ / ২৫৬ GB",
		"179 g": "১৭৯ g",
		"2400 × 1080 px": "২৪০০ × ১০৮০ পিক্সেল",
		"32 MP": "৩২ মেগাপিক্সেল",
		"4200 mAh": "৪২০০ এমএএইচ",
		"44 W wired": "৪৪ W তারযুক্ত",
		"4K @ 30fps; 1080p @ 30/60fps": "৪K @ ৩০fps; ১০৮০p @ ৩০/৬০fps",
		"5.2": "৫.২",
		"6.44 inches": "৬.৪৪ ইঞ্চি",
		"64 MP + 8 MP + 2 MP": "৬৪ মেগাপিক্সেল + ৮ মেগাপিক্সেল + ২ মেগাপিক্সেল",
		"8 / 12 GB": "৮ / ১২ GB",
		"AMOLED display, 120 Hz refresh rate": "অ্যামোলেড display, ১২০ Hz refresh rate",
		"AMOLED, 120 Hz": "অ্যামোলেড, ১২০ Hz",
		"Android 12, Funtouch OS 12": "অ্যান্ড্রয়েড ১২, Funথেকেuch OS ১২",
		"Black, Blue, Silver": "কালো, নীল, রূপালী",
		"Dimensity 920 (6 nm)": "ডাইমেনসিটি ৯২০ (৬ ন্যানোমিটার)",
		"Dual SIM (Nano-SIM)": "ডুয়াল সিম (ন্যানো-সিম)",
		"Fingerprint (under display), Accelerometer, Gyro, Proximity, Compass": "ফিঙ্গারপ্রিন্ট (ডিসপ্লের নিচে), অ্যাক্সিলেরোমিটার, জাইরো, প্রক্সিমিটি, কম্পাস",
		"GSM / HSPA / LTE / 5G": "জিএসএম / এইচএসপিএ / এলটিই / ৫G",
		"Glass front/back, aluminum frame": "গ্লাস সামনে/পেছনে, অ্যালুমিনিয়াম ফ্রেম",
		"Li-Ion (non-removable)": "লি-আয়ন (অপসারণযোগ্য নয়)",
		"Mali-G68 MC4": "মালি-G৬৮ MC৪",
		"May 2022": "মে ২০২২",
		"MediaTek Dimensity 920": "মিডিয়াটেক ডাইমেনসিটি ৯২০",
		"USB-C 2.0": "ইউএসবি-সি ২.০",
		"Wi-Fi 802.11 a/b/g/n/ac/6": "ওয়াই-ফাই ৮০২.১১ a/b/g/n/ac/৬",
	}
}

// Seed inserts specification records for the product identified by slug 'vivo-v40'
func (s *SpecificationSeederMobileVivoV40) Seed(db *gorm.DB) error {
	productSlug := "vivo-v40"

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

	// Override model-specific values for Vivo V40
	specs["Display Size"] = "6.44 inches"
	specs["Processor"] = "MediaTek Dimensity 920"
	specs["Chipset"] = "Dimensity 920 (6 nm)"
	specs["Cpu Type"] = "Octa-core"
	specs["Gpu Type"] = "Mali-G68 MC4"
	specs["Ram"] = "8 / 12 GB"
	specs["Storage"] = "128 / 256 GB"
	specs["Display Type"] = "AMOLED, 120 Hz"
	specs["Resolution"] = "2400 × 1080 px"
	specs["Refresh Rate"] = "120 Hz"
	specs["Build Material"] = "Glass front/back, aluminum frame"
	specs["Weight"] = "179 g"
	specs["Network Technology"] = "GSM / HSPA / LTE / 5G"
	specs["Wifi Support"] = "Wi-Fi 802.11 a/b/g/n/ac/6"
	specs["Bluetooth Version"] = "5.2"
	specs["Nfc Support"] = "Yes"
	specs["Usb Type"] = "USB-C 2.0"
	specs["Rear Camera"] = "64 MP + 8 MP + 2 MP"
	specs["Camera Features"] = "LED flash, OIS (main)"
	specs["Camera Video Resolution"] = "4K @ 30fps; 1080p @ 30/60fps"
	specs["Optical Zoom"] = "None"
	specs["Front Camera"] = "32 MP"
	specs["Front Camera Video Resolution"] = "1080p @ 30fps"
	specs["Operating System"] = "Android 12, Funtouch OS 12"
	specs["Battery"] = "4200 mAh"
	specs["Battery Type"] = "Li-Ion (non-removable)"
	specs["Fast Charging"] = "44 W wired"
	specs["Wireless Charging"] = "No"
	specs["5G Support"] = "Yes"
	specs["Positioning System"] = "GPS, GLONASS, GALILEO, BDS"
	specs["Sensors"] = "Fingerprint (under display), Accelerometer, Gyro, Proximity, Compass"
	specs["Special Features"] = "AMOLED display, 120 Hz refresh rate"
	specs["Sim Card Type"] = "Dual SIM (Nano-SIM)"
	specs["Loudspeaker Quality"] = "Stereo"
	specs["Audio Quality"] = "Hi-Res Audio"
	specs["Audio Jack"] = "No"
	specs["Available Colors"] = "Black, Blue, Silver"
	specs["Announcement Date"] = "May 2022"
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
