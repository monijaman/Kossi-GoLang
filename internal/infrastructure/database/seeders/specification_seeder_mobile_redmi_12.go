package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileRedmi12 seeds specifications/options for product 'redmi-12'
type SpecificationSeederMobileRedmi12 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileRedmi12 creates a new seeder instance
func NewSpecificationSeederMobileRedmi12() *SpecificationSeederMobileRedmi12 {
	return &SpecificationSeederMobileRedmi12{BaseSeeder: BaseSeeder{name: "Specifications for Redmi 12"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileRedmi12) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080p @ 30fps": "১০৮০p @ ৩০fps",
		"128 / 256 GB + microSD": "১২৮ / ২৫৬ GB + microSD",
		"1650 × 720 px": "১৬৫০ × ৭২০ px",
		"168.8 × 76.5 × 8.9 mm": "১৬৮.৮ × ৭৬.৫ × ৮.৯ মিমি",
		"18 W wired": "১৮ W তারযুক্ত",
		"198 g": "১৯৮ g",
		"2022": "২০২২",
		"3.5 mm": "৩.৫ মিমি",
		"4 / 6 GB": "৪ / ৬ GB",
		"5.1": "৫.১",
		"50 MP + 2 MP": "৫০ MP + ২ MP",
		"5000 mAh": "৫০০০ এমএএইচ",
		"6.71 inches": "৬.৭১ ইঞ্চি",
		"8 MP": "৮ MP",
		"90 Hz": "৯০ Hz",
		"Android 13, MIUI 14": "Android ১৩, MIUI ১৪",
		"Black, Sky Blue, Green": "কালো, Sky নীল, সবুজ",
		"Fingerprint (side), Accelerometer, Compass": "ফিঙ্গারপ্রিন্ট (side), অ্যাক্সিলেরোমিটার, কম্পাস",
		"Helio G99": "Helio G৯৯",
		"IPS LCD, 90 Hz": "IPS LCD, ৯০ Hz",
		"Li‑Ion (non-removable)": "Li‑Ion (অপসারণযোগ্য নয়)",
		"Mali-G57 MC2": "Mali-G৫৭ MC২",
		"MediaTek Helio G99": "MediaTek Helio G৯৯",
		"Nano-SIM + Nano-SIM": "ন্যানো-সিম + ন্যানো-সিম",
		"Wi‑Fi 802.11 a/b/g/n/ac": "Wi‑Fi ৮০২.১১ a/b/g/n/ac",
	}
}

// Seed inserts specification records for the product identified by slug 'redmi-12'
func (s *SpecificationSeederMobileRedmi12) Seed(db *gorm.DB) error {
	productSlug := "redmi-12"

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

	// Override model-specific values for Redmi 12
	specs["Display Size"] = "6.71 inches"
	specs["Processor"] = "MediaTek Helio G99"
	specs["Chipset"] = "Helio G99"
	specs["Cpu Type"] = "Octa-core"
	specs["Gpu Type"] = "Mali-G57 MC2"
	specs["Ram"] = "4 / 6 GB"
	specs["Storage"] = "128 / 256 GB + microSD"
	specs["Display Type"] = "IPS LCD, 90 Hz"
	specs["Resolution"] = "1650 × 720 px"
	specs["Refresh Rate"] = "90 Hz"
	specs["Build Material"] = "Plastic body"
	specs["Weight"] = "198 g"
	specs["Dimensions"] = "168.8 × 76.5 × 8.9 mm"
	specs["Network Technology"] = "GSM / HSPA / LTE"
	specs["Wifi Support"] = "Wi‑Fi 802.11 a/b/g/n/ac"
	specs["Bluetooth Version"] = "5.1"
	specs["Nfc Support"] = "Depends on region"
	specs["Usb Type"] = "USB‑C"
	specs["Rear Camera"] = "50 MP + 2 MP"
	specs["Camera Features"] = "LED flash, HDR"
	specs["Camera Video Resolution"] = "1080p @ 30fps"
	specs["Optical Zoom"] = "None"
	specs["Front Camera"] = "8 MP"
	specs["Front Camera Video Resolution"] = "1080p @ 30fps"
	specs["Operating System"] = "Android 13, MIUI 14"
	specs["Battery"] = "5000 mAh"
	specs["Battery Type"] = "Li‑Ion (non-removable)"
	specs["Fast Charging"] = "18 W wired"
	specs["Wireless Charging"] = "No"
	specs["5G Support"] = "No"
	specs["Positioning System"] = "GPS, GLONASS, GALILEO"
	specs["Sensors"] = "Fingerprint (side), Accelerometer, Compass"
	specs["Sim Card Type"] = "Nano-SIM + Nano-SIM"
	specs["Loudspeaker Quality"] = "Single speaker"
	specs["Audio Jack"] = "3.5 mm"
	specs["Available Colors"] = "Black, Sky Blue, Green"
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
