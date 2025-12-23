package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederMobileWaltonZENX2 seeds specifications/options for product 'walton-zenx-2'
type SpecificationSeederMobileWaltonZENX2 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileWaltonZENX2 creates a new seeder instance
func NewSpecificationSeederMobileWaltonZENX2() *SpecificationSeederMobileWaltonZENX2 {
	return &SpecificationSeederMobileWaltonZENX2{BaseSeeder: BaseSeeder{name: "Specifications for Walton ZENX 2"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileWaltonZENX2) getBanglaTranslations() map[string]string {
	specific := map[string]string{
		"64 MP + 8 MP + 2 MP":               "৬৪ এমপি + ৮ এমপি + ২ এমপি",
		"5,000 mAh":                         "৫,০০০ এমএএইচ",
		"November 2023":                     "নভেম্বর ২০২৩",
		"6.8 inches":                        "৬.৮ ইঞ্চি",
		"IPS LCD, 120Hz":                    "আইপিএস এলসিডি, ১২০হার্টজ",
		"16 MP":                             "১৬ এমপি",
		"Available":                         "উপলব্ধ",
		"Qualcomm Snapdragon 695 5G (6 nm)": "কোয়ালকম স্ন্যাপড্রাগন ৬৯৫ ৫জি (৬ ন্যানোমিটার)",
		"8 GB":                              "৮ জিবি",
		"1080 x 2460 pixels":                "১০৮০ x ২৪৬০ পিক্সেল",
		"Corning Gorilla Glass 3":           "কর্নিং গরিলা গ্লাস ৩",
		"Glass front, plastic back":         "গ্লাস ফ্রন্ট, প্লাস্টিক ব্যাক",
		"No":                                "না",
		"128 GB":                            "১২৮ জিবি",
		"120Hz":                             "১২০হার্টজ",
		"168.5 x 76.5 x 9.2 mm":             "১৬৮.৫ x ৭৬.৫ x ৯.২ মিমি",
		"Android 13":                        "অ্যান্ড্রয়েড ১৩",
		"Silver, Black":                     "সিলভার, ব্ল্যাক",
		"Snapdragon 695":                    "স্ন্যাপড্রাগন ৬৯৫",
		"Octa-core":                         "অক্টা-কোর",
		"Adreno 619":                        "অ্যাড্রেনো ৬১৯",
		"205 g":                             "২০৫ গ্রাম",
		"5G":                                "৫জি",
		"GSM / HSPA / LTE / 5G":             "জিএসএম / এইচএসপিএ / এলটিই / ৫জি",
		"18W wired":                         "১৮ডব্লিউ ওয়্যার্ড",
		"GPS, GLONASS, BDS, GALILEO":        "জিপিএস, গ্লোনাস, বিডিএস, গ্যালিলিও",
		"Fingerprint (side-mounted), accelerometer, proximity, compass": "ফিঙ্গারপ্রিন্ট (সাইড-মাউন্টেড), অ্যাকসেলেরোমিটার, প্রক্সিমিটি, কম্পাস",
		"Walton Prism UI": "ওয়ালটন প্রিজম ইউআই",
		"Dual Nano-SIM":   "ডুয়াল ন্যানো-সিম",
		"Wi-Fi 802.11 a/b/g/n/ac, dual-band, Wi-Fi Direct": "ওয়াই-ফাই ৮০২.১১ এ/বি/জি/এন/এসি, ডুয়াল-ব্যান্ড, ওয়াই-ফাই ডাইরেক্ট",
		"5.1, A2DP, LE":                   "৫.১, এ২ডিপি, এলই",
		"USB Type-C 2.0":                  "ইউএসবি টাইপ-সি ২.০",
		"24-bit/192kHz audio":             "২৪-বিট/১৯২কেএইচজেড অডিও",
		"Mono speaker":                    "মোনো স্পিকার",
		"Yes":                             "হ্যাঁ",
		"1080p@30fps":                     "১০৮০পি@৩০এফপিএস",
		"Li-Po (non-removable)":           "লি-পো (নন-রিমুভেবল)",
		"UFS 2.2":                         "ইউএফএস ২.২",
		"microSDXC (dedicated slot)":      "মাইক্রোএসডিএক্সসি (ডেডিকেটেড স্লট)",
		"GSM 850 / 900 / 1800 / 1900 MHz": "জিএসএম ৮৫০ / ৯০০ / ১৮০০ / ১৯০০ মেগাহার্টজ",
		"HSDPA 850 / 900 / 2100 MHz":      "এইচএসডিপিএ ৮৫০ / ৯০০ / ২১০০ মেগাহার্টজ",
		"LTE bands 1, 3, 5, 7, 8, 20, 28, 38, 40, 41":                "এলটিই ব্যান্ডস ১, ৩, ৫, ৭, ৮, ২০, ২৮, ৩৮, ৪০, ৪১",
		"n1, n3, n5, n7, n8, n20, n28, n38, n40, n41, n66, n77, n78": "এন১, এন৩, এন৫, এন৭, এন৮, এন২০, এন২৮, এন৩৮, এন৪০, এন৪১, এন৬৬, এন৭৭, এন৭৮",
	}
	common := s.GetCommonBanglaTranslations()
	for k, v := range specific {
		common[k] = v
	}
	return common
}

// Seed inserts specification records for the product identified by slug 'walton-zenx-2'
func (s *SpecificationSeederMobileWaltonZENX2) Seed(db *gorm.DB) error {
	productSlug := "walton-zenx-2"
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
		"rear_camera":                   "64 MP + 8 MP + 2 MP",
		"front_camera":                  "16 MP",
		"gpu_type":                      "Adreno 619",
		"ram":                           "8 GB",
		"storage":                       "128 GB",
		"display_type":                  "IPS LCD, 120Hz",
		"water_resistance":              "No",
		"processor":                     "Snapdragon 695",
		"cpu_type":                      "Octa-core",
		"weight":                        "205 g",
		"battery":                       "5,000 mAh",
		"available_colors":              "Silver, Black",
		"display_size":                  "6.8 inches",
		"screen_protection":             "Corning Gorilla Glass 3",
		"refresh_rate":                  "120Hz",
		"operating_system":              "Android 13",
		"announcement_date":             "November 2023",
		"device_status":                 "Available",
		"chipset":                       "Qualcomm Snapdragon 695 5G (6 nm)",
		"resolution":                    "1080 x 2460 pixels",
		"build_material":                "Glass front, plastic back",
		"dimensions":                    "168.5 x 76.5 x 9.2 mm",
		"fast_charging":                 "18W wired",
		"charging_speed":                "18W wired",
		"positioning_system":            "GPS, GLONASS, BDS, GALILEO",
		"sensors":                       "Fingerprint (side-mounted), accelerometer, proximity, compass",
		"special_features":              "Walton Prism UI",
		"sim_card_type":                 "Dual Nano-SIM",
		"wifi_support":                  "Wi-Fi 802.11 a/b/g/n/ac, dual-band, Wi-Fi Direct",
		"bluetooth_version":             "5.1, A2DP, LE",
		"nfc_support":                   "No",
		"usb_type":                      "USB Type-C 2.0",
		"audio_quality":                 "24-bit/192kHz audio",
		"loudspeaker_quality":           "Mono speaker",
		"audio_jack":                    "Yes",
		"camera_video_resolution":       "1080p@30fps",
		"front_camera_video_resolution": "1080p@30fps",
		"battery_type":                  "Li-Po (non-removable)",
		"internal_memory_capacity":      "UFS 2.2",
		"card_slot_type":                "microSDXC (dedicated slot)",
		"5g_support":                    "Yes",
		"2g_bands":                      "GSM 850 / 900 / 1800 / 1900 MHz",
		"3g_bands":                      "HSDPA 850 / 900 / 2100 MHz",
		"4g_bands":                      "LTE bands 1, 3, 5, 7, 8, 20, 28, 38, 40, 41",
		"5g_bands":                      "n1, n3, n5, n7, n8, n20, n28, n38, n40, n41, n66, n77, n78",
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
