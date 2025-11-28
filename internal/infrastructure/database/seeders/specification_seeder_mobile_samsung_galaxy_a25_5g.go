package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileSamsungGalaxyA255g seeds specifications/options for product 'samsung-galaxy-a25-5g'
type SpecificationSeederMobileSamsungGalaxyA255g struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSamsungGalaxyA255g creates a new seeder instance
func NewSpecificationSeederMobileSamsungGalaxyA255g() *SpecificationSeederMobileSamsungGalaxyA255g {
	return &SpecificationSeederMobileSamsungGalaxyA255g{BaseSeeder: BaseSeeder{name: "Specifications for Samsung Galaxy A25 5G"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSamsungGalaxyA255g) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080p @ 30fps": "১০৮০p @ ৩০fps",
		"120 Hz": "১২০ Hz",
		"128 / 256 GB + microSD (up to 1 TB)": "১২৮ / ২৫৬ জিবি + মাইক্রোএসডি (পর্যন্ত ১ টিবি)",
		"13 MP": "১৩ মেগাপিক্সেল",
		"161 × 76.5 × 8.3 mm": "১৬১ × ৭৬.৫ × ৮.৩ মিমি",
		"197 g": "১৯৭ g",
		"2340 × 1080 pixels (~396 ppi)": "২৩৪০ × ১০৮০ পিক্সেল (~৩৯৬ পিপিআই)",
		"25 W wired": "২৫ W তারযুক্ত",
		"4K @ 30fps; 1080p @ 30/60fps": "৪K @ ৩০fps; ১০৮০p @ ৩০/৬০fps",
		"5.3": "৫.৩",
		"50 MP + 8 MP + 2 MP": "৫০ মেগাপিক্সেল + ৮ মেগাপিক্সেল + ২ মেগাপিক্সেল",
		"5000 mAh": "৫০০০ এমএএইচ",
		"6 / 8 GB": "৬ / ৮ GB",
		"6.5 inches": "৬.৫ ইঞ্চি",
		"Accelerometer, Gyro, Proximity, Compass": "অ্যাক্সিলেরোমিটার, জাইরো, প্রক্সিমিটি, কম্পাস",
		"Android 14, One UI 6": "অ্যান্ড্রয়েড ১৪, ওয়ান ইউআই ৬",
		"Blue Black, Blue, Light Blue, Yellow": "নীল কালো, নীল, Light নীল, হলুদ",
		"December 2023": "ডিসেম্বর ২০২৩",
		"Exynos 1280": "এক্সিনস ১২৮০",
		"Exynos 1280 (5 nm)": "এক্সিনস ১২৮০ (৫ ন্যানোমিটার)",
		"GSM / HSPA / LTE / 5G": "জিএসএম / এইচএসপিএ / এলটিই / ৫G",
		"Glass front, plastic frame / back": "গ্লাস সামনে, প্লাস্টিক ফ্রেম / পেছনে",
		"Li-Ion (non-removable)": "লি-আয়ন (অপসারণযোগ্য নয়)",
		"Mali-G68": "মালি-G৬৮",
		"Nano-SIM (or hybrid)": "ন্যানো-সিম (বা hybrid)",
		"Octa-core (2×2.4 GHz Cortex-A78 + 6×2.0 GHz Cortex-A55)": "অক্টা-কোর (২×২.৪ গিগাহার্টজ Cবাtex-A৭৮ + ৬×২.০ গিগাহার্টজ Cবাtex-A৫৫)",
		"Super AMOLED, 120 Hz": "Super অ্যামোলেড, ১২০ Hz",
		"USB-C 2.0": "ইউএসবি-সি ২.০",
		"Wi-Fi 802.11 a/b/g/n/ac": "ওয়াই-ফাই ৮০২.১১ a/b/g/n/ac",
	}
}

// Seed inserts specification records for the product identified by slug 'samsung-galaxy-a25-5g'
func (s *SpecificationSeederMobileSamsungGalaxyA255g) Seed(db *gorm.DB) error {
	productSlug := "samsung-galaxy-a25-5g"

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

	// Override model-specific values for Samsung Galaxy A25 5G
	specs["Display Size"] = "6.5 inches"
	specs["Processor"] = "Exynos 1280"
	specs["Chipset"] = "Exynos 1280 (5 nm)"
	specs["Cpu Type"] = "Octa-core (2×2.4 GHz Cortex-A78 + 6×2.0 GHz Cortex-A55)"
	specs["Gpu Type"] = "Mali-G68"
	specs["Ram"] = "6 / 8 GB"
	specs["Storage"] = "128 / 256 GB + microSD (up to 1 TB)"
	specs["Display Type"] = "Super AMOLED, 120 Hz"
	specs["Resolution"] = "2340 × 1080 pixels (~396 ppi)"
	specs["Refresh Rate"] = "120 Hz"
	specs["Build Material"] = "Glass front, plastic frame / back"
	specs["Weight"] = "197 g"
	specs["Dimensions"] = "161 × 76.5 × 8.3 mm"
	specs["Water Resistance"] = "None"
	specs["Network Technology"] = "GSM / HSPA / LTE / 5G"
	specs["Wifi Support"] = "Wi-Fi 802.11 a/b/g/n/ac"
	specs["Bluetooth Version"] = "5.3"
	specs["Nfc Support"] = "Yes"
	specs["Usb Type"] = "USB-C 2.0"
	specs["Rear Camera"] = "50 MP + 8 MP + 2 MP"
	specs["Camera Features"] = "OIS (main), LED flash"
	specs["Camera Video Resolution"] = "4K @ 30fps; 1080p @ 30/60fps"
	specs["Optical Zoom"] = "None"
	specs["Front Camera"] = "13 MP"
	specs["Front Camera Video Resolution"] = "1080p @ 30fps"
	specs["Operating System"] = "Android 14, One UI 6"
	specs["Battery"] = "5000 mAh"
	specs["Battery Type"] = "Li-Ion (non-removable)"
	specs["Fast Charging"] = "25 W wired"
	specs["Wireless Charging"] = "No"
	specs["5G Support"] = "Yes"
	specs["Positioning System"] = "GPS, GLONASS, Galileo, BDS"
	specs["Sensors"] = "Accelerometer, Gyro, Proximity, Compass"
	specs["Special Features"] = "None"
	specs["Sim Card Type"] = "Nano-SIM (or hybrid)"
	specs["Loudspeaker Quality"] = "Stereo"
	specs["Audio Quality"] = "Not specified"
	specs["Audio Jack"] = "No"
	specs["Available Colors"] = "Blue Black, Blue, Light Blue, Yellow"
	specs["Announcement Date"] = "December 2023"
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
