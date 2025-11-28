package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileGooglePixel9ProXl seeds specifications/options for product 'google-pixel-9-pro-xl'
type SpecificationSeederMobileGooglePixel9ProXl struct {
	BaseSeeder
}

// NewSpecificationSeederMobileGooglePixel9ProXl creates a new seeder instance
func NewSpecificationSeederMobileGooglePixel9ProXl() *SpecificationSeederMobileGooglePixel9ProXl {
	return &SpecificationSeederMobileGooglePixel9ProXl{BaseSeeder: BaseSeeder{name: "Specifications for Google Pixel 9 Pro XL"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileGooglePixel9ProXl) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1-120Hz (LTPO)": "১-১২০Hz (এলটিপিও)",
		"1.15 W/kg (head), 1.36 W/kg (body)": "১.১৫ W/kg (head), ১.৩৬ W/kg (body)",
		"1.17 W/kg (head), 1.24 W/kg (body)": "১.১৭ W/kg (head), ১.২৪ W/kg (body)",
		"1344 x 2992 pixels (~486 ppi density)": "১৩৪৪ x ২৯৯২ পিক্সেল (~৪৮৬ পিপিআই density)",
		"16 GB": "১৬ GB",
		"162.8 x 76.6 x 8.5 mm (6.41 x 3.02 x 0.33 in)": "১৬২.৮ x ৭৬.৬ x ৮.৫ মিমি (৬.৪১ x ৩.০২ x ০.৩৩ in)",
		"1x3.1 GHz + 3x2.6 GHz + 4x1.92 GHz": "১x৩.১ গিগাহার্টজ + ৩x২.৬ গিগাহার্টজ + ৪x১.৯২ গিগাহার্টজ",
		"221 g (7.80 oz)": "২২১ গ্রাম (৭.৮০ oz)",
		"24-bit/192kHz audio": "২৪-bit/১৯২kHz audio",
		"256 GB / 512 GB / 1 TB": "২৫৬ জিবি / ৫১২ জিবি / ১ টিবি",
		"37W wired (50% in 28 min), 23W wireless, 5W reverse wireless": "৩৭W তারযুক্ত (৫০% in ২৮ মিনিট), ২৩W বেতার, ৫W রিভার্স বেতার",
		"37W wired, 23W wireless, 5W reverse wireless": "৩৭W তারযুক্ত, ২৩W বেতার, ৫W রিভার্স বেতার",
		"42 MP, f/2.2, 17mm (ultrawide), PDAF": "৪২ মেগাপিক্সেল, f/২.২, ১৭mm (আল্ট্রাওয়াইড), PDAF",
		"4K@24/30/60fps, 1080p@24/30/60/120/240fps, gyro-EIS, OIS, 10-bit HDR": "৪K@২৪/৩০/৬০fps, ১০৮০p@২৪/৩০/৬০/১২০/২৪০fps, gyro-EIS, OIS, ১০-bit এইচডিআর",
		"4K@24/30/60fps, 1080p@30/60fps": "৪K@২৪/৩০/৬০fps, ১০৮০p@৩০/৬০fps",
		"5,060 mAh": "৫,০৬০ এমএএইচ",
		"5.3, A2DP, LE, aptX HD": "৫.৩, A২DP, LE, aptX HD",
		"50 MP + 48 MP + 48 MP": "৫০ মেগাপিক্সেল + ৪৮ মেগাপিক্সেল + ৪৮ মেগাপিক্সেল",
		"5x Optical Zoom": "৫x অপটিক্যাল জুম",
		"6.8 inches": "৬.৮ ইঞ্চি",
		"Android 15, upgradable to Android 22": "অ্যান্ড্রয়েড ১৫, আপগ্রেডযোগ্য থেকে অ্যান্ড্রয়েড ২২",
		"August 13, 2024": "আগস্ট ১৩, ২০২৪",
		"Available. Released August 22, 2024": "উপলব্ধ. প্রকাশিত আগস্ট ২২, ২০২৪",
		"Corning Gorilla Glass Victus 2": "কর্নিং গরিলা গ্লাস ভিক্টাস ২",
		"Fingerprint (under display, ultrasonic), accelerometer, gyro, proximity, compass, barometer, thermometer (skin temperature)": "ফিঙ্গারপ্রিন্ট (ডিসপ্লের নিচে, আল্ট্রাসনিক), অ্যাক্সিলেরোমিটার, জাইরো, প্রক্সিমিটি, কম্পাস, ব্যারোমিটার, থার্মোমিটার (ত্বকের তাপমাত্রা)",
		"GA05844-US, GA05845-US, GA05846-US": "GA০৫৮৪৪-US, GA০৫৮৪৫-US, GA০৫৮৪৬-US",
		"GPS (L1+L5), GLONASS (G1), GALILEO (E1+E5a), BDS (B1I+B1c+B2a+B2b), QZSS (L1+L5)": "জিপিএস (L১+L৫), GLONASS (G১), GALILEO (E১+E৫a), BDS (B১I+B১c+B২a+B২b), QZSS (L১+L৫)",
		"GSM / CDMA / HSPA / EVDO / LTE / 5G": "জিএসএম / CDMA / এইচএসপিএ / EVDO / এলটিই / ৫G",
		"GSM 850 / 900 / 1800 / 1900 MHz": "জিএসএম ৮৫০ / ৯০০ / ১৮০০ / ১৯০০ মেগাহার্টজ",
		"Gemini AI assistant, Circle to Search, Magic Eraser, Best Take, Audio Magic Eraser, Call Screen, Hold for Me": "জেমিনি এআই সহায়ক, সার্কেল টু সার্চ, ম্যাজিক ইরেজার, বেস্ট টেক, অডিও ম্যাজিক ইরেজার, কল স্ক্রিন, হোল্ড জন্য মি",
		"Glass front (Gorilla Glass Victus 2), glass back, aluminum frame": "গ্লাস সামনে (গরিলা গ্লাস Victus ২), গ্লাস পেছনে, অ্যালুমিনিয়াম ফ্রেম",
		"Google Tensor G4": "গুগল টেনসর G৪",
		"Google Tensor G4 (4 nm)": "গুগল টেনসর G৪ (৪ ন্যানোমিটার)",
		"HSDPA 800 / 850 / 900 / 1700(AWS) / 1900 / 2100 MHz": "HSDPA ৮০০ / ৮৫০ / ৯০০ / ১৭০০(AWS) / ১৯০০ / ২১০০ মেগাহার্টজ",
		"IP68 dust tight and water resistant (up to 1.5m for 30 min)": "IP৬৮ ধুলো প্রতিরোধী এবং জল প্রতিরোধী (পর্যন্ত ১.৫m জন্য ৩০ min)",
		"LTE Band 1, 2, 3, 4, 5, 7, 8, 12, 13, 14, 17, 18, 19, 20, 25, 26, 28, 29, 30, 32, 38, 39, 40, 41, 42, 46, 48, 66, 71": "এলটিই Bএবং ১, ২, ৩, ৪, ৫, ৭, ৮, ১২, ১৩, ১৪, ১৭, ১৮, ১৯, ২০, ২৫, ২৬, ২৮, ২৯, ৩০, ৩২, ৩৮, ৩৯, ৪০, ৪১, ৪২, ৪৬, ৪৮, ৬৬, ৭১",
		"LTPO OLED, 120Hz, HDR10+, 3000 nits (peak)": "এলটিপিও ওএলইডি, ১২০Hz, এইচডিআর১০+, ৩০০০ nits (পিক)",
		"Li-Ion (non-removable)": "লি-আয়ন (অপসারণযোগ্য নয়)",
		"Mali-G715 MC7": "মালি-G৭১৫ MC৭",
		"Octa-core (1x3.1 GHz Cortex-X4 & 3x2.6 GHz Cortex-A720 & 4x1.92 GHz Cortex-A520)": "অক্টা-কোর (১x৩.১ গিগাহার্টজ Cবাtex-X৪ & ৩x২.৬ গিগাহার্টজ Cবাtex-A৭২০ & ৪x১.৯২ গিগাহার্টজ Cবাtex-A৫২০)",
		"UFS 3.1": "UFS ৩.১",
		"USB Type-C 3.2, DisplayPort 1.4, OTG": "ইউএসবি টাইপ-সি ৩.২, DisplayPবাt ১.৪, OTG",
		"Wi-Fi 802.11 a/b/g/n/ac/6e/7, tri-band, Wi-Fi Direct": "ওয়াই-ফাই ৮০২.১১ a/b/g/n/ac/৬e/৭, tri-bএবং, ওয়াই-ফাই Direct",
		"Yes - 23W wireless": "Yes - ২৩W বেতার",
		"n1, n2, n3, n5, n7, n8, n12, n14, n20, n25, n26, n28, n30, n38, n40, n41, n48, n66, n71, n75, n76, n77, n78": "n১, n২, n৩, n৫, n৭, n৮, n১২, n১৪, n২০, n২৫, n২৬, n২৮, n৩০, n৩৮, n৪০, n৪১, n৪৮, n৬৬, n৭১, n৭৫, n৭৬, n৭৭, n৭৮",
	}
}

