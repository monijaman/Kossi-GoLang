package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileVivoV40Lite seeds specifications/options for product 'vivo-v40-lite'
type SpecificationSeederMobileVivoV40Lite struct {
	BaseSeeder
}

// NewSpecificationSeederMobileVivoV40Lite creates a new seeder instance
func NewSpecificationSeederMobileVivoV40Lite() *SpecificationSeederMobileVivoV40Lite {
	return &SpecificationSeederMobileVivoV40Lite{BaseSeeder: BaseSeeder{name: "Specifications for Vivo V40 Lite"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileVivoV40Lite) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080p @ 30fps": "১০৮০p @ ৩০fps",
		"128 GB": "১২৮ GB",
		"16 MP": "১৬ মেগাপিক্সেল",
		"176 g": "১৭৬ g",
		"18 W wired": "১৮ W তারযুক্ত",
		"2400 × 1080 px": "২৪০০ × ১০৮০ পিক্সেল",
		"4100 mAh": "৪১০০ এমএএইচ",
		"48 MP + 2 MP": "৪৮ মেগাপিক্সেল + ২ মেগাপিক্সেল",
		"5.1": "৫.১",
		"6 / 8 GB": "৬ / ৮ GB",
		"6.38 inches": "৬.৩৮ ইঞ্চি",
		"60 Hz": "৬০ Hz",
		"AMOLED, 60 Hz": "অ্যামোলেড, ৬০ Hz",
		"Android 12, Funtouch OS 12": "অ্যান্ড্রয়েড ১২, Funথেকেuch OS ১২",
		"Black, Blue, Pink": "কালো, নীল, গোলাপী",
		"Dimensity 810 (6 nm)": "ডাইমেনসিটি ৮১০ (৬ ন্যানোমিটার)",
		"Dual SIM (Nano-SIM)": "ডুয়াল সিম (ন্যানো-সিম)",
		"Fingerprint (side-mounted), Accelerometer, Gyro, Proximity, Compass": "ফিঙ্গারপ্রিন্ট (পাশে মাউন্ট), অ্যাক্সিলেরোমিটার, জাইরো, প্রক্সিমিটি, কম্পাস",
		"GSM / HSPA / LTE / 5G": "জিএসএম / এইচএসপিএ / এলটিই / ৫G",
		"Glass front, plastic back/frame": "গ্লাস সামনে, প্লাস্টিক পেছনে/ফ্রেম",
		"July 2022": "জুলাই ২০২২",
		"Li-Ion (non-removable)": "লি-আয়ন (অপসারণযোগ্য নয়)",
		"Mali-G57 MC2": "মালি-G৫৭ MC২",
		"MediaTek Dimensity 810": "মিডিয়াটেক ডাইমেনসিটি ৮১০",
		"USB-C 2.0": "ইউএসবি-সি ২.০",
		"Wi-Fi 802.11 a/b/g/n/ac": "ওয়াই-ফাই ৮০২.১১ a/b/g/n/ac",
		"Yes, 3.5 mm": "হ্যাঁ, ৩.৫ মিমি",
	}
}

// Seed inserts specification records for the product identified by slug 'vivo-v40-lite'
func (s *SpecificationSeederMobileVivoV40Lite) Seed(db *gorm.DB) error {
	productSlug := "vivo-v40-lite"

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

	// Override model-specific values for Vivo V40 Lite
	specs["Display Size"] = "6.38 inches"
	specs["Processor"] = "MediaTek Dimensity 810"
	specs["Chipset"] = "Dimensity 810 (6 nm)"
	specs["Cpu Type"] = "Octa-core"
	specs["Gpu Type"] = "Mali-G57 MC2"
	specs["Ram"] = "6 / 8 GB"
	specs["Storage"] = "128 GB"
	specs["Display Type"] = "AMOLED, 60 Hz"
	specs["Resolution"] = "2400 × 1080 px"
	specs["Refresh Rate"] = "60 Hz"
	specs["Build Material"] = "Glass front, plastic back/frame"
	specs["Weight"] = "176 g"
	specs["Network Technology"] = "GSM / HSPA / LTE / 5G"
	specs["Wifi Support"] = "Wi-Fi 802.11 a/b/g/n/ac"
	specs["Bluetooth Version"] = "5.1"
	specs["Nfc Support"] = "Yes"
	specs["Usb Type"] = "USB-C 2.0"
	specs["Rear Camera"] = "48 MP + 2 MP"
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
	specs["Sensors"] = "Fingerprint (side-mounted), Accelerometer, Gyro, Proximity, Compass"
	specs["Special Features"] = "AMOLED display"
	specs["Sim Card Type"] = "Dual SIM (Nano-SIM)"
	specs["Loudspeaker Quality"] = "Mono"
	specs["Audio Quality"] = "Standard"
	specs["Audio Jack"] = "Yes, 3.5 mm"
	specs["Available Colors"] = "Black, Blue, Pink"
	specs["Announcement Date"] = "July 2022"
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
