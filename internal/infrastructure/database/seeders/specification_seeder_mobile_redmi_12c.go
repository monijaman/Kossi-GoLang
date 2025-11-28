package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileRedmi12c seeds specifications/options for product 'redmi-12c'
type SpecificationSeederMobileRedmi12c struct {
	BaseSeeder
}

// NewSpecificationSeederMobileRedmi12c creates a new seeder instance
func NewSpecificationSeederMobileRedmi12c() *SpecificationSeederMobileRedmi12c {
	return &SpecificationSeederMobileRedmi12c{BaseSeeder: BaseSeeder{name: "Specifications for Redmi 12C"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileRedmi12c) getBanglaTranslations() map[string]string {
	return map[string]string{
		"10 W wired": "১০ W তারযুক্ত",
		"1080p @ 30fps": "১০৮০p @ ৩০fps",
		"13 MP + 2 MP": "১৩ মেগাপিক্সেল + ২ মেগাপিক্সেল",
		"1650 × 720 px": "১৬৫০ × ৭২০ পিক্সেল",
		"168.8 × 76.5 × 8.9 mm": "১৬৮.৮ × ৭৬.৫ × ৮.৯ মিমি",
		"196 g": "১৯৬ g",
		"2022": "২০২২",
		"3 / 4 / 6 GB": "৩ / ৪ / ৬ GB",
		"3.5 mm": "৩.৫ মিমি",
		"5 MP": "৫ মেগাপিক্সেল",
		"5.1": "৫.১",
		"5000 mAh": "৫০০০ এমএএইচ",
		"6.71 inches": "৬.৭১ ইঞ্চি",
		"60 Hz": "৬০ Hz",
		"64 / 128 GB + microSD": "৬৪ / ১২৮ জিবি + মাইক্রোএসডি",
		"Accelerometer, Proximity": "অ্যাক্সিলেরোমিটার, প্রক্সিমিটি",
		"Android 13 (Go Edition) / MIUI": "অ্যান্ড্রয়েড ১৩ (Go Edition) / এমআইইউআই",
		"Graphite Gray, Coral Green, Sky Blue": "গ্রাফাইট ধূসর, কোরাল সবুজ, স্কাই নীল",
		"Helio G85": "হেলিও G৮৫",
		"Li‑Ion (non-removable)": "Li‑Ion (অপসারণযোগ্য নয়)",
		"Mali-G52": "মালি-G৫২",
		"MediaTek Helio G85": "মিডিয়াটেক হেলিও G৮৫",
		"Nano-SIM + Nano-SIM": "ন্যানো-সিম + ন্যানো-সিম",
		"Plastic back, glass front": "প্লাস্টিক পেছনে, গ্লাস সামনে",
		"Wi‑Fi 802.11 a/b/g/n/ac": "Wi‑Fi ৮০২.১১ a/b/g/n/ac",
	}
}

// Seed inserts specification records for the product identified by slug 'redmi-12c'
func (s *SpecificationSeederMobileRedmi12c) Seed(db *gorm.DB) error {
	productSlug := "redmi-12c"

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

	// Override model-specific values for Redmi 12C
	specs["Display Size"] = "6.71 inches"
	specs["Processor"] = "MediaTek Helio G85"
	specs["Chipset"] = "Helio G85"
	specs["Cpu Type"] = "Octa-core"
	specs["Gpu Type"] = "Mali-G52"
	specs["Ram"] = "3 / 4 / 6 GB"
	specs["Storage"] = "64 / 128 GB + microSD"
	specs["Display Type"] = "IPS LCD"
	specs["Resolution"] = "1650 × 720 px"
	specs["Refresh Rate"] = "60 Hz"
	specs["Build Material"] = "Plastic back, glass front"
	specs["Weight"] = "196 g"
	specs["Dimensions"] = "168.8 × 76.5 × 8.9 mm"
	specs["Network Technology"] = "GSM / HSPA / LTE"
	specs["Wifi Support"] = "Wi‑Fi 802.11 a/b/g/n/ac"
	specs["Bluetooth Version"] = "5.1"
	specs["Nfc Support"] = "Sometimes, based on region"
	specs["Usb Type"] = "USB‑C"
	specs["Rear Camera"] = "13 MP + 2 MP"
	specs["Camera Features"] = "LED flash, Portrait"
	specs["Camera Video Resolution"] = "1080p @ 30fps"
	specs["Optical Zoom"] = "None"
	specs["Front Camera"] = "5 MP"
	specs["Front Camera Video Resolution"] = "1080p @ 30fps"
	specs["Operating System"] = "Android 13 (Go Edition) / MIUI"
	specs["Battery"] = "5000 mAh"
	specs["Battery Type"] = "Li‑Ion (non-removable)"
	specs["Fast Charging"] = "10 W wired"
	specs["Wireless Charging"] = "No"
	specs["5G Support"] = "No"
	specs["Positioning System"] = "GPS, GLONASS"
	specs["Sensors"] = "Accelerometer, Proximity"
	specs["Sim Card Type"] = "Nano-SIM + Nano-SIM"
	specs["Loudspeaker Quality"] = "Mono"
	specs["Audio Jack"] = "3.5 mm"
	specs["Available Colors"] = "Graphite Gray, Coral Green, Sky Blue"
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
