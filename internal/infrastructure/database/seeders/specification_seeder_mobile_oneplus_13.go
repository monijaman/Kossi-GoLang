package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileOneplus13 seeds specifications/options for product 'oneplus-13'
type SpecificationSeederMobileOneplus13 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileOneplus13 creates a new seeder instance
func NewSpecificationSeederMobileOneplus13() *SpecificationSeederMobileOneplus13 {
	return &SpecificationSeederMobileOneplus13{BaseSeeder: BaseSeeder{name: "Specifications for OnePlus 13"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileOneplus13) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1-120Hz (LTPO)": "১-১২০Hz (এলটিপিও)",
		"100W wired (100% in 36 min), 50W wireless (100% in 56 min), 10W reverse wireless": "১০০W তারযুক্ত (১০০% in ৩৬ মিনিট), ৫০W বেতার (১০০% in ৫৬ মিনিট), ১০W রিভার্স বেতার",
		"100W wired, 50W wireless, 10W reverse wireless": "১০০W তারযুক্ত, ৫০W বেতার, ১০W রিভার্স বেতার",
		"12 GB / 16 GB / 24 GB": "১২ জিবি / ১৬ জিবি / ২৪ GB",
		"1440 x 3168 pixels (~510 ppi density)": "১৪৪০ x ৩১৬৮ পিক্সেল (~৫১০ পিপিআই density)",
		"162.9 x 76.5 x 8.5 mm (glass) / 8.9 mm (leather)": "১৬২.৯ x ৭৬.৫ x ৮.৫ মিমি (গ্লাস) / ৮.৯ মিমি (চামড়া)",
		"213 g (glass) / 210 g (leather) (7.51 oz)": "২১৩ গ্রাম (গ্লাস) / ২১০ গ্রাম (চামড়া) (৭.৫১ oz)",
		"24-bit/192kHz audio, Hi-Res Audio": "২৪-bit/১৯২kHz audio, Hi-Res অডিও",
		"256 GB / 512 GB / 1 TB": "২৫৬ জিবি / ৫১২ জিবি / ১ টিবি",
		"2x4.32 GHz + 6x3.53 GHz": "২x৪.৩২ গিগাহার্টজ + ৬x৩.৫৩ গিগাহার্টজ",
		"32 MP, f/2.4, 21mm (wide), Auto-HDR": "৩২ মেগাপিক্সেল, f/২.৪, ২১mm (ওয়াইড), Auথেকে-এইচডিআর",
		"3x Optical Zoom": "৩x অপটিক্যাল জুম",
		"4K@30/60fps, 1080p@30/60fps": "৪K@৩০/৬০fps, ১০৮০p@৩০/৬০fps",
		"5.4, A2DP, LE, aptX HD": "৫.৪, A২DP, LE, aptX HD",
		"50 MP + 50 MP + 50 MP": "৫০ মেগাপিক্সেল + ৫০ মেগাপিক্সেল + ৫০ মেগাপিক্সেল",
		"50MP (wide, f/1.6, OIS) + 50MP (periscope telephoto 3x, f/2.6, OIS) + 50MP (ultrawide, f/2.0)": "৫০MP (ওয়াইড, f/১.৬, OIS) + ৫০MP (periscope টেলিফটো ৩x, f/২.৬, OIS) + ৫০MP (আল্ট্রাওয়াইড, f/২.০)",
		"6,000 mAh": "৬,০০০ এমএএইচ",
		"6.82 inches": "৬.৮২ ইঞ্চি",
		"8K@24fps, 4K@30/60fps, 1080p@30/60/240fps, Auto HDR, gyro-EIS, Dolby Vision": "৮K@২৪fps, ৪K@৩০/৬০fps, ১০৮০p@৩০/৬০/২৪০fps, Auথেকে এইচডিআর, gyro-EIS, ডলবি ভিশন",
		"Adreno 830": "অ্যাড্রেনো ৮৩০",
		"Android 15, OxygenOS 15 (Global) / ColorOS 15 (China)": "অ্যান্ড্রয়েড ১৫, OxygenOS ১৫ (Global) / কালারওএস ১৫ (China)",
		"Available. Released January 07, 2025": "উপলব্ধ. প্রকাশিত জানুয়ারি ০৭, ২০২৫",
		"CPH2649, PJZ110": "CPH২৬৪৯, PJZ১১০",
		"Fingerprint (under display, optical), accelerometer, gyro, proximity, compass, barometer": "ফিঙ্গারপ্রিন্ট (ডিসপ্লের নিচে, optical), অ্যাক্সিলেরোমিটার, জাইরো, প্রক্সিমিটি, কম্পাস, ব্যারোমিটার",
		"GPS (L1+L5), GLONASS (G1), BDS (B1I+B1c+B2a+B2b), GALILEO (E1+E5a+E5b), QZSS (L1+L5), NavIC (L5)": "জিপিএস (L১+L৫), GLONASS (G১), BDS (B১I+B১c+B২a+B২b), GALILEO (E১+E৫a+E৫b), QZSS (L১+L৫), NavIC (L৫)",
		"GSM / CDMA / HSPA / LTE / 5G": "জিএসএম / CDMA / এইচএসপিএ / এলটিই / ৫G",
		"GSM 850 / 900 / 1800 / 1900 MHz": "জিএসএম ৮৫০ / ৯০০ / ১৮০০ / ১৯০০ মেগাহার্টজ",
		"Glass front (Ceramic Guard), glass back or eco leather back, Aluminum frame": "গ্লাস সামনে (Ceramic Guard), গ্লাস পেছনে বা eco চামড়া পেছনে, অ্যালুমিনিয়াম ফ্রেম",
		"HSDPA 800 / 850 / 900 / 1700(AWS) / 1900 / 2100 MHz": "HSDPA ৮০০ / ৮৫০ / ৯০০ / ১৭০০(AWS) / ১৯০০ / ২১০০ মেগাহার্টজ",
		"IP68/IP69 dust tight and water resistant": "IP৬৮/IP৬৯ ধুলো প্রতিরোধী এবং জল প্রতিরোধী",
		"LTE Band 1, 2, 3, 4, 5, 7, 8, 12, 13, 17, 18, 19, 20, 25, 26, 28, 32, 34, 38, 39, 40, 41, 42, 46, 48, 66, 71": "এলটিই Bএবং ১, ২, ৩, ৪, ৫, ৭, ৮, ১২, ১৩, ১৭, ১৮, ১৯, ২০, ২৫, ২৬, ২৮, ৩২, ৩৪, ৩৮, ৩৯, ৪০, ৪১, ৪২, ৪৬, ৪৮, ৬৬, ৭১",
		"LTPO AMOLED, 1B colors, 120Hz, Dolby Vision, HDR10+, 4500 nits (peak)": "এলটিপিও অ্যামোলেড, ১B রঙ, ১২০Hz, ডলবি ভিশন, এইচডিআর১০+, ৪৫০০ nits (পিক)",
		"Octa-core (2x4.32 GHz Oryon V2 Phoenix L + 6x3.53 GHz Oryon V2 Phoenix M)": "অক্টা-কোর (২x৪.৩২ গিগাহার্টজ Oryon V২ Phoenix L + ৬x৩.৫৩ গিগাহার্টজ Oryon V২ Phoenix M)",
		"October 31, 2024": "অক্টোবর ৩১, ২০২৪",
		"Qualcomm SM8750 Snapdragon 8 Elite (3 nm)": "কোয়ালকম SM৮৭৫০ স্ন্যাপড্রাগন ৮ এলিট (৩ ন্যানোমিটার)",
		"Si/C 6000 mAh (non-removable)": "Si/C ৬০০০ এমএএইচ (অপসারণযোগ্য নয়)",
		"Snapdragon 8 Elite": "স্ন্যাপড্রাগন ৮ এলিট",
		"USB Type-C 3.2, DisplayPort 1.4, OTG": "ইউএসবি টাইপ-সি ৩.২, DisplayPবাt ১.৪, OTG",
		"Wi-Fi 802.11 a/b/g/n/ac/6e/7, tri-band, Wi-Fi Direct": "ওয়াই-ফাই ৮০২.১১ a/b/g/n/ac/৬e/৭, tri-bএবং, ওয়াই-ফাই Direct",
		"Yes - 50W AIRVOOC wireless": "Yes - ৫০W এআইRVOOC বেতার",
		"n1, n3, n5, n7, n8, n20, n28, n38, n40, n41, n77, n78": "n১, n৩, n৫, n৭, n৮, n২০, n২৮, n৩৮, n৪০, n৪১, n৭৭, n৭৮",
	}
}

