package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileSamsungGalaxyA345g seeds specifications/options for product 'samsung-galaxy-a34-5g'
type SpecificationSeederMobileSamsungGalaxyA345g struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSamsungGalaxyA345g creates a new seeder instance
func NewSpecificationSeederMobileSamsungGalaxyA345g() *SpecificationSeederMobileSamsungGalaxyA345g {
	return &SpecificationSeederMobileSamsungGalaxyA345g{BaseSeeder: BaseSeeder{name: "Specifications for Samsung Galaxy A34 5G"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSamsungGalaxyA345g) getBanglaTranslations() map[string]string {
	return map[string]string{
		"120 Hz": "১২০ Hz",
		"128 / 256 GB + microSD (expandable)": "১২৮ / ২৫৬ জিবি + মাইক্রোএসডি (expএবংable)",
		"13 MP": "১৩ মেগাপিক্সেল",
		"161.3 × 78.1 × 8.2 mm": "১৬১.৩ × ৭৮.১ × ৮.২ মিমি",
		"199 g": "১৯৯ g",
		"2340 × 1080 pixels (~390 ppi)": "২৩৪০ × ১০৮০ পিক্সেল (~৩৯০ পিপিআই)",
		"25 W wired": "২৫ W তারযুক্ত",
		"4 / 6 / 8 GB": "৪ / ৬ / ৮ GB",
		"4 OS updates, 5 years security updates": "৪ OS updates, ৫ years security updates",
		"48 MP + 8 MP + 5 MP": "৪৮ মেগাপিক্সেল + ৮ মেগাপিক্সেল + ৫ মেগাপিক্সেল",
		"4K @ 30fps; 1080p @ 30fps": "৪K @ ৩০fps; ১০৮০p @ ৩০fps",
		"4K @ 30fps; 1080p @30/60fps": "৪K @ ৩০fps; ১০৮০p @৩০/৬০fps",
		"5.3": "৫.৩",
		"5000 mAh (typical)": "৫০০০ এমএএইচ (সাধারণ)",
		"6.6 inches": "৬.৬ ইঞ্চি",
		"Android 13, One UI 5.1": "অ্যান্ড্রয়েড ১৩, ওয়ান ইউআই ৫.১",
		"Awesome Lime, Awesome Graphite, Awesome Silver, Awesome Violet": "Awesome Lime, Awesome গ্রাফাইট, Awesome রূপালী, Awesome ভায়োলেট",
		"Dimensity 1080": "ডাইমেনসিটি ১০৮০",
		"Fingerprint (side), Accelerometer, Gyro, Proximity, Compass": "ফিঙ্গারপ্রিন্ট (side), অ্যাক্সিলেরোমিটার, জাইরো, প্রক্সিমিটি, কম্পাস",
		"GSM / HSPA / LTE / 5G": "জিএসএম / এইচএসপিএ / এলটিই / ৫G",
		"Glass front (Gorilla Glass 5), plastic frame & back": "গ্লাস সামনে (গরিলা গ্লাস ৫), প্লাস্টিক ফ্রেম & পেছনে",
		"IP67 (dust / water up to 1m for 30 min)": "IP৬৭ (dust / water পর্যন্ত ১m জন্য ৩০ min)",
		"Li-Ion (non-removable)": "লি-আয়ন (অপসারণযোগ্য নয়)",
		"Mali-G68 MC4": "মালি-G৬৮ MC৪",
		"March 2023": "মার্চ ২০২৩",
		"MediaTek Dimensity 1080": "মিডিয়াটেক ডাইমেনসিটি ১০৮০",
		"Octa-core (2×2.6 GHz Cortex-A78 + 6×2.0 GHz Cortex-A55)": "অক্টা-কোর (২×২.৬ গিগাহার্টজ Cবাtex-A৭৮ + ৬×২.০ গিগাহার্টজ Cবাtex-A৫৫)",
		"Super AMOLED, 120 Hz": "Super অ্যামোলেড, ১২০ Hz",
		"Wi-Fi 802.11 a/b/g/n/ac": "ওয়াই-ফাই ৮০২.১১ a/b/g/n/ac",
		"Yes (market dependent)": "হ্যাঁ (বাজার নির্ভরশীল)",
	}
}

// Seed inserts specification records for the product identified by slug 'samsung-galaxy-a34-5g'
func (s *SpecificationSeederMobileSamsungGalaxyA345g) Seed(db *gorm.DB) error {
	productSlug := "samsung-galaxy-a34-5g"

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

	// Override model-specific values for Samsung Galaxy A34 5G
	specs["Display Size"] = "6.6 inches"
	specs["Processor"] = "MediaTek Dimensity 1080"
	specs["Chipset"] = "Dimensity 1080"
	specs["Cpu Type"] = "Octa-core (2×2.6 GHz Cortex-A78 + 6×2.0 GHz Cortex-A55)"
	specs["Gpu Type"] = "Mali-G68 MC4"
	specs["Ram"] = "4 / 6 / 8 GB"
	specs["Storage"] = "128 / 256 GB + microSD (expandable)"
	specs["Display Type"] = "Super AMOLED, 120 Hz"
	specs["Resolution"] = "2340 × 1080 pixels (~390 ppi)"
	specs["Refresh Rate"] = "120 Hz"
	specs["Build Material"] = "Glass front (Gorilla Glass 5), plastic frame & back"
	specs["Weight"] = "199 g"
	specs["Dimensions"] = "161.3 × 78.1 × 8.2 mm"
	specs["Water Resistance"] = "IP67 (dust / water up to 1m for 30 min)"
	specs["Network Technology"] = "GSM / HSPA / LTE / 5G"
	specs["Wifi Support"] = "Wi-Fi 802.11 a/b/g/n/ac"
	specs["Bluetooth Version"] = "5.3"
	specs["Nfc Support"] = "Yes (market dependent)"
	specs["Usb Type"] = "USB-C"
	specs["Rear Camera"] = "48 MP + 8 MP + 5 MP"
	specs["Camera Features"] = "OIS (main)"
	specs["Camera Video Resolution"] = "4K @ 30fps; 1080p @30/60fps"
	specs["Optical Zoom"] = "None"
	specs["Front Camera"] = "13 MP"
	specs["Front Camera Video Resolution"] = "4K @ 30fps; 1080p @ 30fps"
	specs["Operating System"] = "Android 13, One UI 5.1"
	specs["Battery"] = "5000 mAh (typical)"
	specs["Battery Type"] = "Li-Ion (non-removable)"
	specs["Fast Charging"] = "25 W wired"
	specs["Wireless Charging"] = "No"
	specs["5G Support"] = "Yes"
	specs["Positioning System"] = "GPS, GLONASS, GALILEO, BDS, QZSS"
	specs["Sensors"] = "Fingerprint (side), Accelerometer, Gyro, Proximity, Compass"
	specs["Special Features"] = "4 OS updates, 5 years security updates"
	specs["Sim Card Type"] = "Nano-SIM or Hybrid Dual SIM"
	specs["Loudspeaker Quality"] = "Stereo"
	specs["Audio Quality"] = "Not specified"
	specs["Audio Jack"] = "No"
	specs["Available Colors"] = "Awesome Lime, Awesome Graphite, Awesome Silver, Awesome Violet"
	specs["Announcement Date"] = "March 2023"
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
