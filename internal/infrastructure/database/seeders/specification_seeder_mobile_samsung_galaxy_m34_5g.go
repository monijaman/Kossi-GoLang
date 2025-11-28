package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileSamsungGalaxyM345g seeds specifications/options for product 'samsung-galaxy-m34-5g'
type SpecificationSeederMobileSamsungGalaxyM345g struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSamsungGalaxyM345g creates a new seeder instance
func NewSpecificationSeederMobileSamsungGalaxyM345g() *SpecificationSeederMobileSamsungGalaxyM345g {
	return &SpecificationSeederMobileSamsungGalaxyM345g{BaseSeeder: BaseSeeder{name: "Specifications for Samsung Galaxy M34 5G"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSamsungGalaxyM345g) getBanglaTranslations() map[string]string {
	return map[string]string{
		"120 Hz": "১২০ Hz",
		"128 / 256 GB + microSD (up to 1TB)": "১২৮ / ২৫৬ জিবি + মাইক্রোএসডি (পর্যন্ত ১TB)",
		"13 MP": "১৩ মেগাপিক্সেল",
		"161.7 × 77.2 × 8.8 mm": "১৬১.৭ × ৭৭.২ × ৮.৮ মিমি",
		"208 g": "২০৮ g",
		"2340 × 1080 px": "২৩৪০ × ১০৮০ পিক্সেল",
		"25 W wired": "২৫ W তারযুক্ত",
		"50 MP + 8 MP + 2 MP": "৫০ মেগাপিক্সেল + ৮ মেগাপিক্সেল + ২ মেগাপিক্সেল",
		"6 / 8 GB": "৬ / ৮ GB",
		"6.6 inches": "৬.৬ ইঞ্চি",
		"6000 mAh": "৬০০০ এমএএইচ",
		"Accelerometer, Gyro, Proximity, Compass, fingerprint (side)": "অ্যাক্সিলেরোমিটার, জাইরো, প্রক্সিমিটি, কম্পাস, ফিঙ্গারপ্রিন্ট (side)",
		"Android 13, One UI 5 (upgradable)": "অ্যান্ড্রয়েড ১৩, ওয়ান ইউআই ৫ (আপগ্রেডযোগ্য)",
		"Exynos 1280": "এক্সিনস ১২৮০",
		"Exynos 1280 (5 nm)": "এক্সিনস ১২৮০ (৫ ন্যানোমিটার)",
		"GSM / HSPA / LTE / 5G": "জিএসএম / এইচএসপিএ / এলটিই / ৫G",
		"Glass front (Gorilla Glass 5), plastic frame / back": "গ্লাস সামনে (গরিলা গ্লাস ৫), প্লাস্টিক ফ্রেম / পেছনে",
		"Hybrid Dual SIM (nano + microSD)": "Hybrid ডুয়াল সিম (nano + মাইক্রোএসডি)",
		"July 2023": "জুলাই ২০২৩",
		"Li-Po (non-removable)": "লি-পো (অপসারণযোগ্য নয়)",
		"Mali-G68 MP4": "মালি-G৬৮ মেগাপিক্সেল৪",
		"Midnight Blue, Prism Silver, Waterfall Blue": "মিডনাইট নীল, প্রিজম রূপালী, ওয়াটারফল নীল",
		"Octa-core (2×2.4 GHz + 6×2.0 GHz)": "অক্টা-কোর (২×২.৪ গিগাহার্টজ + ৬×২.০ গিগাহার্টজ)",
		"Super AMOLED, 120 Hz": "Super অ্যামোলেড, ১২০ Hz",
		"USB-C 2.0": "ইউএসবি-সি ২.০",
		"Wi-Fi 802.11 a/b/g/n/ac": "ওয়াই-ফাই ৮০২.১১ a/b/g/n/ac",
	}
}

// Seed inserts specification records for the product identified by slug 'samsung-galaxy-m34-5g'
func (s *SpecificationSeederMobileSamsungGalaxyM345g) Seed(db *gorm.DB) error {
	productSlug := "samsung-galaxy-m34-5g"

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

	// Override model-specific values for Samsung Galaxy M34 5G
	specs["Display Size"] = "6.6 inches"
	specs["Processor"] = "Exynos 1280"
	specs["Chipset"] = "Exynos 1280 (5 nm)"
	specs["Cpu Type"] = "Octa-core (2×2.4 GHz + 6×2.0 GHz)"
	specs["Gpu Type"] = "Mali-G68 MP4"
	specs["Ram"] = "6 / 8 GB"
	specs["Storage"] = "128 / 256 GB + microSD (up to 1TB)"
	specs["Display Type"] = "Super AMOLED, 120 Hz"
	specs["Resolution"] = "2340 × 1080 px"
	specs["Refresh Rate"] = "120 Hz"
	specs["Build Material"] = "Glass front (Gorilla Glass 5), plastic frame / back"
	specs["Weight"] = "208 g"
	specs["Dimensions"] = "161.7 × 77.2 × 8.8 mm"
	specs["Network Technology"] = "GSM / HSPA / LTE / 5G"
	specs["Wifi Support"] = "Wi-Fi 802.11 a/b/g/n/ac"
	specs["Usb Type"] = "USB-C 2.0"
	specs["Rear Camera"] = "50 MP + 8 MP + 2 MP"
	specs["Camera Features"] = "OIS (main), LED flash"
	specs["Optical Zoom"] = "None"
	specs["Front Camera"] = "13 MP"
	specs["Operating System"] = "Android 13, One UI 5 (upgradable)"
	specs["Battery"] = "6000 mAh"
	specs["Battery Type"] = "Li-Po (non-removable)"
	specs["Fast Charging"] = "25 W wired"
	specs["Wireless Charging"] = "No"
	specs["5G Support"] = "Yes"
	specs["Sensors"] = "Accelerometer, Gyro, Proximity, Compass, fingerprint (side)"
	specs["Sim Card Type"] = "Hybrid Dual SIM (nano + microSD)"
	specs["Loudspeaker Quality"] = "Mono / Stereo (depends)"
	specs["Available Colors"] = "Midnight Blue, Prism Silver, Waterfall Blue"
	specs["Announcement Date"] = "July 2023"
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
