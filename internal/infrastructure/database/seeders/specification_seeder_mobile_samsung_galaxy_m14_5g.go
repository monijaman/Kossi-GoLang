package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileSamsungGalaxyM145g seeds specifications/options for product 'samsung-galaxy-m14-5g'
type SpecificationSeederMobileSamsungGalaxyM145g struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSamsungGalaxyM145g creates a new seeder instance
func NewSpecificationSeederMobileSamsungGalaxyM145g() *SpecificationSeederMobileSamsungGalaxyM145g {
	return &SpecificationSeederMobileSamsungGalaxyM145g{BaseSeeder: BaseSeeder{name: "Specifications for Samsung Galaxy M14 5G"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSamsungGalaxyM145g) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080p @ 30fps": "১০৮০p @ ৩০fps",
		"13 MP": "১৩ MP",
		"15 W / 25 W (region dependent)": "১৫ W / ২৫ W (region dependent)",
		"166.8 × 77.2 × 9.4 mm": "১৬৬.৮ × ৭৭.২ × ৯.৪ মিমি",
		"206 g": "২০৬ g",
		"2408 × 1080 px": "২৪০৮ × ১০৮০ px",
		"3.5 mm": "৩.৫ মিমি",
		"4 / 6 GB": "৪ / ৬ GB",
		"50 MP + 2 MP + 2 MP": "৫০ MP + ২ MP + ২ MP",
		"6.6 inches": "৬.৬ ইঞ্চি",
		"6000 mAh": "৬০০০ এমএএইচ",
		"64 / 128 GB + microSD": "৬৪ / ১২৮ GB + microSD",
		"90 Hz": "৯০ Hz",
		"Android 13, One UI 5.1": "Android ১৩, One UI ৫.১",
		"Exynos 1330": "Exynos ১৩৩০",
		"Exynos 1330 (5 nm)": "Exynos ১৩৩০ (৫ nm)",
		"GSM / HSPA / LTE / 5G": "GSM / HSPA / LTE / ৫G",
		"Glass front, plastic back/frame": "গ্লাস সামনে, প্লাস্টিক পেছনে/frame",
		"Li-Ion (non-removable)": "লি-আয়ন (অপসারণযোগ্য নয়)",
		"Mali-G68 MP2": "Mali-G৬৮ MP২",
		"March 2023": "March ২০২৩",
		"Nano-SIM or Dual SIM (hybrid)": "ন্যানো-সিম or ডুয়াল সিম (hybrid)",
		"Navy Blue, Light Blue, Silver": "Navy নীল, Light নীল, রূপালী",
		"Octa-core (2×2.4 GHz Cortex-A78 + 6×2.0 GHz Cortex-A55)": "Octa-core (২×২.৪ GHz Cortex-A৭৮ + ৬×২.০ GHz Cortex-A৫৫)",
		"Side fingerprint, Accelerometer, Gyro, Proximity, Compass": "Side ফিঙ্গারপ্রিন্ট, অ্যাক্সিলেরোমিটার, জাইরো, প্রক্সিমিটি, কম্পাস",
		"USB-C 2.0": "USB-C ২.০",
		"Wi-Fi 802.11 a/b/g/n/ac": "Wi-Fi ৮০২.১১ a/b/g/n/ac",
	}
}

// Seed inserts specification records for the product identified by slug 'samsung-galaxy-m14-5g'
func (s *SpecificationSeederMobileSamsungGalaxyM145g) Seed(db *gorm.DB) error {
	productSlug := "samsung-galaxy-m14-5g"

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

	// Override model-specific values for Samsung Galaxy M14 5G
	specs["Display Size"] = "6.6 inches"
	specs["Processor"] = "Exynos 1330"
	specs["Chipset"] = "Exynos 1330 (5 nm)"
	specs["Cpu Type"] = "Octa-core (2×2.4 GHz Cortex-A78 + 6×2.0 GHz Cortex-A55)"
	specs["Gpu Type"] = "Mali-G68 MP2"
	specs["Ram"] = "4 / 6 GB"
	specs["Storage"] = "64 / 128 GB + microSD"
	specs["Display Type"] = "PLS LCD"
	specs["Resolution"] = "2408 × 1080 px"
	specs["Refresh Rate"] = "90 Hz"
	specs["Build Material"] = "Glass front, plastic back/frame"
	specs["Weight"] = "206 g"
	specs["Dimensions"] = "166.8 × 77.2 × 9.4 mm"
	specs["Network Technology"] = "GSM / HSPA / LTE / 5G"
	specs["Wifi Support"] = "Wi-Fi 802.11 a/b/g/n/ac"
	specs["Usb Type"] = "USB-C 2.0"
	specs["Rear Camera"] = "50 MP + 2 MP + 2 MP"
	specs["Camera Features"] = "LED flash, Depth"
	specs["Camera Video Resolution"] = "1080p @ 30fps"
	specs["Optical Zoom"] = "None"
	specs["Front Camera"] = "13 MP"
	specs["Front Camera Video Resolution"] = "1080p @ 30fps"
	specs["Operating System"] = "Android 13, One UI 5.1"
	specs["Battery"] = "6000 mAh"
	specs["Battery Type"] = "Li-Ion (non-removable)"
	specs["Fast Charging"] = "15 W / 25 W (region dependent)"
	specs["Wireless Charging"] = "No"
	specs["5G Support"] = "Yes"
	specs["Sensors"] = "Side fingerprint, Accelerometer, Gyro, Proximity, Compass"
	specs["Sim Card Type"] = "Nano-SIM or Dual SIM (hybrid)"
	specs["Loudspeaker Quality"] = "Mono"
	specs["Audio Jack"] = "3.5 mm"
	specs["Available Colors"] = "Navy Blue, Light Blue, Silver"
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
