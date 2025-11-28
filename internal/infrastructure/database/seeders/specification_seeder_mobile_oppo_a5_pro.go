package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileOppoA5Pro seeds specifications/options for product 'oppo-a5-pro'
type SpecificationSeederMobileOppoA5Pro struct {
	BaseSeeder
}

// NewSpecificationSeederMobileOppoA5Pro creates a new seeder instance
func NewSpecificationSeederMobileOppoA5Pro() *SpecificationSeederMobileOppoA5Pro {
	return &SpecificationSeederMobileOppoA5Pro{BaseSeeder: BaseSeeder{name: "Specifications for Oppo A5 Pro"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileOppoA5Pro) getBanglaTranslations() map[string]string {
	return map[string]string{
		"10 W wired": "১০ W তারযুক্ত",
		"1080p @ 30fps": "১০৮০p @ ৩০fps",
		"12 MP + 2 MP": "১২ মেগাপিক্সেল + ২ মেগাপিক্সেল",
		"1600 × 720 px": "১৬০০ × ৭২০ পিক্সেল",
		"190 g": "১৯০ g",
		"2022": "২০২২",
		"3.5 mm headphone jack": "৩.৫ মিমি হেডফোন জ্যাক",
		"4 / 6 GB": "৪ / ৬ GB",
		"5.0": "৫.০",
		"5000 mAh": "৫০০০ এমএএইচ",
		"6.5 inches": "৬.৫ ইঞ্চি",
		"60 Hz": "৬০ Hz",
		"64 / 128 GB + microSD": "৬৪ / ১২৮ জিবি + মাইক্রোএসডি",
		"8 MP": "৮ মেগাপিক্সেল",
		"Android 12, ColorOS 12": "অ্যান্ড্রয়েড ১২, কালারওএস ১২",
		"Black, Blue": "কালো, নীল",
		"Dual SIM (Nano-SIM)": "ডুয়াল সিম (ন্যানো-সিম)",
		"Fingerprint (rear), Accelerometer, Proximity, Compass": "ফিঙ্গারপ্রিন্ট (rear), অ্যাক্সিলেরোমিটার, প্রক্সিমিটি, কম্পাস",
		"Glass front, plastic frame/back": "গ্লাস সামনে, প্লাস্টিক ফ্রেম/পেছনে",
		"Helio P35 (12 nm)": "হেলিও P৩৫ (১২ ন্যানোমিটার)",
		"Li-Ion (non-removable)": "লি-আয়ন (অপসারণযোগ্য নয়)",
		"MediaTek Helio P35": "মিডিয়াটেক হেলিও P৩৫",
		"PowerVR GE8320": "পাওয়ারভিআর GE৮৩২০",
		"Wi-Fi 802.11 a/b/g/n": "ওয়াই-ফাই ৮০২.১১ a/b/g/n",
	}
}

// Seed inserts specification records for the product identified by slug 'oppo-a5-pro'
func (s *SpecificationSeederMobileOppoA5Pro) Seed(db *gorm.DB) error {
	productSlug := "oppo-a5-pro"

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

	// Override model-specific values for Oppo A5 Pro
	specs["Display Size"] = "6.5 inches"
	specs["Processor"] = "MediaTek Helio P35"
	specs["Chipset"] = "Helio P35 (12 nm)"
	specs["Cpu Type"] = "Octa-core"
	specs["Gpu Type"] = "PowerVR GE8320"
	specs["Ram"] = "4 / 6 GB"
	specs["Storage"] = "64 / 128 GB + microSD"
	specs["Display Type"] = "IPS LCD"
	specs["Resolution"] = "1600 × 720 px"
	specs["Refresh Rate"] = "60 Hz"
	specs["Build Material"] = "Glass front, plastic frame/back"
	specs["Weight"] = "190 g"
	specs["Network Technology"] = "GSM / HSPA / LTE"
	specs["Wifi Support"] = "Wi-Fi 802.11 a/b/g/n"
	specs["Bluetooth Version"] = "5.0"
	specs["Nfc Support"] = "No"
	specs["Usb Type"] = "Micro-USB"
	specs["Rear Camera"] = "12 MP + 2 MP"
	specs["Camera Features"] = "LED flash, HDR"
	specs["Camera Video Resolution"] = "1080p @ 30fps"
	specs["Optical Zoom"] = "None"
	specs["Front Camera"] = "8 MP"
	specs["Front Camera Video Resolution"] = "1080p @ 30fps"
	specs["Operating System"] = "Android 12, ColorOS 12"
	specs["Battery"] = "5000 mAh"
	specs["Battery Type"] = "Li-Ion (non-removable)"
	specs["Fast Charging"] = "10 W wired"
	specs["Wireless Charging"] = "No"
	specs["5G Support"] = "No"
	specs["Positioning System"] = "GPS, GLONASS"
	specs["Sensors"] = "Fingerprint (rear), Accelerometer, Proximity, Compass"
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
