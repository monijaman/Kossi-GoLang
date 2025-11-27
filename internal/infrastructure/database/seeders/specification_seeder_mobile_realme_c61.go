package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileRealmeC61 seeds specifications/options for product 'realme-c61'
type SpecificationSeederMobileRealmeC61 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileRealmeC61 creates a new seeder instance
func NewSpecificationSeederMobileRealmeC61() *SpecificationSeederMobileRealmeC61 {
	return &SpecificationSeederMobileRealmeC61{BaseSeeder: BaseSeeder{name: "Specifications for Realme C61"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileRealmeC61) getBanglaTranslations() map[string]string {
	return map[string]string{
		"10 W wired": "১০ W তারযুক্ত",
		"1080p @ 30fps": "১০৮০p @ ৩০fps",
		"13 MP + 2 MP": "১৩ MP + ২ MP",
		"1600 × 720 px (~270 ppi)": "১৬০০ × ৭২০ px (~২৭০ ppi)",
		"182 g": "১৮২ g",
		"2022": "২০২২",
		"3 / 4 GB": "৩ / ৪ GB",
		"3.5 mm headphone jack": "৩.৫ মিমি headphone jack",
		"32 / 64 GB": "৩২ / ৬৪ GB",
		"5 MP": "৫ MP",
		"5.0": "৫.০",
		"5000 mAh": "৫০০০ এমএএইচ",
		"6.5 inches": "৬.৫ ইঞ্চি",
		"60 Hz": "৬০ Hz",
		"Android 12, Realme UI Go Edition": "Android ১২, Realme UI Go Edition",
		"Black, Blue": "কালো, নীল",
		"Dual SIM (Nano-SIM)": "ডুয়াল সিম (ন্যানো-সিম)",
		"Fingerprint (rear), Accelerometer, Proximity": "ফিঙ্গারপ্রিন্ট (rear), অ্যাক্সিলেরোমিটার, প্রক্সিমিটি",
		"Li-Ion (non-removable)": "লি-আয়ন (অপসারণযোগ্য নয়)",
		"Mali-G57": "Mali-G৫৭",
		"Octa-core (2×1.8 GHz Cortex-A75 + 6×1.8 GHz Cortex-A55)": "Octa-core (২×১.৮ GHz Cortex-A৭৫ + ৬×১.৮ GHz Cortex-A৫৫)",
		"T612 (12 nm)": "T৬১২ (১২ nm)",
		"Unisoc T612": "Unisoc T৬১২",
		"Wi-Fi 802.11 a/b/g/n": "Wi-Fi ৮০২.১১ a/b/g/n",
	}
}

// Seed inserts specification records for the product identified by slug 'realme-c61'
func (s *SpecificationSeederMobileRealmeC61) Seed(db *gorm.DB) error {
	productSlug := "realme-c61"

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

	// Override model-specific values for Realme C61
	specs["Display Size"] = "6.5 inches"
	specs["Processor"] = "Unisoc T612"
	specs["Chipset"] = "T612 (12 nm)"
	specs["Cpu Type"] = "Octa-core (2×1.8 GHz Cortex-A75 + 6×1.8 GHz Cortex-A55)"
	specs["Gpu Type"] = "Mali-G57"
	specs["Ram"] = "3 / 4 GB"
	specs["Storage"] = "32 / 64 GB"
	specs["Display Type"] = "IPS LCD"
	specs["Resolution"] = "1600 × 720 px (~270 ppi)"
	specs["Refresh Rate"] = "60 Hz"
	specs["Build Material"] = "Plastic frame/back"
	specs["Weight"] = "182 g"
	specs["Network Technology"] = "GSM / HSPA / LTE"
	specs["Wifi Support"] = "Wi-Fi 802.11 a/b/g/n"
	specs["Bluetooth Version"] = "5.0"
	specs["Nfc Support"] = "No"
	specs["Usb Type"] = "USB-C"
	specs["Rear Camera"] = "13 MP + 2 MP"
	specs["Camera Features"] = "LED flash"
	specs["Camera Video Resolution"] = "1080p @ 30fps"
	specs["Optical Zoom"] = "None"
	specs["Front Camera"] = "5 MP"
	specs["Front Camera Video Resolution"] = "1080p @ 30fps"
	specs["Operating System"] = "Android 12, Realme UI Go Edition"
	specs["Battery"] = "5000 mAh"
	specs["Battery Type"] = "Li-Ion (non-removable)"
	specs["Fast Charging"] = "10 W wired"
	specs["Wireless Charging"] = "No"
	specs["5G Support"] = "No"
	specs["Positioning System"] = "GPS"
	specs["Sensors"] = "Fingerprint (rear), Accelerometer, Proximity"
	specs["Sim Card Type"] = "Dual SIM (Nano-SIM)"
	specs["Loudspeaker Quality"] = "Mono"
	specs["Audio Jack"] = "3.5 mm headphone jack"
	specs["Available Colors"] = "Black, Blue"
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
