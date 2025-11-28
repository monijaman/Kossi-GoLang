package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileSymphonyHelio50 seeds specifications/options for product 'symphony-helio-50'
type SpecificationSeederMobileSymphonyHelio50 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSymphonyHelio50 creates a new seeder instance
func NewSpecificationSeederMobileSymphonyHelio50() *SpecificationSeederMobileSymphonyHelio50 {
	return &SpecificationSeederMobileSymphonyHelio50{BaseSeeder: BaseSeeder{name: "Specifications for Symphony Helio 50"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSymphonyHelio50) getBanglaTranslations() map[string]string {
	return map[string]string{
		"108 MP + 2 MP + 0.8 MP": "১০৮ মেগাপিক্সেল + ২ মেগাপিক্সেল + ০.৮ মেগাপিক্সেল",
		"1080 × 2400 pixels": "১০৮০ × ২৪০০ পিক্সেল",
		"1080p@30fps": "১০৮০p@৩০fps",
		"128 GB + microSD": "১২৮ জিবি + মাইক্রোএসডি",
		"167 x 77.6 x 8.79 mm": "১৬৭ x ৭৭.৬ x ৮.৭৯ মিমি",
		"215 g": "২১৫ g",
		"2K@30fps, 1080p@30fps": "২K@৩০fps, ১০৮০p@৩০fps",
		"32 MP": "৩২ মেগাপিক্সেল",
		"33 W wired": "৩৩ W তারযুক্ত",
		"5.x": "৫.x",
		"5000 mAh": "৫০০০ এমএএইচ",
		"6 GB": "৬ GB",
		"6.72 inches": "৬.৭২ ইঞ্চি",
		"90Hz": "৯০Hz",
		"Android 13": "অ্যান্ড্রয়েড ১৩",
		"Dual Nano-SIM / Hybrid": "ডুয়াল ন্যানো-সিম / Hybrid",
		"Glass front, plastic back": "গ্লাস সামনে, প্লাস্টিক পেছনে",
		"June 2024": "জুন ২০২৪",
		"Li-Po (non-removable)": "লি-পো (অপসারণযোগ্য নয়)",
		"Mali-G52": "মালি-G৫২",
		"MediaTek Helio G88": "মিডিয়াটেক হেলিও G৮৮",
		"MediaTek Helio G88 (12 nm)": "মিডিয়াটেক হেলিও G৮৮ (১২ ন্যানোমিটার)",
		"Octa-core, up to 2.0 GHz": "অক্টা-কোর, পর্যন্ত ২.০ গিগাহার্টজ",
		"Radium Green, Cosmic Gold, Honey Dew Green": "রেডিয়াম সবুজ, কসমিক সোনালী, Honey Dew সবুজ",
		"Side fingerprint, proximity, light, gyroscope": "সাইড ফিঙ্গারপ্রিন্ট, প্রক্সিমিটি, লাইট, gyroscope",
		"USB Type-C 2.0": "ইউএসবি টাইপ-সি ২.০",
		"Wi-Fi 802.11 b/g/n": "ওয়াই-ফাই ৮০২.১১ b/g/n",
		"Yes, 3.5 mm": "হ্যাঁ, ৩.৫ মিমি",
	}
}

// Seed inserts specification records for the product identified by slug 'symphony-helio-50'
func (s *SpecificationSeederMobileSymphonyHelio50) Seed(db *gorm.DB) error {
	productSlug := "symphony-helio-50"

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

	// Override model-specific values for Symphony Helio 50
	specs["Display Size"] = "6.72 inches"
	specs["Processor"] = "MediaTek Helio G88"
	specs["Chipset"] = "MediaTek Helio G88 (12 nm)"
	specs["Cpu Type"] = "Octa-core, up to 2.0 GHz"
	specs["Gpu Type"] = "Mali-G52"
	specs["Ram"] = "6 GB"
	specs["Storage"] = "128 GB + microSD"
	specs["Display Type"] = "IPS LCD"
	specs["Resolution"] = "1080 × 2400 pixels"
	specs["Refresh Rate"] = "90Hz"
	specs["Build Material"] = "Glass front, plastic back"
	specs["Weight"] = "215 g"
	specs["Dimensions"] = "167 x 77.6 x 8.79 mm"
	specs["Water Resistance"] = "No"
	specs["Network Technology"] = "GSM / HSPA / LTE"
	specs["Wifi Support"] = "Wi-Fi 802.11 b/g/n"
	specs["Bluetooth Version"] = "5.x"
	specs["Nfc Support"] = "No"
	specs["Usb Type"] = "USB Type-C 2.0"
	specs["Rear Camera"] = "108 MP + 2 MP + 0.8 MP"
	specs["Camera Features"] = "AI, RAW, Slow motion, Portrait, HDR"
	specs["Camera Video Resolution"] = "2K@30fps, 1080p@30fps"
	specs["Optical Zoom"] = "None"
	specs["Front Camera"] = "32 MP"
	specs["Front Camera Video Resolution"] = "1080p@30fps"
	specs["Operating System"] = "Android 13"
	specs["Battery"] = "5000 mAh"
	specs["Battery Type"] = "Li-Po (non-removable)"
	specs["Fast Charging"] = "33 W wired"
	specs["Wireless Charging"] = "No"
	specs["5G Support"] = "No"
	specs["Positioning System"] = "GPS, A-GPS"
	specs["Sensors"] = "Side fingerprint, proximity, light, gyroscope"
	specs["Special Features"] = "Dual speakers"
	specs["Sim Card Type"] = "Dual Nano-SIM / Hybrid"
	specs["Loudspeaker Quality"] = "Stereo"
	specs["Audio Quality"] = "Standard"
	specs["Audio Jack"] = "Yes, 3.5 mm"
	specs["Available Colors"] = "Radium Green, Cosmic Gold, Honey Dew Green"
	specs["Announcement Date"] = "June 2024"
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
