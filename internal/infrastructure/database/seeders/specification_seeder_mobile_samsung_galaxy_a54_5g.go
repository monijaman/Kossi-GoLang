package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileSamsungGalaxyA545g seeds specifications/options for product 'samsung-galaxy-a54-5g'
type SpecificationSeederMobileSamsungGalaxyA545g struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSamsungGalaxyA545g creates a new seeder instance
func NewSpecificationSeederMobileSamsungGalaxyA545g() *SpecificationSeederMobileSamsungGalaxyA545g {
	return &SpecificationSeederMobileSamsungGalaxyA545g{BaseSeeder: BaseSeeder{name: "Specifications for Samsung Galaxy A54 5G"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSamsungGalaxyA545g) getBanglaTranslations() map[string]string {
	return map[string]string{
		"120 Hz": "১২০ Hz",
		"128 / 256 GB + microSD": "১২৮ / ২৫৬ GB + microSD",
		"158.2 × 76.7 × 8.2 mm": "১৫৮.২ × ৭৬.৭ × ৮.২ মিমি",
		"202 g": "২০২ g",
		"2340 × 1080 px (19.5:9)": "২৩৪০ × ১০৮০ px (১৯.৫:৯)",
		"25 W wired": "২৫ W তারযুক্ত",
		"32 MP": "৩২ MP",
		"4 OS upgrades, 5 years of security updates": "৪ OS upgrades, ৫ years of security updates",
		"4K @ 30fps; 1080p @ 30/60fps": "৪K @ ৩০fps; ১০৮০p @ ৩০/৬০fps",
		"4K @ 30fps; 1080p @ 30fps": "৪K @ ৩০fps; ১০৮০p @ ৩০fps",
		"5.2": "৫.২",
		"50 MP + 8 MP + 5 MP": "৫০ MP + ৮ MP + ৫ MP",
		"5000 mAh": "৫০০০ এমএএইচ",
		"6 / 8 GB": "৬ / ৮ GB",
		"6.4 inches": "৬.৪ ইঞ্চি",
		"Android 13, One UI 5.1": "Android ১৩, One UI ৫.১",
		"Exynos 1380": "Exynos ১৩৮০",
		"Exynos 1380 (5 nm)": "Exynos ১৩৮০ (৫ nm)",
		"Fingerprint (in-display, optical), Accelerometer, Gyro, Compass, Barometer": "ফিঙ্গারপ্রিন্ট (in-display, optical), অ্যাক্সিলেরোমিটার, জাইরো, কম্পাস, ব্যারোমিটার",
		"GSM / HSPA / LTE / 5G": "GSM / HSPA / LTE / ৫G",
		"Glass front (Gorilla Glass 5), plastic frame / back": "গ্লাস সামনে (Gorilla Glass ৫), plastic frame / back",
		"IP67 (up to 1 m for 30 min)": "IP৬৭ (up to ১ m for ৩০ min)",
		"Li-Ion (non-removable)": "লি-আয়ন (অপসারণযোগ্য নয়)",
		"Lime, Graphite, Violet, White": "Lime, Graphite, Violet, সাদা",
		"Mali-G68 (MP5)": "Mali-G৬৮ (MP৫)",
		"March 2023": "March ২০২৩",
		"Nano-SIM or Hybrid Dual SIM (market dependent)": "ন্যানো-সিম or Hybrid ডুয়াল সিম (market dependent)",
		"Octa-core (4×2.4 GHz Cortex-A78 + 4×2.0 GHz Cortex-A55)": "Octa-core (৪×২.৪ GHz Cortex-A৭৮ + ৪×২.০ GHz Cortex-A৫৫)",
		"Super AMOLED, 120 Hz": "Super AMOLED, ১২০ Hz",
		"USB-C 2.0": "USB-C ২.০",
		"Wi-Fi 802.11 a/b/g/n/ac, dual-band": "Wi-Fi ৮০২.১১ a/b/g/n/ac, dual-band",
		"Yes (region dependent)": "হ্যাঁ (region dependent)",
	}
}

// Seed inserts specification records for the product identified by slug 'samsung-galaxy-a54-5g'
func (s *SpecificationSeederMobileSamsungGalaxyA545g) Seed(db *gorm.DB) error {
	productSlug := "samsung-galaxy-a54-5g"

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

	// Override model-specific values for Samsung Galaxy A54 5G
	specs["Display Size"] = "6.4 inches"
	specs["Processor"] = "Exynos 1380"
	specs["Chipset"] = "Exynos 1380 (5 nm)"
	specs["Cpu Type"] = "Octa-core (4×2.4 GHz Cortex-A78 + 4×2.0 GHz Cortex-A55)"
	specs["Gpu Type"] = "Mali-G68 (MP5)"
	specs["Ram"] = "6 / 8 GB"
	specs["Storage"] = "128 / 256 GB + microSD"
	specs["Display Type"] = "Super AMOLED, 120 Hz"
	specs["Resolution"] = "2340 × 1080 px (19.5:9)"
	specs["Refresh Rate"] = "120 Hz"
	specs["Build Material"] = "Glass front (Gorilla Glass 5), plastic frame / back"
	specs["Weight"] = "202 g"
	specs["Dimensions"] = "158.2 × 76.7 × 8.2 mm"
	specs["Water Resistance"] = "IP67 (up to 1 m for 30 min)"
	specs["Network Technology"] = "GSM / HSPA / LTE / 5G"
	specs["Wifi Support"] = "Wi-Fi 802.11 a/b/g/n/ac, dual-band"
	specs["Bluetooth Version"] = "5.2"
	specs["Nfc Support"] = "Yes (region dependent)"
	specs["Usb Type"] = "USB-C 2.0"
	specs["Rear Camera"] = "50 MP + 8 MP + 5 MP"
	specs["Camera Features"] = "OIS (main), LED flash"
	specs["Camera Video Resolution"] = "4K @ 30fps; 1080p @ 30/60fps"
	specs["Optical Zoom"] = "None"
	specs["Front Camera"] = "32 MP"
	specs["Front Camera Video Resolution"] = "4K @ 30fps; 1080p @ 30fps"
	specs["Operating System"] = "Android 13, One UI 5.1"
	specs["Battery"] = "5000 mAh"
	specs["Battery Type"] = "Li-Ion (non-removable)"
	specs["Fast Charging"] = "25 W wired"
	specs["Wireless Charging"] = "No"
	specs["5G Support"] = "Yes"
	specs["Positioning System"] = "GPS, GLONASS, GALILEO, BDS"
	specs["Sensors"] = "Fingerprint (in-display, optical), Accelerometer, Gyro, Compass, Barometer"
	specs["Special Features"] = "4 OS upgrades, 5 years of security updates"
	specs["Sim Card Type"] = "Nano-SIM or Hybrid Dual SIM (market dependent)"
	specs["Loudspeaker Quality"] = "Stereo"
	specs["Audio Quality"] = "Not specified"
	specs["Audio Jack"] = "No"
	specs["Available Colors"] = "Lime, Graphite, Violet, White"
	specs["Announcement Date"] = "March 2023"
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
