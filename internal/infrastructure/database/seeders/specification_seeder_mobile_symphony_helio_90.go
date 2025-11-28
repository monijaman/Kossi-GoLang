package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileSymphonyHelio90 seeds specifications/options for product 'symphony-helio-90'
type SpecificationSeederMobileSymphonyHelio90 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSymphonyHelio90 creates a new seeder instance
func NewSpecificationSeederMobileSymphonyHelio90() *SpecificationSeederMobileSymphonyHelio90 {
	return &SpecificationSeederMobileSymphonyHelio90{BaseSeeder: BaseSeeder{name: "Specifications for Symphony Helio 90"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSymphonyHelio90) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080 × 2412 pixels": "১০৮০ × ২৪১২ পিক্সেল",
		"1080p@30fps": "১০৮০p@৩০fps",
		"120Hz": "১২০Hz",
		"128 GB + microSD": "১২৮ জিবি + মাইক্রোএসডি",
		"163.8 x 76.3 x 7.96 mm": "১৬৩.৮ x ৭৬.৩ x ৭.৯৬ মিমি",
		"188 g": "১৮৮ g",
		"2K@30fps, 1080p@30fps": "২K@৩০fps, ১০৮০p@৩০fps",
		"32 MP": "৩২ মেগাপিক্সেল",
		"33 W wired": "৩৩ W তারযুক্ত",
		"5.x": "৫.x",
		"5000 mAh": "৫০০০ এমএএইচ",
		"6.7 inches": "৬.৭ ইঞ্চি",
		"64 MP + 2 MP + 0.8 MP": "৬৪ মেগাপিক্সেল + ২ মেগাপিক্সেল + ০.৮ মেগাপিক্সেল",
		"8 GB": "৮ GB",
		"Android 14": "অ্যান্ড্রয়েড ১৪",
		"Dual Nano-SIM / Hybrid": "ডুয়াল ন্যানো-সিম / Hybrid",
		"Glass front, plastic back": "গ্লাস সামনে, প্লাস্টিক পেছনে",
		"In-display fingerprint, proximity, light, gyroscope": "In-display ফিঙ্গারপ্রিন্ট, প্রক্সিমিটি, লাইট, gyroscope",
		"Li-Po (non-removable)": "লি-পো (অপসারণযোগ্য নয়)",
		"Mali-G57 MC2": "মালি-G৫৭ MC২",
		"MediaTek Helio G99": "মিডিয়াটেক হেলিও G৯৯",
		"MediaTek Helio G99 (6 nm)": "মিডিয়াটেক হেলিও G৯৯ (৬ ন্যানোমিটার)",
		"Octa-core, up to 2.2 GHz": "অক্টা-কোর, পর্যন্ত ২.২ গিগাহার্টজ",
		"September 2024": "সেপ্টেম্বর ২০২৪",
		"Space Black, Thunder White, Cosmic Gold": "স্পেস কালো, থান্ডার সাদা, কসমিক সোনালী",
		"USB Type-C 2.0": "ইউএসবি টাইপ-সি ২.০",
		"Wi-Fi 2.4 / 5 GHz": "ওয়াই-ফাই ২.৪ / ৫ গিগাহার্টজ",
		"Yes, 3.5 mm": "হ্যাঁ, ৩.৫ মিমি",
	}
}

// Seed inserts specification records for the product identified by slug 'symphony-helio-90'
func (s *SpecificationSeederMobileSymphonyHelio90) Seed(db *gorm.DB) error {
	productSlug := "symphony-helio-90"

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

	// Override model-specific values for Symphony Helio 90
	specs["Display Size"] = "6.7 inches"
	specs["Processor"] = "MediaTek Helio G99"
	specs["Chipset"] = "MediaTek Helio G99 (6 nm)"
	specs["Cpu Type"] = "Octa-core, up to 2.2 GHz"
	specs["Gpu Type"] = "Mali-G57 MC2"
	specs["Ram"] = "8 GB"
	specs["Storage"] = "128 GB + microSD"
	specs["Display Type"] = "AMOLED"
	specs["Resolution"] = "1080 × 2412 pixels"
	specs["Refresh Rate"] = "120Hz"
	specs["Build Material"] = "Glass front, plastic back"
	specs["Weight"] = "188 g"
	specs["Dimensions"] = "163.8 x 76.3 x 7.96 mm"
	specs["Water Resistance"] = "No"
	specs["Network Technology"] = "GSM / HSPA / LTE"
	specs["Wifi Support"] = "Wi-Fi 2.4 / 5 GHz"
	specs["Bluetooth Version"] = "5.x"
	specs["Nfc Support"] = "No"
	specs["Usb Type"] = "USB Type-C 2.0"
	specs["Rear Camera"] = "64 MP + 2 MP + 0.8 MP"
	specs["Camera Features"] = "AI, Portrait, Macro, Night mode, EIS"
	specs["Camera Video Resolution"] = "2K@30fps, 1080p@30fps"
	specs["Optical Zoom"] = "None"
	specs["Front Camera"] = "32 MP"
	specs["Front Camera Video Resolution"] = "1080p@30fps"
	specs["Operating System"] = "Android 14"
	specs["Battery"] = "5000 mAh"
	specs["Battery Type"] = "Li-Po (non-removable)"
	specs["Fast Charging"] = "33 W wired"
	specs["Wireless Charging"] = "No"
	specs["5G Support"] = "No"
	specs["Positioning System"] = "GPS, GLONASS, BDS"
	specs["Sensors"] = "In-display fingerprint, proximity, light, gyroscope"
	specs["Special Features"] = "Always-on display"
	specs["Sim Card Type"] = "Dual Nano-SIM / Hybrid"
	specs["Loudspeaker Quality"] = "Stereo"
	specs["Audio Quality"] = "Standard"
	specs["Audio Jack"] = "Yes, 3.5 mm"
	specs["Available Colors"] = "Space Black, Thunder White, Cosmic Gold"
	specs["Announcement Date"] = "September 2024"
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
