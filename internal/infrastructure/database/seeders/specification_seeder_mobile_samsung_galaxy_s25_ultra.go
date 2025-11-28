package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileSamsungGalaxyS25Ultra seeds specifications/options for product 'samsung-galaxy-s25-ultra'
type SpecificationSeederMobileSamsungGalaxyS25Ultra struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSamsungGalaxyS25Ultra creates a new seeder instance
func NewSpecificationSeederMobileSamsungGalaxyS25Ultra() *SpecificationSeederMobileSamsungGalaxyS25Ultra {
	return &SpecificationSeederMobileSamsungGalaxyS25Ultra{BaseSeeder: BaseSeeder{name: "Specifications for Samsung Galaxy S25 Ultra"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSamsungGalaxyS25Ultra) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1.25 W/kg (head), 1.42 W/kg (body)": "১.২৫ W/kg (head), ১.৪২ W/kg (body)",
		"1.26 W/kg (head), 0.64 W/kg (body)": "১.২৬ W/kg (head), ০.৬৪ W/kg (body)",
		"12 GB / 16 GB": "১২ জিবি / ১৬ GB",
		"12 MP, f/2.2, 26mm (wide), 1.12µm, dual pixel PDAF": "১২ মেগাপিক্সেল, f/২.২, ২৬mm (ওয়াইড), ১.১২µm, dual পিক্সেল PDAF",
		"120Hz": "১২০Hz",
		"1440 x 3120 pixels, 19.5:9 ratio (~498 ppi density)": "১৪৪০ x ৩১২০ পিক্সেল, ১৯.৫:৯ ratio (~৪৯৮ পিপিআই density)",
		"162.8 x 77.6 x 8.2 mm (6.41 x 3.06 x 0.32 in)": "১৬২.৮ x ৭৭.৬ x ৮.২ মিমি (৬.৪১ x ৩.০৬ x ০.৩২ in)",
		"200 MP + 10 MP + 50 MP + 50 MP": "২০০ মেগাপিক্সেল + ১০ মেগাপিক্সেল + ৫০ মেগাপিক্সেল + ৫০ মেগাপিক্সেল",
		"200MP (wide, f/1.7) + 10MP (telephoto 3x, f/2.4) + 50MP (periscope 5x, f/3.4) + 50MP (ultrawide, f/1.9)": "২০০MP (ওয়াইড, f/১.৭) + ১০MP (টেলিফটো ৩x, f/২.৪) + ৫০MP (periscope ৫x, f/৩.৪) + ৫০MP (আল্ট্রাওয়াইড, f/১.৯)",
		"218 g (7.69 oz)": "২১৮ গ্রাম (৭.৬৯ oz)",
		"256 GB / 512 GB / 1 TB": "২৫৬ জিবি / ৫১২ জিবি / ১ টিবি",
		"2x4.47 GHz (Performance) + 6x3.53 GHz (Efficiency)": "২x৪.৪৭ গিগাহার্টজ (Perজন্যmance) + ৬x৩.৫৩ গিগাহার্টজ (Efficiency)",
		"3x + 5x Optical Zoom": "৩x + ৫x অপটিক্যাল জুম",
		"45W wired (PD3.0), 65% in 30 min": "৪৫W তারযুক্ত (PD৩.০), ৬৫% in ৩০ মিনিট",
		"45W wired (PD3.0), 65% in 30 min; 15W wireless (Qi2 Ready); 4.5W reverse wireless": "৪৫W তারযুক্ত (PD৩.০), ৬৫% in ৩০ মিনিট; ১৫W বেতার (Qi২ Ready); ৪.৫W রিভার্স বেতার",
		"4K@30/60fps, 1080p@30fps": "৪K@৩০/৬০fps, ১০৮০p@৩০fps",
		"5.4, A2DP, LE": "৫.৪, A২DP, LE",
		"5000 mAh": "৫০০০ এমএএইচ",
		"6.9 inches": "৬.৯ ইঞ্চি",
		"8K@24/30fps, 4K@30/60/120fps, 1080p@30/60/120/240fps, 10-bit HDR, HDR10+, stereo sound rec., gyro-EIS": "৮K@২৪/৩০fps, ৪K@৩০/৬০/১২০fps, ১০৮০p@৩০/৬০/১২০/২৪০fps, ১০-bit এইচডিআর, এইচডিআর১০+, stereo sound rec., gyro-EIS",
		"Adreno 830 (1200 MHz)": "অ্যাড্রেনো ৮৩০ (১২০০ মেগাহার্টজ)",
		"Android 15, One UI 7": "অ্যান্ড্রয়েড ১৫, ওয়ান ইউআই ৭",
		"Available. Released February 03, 2025": "উপলব্ধ. প্রকাশিত ফেব্রুয়ারি ০৩, ২০২৫",
		"Corning Gorilla Armor 2, Mohs level 6; DX anti-reflective coating": "Cবাning Gবাilla Armবা ২, Mohs level ৬; DX anti-reflective coating",
		"Dynamic LTPO AMOLED 2X, 120Hz, HDR10+, 2600 nits (peak)": "Dynamic এলটিপিও অ্যামোলেড ২X, ১২০Hz, এইচডিআর১০+, ২৬০০ nits (পিক)",
		"Fingerprint (under display, ultrasonic), accelerometer, gyro, proximity, compass, barometer; UWB support": "ফিঙ্গারপ্রিন্ট (ডিসপ্লের নিচে, আল্ট্রাসনিক), অ্যাক্সিলেরোমিটার, জাইরো, প্রক্সিমিটি, কম্পাস, ব্যারোমিটার; UWB suppবাt",
		"GSM / CDMA / HSPA / EVDO / LTE / 5G": "জিএসএম / CDMA / এইচএসপিএ / EVDO / এলটিই / ৫G",
		"GSM 850 / 900 / 1800 / 1900 MHz": "জিএসএম ৮৫০ / ৯০০ / ১৮০০ / ১৯০০ মেগাহার্টজ",
		"Glass front (Gorilla Armor 2), glass back (Victus 2), titanium frame (grade 5)": "গ্লাস সামনে (Gবাilla Armবা ২), গ্লাস পেছনে (Victus ২), টাইটানিয়াম ফ্রেম (grade ৫)",
		"HSPA 850 / 900 / 1900 / 2100 MHz; EVDO Rev. A 800 / 1900 / 2100 MHz": "এইচএসপিএ ৮৫০ / ৯০০ / ১৯০০ / ২১০০ মেগাহার্টজ; EVDO Rev. A ৮০০ / ১৯০০ / ২১০০ মেগাহার্টজ",
		"IP68 dust tight and water resistant (immersible up to 1.5m for 30 min)": "IP৬৮ ধুলো প্রতিরোধী এবং জল প্রতিরোধী (immersible পর্যন্ত ১.৫m জন্য ৩০ min)",
		"January 22, 2025": "জানুয়ারি ২২, ২০২৫",
		"LTE Band 1, 2, 3, 4, 5, 7, 8, 12, 13, 14, 17, 20, 32, 66": "এলটিই Bএবং ১, ২, ৩, ৪, ৫, ৭, ৮, ১২, ১৩, ১৪, ১৭, ২০, ৩২, ৬৬",
		"Li-Ion 5000 mAh": "Li-Ion ৫০০০ mAh",
		"Nano-SIM + Nano-SIM + eSIM + eSIM (max 2 at a time)": "ন্যানো-সিম + ন্যানো-সিম + ই-সিম + ই-সিম (max ২ at a time)",
		"Octa-core (2x4.47 GHz Oryon V2 Phoenix L + 6x3.53 GHz Oryon V2 Phoenix M)": "অক্টা-কোর (২x৪.৪৭ গিগাহার্টজ Oryon V২ Phoenix L + ৬x৩.৫৩ গিগাহার্টজ Oryon V২ Phoenix M)",
		"Qualcomm SM8750-AC Snapdragon 8 Elite (3 nm)": "কোয়ালকম SM৮৭৫০-AC স্ন্যাপড্রাগন ৮ এলিট (৩ ন্যানোমিটার)",
		"SM-S938B, SM-S938B/DS, SM-S938U, SM-S938U1, SM-S938W, SM-S938N, SM-S9380, SM-S938E, SM-S938E/DS": "SM-S৯৩৮B, SM-S৯৩৮B/DS, SM-S৯৩৮U, SM-S৯৩৮U১, SM-S৯৩৮W, SM-S৯৩৮N, SM-S৯৩৮০, SM-S৯৩৮E, SM-S৯৩৮E/DS",
		"Samsung DeX, Samsung Wireless DeX, Ultra Wideband (UWB) support": "Samsung DeX, Samsung বেতার DeX, আল্ট্রা Widebএবং (UWB) suppবাt",
		"Snapdragon 8 Elite": "স্ন্যাপড্রাগন ৮ এলিট",
		"Titanium Silver Blue, Titanium Black, Titanium White Silver, Titanium Gray, Titanium Jade Green, Titanium Jet Black, Titanium Pink Gold": "টাইটানিয়াম রূপালী নীল, টাইটানিয়াম কালো, টাইটানিয়াম সাদা রূপালী, টাইটানিয়াম ধূসর, টাইটানিয়াম জেড সবুজ, টাইটানিয়াম জেট কালো, টাইটানিয়াম গোলাপী সোনালী",
		"UFS 4.0": "UFS ৪.০",
		"USB Type-C 3.2, DisplayPort 1.2, OTG": "ইউএসবি টাইপ-সি ৩.২, DisplayPবাt ১.২, OTG",
		"Wi-Fi 802.11 a/b/g/n/ac/6e/7, tri-band, Wi-Fi Direct": "ওয়াই-ফাই ৮০২.১১ a/b/g/n/ac/৬e/৭, tri-bএবং, ওয়াই-ফাই Direct",
		"Yes - 15W wireless (Qi2 Ready); 4.5W reverse wireless": "হ্যাঁ - ১৫W বেতার (Qi২ Ready); ৪.৫W রিভার্স বেতার",
		"Yes, stereo speakers": "হ্যাঁ, স্টেরিও স্পিকার",
		"n1, n3, n5, n7, n8, n20, n28, n32, n38, n40, n41, n48, n77, n78, n79": "n১, n৩, n৫, n৭, n৮, n২০, n২৮, n৩২, n৩৮, n৪০, n৪১, n৪৮, n৭৭, n৭৮, n৭৯",
	}
}

