package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobilePocoC75 seeds specifications/options for product 'poco-c75'
type SpecificationSeederMobilePocoC75 struct {
	BaseSeeder
}

// NewSpecificationSeederMobilePocoC75 creates a new seeder instance
func NewSpecificationSeederMobilePocoC75() *SpecificationSeederMobilePocoC75 {
	return &SpecificationSeederMobilePocoC75{BaseSeeder: BaseSeeder{name: "Specifications for POCO C75"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobilePocoC75) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080p @ 30fps": "১০৮০p @ ৩০fps",
		"18 W wired": "১৮ W তারযুক্ত",
		"190 g": "১৯০ g",
		"2023": "২০২৩",
		"3 / 4 GB": "৩ / ৪ GB",
		"3.5 mm": "৩.৫ মিমি",
		"5.0": "৫.০",
		"50 MP + 2 MP": "৫০ মেগাপিক্সেল + ২ মেগাপিক্সেল",
		"5000 mAh": "৫০০০ এমএএইচ",
		"6.6 inches": "৬.৬ ইঞ্চি",
		"64 / 128 GB": "৬৪ / ১২৮ GB",
		"720 × 1600 px": "৭২০ × ১৬০০ পিক্সেল",
		"8 MP": "৮ মেগাপিক্সেল",
		"90 Hz": "৯০ Hz",
		"Accelerometer, Proximity, Compass": "অ্যাক্সিলেরোমিটার, প্রক্সিমিটি, কম্পাস",
		"Android 12, MIUI 13": "অ্যান্ড্রয়েড ১২, এমআইইউআই ১৩",
		"Black, Blue, Red": "কালো, নীল, লাল",
		"Helio G88 (12 nm)": "হেলিও G৮৮ (১২ ন্যানোমিটার)",
		"Li-Ion (non-removable)": "লি-আয়ন (অপসারণযোগ্য নয়)",
		"Mali-G52 MC2": "মালি-G৫২ MC২",
		"MediaTek Helio G88": "মিডিয়াটেক হেলিও G৮৮",
		"Nano-SIM + Nano-SIM": "ন্যানো-সিম + ন্যানো-সিম",
		"Plastic frame / back, Glass front": "Plastic ফ্রেম / পেছনে, গ্লাস সামনে",
		"Wi-Fi 802.11 b/g/n": "ওয়াই-ফাই ৮০২.১১ b/g/n",
	}
}

// Seed inserts specification records for the product identified by slug 'poco-c75'
func (s *SpecificationSeederMobilePocoC75) Seed(db *gorm.DB) error {
	productSlug := "poco-c75"

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

	// Override model-specific values for POCO C75
	specs["Display Size"] = "6.6 inches"
	specs["Processor"] = "MediaTek Helio G88"
	specs["Chipset"] = "Helio G88 (12 nm)"
	specs["Cpu Type"] = "Octa-core"
	specs["Gpu Type"] = "Mali-G52 MC2"
	specs["Ram"] = "3 / 4 GB"
	specs["Storage"] = "64 / 128 GB"
	specs["Display Type"] = "IPS LCD"
	specs["Resolution"] = "720 × 1600 px"
	specs["Refresh Rate"] = "90 Hz"
	specs["Build Material"] = "Plastic frame / back, Glass front"
	specs["Weight"] = "190 g"
	specs["Network Technology"] = "GSM / HSPA / LTE"
	specs["Wifi Support"] = "Wi-Fi 802.11 b/g/n"
	specs["Bluetooth Version"] = "5.0"
	specs["Nfc Support"] = "No"
	specs["Usb Type"] = "USB-C"
	specs["Rear Camera"] = "50 MP + 2 MP"
	specs["Camera Features"] = "LED flash"
	specs["Camera Video Resolution"] = "1080p @ 30fps"
	specs["Optical Zoom"] = "None"
	specs["Front Camera"] = "8 MP"
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
	specs["Available Colors"] = "Black, Blue, Red"
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
