package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileSamsungGalaxyM565g seeds specifications/options for product 'samsung-galaxy-m56-5g'
type SpecificationSeederMobileSamsungGalaxyM565g struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSamsungGalaxyM565g creates a new seeder instance
func NewSpecificationSeederMobileSamsungGalaxyM565g() *SpecificationSeederMobileSamsungGalaxyM565g {
	return &SpecificationSeederMobileSamsungGalaxyM565g{BaseSeeder: BaseSeeder{name: "Specifications for Samsung Galaxy M56 5G"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSamsungGalaxyM565g) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080 × 2340 px": "১০৮০ × ২৩৪০ px",
		"12 MP": "১২ MP",
		"120 Hz": "১২০ Hz",
		"128 / 256 GB + microSDXC": "১২৮ / ২৫৬ GB + microSDXC",
		"162 × 77.3 × 7.2 mm": "১৬২ × ৭৭.৩ × ৭.২ মিমি",
		"180 g": "১৮০ g",
		"45 W wired": "৪৫ W তারযুক্ত",
		"4K @ 30fps; 1080p @ 30/60fps": "৪K @ ৩০fps; ১০৮০p @ ৩০/৬০fps",
		"4K @ 30fps; 1080p @ 30fps": "৪K @ ৩০fps; ১০৮০p @ ৩০fps",
		"5.3": "৫.৩",
		"50 MP + 8 MP + 2 MP": "৫০ MP + ৮ MP + ২ MP",
		"5000 mAh": "৫০০০ এমএএইচ",
		"6.74 inches": "৬.৭৪ ইঞ্চি",
		"8 GB": "৮ GB",
		"Accelerometer, Gyro, Proximity, Compass, fingerprint (side?)": "অ্যাক্সিলেরোমিটার, জাইরো, প্রক্সিমিটি, কম্পাস, ফিঙ্গারপ্রিন্ট (side?)",
		"Android 15, One UI 7": "Android ১৫, One UI ৭",
		"April 2025": "April ২০২৫",
		"Black, Light Green": "কালো, Light সবুজ",
		"Exynos 1480": "Exynos ১৪৮০",
		"Exynos 1480 (4 nm)": "Exynos ১৪৮০ (৪ nm)",
		"GSM / HSPA / LTE / 5G": "GSM / HSPA / LTE / ৫G",
		"Glass front (Gorilla Glass Victus+), plastic frame, glass back": "গ্লাস সামনে (Gorilla Glass Victus+), plastic frame, গ্লাস পেছনে",
		"Li-Po (non-removable)": "লি-পো (অপসারণযোগ্য নয়)",
		"Nano-SIM + Nano-SIM": "ন্যানো-সিম + ন্যানো-সিম",
		"Octa-core (4×2.75 GHz + 4×2.0 GHz)": "Octa-core (৪×২.৭৫ GHz + ৪×২.০ GHz)",
		"Super AMOLED+ 120 Hz": "Super AMOLED+ ১২০ Hz",
		"Wi-Fi 802.11 a/b/g/n/ac/6": "Wi-Fi ৮০২.১১ a/b/g/n/ac/৬",
		"Xclipse 530": "Xclipse ৫৩০",
	}
}

// Seed inserts specification records for the product identified by slug 'samsung-galaxy-m56-5g'
func (s *SpecificationSeederMobileSamsungGalaxyM565g) Seed(db *gorm.DB) error {
	productSlug := "samsung-galaxy-m56-5g"

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

	// Override model-specific values for Samsung Galaxy M56 5G
	specs["Display Size"] = "6.74 inches"
	specs["Processor"] = "Exynos 1480"
	specs["Chipset"] = "Exynos 1480 (4 nm)"
	specs["Cpu Type"] = "Octa-core (4×2.75 GHz + 4×2.0 GHz)"
	specs["Gpu Type"] = "Xclipse 530"
	specs["Ram"] = "8 GB"
	specs["Storage"] = "128 / 256 GB + microSDXC"
	specs["Display Type"] = "Super AMOLED+ 120 Hz"
	specs["Resolution"] = "1080 × 2340 px"
	specs["Refresh Rate"] = "120 Hz"
	specs["Build Material"] = "Glass front (Gorilla Glass Victus+), plastic frame, glass back"
	specs["Weight"] = "180 g"
	specs["Dimensions"] = "162 × 77.3 × 7.2 mm"
	specs["Network Technology"] = "GSM / HSPA / LTE / 5G"
	specs["Wifi Support"] = "Wi-Fi 802.11 a/b/g/n/ac/6"
	specs["Bluetooth Version"] = "5.3"
	specs["Usb Type"] = "USB-C"
	specs["Rear Camera"] = "50 MP + 8 MP + 2 MP"
	specs["Camera Features"] = "OIS, LED flash"
	specs["Camera Video Resolution"] = "4K @ 30fps; 1080p @ 30/60fps"
	specs["Optical Zoom"] = "None"
	specs["Front Camera"] = "12 MP"
	specs["Front Camera Video Resolution"] = "4K @ 30fps; 1080p @ 30fps"
	specs["Operating System"] = "Android 15, One UI 7"
	specs["Battery"] = "5000 mAh"
	specs["Battery Type"] = "Li-Po (non-removable)"
	specs["Fast Charging"] = "45 W wired"
	specs["Wireless Charging"] = "No"
	specs["5G Support"] = "Yes"
	specs["Sensors"] = "Accelerometer, Gyro, Proximity, Compass, fingerprint (side?)"
	specs["Sim Card Type"] = "Nano-SIM + Nano-SIM"
	specs["Loudspeaker Quality"] = "Stereo"
	specs["Audio Jack"] = "No"
	specs["Available Colors"] = "Black, Light Green"
	specs["Announcement Date"] = "April 2025"
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
