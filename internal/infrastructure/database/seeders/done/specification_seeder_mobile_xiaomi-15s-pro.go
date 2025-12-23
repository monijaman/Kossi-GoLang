package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederMobileXiaomi15SPro seeds specifications/options for product 'xiaomi-15s-pro'
type SpecificationSeederMobileXiaomi15SPro struct {
	BaseSeeder
}

// NewSpecificationSeederMobileXiaomi15SPro creates a new seeder instance
func NewSpecificationSeederMobileXiaomi15SPro() *SpecificationSeederMobileXiaomi15SPro {
	return &SpecificationSeederMobileXiaomi15SPro{BaseSeeder: BaseSeeder{name: "Specifications for Xiaomi 15S Pro"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileXiaomi15SPro) getBanglaTranslations() map[string]string {
	specific := map[string]string{
		"Octa-core":                                "অক্টা-কোর",
		"GSM / HSPA / LTE / 5G":                    "জিএসএম / এইচএসপিএ / এলটিই / ৫জি",
		"120W wired, 50W wireless":                 "১২০ডব্লিউ ওয়্যার্ড, ৫০ডব্লিউ ওয়্যারলেস",
		"Black, White, Green":                      "ব্ল্যাক, হোয়াইট, গ্রিন",
		"6.73 inches":                              "৬.৭৩ ইঞ্চি",
		"Adreno 830":                               "অ্যাড্রেনো ৮৩০",
		"12 GB / 16 GB":                            "১২ জিবি / ১৬ জিবি",
		"256 GB / 512 GB / 1 TB":                   "২৫৬ জিবি / ৫১২ জিবি / ১ টিবি",
		"LTPO AMOLED, 120Hz, Dolby Vision, HDR10+": "এলটিপিও এএমওএলইডি, ১২০হার্টজ, ডলবি ভিশন, এইচডিআর১০+",
		"Glass front, glass back or eco leather back, aluminum frame": "গ্লাস ফ্রন্ট, গ্লাস ব্যাক বা ইকো লেদার ব্যাক, অ্যালুমিনিয়াম ফ্রেম",
		"5000 mAh":                    "৫০০০ এমএএইচ",
		"Android 15, HyperOS":         "অ্যান্ড্রয়েড ১৫, হাইপারওএস",
		"Snapdragon 8 Gen 4":          "স্ন্যাপড্রাগন ৮ জেন ৪",
		"1440 x 3200 pixels":          "১৪৪০ x ৩২০০ পিক্সেল",
		"IP68":                        "আইপি৬৮",
		"32 MP":                       "৩২ এমপি",
		"May 2025":                    "মে ২০২৫",
		"Upcoming":                    "আসন্ন",
		"215 g":                       "২১৫ গ্রাম",
		"Qualcomm Snapdragon 8 Gen 4": "কোয়ালকম স্ন্যাপড্রাগন ৮ জেন ৪",
		"120Hz":                       "১২০হার্টজ",
		"50 MP + 50 MP + 50 MP":       "৫০ এমপি + ৫০ এমপি + ৫০ এমপি",
		"Xiaomi Dragon Crystal Glass": "শাওমি ড্রাগন ক্রিস্টাল গ্লাস",
		"Li-Po":                       "লি-পো",
		"Yes":                         "হ্যাঁ",
		"24-bit/192kHz":               "২৪-বিট/১৯২কেএইচজেড",
		"Stereo":                      "স্টেরিও",
		"No":                          "না",
		"161.3 x 75.3 x 8.4 mm":       "১৬১.৩ x ৭৫.৩ x ৮.৪ মিমি",
		"GSM 850 / 900 / 1800 / 1900": "জিএসএম ৮৫০ / ৯০০ / ১৮০০ / ১৯০০",
		"HSDPA 850 / 900 / 1700(AWS) / 1900 / 2100": "এইচএসডিপিএ ৮৫০ / ৯০০ / ১৭০০(এডব্লিউএস) / ১৯০০ / ২১০০",
		"LTE":                          "এলটিই",
		"5G":                           "৫জি",
		"Wi-Fi 802.11 a/b/g/n/ac/6e/7": "ওয়াই-ফাই ৮০২.১১ এ/বি/জি/এন/এসি/৬ই/৭",
		"5.4":                          "৫.৪",
		"USB Type-C 3.2":               "ইউএসবি টাইপ-সি ৩.২",
		"GPS, GLONASS, BDS, GALILEO":   "জিপিএস, গ্লোনাস, বিডিএস, গ্যালিলিও",
		"Fingerprint, accelerometer, gyro, proximity, compass, barometer":                "ফিঙ্গারপ্রিন্ট, অ্যাক্সেলেরোমিটার, জাইরো, প্রক্সিমিটি, কম্পাস, ব্যারোমিটার",
		"AI image processing, Leica camera mode":                                         "এআই ইমেজ প্রসেসিং, লাইকা ক্যামেরা মোড",
		"Dual Nano-SIM":                                                                  "ডুয়াল ন্যানো-সিম",
		"8K@24fps, 4K@30/60fps, 1080p@30/60/120/240fps, gyro-EIS, HDR10+, Dolby Vision":  "৮কে@২৪এফপিএস, ৪কে@৩০/৬০এফপিএস, ১০৮০পি@৩০/৬০/১২০/২৪০এফপিএস, জাইরো-ইআইএস, এইচডিআর১০+, ডলবি ভিশন",
		"4K@30/60fps, 1080p@30/60fps":                                                    "৪কে@৩০/৬০এফপিএস, ১০৮০পি@৩০/৬০এফপিএস",
		"2x Optical Zoom":                                                                "২x অপটিক্যাল জুম",
		"50MP (wide, f/1.4) + 50MP (telephoto 2x, f/1.6) + 50MP (telephoto 3.2x, f/1.8)": "৫০এমপি (ওয়াইড, এফ/১.৪) + ৫০এমপি (টেলিফটো ২x, এফ/১.৬) + ৫০এমপি (টেলিফটো ৩.২x, এফ/১.৮)",
		"UFS 2.2":                           "ইউএফএস ২.২",
		"0.99 W/kg (head) 0.99 W/kg (body)": "০.৯৯ ডব্লিউ/কেজি (হেড) ০.৯৯ ডব্লিউ/কেজি (বডি)",
		"Xiaomi 15S Pro":                    "শাওমি ১৫এস প্রো",
		"Octa-core (1x3.4 GHz Cortex-X5 & 4x3.2 GHz Cortex-X4 & 2x2.8 GHz Cortex-A720 & 1x2.1 GHz Cortex-A520)": "অক্টা-কোর (১x৩.৪ গিগাহার্টজ কর্টেক্স-এক্স৫ & ৪x৩.২ গিগাহার্টজ কর্টেক্স-এক্স৪ & ২x২.৮ গিগাহার্টজ কর্টেক্স-এ৭২০ & ১x২.১ গিগাহার্টজ কর্টেক্স-এ৫২০)",
	}
	common := s.GetCommonBanglaTranslations()
	for k, v := range specific {
		common[k] = v
	}
	return common
}

// Seed inserts specification records for the product identified by slug 'xiaomi-15s-pro'
func (s *SpecificationSeederMobileXiaomi15SPro) Seed(db *gorm.DB) error {
	productSlug := "xiaomi-15s-pro"
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
		"cpu_type":                      "Octa-core",
		"network_technology":            "GSM / HSPA / LTE / 5G",
		"fast_charging":                 "120W wired, 50W wireless",
		"available_colors":              "Black, White, Green",
		"display_size":                  "6.73 inches",
		"gpu_type":                      "Adreno 830",
		"ram":                           "12 GB / 16 GB",
		"storage":                       "256 GB / 512 GB / 1 TB",
		"display_type":                  "LTPO AMOLED, 120Hz, Dolby Vision, HDR10+",
		"build_material":                "Glass front, glass back or eco leather back, aluminum frame",
		"battery":                       "5000 mAh",
		"operating_system":              "Android 15, HyperOS",
		"processor":                     "Snapdragon 8 Gen 4",
		"resolution":                    "1440 x 3200 pixels",
		"water_resistance":              "IP68",
		"front_camera":                  "32 MP",
		"announcement_date":             "May 2025",
		"device_status":                 "Upcoming",
		"weight":                        "215 g",
		"chipset":                       "Qualcomm Snapdragon 8 Gen 4",
		"refresh_rate":                  "120Hz",
		"rear_camera":                   "50 MP + 50 MP + 50 MP",
		"screen_protection":             "Xiaomi Dragon Crystal Glass",
		"battery_type":                  "Li-Po",
		"charging_speed":                "120W wired, 50W wireless",
		"wireless_charging":             "Yes",
		"audio_quality":                 "24-bit/192kHz",
		"loudspeaker_quality":           "Stereo",
		"audio_jack":                    "No",
		"dimensions":                    "161.3 x 75.3 x 8.4 mm",
		"2g_bands":                      "GSM 850 / 900 / 1800 / 1900",
		"3g_bands":                      "HSDPA 850 / 900 / 1700(AWS) / 1900 / 2100",
		"4g_bands":                      "LTE",
		"5g_bands":                      "5G",
		"wifi_support":                  "Wi-Fi 802.11 a/b/g/n/ac/6e/7",
		"bluetooth_version":             "5.4",
		"nfc_support":                   "Yes",
		"usb_type":                      "USB Type-C 3.2",
		"positioning_system":            "GPS, GLONASS, BDS, GALILEO",
		"sensors":                       "Fingerprint, accelerometer, gyro, proximity, compass, barometer",
		"special_features":              "AI image processing, Leica camera mode",
		"sim_card_type":                 "Dual Nano-SIM",
		"camera_video_resolution":       "8K@24fps, 4K@30/60fps, 1080p@30/60/120/240fps, gyro-EIS, HDR10+, Dolby Vision",
		"front_camera_video_resolution": "4K@30/60fps, 1080p@30/60fps",
		"optical_zoom":                  "2x Optical Zoom",
		"quad_camera_setup":             "50MP (wide, f/1.4) + 50MP (telephoto 2x, f/1.6) + 50MP (telephoto 3.2x, f/1.8)",
		"internal_memory_capacity":      "UFS 2.2",
		"card_slot_type":                "No",
		"sar_rating":                    "0.99 W/kg (head) 0.99 W/kg (body)",
		"sar_rating_eu":                 "0.99 W/kg (head) 0.99 W/kg (body)",
		"model_variants":                "Xiaomi 15S Pro",
		"processor_speed":               "Octa-core (1x3.4 GHz Cortex-X5 & 4x3.2 GHz Cortex-X4 & 2x2.8 GHz Cortex-A720 & 1x2.1 GHz Cortex-A520)",
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
