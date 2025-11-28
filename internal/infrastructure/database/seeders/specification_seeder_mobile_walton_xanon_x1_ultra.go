package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileWaltonXanonX1Ultra seeds specifications/options for product 'walton-xanon-x1-ultra'
type SpecificationSeederMobileWaltonXanonX1Ultra struct {
	BaseSeeder
}

// NewSpecificationSeederMobileWaltonXanonX1Ultra creates a new seeder instance
func NewSpecificationSeederMobileWaltonXanonX1Ultra() *SpecificationSeederMobileWaltonXanonX1Ultra {
	return &SpecificationSeederMobileWaltonXanonX1Ultra{BaseSeeder: BaseSeeder{name: "Specifications for Walton XANON X1 Ultra"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileWaltonXanonX1Ultra) getBanglaTranslations() map[string]string {
	return map[string]string{
		"108 MP + 50 MP + 20 MP": "১০৮ মেগাপিক্সেল + ৫০ মেগাপিক্সেল + ২০ মেগাপিক্সেল",
		"1080 x 2412 pixels (~394 ppi density)": "১০৮০ x ২৪১২ পিক্সেল (~৩৯৪ পিপিআই density)",
		"12 GB": "১২ GB",
		"120Hz": "১২০Hz",
		"161.8 x 73.9 x 8.4 mm (6.37 x 2.91 x 0.33 in)": "১৬১.৮ x ৭৩.৯ x ৮.৪ মিমি (৬.৩৭ x ২.৯১ x ০.৩৩ in)",
		"195 g (6.88 oz)": "১৯৫ গ্রাম (৬.৮৮ oz)",
		"1x3.25 GHz + 3x2.85 GHz + 4x2.0 GHz": "১x৩.২৫ গিগাহার্টজ + ৩x২.৮৫ গিগাহার্টজ + ৪x২.০ গিগাহার্টজ",
		"24-bit/192kHz audio, Hi-Res Audio": "২৪-bit/১৯২kHz audio, Hi-Res অডিও",
		"256 GB": "২৫৬ GB",
		"32 MP, f/2.4, (wide), HDR": "৩২ মেগাপিক্সেল, f/২.৪, (ওয়াইড), এইচডিআর",
		"3x Optical Zoom": "৩x অপটিক্যাল জুম",
		"4K@30/60fps, 1080p@30/60/120fps, gyro-EIS": "৪K@৩০/৬০fps, ১০৮০p@৩০/৬০/১২০fps, জাইরো-EIS",
		"4K@30fps, 1080p@30/60fps": "৪K@৩০fps, ১০৮০p@৩০/৬০fps",
		"5,000 mAh": "৫,০০০ এমএএইচ",
		"5.3, A2DP, LE, aptX HD": "৫.৩, A২DP, LE, aptX HD",
		"6.7 inches": "৬.৭ ইঞ্চি",
		"66W wired": "৬৬W তারযুক্ত",
		"66W wired (100% in 45 min)": "৬৬W তারযুক্ত (১০০% in ৪৫ মিনিট)",
		"AMOLED, 120Hz, HDR10+, 1200 nits (peak)": "অ্যামোলেড, ১২০Hz, এইচডিআর১০+, ১২০০ nits (পিক)",
		"Android 14, Walton WaltonOS": "অ্যান্ড্রয়েড ১৪, Walথেকেn WalথেকেnOS",
		"December 2024": "ডিসেম্বর ২০২৪",
		"Dual stereo speakers, Hi-Res Audio": "ডুয়াল স্টেরিও স্পিকার, Hi-Res অডিও",
		"Fingerprint (under display, optical), accelerometer, gyro, proximity, compass": "ফিঙ্গারপ্রিন্ট (ডিসপ্লের নিচে, optical), অ্যাক্সিলেরোমিটার, জাইরো, প্রক্সিমিটি, কম্পাস",
		"GSM / HSPA / LTE / 5G": "জিএসএম / এইচএসপিএ / এলটিই / ৫G",
		"GSM 850 / 900 / 1800 / 1900 MHz": "জিএসএম ৮৫০ / ৯০০ / ১৮০০ / ১৯০০ মেগাহার্টজ",
		"Glass front (Gorilla Glass Victus), glass back, aluminum frame": "গ্লাস সামনে (গরিলা গ্লাস Victus), গ্লাস পেছনে, অ্যালুমিনিয়াম ফ্রেম",
		"HSDPA 850 / 900 / 1900 / 2100 MHz": "HSDPA ৮৫০ / ৯০০ / ১৯০০ / ২১০০ মেগাহার্টজ",
		"IP67 dust tight and water resistant (up to 1m for 30 min)": "IP৬৭ ধুলো প্রতিরোধী এবং জল প্রতিরোধী (পর্যন্ত ১m জন্য ৩০ min)",
		"LTE Band 1, 3, 5, 7, 8, 20, 28, 38, 40, 41": "এলটিই Bএবং ১, ৩, ৫, ৭, ৮, ২০, ২৮, ৩৮, ৪০, ৪১",
		"Li-Po (non-removable)": "লি-পো (অপসারণযোগ্য নয়)",
		"Mali-G720 Immortalis MC12": "মালি-G৭২০ Immবাtalis MC১২",
		"MediaTek Dimensity 9300": "মিডিয়াটেক ডাইমেনসিটি ৯৩০০",
		"MediaTek Dimensity 9300 (4 nm)": "মিডিয়াটেক ডাইমেনসিটি ৯৩০০ (৪ ন্যানোমিটার)",
		"Midnight Black, Pearl White, Ocean Blue": "মিডনাইট কালো, পার্ল সাদা, ওশান নীল",
		"Octa-core (1x3.25 GHz Cortex-X4 & 3x2.85 GHz Cortex-X4 & 4x2.0 GHz Cortex-A720)": "অক্টা-কোর (১x৩.২৫ গিগাহার্টজ Cবাtex-X৪ & ৩x২.৮৫ গিগাহার্টজ Cবাtex-X৪ & ৪x২.০ গিগাহার্টজ Cবাtex-A৭২০)",
		"UFS 4.0": "UFS ৪.০",
		"USB Type-C 3.1, OTG": "ইউএসবি টাইপ-সি ৩.১, OTG",
		"WX1U-001": "WX১U-০০১",
		"Wi-Fi 802.11 a/b/g/n/ac/6e, dual-band, Wi-Fi Direct": "ওয়াই-ফাই ৮০২.১১ a/b/g/n/ac/৬e, dual-bএবং, ওয়াই-ফাই Direct",
		"microSD card slot up to 1TB": "মাইক্রোএসডি card slot পর্যন্ত ১TB",
		"n1, n3, n5, n7, n8, n20, n28, n38, n40, n41, n77, n78": "n১, n৩, n৫, n৭, n৮, n২০, n২৮, n৩৮, n৪০, n৪১, n৭৭, n৭৮",
	}
}

// Seed inserts specification records for the product identified by slug 'walton-xanon-x1-ultra'
func (s *SpecificationSeederMobileWaltonXanonX1Ultra) Seed(db *gorm.DB) error {
	productSlug := "walton-xanon-x1-ultra"

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

	// Override model-specific values for Walton XANON X1 Ultra
	specs["Display Size"] = "6.7 inches"
	specs["Processor"] = "MediaTek Dimensity 9300"
	specs["Chipset"] = "MediaTek Dimensity 9300 (4 nm)"
	specs["Cpu Type"] = "Octa-core (1x3.25 GHz Cortex-X4 & 3x2.85 GHz Cortex-X4 & 4x2.0 GHz Cortex-A720)"
	specs["Gpu Type"] = "Mali-G720 Immortalis MC12"
	specs["Processor Speed"] = "1x3.25 GHz + 3x2.85 GHz + 4x2.0 GHz"
	specs["Ram"] = "12 GB"
	specs["Storage"] = "256 GB"
	specs["Internal Memory Capacity"] = "UFS 4.0"
	specs["Card Slot Type"] = "microSD card slot up to 1TB"
	specs["Display Type"] = "AMOLED, 120Hz, HDR10+, 1200 nits (peak)"
	specs["Resolution"] = "1080 x 2412 pixels (~394 ppi density)"
	specs["Screen Protection"] = "Corning Gorilla Glass Victus"
	specs["Refresh Rate"] = "120Hz"
	specs["Build Material"] = "Glass front (Gorilla Glass Victus), glass back, aluminum frame"
	specs["Weight"] = "195 g (6.88 oz)"
	specs["Dimensions"] = "161.8 x 73.9 x 8.4 mm (6.37 x 2.91 x 0.33 in)"
	specs["Water Resistance"] = "IP67 dust tight and water resistant (up to 1m for 30 min)"
	specs["Network Technology"] = "GSM / HSPA / LTE / 5G"
	specs["2G Bands"] = "GSM 850 / 900 / 1800 / 1900 MHz"
	specs["3G Bands"] = "HSDPA 850 / 900 / 1900 / 2100 MHz"
	specs["4G Bands"] = "LTE Band 1, 3, 5, 7, 8, 20, 28, 38, 40, 41"
	specs["5G Bands"] = "n1, n3, n5, n7, n8, n20, n28, n38, n40, n41, n77, n78"
	specs["Wifi Support"] = "Wi-Fi 802.11 a/b/g/n/ac/6e, dual-band, Wi-Fi Direct"
	specs["Bluetooth Version"] = "5.3, A2DP, LE, aptX HD"
	specs["Nfc Support"] = "Yes"
	specs["Usb Type"] = "USB Type-C 3.1, OTG"
	specs["Rear Camera"] = "108 MP + 50 MP + 20 MP"
	specs["Camera Features"] = "LED flash, HDR, panorama"
	specs["Camera Video Resolution"] = "4K@30/60fps, 1080p@30/60/120fps, gyro-EIS"
	specs["Optical Zoom"] = "3x Optical Zoom"
	specs["Front Camera"] = "32 MP, f/2.4, (wide), HDR"
	specs["Front Camera Video Resolution"] = "4K@30fps, 1080p@30/60fps"
	specs["Operating System"] = "Android 14, Walton WaltonOS"
	specs["Battery"] = "5,000 mAh"
	specs["Battery Type"] = "Li-Po (non-removable)"
	specs["Fast Charging"] = "66W wired"
	specs["Charging Speed"] = "66W wired (100% in 45 min)"
	specs["Wireless Charging"] = "No"
	specs["5G Support"] = "Yes"
	specs["Positioning System"] = "GPS, GLONASS, BDS, GALILEO"
	specs["Sensors"] = "Fingerprint (under display, optical), accelerometer, gyro, proximity, compass"
	specs["Special Features"] = "Dual stereo speakers, Hi-Res Audio"
	specs["Sim Card Type"] = "Dual Nano-SIM"
	specs["Loudspeaker Quality"] = "Dual stereo speakers"
	specs["Audio Quality"] = "24-bit/192kHz audio, Hi-Res Audio"
	specs["Audio Jack"] = "No"
	specs["Available Colors"] = "Midnight Black, Pearl White, Ocean Blue"
	specs["Model Variants"] = "WX1U-001"
	specs["Announcement Date"] = "December 2024"
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
