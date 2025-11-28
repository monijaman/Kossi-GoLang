package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileVivoX300 seeds specifications/options for product 'vivo-x300'
type SpecificationSeederMobileVivoX300 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileVivoX300 creates a new seeder instance
func NewSpecificationSeederMobileVivoX300() *SpecificationSeederMobileVivoX300 {
	return &SpecificationSeederMobileVivoX300{BaseSeeder: BaseSeeder{name: "Specifications for Vivo X300"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileVivoX300) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080p @ 30fps": "১০৮০p @ ৩০fps",
		"120 Hz": "১২০ Hz",
		"128 / 256 GB": "১২৮ / ২৫৬ GB",
		"202 g": "২০২ g",
		"2x + 5x": "২x + ৫x",
		"32 MP": "৩২ মেগাপিক্সেল",
		"3200 × 1440 px": "৩২০০ × ১৪৪০ পিক্সেল",
		"4400 mAh": "৪৪০০ এমএএইচ",
		"48 MP + 13 MP + 13 MP": "৪৮ মেগাপিক্সেল + ১৩ মেগাপিক্সেল + ১৩ মেগাপিক্সেল",
		"4K @ 30/60fps; 1080p @ 30/120fps": "৪K @ ৩০/৬০fps; ১০৮০p @ ৩০/১২০fps",
		"5.2": "৫.২",
		"6.78 inches": "৬.৭৮ ইঞ্চি",
		"66 W wired": "৬৬ W তারযুক্ত",
		"8 / 12 GB": "৮ / ১২ GB",
		"AMOLED, 120 Hz": "অ্যামোলেড, ১২০ Hz",
		"Adreno 650": "অ্যাড্রেনো ৬৫০",
		"Android 12, Funtouch OS 12": "অ্যান্ড্রয়েড ১২, Funথেকেuch OS ১২",
		"Black, Blue, Silver": "কালো, নীল, রূপালী",
		"Dual SIM (Nano-SIM)": "ডুয়াল সিম (ন্যানো-সিম)",
		"Fingerprint (under display), Accelerometer, Gyro, Proximity, Compass": "ফিঙ্গারপ্রিন্ট (ডিসপ্লের নিচে), অ্যাক্সিলেরোমিটার, জাইরো, প্রক্সিমিটি, কম্পাস",
		"GSM / HSPA / LTE / 5G": "জিএসএম / এইচএসপিএ / এলটিই / ৫G",
		"Glass front/back, aluminum frame": "গ্লাস সামনে/পেছনে, অ্যালুমিনিয়াম ফ্রেম",
		"HDR10+, AMOLED display": "এইচডিআর১০+, অ্যামোলেড display",
		"January 2023": "জানুয়ারি ২০২৩",
		"Li-Ion (non-removable)": "লি-আয়ন (অপসারণযোগ্য নয়)",
		"Snapdragon 870": "স্ন্যাপড্রাগন ৮৭০",
		"Snapdragon 870 (7 nm)": "স্ন্যাপড্রাগন ৮৭০ (৭ ন্যানোমিটার)",
		"USB-C 3.1": "ইউএসবি-সি ৩.১",
		"Wi-Fi 802.11 a/b/g/n/ac/6": "ওয়াই-ফাই ৮০২.১১ a/b/g/n/ac/৬",
	}
}

// Seed inserts specification records for the product identified by slug 'vivo-x300'
func (s *SpecificationSeederMobileVivoX300) Seed(db *gorm.DB) error {
	productSlug := "vivo-x300"

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

	// Override model-specific values for Vivo X300
	specs["Display Size"] = "6.78 inches"
	specs["Processor"] = "Snapdragon 870"
	specs["Chipset"] = "Snapdragon 870 (7 nm)"
	specs["Cpu Type"] = "Octa-core"
	specs["Gpu Type"] = "Adreno 650"
	specs["Ram"] = "8 / 12 GB"
	specs["Storage"] = "128 / 256 GB"
	specs["Display Type"] = "AMOLED, 120 Hz"
	specs["Resolution"] = "3200 × 1440 px"
	specs["Refresh Rate"] = "120 Hz"
	specs["Build Material"] = "Glass front/back, aluminum frame"
	specs["Weight"] = "202 g"
	specs["Network Technology"] = "GSM / HSPA / LTE / 5G"
	specs["Wifi Support"] = "Wi-Fi 802.11 a/b/g/n/ac/6"
	specs["Bluetooth Version"] = "5.2"
	specs["Nfc Support"] = "Yes"
	specs["Usb Type"] = "USB-C 3.1"
	specs["Rear Camera"] = "48 MP + 13 MP + 13 MP"
	specs["Camera Features"] = "OIS, LED flash"
	specs["Camera Video Resolution"] = "4K @ 30/60fps; 1080p @ 30/120fps"
	specs["Optical Zoom"] = "2x + 5x"
	specs["Front Camera"] = "32 MP"
	specs["Front Camera Video Resolution"] = "1080p @ 30fps"
	specs["Operating System"] = "Android 12, Funtouch OS 12"
	specs["Battery"] = "4400 mAh"
	specs["Battery Type"] = "Li-Ion (non-removable)"
	specs["Fast Charging"] = "66 W wired"
	specs["Wireless Charging"] = "No"
	specs["5G Support"] = "Yes"
	specs["Positioning System"] = "GPS, GLONASS, GALILEO, BDS"
	specs["Sensors"] = "Fingerprint (under display), Accelerometer, Gyro, Proximity, Compass"
	specs["Special Features"] = "HDR10+, AMOLED display"
	specs["Sim Card Type"] = "Dual SIM (Nano-SIM)"
	specs["Loudspeaker Quality"] = "Stereo"
	specs["Audio Quality"] = "Hi-Res Audio"
	specs["Audio Jack"] = "No"
	specs["Available Colors"] = "Black, Blue, Silver"
	specs["Announcement Date"] = "January 2023"
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
