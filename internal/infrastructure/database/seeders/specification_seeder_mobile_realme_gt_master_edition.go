package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileRealmeGtMasterEdition seeds specifications/options for product 'realme-gt-master-edition'
type SpecificationSeederMobileRealmeGtMasterEdition struct {
	BaseSeeder
}

// NewSpecificationSeederMobileRealmeGtMasterEdition creates a new seeder instance
func NewSpecificationSeederMobileRealmeGtMasterEdition() *SpecificationSeederMobileRealmeGtMasterEdition {
	return &SpecificationSeederMobileRealmeGtMasterEdition{BaseSeeder: BaseSeeder{name: "Specifications for Realme GT Master Edition"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileRealmeGtMasterEdition) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080p @ 30fps": "১০৮০p @ ৩০fps",
		"120 Hz": "১২০ Hz",
		"128 / 256 GB": "১২৮ / ২৫৬ GB",
		"174 g": "১৭৪ g",
		"2021": "২০২১",
		"2400 × 1080 px (~409 ppi)": "২৪০০ × ১০৮০ px (~৪০৯ ppi)",
		"32 MP": "৩২ MP",
		"4300 mAh": "৪৩০০ এমএএইচ",
		"4K @ 30fps, 1080p @ 30/60/120fps": "৪K @ ৩০fps, ১০৮০p @ ৩০/৬০/১২০fps",
		"5.2": "৫.২",
		"6 / 8 GB": "৬ / ৮ GB",
		"6.43 inches": "৬.৪৩ ইঞ্চি",
		"64 MP + 8 MP + 2 MP": "৬৪ MP + ৮ MP + ২ MP",
		"65 W wired": "৬৫ W তারযুক্ত",
		"Adreno 642L": "Adreno ৬৪২L",
		"Android 11, Realme UI 2.0": "Android ১১, Realme UI ২.০",
		"Dual SIM (Nano-SIM)": "ডুয়াল সিম (ন্যানো-সিম)",
		"Fingerprint (under display), Accelerometer, Gyro, Proximity, Compass": "ফিঙ্গারপ্রিন্ট (ডিসপ্লের নিচে), অ্যাক্সিলেরোমিটার, জাইরো, প্রক্সিমিটি, কম্পাস",
		"GSM / HSPA / LTE / 5G": "GSM / HSPA / LTE / ৫G",
		"Glass front, plastic frame/back": "গ্লাস সামনে, plastic frame/back",
		"Li-Ion (non-removable)": "লি-আয়ন (অপসারণযোগ্য নয়)",
		"Octa-core (4×2.4 GHz Cortex-A78 + 4×1.8 GHz Cortex-A55)": "Octa-core (৪×২.৪ GHz Cortex-A৭৮ + ৪×১.৮ GHz Cortex-A৫৫)",
		"Qualcomm Snapdragon 778G": "Qualcomm Snapdragon ৭৭৮G",
		"Snapdragon 778G (6 nm)": "Snapdragon ৭৭৮G (৬ nm)",
		"Super AMOLED, 120 Hz": "Super AMOLED, ১২০ Hz",
		"USB-C 2.0": "USB-C ২.০",
		"Voyager Grey, Daybreak Blue": "Voyager ধূসর, Daybreak নীল",
		"Wi-Fi 802.11 a/b/g/n/ac/6": "Wi-Fi ৮০২.১১ a/b/g/n/ac/৬",
	}
}

// Seed inserts specification records for the product identified by slug 'realme-gt-master-edition'
func (s *SpecificationSeederMobileRealmeGtMasterEdition) Seed(db *gorm.DB) error {
	productSlug := "realme-gt-master-edition"

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

	// Override model-specific values for Realme GT Master Edition
	specs["Display Size"] = "6.43 inches"
	specs["Processor"] = "Qualcomm Snapdragon 778G"
	specs["Chipset"] = "Snapdragon 778G (6 nm)"
	specs["Cpu Type"] = "Octa-core (4×2.4 GHz Cortex-A78 + 4×1.8 GHz Cortex-A55)"
	specs["Gpu Type"] = "Adreno 642L"
	specs["Ram"] = "6 / 8 GB"
	specs["Storage"] = "128 / 256 GB"
	specs["Display Type"] = "Super AMOLED, 120 Hz"
	specs["Resolution"] = "2400 × 1080 px (~409 ppi)"
	specs["Refresh Rate"] = "120 Hz"
	specs["Build Material"] = "Glass front, plastic frame/back"
	specs["Weight"] = "174 g"
	specs["Network Technology"] = "GSM / HSPA / LTE / 5G"
	specs["Wifi Support"] = "Wi-Fi 802.11 a/b/g/n/ac/6"
	specs["Bluetooth Version"] = "5.2"
	specs["Nfc Support"] = "Yes"
	specs["Usb Type"] = "USB-C 2.0"
	specs["Rear Camera"] = "64 MP + 8 MP + 2 MP"
	specs["Camera Features"] = "OIS (main), LED flash, HDR"
	specs["Camera Video Resolution"] = "4K @ 30fps, 1080p @ 30/60/120fps"
	specs["Optical Zoom"] = "None"
	specs["Front Camera"] = "32 MP"
	specs["Front Camera Video Resolution"] = "1080p @ 30fps"
	specs["Operating System"] = "Android 11, Realme UI 2.0"
	specs["Battery"] = "4300 mAh"
	specs["Battery Type"] = "Li-Ion (non-removable)"
	specs["Fast Charging"] = "65 W wired"
	specs["Wireless Charging"] = "No"
	specs["5G Support"] = "Yes"
	specs["Positioning System"] = "GPS, GLONASS, GALILEO, BDS"
	specs["Sensors"] = "Fingerprint (under display), Accelerometer, Gyro, Proximity, Compass"
	specs["Sim Card Type"] = "Dual SIM (Nano-SIM)"
	specs["Loudspeaker Quality"] = "Stereo"
	specs["Audio Jack"] = "No"
	specs["Available Colors"] = "Voyager Grey, Daybreak Blue"
	specs["Announcement Date"] = "2021"
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
