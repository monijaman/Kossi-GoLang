package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileSamsungGalaxyM065g seeds specifications/options for product 'samsung-galaxy-m06-5g'
type SpecificationSeederMobileSamsungGalaxyM065g struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSamsungGalaxyM065g creates a new seeder instance
func NewSpecificationSeederMobileSamsungGalaxyM065g() *SpecificationSeederMobileSamsungGalaxyM065g {
	return &SpecificationSeederMobileSamsungGalaxyM065g{BaseSeeder: BaseSeeder{name: "Specifications for Samsung Galaxy M06 5G"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSamsungGalaxyM065g) getBanglaTranslations() map[string]string {
	return map[string]string{
		"167.4 × 77.4 × 8 mm": "১৬৭.৪ × ৭৭.৪ × ৮ মিমি",
		"2025": "২০২৫",
		"25 W wired (likely)": "২৫ W তারযুক্ত (likely)",
		"3.5 mm (some sources)": "৩.৫ মিমি (some sources)",
		"4 / 6 GB": "৪ / ৬ GB",
		"5.3": "৫.৩",
		"50 MP + 2 MP": "৫০ মেগাপিক্সেল + ২ মেগাপিক্সেল",
		"5000 mAh": "৫০০০ এমএএইচ",
		"6.74 inches": "৬.৭৪ ইঞ্চি",
		"64 / 128 GB": "৬৪ / ১২৮ GB",
		"720 × 1600 px (reportedly)": "৭২০ × ১৬০০ পিক্সেল (repবাtedly)",
		"8 MP": "৮ মেগাপিক্সেল",
		"90 Hz": "৯০ Hz",
		"Accelerometer, Gyro, Proximity, Compass, Fingerprint (side?)": "অ্যাক্সিলেরোমিটার, জাইরো, প্রক্সিমিটি, কম্পাস, ফিঙ্গারপ্রিন্ট (side?)",
		"Blazing Black, Sage Green": "Blazing কালো, Sage সবুজ",
		"Dimensity 6300 (6 nm)": "ডাইমেনসিটি ৬৩০০ (৬ ন্যানোমিটার)",
		"GSM / HSPA / LTE / 5G": "জিএসএম / এইচএসপিএ / এলটিই / ৫G",
		"Glass front, plastic back": "গ্লাস সামনে, প্লাস্টিক পেছনে",
		"Li-Ion (non-removable)": "লি-আয়ন (অপসারণযোগ্য নয়)",
		"MediaTek Dimensity 6300": "মিডিয়াটেক ডাইমেনসিটি ৬৩০০",
		"Nano-SIM + Nano-SIM or hybrid": "ন্যানো-সিম + ন্যানো-সিম বা hybrid",
		"Wi-Fi 802.11 a/b/g/n/ac": "ওয়াই-ফাই ৮০২.১১ a/b/g/n/ac",
	}
}

// Seed inserts specification records for the product identified by slug 'samsung-galaxy-m06-5g'
func (s *SpecificationSeederMobileSamsungGalaxyM065g) Seed(db *gorm.DB) error {
	productSlug := "samsung-galaxy-m06-5g"

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

	// Override model-specific values for Samsung Galaxy M06 5G
	specs["Display Size"] = "6.74 inches"
	specs["Processor"] = "MediaTek Dimensity 6300"
	specs["Chipset"] = "Dimensity 6300 (6 nm)"
	specs["Cpu Type"] = "Octa-core"
	specs["Ram"] = "4 / 6 GB"
	specs["Storage"] = "64 / 128 GB"
	specs["Display Type"] = "PLS LCD"
	specs["Resolution"] = "720 × 1600 px (reportedly)"
	specs["Refresh Rate"] = "90 Hz"
	specs["Build Material"] = "Glass front, plastic back"
	specs["Dimensions"] = "167.4 × 77.4 × 8 mm"
	specs["Network Technology"] = "GSM / HSPA / LTE / 5G"
	specs["Wifi Support"] = "Wi-Fi 802.11 a/b/g/n/ac"
	specs["Bluetooth Version"] = "5.3"
	specs["Usb Type"] = "USB-C"
	specs["Rear Camera"] = "50 MP + 2 MP"
	specs["Camera Features"] = "LED flash, Depth"
	specs["Optical Zoom"] = "None"
	specs["Front Camera"] = "8 MP"
	specs["Operating System"] = "Android (version not clearly confirmed)"
	specs["Battery"] = "5000 mAh"
	specs["Battery Type"] = "Li-Ion (non-removable)"
	specs["Fast Charging"] = "25 W wired (likely)"
	specs["Wireless Charging"] = "No"
	specs["5G Support"] = "Yes"
	specs["Sensors"] = "Accelerometer, Gyro, Proximity, Compass, Fingerprint (side?)"
	specs["Sim Card Type"] = "Nano-SIM + Nano-SIM or hybrid"
	specs["Loudspeaker Quality"] = "Mono"
	specs["Audio Jack"] = "3.5 mm (some sources)"
	specs["Available Colors"] = "Blazing Black, Sage Green"
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
