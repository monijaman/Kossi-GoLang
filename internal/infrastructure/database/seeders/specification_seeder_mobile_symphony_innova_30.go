package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileSymphonyInnova30 seeds specifications/options for product 'symphony-innova-30'
type SpecificationSeederMobileSymphonyInnova30 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSymphonyInnova30 creates a new seeder instance
func NewSpecificationSeederMobileSymphonyInnova30() *SpecificationSeederMobileSymphonyInnova30 {
	return &SpecificationSeederMobileSymphonyInnova30{BaseSeeder: BaseSeeder{name: "Specifications for Symphony Innova 30"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSymphonyInnova30) getBanglaTranslations() map[string]string {
	return map[string]string{
		"108 MP + 2 MP + 2 MP": "১০৮ মেগাপিক্সেল + ২ মেগাপিক্সেল + ২ মেগাপিক্সেল",
		"1080p@30fps": "১০৮০p@৩০fps",
		"128 GB + microSD": "১২৮ জিবি + মাইক্রোএসডি",
		"164.27 x 76.0 x 8.45 mm": "১৬৪.২৭ x ৭৬.০ x ৮.৪৫ মিমি",
		"18 W wired": "১৮ W তারযুক্ত",
		"193 g": "১৯৩ g",
		"5.0": "৫.০",
		"5000 mAh": "৫০০০ এমএএইচ",
		"6 / 8 GB": "৬ / ৮ GB",
		"6.56 inches": "৬.৫৬ ইঞ্চি",
		"720 × 1612 pixels": "৭২০ × ১৬১২ পিক্সেল",
		"8 MP": "৮ মেগাপিক্সেল",
		"90Hz": "৯০Hz",
		"AI, UHD, Night mode, Portrait, Panorama": "এআই, ইউএইচডি, নাইট মোড, পোর্ট্রেট, প্যানোরামা",
		"Android 13": "অ্যান্ড্রয়েড ১৩",
		"Li-Po (non-removable)": "লি-পো (অপসারণযোগ্য নয়)",
		"Mali-G57 @ 750 MHz": "মালি-G৫৭ @ ৭৫০ মেগাহার্টজ",
		"March 2024": "মার্চ ২০২৪",
		"Mirror White, Reflective Green, Space Green": "মিরর সাদা, রিফ্লেক্টিভ সবুজ, স্পেস সবুজ",
		"Octa-core (2.0 GHz)": "অক্টা-কোর (২.০ গিগাহার্টজ)",
		"Side fingerprint, face unlock, accelerometer, proximity, light, compass": "সাইড ফিঙ্গারপ্রিন্ট, ফেস আনলক, অ্যাক্সিলেরোমিটার, প্রক্সিমিটি, লাইট, কম্পাস",
		"USB Type-C 2.0": "ইউএসবি টাইপ-সি ২.০",
		"UniSOC T616": "UniSOC T৬১৬",
		"UniSOC T616 (12 nm)": "UniSOC T৬১৬ (১২ ন্যানোমিটার)",
		"Wi-Fi 802.11 b/g/n": "ওয়াই-ফাই ৮০২.১১ b/g/n",
		"Yes, 3.5 mm": "হ্যাঁ, ৩.৫ মিমি",
	}
}

// Seed inserts specification records for the product identified by slug 'symphony-innova-30'
func (s *SpecificationSeederMobileSymphonyInnova30) Seed(db *gorm.DB) error {
	productSlug := "symphony-innova-30"

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

	// Override model-specific values for Symphony Innova 30
	specs["Display Size"] = "6.56 inches"
	specs["Processor"] = "UniSOC T616"
	specs["Chipset"] = "UniSOC T616 (12 nm)"
	specs["Cpu Type"] = "Octa-core (2.0 GHz)"
	specs["Gpu Type"] = "Mali-G57 @ 750 MHz"
	specs["Ram"] = "6 / 8 GB"
	specs["Storage"] = "128 GB + microSD"
	specs["Display Type"] = "IPS Incell"
	specs["Resolution"] = "720 × 1612 pixels"
	specs["Refresh Rate"] = "90Hz"
	specs["Build Material"] = "PMMA plastic"
	specs["Weight"] = "193 g"
	specs["Dimensions"] = "164.27 x 76.0 x 8.45 mm"
	specs["Water Resistance"] = "No"
	specs["Network Technology"] = "GSM / HSPA / LTE"
	specs["Wifi Support"] = "Wi-Fi 802.11 b/g/n"
	specs["Bluetooth Version"] = "5.0"
	specs["Nfc Support"] = "No"
	specs["Usb Type"] = "USB Type-C 2.0"
	specs["Rear Camera"] = "108 MP + 2 MP + 2 MP"
	specs["Camera Features"] = "AI, UHD, Night mode, Portrait, Panorama"
	specs["Camera Video Resolution"] = "1080p@30fps"
	specs["Optical Zoom"] = "None"
	specs["Front Camera"] = "8 MP"
	specs["Front Camera Video Resolution"] = "1080p@30fps"
	specs["Operating System"] = "Android 13"
	specs["Battery"] = "5000 mAh"
	specs["Battery Type"] = "Li-Po (non-removable)"
	specs["Fast Charging"] = "18 W wired"
	specs["Wireless Charging"] = "No"
	specs["5G Support"] = "No"
	specs["Positioning System"] = "GPS, A-GPS"
	specs["Sensors"] = "Side fingerprint, face unlock, accelerometer, proximity, light, compass"
	specs["Special Features"] = "Dynamic Island, Game Booster"
	specs["Sim Card Type"] = "Dual Nano-SIM"
	specs["Loudspeaker Quality"] = "Mono"
	specs["Audio Quality"] = "Standard"
	specs["Audio Jack"] = "Yes, 3.5 mm"
	specs["Available Colors"] = "Mirror White, Reflective Green, Space Green"
	specs["Announcement Date"] = "March 2024"
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
