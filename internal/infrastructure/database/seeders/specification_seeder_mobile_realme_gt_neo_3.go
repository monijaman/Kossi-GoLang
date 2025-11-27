package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileRealmeGtNeo3 seeds specifications/options for product 'realme-gt-neo-3'
type SpecificationSeederMobileRealmeGtNeo3 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileRealmeGtNeo3 creates a new seeder instance
func NewSpecificationSeederMobileRealmeGtNeo3() *SpecificationSeederMobileRealmeGtNeo3 {
	return &SpecificationSeederMobileRealmeGtNeo3{BaseSeeder: BaseSeeder{name: "Specifications for Realme GT Neo 3"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileRealmeGtNeo3) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080p @ 30fps": "১০৮০p @ ৩০fps",
		"120 Hz": "১২০ Hz",
		"128 / 256 GB": "১২৮ / ২৫৬ GB",
		"150 W wired": "১৫০ W তারযুক্ত",
		"16 MP": "১৬ MP",
		"188 g": "১৮৮ g",
		"2022": "২০২২",
		"2412 × 1080 px (~394 ppi)": "২৪১২ × ১০৮০ px (~৩৯৪ ppi)",
		"4500 mAh": "৪৫০০ এমএএইচ",
		"4K @ 30/60fps, 1080p @ 30/60/120fps": "৪K @ ৩০/৬০fps, ১০৮০p @ ৩০/৬০/১২০fps",
		"5.2": "৫.২",
		"50 MP + 8 MP + 2 MP": "৫০ MP + ৮ MP + ২ MP",
		"6.7 inches": "৬.৭ ইঞ্চি",
		"8 / 12 GB": "৮ / ১২ GB",
		"AMOLED, 120 Hz": "AMOLED, ১২০ Hz",
		"Android 12, Realme UI 3.0": "Android ১২, Realme UI ৩.০",
		"Dimensity 8100-Max (5 nm)": "Dimensity ৮১০০-Max (৫ nm)",
		"Dual SIM (Nano-SIM)": "ডুয়াল সিম (ন্যানো-সিম)",
		"Fingerprint (under display), Accelerometer, Gyro, Proximity, Compass": "ফিঙ্গারপ্রিন্ট (ডিসপ্লের নিচে), অ্যাক্সিলেরোমিটার, জাইরো, প্রক্সিমিটি, কম্পাস",
		"GSM / HSPA / LTE / 5G": "GSM / HSPA / LTE / ৫G",
		"Glass front, plastic frame/back": "গ্লাস সামনে, plastic frame/back",
		"Li-Ion (non-removable)": "লি-আয়ন (অপসারণযোগ্য নয়)",
		"Mali-G610 MC6": "Mali-G৬১০ MC৬",
		"MediaTek Dimensity 8100-Max": "MediaTek Dimensity ৮১০০-Max",
		"Octa-core (4×2.85 GHz Cortex-A78 + 4×2.0 GHz Cortex-A55)": "Octa-core (৪×২.৮৫ GHz Cortex-A৭৮ + ৪×২.০ GHz Cortex-A৫৫)",
		"Sprint Blue, Nitro Blue, Shade Black": "Sprint নীল, Nitro নীল, Shade কালো",
		"USB-C 2.0": "USB-C ২.০",
		"Wi-Fi 802.11 a/b/g/n/ac/6": "Wi-Fi ৮০২.১১ a/b/g/n/ac/৬",
	}
}

// Seed inserts specification records for the product identified by slug 'realme-gt-neo-3'
func (s *SpecificationSeederMobileRealmeGtNeo3) Seed(db *gorm.DB) error {
	productSlug := "realme-gt-neo-3"

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

	// Override model-specific values for Realme GT Neo 3
	specs["Display Size"] = "6.7 inches"
	specs["Processor"] = "MediaTek Dimensity 8100-Max"
	specs["Chipset"] = "Dimensity 8100-Max (5 nm)"
	specs["Cpu Type"] = "Octa-core (4×2.85 GHz Cortex-A78 + 4×2.0 GHz Cortex-A55)"
	specs["Gpu Type"] = "Mali-G610 MC6"
	specs["Ram"] = "8 / 12 GB"
	specs["Storage"] = "128 / 256 GB"
	specs["Display Type"] = "AMOLED, 120 Hz"
	specs["Resolution"] = "2412 × 1080 px (~394 ppi)"
	specs["Refresh Rate"] = "120 Hz"
	specs["Build Material"] = "Glass front, plastic frame/back"
	specs["Weight"] = "188 g"
	specs["Network Technology"] = "GSM / HSPA / LTE / 5G"
	specs["Wifi Support"] = "Wi-Fi 802.11 a/b/g/n/ac/6"
	specs["Bluetooth Version"] = "5.2"
	specs["Nfc Support"] = "Yes"
	specs["Usb Type"] = "USB-C 2.0"
	specs["Rear Camera"] = "50 MP + 8 MP + 2 MP"
	specs["Camera Features"] = "OIS (main), LED flash, HDR, panorama"
	specs["Camera Video Resolution"] = "4K @ 30/60fps, 1080p @ 30/60/120fps"
	specs["Optical Zoom"] = "None"
	specs["Front Camera"] = "16 MP"
	specs["Front Camera Video Resolution"] = "1080p @ 30fps"
	specs["Operating System"] = "Android 12, Realme UI 3.0"
	specs["Battery"] = "4500 mAh"
	specs["Battery Type"] = "Li-Ion (non-removable)"
	specs["Fast Charging"] = "150 W wired"
	specs["Wireless Charging"] = "No"
	specs["5G Support"] = "Yes"
	specs["Positioning System"] = "GPS, GLONASS, GALILEO, BDS"
	specs["Sensors"] = "Fingerprint (under display), Accelerometer, Gyro, Proximity, Compass"
	specs["Sim Card Type"] = "Dual SIM (Nano-SIM)"
	specs["Loudspeaker Quality"] = "Stereo"
	specs["Audio Jack"] = "No"
	specs["Available Colors"] = "Sprint Blue, Nitro Blue, Shade Black"
	specs["Announcement Date"] = "2022"
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
