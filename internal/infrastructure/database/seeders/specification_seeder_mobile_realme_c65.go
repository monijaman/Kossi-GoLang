package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileRealmeC65 seeds specifications/options for product 'realme-c65'
type SpecificationSeederMobileRealmeC65 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileRealmeC65 creates a new seeder instance
func NewSpecificationSeederMobileRealmeC65() *SpecificationSeederMobileRealmeC65 {
	return &SpecificationSeederMobileRealmeC65{BaseSeeder: BaseSeeder{name: "Specifications for Realme C65"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileRealmeC65) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080p @ 30fps": "১০৮০p @ ৩০fps",
		"1600 × 720 px (~270 ppi)": "১৬০০ × ৭২০ পিক্সেল (~২৭০ পিপিআই)",
		"18 W wired": "১৮ W তারযুক্ত",
		"185 g": "১৮৫ g",
		"2022": "২০২২",
		"3.5 mm headphone jack": "৩.৫ মিমি হেডফোন জ্যাক",
		"4 / 6 GB": "৪ / ৬ GB",
		"5.0": "৫.০",
		"50 MP + 2 MP": "৫০ মেগাপিক্সেল + ২ মেগাপিক্সেল",
		"5000 mAh": "৫০০০ এমএএইচ",
		"6.5 inches": "৬.৫ ইঞ্চি",
		"60 Hz": "৬০ Hz",
		"64 / 128 GB": "৬৪ / ১২৮ GB",
		"8 MP": "৮ মেগাপিক্সেল",
		"Android 12, Realme UI 3.0": "অ্যান্ড্রয়েড ১২, Realme UI ৩.০",
		"Black, Blue, Green": "কালো, নীল, সবুজ",
		"Dual SIM (Nano-SIM)": "ডুয়াল সিম (ন্যানো-সিম)",
		"Fingerprint (rear), Accelerometer, Proximity": "ফিঙ্গারপ্রিন্ট (rear), অ্যাক্সিলেরোমিটার, প্রক্সিমিটি",
		"Helio G85 (12 nm)": "হেলিও G৮৫ (১২ ন্যানোমিটার)",
		"Li-Ion (non-removable)": "লি-আয়ন (অপসারণযোগ্য নয়)",
		"Mali-G52 MC2": "মালি-G৫২ MC২",
		"MediaTek Helio G85": "মিডিয়াটেক হেলিও G৮৫",
		"Octa-core (2×2.0 GHz Cortex-A75 + 6×1.8 GHz Cortex-A55)": "অক্টা-কোর (২×২.০ গিগাহার্টজ Cবাtex-A৭৫ + ৬×১.৮ গিগাহার্টজ Cবাtex-A৫৫)",
		"Wi-Fi 802.11 a/b/g/n": "ওয়াই-ফাই ৮০২.১১ a/b/g/n",
	}
}

// Seed inserts specification records for the product identified by slug 'realme-c65'
func (s *SpecificationSeederMobileRealmeC65) Seed(db *gorm.DB) error {
	productSlug := "realme-c65"

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

	// Override model-specific values for Realme C65
	specs["Display Size"] = "6.5 inches"
	specs["Processor"] = "MediaTek Helio G85"
	specs["Chipset"] = "Helio G85 (12 nm)"
	specs["Cpu Type"] = "Octa-core (2×2.0 GHz Cortex-A75 + 6×1.8 GHz Cortex-A55)"
	specs["Gpu Type"] = "Mali-G52 MC2"
	specs["Ram"] = "4 / 6 GB"
	specs["Storage"] = "64 / 128 GB"
	specs["Display Type"] = "IPS LCD"
	specs["Resolution"] = "1600 × 720 px (~270 ppi)"
	specs["Refresh Rate"] = "60 Hz"
	specs["Build Material"] = "Plastic frame/back"
	specs["Weight"] = "185 g"
	specs["Network Technology"] = "GSM / HSPA / LTE"
	specs["Wifi Support"] = "Wi-Fi 802.11 a/b/g/n"
	specs["Bluetooth Version"] = "5.0"
	specs["Nfc Support"] = "No"
	specs["Usb Type"] = "USB-C"
	specs["Rear Camera"] = "50 MP + 2 MP"
	specs["Camera Features"] = "LED flash"
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
	specs["Positioning System"] = "GPS"
	specs["Sensors"] = "Fingerprint (rear), Accelerometer, Proximity"
	specs["Sim Card Type"] = "Dual SIM (Nano-SIM)"
	specs["Loudspeaker Quality"] = "Mono"
	specs["Audio Jack"] = "3.5 mm headphone jack"
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
