package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileSamsungGalaxyA24 seeds specifications/options for product 'samsung-galaxy-a24'
type SpecificationSeederMobileSamsungGalaxyA24 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSamsungGalaxyA24 creates a new seeder instance
func NewSpecificationSeederMobileSamsungGalaxyA24() *SpecificationSeederMobileSamsungGalaxyA24 {
	return &SpecificationSeederMobileSamsungGalaxyA24{BaseSeeder: BaseSeeder{name: "Specifications for Samsung Galaxy A24"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSamsungGalaxyA24) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080p @ 30fps": "১০৮০p @ ৩০fps",
		"128 GB + microSD": "১২৮ GB + microSD",
		"13 MP": "১৩ MP",
		"165.4 × 76.9 × 8.4 mm": "১৬৫.৪ × ৭৬.৯ × ৮.৪ মিমি",
		"195 g": "১৯৫ g",
		"2408 × 1080 pixels (~406 ppi)": "২৪০৮ × ১০৮০ pixels (~৪০৬ ppi)",
		"25 W wired": "২৫ W তারযুক্ত",
		"4 / 6 / 8 GB": "৪ / ৬ / ৮ GB",
		"5.3": "৫.৩",
		"50 MP + 2 MP + 2 MP": "৫০ MP + ২ MP + ২ MP",
		"5000 mAh": "৫০০০ এমএএইচ",
		"6.5 inches": "৬.৫ ইঞ্চি",
		"90 Hz": "৯০ Hz",
		"Android 13, One UI 5.1": "Android ১৩, One UI ৫.১",
		"April 2023": "April ২০২৩",
		"Black, Dark Red, Light Green, Silver": "কালো, Dark লাল, Light সবুজ, রূপালী",
		"Eye Care Certification (low blue light)": "Eye Care Certification (low নীল light)",
		"GSM / HSPA / LTE (no 5G)": "GSM / HSPA / LTE (no ৫G)",
		"Glass front, plastic frame / back": "গ্লাস সামনে, plastic frame / back",
		"Helio G99 (6 nm)": "Helio G৯৯ (৬ nm)",
		"Li-Ion (non-removable)": "লি-আয়ন (অপসারণযোগ্য নয়)",
		"Mali-G57 MC2": "Mali-G৫৭ MC২",
		"MediaTek Helio G99": "MediaTek Helio G৯৯",
		"Octa-core (2×2.2 GHz Cortex-A76 + 6×2.0 GHz Cortex-A55)": "Octa-core (২×২.২ GHz Cortex-A৭৬ + ৬×২.০ GHz Cortex-A৫৫)",
		"Side-mounted fingerprint, Accelerometer, Gyro, Proximity, Compass": "পাশে মাউন্ট ফিঙ্গারপ্রিন্ট, অ্যাক্সিলেরোমিটার, জাইরো, প্রক্সিমিটি, কম্পাস",
		"USB-C 2.0": "USB-C ২.০",
		"Wi-Fi 802.11 a/b/g/n/ac": "Wi-Fi ৮০২.১১ a/b/g/n/ac",
		"Yes (market dependent)": "হ্যাঁ (market dependent)",
		"Yes, 3.5 mm": "হ্যাঁ, ৩.৫ মিমি",
	}
}

// Seed inserts specification records for the product identified by slug 'samsung-galaxy-a24'
func (s *SpecificationSeederMobileSamsungGalaxyA24) Seed(db *gorm.DB) error {
	productSlug := "samsung-galaxy-a24"

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

	// Override model-specific values for Samsung Galaxy A24
	specs["Display Size"] = "6.5 inches"
	specs["Processor"] = "MediaTek Helio G99"
	specs["Chipset"] = "Helio G99 (6 nm)"
	specs["Cpu Type"] = "Octa-core (2×2.2 GHz Cortex-A76 + 6×2.0 GHz Cortex-A55)"
	specs["Gpu Type"] = "Mali-G57 MC2"
	specs["Ram"] = "4 / 6 / 8 GB"
	specs["Storage"] = "128 GB + microSD"
	specs["Display Type"] = "Super AMOLED"
	specs["Resolution"] = "2408 × 1080 pixels (~406 ppi)"
	specs["Refresh Rate"] = "90 Hz"
	specs["Build Material"] = "Glass front, plastic frame / back"
	specs["Weight"] = "195 g"
	specs["Dimensions"] = "165.4 × 76.9 × 8.4 mm"
	specs["Water Resistance"] = "None"
	specs["Network Technology"] = "GSM / HSPA / LTE (no 5G)"
	specs["Wifi Support"] = "Wi-Fi 802.11 a/b/g/n/ac"
	specs["Bluetooth Version"] = "5.3"
	specs["Nfc Support"] = "Yes (market dependent)"
	specs["Usb Type"] = "USB-C 2.0"
	specs["Rear Camera"] = "50 MP + 2 MP + 2 MP"
	specs["Camera Features"] = "OIS (main)"
	specs["Camera Video Resolution"] = "1080p @ 30fps"
	specs["Optical Zoom"] = "None"
	specs["Front Camera"] = "13 MP"
	specs["Front Camera Video Resolution"] = "1080p @ 30fps"
	specs["Operating System"] = "Android 13, One UI 5.1"
	specs["Battery"] = "5000 mAh"
	specs["Battery Type"] = "Li-Ion (non-removable)"
	specs["Fast Charging"] = "25 W wired"
	specs["Wireless Charging"] = "No"
	specs["5G Support"] = "No"
	specs["Positioning System"] = "GPS, GLONASS, GALILEO, BDS"
	specs["Sensors"] = "Side-mounted fingerprint, Accelerometer, Gyro, Proximity, Compass"
	specs["Special Features"] = "Eye Care Certification (low blue light)"
	specs["Sim Card Type"] = "Nano-SIM or Dual SIM"
	specs["Loudspeaker Quality"] = "Stereo"
	specs["Audio Quality"] = "Not specified"
	specs["Audio Jack"] = "Yes, 3.5 mm"
	specs["Available Colors"] = "Black, Dark Red, Light Green, Silver"
	specs["Announcement Date"] = "April 2023"
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
