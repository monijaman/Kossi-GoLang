package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileSamsungGalaxyF165g seeds specifications/options for product 'samsung-galaxy-f16-5g'
type SpecificationSeederMobileSamsungGalaxyF165g struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSamsungGalaxyF165g creates a new seeder instance
func NewSpecificationSeederMobileSamsungGalaxyF165g() *SpecificationSeederMobileSamsungGalaxyF165g {
	return &SpecificationSeederMobileSamsungGalaxyF165g{BaseSeeder: BaseSeeder{name: "Specifications for Samsung Galaxy F16 5G"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSamsungGalaxyF165g) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080 × 2408 px": "১০৮০ × ২৪০৮ পিক্সেল",
		"1080p @ 30fps": "১০৮০p @ ৩০fps",
		"15 W wired": "১৫ W তারযুক্ত",
		"164 × 76 × 8.4 mm": "১৬৪ × ৭৬ × ৮.৪ মিমি",
		"190 g": "১৯০ g",
		"2024": "২০২৪",
		"4 / 6 GB": "৪ / ৬ GB",
		"5.1": "৫.১",
		"50 MP + 2 MP + 2 MP": "৫০ মেগাপিক্সেল + ২ মেগাপিক্সেল + ২ মেগাপিক্সেল",
		"5000 mAh": "৫০০০ এমএএইচ",
		"6.6 inches": "৬.৬ ইঞ্চি",
		"64 / 128 GB + microSD": "৬৪ / ১২৮ জিবি + মাইক্রোএসডি",
		"8 MP": "৮ মেগাপিক্সেল",
		"90 Hz": "৯০ Hz",
		"Accelerometer, Proximity, Compass, Side fingerprint": "অ্যাক্সিলেরোমিটার, প্রক্সিমিটি, কম্পাস, সাইড ফিঙ্গারপ্রিন্ট",
		"Android 13, One UI Core 5": "অ্যান্ড্রয়েড ১৩, ওয়ান ইউআই Cবাe ৫",
		"Black, Blue, Green": "কালো, নীল, সবুজ",
		"Dimensity 6100+": "ডাইমেনসিটি ৬১০০+",
		"GSM / HSPA / LTE / 5G": "জিএসএম / এইচএসপিএ / এলটিই / ৫G",
		"Li-Ion (non-removable)": "লি-আয়ন (অপসারণযোগ্য নয়)",
		"Mali-G57 MC2": "মালি-G৫৭ MC২",
		"MediaTek Dimensity 6100+": "মিডিয়াটেক ডাইমেনসিটি ৬১০০+",
		"Nano-SIM + Nano-SIM": "ন্যানো-সিম + ন্যানো-সিম",
		"Octa-core (2×2.0 GHz + 6×1.8 GHz)": "অক্টা-কোর (২×২.০ গিগাহার্টজ + ৬×১.৮ গিগাহার্টজ)",
		"PLS LCD, 90 Hz": "PLS এলসিডি, ৯০ Hz",
		"Plastic frame & back, glass front": "Plastic ফ্রেম & পেছনে, গ্লাস সামনে",
		"USB-C 2.0": "ইউএসবি-সি ২.০",
		"Wi-Fi 802.11 a/b/g/n/ac": "ওয়াই-ফাই ৮০২.১১ a/b/g/n/ac",
		"Yes (market dependent)": "হ্যাঁ (বাজার নির্ভরশীল)",
		"Yes, 3.5 mm": "হ্যাঁ, ৩.৫ মিমি",
	}
}

// Seed inserts specification records for the product identified by slug 'samsung-galaxy-f16-5g'
func (s *SpecificationSeederMobileSamsungGalaxyF165g) Seed(db *gorm.DB) error {
	productSlug := "samsung-galaxy-f16-5g"

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

	// Override model-specific values for Samsung Galaxy F16 5G
	specs["Display Size"] = "6.6 inches"
	specs["Processor"] = "MediaTek Dimensity 6100+"
	specs["Chipset"] = "Dimensity 6100+"
	specs["Cpu Type"] = "Octa-core (2×2.0 GHz + 6×1.8 GHz)"
	specs["Gpu Type"] = "Mali-G57 MC2"
	specs["Ram"] = "4 / 6 GB"
	specs["Storage"] = "64 / 128 GB + microSD"
	specs["Display Type"] = "PLS LCD, 90 Hz"
	specs["Resolution"] = "1080 × 2408 px"
	specs["Refresh Rate"] = "90 Hz"
	specs["Build Material"] = "Plastic frame & back, glass front"
	specs["Weight"] = "190 g"
	specs["Dimensions"] = "164 × 76 × 8.4 mm"
	specs["Network Technology"] = "GSM / HSPA / LTE / 5G"
	specs["Wifi Support"] = "Wi-Fi 802.11 a/b/g/n/ac"
	specs["Bluetooth Version"] = "5.1"
	specs["Nfc Support"] = "Yes (market dependent)"
	specs["Usb Type"] = "USB-C 2.0"
	specs["Rear Camera"] = "50 MP + 2 MP + 2 MP"
	specs["Camera Features"] = "LED flash"
	specs["Camera Video Resolution"] = "1080p @ 30fps"
	specs["Optical Zoom"] = "None"
	specs["Front Camera"] = "8 MP"
	specs["Front Camera Video Resolution"] = "1080p @ 30fps"
	specs["Operating System"] = "Android 13, One UI Core 5"
	specs["Battery"] = "5000 mAh"
	specs["Battery Type"] = "Li-Ion (non-removable)"
	specs["Fast Charging"] = "15 W wired"
	specs["Wireless Charging"] = "No"
	specs["5G Support"] = "Yes"
	specs["Positioning System"] = "GPS, GLONASS, GALILEO, BDS"
	specs["Sensors"] = "Accelerometer, Proximity, Compass, Side fingerprint"
	specs["Sim Card Type"] = "Nano-SIM + Nano-SIM"
	specs["Loudspeaker Quality"] = "Mono"
	specs["Audio Jack"] = "Yes, 3.5 mm"
	specs["Available Colors"] = "Black, Blue, Green"
	specs["Announcement Date"] = "2024"
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
