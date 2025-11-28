package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileRealmeC71 seeds specifications/options for product 'realme-c71'
type SpecificationSeederMobileRealmeC71 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileRealmeC71 creates a new seeder instance
func NewSpecificationSeederMobileRealmeC71() *SpecificationSeederMobileRealmeC71 {
	return &SpecificationSeederMobileRealmeC71{BaseSeeder: BaseSeeder{name: "Specifications for Realme C71"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileRealmeC71) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080p @ 30fps": "১০৮০p @ ৩০fps",
		"13 MP + 2 MP": "১৩ মেগাপিক্সেল + ২ মেগাপিক্সেল",
		"1600 × 720 px (~269 ppi)": "১৬০০ × ৭২০ পিক্সেল (~২৬৯ পিপিআই)",
		"18 W wired": "১৮ W তারযুক্ত",
		"190 g": "১৯০ g",
		"2023": "২০২৩",
		"3.5 mm headphone jack": "৩.৫ মিমি হেডফোন জ্যাক",
		"4 GB": "৪ GB",
		"5.0": "৫.০",
		"5000 mAh": "৫০০০ এমএএইচ",
		"6.52 inches": "৬.৫২ ইঞ্চি",
		"60 Hz": "৬০ Hz",
		"64 / 128 GB": "৬৪ / ১২৮ GB",
		"8 MP": "৮ মেগাপিক্সেল",
		"Android 12, Realme UI 3.0": "অ্যান্ড্রয়েড ১২, Realme UI ৩.০",
		"Black, Blue": "কালো, নীল",
		"Dual SIM (Nano-SIM)": "ডুয়াল সিম (ন্যানো-সিম)",
		"Fingerprint (rear), Accelerometer, Proximity, Compass": "ফিঙ্গারপ্রিন্ট (rear), অ্যাক্সিলেরোমিটার, প্রক্সিমিটি, কম্পাস",
		"Helio G88 (12 nm)": "হেলিও G৮৮ (১২ ন্যানোমিটার)",
		"Li-Ion (non-removable)": "লি-আয়ন (অপসারণযোগ্য নয়)",
		"Mali-G52 MC2": "মালি-G৫২ MC২",
		"MediaTek Helio G88": "মিডিয়াটেক হেলিও G৮৮",
		"Octa-core (2×2.0 GHz Cortex-A75 + 6×1.8 GHz Cortex-A55)": "অক্টা-কোর (২×২.০ গিগাহার্টজ Cবাtex-A৭৫ + ৬×১.৮ গিগাহার্টজ Cবাtex-A৫৫)",
		"Wi-Fi 802.11 a/b/g/n": "ওয়াই-ফাই ৮০২.১১ a/b/g/n",
	}
}

// Seed inserts specification records for the product identified by slug 'realme-c71'
func (s *SpecificationSeederMobileRealmeC71) Seed(db *gorm.DB) error {
	productSlug := "realme-c71"

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

	// Override model-specific values for Realme C71
	specs["Display Size"] = "6.52 inches"
	specs["Processor"] = "MediaTek Helio G88"
	specs["Chipset"] = "Helio G88 (12 nm)"
	specs["Cpu Type"] = "Octa-core (2×2.0 GHz Cortex-A75 + 6×1.8 GHz Cortex-A55)"
	specs["Gpu Type"] = "Mali-G52 MC2"
	specs["Ram"] = "4 GB"
	specs["Storage"] = "64 / 128 GB"
	specs["Display Type"] = "IPS LCD"
	specs["Resolution"] = "1600 × 720 px (~269 ppi)"
	specs["Refresh Rate"] = "60 Hz"
	specs["Build Material"] = "Plastic frame/back"
	specs["Weight"] = "190 g"
	specs["Network Technology"] = "GSM / HSPA / LTE"
	specs["Wifi Support"] = "Wi-Fi 802.11 a/b/g/n"
	specs["Bluetooth Version"] = "5.0"
	specs["Nfc Support"] = "Yes"
	specs["Usb Type"] = "USB-C"
	specs["Rear Camera"] = "13 MP + 2 MP"
	specs["Camera Features"] = "LED flash, HDR"
	specs["Camera Video Resolution"] = "1080p @ 30fps"
	specs["Optical Zoom"] = "None"
	specs["Front Camera"] = "8 MP"
	specs["Front Camera Video Resolution"] = "1080p @ 30fps"
	specs["Operating System"] = "Android 12, Realme UI 3.0"
	specs["Battery"] = "5000 mAh"
	specs["Battery Type"] = "Li-Ion (non-removable)"
	specs["Fast Charging"] = "18 W wired"
	specs["Wireless Charging"] = "No"
	specs["5G Support"] = "No"
	specs["Positioning System"] = "GPS, GLONASS, GALILEO, BDS"
	specs["Sensors"] = "Fingerprint (rear), Accelerometer, Proximity, Compass"
	specs["Sim Card Type"] = "Dual SIM (Nano-SIM)"
	specs["Loudspeaker Quality"] = "Mono"
	specs["Audio Jack"] = "3.5 mm headphone jack"
	specs["Available Colors"] = "Black, Blue"
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
