package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederMobileXiaomi13Pro seeds specifications/options for product 'xiaomi-13-pro'
type SpecificationSeederMobileXiaomi13Pro struct {
	BaseSeeder
}

// NewSpecificationSeederMobileXiaomi13Pro creates a new seeder instance
func NewSpecificationSeederMobileXiaomi13Pro() *SpecificationSeederMobileXiaomi13Pro {
	return &SpecificationSeederMobileXiaomi13Pro{BaseSeeder: BaseSeeder{name: "Specifications for Xiaomi 13 Pro"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileXiaomi13Pro) getBanglaTranslations() map[string]string {
	specific := map[string]string{
		"120Hz": "১২০ হার্টজ",
		"Qualcomm SM8550-AB Snapdragon 8 Gen 2 (4 nm)": "কোয়ালকম এসএম৮৫৫০-এবি স্ন্যাপড্রাগন ৮ জেন ২ (৪ ন্যানোমিটার)",
		"Adreno 740": "অ্যাড্রেনো ৭৪০",
		"210 g":      "২১০ গ্রাম",
		"Ceramic White, Ceramic Black, Ceramic Flora Green, Mountain Blue": "সেরামিক হোয়াইট, সেরামিক ব্ল্যাক, সেরামিক ফ্লোরা গ্রিন, মাউন্টেন ব্লু",
		"Glass front, ceramic/polymer back, aluminum frame":                "গ্লাস ফ্রন্ট, সেরামিক/পলিমার ব্যাক, অ্যালুমিনিয়াম ফ্রেম",
		"IP68":                     "আইপি৬৮",
		"5G":                       "৫জি",
		"50 MP + 50 MP + 50 MP":    "৫০ এমপি + ৫০ এমপি + ৫০ এমপি",
		"Android 13, upgradable":   "অ্যান্ড্রয়েড ১৩, আপগ্রেডেবল",
		"Available":                "উপলব্ধ",
		"Snapdragon 8 Gen 2":       "স্ন্যাপড্রাগন ৮ জেন ২",
		"8 GB / 12 GB":             "৮ জিবি / ১২ জিবি",
		"128 GB / 256 GB / 512 GB": "১২৮ জিবি / ২৫৬ জিবি / ৫১২ জিবি",
		"LTPO AMOLED, 120Hz, Dolby Vision, HDR10+, 1900 nits": "এলটিপিও অ্যামোলেড, ১২০হার্টজ, ডলবি ভিশন, এইচডিআর১০+, ১৯০০ নিটস",
		"162.9 x 74.6 x 8.4 mm":                               "১৬২.৯ x ৭৪.৬ x ৮.৪ মিমি",
		"32 MP":                                               "৩২ এমপি",
		"4,820 mAh":                                           "৪,৮২০ এমএএইচ",
		"December 2022":                                       "ডিসেম্বর ২০২২",
		"6.73 inches":                                         "৬.৭৩ ইঞ্চি",
		"Octa-core":                                           "অক্টা-কোর",
		"1440 x 3200 pixels":                                  "১৪৪০ x ৩২০০ পিক্সেল",
		"Corning Gorilla Glass Victus":                        "কর্নিং গরিলা গ্লাস ভিক্টাস",
	}
	common := s.GetCommonBanglaTranslations()
	for k, v := range specific {
		common[k] = v
	}
	return common
}

// Seed inserts specification records for the product identified by slug 'xiaomi-13-pro'
func (s *SpecificationSeederMobileXiaomi13Pro) Seed(db *gorm.DB) error {
	productSlug := "xiaomi-13-pro"
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
		"rear_camera":                   "50 MP + 50 MP + 50 MP",
		"front_camera":                  "32 MP",
		"gpu_type":                      "Adreno 740",
		"ram":                           "8 GB / 12 GB",
		"storage":                       "128 GB / 256 GB / 512 GB",
		"display_type":                  "LTPO AMOLED, 120Hz, Dolby Vision, HDR10+, 1900 nits",
		"water_resistance":              "IP68",
		"processor":                     "Snapdragon 8 Gen 2",
		"cpu_type":                      "Octa-core",
		"weight":                        "210 g",
		"battery":                       "4,820 mAh",
		"available_colors":              "Ceramic White, Ceramic Black, Ceramic Flora Green, Mountain Blue",
		"display_size":                  "6.73 inches",
		"screen_protection":             "Corning Gorilla Glass Victus",
		"refresh_rate":                  "120Hz",
		"operating_system":              "Android 13, upgradable",
		"announcement_date":             "December 2022",
		"device_status":                 "Available",
		"chipset":                       "Qualcomm SM8550-AB Snapdragon 8 Gen 2 (4 nm)",
		"resolution":                    "1440 x 3200 pixels",
		"build_material":                "Glass front, ceramic/polymer back, aluminum frame",
		"dimensions":                    "162.9 x 74.6 x 8.4 mm",
		"fast_charging":                 "120W wired, 50W wireless, 10W reverse wireless",
		"charging_speed":                "120W wired (100% in 19 min), 50W wireless (100% in 40 min), 10W reverse wireless",
		"positioning_system":            "GPS (L1+L5), GLONASS (G1), BDS (B1I+B1c+B2a), GALILEO (E1+E5a), QZSS (L1+L5)",
		"sensors":                       "Fingerprint (under display, optical), accelerometer, gyro, proximity, compass, barometer, color spectrum",
		"special_features":              "Xiaomi HyperOS, Leica camera mode",
		"sim_card_type":                 "Dual Nano-SIM",
		"wifi_support":                  "Wi-Fi 802.11 a/b/g/n/ac/6e, tri-band, Wi-Fi Direct",
		"bluetooth_version":             "5.3, A2DP, LE, aptX HD, aptX Adaptive",
		"nfc_support":                   "Yes",
		"usb_type":                      "USB Type-C 3.2, DisplayPort 1.4",
		"audio_quality":                 "24-bit/192kHz audio",
		"loudspeaker_quality":           "Stereo speakers, Dolby Atmos",
		"audio_jack":                    "No",
		"camera_video_resolution":       "8K@24fps, 4K@30/60fps, 1080p@30/60/120/240fps, gyro-EIS, HDR10+, Dolby Vision",
		"front_camera_video_resolution": "4K@30/60fps, 1080p@30/60fps",
		"optical_zoom":                  "2x Optical Zoom",
		"quad_camera_setup":             "50MP (wide, f/1.9) + 50MP (telephoto 2x, f/2.0) + 50MP (ultrawide, f/2.2)",
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