// Seed inserts specification records for the product identified by slug 'oneplus-13'
func (s *SpecificationSeederMobileOneplus13) Seed(db *gorm.DB) error {
	productSlug := "oneplus-13"

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

	// Override model-specific values for OnePlus 13
	specs["Display Size"] = "6.82 inches"
	specs["Processor"] = "Snapdragon 8 Elite"
	specs["Chipset"] = "Qualcomm SM8750 Snapdragon 8 Elite (3 nm)"
	specs["Cpu Type"] = "Octa-core (2x4.32 GHz Oryon V2 Phoenix L + 6x3.53 GHz Oryon V2 Phoenix M)"
	specs["Gpu Type"] = "Adreno 830"
	specs["Processor Speed"] = "2x4.32 GHz + 6x3.53 GHz"
	specs["Ram"] = "12 GB / 16 GB / 24 GB"
	specs["Storage"] = "256 GB / 512 GB / 1 TB"
	specs["Display Type"] = "LTPO AMOLED, 1B colors, 120Hz, Dolby Vision, HDR10+, 4500 nits (peak)"
	specs["Resolution"] = "1440 x 3168 pixels (~510 ppi density)"
	specs["Screen Protection"] = "Ceramic Guard"
	specs["Refresh Rate"] = "1-120Hz (LTPO)"
	specs["Build Material"] = "Glass front (Ceramic Guard), glass back or eco leather back, Aluminum frame"
	specs["Weight"] = "213 g (glass) / 210 g (leather) (7.51 oz)"
	specs["Dimensions"] = "162.9 x 76.5 x 8.5 mm (glass) / 8.9 mm (leather)"
	specs["Water Resistance"] = "IP68/IP69 dust tight and water resistant"
	specs["Network Technology"] = "GSM / CDMA / HSPA / LTE / 5G"
	specs["2G Bands"] = "GSM 850 / 900 / 1800 / 1900 MHz"
	specs["3G Bands"] = "HSDPA 800 / 850 / 900 / 1700(AWS) / 1900 / 2100 MHz"
	specs["4G Bands"] = "LTE Band 1, 2, 3, 4, 5, 7, 8, 12, 13, 17, 18, 19, 20, 25, 26, 28, 32, 34, 38, 39, 40, 41, 42, 46, 48, 66, 71"
	specs["5G Bands"] = "n1, n3, n5, n7, n8, n20, n28, n38, n40, n41, n77, n78"
	specs["Wifi Support"] = "Wi-Fi 802.11 a/b/g/n/ac/6e/7, tri-band, Wi-Fi Direct"
	specs["Bluetooth Version"] = "5.4, A2DP, LE, aptX HD"
	specs["Nfc Support"] = "Yes"
	specs["Usb Type"] = "USB Type-C 3.2, DisplayPort 1.4, OTG"
	specs["Rear Camera"] = "50 MP + 50 MP + 50 MP"
	specs["Quad Camera Setup"] = "50MP (wide, f/1.6, OIS) + 50MP (periscope telephoto 3x, f/2.6, OIS) + 50MP (ultrawide, f/2.0)"
	specs["Camera Features"] = "Hasselblad Color Calibration, LED flash, HDR, panorama"
	specs["Camera Video Resolution"] = "8K@24fps, 4K@30/60fps, 1080p@30/60/240fps, Auto HDR, gyro-EIS, Dolby Vision"
	specs["Optical Zoom"] = "3x Optical Zoom"
	specs["Front Camera"] = "32 MP, f/2.4, 21mm (wide), Auto-HDR"
	specs["Front Camera Video Resolution"] = "4K@30/60fps, 1080p@30/60fps"
	specs["Operating System"] = "Android 15, OxygenOS 15 (Global) / ColorOS 15 (China)"
	specs["Battery"] = "6,000 mAh"
	specs["Battery Type"] = "Si/C 6000 mAh (non-removable)"
	specs["Fast Charging"] = "100W wired, 50W wireless, 10W reverse wireless"
	specs["Charging Speed"] = "100W wired (100% in 36 min), 50W wireless (100% in 56 min), 10W reverse wireless"
	specs["Wireless Charging"] = "Yes - 50W AIRVOOC wireless"
	specs["5G Support"] = "Yes"
	specs["Positioning System"] = "GPS (L1+L5), GLONASS (G1), BDS (B1I+B1c+B2a+B2b), GALILEO (E1+E5a+E5b), QZSS (L1+L5), NavIC (L5)"
	specs["Sensors"] = "Fingerprint (under display, optical), accelerometer, gyro, proximity, compass, barometer"
	specs["Special Features"] = "Alert slider, Dolby Atmos, OnePlus AI features"
	specs["Sim Card Type"] = "Dual Nano-SIM"
	specs["Loudspeaker Quality"] = "Stereo speakers, Dolby Atmos"
	specs["Audio Quality"] = "24-bit/192kHz audio, Hi-Res Audio"
	specs["Audio Jack"] = "No"
	specs["Available Colors"] = "Midnight Ocean, Black Eclipse, Arctic Dawn, White Dew (China only)"
	specs["Model Variants"] = "CPH2649, PJZ110"
	specs["Announcement Date"] = "October 31, 2024"
	specs["Device Status"] = "Available. Released January 07, 2025"

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
