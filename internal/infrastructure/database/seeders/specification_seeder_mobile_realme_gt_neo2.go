package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileRealmeGtNeo2 seeds specifications/options for product 'realme-gt-neo2'
type SpecificationSeederMobileRealmeGtNeo2 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileRealmeGtNeo2 creates a new seeder instance
func NewSpecificationSeederMobileRealmeGtNeo2() *SpecificationSeederMobileRealmeGtNeo2 {
	return &SpecificationSeederMobileRealmeGtNeo2{BaseSeeder: BaseSeeder{name: "Specifications for Realme GT Neo2"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileRealmeGtNeo2) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080p @ 30fps": "১০৮০p @ ৩০fps",
		"120 Hz": "১২০ Hz",
		"128 / 256 GB": "১২৮ / ২৫৬ GB",
		"16 MP": "১৬ মেগাপিক্সেল",
		"200 g": "২০০ g",
		"2021": "২০২১",
		"2400 × 1080 px (~402 ppi)": "২৪০০ × ১০৮০ পিক্সেল (~৪০২ পিপিআই)",
		"4K @ 30/60fps, 1080p @ 30/60/120fps": "৪K @ ৩০/৬০fps, ১০৮০p @ ৩০/৬০/১২০fps",
		"5.1": "৫.১",
		"5000 mAh": "৫০০০ এমএএইচ",
		"6.6 inches": "৬.৬ ইঞ্চি",
		"64 MP + 8 MP + 2 MP": "৬৪ মেগাপিক্সেল + ৮ মেগাপিক্সেল + ২ মেগাপিক্সেল",
		"65 W wired": "৬৫ W তারযুক্ত",
		"8 / 12 GB": "৮ / ১২ GB",
		"AMOLED, 120 Hz": "অ্যামোলেড, ১২০ Hz",
		"Adreno 650": "অ্যাড্রেনো ৬৫০",
		"Android 11, Realme UI 2.0": "অ্যান্ড্রয়েড ১১, Realme UI ২.০",
		"Dual SIM (Nano-SIM)": "ডুয়াল সিম (ন্যানো-সিম)",
		"Fingerprint (under display), Accelerometer, Gyro, Proximity, Compass": "ফিঙ্গারপ্রিন্ট (ডিসপ্লের নিচে), অ্যাক্সিলেরোমিটার, জাইরো, প্রক্সিমিটি, কম্পাস",
		"GSM / HSPA / LTE / 5G": "জিএসএম / এইচএসপিএ / এলটিই / ৫G",
		"Glass front, plastic frame/back": "গ্লাস সামনে, প্লাস্টিক ফ্রেম/পেছনে",
		"Li-Ion (non-removable)": "লি-আয়ন (অপসারণযোগ্য নয়)",
		"Neo Green, Neo Blue, Neo Black": "Neo সবুজ, Neo নীল, Neo কালো",
		"Octa-core (1×3.2 GHz Cortex-A78 + 3×2.42 GHz Cortex-A78 + 4×1.8 GHz Cortex-A55)": "অক্টা-কোর (১×৩.২ গিগাহার্টজ Cবাtex-A৭৮ + ৩×২.৪২ গিগাহার্টজ Cবাtex-A৭৮ + ৪×১.৮ গিগাহার্টজ Cবাtex-A৫৫)",
		"Qualcomm Snapdragon 870": "কোয়ালকম স্ন্যাপড্রাগন ৮৭০",
		"Snapdragon 870 (7 nm)": "স্ন্যাপড্রাগন ৮৭০ (৭ ন্যানোমিটার)",
		"USB-C 2.0": "ইউএসবি-সি ২.০",
		"Wi-Fi 802.11 a/b/g/n/ac/6": "ওয়াই-ফাই ৮০২.১১ a/b/g/n/ac/৬",
	}
}

// Seed inserts specification records for the product identified by slug 'realme-gt-neo2'
func (s *SpecificationSeederMobileRealmeGtNeo2) Seed(db *gorm.DB) error {
	productSlug := "realme-gt-neo2"

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

	// Override model-specific values for Realme GT Neo2
	specs["Display Size"] = "6.6 inches"
	specs["Processor"] = "Qualcomm Snapdragon 870"
	specs["Chipset"] = "Snapdragon 870 (7 nm)"
	specs["Cpu Type"] = "Octa-core (1×3.2 GHz Cortex-A78 + 3×2.42 GHz Cortex-A78 + 4×1.8 GHz Cortex-A55)"
	specs["Gpu Type"] = "Adreno 650"
	specs["Ram"] = "8 / 12 GB"
	specs["Storage"] = "128 / 256 GB"
	specs["Display Type"] = "AMOLED, 120 Hz"
	specs["Resolution"] = "2400 × 1080 px (~402 ppi)"
	specs["Refresh Rate"] = "120 Hz"
	specs["Build Material"] = "Glass front, plastic frame/back"
	specs["Weight"] = "200 g"
	specs["Network Technology"] = "GSM / HSPA / LTE / 5G"
	specs["Wifi Support"] = "Wi-Fi 802.11 a/b/g/n/ac/6"
	specs["Bluetooth Version"] = "5.1"
	specs["Nfc Support"] = "Yes"
	specs["Usb Type"] = "USB-C 2.0"
	specs["Rear Camera"] = "64 MP + 8 MP + 2 MP"
	specs["Camera Features"] = "OIS (main), LED flash, HDR"
	specs["Camera Video Resolution"] = "4K @ 30/60fps, 1080p @ 30/60/120fps"
	specs["Optical Zoom"] = "None"
	specs["Front Camera"] = "16 MP"
	specs["Front Camera Video Resolution"] = "1080p @ 30fps"
	specs["Operating System"] = "Android 11, Realme UI 2.0"
	specs["Battery"] = "5000 mAh"
	specs["Battery Type"] = "Li-Ion (non-removable)"
	specs["Fast Charging"] = "65 W wired"
	specs["Wireless Charging"] = "No"
	specs["5G Support"] = "Yes"
	specs["Positioning System"] = "GPS, GLONASS, GALILEO, BDS"
	specs["Sensors"] = "Fingerprint (under display), Accelerometer, Gyro, Proximity, Compass"
	specs["Sim Card Type"] = "Dual SIM (Nano-SIM)"
	specs["Loudspeaker Quality"] = "Stereo"
	specs["Audio Jack"] = "No"
	specs["Available Colors"] = "Neo Green, Neo Blue, Neo Black"
	specs["Announcement Date"] = "2021"
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
