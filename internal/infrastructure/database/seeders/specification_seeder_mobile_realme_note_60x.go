package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileRealmeNote60x seeds specifications/options for product 'realme-note-60x'
type SpecificationSeederMobileRealmeNote60x struct {
	BaseSeeder
}

// NewSpecificationSeederMobileRealmeNote60x creates a new seeder instance
func NewSpecificationSeederMobileRealmeNote60x() *SpecificationSeederMobileRealmeNote60x {
	return &SpecificationSeederMobileRealmeNote60x{BaseSeeder: BaseSeeder{name: "Specifications for Realme Note 60x"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileRealmeNote60x) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080p @ 30fps": "১০৮০p @ ৩০fps",
		"190 g": "১৯০ g",
		"2023": "২০২৩",
		"2400 × 1080 px (~405 ppi)": "২৪০০ × ১০৮০ পিক্সেল (~৪০৫ পিপিআই)",
		"3.5 mm headphone jack": "৩.৫ মিমি হেডফোন জ্যাক",
		"33 W wired": "৩৩ W তারযুক্ত",
		"4 / 6 GB": "৪ / ৬ GB",
		"5.1": "৫.১",
		"5000 mAh": "৫০০০ এমএএইচ",
		"6.5 inches": "৬.৫ ইঞ্চি",
		"64 / 128 GB": "৬৪ / ১২৮ GB",
		"64 MP + 2 MP": "৬৪ মেগাপিক্সেল + ২ মেগাপিক্সেল",
		"8 MP": "৮ মেগাপিক্সেল",
		"90 Hz": "৯০ Hz",
		"Android 13, Realme UI 4.0": "অ্যান্ড্রয়েড ১৩, Realme UI ৪.০",
		"Black, Blue": "কালো, নীল",
		"Dimensity 6100+ (6 nm)": "ডাইমেনসিটি ৬১০০+ (৬ ন্যানোমিটার)",
		"Dual SIM (Nano-SIM)": "ডুয়াল সিম (ন্যানো-সিম)",
		"Fingerprint (side), Accelerometer, Gyro, Proximity": "ফিঙ্গারপ্রিন্ট (side), অ্যাক্সিলেরোমিটার, জাইরো, প্রক্সিমিটি",
		"GSM / HSPA / LTE / 5G": "জিএসএম / এইচএসপিএ / এলটিই / ৫G",
		"Li-Ion (non-removable)": "লি-আয়ন (অপসারণযোগ্য নয়)",
		"Mali-G57 MC2": "মালি-G৫৭ MC২",
		"MediaTek Dimensity 6100+": "মিডিয়াটেক ডাইমেনসিটি ৬১০০+",
		"Octa-core (2×2.0 GHz Cortex-A75 + 6×1.8 GHz Cortex-A55)": "অক্টা-কোর (২×২.০ গিগাহার্টজ Cবাtex-A৭৫ + ৬×১.৮ গিগাহার্টজ Cবাtex-A৫৫)",
		"USB-C 2.0": "ইউএসবি-সি ২.০",
		"Wi-Fi 802.11 a/b/g/n/ac": "ওয়াই-ফাই ৮০২.১১ a/b/g/n/ac",
	}
}

// Seed inserts specification records for the product identified by slug 'realme-note-60x'
func (s *SpecificationSeederMobileRealmeNote60x) Seed(db *gorm.DB) error {
	productSlug := "realme-note-60x"

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

	// Override model-specific values for Realme Note 60x
	specs["Display Size"] = "6.5 inches"
	specs["Processor"] = "MediaTek Dimensity 6100+"
	specs["Chipset"] = "Dimensity 6100+ (6 nm)"
	specs["Cpu Type"] = "Octa-core (2×2.0 GHz Cortex-A75 + 6×1.8 GHz Cortex-A55)"
	specs["Gpu Type"] = "Mali-G57 MC2"
	specs["Ram"] = "4 / 6 GB"
	specs["Storage"] = "64 / 128 GB"
	specs["Display Type"] = "IPS LCD"
	specs["Resolution"] = "2400 × 1080 px (~405 ppi)"
	specs["Refresh Rate"] = "90 Hz"
	specs["Build Material"] = "Plastic frame/back"
	specs["Weight"] = "190 g"
	specs["Network Technology"] = "GSM / HSPA / LTE / 5G"
	specs["Wifi Support"] = "Wi-Fi 802.11 a/b/g/n/ac"
	specs["Bluetooth Version"] = "5.1"
	specs["Nfc Support"] = "Yes"
	specs["Usb Type"] = "USB-C 2.0"
	specs["Rear Camera"] = "64 MP + 2 MP"
	specs["Camera Features"] = "LED flash, HDR"
	specs["Camera Video Resolution"] = "1080p @ 30fps"
	specs["Optical Zoom"] = "None"
	specs["Front Camera"] = "8 MP"
	specs["Front Camera Video Resolution"] = "1080p @ 30fps"
	specs["Operating System"] = "Android 13, Realme UI 4.0"
	specs["Battery"] = "5000 mAh"
	specs["Battery Type"] = "Li-Ion (non-removable)"
	specs["Fast Charging"] = "33 W wired"
	specs["Wireless Charging"] = "No"
	specs["5G Support"] = "Yes"
	specs["Positioning System"] = "GPS, GLONASS, GALILEO, BDS"
	specs["Sensors"] = "Fingerprint (side), Accelerometer, Gyro, Proximity"
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
