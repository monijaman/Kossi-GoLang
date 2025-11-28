package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobilePocoC65 seeds specifications/options for product 'poco-c65'
type SpecificationSeederMobilePocoC65 struct {
	BaseSeeder
}

// NewSpecificationSeederMobilePocoC65 creates a new seeder instance
func NewSpecificationSeederMobilePocoC65() *SpecificationSeederMobilePocoC65 {
	return &SpecificationSeederMobilePocoC65{BaseSeeder: BaseSeeder{name: "Specifications for POCO C65"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobilePocoC65) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080p @ 30fps": "১০৮০p @ ৩০fps",
		"13 MP + 2 MP": "১৩ মেগাপিক্সেল + ২ মেগাপিক্সেল",
		"18 W wired": "১৮ W তারযুক্ত",
		"185 g": "১৮৫ g",
		"2023": "২০২৩",
		"3 / 4 GB": "৩ / ৪ GB",
		"3.5 mm": "৩.৫ মিমি",
		"32 / 64 GB": "৩২ / ৬৪ GB",
		"5 MP": "৫ মেগাপিক্সেল",
		"5.0": "৫.০",
		"5000 mAh": "৫০০০ এমএএইচ",
		"6.53 inches": "৬.৫৩ ইঞ্চি",
		"60 Hz": "৬০ Hz",
		"720 × 1600 px": "৭২০ × ১৬০০ পিক্সেল",
		"Accelerometer, Proximity, Compass": "অ্যাক্সিলেরোমিটার, প্রক্সিমিটি, কম্পাস",
		"Android 12, MIUI 13": "অ্যান্ড্রয়েড ১২, এমআইইউআই ১৩",
		"Black, Blue, Green": "কালো, নীল, সবুজ",
		"Helio G35 (12 nm)": "হেলিও G৩৫ (১২ ন্যানোমিটার)",
		"Li-Ion (non-removable)": "লি-আয়ন (অপসারণযোগ্য নয়)",
		"MediaTek Helio G35": "মিডিয়াটেক হেলিও G৩৫",
		"Nano-SIM + Nano-SIM": "ন্যানো-সিম + ন্যানো-সিম",
		"Plastic frame / back, Glass front": "Plastic ফ্রেম / পেছনে, গ্লাস সামনে",
		"PowerVR GE8320": "পাওয়ারভিআর GE৮৩২০",
		"Wi-Fi 802.11 b/g/n": "ওয়াই-ফাই ৮০২.১১ b/g/n",
	}
}

// Seed inserts specification records for the product identified by slug 'poco-c65'
func (s *SpecificationSeederMobilePocoC65) Seed(db *gorm.DB) error {
	productSlug := "poco-c65"

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

	// Override model-specific values for POCO C65
	specs["Display Size"] = "6.53 inches"
	specs["Processor"] = "MediaTek Helio G35"
	specs["Chipset"] = "Helio G35 (12 nm)"
	specs["Cpu Type"] = "Octa-core"
	specs["Gpu Type"] = "PowerVR GE8320"
	specs["Ram"] = "3 / 4 GB"
	specs["Storage"] = "32 / 64 GB"
	specs["Display Type"] = "IPS LCD"
	specs["Resolution"] = "720 × 1600 px"
	specs["Refresh Rate"] = "60 Hz"
	specs["Build Material"] = "Plastic frame / back, Glass front"
	specs["Weight"] = "185 g"
	specs["Network Technology"] = "GSM / HSPA / LTE"
	specs["Wifi Support"] = "Wi-Fi 802.11 b/g/n"
	specs["Bluetooth Version"] = "5.0"
	specs["Nfc Support"] = "No"
	specs["Usb Type"] = "USB-C"
	specs["Rear Camera"] = "13 MP + 2 MP"
	specs["Camera Features"] = "LED flash"
	specs["Camera Video Resolution"] = "1080p @ 30fps"
	specs["Optical Zoom"] = "None"
	specs["Front Camera"] = "5 MP"
	specs["Front Camera Video Resolution"] = "1080p @ 30fps"
	specs["Operating System"] = "Android 12, MIUI 13"
	specs["Battery"] = "5000 mAh"
	specs["Battery Type"] = "Li-Ion (non-removable)"
	specs["Fast Charging"] = "18 W wired"
	specs["Wireless Charging"] = "No"
	specs["5G Support"] = "No"
	specs["Positioning System"] = "GPS"
	specs["Sensors"] = "Accelerometer, Proximity, Compass"
	specs["Sim Card Type"] = "Nano-SIM + Nano-SIM"
	specs["Loudspeaker Quality"] = "Mono"
	specs["Audio Jack"] = "3.5 mm"
	specs["Available Colors"] = "Black, Blue, Green"
	specs["Announcement Date"] = "2023"
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
