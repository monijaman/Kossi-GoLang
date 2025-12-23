package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederMobileXiaomi14TPro seeds specifications/options for product 'xiaomi-14t-pro'
type SpecificationSeederMobileXiaomi14TPro struct {
	BaseSeeder
}

// NewSpecificationSeederMobileXiaomi14TPro creates a new seeder instance
func NewSpecificationSeederMobileXiaomi14TPro() *SpecificationSeederMobileXiaomi14TPro {
	return &SpecificationSeederMobileXiaomi14TPro{BaseSeeder: BaseSeeder{name: "Specifications for Xiaomi 14T Pro"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileXiaomi14TPro) getBanglaTranslations() map[string]string {
	specific := map[string]string{
		"Android 14, HyperOS":                     "অ্যান্ড্রয়েড ১৪, হাইপারওএস",
		"Mediatek Dimensity 9300+ (4 nm)":         "মিডিয়াটেক ডাইমেনসিটি ৯৩০০+ (৪ ন্যানোমিটার)",
		"Octa-core":                               "অক্টা-কোর",
		"12 GB / 16 GB":                           "১২ জিবি / ১৬ জিবি",
		"1220 x 2712 pixels":                      "১২২০ x ২৭১২ পিক্সেল",
		"Glass front, aluminum frame, glass back": "গ্লাস ফ্রন্ট, অ্যালুমিনিয়াম ফ্রেম, গ্লাস ব্যাক",
		"GSM / HSPA / LTE / 5G":                   "জিএসএম / এইচএসপিএ / এলটিই / ৫জি",
		"Titan Gray, Titan Blue, Titan Black":     "টাইটান গ্রে, টাইটান ব্লু, টাইটান ব্ল্যাক",
		"6.67 inches":                             "৬.৬৭ ইঞ্চি",
		"Immortalis-G720 MC12":                    "ইমর্টালিস-জি৭২০ এমসি১২",
		"AMOLED, 144Hz, Dolby Vision, HDR10+":     "এমোলেড, ১৪৪হার্টজ, ডলবি ভিশন, এইচডিআর১০+",
		"September 2024":                          "সেপ্টেম্বর ২০২৪",
		"Available":                               "উপলব্ধ",
		"256 GB / 512 GB / 1 TB":                  "২৫৬ জিবি / ৫১২ জিবি / ১ টিবি",
		"50 MP + 50 MP + 12 MP":                   "৫০ এমপি + ৫০ এমপি + ১২ এমপি",
		"5000 mAh":                                "৫০০০ এমএএইচ",
		"Dimensity 9300+":                         "ডাইমেনসিটি ৯৩০০+",
		"144Hz":                                   "১৪৪ হার্টজ",
		"209 g":                                   "২০৯ গ্রাম",
		"IP68":                                    "আইপি৬৮",
		"32 MP":                                   "৩২ এমপি",
		"120W wired, 50W wireless":                "১২০ডব্লিউ ওয়্যার্ড, ৫০ডব্লিউ ওয়্যারলেস",
		"160.4 x 75.1 x 8.4 mm":                   "১৬০.৪ x ৭৫.১ x ৮.৪ মিমি",
		"120W wired (100% in 19 min), 50W wireless":                                                                "১২০ডব্লিউ ওয়্যার্ড (১০০% ১৯ মিনিটে), ৫০ডব্লিউ ওয়্যারলেস",
		"GPS (L1+L5), GLONASS (G1), BDS (B1I+B1c+B2a), GALILEO (E1+E5a), QZSS (L1+L5)":                             "জিপিএস (এল১+এল৫), গ্লোনাস (জি১), বিডিএস (বি১আই+বি১সি+বি২এ), গ্যালিলিও (ই১+ই৫এ), কিউজেডএসএস (এল১+এল৫)",
		"Fingerprint (under display, optical), accelerometer, gyro, proximity, compass, barometer, color spectrum": "ফিঙ্গারপ্রিন্ট (ডিসপ্লের নিচে, অপটিক্যাল), অ্যাকসেলেরোমিটার, জাইরো, প্রক্সিমিটি, কম্পাস, ব্যারোমিটার, কালার স্পেকট্রাম",
		"Xiaomi HyperOS, Leica camera mode":                                                                        "শাওমি হাইপারওএস, লাইকা ক্যামেরা মোড",
		"Dual Nano-SIM":                                                                                            "ডুয়াল ন্যানো-সিম",
		"Wi-Fi 802.11 a/b/g/n/ac/6e/7, tri-band, Wi-Fi Direct":                                                     "ওয়াই-ফাই ৮০২.১১ এ/বি/জি/এন/এসি/৬ই/৭, ট্রাই-ব্যান্ড, ওয়াই-ফাই ডাইরেক্ট",
		"5.4, A2DP, LE, aptX HD, aptX Adaptive":                                                                    "৫.৪, এ২ডিপি, এলই, অ্যাপটিএক্স এইচডি, অ্যাপটিএক্স অ্যাডাপটিভ",
		"Yes":                                                                                                      "হ্যাঁ",
		"USB Type-C 3.2, DisplayPort 1.4":                                                                          "ইউএসবি টাইপ-সি ৩.২, ডিসপ্লেপোর্ট ১.৪",
		"24-bit/192kHz audio":                                                                                      "২৪-বিট/১৯২কেএইচজেড অডিও",
		"Stereo speakers, Dolby Atmos":                                                                             "স্টেরিও স্পিকার, ডলবি অ্যাটমস",
		"No":                                                                                                       "না",
		"8K@24fps, 4K@30/60fps, 1080p@30/60/120/240fps, gyro-EIS, HDR10+, Dolby Vision": "৮কে@২৪এফপিএস, ৪কে@৩০/৬০এফপিএস, ১০৮০পি@৩০/৬০/১২০/২৪০এফপিএস, জাইরো-ইআইএস, এইচডিআর১০+, ডলবি ভিশন",
		"4K@30/60fps, 1080p@30/60fps": "৪কে@৩০/৬০এফপিএস, ১০৮০পি@৩০/৬০এফপিএস",
		"2x Optical Zoom":             "২x অপটিক্যাল জুম",
		"50MP (wide, f/1.9) + 50MP (telephoto 2x, f/2.0) + 12MP (ultrawide, f/2.2)": "৫০এমপি (ওয়াইড, এফ/১.৯) + ৫০এমপি (টেলিফটো ২x, এফ/২.০) + ১২এমপি (আল্ট্রাওয়াইড, এফ/২.২)",
		"Li-Po (non-removable)":                         "লি-পো (নন-রিমুভেবল)",
		"UFS 4.0":                                       "ইউএফএস ৪.০",
		"No microSD card slot":                          "কোন মাইক্রোএসডি কার্ড স্লট নেই",
		"GSM 850 / 900 / 1800 / 1900 MHz":               "জিএসএম ৮৫০ / ৯০০ / ১৮০০ / ১৯০০ মেগাহার্টজ",
		"HSDPA 850 / 900 / 1700(AWS) / 1900 / 2100 MHz": "এইচএসডিপিএ ৮৫০ / ৯০০ / ১৭০০(এডব্লিউএস) / ১৯০০ / ২১০০ মেগাহার্টজ",
		"LTE bands 1, 2, 3, 4, 5, 7, 8, 12, 13, 17, 18, 19, 20, 26, 28, 32, 38, 39, 40, 41, 42, 48, 66": "এলটিই ব্যান্ডস ১, ২, ৩, ৪, ৫, ৭, ৮, ১২, ১৩, ১৭, ১৮, ১৯, ২০, ২৬, ২৮, ৩২, ৩৮, ৩৯, ৪০, ৪১, ৪২, ৪৮, ৬৬",
		"n1, n3, n5, n7, n8, n20, n28, n38, n40, n41, n77, n78":                                         "এন১, এন৩, এন৫, এন৭, এন৮, এন২০, এন২৮, এন৩৮, এন৪০, এন৪১, এন৭৭, এন৭৮",
	}
	common := s.GetCommonBanglaTranslations()
	for k, v := range specific {
		common[k] = v
	}
	return common
}

// Seed inserts specification records for the product identified by slug 'xiaomi-14t-pro'
func (s *SpecificationSeederMobileXiaomi14TPro) Seed(db *gorm.DB) error {
	productSlug := "xiaomi-14t-pro"
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
		"network_technology":            "GSM / HSPA / LTE / 5G",
		"rear_camera":                   "50 MP + 50 MP + 12 MP",
		"front_camera":                  "32 MP",
		"gpu_type":                      "Immortalis-G720 MC12",
		"ram":                           "12 GB / 16 GB",
		"storage":                       "256 GB / 512 GB / 1 TB",
		"display_type":                  "AMOLED, 144Hz, Dolby Vision, HDR10+",
		"water_resistance":              "IP68",
		"processor":                     "Dimensity 9300+",
		"cpu_type":                      "Octa-core",
		"weight":                        "209 g",
		"battery":                       "5000 mAh",
		"available_colors":              "Titan Gray, Titan Blue, Titan Black",
		"display_size":                  "6.67 inches",
		"refresh_rate":                  "144Hz",
		"operating_system":              "Android 14, HyperOS",
		"announcement_date":             "September 2024",
		"device_status":                 "Available",
		"chipset":                       "Mediatek Dimensity 9300+ (4 nm)",
		"resolution":                    "1220 x 2712 pixels",
		"build_material":                "Glass front, aluminum frame, glass back",
		"dimensions":                    "160.4 x 75.1 x 8.4 mm",
		"fast_charging":                 "120W wired, 50W wireless",
		"charging_speed":                "120W wired (100% in 19 min), 50W wireless",
		"positioning_system":            "GPS (L1+L5), GLONASS (G1), BDS (B1I+B1c+B2a), GALILEO (E1+E5a), QZSS (L1+L5)",
		"sensors":                       "Fingerprint (under display, optical), accelerometer, gyro, proximity, compass, barometer, color spectrum",
		"special_features":              "Xiaomi HyperOS, Leica camera mode",
		"sim_card_type":                 "Dual Nano-SIM",
		"wifi_support":                  "Wi-Fi 802.11 a/b/g/n/ac/6e/7, tri-band, Wi-Fi Direct",
		"bluetooth_version":             "5.4, A2DP, LE, aptX HD, aptX Adaptive",
		"nfc_support":                   "Yes",
		"usb_type":                      "USB Type-C 3.2, DisplayPort 1.4",
		"audio_quality":                 "24-bit/192kHz audio",
		"loudspeaker_quality":           "Stereo speakers, Dolby Atmos",
		"audio_jack":                    "No",
		"camera_video_resolution":       "8K@24fps, 4K@30/60fps, 1080p@30/60/120/240fps, gyro-EIS, HDR10+, Dolby Vision",
		"front_camera_video_resolution": "4K@30/60fps, 1080p@30/60fps",
		"optical_zoom":                  "2x Optical Zoom",
		"quad_camera_setup":             "50MP (wide, f/1.9) + 50MP (telephoto 2x, f/2.0) + 12MP (ultrawide, f/2.2)",
		"battery_type":                  "Li-Po (non-removable)",
		"internal_memory_capacity":      "UFS 4.0",
		"card_slot_type":                "No microSD card slot",
		"5g_support":                    "Yes",
		"2g_bands":                      "GSM 850 / 900 / 1800 / 1900 MHz",
		"3g_bands":                      "HSDPA 850 / 900 / 1700(AWS) / 1900 / 2100 MHz",
		"4g_bands":                      "LTE bands 1, 2, 3, 4, 5, 7, 8, 12, 13, 17, 18, 19, 20, 26, 28, 32, 38, 39, 40, 41, 42, 48, 66",
		"5g_bands":                      "n1, n3, n5, n7, n8, n20, n28, n38, n40, n41, n77, n78",
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
