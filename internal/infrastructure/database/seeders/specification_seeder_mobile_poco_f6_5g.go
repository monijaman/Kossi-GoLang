package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobilePocoF65g seeds specifications/options for product 'poco-f6-5g'
type SpecificationSeederMobilePocoF65g struct {
	BaseSeeder
}

// NewSpecificationSeederMobilePocoF65g creates a new seeder instance
func NewSpecificationSeederMobilePocoF65g() *SpecificationSeederMobilePocoF65g {
	return &SpecificationSeederMobilePocoF65g{BaseSeeder: BaseSeeder{name: "Specifications for POCO F6 5G"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobilePocoF65g) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080p @ 30fps": "১০৮০p @ ৩০fps",
		"120 Hz": "১২০ Hz",
		"128 / 256 GB": "১২৮ / ২৫৬ GB",
		"198 g": "১৯৮ g",
		"20 MP": "২০ MP",
		"2022": "২০২২",
		"2400 × 1080 px": "২৪০০ × ১০৮০ px",
		"33 W wired": "৩৩ W তারযুক্ত",
		"4700 mAh": "৪৭০০ এমএএইচ",
		"4K @ 30/60fps; 1080p @ 30/60/120fps": "৪K @ ৩০/৬০fps; ১০৮০p @ ৩০/৬০/১২০fps",
		"5.1": "৫.১",
		"6 / 8 GB": "৬ / ৮ GB",
		"6.67 inches": "৬.৬৭ ইঞ্চি",
		"64 MP + 8 MP + 2 MP": "৬৪ MP + ৮ MP + ২ MP",
		"AMOLED, 120 Hz": "AMOLED, ১২০ Hz",
		"Adreno 650": "Adreno ৬৫০",
		"Android 12, MIUI 13": "Android ১২, MIUI ১৩",
		"Black, Blue, Green": "কালো, নীল, সবুজ",
		"Fingerprint (side), Accelerometer, Gyro, Compass": "ফিঙ্গারপ্রিন্ট (side), অ্যাক্সিলেরোমিটার, জাইরো, কম্পাস",
		"GSM / HSPA / LTE / 5G": "GSM / HSPA / LTE / ৫G",
		"Glass front/back, aluminum frame": "গ্লাস সামনে/back, অ্যালুমিনিয়াম ফ্রেম",
		"Li-Ion (non-removable)": "লি-আয়ন (অপসারণযোগ্য নয়)",
		"Nano-SIM + Nano-SIM": "ন্যানো-সিম + ন্যানো-সিম",
		"Octa-core (1×3.2 GHz Cortex-A78 + 3×2.42 GHz Cortex-A78 + 4×1.8 GHz Cortex-A55)": "Octa-core (১×৩.২ GHz Cortex-A৭৮ + ৩×২.৪২ GHz Cortex-A৭৮ + ৪×১.৮ GHz Cortex-A৫৫)",
		"Snapdragon 870": "Snapdragon ৮৭০",
		"Snapdragon 870 (7 nm)": "Snapdragon ৮৭০ (৭ nm)",
		"USB-C 2.0": "USB-C ২.০",
		"Wi-Fi 802.11 a/b/g/n/ac/6": "Wi-Fi ৮০২.১১ a/b/g/n/ac/৬",
	}
}

// Seed inserts specification records for the product identified by slug 'poco-f6-5g'
func (s *SpecificationSeederMobilePocoF65g) Seed(db *gorm.DB) error {
	productSlug := "poco-f6-5g"

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

	// Override model-specific values for POCO F6 5G
	specs["Display Size"] = "6.67 inches"
	specs["Processor"] = "Snapdragon 870"
	specs["Chipset"] = "Snapdragon 870 (7 nm)"
	specs["Cpu Type"] = "Octa-core (1×3.2 GHz Cortex-A78 + 3×2.42 GHz Cortex-A78 + 4×1.8 GHz Cortex-A55)"
	specs["Gpu Type"] = "Adreno 650"
	specs["Ram"] = "6 / 8 GB"
	specs["Storage"] = "128 / 256 GB"
	specs["Display Type"] = "AMOLED, 120 Hz"
	specs["Resolution"] = "2400 × 1080 px"
	specs["Refresh Rate"] = "120 Hz"
	specs["Build Material"] = "Glass front/back, aluminum frame"
	specs["Weight"] = "198 g"
	specs["Network Technology"] = "GSM / HSPA / LTE / 5G"
	specs["Wifi Support"] = "Wi-Fi 802.11 a/b/g/n/ac/6"
	specs["Bluetooth Version"] = "5.1"
	specs["Nfc Support"] = "Yes"
	specs["Usb Type"] = "USB-C 2.0"
	specs["Rear Camera"] = "64 MP + 8 MP + 2 MP"
	specs["Camera Features"] = "OIS, LED flash, HDR"
	specs["Camera Video Resolution"] = "4K @ 30/60fps; 1080p @ 30/60/120fps"
	specs["Optical Zoom"] = "None"
	specs["Front Camera"] = "20 MP"
	specs["Front Camera Video Resolution"] = "1080p @ 30fps"
	specs["Operating System"] = "Android 12, MIUI 13"
	specs["Battery"] = "4700 mAh"
	specs["Battery Type"] = "Li-Ion (non-removable)"
	specs["Fast Charging"] = "33 W wired"
	specs["Wireless Charging"] = "No"
	specs["5G Support"] = "Yes"
	specs["Positioning System"] = "GPS, GLONASS, GALILEO, BDS"
	specs["Sensors"] = "Fingerprint (side), Accelerometer, Gyro, Compass"
	specs["Sim Card Type"] = "Nano-SIM + Nano-SIM"
	specs["Loudspeaker Quality"] = "Stereo"
	specs["Audio Jack"] = "No"
	specs["Available Colors"] = "Black, Blue, Green"
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
