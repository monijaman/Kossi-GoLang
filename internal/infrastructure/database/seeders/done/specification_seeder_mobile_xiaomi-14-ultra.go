package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederMobileXiaomi14Ultra seeds specifications/options for product 'xiaomi-14-ultra'
type SpecificationSeederMobileXiaomi14Ultra struct {
	BaseSeeder
}

// NewSpecificationSeederMobileXiaomi14Ultra creates a new seeder instance
func NewSpecificationSeederMobileXiaomi14Ultra() *SpecificationSeederMobileXiaomi14Ultra {
	return &SpecificationSeederMobileXiaomi14Ultra{BaseSeeder: BaseSeeder{name: "Specifications for Xiaomi 14 Ultra"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileXiaomi14Ultra) getBanglaTranslations() map[string]string {
	specific := map[string]string{
		"5G":                            "৫জি",
		"50 MP + 50 MP + 50 MP + 50 MP": "৫০ এমপি + ৫০ এমপি + ৫০ এমপি + ৫০ এমপি",
		"32 MP":                         "৩২ এমপি",
		"Adreno 750":                    "অ্যাড্রেনো ৭৫০",
		"12 GB / 16 GB":                 "১২ জিবি / ১৬ জিবি",
		"256 GB / 512 GB / 1 TB":        "২৫৬ জিবি / ৫১২ জিবি / ১ টিবি",
		"LTPO AMOLED, 120Hz, Dolby Vision, HDR10+, 3000 nits": "এলটিপিও এএমওএলইডি, ১২০হার্টজ, ডলবি ভিশন, এইচডিআর১০+, ৩০০০ নিটস",
		"IP68":                              "আইপি৬৮",
		"Snapdragon 8 Gen 3":                "স্ন্যাপড্রাগন ৮ জেন ৩",
		"Octa-core":                         "অক্টা-কোর",
		"219 g":                             "২১৯ গ্রাম",
		"5,000 mAh":                         "৫,০০০ এমএএইচ",
		"Black, White, Blue, Titanium Gray": "ব্ল্যাক, হোয়াইট, ব্লু, টাইটানিয়াম গ্রে",
		"6.73 inches":                       "৬.৭৩ ইঞ্চি",
		"Xiaomi Shield Glass":               "শাওমি শিল্ড গ্লাস",
		"120Hz":                             "১২০হার্টজ",
		"Android 14, HyperOS":               "অ্যান্ড্রয়েড ১৪, হাইপারওএস",
		"February 2024":                     "ফেব্রুয়ারি ২০২৪",
		"Available":                         "উপলব্ধ",
		"Qualcomm SM8650-AB Snapdragon 8 Gen 3 (4 nm)":  "কোয়ালকম এসএম৮৬৫০-এবি স্ন্যাপড্রাগন ৮ জেন ৩ (৪ ন্যানোমিটার)",
		"1440 x 3200 pixels":                            "১৪৪০ x ৩২০০ পিক্সেল",
		"Glass front, eco leather back, titanium frame": "গ্লাস ফ্রন্ট, ইকো লেদার ব্যাক, টাইটানিয়াম ফ্রেম",
		"161.4 x 75.3 x 9.2 mm":                         "১৬১.৪ x ৭৫.৩ x ৯.২ মিমি",
		"GSM / HSPA / LTE / 5G":                         "জিএসএম / এইচএসপিএ / এলটিই / ৫জি",
		"90W wired, 80W wireless, 10W reverse wireless": "৯০ডব্লিউ ওয়্যার্ড, ৮০ডব্লিউ ওয়্যারলেস, ১০ডব্লিউ রিভার্স ওয়্যারলেস",
		"90W wired (100% in 35 min), 80W wireless (100% in 40 min), 10W reverse wireless":                          "৯০ডব্লিউ ওয়্যার্ড (১০০% ৩৫ মিনিটে), ৮০ডব্লিউ ওয়্যারলেস (১০০% ৪০ মিনিটে), ১০ডব্লিউ রিভার্স ওয়্যারলেস",
		"GPS (L1+L5), GLONASS (G1), BDS (B1I+B1c+B2a), GALILEO (E1+E5a), QZSS (L1+L5)":                             "জিপিএস (এল১+এল৫), গ্লোনাস (জি১), বিডিএস (বি১আই+বি১সি+বি২এ), গ্যালিলিও (ই১+ই৫এ), কিউজেডএসএস (এল১+এল৫)",
		"Fingerprint (under display, optical), accelerometer, gyro, proximity, compass, barometer, color spectrum": "ফিঙ্গারপ্রিন্ট (ডিসপ্লের নিচে, অপটিক্যাল), অ্যাকসেলেরোমিটার, জাইরো, প্রক্সিমিটি, কম্পাস, ব্যারোমিটার, কালার স্পেকট্রাম",
		"Xiaomi HyperOS, Leica camera mode":                                                                        "শাওমি হাইপারওএস, লাইকা ক্যামেরা মোড",
		"Dual Nano-SIM":                                                                                            "ডুয়াল ন্যানো-সিম",
		"Wi-Fi 802.11 a/b/g/n/ac/6e, tri-band, Wi-Fi Direct":                                                       "ওয়াই-ফাই ৮০২.১১ এ/বি/জি/এন/এসি/৬ই, ট্রাই-ব্যান্ড, ওয়াই-ফাই ডাইরেক্ট",
		"5.4, A2DP, LE, aptX HD, aptX Adaptive":                                                                    "৫.৪, এ২ডিপি, এলই, অ্যাপটিএক্স এইচডি, অ্যাপটিএক্স অ্যাডাপটিভ",
		"Yes":                                                                                                      "হ্যাঁ",
		"USB Type-C 3.2, DisplayPort 1.4":                                                                          "ইউএসবি টাইপ-সি ৩.২, ডিসপ্লেপোর্ট ১.৪",
		"24-bit/192kHz audio":                                                                                      "২৪-বিট/১৯২কেএইচজেড অডিও",
		"Stereo speakers, Dolby Atmos":                                                                             "স্টেরিও স্পিকার, ডলবি অ্যাটমস",
		"No":                                                                                                       "না",
		"8K@24fps, 4K@30/60fps, 1080p@30/60/120/240fps, gyro-EIS, HDR10+, Dolby Vision": "৮কে@২৪এফপিএস, ৪কে@৩০/৬০এফপিএস, ১০৮০পি@৩০/৬০/১২০/২৪০এফপিএস, জাইরো-ইআইএস, এইচডিআর১০+, ডলবি ভিশন",
		"4K@30/60fps, 1080p@30/60fps": "৪কে@৩০/৬০এফপিএস, ১০৮০পি@৩০/৬০এফপিএস",
		"3.2x + 5x Optical Zoom":      "৩.২x + ৫x অপটিক্যাল জুম",
		"50MP (wide, f/1.6) + 50MP (periscope telephoto 5x, f/2.5) + 50MP (telephoto 3.2x, f/2.0) + 50MP (ultrawide, f/1.8)": "৫০এমপি (ওয়াইড, এফ/১.৬) + ৫০এমপি (পেরিস্কোপ টেলিফটো ৫x, এফ/২.৫) + ৫০এমপি (টেলিফটো ৩.২x, এফ/২.০) + ৫০এমপি (আল্ট্রাওয়াইড, এফ/১.৮)",
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

// Seed inserts specification records for the product identified by slug 'xiaomi-14-ultra'
func (s *SpecificationSeederMobileXiaomi14Ultra) Seed(db *gorm.DB) error {
	productSlug := "xiaomi-14-ultra"
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
		"rear_camera":                   "50 MP + 50 MP + 50 MP + 50 MP",
		"front_camera":                  "32 MP",
		"gpu_type":                      "Adreno 750",
		"ram":                           "12 GB / 16 GB",
		"storage":                       "256 GB / 512 GB / 1 TB",
		"display_type":                  "LTPO AMOLED, 120Hz, Dolby Vision, HDR10+, 3000 nits",
		"water_resistance":              "IP68",
		"processor":                     "Snapdragon 8 Gen 3",
		"cpu_type":                      "Octa-core",
		"weight":                        "219 g",
		"battery":                       "5,000 mAh",
		"available_colors":              "Black, White, Blue, Titanium Gray",
		"display_size":                  "6.73 inches",
		"screen_protection":             "Xiaomi Shield Glass",
		"refresh_rate":                  "120Hz",
		"operating_system":              "Android 14, HyperOS",
		"announcement_date":             "February 2024",
		"device_status":                 "Available",
		"chipset":                       "Qualcomm SM8650-AB Snapdragon 8 Gen 3 (4 nm)",
		"resolution":                    "1440 x 3200 pixels",
		"build_material":                "Glass front, eco leather back, titanium frame",
		"dimensions":                    "161.4 x 75.3 x 9.2 mm",
		"fast_charging":                 "90W wired, 80W wireless, 10W reverse wireless",
		"charging_speed":                "90W wired (100% in 35 min), 80W wireless (100% in 40 min), 10W reverse wireless",
		"positioning_system":            "GPS (L1+L5), GLONASS (G1), BDS (B1I+B1c+B2a), GALILEO (E1+E5a), QZSS (L1+L5)",
		"sensors":                       "Fingerprint (under display, optical), accelerometer, gyro, proximity, compass, barometer, color spectrum",
		"special_features":              "Xiaomi HyperOS, Leica camera mode",
		"sim_card_type":                 "Dual Nano-SIM",
		"wifi_support":                  "Wi-Fi 802.11 a/b/g/n/ac/6e, tri-band, Wi-Fi Direct",
		"bluetooth_version":             "5.4, A2DP, LE, aptX HD, aptX Adaptive",
		"nfc_support":                   "Yes",
		"usb_type":                      "USB Type-C 3.2, DisplayPort 1.4",
		"audio_quality":                 "24-bit/192kHz audio",
		"loudspeaker_quality":           "Stereo speakers, Dolby Atmos",
		"audio_jack":                    "No",
		"camera_video_resolution":       "8K@24fps, 4K@30/60fps, 1080p@30/60/120/240fps, gyro-EIS, HDR10+, Dolby Vision",
		"front_camera_video_resolution": "4K@30/60fps, 1080p@30/60fps",
		"optical_zoom":                  "3.2x + 5x Optical Zoom",
		"quad_camera_setup":             "50MP (wide, f/1.6) + 50MP (periscope telephoto 5x, f/2.5) + 50MP (telephoto 3.2x, f/2.0) + 50MP (ultrawide, f/1.8)",
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
