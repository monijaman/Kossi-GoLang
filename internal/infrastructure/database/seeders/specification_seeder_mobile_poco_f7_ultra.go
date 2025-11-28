package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobilePocoF7Ultra seeds specifications/options for product 'poco-f7-ultra'
type SpecificationSeederMobilePocoF7Ultra struct {
	BaseSeeder
}

// NewSpecificationSeederMobilePocoF7Ultra creates a new seeder instance
func NewSpecificationSeederMobilePocoF7Ultra() *SpecificationSeederMobilePocoF7Ultra {
	return &SpecificationSeederMobilePocoF7Ultra{BaseSeeder: BaseSeeder{name: "Specifications for POCO F7 Ultra"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobilePocoF7Ultra) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080p @ 60fps": "১০৮০p @ ৬০fps",
		"120 Hz": "১২০ Hz",
		"120 W wired": "১২০ W তারযুক্ত",
		"128 / 256 GB": "১২৮ / ২৫৬ GB",
		"200 MP + 13 MP + 5 MP": "২০০ মেগাপিক্সেল + ১৩ মেগাপিক্সেল + ৫ মেগাপিক্সেল",
		"2024": "২০২৪",
		"210 g": "২১০ g",
		"32 MP": "৩২ মেগাপিক্সেল",
		"3200 × 1440 px": "৩২০০ × ১৪৪০ পিক্সেল",
		"5.3": "৫.৩",
		"5000 mAh": "৫০০০ এমএএইচ",
		"5x": "৫x",
		"6.67 inches": "৬.৬৭ ইঞ্চি",
		"8 / 12 GB": "৮ / ১২ GB",
		"8K @ 30fps; 4K @ 60fps; 1080p @ 120fps": "৮K @ ৩০fps; ৪K @ ৬০fps; ১০৮০p @ ১২০fps",
		"AMOLED, 120 Hz": "অ্যামোলেড, ১২০ Hz",
		"Android 13, MIUI 14": "অ্যান্ড্রয়েড ১৩, এমআইইউআই ১৪",
		"Black, Blue": "কালো, নীল",
		"Dimensity 9200+ (4 nm)": "ডাইমেনসিটি ৯২০০+ (৪ ন্যানোমিটার)",
		"Fingerprint (under-display), Accelerometer, Gyro, Compass": "ফিঙ্গারপ্রিন্ট (under-display), অ্যাক্সিলেরোমিটার, জাইরো, কম্পাস",
		"GSM / HSPA / LTE / 5G": "জিএসএম / এইচএসপিএ / এলটিই / ৫G",
		"Glass front/back, aluminum frame": "গ্লাস সামনে/পেছনে, অ্যালুমিনিয়াম ফ্রেম",
		"Li-Ion (non-removable)": "লি-আয়ন (অপসারণযোগ্য নয়)",
		"Mali-G710 MC10": "মালি-G৭১০ MC১০",
		"MediaTek Dimensity 9200+": "মিডিয়াটেক ডাইমেনসিটি ৯২০০+",
		"Nano-SIM + Nano-SIM": "ন্যানো-সিম + ন্যানো-সিম",
		"USB-C 3.1": "ইউএসবি-সি ৩.১",
		"Wi-Fi 802.11 a/b/g/n/ac/6e": "ওয়াই-ফাই ৮০২.১১ a/b/g/n/ac/৬e",
	}
}

// Seed inserts specification records for the product identified by slug 'poco-f7-ultra'
func (s *SpecificationSeederMobilePocoF7Ultra) Seed(db *gorm.DB) error {
	productSlug := "poco-f7-ultra"

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

	// Override model-specific values for POCO F7 Ultra
	specs["Display Size"] = "6.67 inches"
	specs["Processor"] = "MediaTek Dimensity 9200+"
	specs["Chipset"] = "Dimensity 9200+ (4 nm)"
	specs["Cpu Type"] = "Octa-core"
	specs["Gpu Type"] = "Mali-G710 MC10"
	specs["Ram"] = "8 / 12 GB"
	specs["Storage"] = "128 / 256 GB"
	specs["Display Type"] = "AMOLED, 120 Hz"
	specs["Resolution"] = "3200 × 1440 px"
	specs["Refresh Rate"] = "120 Hz"
	specs["Build Material"] = "Glass front/back, aluminum frame"
	specs["Weight"] = "210 g"
	specs["Network Technology"] = "GSM / HSPA / LTE / 5G"
	specs["Wifi Support"] = "Wi-Fi 802.11 a/b/g/n/ac/6e"
	specs["Bluetooth Version"] = "5.3"
	specs["Nfc Support"] = "Yes"
	specs["Usb Type"] = "USB-C 3.1"
	specs["Rear Camera"] = "200 MP + 13 MP + 5 MP"
	specs["Camera Features"] = "OIS, LED flash, HDR"
	specs["Camera Video Resolution"] = "8K @ 30fps; 4K @ 60fps; 1080p @ 120fps"
	specs["Optical Zoom"] = "5x"
	specs["Front Camera"] = "32 MP"
	specs["Front Camera Video Resolution"] = "1080p @ 60fps"
	specs["Operating System"] = "Android 13, MIUI 14"
	specs["Battery"] = "5000 mAh"
	specs["Battery Type"] = "Li-Ion (non-removable)"
	specs["Fast Charging"] = "120 W wired"
	specs["Wireless Charging"] = "No"
	specs["5G Support"] = "Yes"
	specs["Positioning System"] = "GPS, GLONASS, GALILEO, BDS, QZSS"
	specs["Sensors"] = "Fingerprint (under-display), Accelerometer, Gyro, Compass"
	specs["Sim Card Type"] = "Nano-SIM + Nano-SIM"
	specs["Loudspeaker Quality"] = "Stereo"
	specs["Audio Jack"] = "No"
	specs["Available Colors"] = "Black, Blue"
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
