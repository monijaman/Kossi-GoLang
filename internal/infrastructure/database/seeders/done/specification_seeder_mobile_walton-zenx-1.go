package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederMobileWaltonZENX1 seeds specifications/options for product 'walton-zenx-1'
type SpecificationSeederMobileWaltonZENX1 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileWaltonZENX1 creates a new seeder instance
func NewSpecificationSeederMobileWaltonZENX1() *SpecificationSeederMobileWaltonZENX1 {
	return &SpecificationSeederMobileWaltonZENX1{BaseSeeder: BaseSeeder{name: "Specifications for Walton ZENX 1"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileWaltonZENX1) getBanglaTranslations() map[string]string {
	translations := map[string]string{
		"168.5 x 76.5 x 9.2 mm":             "168.5x76.5x9.2 মিমি",
		"Snapdragon 695":                    "স্ন্যাপড্রাগন 695",
		"Qualcomm Snapdragon 695 5G (6 nm)": "কোয়ালকম স্ন্যাপড্রাগন 695 5G",
		"8 GB":                              "8 জিবি",
		"128 GB":                            "128 জিবি",
		"1080 x 2460 pixels":                "1080x2460",
		"Glass front, plastic back":         "গ্লাস ফ্রন্ট, প্লাস্টিক ব্যাক",
		"No":                                "না",
		"Silver, Black":                     "সিলভার, ব্ল্যাক",
		"205 g":                             "205 গ্রাম",
		"5G":                                "5G",
		"5,000 mAh":                         "5000 mAh",
		"May 2023":                          "মে 2023",
		"6.8 inches":                        "6.8 ইঞ্চি",
		"120Hz":                             "120Hz",
		"64 MP + 8 MP + 2 MP":               "64+8+2 মেগা",
		"16 MP":                             "16 মেগা",
		"Android 13":                        "অ্যান্ড্রয়েড 13",
		"Available":                         "উপলব্ধ",
		"Octa-core":                         "অক্টা-কোর",
		"Adreno 619":                        "অ্যাড্রেনো 619",
		"IPS LCD, 120Hz":                    "IPS LCD 120Hz",
		"Corning Gorilla Glass 3":           "কর্নিং গরিলা গ্লাস 3",
		"Fingerprint (side-mounted), accelerometer, proximity, compass": "ফিঙ্গারপ্রিন্ট (সাইড), অ্যাক্সেলেরোমিটার, প্রক্সিমিটি, কম্পাস",
		"GPS, GLONASS, BDS":                 "GPS, গ্লোনাস, বিডিএস",
		"5.1":                               "5.1",
		"Wi-Fi 802.11 a/b/g/n/ac":           "ওয়াই-ফাই 5",
		"USB-C 2.0":                         "USB-C",
		"Dual SIM (Nano-SIM)":               "ডুয়াল সিম",
		"Yes, 3.5 mm":                       "হ্যাঁ, 3.5 mm",
		"Single":                            "সিঙ্গেল",
		"Mono":                              "মনো",
		"Li-Ion (non-removable)":            "লি-আয়ন",
		"18 W wired":                        "18W",
		"microSDXC":                         "মাইক্রো এসডিএক্সসি",
		"2.2 GHz":                           "2.2 GHz",
		"LED flash, HDR":                    "LED ফ্ল্যাশ, HDR",
		"4K @ 30fps; 1080p @ 30fps":         "4K 30fps",
		"2x":                                "2x",
		"1080p @ 30fps":                     "1080p 30fps",
		"0.35 W/kg (head) 1.33 W/kg (body)": "0.35 W/kg (হেড) 1.33 W/kg (বডি)",
		"0.19 W/kg (head) 1.33 W/kg (body)": "0.19 W/kg (হেড) 1.33 W/kg (বডি)",
		"ZENX 1":                            "ZENX 1",
		"Fast charging, reverse charging":   "ফাস্ট চার্জিং, রিভার্স চার্জিং",
	}
	for k, v := range s.GetCommonBanglaTranslations() {
		translations[k] = v
	}
	return translations
}

// Seed inserts specification records for the product identified by slug 'walton-zenx-1'
func (s *SpecificationSeederMobileWaltonZENX1) Seed(db *gorm.DB) error {
	productSlug := "walton-zenx-1"
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
		"dimensions":                    "168.5 x 76.5 x 9.2 mm",
		"network_technology":            "5G",
		"rear_camera":                   "64 MP + 8 MP + 2 MP",
		"available_colors":              "Silver, Black",
		"device_status":                 "Available",
		"chipset":                       "Qualcomm Snapdragon 695 5G (6 nm)",
		"gpu_type":                      "Adreno 619",
		"ram":                           "8 GB",
		"storage":                       "128 GB",
		"battery":                       "5,000 mAh",
		"operating_system":              "Android 13",
		"display_size":                  "6.8 inches",
		"cpu_type":                      "Octa-core",
		"display_type":                  "IPS LCD, 120Hz",
		"resolution":                    "1080 x 2460 pixels",
		"refresh_rate":                  "120Hz",
		"weight":                        "205 g",
		"processor":                     "Snapdragon 695",
		"screen_protection":             "Corning Gorilla Glass 3",
		"build_material":                "Glass front, plastic back",
		"water_resistance":              "No",
		"front_camera":                  "16 MP",
		"announcement_date":             "May 2023",
		"sensors":                       "Fingerprint (side-mounted), accelerometer, proximity, compass",
		"positioning_system":            "GPS, GLONASS, BDS",
		"bluetooth_version":             "5.1",
		"wifi_support":                  "Wi-Fi 802.11 a/b/g/n/ac",
		"usb_type":                      "USB-C 2.0",
		"sim_card_type":                 "Dual SIM (Nano-SIM)",
		"audio_jack":                    "Yes, 3.5 mm",
		"loudspeaker_quality":           "Single",
		"audio_quality":                 "Mono",
		"battery_type":                  "Li-Ion (non-removable)",
		"charging_speed":                "18 W wired",
		"fast_charging":                 "Yes",
		"wireless_charging":             "No",
		"card_slot_type":                "microSDXC",
		"internal_memory_capacity":      "128 GB",
		"processor_speed":               "2.2 GHz",
		"camera_features":               "LED flash, HDR",
		"quad_camera_setup":             "64 MP + 8 MP + 2 MP",
		"camera_video_resolution":       "4K @ 30fps; 1080p @ 30fps",
		"optical_zoom":                  "2x",
		"front_camera_video_resolution": "1080p @ 30fps",
		"nfc_support":                   "No",
		"sar_rating":                    "0.35 W/kg (head) 1.33 W/kg (body)",
		"sar_rating_eu":                 "0.19 W/kg (head) 1.33 W/kg (body)",
		"model_variants":                "ZENX 1",
		"special_features":              "Fast charging, reverse charging",
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