// Seed inserts specification records for the product identified by slug 'google-pixel-9-pro-xl'
func (s *SpecificationSeederMobileGooglePixel9ProXl) Seed(db *gorm.DB) error {
	productSlug := "google-pixel-9-pro-xl"

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

	// Override model-specific values for Google Pixel 9 Pro XL
	specs["Display Size"] = "6.8 inches"
	specs["Processor"] = "Google Tensor G4"
	specs["Chipset"] = "Google Tensor G4 (4 nm)"
	specs["Cpu Type"] = "Octa-core (1x3.1 GHz Cortex-X4 & 3x2.6 GHz Cortex-A720 & 4x1.92 GHz Cortex-A520)"
	specs["Gpu Type"] = "Mali-G715 MC7"
	specs["Processor Speed"] = "1x3.1 GHz + 3x2.6 GHz + 4x1.92 GHz"
	specs["Ram"] = "16 GB"
	specs["Storage"] = "256 GB / 512 GB / 1 TB"
	specs["Internal Memory Capacity"] = "UFS 3.1"
	specs["Card Slot Type"] = "No microSD card slot"
	specs["Display Type"] = "LTPO OLED, 120Hz, HDR10+, 3000 nits (peak)"
	specs["Resolution"] = "1344 x 2992 pixels (~486 ppi density)"
	specs["Screen Protection"] = "Corning Gorilla Glass Victus 2"
	specs["Refresh Rate"] = "1-120Hz (LTPO)"
	specs["Build Material"] = "Glass front (Gorilla Glass Victus 2), glass back, aluminum frame"
	specs["Weight"] = "221 g (7.80 oz)"
	specs["Dimensions"] = "162.8 x 76.6 x 8.5 mm (6.41 x 3.02 x 0.33 in)"
	specs["Water Resistance"] = "IP68 dust tight and water resistant (up to 1.5m for 30 min)"
	specs["Network Technology"] = "GSM / CDMA / HSPA / EVDO / LTE / 5G"
	specs["2G Bands"] = "GSM 850 / 900 / 1800 / 1900 MHz"
	specs["3G Bands"] = "HSDPA 800 / 850 / 900 / 1700(AWS) / 1900 / 2100 MHz"
	specs["4G Bands"] = "LTE Band 1, 2, 3, 4, 5, 7, 8, 12, 13, 14, 17, 18, 19, 20, 25, 26, 28, 29, 30, 32, 38, 39, 40, 41, 42, 46, 48, 66, 71"
	specs["5G Bands"] = "n1, n2, n3, n5, n7, n8, n12, n14, n20, n25, n26, n28, n30, n38, n40, n41, n48, n66, n71, n75, n76, n77, n78"
	specs["Wifi Support"] = "Wi-Fi 802.11 a/b/g/n/ac/6e/7, tri-band, Wi-Fi Direct"
	specs["Bluetooth Version"] = "5.3, A2DP, LE, aptX HD"
	specs["Nfc Support"] = "Yes"
	specs["Usb Type"] = "USB Type-C 3.2, DisplayPort 1.4, OTG"
	specs["Rear Camera"] = "50 MP + 48 MP + 48 MP"
	specs["Camera Features"] = "Dual-LED flash, Pixel Shift, Ultra-HDR, panorama, Best Take"
	specs["Camera Video Resolution"] = "4K@24/30/60fps, 1080p@24/30/60/120/240fps, gyro-EIS, OIS, 10-bit HDR"
	specs["Optical Zoom"] = "5x Optical Zoom"
	specs["Front Camera"] = "42 MP, f/2.2, 17mm (ultrawide), PDAF"
	specs["Front Camera Video Resolution"] = "4K@24/30/60fps, 1080p@30/60fps"
	specs["Operating System"] = "Android 15, upgradable to Android 22"
	specs["Battery"] = "5,060 mAh"
	specs["Battery Type"] = "Li-Ion (non-removable)"
	specs["Fast Charging"] = "37W wired, 23W wireless, 5W reverse wireless"
	specs["Charging Speed"] = "37W wired (50% in 28 min), 23W wireless, 5W reverse wireless"
	specs["Wireless Charging"] = "Yes - 23W wireless"
	specs["5G Support"] = "Yes"
	specs["Positioning System"] = "GPS (L1+L5), GLONASS (G1), GALILEO (E1+E5a), BDS (B1I+B1c+B2a+B2b), QZSS (L1+L5)"
	specs["Sensors"] = "Fingerprint (under display, ultrasonic), accelerometer, gyro, proximity, compass, barometer, thermometer (skin temperature)"
	specs["Special Features"] = "Gemini AI assistant, Circle to Search, Magic Eraser, Best Take, Audio Magic Eraser, Call Screen, Hold for Me"
	specs["Sim Card Type"] = "Nano-SIM and eSIM"
	specs["Loudspeaker Quality"] = "Stereo speakers"
	specs["Audio Quality"] = "24-bit/192kHz audio"
	specs["Audio Jack"] = "No"
	specs["Sar Rating"] = "1.15 W/kg (head), 1.36 W/kg (body)"
	specs["Sar Rating Eu"] = "1.17 W/kg (head), 1.24 W/kg (body)"
	specs["Available Colors"] = "Obsidian, Porcelain, Hazel, Rose Quartz"
	specs["Model Variants"] = "GA05844-US, GA05845-US, GA05846-US"
	specs["Announcement Date"] = "August 13, 2024"
	specs["Device Status"] = "Available. Released August 22, 2024"

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
