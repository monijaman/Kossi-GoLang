package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileIphone17Plus seeds specifications/options for product 'iphone-17-plus'
type SpecificationSeederMobileIphone17Plus struct {
	BaseSeeder
}

// NewSpecificationSeederMobileIphone17Plus creates a new seeder instance
func NewSpecificationSeederMobileIphone17Plus() *SpecificationSeederMobileIphone17Plus {
	return &SpecificationSeederMobileIphone17Plus{BaseSeeder: BaseSeeder{name: "Specifications for iPhone 17 Plus"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileIphone17Plus) getBanglaTranslations() map[string]string {
	return map[string]string{
		"12 MP, f/1.9": "১২ MP, f/১.৯",
		"128 GB / 256 GB / 512 GB": "১২৮ GB / ২৫৬ GB / ৫১২ GB",
		"160.9 x 77.8 x 7.8 mm": "১৬০.৯ x ৭৭.৮ x ৭.৮ মিমি",
		"203 g (7.16 oz)": "২০৩ g (৭.১৬ oz)",
		"20W wired, 15W MagSafe wireless": "২০W wired, ১৫W MagSafe wireless",
		"2796 x 1290 pixels (~460 ppi density)": "২৭৯৬ x ১২৯০ pixels (~৪৬০ ppi density)",
		"2x Optical Zoom (via crop)": "২x অপটিক্যাল জুম (ক্রপের মাধ্যমে)",
		"4,383 mAh": "৪,৩৮৩ এমএএইচ",
		"48 MP + 12 MP": "৪৮ MP + ১২ MP",
		"4K@24/25/30/60fps, 1080p@25/30/60/120/240fps, HDR, Dolby Vision HDR": "৪K@২৪/২৫/৩০/৬০fps, ১০৮০p@২৫/৩০/৬০/১২০/২৪০fps, HDR, Dolby Vision HDR",
		"4K@24/25/30/60fps, 1080p@25/30/60/120fps": "৪K@২৪/২৫/৩০/৬০fps, ১০৮০p@২৫/৩০/৬০/১২০fps",
		"5-core Apple GPU": "৫-core Apple GPU",
		"5.3, A2DP, LE": "৫.৩, A২DP, LE",
		"6.7 inches": "৬.৭ ইঞ্চি",
		"60Hz": "৬০Hz",
		"8 GB": "৮ GB",
		"Apple A19": "Apple A১৯",
		"Apple A19 (3 nm)": "Apple A১৯ (৩ nm)",
		"Black, Blue, Green, Yellow, Pink": "কালো, নীল, সবুজ, হলুদ, গোলাপী",
		"GSM / CDMA / HSPA / EVDO / LTE / 5G": "GSM / CDMA / HSPA / EVDO / LTE / ৫G",
		"Hexa-core (2x Performance + 4x Efficiency)": "Hexa-core (২x Performance + ৪x Efficiency)",
		"IP68 dust/water resistant (up to 6m for 30 min)": "IP৬৮ dust/water resistant (up to ৬m for ৩০ min)",
		"Li-Ion (non-removable)": "লি-আয়ন (অপসারণযোগ্য নয়)",
		"No 3.5mm jack": "না ৩.৫mm jack",
		"September 2024": "September ২০২৪",
		"Super Retina XDR OLED, HDR10, Dolby Vision, 2000 nits (HBM)": "Super Retina XDR OLED, HDR১০, Dolby Vision, ২০০০ nits (HBM)",
		"USB Type-C 2.0": "USB Type-C ২.০",
		"Wi-Fi 802.11 a/b/g/n/ac/6": "Wi-Fi ৮০২.১১ a/b/g/n/ac/৬",
		"iOS 18": "iOS ১৮",
	}
}

// Seed inserts specification records for the product identified by slug 'iphone-17-plus'
func (s *SpecificationSeederMobileIphone17Plus) Seed(db *gorm.DB) error {
	productSlug := "iphone-17-plus"

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

	// Override model-specific values for iPhone 17 Plus
	specs["Display Size"] = "6.7 inches"
	specs["Processor"] = "Apple A19"
	specs["Chipset"] = "Apple A19 (3 nm)"
	specs["Cpu Type"] = "Hexa-core (2x Performance + 4x Efficiency)"
	specs["Gpu Type"] = "5-core Apple GPU"
	specs["Ram"] = "8 GB"
	specs["Storage"] = "128 GB / 256 GB / 512 GB"
	specs["Display Type"] = "Super Retina XDR OLED, HDR10, Dolby Vision, 2000 nits (HBM)"
	specs["Resolution"] = "2796 x 1290 pixels (~460 ppi density)"
	specs["Screen Protection"] = "Ceramic Shield"
	specs["Refresh Rate"] = "60Hz"
	specs["Build Material"] = "Glass front (Ceramic Shield), glass back, aluminum frame"
	specs["Weight"] = "203 g (7.16 oz)"
	specs["Dimensions"] = "160.9 x 77.8 x 7.8 mm"
	specs["Water Resistance"] = "IP68 dust/water resistant (up to 6m for 30 min)"
	specs["Network Technology"] = "GSM / CDMA / HSPA / EVDO / LTE / 5G"
	specs["Wifi Support"] = "Wi-Fi 802.11 a/b/g/n/ac/6"
	specs["Bluetooth Version"] = "5.3, A2DP, LE"
	specs["Nfc Support"] = "Yes"
	specs["Usb Type"] = "USB Type-C 2.0"
	specs["Rear Camera"] = "48 MP + 12 MP"
	specs["Camera Features"] = "Dual-LED dual-tone flash, HDR (photo/panorama)"
	specs["Camera Video Resolution"] = "4K@24/25/30/60fps, 1080p@25/30/60/120/240fps, HDR, Dolby Vision HDR"
	specs["Optical Zoom"] = "2x Optical Zoom (via crop)"
	specs["Front Camera"] = "12 MP, f/1.9"
	specs["Front Camera Video Resolution"] = "4K@24/25/30/60fps, 1080p@25/30/60/120fps"
	specs["Operating System"] = "iOS 18"
	specs["Battery"] = "4,383 mAh"
	specs["Battery Type"] = "Li-Ion (non-removable)"
	specs["Fast Charging"] = "20W wired, 15W MagSafe wireless"
	specs["Wireless Charging"] = "Yes"
	specs["5G Support"] = "Yes"
	specs["Positioning System"] = "GPS, GLONASS, GALILEO, BDS, QZSS"
	specs["Sensors"] = "Face ID, accelerometer, gyro, proximity, compass, barometer"
	specs["Special Features"] = "Emergency SOS via satellite"
	specs["Sim Card Type"] = "Nano-SIM and eSIM"
	specs["Loudspeaker Quality"] = "Stereo speakers"
	specs["Audio Quality"] = "No 3.5mm jack"
	specs["Audio Jack"] = "No"
	specs["Available Colors"] = "Black, Blue, Green, Yellow, Pink"
	specs["Announcement Date"] = "September 2024"
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
