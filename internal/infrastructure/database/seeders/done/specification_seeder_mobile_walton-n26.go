package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederMobileWaltonN26 seeds specifications/options for product 'walton-n26'
type SpecificationSeederMobileWaltonN26 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileWaltonN26 creates a new seeder instance
func NewSpecificationSeederMobileWaltonN26() *SpecificationSeederMobileWaltonN26 {
	return &SpecificationSeederMobileWaltonN26{BaseSeeder: BaseSeeder{name: "Specifications for Walton N26"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileWaltonN26) getBanglaTranslations() map[string]string {
	translations := map[string]string{
		"4 GB":                       "4 জিবি",
		"IPS LCD":                    "IPS LCD",
		"Glass front":                "গ্লাস ফ্রন্ট",
		"Plastic body":               "প্লাস্টিক বডি",
		"Blue, Black":                "নীল, কালো",
		"Available":                  "উপলব্ধ",
		"6.5 inches":                 "6.5 ইঞ্চি",
		"Helio G36":                  "হেলিও G36",
		"Octa-core":                  "অক্টা-কোর",
		"720 x 1600 pixels":          "720x1600",
		"164 x 75.8 x 8.9 mm":        "164x75.8x8.9 মিমি",
		"4G":                         "4G",
		"13 MP + 2 MP":               "13+2 মেগা",
		"5,000 mAh":                  "5000 mAh",
		"Mediatek Helio G36 (12 nm)": "মিডিয়াটেক হেলিও G36",
		"PowerVR GE8320":             "পাওয়ারভিআর GE8320",
		"192 g":                      "192 গ্রাম",
		"5 MP":                       "5 মেগা",
		"Android 13 Go Edition":      "অ্যান্ড্রয়েড 13 গো",
		"August 2023":                "আগস্ট 2023",
		"64 GB":                      "64 জিবি",
		"60Hz":                       "60Hz",
		"No":                         "না",
		"Fingerprint (rear-mounted), accelerometer, proximity": "ফিঙ্গারপ্রিন্ট (রিয়ার), অ্যাক্সেলেরোমিটার, প্রক্সিমিটি",
		"GPS, GLONASS":                    "GPS, গ্লোনাস",
		"5.0":                             "5.0",
		"Wi-Fi 802.11 b/g/n":              "ওয়াই-ফাই 4",
		"USB-C 2.0":                       "USB-C",
		"Dual SIM (Nano-SIM)":             "ডুয়াল সিম",
		"Yes, 3.5 mm":                     "হ্যাঁ, 3.5 mm",
		"Single":                          "সিঙ্গেল",
		"Mono":                            "মনো",
		"Li-Ion (non-removable)":          "লি-আয়ন",
		"10 W wired":                      "10W",
		"microSDXC":                       "মাইক্রো এসডিএক্সসি",
		"2.2 GHz":                         "2.2 GHz",
		"LED flash":                       "LED ফ্ল্যাশ",
		"1080p @ 30fps":                   "1080p 30fps",
		"None":                            "না",
		"720p @ 30fps":                    "720p 30fps",
		"0.5 W/kg (head) 1.0 W/kg (body)": "0.5 W/kg (হেড) 1.0 W/kg (বডি)",
		"0.4 W/kg (head) 1.0 W/kg (body)": "0.4 W/kg (হেড) 1.0 W/kg (বডি)",
		"N26":                             "N26",
		"Face unlock":                     "ফেস আনলক",
	}
	for k, v := range s.GetCommonBanglaTranslations() {
		translations[k] = v
	}
	return translations
}

// Seed inserts specification records for the product identified by slug 'walton-n26'
func (s *SpecificationSeederMobileWaltonN26) Seed(db *gorm.DB) error {
	productSlug := "walton-n26"
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
		"ram":                           "4 GB",
		"display_type":                  "IPS LCD",
		"screen_protection":             "Glass front",
		"build_material":                "Plastic body",
		"available_colors":              "Blue, Black",
		"device_status":                 "Available",
		"display_size":                  "6.5 inches",
		"processor":                     "Helio G36",
		"cpu_type":                      "Octa-core",
		"resolution":                    "720 x 1600 pixels",
		"dimensions":                    "164 x 75.8 x 8.9 mm",
		"network_technology":            "4G",
		"rear_camera":                   "13 MP + 2 MP",
		"battery":                       "5,000 mAh",
		"chipset":                       "Mediatek Helio G36 (12 nm)",
		"gpu_type":                      "PowerVR GE8320",
		"weight":                        "192 g",
		"front_camera":                  "5 MP",
		"operating_system":              "Android 13 Go Edition",
		"announcement_date":             "August 2023",
		"storage":                       "64 GB",
		"refresh_rate":                  "60Hz",
		"water_resistance":              "No",
		"sensors":                       "Fingerprint (rear-mounted), accelerometer, proximity",
		"positioning_system":            "GPS, GLONASS",
		"bluetooth_version":             "5.0",
		"wifi_support":                  "Wi-Fi 802.11 b/g/n",
		"usb_type":                      "USB-C 2.0",
		"sim_card_type":                 "Dual SIM (Nano-SIM)",
		"audio_jack":                    "Yes, 3.5 mm",
		"loudspeaker_quality":           "Single",
		"audio_quality":                 "Mono",
		"battery_type":                  "Li-Ion (non-removable)",
		"charging_speed":                "10 W wired",
		"wireless_charging":             "No",
		"card_slot_type":                "microSDXC",
		"internal_memory_capacity":      "64 GB",
		"processor_speed":               "2.2 GHz",
		"camera_features":               "LED flash",
		"quad_camera_setup":             "13 MP + 2 MP",
		"camera_video_resolution":       "1080p @ 30fps",
		"optical_zoom":                  "None",
		"front_camera_video_resolution": "720p @ 30fps",
		"nfc_support":                   "No",
		"sar_rating":                    "0.5 W/kg (head) 1.0 W/kg (body)",
		"sar_rating_eu":                 "0.4 W/kg (head) 1.0 W/kg (body)",
		"model_variants":                "N26",
		"special_features":              "Face unlock",
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
