package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileSamsungGalaxyA23 seeds specifications/options for product 'samsung-galaxy-a23'
type SpecificationSeederMobileSamsungGalaxyA23 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSamsungGalaxyA23 creates a new seeder instance
func NewSpecificationSeederMobileSamsungGalaxyA23() *SpecificationSeederMobileSamsungGalaxyA23 {
	return &SpecificationSeederMobileSamsungGalaxyA23{BaseSeeder: BaseSeeder{name: "Specifications for Samsung Galaxy A23"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSamsungGalaxyA23) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080p @ 30fps": "১০৮০p @ ৩০fps",
		"165.4 × 76.9 × 8.4 mm": "১৬৫.৪ × ৭৬.৯ × ৮.৪ মিমি",
		"195 g": "১৯৫ g",
		"2408 × 1080 px (~400 ppi)": "২৪০৮ × ১০৮০ px (~৪০০ ppi)",
		"25 W wired": "২৫ W তারযুক্ত",
		"3.5 mm headphone jack": "৩.৫ মিমি headphone jack",
		"4 / 6 / 8 GB": "৪ / ৬ / ৮ GB",
		"5.0": "৫.০",
		"50 MP + 5 MP + 2 MP + 2 MP": "৫০ MP + ৫ MP + ২ MP + ২ MP",
		"5000 mAh": "৫০০০ এমএএইচ",
		"6.6 inches": "৬.৬ ইঞ্চি",
		"64 / 128 GB + microSD": "৬৪ / ১২৮ GB + microSD",
		"8 MP": "৮ MP",
		"90 Hz": "৯০ Hz",
		"Adreno 610": "Adreno ৬১০",
		"Android 12, One UI 4.1 (upgradable)": "Android ১২, One UI ৪.১ (upgradable)",
		"Black, White, Peach, Blue": "কালো, সাদা, Peach, নীল",
		"GSM / HSPA / LTE (4G)": "GSM / HSPA / LTE (৪G)",
		"Glass front (Gorilla Glass 5), plastic frame/back": "গ্লাস সামনে (Gorilla Glass ৫), plastic frame/back",
		"Li-Ion (non-removable)": "লি-আয়ন (অপসারণযোগ্য নয়)",
		"March 2022": "March ২০২২",
		"Octa-core (4×2.4 GHz + 4×1.9 GHz)": "Octa-core (৪×২.৪ GHz + ৪×১.৯ GHz)",
		"PLS LCD, 90 Hz": "PLS LCD, ৯০ Hz",
		"Side fingerprint, Accelerometer, Gyro, Proximity, Compass": "Side ফিঙ্গারপ্রিন্ট, অ্যাক্সিলেরোমিটার, জাইরো, প্রক্সিমিটি, কম্পাস",
		"Snapdragon 680 (4G)": "Snapdragon ৬৮০ (৪G)",
		"Snapdragon 680 (6 nm)": "Snapdragon ৬৮০ (৬ nm)",
		"Some variants (A23 5G), base A23 is 4G": "Some variants (A২৩ ৫G), base A২৩ is ৪G",
		"USB-C 2.0": "USB-C ২.০",
		"Wi-Fi 802.11 a/b/g/n/ac": "Wi-Fi ৮০২.১১ a/b/g/n/ac",
		"Yes (market dependent)": "হ্যাঁ (market dependent)",
	}
}

// Seed inserts specification records for the product identified by slug 'samsung-galaxy-a23'
func (s *SpecificationSeederMobileSamsungGalaxyA23) Seed(db *gorm.DB) error {
	productSlug := "samsung-galaxy-a23"

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

	// Override model-specific values for Samsung Galaxy A23
	specs["Display Size"] = "6.6 inches"
	specs["Processor"] = "Snapdragon 680 (4G)"
	specs["Chipset"] = "Snapdragon 680 (6 nm)"
	specs["Cpu Type"] = "Octa-core (4×2.4 GHz + 4×1.9 GHz)"
	specs["Gpu Type"] = "Adreno 610"
	specs["Ram"] = "4 / 6 / 8 GB"
	specs["Storage"] = "64 / 128 GB + microSD"
	specs["Display Type"] = "PLS LCD, 90 Hz"
	specs["Resolution"] = "2408 × 1080 px (~400 ppi)"
	specs["Refresh Rate"] = "90 Hz"
	specs["Build Material"] = "Glass front (Gorilla Glass 5), plastic frame/back"
	specs["Weight"] = "195 g"
	specs["Dimensions"] = "165.4 × 76.9 × 8.4 mm"
	specs["Water Resistance"] = "None"
	specs["Network Technology"] = "GSM / HSPA / LTE (4G)"
	specs["Wifi Support"] = "Wi-Fi 802.11 a/b/g/n/ac"
	specs["Bluetooth Version"] = "5.0"
	specs["Nfc Support"] = "Yes (market dependent)"
	specs["Usb Type"] = "USB-C 2.0"
	specs["Rear Camera"] = "50 MP + 5 MP + 2 MP + 2 MP"
	specs["Camera Features"] = "OIS (main), LED flash"
	specs["Camera Video Resolution"] = "1080p @ 30fps"
	specs["Optical Zoom"] = "None"
	specs["Front Camera"] = "8 MP"
	specs["Front Camera Video Resolution"] = "1080p @ 30fps"
	specs["Operating System"] = "Android 12, One UI 4.1 (upgradable)"
	specs["Battery"] = "5000 mAh"
	specs["Battery Type"] = "Li-Ion (non-removable)"
	specs["Fast Charging"] = "25 W wired"
	specs["Wireless Charging"] = "No"
	specs["5G Support"] = "Some variants (A23 5G), base A23 is 4G"
	specs["Positioning System"] = "GPS, GLONASS, GALILEO, BDS"
	specs["Sensors"] = "Side fingerprint, Accelerometer, Gyro, Proximity, Compass"
	specs["Special Features"] = "Dolby Atmos"
	specs["Sim Card Type"] = "Nano-SIM or Dual SIM"
	specs["Loudspeaker Quality"] = "Mono / Stereo depending on variant"
	specs["Audio Quality"] = "Not specified"
	specs["Audio Jack"] = "3.5 mm headphone jack"
	specs["Available Colors"] = "Black, White, Peach, Blue"
	specs["Announcement Date"] = "March 2022"
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