// Seed inserts specification records for the product identified by slug 'samsung-galaxy-s25-ultra'
func (s *SpecificationSeederMobileSamsungGalaxyS25Ultra) Seed(db *gorm.DB) error {
	productSlug := "samsung-galaxy-s25-ultra"

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

	// Override model-specific values for Samsung Galaxy S25 Ultra
	specs["Display Size"] = "6.9 inches"
	specs["Processor"] = "Snapdragon 8 Elite"
	specs["Chipset"] = "Qualcomm SM8750-AC Snapdragon 8 Elite (3 nm)"
	specs["Cpu Type"] = "Octa-core (2x4.47 GHz Oryon V2 Phoenix L + 6x3.53 GHz Oryon V2 Phoenix M)"
	specs["Gpu Type"] = "Adreno 830 (1200 MHz)"
	specs["Processor Speed"] = "2x4.47 GHz (Performance) + 6x3.53 GHz (Efficiency)"
	specs["Ram"] = "12 GB / 16 GB"
	specs["Storage"] = "256 GB / 512 GB / 1 TB"
	specs["Internal Memory Capacity"] = "UFS 4.0"
	specs["Card Slot Type"] = "No microSD card slot"
	specs["Display Type"] = "Dynamic LTPO AMOLED 2X, 120Hz, HDR10+, 2600 nits (peak)"
	specs["Resolution"] = "1440 x 3120 pixels, 19.5:9 ratio (~498 ppi density)"
	specs["Screen Protection"] = "Corning Gorilla Armor 2, Mohs level 6; DX anti-reflective coating"
	specs["Refresh Rate"] = "120Hz"
	specs["Build Material"] = "Glass front (Gorilla Armor 2), glass back (Victus 2), titanium frame (grade 5)"
	specs["Weight"] = "218 g (7.69 oz)"
	specs["Dimensions"] = "162.8 x 77.6 x 8.2 mm (6.41 x 3.06 x 0.32 in)"
	specs["Water Resistance"] = "IP68 dust tight and water resistant (immersible up to 1.5m for 30 min)"
	specs["Network Technology"] = "GSM / CDMA / HSPA / EVDO / LTE / 5G"
	specs["2G Bands"] = "GSM 850 / 900 / 1800 / 1900 MHz"
	specs["3G Bands"] = "HSPA 850 / 900 / 1900 / 2100 MHz; EVDO Rev. A 800 / 1900 / 2100 MHz"
	specs["4G Bands"] = "LTE Band 1, 2, 3, 4, 5, 7, 8, 12, 13, 14, 17, 20, 32, 66"
	specs["5G Bands"] = "n1, n3, n5, n7, n8, n20, n28, n32, n38, n40, n41, n48, n77, n78, n79"
	specs["Wifi Support"] = "Wi-Fi 802.11 a/b/g/n/ac/6e/7, tri-band, Wi-Fi Direct"
	specs["Bluetooth Version"] = "5.4, A2DP, LE"
	specs["Nfc Support"] = "Yes"
	specs["Usb Type"] = "USB Type-C 3.2, DisplayPort 1.2, OTG"
	specs["Rear Camera"] = "200 MP + 10 MP + 50 MP + 50 MP"
	specs["Quad Camera Setup"] = "200MP (wide, f/1.7) + 10MP (telephoto 3x, f/2.4) + 50MP (periscope 5x, f/3.4) + 50MP (ultrawide, f/1.9)"
	specs["Camera Features"] = "Laser AF, Best Face, LED flash, auto-HDR, panorama"
	specs["Camera Video Resolution"] = "8K@24/30fps, 4K@30/60/120fps, 1080p@30/60/120/240fps, 10-bit HDR, HDR10+, stereo sound rec., gyro-EIS"
	specs["Optical Zoom"] = "3x + 5x Optical Zoom"
	specs["Front Camera"] = "12 MP, f/2.2, 26mm (wide), 1.12µm, dual pixel PDAF"
	specs["Front Camera Video Resolution"] = "4K@30/60fps, 1080p@30fps"
	specs["Operating System"] = "Android 15, One UI 7"
	specs["Battery"] = "5000 mAh"
	specs["Battery Type"] = "Li-Ion 5000 mAh"
	specs["Fast Charging"] = "45W wired (PD3.0), 65% in 30 min"
	specs["Charging Speed"] = "45W wired (PD3.0), 65% in 30 min; 15W wireless (Qi2 Ready); 4.5W reverse wireless"
	specs["Wireless Charging"] = "Yes - 15W wireless (Qi2 Ready); 4.5W reverse wireless"
	specs["5G Support"] = "Yes"
	specs["Positioning System"] = "GPS, GLONASS, BDS, GALILEO, QZSS"
	specs["Sensors"] = "Fingerprint (under display, ultrasonic), accelerometer, gyro, proximity, compass, barometer; UWB support"
	specs["Special Features"] = "Samsung DeX, Samsung Wireless DeX, Ultra Wideband (UWB) support"
	specs["Sim Card Type"] = "Nano-SIM + Nano-SIM + eSIM + eSIM (max 2 at a time)"
	specs["Loudspeaker Quality"] = "Yes, stereo speakers"
	specs["Audio Quality"] = "High-bitrate audio support"
	specs["Audio Jack"] = "No"
	specs["Sar Rating"] = "1.26 W/kg (head), 0.64 W/kg (body)"
	specs["Sar Rating Eu"] = "1.25 W/kg (head), 1.42 W/kg (body)"
	specs["Available Colors"] = "Titanium Silver Blue, Titanium Black, Titanium White Silver, Titanium Gray, Titanium Jade Green, Titanium Jet Black, Titanium Pink Gold"
	specs["Model Variants"] = "SM-S938B, SM-S938B/DS, SM-S938U, SM-S938U1, SM-S938W, SM-S938N, SM-S9380, SM-S938E, SM-S938E/DS"
	specs["Announcement Date"] = "January 22, 2025"
	specs["Device Status"] = "Available. Released February 03, 2025"

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
