package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileSamsungGalaxyF345g seeds specifications/options for product 'samsung-galaxy-f34-5g'
type SpecificationSeederMobileSamsungGalaxyF345g struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSamsungGalaxyF345g creates a new seeder instance
func NewSpecificationSeederMobileSamsungGalaxyF345g() *SpecificationSeederMobileSamsungGalaxyF345g {
	return &SpecificationSeederMobileSamsungGalaxyF345g{BaseSeeder: BaseSeeder{name: "Specifications for Samsung Galaxy F34 5G"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSamsungGalaxyF345g) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080 × 2340 px": "১০৮০ × ২৩৪০ পিক্সেল",
		"1080p @ 30fps": "১০৮০p @ ৩০fps",
		"120 Hz": "১২০ Hz",
		"128 / 256 GB + microSD": "১২৮ / ২৫৬ জিবি + মাইক্রোএসডি",
		"13 MP": "১৩ মেগাপিক্সেল",
		"161 × 78 × 8.2 mm": "১৬১ × ৭৮ × ৮.২ মিমি",
		"199 g": "১৯৯ g",
		"2024": "২০২৪",
		"25 W wired": "২৫ W তারযুক্ত",
		"48 MP + 8 MP + 5 MP": "৪৮ মেগাপিক্সেল + ৮ মেগাপিক্সেল + ৫ মেগাপিক্সেল",
		"4K @ 30fps; 1080p @ 30/60fps": "৪K @ ৩০fps; ১০৮০p @ ৩০/৬০fps",
		"5.3": "৫.৩",
		"5000 mAh": "৫০০০ এমএএইচ",
		"6 / 8 GB": "৬ / ৮ GB",
		"6.6 inches": "৬.৬ ইঞ্চি",
		"Android 13, One UI 5.1": "অ্যান্ড্রয়েড ১৩, ওয়ান ইউআই ৫.১",
		"Dimensity 1080": "ডাইমেনসিটি ১০৮০",
		"GSM / HSPA / LTE / 5G": "জিএসএম / এইচএসপিএ / এলটিই / ৫G",
		"Glass front, plastic back": "গ্লাস সামনে, প্লাস্টিক পেছনে",
		"IP67": "IP৬৭",
		"Li-Ion (non-removable)": "লি-আয়ন (অপসারণযোগ্য নয়)",
		"Mali-G68 MC4": "মালি-G৬৮ MC৪",
		"MediaTek Dimensity 1080": "মিডিয়াটেক ডাইমেনসিটি ১০৮০",
		"Nano-SIM + Nano-SIM": "ন্যানো-সিম + ন্যানো-সিম",
		"Side fingerprint, Accelerometer, Gyro, Compass": "সাইড ফিঙ্গারপ্রিন্ট, অ্যাক্সিলেরোমিটার, জাইরো, কম্পাস",
		"Super AMOLED, 120 Hz": "Super অ্যামোলেড, ১২০ Hz",
		"Wi-Fi 802.11 a/b/g/n/ac": "ওয়াই-ফাই ৮০২.১১ a/b/g/n/ac",
	}
}

// Seed inserts specification records for the product identified by slug 'samsung-galaxy-f34-5g'
func (s *SpecificationSeederMobileSamsungGalaxyF345g) Seed(db *gorm.DB) error {
	productSlug := "samsung-galaxy-f34-5g"

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

	// Override model-specific values for Samsung Galaxy F34 5G
	specs["Display Size"] = "6.6 inches"
	specs["Processor"] = "MediaTek Dimensity 1080"
	specs["Chipset"] = "Dimensity 1080"
	specs["Cpu Type"] = "Octa-core"
	specs["Gpu Type"] = "Mali-G68 MC4"
	specs["Ram"] = "6 / 8 GB"
	specs["Storage"] = "128 / 256 GB + microSD"
	specs["Display Type"] = "Super AMOLED, 120 Hz"
	specs["Resolution"] = "1080 × 2340 px"
	specs["Refresh Rate"] = "120 Hz"
	specs["Build Material"] = "Glass front, plastic back"
	specs["Weight"] = "199 g"
	specs["Dimensions"] = "161 × 78 × 8.2 mm"
	specs["Water Resistance"] = "IP67"
	specs["Network Technology"] = "GSM / HSPA / LTE / 5G"
	specs["Wifi Support"] = "Wi-Fi 802.11 a/b/g/n/ac"
	specs["Bluetooth Version"] = "5.3"
	specs["Nfc Support"] = "Yes"
	specs["Usb Type"] = "USB-C"
	specs["Rear Camera"] = "48 MP + 8 MP + 5 MP"
	specs["Camera Features"] = "OIS"
	specs["Camera Video Resolution"] = "4K @ 30fps; 1080p @ 30/60fps"
	specs["Optical Zoom"] = "None"
	specs["Front Camera"] = "13 MP"
	specs["Front Camera Video Resolution"] = "1080p @ 30fps"
	specs["Operating System"] = "Android 13, One UI 5.1"
	specs["Battery"] = "5000 mAh"
	specs["Battery Type"] = "Li-Ion (non-removable)"
	specs["Fast Charging"] = "25 W wired"
	specs["Wireless Charging"] = "No"
	specs["5G Support"] = "Yes"
	specs["Positioning System"] = "GPS, GLONASS, GALILEO, BDS"
	specs["Sensors"] = "Side fingerprint, Accelerometer, Gyro, Compass"
	specs["Sim Card Type"] = "Nano-SIM + Nano-SIM"
	specs["Loudspeaker Quality"] = "Stereo"
	specs["Audio Jack"] = "No"
	specs["Available Colors"] = "Awesome Lime, Awesome Violet, Awesome Graphite"
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
