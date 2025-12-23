package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederMobileWaltonXANONX1Ultra seeds specifications/options for product 'walton-xanon-x1-ultra'
type SpecificationSeederMobileWaltonXANONX1Ultra struct {
	BaseSeeder
}

// NewSpecificationSeederMobileWaltonXANONX1Ultra creates a new seeder instance
func NewSpecificationSeederMobileWaltonXANONX1Ultra() *SpecificationSeederMobileWaltonXANONX1Ultra {
	return &SpecificationSeederMobileWaltonXANONX1Ultra{BaseSeeder: BaseSeeder{name: "Specifications for Walton XANON X1 Ultra"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileWaltonXANONX1Ultra) getBanglaTranslations() map[string]string {
	specific := map[string]string{
		"Fingerprint (under display, optical), accelerometer, gyro, proximity, compass": "ফিঙ্গারপ্রিন্ট (আন্ডার ডিসপ্লে, অপটিকাল), অ্যাক্সেলেরোমিটার, জাইরো, প্রক্সিমিটি, কম্পাস",
		"GPS, GLONASS, BDS, GALILEO":                          "জিপিএস, গ্লোনাস, বিডিএস, গ্যালিলিও",
		"5.3, A2DP, LE, aptX HD":                              "৫.৩, এ২ডিপি, এলই, অ্যাপটিএক্স এইচডি",
		"Wi-Fi 802.11 a/b/g/n/ac/6e, dual-band, Wi-Fi Direct": "ওয়াই-ফাই ৮০২.১১ এ/বি/জি/এন/এসি/৬ই, ডুয়াল-ব্যান্ড, ওয়াই-ফাই ডাইরেক্ট",
		"USB Type-C 3.1, OTG":                                 "ইউএসবি টাইপ-সি ৩.১, ওটিজি",
		"Dual Nano-SIM":                                       "ডুয়াল ন্যানো-সিম",
		"No":                                                  "না",
		"Dual stereo speakers":                                "ডুয়াল স্টেরিও স্পিকার",
		"24-bit/192kHz audio, Hi-Res Audio":                   "২৪-বিট/১৯২কেএইচজেড অডিও, হাই-রেজ অডিও",
		"Li-Po (non-removable)":                               "লি-পো (নন-রিমুভেবল)",
		"66W wired (100% in 45 min)":                          "৬৬ডব্লিউ ওয়্যার্ড (১০০% ইন ৪৫ মিন)",
		"microSD card slot up to 1TB":                         "মাইক্রোএসডি কার্ড স্লট আপ টু ১টিবি",
		"UFS 4.0":                                             "ইউএফএস ৪.০",
		"1080 x 2412 pixels (~394 ppi density)":               "১০৮০ x ২৪১২ পিক্সেল (~৩৯৪ পিপিআই ডেনসিটি)",
		"32 MP, f/2.4, (wide), HDR":                           "৩২ এমপি, এফ/২.৪, (ওয়াইড), এইচডিআর",
		"5,000 mAh":                                           "৫,০০০ এমএএইচ",
		"WX1U-001":                                            "ডব্লিউএক্স১ইউ-০০১",
		"6.7 inches":                                          "৬.৭ ইঞ্চি",
		"120Hz":                                               "১২০হার্টজ",
		"161.8 x 73.9 x 8.4 mm (6.37 x 2.91 x 0.33 in)":             "১৬১.৮ x ৭৩.৯ x ৮.৪ মিমি (৬.৩৭ x ২.৯১ x ০.৩৩ ইন)",
		"IP67 dust tight and water resistant (up to 1m for 30 min)": "আইপি৬৭ ডাস্ট টাইট অ্যান্ড ওয়াটার রেসিস্ট্যান্ট (আপ টু ১মি ফর ৩০ মিন)",
		"GSM 850 / 900 / 1800 / 1900 MHz":                           "জিএসএম ৮৫০ / ৯০০ / ১৮০০ / ১৯০০ মেগাহার্টজ",
		"Yes":                                                       "হ্যাঁ",
		"Android 14, Walton WaltonOS":                               "অ্যান্ড্রয়েড ১৪, ওয়ালটন ওয়ালটনওএস",
		"195 g (6.88 oz)":                                           "১৯৫ জি (৬.৮৮ আওজ)",
		"GSM / HSPA / LTE / 5G":                                     "জিএসএম / এইচএসপিএ / এলটিই / ৫জি",
		"n1, n3, n5, n7, n8, n20, n28, n38, n40, n41, n77, n78":     "এন১, এন৩, এন৫, এন৭, এন৮, এন২০, এন২৮, এন৩৮, এন৪০, এন৪১, এন৭৭, এন৭৮",
		"Corning Gorilla Glass Victus":                              "কর্নিং গরিলা গ্লাস ভিক্টাস",
		"LED flash, HDR, panorama":                                  "এলইডি ফ্ল্যাশ, এইচডিআর, প্যানোরামা",
		"4K@30/60fps, 1080p@30/60/120fps, gyro-EIS":                 "৪কে@৩০/৬০এফপিএস, ১০৮০পি@৩০/৬০/১২০এফপিএস, জাইরো-ইআইএস",
		"4K@30fps, 1080p@30/60fps":                                  "৪কে@৩০এফপিএস, ১০৮০পি@৩০/৬০এফপিএস",
		"MediaTek Dimensity 9300":                                   "মিডিয়াটেক ডাইমেনসিটি ৯৩০০",
		"Mali-G720 Immortalis MC12":                                 "মালি-জি৭২০ ইমর্টালিস এমসি১২",
		"LTE Band 1, 3, 5, 7, 8, 20, 28, 38, 40, 41":                "এলটিই ব্যান্ড ১, ৩, ৫, ৭, ৮, ২০, ২৮, ৩৮, ৪০, ৪১",
		"3x Optical Zoom":                                           "৩এক্স অপটিকাল জুম",
		"108 MP + 50 MP + 20 MP":                                    "১০৮ এমপি + ৫০ এমপি + ২০ এমপি",
		"12 GB":                                                     "১২ জিবি",
		"256 GB":                                                    "২৫৬ জিবি",
		"HSDPA 850 / 900 / 1900 / 2100 MHz":                         "এইচএসডিপিএ ৮৫০ / ৯০০ / ১৯০০ / ২১০০ মেগাহার্টজ",
		"1x3.25 GHz Cortex-X4 & 3x2.85 GHz Cortex-X4 & 4x2.0 GHz Cortex-A720": "১এক্স৩.২৫ গিগাহার্টজ কর্টেক্স-এক্স৪ & ৩এক্স২.৮৫ গিগাহার্টজ কর্টেক্স-এক্স৪ & ৪এক্স২.০ গিগাহার্টজ কর্টেক্স-এ৭২০",
		"1x3.25 GHz + 3x2.85 GHz + 4x2.0 GHz":                                 "১এক্স৩.২৫ গিগাহার্টজ + ৩এক্স২.৮৫ গিগাহার্টজ + ৪এক্স২.০ গিগাহার্টজ",
		"Midnight Black, Pearl White, Ocean Blue":                             "মিডনাইট ব্ল্যাক, পার্ল হোয়াইট, ওশান ব্লু",
		"0.43 W/kg (head) 1.29 W/kg (body)":                                   "০.৪৩ ডব্লিউ/কেজি (হেড) ১.২৯ ডব্লিউ/কেজি (বডি)",
		"0.32 W/kg (head) 1.50 W/kg (body)":                                   "০.৩২ ডব্লিউ/কেজি (হেড) ১.৫০ ডব্লিউ/কেজি (বডি)",
		"AMOLED, 120Hz, HDR10+, 1200 nits (peak)":                             "এএমওএলইডি, ১২০হার্টজ, এইচডিআর১০+, ১২০০ নিটস (পিক)",
		"Glass front (Gorilla Glass Victus), glass back, aluminum frame":      "গ্লাস ফ্রন্ট (গরিলা গ্লাস ভিক্টাস), গ্লাস ব্যাক, অ্যালুমিনিয়াম ফ্রেম",
		"Available":                          "উপলব্ধ",
		"Dual stereo speakers, Hi-Res Audio": "ডুয়াল স্টেরিও স্পিকার, হাই-রেজ অডিও",
		"December 2024":                      "ডিসেম্বর ২০২৪",
		"Octa-core (1x3.25 GHz Cortex-X4 & 3x2.85 GHz Cortex-X4 & 4x2.0 GHz Cortex-A720)": "অক্টা-কোর (১x৩.২৫ গিগাহার্টজ কর্টেক্স-X4 & ৩x২.৮৫ গিগাহার্টজ কর্টেক্স-X4 & ৪x২.০ গিগাহার্টজ কর্টেক্স-A720)",
	}
	common := s.GetCommonBanglaTranslations()
	for k, v := range specific {
		common[k] = v
	}
	return common
}

// Seed inserts specification records for the product identified by slug 'walton-xanon-x1-ultra'
func (s *SpecificationSeederMobileWaltonXANONX1Ultra) Seed(db *gorm.DB) error {
	productSlug := "walton-xanon-x1-ultra"
	var prod models.ProductModel
	if err := db.Where("slug = ?", productSlug).First(&prod).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil
		}
		return err
	}
	productID := prod.ID
	existingkeyMapping := map[string]uint{
		// Core display
		"display_size":      66, // Screen Size
		"display_type":      27, // Display Type
		"resolution":        62, // Resolution
		"refresh_rate":      61, // Refresh Rate
		"screen_protection": 65, // Screen Protection
		// Performance
		"processor":       535, // processor (lowercase)
		"chipset":         19,  // Chipset
		"cpu_type":        21,  // CPU Type
		"gpu_type":        33,  // GPU Type
		"processor_speed": 57,  // Processor Speed
		// Memory & storage
		"ram":                      60, // RAM
		"storage":                  71, // Storage Capacity
		"internal_memory_capacity": 36, // Internal Memory Capacity
		"card_slot_type":           17, // Card Slot Type
		// Cameras
		"rear_camera":                   122, // Rear Camera
		"camera_features":               16,  // Camera Features
		"quad_camera_setup":             58,  // Quad Camera Setup
		"camera_video_resolution":       40,  // Main Camera Video Resolution
		"optical_zoom":                  193, // Optical Zoom
		"front_camera":                  32,  // Front Camera
		"front_camera_video_resolution": 32,  // Front Camera Video Resolution
		// Battery & charging
		"battery":           11,  // Battery Capacity
		"battery_type":      12,  // Battery Type
		"fast_charging":     553, // fast_charging (lowercase)
		"charging_speed":    18,  // Charging Speed
		"wireless_charging": 82,  // Wireless Charging
		// Software
		"operating_system": 49, // Operating System
		// Audio
		"audio_quality":       9,  // Audio Quality
		"loudspeaker_quality": 38, // Loudspeaker Quality
		"audio_jack":          2,  // 3.5mm Audio Jack
		// Build & body
		"build_material":   14, // Build Material
		"weight":           80, // Weight
		"dimensions":       25, // Dimensions
		"water_resistance": 78, // Water Resistance
		// Network & connectivity
		"network_technology": 653, // Network Technology
		"2g_bands":           1,   // 2G Bands
		"3g_bands":           3,   // 3G Bands
		"4g_bands":           4,   // 4G Bands
		"5g_bands":           5,   // 5G Bands
		"5g_support":         5,   // 5G Support (maps to 5G Bands)
		"wifi_support":       81,  // WiFi
		"bluetooth_version":  13,  // Bluetooth Version
		"nfc_support":        46,  // NFC Support
		"usb_type":           76,  // USB Type
		// Positioning & sensors
		"positioning_system": 54, // Positioning System
		"sensors":            67, // Sensors
		"special_features":   69, // Special Features
		// SIM & compliance
		"sim_card_type": 68, // SIM Card Type
		"sar_rating":    63, // SAR (Specific Absorption Rate)
		"sar_rating_eu": 64, // SAR (Specific Absorption Rate) EU
		// Commercial info
		"available_colors":  10, // Available Colors
		"model_variants":    42, // Model Variants
		"announcement_date": 8,  // Announcement Date
		"device_status":     23, // Device Status
	}
	specs := map[string]string{
		"rear_camera":                   "108 MP + 50 MP + 20 MP",
		"special_features":              "Dual stereo speakers, Hi-Res Audio",
		"sim_card_type":                 "Dual Nano-SIM",
		"display_type":                  "AMOLED, 120Hz, HDR10+, 1200 nits (peak)",
		"build_material":                "Glass front (Gorilla Glass Victus), glass back, aluminum frame",
		"fast_charging":                 "66W wired",
		"audio_jack":                    "No",
		"device_status":                 "Available",
		"processor":                     "MediaTek Dimensity 9300",
		"gpu_type":                      "Mali-G720 Immortalis MC12",
		"4g_bands":                      "LTE Band 1, 3, 5, 7, 8, 20, 28, 38, 40, 41",
		"wifi_support":                  "Wi-Fi 802.11 a/b/g/n/ac/6e, dual-band, Wi-Fi Direct",
		"usb_type":                      "USB Type-C 3.1, OTG",
		"positioning_system":            "GPS, GLONASS, BDS, GALILEO",
		"loudspeaker_quality":           "Dual stereo speakers",
		"chipset":                       "MediaTek Dimensity 9300 (4 nm)",
		"card_slot_type":                "microSD card slot up to 1TB",
		"optical_zoom":                  "3x Optical Zoom",
		"wireless_charging":             "No",
		"internal_memory_capacity":      "UFS 4.0",
		"resolution":                    "1080 x 2412 pixels (~394 ppi density)",
		"bluetooth_version":             "5.3, A2DP, LE, aptX HD",
		"front_camera":                  "32 MP, f/2.4, (wide), HDR",
		"battery":                       "5,000 mAh",
		"model_variants":                "WX1U-001",
		"display_size":                  "6.7 inches",
		"refresh_rate":                  "120Hz",
		"dimensions":                    "161.8 x 73.9 x 8.4 mm (6.37 x 2.91 x 0.33 in)",
		"water_resistance":              "IP67 dust tight and water resistant (up to 1m for 30 min)",
		"2g_bands":                      "GSM 850 / 900 / 1800 / 1900 MHz",
		"nfc_support":                   "Yes",
		"operating_system":              "Android 14, Walton WaltonOS",
		"battery_type":                  "Li-Po (non-removable)",
		"weight":                        "195 g (6.88 oz)",
		"network_technology":            "GSM / HSPA / LTE / 5G",
		"5g_bands":                      "n1, n3, n5, n7, n8, n20, n28, n38, n40, n41, n77, n78",
		"sensors":                       "Fingerprint (under display, optical), accelerometer, gyro, proximity, compass",
		"audio_quality":                 "24-bit/192kHz audio, Hi-Res Audio",
		"announcement_date":             "December 2024",
		"ram":                           "12 GB",
		"screen_protection":             "Corning Gorilla Glass Victus",
		"camera_features":               "LED flash, HDR, panorama",
		"camera_video_resolution":       "4K@30/60fps, 1080p@30/60/120fps, gyro-EIS",
		"front_camera_video_resolution": "4K@30fps, 1080p@30/60fps",
		"charging_speed":                "66W wired (100% in 45 min)",
		"5g_support":                    "Yes",
		"available_colors":              "Midnight Black, Pearl White, Ocean Blue",
		"cpu_type":                      "Octa-core (1x3.25 GHz Cortex-X4 & 3x2.85 GHz Cortex-X4 & 4x2.0 GHz Cortex-A720)",
		"processor_speed":               "1x3.25 GHz + 3x2.85 GHz + 4x2.0 GHz",
		"storage":                       "256 GB",
		"3g_bands":                      "HSDPA 850 / 900 / 1900 / 2100 MHz",
		"sar_rating":                    "0.43 W/kg (head) 1.29 W/kg (body)",
		"sar_rating_eu":                 "0.32 W/kg (head) 1.50 W/kg (body)",
	}
	banglaTranslations := s.getBanglaTranslations()
	for key, value := range specs {
		specKeyID, exists := existingkeyMapping[key]
		if !exists {
			log.Printf("⚠️  Key not found in existingkeyMapping: '%!s(MISSING)' (used in product: %!s(MISSING))", key, productSlug)
			continue // Skip keys that don't exist in mapping
		}
		var existing models.SpecificationModel
		if err := db.Where("product_id = ? AND specification_key_id = ?", productID, specKeyID).First(&existing).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				sModel := &models.SpecificationModel{
					ProductID:          productID,
					SpecificationKeyID: specKeyID,
					Value:              value,
					Status:             1,
				}
				if err := db.Create(sModel).Error; err != nil {
					return err
				}
				// Create Bangla translation for the specification
				banglaValue, exists := banglaTranslations[value]
				if exists && banglaValue != "" {
					translation := &models.SpecificationTranslationModel{
						SpecificationID: sModel.ID,
						Locale:          "bn",
						Value:           banglaValue,
					}
					if err := db.Create(translation).Error; err != nil {
						return err
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
