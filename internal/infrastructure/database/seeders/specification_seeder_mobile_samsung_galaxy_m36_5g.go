package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileSamsungGalaxyM365g seeds specifications/options for product 'samsung-galaxy-m36-5g'
type SpecificationSeederMobileSamsungGalaxyM365g struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSamsungGalaxyM365g creates a new seeder instance
func NewSpecificationSeederMobileSamsungGalaxyM365g() *SpecificationSeederMobileSamsungGalaxyM365g {
	return &SpecificationSeederMobileSamsungGalaxyM365g{BaseSeeder: BaseSeeder{name: "Specifications for Samsung Galaxy M36 5G"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSamsungGalaxyM365g) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080 × 2340 px": "১০৮০ × ২৩৪০ px",
		"120 Hz": "১২০ Hz",
		"128 / 256 GB + microSD": "১২৮ / ২৫৬ GB + microSD",
		"13 MP": "১৩ MP",
		"162.3 × 78.6 × 9.1 mm": "১৬২.৩ × ৭৮.৬ × ৯.১ মিমি",
		"222 g": "২২২ g",
		"25 W wired": "২৫ W তারযুক্ত",
		"4K @ 30fps; 1080p @ 30/60fps": "৪K @ ৩০fps; ১০৮০p @ ৩০/৬০fps",
		"4K @ 30fps; 1080p @ 30fps": "৪K @ ৩০fps; ১০৮০p @ ৩০fps",
		"5.3": "৫.৩",
		"50 MP + 8 MP + 2 MP": "৫০ MP + ৮ MP + ২ MP",
		"5000 mAh": "৫০০০ এমএএইচ",
		"6 / 8 GB": "৬ / ৮ GB",
		"6 years of Android / Security updates": "৬ years of Android / Security updates",
		"6.7 inches": "৬.৭ ইঞ্চি",
		"Android 15 (One UI 7)": "Android ১৫ (One UI ৭)",
		"Exynos 1380": "Exynos ১৩৮০",
		"Exynos 1380 (5 nm)": "Exynos ১৩৮০ (৫ nm)",
		"GSM / HSPA / LTE / 5G": "GSM / HSPA / LTE / ৫G",
		"Glass front (Gorilla Glass Victus+), plastic frame, glass back": "গ্লাস সামনে (Gorilla Glass Victus+), plastic frame, গ্লাস পেছনে",
		"July 2025": "July ২০২৫",
		"Li-Po (non-removable)": "লি-পো (অপসারণযোগ্য নয়)",
		"Mali-G68 MP5": "Mali-G৬৮ MP৫",
		"Nano-SIM + Nano-SIM": "ন্যানো-সিম + ন্যানো-সিম",
		"Octa-core (4×2.4 GHz Cortex-A78 + 4×2.0 GHz Cortex-A55)": "Octa-core (৪×২.৪ GHz Cortex-A৭৮ + ৪×২.০ GHz Cortex-A৫৫)",
		"Side fingerprint, Accelerometer, Gyro, Proximity, Compass": "Side ফিঙ্গারপ্রিন্ট, অ্যাক্সিলেরোমিটার, জাইরো, প্রক্সিমিটি, কম্পাস",
		"Super AMOLED, 120 Hz": "Super AMOLED, ১২০ Hz",
		"Thunder Grey, DayBreak Blue, Moonlight Blue": "Thunder ধূসর, DayBreak নীল, Moonlight নীল",
		"USB-C 2.0": "USB-C ২.০",
		"Wi-Fi 802.11 a/b/g/n/ac": "Wi-Fi ৮০২.১১ a/b/g/n/ac",
	}
}

// Seed inserts specification records for the product identified by slug 'samsung-galaxy-m36-5g'
func (s *SpecificationSeederMobileSamsungGalaxyM365g) Seed(db *gorm.DB) error {
	productSlug := "samsung-galaxy-m36-5g"

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

	// Override model-specific values for Samsung Galaxy M36 5G
	specs["Display Size"] = "6.7 inches"
	specs["Processor"] = "Exynos 1380"
	specs["Chipset"] = "Exynos 1380 (5 nm)"
	specs["Cpu Type"] = "Octa-core (4×2.4 GHz Cortex-A78 + 4×2.0 GHz Cortex-A55)"
	specs["Gpu Type"] = "Mali-G68 MP5"
	specs["Ram"] = "6 / 8 GB"
	specs["Storage"] = "128 / 256 GB + microSD"
	specs["Display Type"] = "Super AMOLED, 120 Hz"
	specs["Resolution"] = "1080 × 2340 px"
	specs["Refresh Rate"] = "120 Hz"
	specs["Build Material"] = "Glass front (Gorilla Glass Victus+), plastic frame, glass back"
	specs["Weight"] = "222 g"
	specs["Dimensions"] = "162.3 × 78.6 × 9.1 mm"
	specs["Network Technology"] = "GSM / HSPA / LTE / 5G"
	specs["Wifi Support"] = "Wi-Fi 802.11 a/b/g/n/ac"
	specs["Bluetooth Version"] = "5.3"
	specs["Nfc Support"] = "Yes"
	specs["Usb Type"] = "USB-C 2.0"
	specs["Rear Camera"] = "50 MP + 8 MP + 2 MP"
	specs["Camera Features"] = "OIS, LED flash"
	specs["Camera Video Resolution"] = "4K @ 30fps; 1080p @ 30/60fps"
	specs["Optical Zoom"] = "None"
	specs["Front Camera"] = "13 MP"
	specs["Front Camera Video Resolution"] = "4K @ 30fps; 1080p @ 30fps"
	specs["Operating System"] = "Android 15 (One UI 7)"
	specs["Battery"] = "5000 mAh"
	specs["Battery Type"] = "Li-Po (non-removable)"
	specs["Fast Charging"] = "25 W wired"
	specs["Wireless Charging"] = "No"
	specs["5G Support"] = "Yes"
	specs["Sensors"] = "Side fingerprint, Accelerometer, Gyro, Proximity, Compass"
	specs["Special Features"] = "6 years of Android / Security updates"
	specs["Sim Card Type"] = "Nano-SIM + Nano-SIM"
	specs["Loudspeaker Quality"] = "Stereo"
	specs["Audio Jack"] = "No"
	specs["Available Colors"] = "Thunder Grey, DayBreak Blue, Moonlight Blue"
	specs["Announcement Date"] = "July 2025"
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
