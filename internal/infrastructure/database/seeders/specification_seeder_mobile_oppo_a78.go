package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileOppoA78 seeds specifications/options for product 'oppo-a78'
type SpecificationSeederMobileOppoA78 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileOppoA78 creates a new seeder instance
func NewSpecificationSeederMobileOppoA78() *SpecificationSeederMobileOppoA78 {
	return &SpecificationSeederMobileOppoA78{BaseSeeder: BaseSeeder{name: "Specifications for Oppo A78"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileOppoA78) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080p @ 30fps": "১০৮০p @ ৩০fps",
		"120 Hz": "১২০ Hz",
		"128 / 256 GB + microSD": "১২৮ / ২৫৬ জিবি + মাইক্রোএসডি",
		"16 MP": "১৬ মেগাপিক্সেল",
		"186 g": "১৮৬ g",
		"2023": "২০২৩",
		"2412 × 1080 px (~401 ppi)": "২৪১২ × ১০৮০ পিক্সেল (~৪০১ পিপিআই)",
		"3.5 mm headphone jack": "৩.৫ মিমি হেডফোন জ্যাক",
		"33 W wired": "৩৩ W তারযুক্ত",
		"4K @ 30fps": "৪K @ ৩০fps",
		"5.2": "৫.২",
		"5000 mAh": "৫০০০ এমএএইচ",
		"6 / 8 GB": "৬ / ৮ GB",
		"6.56 inches": "৬.৫৬ ইঞ্চি",
		"64 MP + 2 MP": "৬৪ মেগাপিক্সেল + ২ মেগাপিক্সেল",
		"Android 13, ColorOS 13": "অ্যান্ড্রয়েড ১৩, কালারওএস ১৩",
		"Black, Blue, Green": "কালো, নীল, সবুজ",
		"Dimensity 930 (6 nm)": "ডাইমেনসিটি ৯৩০ (৬ ন্যানোমিটার)",
		"Dual SIM (Nano-SIM)": "ডুয়াল সিম (ন্যানো-সিম)",
		"Fingerprint (side), Accelerometer, Gyro, Proximity, Compass": "ফিঙ্গারপ্রিন্ট (side), অ্যাক্সিলেরোমিটার, জাইরো, প্রক্সিমিটি, কম্পাস",
		"GSM / HSPA / LTE / 5G": "জিএসএম / এইচএসপিএ / এলটিই / ৫G",
		"Glass front, plastic frame/back": "গ্লাস সামনে, প্লাস্টিক ফ্রেম/পেছনে",
		"Li-Ion (non-removable)": "লি-আয়ন (অপসারণযোগ্য নয়)",
		"Mali-G68 MC4": "মালি-G৬৮ MC৪",
		"MediaTek Dimensity 930": "মিডিয়াটেক ডাইমেনসিটি ৯৩০",
		"Octa-core (2×2.2 GHz Cortex-A78 + 6×2.0 GHz Cortex-A55)": "অক্টা-কোর (২×২.২ গিগাহার্টজ Cবাtex-A৭৮ + ৬×২.০ গিগাহার্টজ Cবাtex-A৫৫)",
		"USB-C 2.0": "ইউএসবি-সি ২.০",
		"Wi-Fi 802.11 a/b/g/n/ac": "ওয়াই-ফাই ৮০২.১১ a/b/g/n/ac",
	}
}

// Seed inserts specification records for the product identified by slug 'oppo-a78'
func (s *SpecificationSeederMobileOppoA78) Seed(db *gorm.DB) error {
	productSlug := "oppo-a78"

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

	// Override model-specific values for Oppo A78
	specs["Display Size"] = "6.56 inches"
	specs["Processor"] = "MediaTek Dimensity 930"
	specs["Chipset"] = "Dimensity 930 (6 nm)"
	specs["Cpu Type"] = "Octa-core (2×2.2 GHz Cortex-A78 + 6×2.0 GHz Cortex-A55)"
	specs["Gpu Type"] = "Mali-G68 MC4"
	specs["Ram"] = "6 / 8 GB"
	specs["Storage"] = "128 / 256 GB + microSD"
	specs["Display Type"] = "AMOLED"
	specs["Resolution"] = "2412 × 1080 px (~401 ppi)"
	specs["Refresh Rate"] = "120 Hz"
	specs["Build Material"] = "Glass front, plastic frame/back"
	specs["Weight"] = "186 g"
	specs["Network Technology"] = "GSM / HSPA / LTE / 5G"
	specs["Wifi Support"] = "Wi-Fi 802.11 a/b/g/n/ac"
	specs["Bluetooth Version"] = "5.2"
	specs["Nfc Support"] = "Yes"
	specs["Usb Type"] = "USB-C 2.0"
	specs["Rear Camera"] = "64 MP + 2 MP"
	specs["Camera Features"] = "LED flash, HDR, OIS"
	specs["Camera Video Resolution"] = "4K @ 30fps"
	specs["Optical Zoom"] = "None"
	specs["Front Camera"] = "16 MP"
	specs["Front Camera Video Resolution"] = "1080p @ 30fps"
	specs["Operating System"] = "Android 13, ColorOS 13"
	specs["Battery"] = "5000 mAh"
	specs["Battery Type"] = "Li-Ion (non-removable)"
	specs["Fast Charging"] = "33 W wired"
	specs["Wireless Charging"] = "No"
	specs["5G Support"] = "Yes"
	specs["Positioning System"] = "GPS, GLONASS, GALILEO, BDS"
	specs["Sensors"] = "Fingerprint (side), Accelerometer, Gyro, Proximity, Compass"
	specs["Sim Card Type"] = "Dual SIM (Nano-SIM)"
	specs["Loudspeaker Quality"] = "Stereo"
	specs["Audio Jack"] = "3.5 mm headphone jack"
	specs["Available Colors"] = "Black, Blue, Green"
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
