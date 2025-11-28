package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileSamsungGalaxyM165g seeds specifications/options for product 'samsung-galaxy-m16-5g'
type SpecificationSeederMobileSamsungGalaxyM165g struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSamsungGalaxyM165g creates a new seeder instance
func NewSpecificationSeederMobileSamsungGalaxyM165g() *SpecificationSeederMobileSamsungGalaxyM165g {
	return &SpecificationSeederMobileSamsungGalaxyM165g{BaseSeeder: BaseSeeder{name: "Specifications for Samsung Galaxy M16 5G"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSamsungGalaxyM165g) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080 × 2340 px": "১০৮০ × ২৩৪০ পিক্সেল",
		"128 GB (maybe more)": "১২৮ জিবি (maybe mবাe)",
		"13 MP": "১৩ মেগাপিক্সেল",
		"164.4 × 77.9 × 7.9 mm": "১৬৪.৪ × ৭৭.৯ × ৭.৯ মিমি",
		"191 g (according to 8 GB variant)": "১৯১ গ্রাম (accবাding থেকে ৮ জিবি variant)",
		"25 W wired": "২৫ W তারযুক্ত",
		"5.3": "৫.৩",
		"50 MP + 5 MP + 2 MP": "৫০ মেগাপিক্সেল + ৫ মেগাপিক্সেল + ২ মেগাপিক্সেল",
		"5000 mAh": "৫০০০ এমএএইচ",
		"6 / 8 GB": "৬ / ৮ GB",
		"6 years of updates (claimed)": "৬ years of updates (claimed)",
		"6.7 inches": "৬.৭ ইঞ্চি",
		"90 Hz": "৯০ Hz",
		"Android 15, One UI 7": "অ্যান্ড্রয়েড ১৫, ওয়ান ইউআই ৭",
		"Dimensity 6300 (6 nm)": "ডাইমেনসিটি ৬৩০০ (৬ ন্যানোমিটার)",
		"February 2025": "ফেব্রুয়ারি ২০২৫",
		"GSM / HSPA / LTE / 5G": "জিএসএম / এইচএসপিএ / এলটিই / ৫G",
		"Glass front, plastic back / frame": "গ্লাস সামনে, প্লাস্টিক পেছনে / ফ্রেম",
		"IP54 (splash)": "IP৫৪ (splash)",
		"Li-Ion (non-removable)": "লি-আয়ন (অপসারণযোগ্য নয়)",
		"Mali-G57 MC2": "মালি-G৫৭ MC২",
		"MediaTek Dimensity 6300": "মিডিয়াটেক ডাইমেনসিটি ৬৩০০",
		"Nano-SIM + Nano-SIM (hybrid)": "ন্যানো-সিম + ন্যানো-সিম (hybrid)",
		"Octa-core (2×2.4 GHz Cortex-A76 + 6×2.0 GHz Cortex-A55)": "অক্টা-কোর (২×২.৪ গিগাহার্টজ Cবাtex-A৭৬ + ৬×২.০ গিগাহার্টজ Cবাtex-A৫৫)",
		"Side fingerprint, Accelerometer, Gyro, Proximity, Compass": "সাইড ফিঙ্গারপ্রিন্ট, অ্যাক্সিলেরোমিটার, জাইরো, প্রক্সিমিটি, কম্পাস",
		"Thunder Black, Mint Green, Blush Pink": "থান্ডার কালো, মিন্ট সবুজ, ব্লাশ গোলাপী",
		"USB-C (no 3.5mm)": "ইউএসবি-সি (no ৩.৫mm)",
		"Wi-Fi 802.11 a/b/g/n/ac": "ওয়াই-ফাই ৮০২.১১ a/b/g/n/ac",
	}
}

// Seed inserts specification records for the product identified by slug 'samsung-galaxy-m16-5g'
func (s *SpecificationSeederMobileSamsungGalaxyM165g) Seed(db *gorm.DB) error {
	productSlug := "samsung-galaxy-m16-5g"

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

	// Override model-specific values for Samsung Galaxy M16 5G
	specs["Display Size"] = "6.7 inches"
	specs["Processor"] = "MediaTek Dimensity 6300"
	specs["Chipset"] = "Dimensity 6300 (6 nm)"
	specs["Cpu Type"] = "Octa-core (2×2.4 GHz Cortex-A76 + 6×2.0 GHz Cortex-A55)"
	specs["Gpu Type"] = "Mali-G57 MC2"
	specs["Ram"] = "6 / 8 GB"
	specs["Storage"] = "128 GB (maybe more)"
	specs["Display Type"] = "Super AMOLED"
	specs["Resolution"] = "1080 × 2340 px"
	specs["Refresh Rate"] = "90 Hz"
	specs["Build Material"] = "Glass front, plastic back / frame"
	specs["Weight"] = "191 g (according to 8 GB variant)"
	specs["Dimensions"] = "164.4 × 77.9 × 7.9 mm"
	specs["Water Resistance"] = "IP54 (splash)"
	specs["Network Technology"] = "GSM / HSPA / LTE / 5G"
	specs["Wifi Support"] = "Wi-Fi 802.11 a/b/g/n/ac"
	specs["Bluetooth Version"] = "5.3"
	specs["Nfc Support"] = "Yes"
	specs["Usb Type"] = "USB-C"
	specs["Rear Camera"] = "50 MP + 5 MP + 2 MP"
	specs["Camera Features"] = "LED flash, Auto-focus, Depth / Macro"
	specs["Optical Zoom"] = "None"
	specs["Front Camera"] = "13 MP"
	specs["Operating System"] = "Android 15, One UI 7"
	specs["Battery"] = "5000 mAh"
	specs["Battery Type"] = "Li-Ion (non-removable)"
	specs["Fast Charging"] = "25 W wired"
	specs["Wireless Charging"] = "No"
	specs["5G Support"] = "Yes"
	specs["Positioning System"] = "GPS, GLONASS, GALILEO, BDS"
	specs["Sensors"] = "Side fingerprint, Accelerometer, Gyro, Proximity, Compass"
	specs["Special Features"] = "6 years of updates (claimed)"
	specs["Sim Card Type"] = "Nano-SIM + Nano-SIM (hybrid)"
	specs["Loudspeaker Quality"] = "Stereo"
	specs["Audio Jack"] = "USB-C (no 3.5mm)"
	specs["Available Colors"] = "Thunder Black, Mint Green, Blush Pink"
	specs["Announcement Date"] = "February 2025"
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
