package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileOppoFindN4Flip seeds specifications/options for product 'oppo-find-n4-flip'
type SpecificationSeederMobileOppoFindN4Flip struct {
	BaseSeeder
}

// NewSpecificationSeederMobileOppoFindN4Flip creates a new seeder instance
func NewSpecificationSeederMobileOppoFindN4Flip() *SpecificationSeederMobileOppoFindN4Flip {
	return &SpecificationSeederMobileOppoFindN4Flip{BaseSeeder: BaseSeeder{name: "Specifications for Oppo Find N4 Flip"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileOppoFindN4Flip) getBanglaTranslations() map[string]string {
	return map[string]string{
		"100 W wired": "১০০ W তারযুক্ত",
		"1080p @ 30fps": "১০৮০p @ ৩০fps",
		"12 / 16 GB": "১২ / ১৬ GB",
		"120 Hz": "১২০ Hz",
		"2024": "২০২৪",
		"210 g": "২১০ g",
		"2412 × 1080 px": "২৪১২ × ১০৮০ px",
		"256 / 512 GB": "২৫৬ / ৫১২ GB",
		"2× optical": "২× optical",
		"32 MP": "৩২ MP",
		"4500 mAh": "৪৫০০ এমএএইচ",
		"4K @ 30/60fps; 1080p @ 30/60/120fps": "৪K @ ৩০/৬০fps; ১০৮০p @ ৩০/৬০/১২০fps",
		"5.3": "৫.৩",
		"50 MP + 13 MP + 12 MP": "৫০ MP + ১৩ MP + ১২ MP",
		"50 W wireless": "৫০ W বেতার",
		"6.8 inches": "৬.৮ ইঞ্চি",
		"AMOLED, 120 Hz": "AMOLED, ১২০ Hz",
		"Adreno 740": "Adreno ৭৪০",
		"Android 14, ColorOS 14": "Android ১৪, ColorOS ১৪",
		"Black, Silver": "কালো, রূপালী",
		"Dual SIM (Nano-SIM)": "ডুয়াল সিম (ন্যানো-সিম)",
		"Fingerprint (under display), Accelerometer, Gyro, Proximity, Compass": "ফিঙ্গারপ্রিন্ট (ডিসপ্লের নিচে), অ্যাক্সিলেরোমিটার, জাইরো, প্রক্সিমিটি, কম্পাস",
		"Foldable design, HDR10+": "Foldable design, HDR১০+",
		"GSM / HSPA / LTE / 5G": "GSM / HSPA / LTE / ৫G",
		"Glass front/back, aluminum frame": "গ্লাস সামনে/back, অ্যালুমিনিয়াম ফ্রেম",
		"Li-Ion (non-removable)": "লি-আয়ন (অপসারণযোগ্য নয়)",
		"Snapdragon 8+ Gen 2": "Snapdragon ৮+ Gen ২",
		"Snapdragon 8+ Gen 2 (4 nm)": "Snapdragon ৮+ Gen ২ (৪ nm)",
		"USB-C 3.1": "USB-C ৩.১",
		"Wi-Fi 802.11 a/b/g/n/ac/6/6E": "Wi-Fi ৮০২.১১ a/b/g/n/ac/৬/৬E",
	}
}

// Seed inserts specification records for the product identified by slug 'oppo-find-n4-flip'
func (s *SpecificationSeederMobileOppoFindN4Flip) Seed(db *gorm.DB) error {
	productSlug := "oppo-find-n4-flip"

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

	// Override model-specific values for Oppo Find N4 Flip
	specs["Display Size"] = "6.8 inches"
	specs["Processor"] = "Snapdragon 8+ Gen 2"
	specs["Chipset"] = "Snapdragon 8+ Gen 2 (4 nm)"
	specs["Cpu Type"] = "Octa-core"
	specs["Gpu Type"] = "Adreno 740"
	specs["Ram"] = "12 / 16 GB"
	specs["Storage"] = "256 / 512 GB"
	specs["Display Type"] = "AMOLED, 120 Hz"
	specs["Resolution"] = "2412 × 1080 px"
	specs["Refresh Rate"] = "120 Hz"
	specs["Build Material"] = "Glass front/back, aluminum frame"
	specs["Weight"] = "210 g"
	specs["Network Technology"] = "GSM / HSPA / LTE / 5G"
	specs["Wifi Support"] = "Wi-Fi 802.11 a/b/g/n/ac/6/6E"
	specs["Bluetooth Version"] = "5.3"
	specs["Nfc Support"] = "Yes"
	specs["Usb Type"] = "USB-C 3.1"
	specs["Rear Camera"] = "50 MP + 13 MP + 12 MP"
	specs["Camera Features"] = "OIS (main), LED flash"
	specs["Camera Video Resolution"] = "4K @ 30/60fps; 1080p @ 30/60/120fps"
	specs["Optical Zoom"] = "2× optical"
	specs["Front Camera"] = "32 MP"
	specs["Front Camera Video Resolution"] = "1080p @ 30fps"
	specs["Operating System"] = "Android 14, ColorOS 14"
	specs["Battery"] = "4500 mAh"
	specs["Battery Type"] = "Li-Ion (non-removable)"
	specs["Fast Charging"] = "100 W wired"
	specs["Wireless Charging"] = "50 W wireless"
	specs["5G Support"] = "Yes"
	specs["Positioning System"] = "GPS, GLONASS, GALILEO, BDS"
	specs["Sensors"] = "Fingerprint (under display), Accelerometer, Gyro, Proximity, Compass"
	specs["Special Features"] = "Foldable design, HDR10+"
	specs["Sim Card Type"] = "Dual SIM (Nano-SIM)"
	specs["Loudspeaker Quality"] = "Stereo"
	specs["Audio Quality"] = "Hi-Res Audio"
	specs["Audio Jack"] = "No"
	specs["Available Colors"] = "Black, Silver"
	specs["Announcement Date"] = "2024"
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
