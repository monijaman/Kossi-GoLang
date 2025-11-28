package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileSamsungGalaxyM535g seeds specifications/options for product 'samsung-galaxy-m53-5g'
type SpecificationSeederMobileSamsungGalaxyM535g struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSamsungGalaxyM535g creates a new seeder instance
func NewSpecificationSeederMobileSamsungGalaxyM535g() *SpecificationSeederMobileSamsungGalaxyM535g {
	return &SpecificationSeederMobileSamsungGalaxyM535g{BaseSeeder: BaseSeeder{name: "Specifications for Samsung Galaxy M53 5G"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSamsungGalaxyM535g) getBanglaTranslations() map[string]string {
	return map[string]string{
		"108 MP + 8 MP + 2 MP + 2 MP": "১০৮ মেগাপিক্সেল + ৮ মেগাপিক্সেল + ২ মেগাপিক্সেল + ২ মেগাপিক্সেল",
		"120 Hz": "১২০ Hz",
		"128 / 256 GB + microSD": "১২৮ / ২৫৬ জিবি + মাইক্রোএসডি",
		"164.7 × 77.0 × 7.4 mm": "১৬৪.৭ × ৭৭.০ × ৭.৪ মিমি",
		"176 g": "১৭৬ g",
		"2408 × 1080 px": "২৪০৮ × ১০৮০ পিক্সেল",
		"25 W wired": "২৫ W তারযুক্ত",
		"32 MP": "৩২ মেগাপিক্সেল",
		"4K @ 30fps; 1080p @ 30/60fps": "৪K @ ৩০fps; ১০৮০p @ ৩০/৬০fps",
		"4K @ 30fps; 1080p @ 30fps": "৪K @ ৩০fps; ১০৮০p @ ৩০fps",
		"5.2": "৫.২",
		"5000 mAh": "৫০০০ এমএএইচ",
		"6 / 8 GB": "৬ / ৮ GB",
		"6.7 inches": "৬.৭ ইঞ্চি",
		"Android 12, One UI 4.1 (or newer)": "অ্যান্ড্রয়েড ১২, ওয়ান ইউআই ৪.১ (বা newer)",
		"April 2022": "এপ্রিল ২০২২",
		"Dimensity 900 (6 nm)": "ডাইমেনসিটি ৯০০ (৬ ন্যানোমিটার)",
		"Dual SIM (Nano + Nano)": "ডুয়াল সিম (ন্যানো + ন্যানো)",
		"GSM / HSPA / LTE / 5G": "জিএসএম / এইচএসপিএ / এলটিই / ৫G",
		"Glass front, plastic frame / back": "গ্লাস সামনে, প্লাস্টিক ফ্রেম / পেছনে",
		"Green, Blue, Emerald Brown": "সবুজ, নীল, Emerald বাদামী",
		"Li-Ion (non-removable)": "লি-আয়ন (অপসারণযোগ্য নয়)",
		"Mali-G68 MC4": "মালি-G৬৮ MC৪",
		"MediaTek Dimensity 900": "মিডিয়াটেক ডাইমেনসিটি ৯০০",
		"Octa-core (2×2.4 GHz Cortex-A78 + 6×2.0 GHz Cortex-A55)": "অক্টা-কোর (২×২.৪ গিগাহার্টজ Cবাtex-A৭৮ + ৬×২.০ গিগাহার্টজ Cবাtex-A৫৫)",
		"Side fingerprint, Accelerometer, Gyro, Proximity, Compass": "সাইড ফিঙ্গারপ্রিন্ট, অ্যাক্সিলেরোমিটার, জাইরো, প্রক্সিমিটি, কম্পাস",
		"Super AMOLED Plus, 120 Hz": "Super অ্যামোলেড প্লাস, ১২০ Hz",
		"USB-C 2.0": "ইউএসবি-সি ২.০",
		"Wi-Fi 802.11 a/b/g/n/ac, dual-band": "ওয়াই-ফাই ৮০২.১১ a/b/g/n/ac, dual-bএবং",
	}
}

// Seed inserts specification records for the product identified by slug 'samsung-galaxy-m53-5g'
func (s *SpecificationSeederMobileSamsungGalaxyM535g) Seed(db *gorm.DB) error {
	productSlug := "samsung-galaxy-m53-5g"

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

	// Override model-specific values for Samsung Galaxy M53 5G
	specs["Display Size"] = "6.7 inches"
	specs["Processor"] = "MediaTek Dimensity 900"
	specs["Chipset"] = "Dimensity 900 (6 nm)"
	specs["Cpu Type"] = "Octa-core (2×2.4 GHz Cortex-A78 + 6×2.0 GHz Cortex-A55)"
	specs["Gpu Type"] = "Mali-G68 MC4"
	specs["Ram"] = "6 / 8 GB"
	specs["Storage"] = "128 / 256 GB + microSD"
	specs["Display Type"] = "Super AMOLED Plus, 120 Hz"
	specs["Resolution"] = "2408 × 1080 px"
	specs["Refresh Rate"] = "120 Hz"
	specs["Build Material"] = "Glass front, plastic frame / back"
	specs["Weight"] = "176 g"
	specs["Dimensions"] = "164.7 × 77.0 × 7.4 mm"
	specs["Network Technology"] = "GSM / HSPA / LTE / 5G"
	specs["Wifi Support"] = "Wi-Fi 802.11 a/b/g/n/ac, dual-band"
	specs["Bluetooth Version"] = "5.2"
	specs["Usb Type"] = "USB-C 2.0"
	specs["Rear Camera"] = "108 MP + 8 MP + 2 MP + 2 MP"
	specs["Camera Features"] = "PDAF, LED flash, HDR"
	specs["Camera Video Resolution"] = "4K @ 30fps; 1080p @ 30/60fps"
	specs["Optical Zoom"] = "None"
	specs["Front Camera"] = "32 MP"
	specs["Front Camera Video Resolution"] = "4K @ 30fps; 1080p @ 30fps"
	specs["Operating System"] = "Android 12, One UI 4.1 (or newer)"
	specs["Battery"] = "5000 mAh"
	specs["Battery Type"] = "Li-Ion (non-removable)"
	specs["Fast Charging"] = "25 W wired"
	specs["Wireless Charging"] = "No"
	specs["5G Support"] = "Yes"
	specs["Positioning System"] = "GPS, GLONASS, GALILEO, BDS"
	specs["Sensors"] = "Side fingerprint, Accelerometer, Gyro, Proximity, Compass"
	specs["Sim Card Type"] = "Dual SIM (Nano + Nano)"
	specs["Loudspeaker Quality"] = "Mono"
	specs["Audio Jack"] = "No"
	specs["Available Colors"] = "Green, Blue, Emerald Brown"
	specs["Announcement Date"] = "April 2022"
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
