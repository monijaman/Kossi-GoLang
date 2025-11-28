package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileSamsungGalaxyA14 seeds specifications/options for product 'samsung-galaxy-a14'
type SpecificationSeederMobileSamsungGalaxyA14 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSamsungGalaxyA14 creates a new seeder instance
func NewSpecificationSeederMobileSamsungGalaxyA14() *SpecificationSeederMobileSamsungGalaxyA14 {
	return &SpecificationSeederMobileSamsungGalaxyA14{BaseSeeder: BaseSeeder{name: "Specifications for Samsung Galaxy A14"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSamsungGalaxyA14) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080p @ 30fps": "১০৮০p @ ৩০fps",
		"13 MP": "১৩ মেগাপিক্সেল",
		"15 W wired": "১৫ W তারযুক্ত",
		"167.7 × 78.0 × 9.1 mm": "১৬৭.৭ × ৭৮.০ × ৯.১ মিমি",
		"201 g": "২০১ g",
		"2408 × 1080 px (FHD+)": "২৪০৮ × ১০৮০ পিক্সেল (FHD+)",
		"3.5 mm headphone jack": "৩.৫ মিমি হেডফোন জ্যাক",
		"4 / 6 GB": "৪ / ৬ GB",
		"5.2 (5G) / 5.0 (4G)": "৫.২ (৫G) / ৫.০ (৪G)",
		"50 MP + 5 MP + 2 MP": "৫০ মেগাপিক্সেল + ৫ মেগাপিক্সেল + ২ মেগাপিক্সেল",
		"5000 mAh": "৫০০০ এমএএইচ",
		"6.6 inches": "৬.৬ ইঞ্চি",
		"60 Hz": "৬০ Hz",
		"64 / 128 GB + microSD support": "৬৪ / ১২৮ জিবি + মাইক্রোএসডি suppবাt",
		"Android 13 (upgradable to Android 14)": "অ্যান্ড্রয়েড ১৩ (আপগ্রেডযোগ্য থেকে অ্যান্ড্রয়েড ১৪)",
		"Black, Silver, Green, Red": "কালো, রূপালী, সবুজ, লাল",
		"GSM / HSPA / LTE (4G) / 5G (A14 5G variant)": "জিএসএম / এইচএসপিএ / এলটিই (৪G) / ৫G (A১৪ ৫G variant)",
		"Glass front, plastic back & frame": "গ্লাস সামনে, প্লাস্টিক পেছনে & ফ্রেম",
		"Helio G80 (12 nm) / Dimensity 700 (7 nm)": "হেলিও G৮০ (১২ ন্যানোমিটার) / ডাইমেনসিটি ৭০০ (৭ ন্যানোমিটার)",
		"January 2023": "জানুয়ারি ২০২৩",
		"Li-Ion (non-removable)": "লি-আয়ন (অপসারণযোগ্য নয়)",
		"Mali-G52 (4G) / Mali-G57 (5G)": "মালি-G৫২ (৪G) / মালি-G৫৭ (৫G)",
		"MediaTek Helio G80 (4G) / Dimensity 700 (5G)": "মিডিয়াটেক হেলিও G৮০ (৪G) / ডাইমেনসিটি ৭০০ (৫G)",
		"Side fingerprint, Accelerometer, Gyro, Proximity, Compass": "সাইড ফিঙ্গারপ্রিন্ট, অ্যাক্সিলেরোমিটার, জাইরো, প্রক্সিমিটি, কম্পাস",
		"USB-C 2.0": "ইউএসবি-সি ২.০",
		"Wi-Fi 802.11 a/b/g/n/ac": "ওয়াই-ফাই ৮০২.১১ a/b/g/n/ac",
		"Yes (market dependent)": "হ্যাঁ (বাজার নির্ভরশীল)",
		"Yes (only A14 5G variant)": "হ্যাঁ (only A১৪ ৫G variant)",
	}
}

// Seed inserts specification records for the product identified by slug 'samsung-galaxy-a14'
func (s *SpecificationSeederMobileSamsungGalaxyA14) Seed(db *gorm.DB) error {
	productSlug := "samsung-galaxy-a14"

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

	// Override model-specific values for Samsung Galaxy A14
	specs["Display Size"] = "6.6 inches"
	specs["Processor"] = "MediaTek Helio G80 (4G) / Dimensity 700 (5G)"
	specs["Chipset"] = "Helio G80 (12 nm) / Dimensity 700 (7 nm)"
	specs["Cpu Type"] = "Octa-core"
	specs["Gpu Type"] = "Mali-G52 (4G) / Mali-G57 (5G)"
	specs["Ram"] = "4 / 6 GB"
	specs["Storage"] = "64 / 128 GB + microSD support"
	specs["Display Type"] = "PLS LCD"
	specs["Resolution"] = "2408 × 1080 px (FHD+)"
	specs["Refresh Rate"] = "60 Hz"
	specs["Build Material"] = "Glass front, plastic back & frame"
	specs["Weight"] = "201 g"
	specs["Dimensions"] = "167.7 × 78.0 × 9.1 mm"
	specs["Water Resistance"] = "None"
	specs["Network Technology"] = "GSM / HSPA / LTE (4G) / 5G (A14 5G variant)"
	specs["Wifi Support"] = "Wi-Fi 802.11 a/b/g/n/ac"
	specs["Bluetooth Version"] = "5.2 (5G) / 5.0 (4G)"
	specs["Nfc Support"] = "Yes (market dependent)"
	specs["Usb Type"] = "USB-C 2.0"
	specs["Rear Camera"] = "50 MP + 5 MP + 2 MP"
	specs["Camera Features"] = "Autofocus, LED flash"
	specs["Camera Video Resolution"] = "1080p @ 30fps"
	specs["Optical Zoom"] = "None"
	specs["Front Camera"] = "13 MP"
	specs["Front Camera Video Resolution"] = "1080p @ 30fps"
	specs["Operating System"] = "Android 13 (upgradable to Android 14)"
	specs["Battery"] = "5000 mAh"
	specs["Battery Type"] = "Li-Ion (non-removable)"
	specs["Fast Charging"] = "15 W wired"
	specs["Wireless Charging"] = "No"
	specs["5G Support"] = "Yes (only A14 5G variant)"
	specs["Positioning System"] = "GPS, GLONASS, GALILEO, BDS"
	specs["Sensors"] = "Side fingerprint, Accelerometer, Gyro, Proximity, Compass"
	specs["Special Features"] = "Virtual RAM, Eye Comfort Shield"
	specs["Sim Card Type"] = "Nano-SIM or Dual SIM"
	specs["Loudspeaker Quality"] = "Mono speaker"
	specs["Audio Jack"] = "3.5 mm headphone jack"
	specs["Available Colors"] = "Black, Silver, Green, Red"
	specs["Announcement Date"] = "January 2023"
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
