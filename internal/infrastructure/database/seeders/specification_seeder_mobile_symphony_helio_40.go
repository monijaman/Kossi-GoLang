package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileSymphonyHelio40 seeds specifications/options for product 'symphony-helio-40'
type SpecificationSeederMobileSymphonyHelio40 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSymphonyHelio40 creates a new seeder instance
func NewSpecificationSeederMobileSymphonyHelio40() *SpecificationSeederMobileSymphonyHelio40 {
	return &SpecificationSeederMobileSymphonyHelio40{BaseSeeder: BaseSeeder{name: "Specifications for Symphony Helio 40"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSymphonyHelio40) getBanglaTranslations() map[string]string {
	return map[string]string{
		"108 MP + 2 MP": "১০৮ মেগাপিক্সেল + ২ মেগাপিক্সেল",
		"1080p@30fps": "১০৮০p@৩০fps",
		"128 GB + microSD": "১২৮ জিবি + মাইক্রোএসডি",
		"16 MP": "১৬ মেগাপিক্সেল",
		"167.4 x 77.6 x 8.54 mm": "১৬৭.৪ x ৭৭.৬ x ৮.৫৪ মিমি",
		"18 W wired": "১৮ W তারযুক্ত",
		"200 g": "২০০ g",
		"2K@30fps, 1080p@30fps": "২K@৩০fps, ১০৮০p@৩০fps",
		"5.0": "৫.০",
		"5000 mAh": "৫০০০ এমএএইচ",
		"6 / 8 GB": "৬ / ৮ GB",
		"6.67 inches": "৬.৬৭ ইঞ্চি",
		"720 × 1604 pixels": "৭২০ × ১৬০৪ পিক্সেল",
		"90Hz": "৯০Hz",
		"Android 14": "অ্যান্ড্রয়েড ১৪",
		"Dual Nano-SIM / Hybrid": "ডুয়াল ন্যানো-সিম / Hybrid",
		"Forest Green, Cosmic Gold, Sunshine Green, Twilight Gold, Sky Blue": "ফরেস্ট সবুজ, কসমিক সোনালী, সানশাইন সবুজ, টোয়াইলাইট সোনালী, স্কাই নীল",
		"Glass front, plastic back": "গ্লাস সামনে, প্লাস্টিক পেছনে",
		"June 2024": "জুন ২০২৪",
		"Li-Po (non-removable)": "লি-পো (অপসারণযোগ্য নয়)",
		"Mali-G52 MC2": "মালি-G৫২ MC২",
		"MediaTek Helio G85": "মিডিয়াটেক হেলিও G৮৫",
		"MediaTek Helio G85 (12 nm)": "মিডিয়াটেক হেলিও G৮৫ (১২ ন্যানোমিটার)",
		"Octa-core, up to 2.0 GHz": "অক্টা-কোর, পর্যন্ত ২.০ গিগাহার্টজ",
		"Side fingerprint, proximity, light, gyroscope": "সাইড ফিঙ্গারপ্রিন্ট, প্রক্সিমিটি, লাইট, gyroscope",
		"USB Type-C 2.0": "ইউএসবি টাইপ-সি ২.০",
		"Wi-Fi 802.11 b/g/n": "ওয়াই-ফাই ৮০২.১১ b/g/n",
		"Yes, 3.5 mm": "হ্যাঁ, ৩.৫ মিমি",
	}
}

// Seed inserts specification records for the product identified by slug 'symphony-helio-40'
func (s *SpecificationSeederMobileSymphonyHelio40) Seed(db *gorm.DB) error {
	productSlug := "symphony-helio-40"

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

	// Override model-specific values for Symphony Helio 40
	specs["Display Size"] = "6.67 inches"
	specs["Processor"] = "MediaTek Helio G85"
	specs["Chipset"] = "MediaTek Helio G85 (12 nm)"
	specs["Cpu Type"] = "Octa-core, up to 2.0 GHz"
	specs["Gpu Type"] = "Mali-G52 MC2"
	specs["Ram"] = "6 / 8 GB"
	specs["Storage"] = "128 GB + microSD"
	specs["Display Type"] = "IPS LCD"
	specs["Resolution"] = "720 × 1604 pixels"
	specs["Refresh Rate"] = "90Hz"
	specs["Build Material"] = "Glass front, plastic back"
	specs["Weight"] = "200 g"
	specs["Dimensions"] = "167.4 x 77.6 x 8.54 mm"
	specs["Water Resistance"] = "No"
	specs["Network Technology"] = "GSM / HSPA / LTE"
	specs["Wifi Support"] = "Wi-Fi 802.11 b/g/n"
	specs["Bluetooth Version"] = "5.0"
	specs["Nfc Support"] = "No"
	specs["Usb Type"] = "USB Type-C 2.0"
	specs["Rear Camera"] = "108 MP + 2 MP"
	specs["Camera Features"] = "AI, Portrait, Pro Mode, Night, HDR, Macro"
	specs["Camera Video Resolution"] = "2K@30fps, 1080p@30fps"
	specs["Optical Zoom"] = "None"
	specs["Front Camera"] = "16 MP"
	specs["Front Camera Video Resolution"] = "1080p@30fps"
	specs["Operating System"] = "Android 14"
	specs["Battery"] = "5000 mAh"
	specs["Battery Type"] = "Li-Po (non-removable)"
	specs["Fast Charging"] = "18 W wired"
	specs["Wireless Charging"] = "No"
	specs["5G Support"] = "No"
	specs["Positioning System"] = "GPS, A-GPS"
	specs["Sensors"] = "Side fingerprint, proximity, light, gyroscope"
	specs["Special Features"] = "Always-on display"
	specs["Sim Card Type"] = "Dual Nano-SIM / Hybrid"
	specs["Loudspeaker Quality"] = "Stereo"
	specs["Audio Quality"] = "Standard"
	specs["Audio Jack"] = "Yes, 3.5 mm"
	specs["Available Colors"] = "Forest Green, Cosmic Gold, Sunshine Green, Twilight Gold, Sky Blue"
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
