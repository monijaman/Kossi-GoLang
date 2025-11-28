package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileSamsungGalaxyF565g seeds specifications/options for product 'samsung-galaxy-f56-5g'
type SpecificationSeederMobileSamsungGalaxyF565g struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSamsungGalaxyF565g creates a new seeder instance
func NewSpecificationSeederMobileSamsungGalaxyF565g() *SpecificationSeederMobileSamsungGalaxyF565g {
	return &SpecificationSeederMobileSamsungGalaxyF565g{BaseSeeder: BaseSeeder{name: "Specifications for Samsung Galaxy F56 5G"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSamsungGalaxyF565g) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080 × 2340 px": "১০৮০ × ২৩৪০ পিক্সেল",
		"1080p @ 30fps": "১০৮০p @ ৩০fps",
		"120 Hz": "১২০ Hz",
		"128 / 256 GB": "১২৮ / ২৫৬ GB",
		"16 MP": "১৬ মেগাপিক্সেল",
		"180 g": "১৮০ g",
		"2025": "২০২৫",
		"25 W wired": "২৫ W তারযুক্ত",
		"4K @ 30fps; 1080p @ 30/60fps": "৪K @ ৩০fps; ১০৮০p @ ৩০/৬০fps",
		"5.3": "৫.৩",
		"50 MP + 8 MP + 2 MP": "৫০ মেগাপিক্সেল + ৮ মেগাপিক্সেল + ২ মেগাপিক্সেল",
		"5000 mAh": "৫০০০ এমএএইচ",
		"6.7 inches": "৬.৭ ইঞ্চি",
		"8 GB": "৮ GB",
		"Android 15, One UI 7": "অ্যান্ড্রয়েড ১৫, ওয়ান ইউআই ৭",
		"Exynos 1480": "এক্সিনস ১৪৮০",
		"GSM / HSPA / LTE / 5G": "জিএসএম / এইচএসপিএ / এলটিই / ৫G",
		"Glass front & back": "গ্লাস সামনে & পেছনে",
		"Graphite, Blue, Violet": "গ্রাফাইট, নীল, ভায়োলেট",
		"Li-Ion (non-removable)": "লি-আয়ন (অপসারণযোগ্য নয়)",
		"Nano-SIM + Nano-SIM": "ন্যানো-সিম + ন্যানো-সিম",
		"Side fingerprint, Accelerometer, Gyro, Compass": "সাইড ফিঙ্গারপ্রিন্ট, অ্যাক্সিলেরোমিটার, জাইরো, কম্পাস",
		"Super AMOLED+, 120 Hz": "Super অ্যামোলেড+, ১২০ Hz",
		"Wi-Fi 802.11 a/b/g/n/ac": "ওয়াই-ফাই ৮০২.১১ a/b/g/n/ac",
	}
}

// Seed inserts specification records for the product identified by slug 'samsung-galaxy-f56-5g'
func (s *SpecificationSeederMobileSamsungGalaxyF565g) Seed(db *gorm.DB) error {
	productSlug := "samsung-galaxy-f56-5g"

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

	// Override model-specific values for Samsung Galaxy F56 5G
	specs["Display Size"] = "6.7 inches"
	specs["Processor"] = "Exynos 1480"
	specs["Chipset"] = "Exynos 1480"
	specs["Ram"] = "8 GB"
	specs["Storage"] = "128 / 256 GB"
	specs["Display Type"] = "Super AMOLED+, 120 Hz"
	specs["Resolution"] = "1080 × 2340 px"
	specs["Refresh Rate"] = "120 Hz"
	specs["Build Material"] = "Glass front & back"
	specs["Weight"] = "180 g"
	specs["Dimensions"] = "Not specified"
	specs["Network Technology"] = "GSM / HSPA / LTE / 5G"
	specs["Wifi Support"] = "Wi-Fi 802.11 a/b/g/n/ac"
	specs["Bluetooth Version"] = "5.3"
	specs["Nfc Support"] = "Yes"
	specs["Usb Type"] = "USB-C"
	specs["Rear Camera"] = "50 MP + 8 MP + 2 MP"
	specs["Camera Features"] = "OIS, LED flash"
	specs["Camera Video Resolution"] = "4K @ 30fps; 1080p @ 30/60fps"
	specs["Optical Zoom"] = "None"
	specs["Front Camera"] = "16 MP"
	specs["Front Camera Video Resolution"] = "1080p @ 30fps"
	specs["Operating System"] = "Android 15, One UI 7"
	specs["Battery"] = "5000 mAh"
	specs["Battery Type"] = "Li-Ion (non-removable)"
	specs["Fast Charging"] = "25 W wired"
	specs["Wireless Charging"] = "No"
	specs["5G Support"] = "Yes"
	specs["Positioning System"] = "GPS, GLONASS, GALILEO, BDS"
	specs["Sensors"] = "Side fingerprint, Accelerometer, Gyro, Compass"
	specs["Sim Card Type"] = "Nano-SIM + Nano-SIM"
	specs["Loudspeaker Quality"] = "Yes"
	specs["Audio Jack"] = "No"
	specs["Available Colors"] = "Graphite, Blue, Violet"
	specs["Announcement Date"] = "2025"
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
