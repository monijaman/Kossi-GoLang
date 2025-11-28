package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileSamsungGalaxyM335g seeds specifications/options for product 'samsung-galaxy-m33-5g'
type SpecificationSeederMobileSamsungGalaxyM335g struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSamsungGalaxyM335g creates a new seeder instance
func NewSpecificationSeederMobileSamsungGalaxyM335g() *SpecificationSeederMobileSamsungGalaxyM335g {
	return &SpecificationSeederMobileSamsungGalaxyM335g{BaseSeeder: BaseSeeder{name: "Specifications for Samsung Galaxy M33 5G"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSamsungGalaxyM335g) getBanglaTranslations() map[string]string {
	return map[string]string{
		"120 Hz": "১২০ Hz",
		"128 GB + microSD": "১২৮ জিবি + মাইক্রোএসডি",
		"165.4 × 76.9 × 9.4 mm": "১৬৫.৪ × ৭৬.৯ × ৯.৪ মিমি",
		"198 g": "১৯৮ g",
		"2408 × 1080 px": "২৪০৮ × ১০৮০ পিক্সেল",
		"25 W wired": "২৫ W তারযুক্ত",
		"3.5 mm": "৩.৫ মিমি",
		"4K @ 30fps; 1080p @ 30/60/720fps": "৪K @ ৩০fps; ১০৮০p @ ৩০/৬০/৭২০fps",
		"4K @ 30fps; 1080p @ 30fps": "৪K @ ৩০fps; ১০৮০p @ ৩০fps",
		"5.1": "৫.১",
		"50 MP + 5 MP + 2 MP + 2 MP": "৫০ মেগাপিক্সেল + ৫ মেগাপিক্সেল + ২ মেগাপিক্সেল + ২ মেগাপিক্সেল",
		"5000 mAh (some markets 6000 mAh)": "৫০০০ এমএএইচ (some markets ৬০০০ এমএএইচ)",
		"6 / 8 GB": "৬ / ৮ GB",
		"6.6 inches": "৬.৬ ইঞ্চি",
		"8 MP": "৮ মেগাপিক্সেল",
		"Android 12 (upgradable)": "অ্যান্ড্রয়েড ১২ (আপগ্রেডযোগ্য)",
		"April 2022": "এপ্রিল ২০২২",
		"Deep Ocean Blue, Mystique Green, Emerald Brown": "Deep ওশান নীল, Mystique সবুজ, Emerald বাদামী",
		"Exynos 1280": "এক্সিনস ১২৮০",
		"Exynos 1280 (5 nm)": "এক্সিনস ১২৮০ (৫ ন্যানোমিটার)",
		"GSM / HSPA / LTE / 5G": "জিএসএম / এইচএসপিএ / এলটিই / ৫G",
		"Glass front (Gorilla Glass 5), plastic frame / back": "গ্লাস সামনে (গরিলা গ্লাস ৫), প্লাস্টিক ফ্রেম / পেছনে",
		"Li-Ion (non-removable)": "লি-আয়ন (অপসারণযোগ্য নয়)",
		"Mali-G68 MP4": "মালি-G৬৮ মেগাপিক্সেল৪",
		"Nano-SIM + Nano-SIM": "ন্যানো-সিম + ন্যানো-সিম",
		"Octa-core (2×2.4 GHz + 6×2.0 GHz)": "অক্টা-কোর (২×২.৪ গিগাহার্টজ + ৬×২.০ গিগাহার্টজ)",
		"Side fingerprint, Accelerometer, Gyro, Proximity, Compass": "সাইড ফিঙ্গারপ্রিন্ট, অ্যাক্সিলেরোমিটার, জাইরো, প্রক্সিমিটি, কম্পাস",
		"TFT LCD, 120 Hz": "TFT এলসিডি, ১২০ Hz",
		"USB-C 2.0": "ইউএসবি-সি ২.০",
		"Wi-Fi 802.11 a/b/g/n/ac, dual-band": "ওয়াই-ফাই ৮০২.১১ a/b/g/n/ac, dual-bএবং",
	}
}

// Seed inserts specification records for the product identified by slug 'samsung-galaxy-m33-5g'
func (s *SpecificationSeederMobileSamsungGalaxyM335g) Seed(db *gorm.DB) error {
	productSlug := "samsung-galaxy-m33-5g"

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

	// Override model-specific values for Samsung Galaxy M33 5G
	specs["Display Size"] = "6.6 inches"
	specs["Processor"] = "Exynos 1280"
	specs["Chipset"] = "Exynos 1280 (5 nm)"
	specs["Cpu Type"] = "Octa-core (2×2.4 GHz + 6×2.0 GHz)"
	specs["Gpu Type"] = "Mali-G68 MP4"
	specs["Ram"] = "6 / 8 GB"
	specs["Storage"] = "128 GB + microSD"
	specs["Display Type"] = "TFT LCD, 120 Hz"
	specs["Resolution"] = "2408 × 1080 px"
	specs["Refresh Rate"] = "120 Hz"
	specs["Build Material"] = "Glass front (Gorilla Glass 5), plastic frame / back"
	specs["Weight"] = "198 g"
	specs["Dimensions"] = "165.4 × 76.9 × 9.4 mm"
	specs["Network Technology"] = "GSM / HSPA / LTE / 5G"
	specs["Wifi Support"] = "Wi-Fi 802.11 a/b/g/n/ac, dual-band"
	specs["Bluetooth Version"] = "5.1"
	specs["Usb Type"] = "USB-C 2.0"
	specs["Rear Camera"] = "50 MP + 5 MP + 2 MP + 2 MP"
	specs["Camera Features"] = "PDAF, LED flash, HDR"
	specs["Camera Video Resolution"] = "4K @ 30fps; 1080p @ 30/60/720fps"
	specs["Optical Zoom"] = "None"
	specs["Front Camera"] = "8 MP"
	specs["Front Camera Video Resolution"] = "4K @ 30fps; 1080p @ 30fps"
	specs["Operating System"] = "Android 12 (upgradable)"
	specs["Battery"] = "5000 mAh (some markets 6000 mAh)"
	specs["Battery Type"] = "Li-Ion (non-removable)"
	specs["Fast Charging"] = "25 W wired"
	specs["Wireless Charging"] = "No"
	specs["5G Support"] = "Yes"
	specs["Positioning System"] = "GPS, GLONASS, GALILEO, BDS"
	specs["Sensors"] = "Side fingerprint, Accelerometer, Gyro, Proximity, Compass"
	specs["Special Features"] = "PowerCool Tech"
	specs["Sim Card Type"] = "Nano-SIM + Nano-SIM"
	specs["Loudspeaker Quality"] = "Mono"
	specs["Audio Jack"] = "3.5 mm"
	specs["Available Colors"] = "Deep Ocean Blue, Mystique Green, Emerald Brown"
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
