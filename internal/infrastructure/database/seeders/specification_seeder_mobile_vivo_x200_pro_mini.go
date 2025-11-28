package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileVivoX200ProMini seeds specifications/options for product 'vivo-x200-pro-mini'
type SpecificationSeederMobileVivoX200ProMini struct {
	BaseSeeder
}

// NewSpecificationSeederMobileVivoX200ProMini creates a new seeder instance
func NewSpecificationSeederMobileVivoX200ProMini() *SpecificationSeederMobileVivoX200ProMini {
	return &SpecificationSeederMobileVivoX200ProMini{BaseSeeder: BaseSeeder{name: "Specifications for Vivo X200 Pro Mini"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileVivoX200ProMini) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080p @ 30fps": "১০৮০p @ ৩০fps",
		"128 GB": "১২৮ GB",
		"185 g": "১৮৫ g",
		"2400 × 1080 px": "২৪০০ × ১০৮০ পিক্সেল",
		"2x": "২x",
		"32 MP": "৩২ মেগাপিক্সেল",
		"33 W wired": "৩৩ W তারযুক্ত",
		"4000 mAh": "৪০০০ এমএএইচ",
		"48 MP + 8 MP + 2 MP": "৪৮ মেগাপিক্সেল + ৮ মেগাপিক্সেল + ২ মেগাপিক্সেল",
		"4K @ 30fps; 1080p @ 30fps": "৪K @ ৩০fps; ১০৮০p @ ৩০fps",
		"5.2": "৫.২",
		"6 / 8 GB": "৬ / ৮ GB",
		"6.44 inches": "৬.৪৪ ইঞ্চি",
		"90 Hz": "৯০ Hz",
		"AMOLED, 90 Hz": "অ্যামোলেড, ৯০ Hz",
		"Android 11, Funtouch OS 11": "অ্যান্ড্রয়েড ১১, Funথেকেuch OS ১১",
		"Black, Blue": "কালো, নীল",
		"Dimensity 1100 (6 nm)": "ডাইমেনসিটি ১১০০ (৬ ন্যানোমিটার)",
		"Dual SIM (Nano-SIM)": "ডুয়াল সিম (ন্যানো-সিম)",
		"Fingerprint (under display), Accelerometer, Gyro, Proximity, Compass": "ফিঙ্গারপ্রিন্ট (ডিসপ্লের নিচে), অ্যাক্সিলেরোমিটার, জাইরো, প্রক্সিমিটি, কম্পাস",
		"GSM / HSPA / LTE / 5G": "জিএসএম / এইচএসপিএ / এলটিই / ৫G",
		"Glass front/back, plastic frame": "গ্লাস সামনে/পেছনে, প্লাস্টিক ফ্রেম",
		"Li-Ion (non-removable)": "লি-আয়ন (অপসারণযোগ্য নয়)",
		"Mali-G77 MC9": "মালি-G৭৭ MC৯",
		"MediaTek Dimensity 1100": "মিডিয়াটেক ডাইমেনসিটি ১১০০",
		"November 2022": "নভেম্বর ২০২২",
		"USB-C 2.0": "ইউএসবি-সি ২.০",
		"Wi-Fi 802.11 a/b/g/n/ac": "ওয়াই-ফাই ৮০২.১১ a/b/g/n/ac",
	}
}

// Seed inserts specification records for the product identified by slug 'vivo-x200-pro-mini'
func (s *SpecificationSeederMobileVivoX200ProMini) Seed(db *gorm.DB) error {
	productSlug := "vivo-x200-pro-mini"

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

	// Override model-specific values for Vivo X200 Pro Mini
	specs["Display Size"] = "6.44 inches"
	specs["Processor"] = "MediaTek Dimensity 1100"
	specs["Chipset"] = "Dimensity 1100 (6 nm)"
	specs["Cpu Type"] = "Octa-core"
	specs["Gpu Type"] = "Mali-G77 MC9"
	specs["Ram"] = "6 / 8 GB"
	specs["Storage"] = "128 GB"
	specs["Display Type"] = "AMOLED, 90 Hz"
	specs["Resolution"] = "2400 × 1080 px"
	specs["Refresh Rate"] = "90 Hz"
	specs["Build Material"] = "Glass front/back, plastic frame"
	specs["Weight"] = "185 g"
	specs["Network Technology"] = "GSM / HSPA / LTE / 5G"
	specs["Wifi Support"] = "Wi-Fi 802.11 a/b/g/n/ac"
	specs["Bluetooth Version"] = "5.2"
	specs["Nfc Support"] = "Yes"
	specs["Usb Type"] = "USB-C 2.0"
	specs["Rear Camera"] = "48 MP + 8 MP + 2 MP"
	specs["Camera Features"] = "OIS, LED flash"
	specs["Camera Video Resolution"] = "4K @ 30fps; 1080p @ 30fps"
	specs["Optical Zoom"] = "2x"
	specs["Front Camera"] = "32 MP"
	specs["Front Camera Video Resolution"] = "1080p @ 30fps"
	specs["Operating System"] = "Android 11, Funtouch OS 11"
	specs["Battery"] = "4000 mAh"
	specs["Battery Type"] = "Li-Ion (non-removable)"
	specs["Fast Charging"] = "33 W wired"
	specs["Wireless Charging"] = "No"
	specs["5G Support"] = "Yes"
	specs["Positioning System"] = "GPS, GLONASS, GALILEO, BDS"
	specs["Sensors"] = "Fingerprint (under display), Accelerometer, Gyro, Proximity, Compass"
	specs["Special Features"] = "AMOLED display"
	specs["Sim Card Type"] = "Dual SIM (Nano-SIM)"
	specs["Loudspeaker Quality"] = "Stereo"
	specs["Audio Quality"] = "Hi-Res Audio"
	specs["Audio Jack"] = "No"
	specs["Available Colors"] = "Black, Blue"
	specs["Announcement Date"] = "November 2022"
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
