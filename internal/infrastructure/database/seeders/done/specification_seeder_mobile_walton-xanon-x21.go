package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederMobileWaltonXANONX21 seeds specifications/options for product 'walton-xanon-x21'
type SpecificationSeederMobileWaltonXANONX21 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileWaltonXANONX21 creates a new seeder instance
func NewSpecificationSeederMobileWaltonXANONX21() *SpecificationSeederMobileWaltonXANONX21 {
	return &SpecificationSeederMobileWaltonXANONX21{BaseSeeder: BaseSeeder{name: "Specifications for Walton XANON X21"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileWaltonXANONX21) getBanglaTranslations() map[string]string {
	translations := map[string]string{
		"163 x 76 x 8 mm":           "163x76x8 মিমি",
		"Helio G99":                 "হেলিও G99",
		"Mediatek Helio G99 (6 nm)": "মিডিয়াটেক হেলিও G99",
		"8 GB":                      "8 জিবি",
		"256 GB":                    "256 জিবি",
		"1080 x 2400 pixels":        "1080x2400",
		"Glass front, plastic frame, plastic back": "গ্লাস ফ্রন্ট, প্লাস্টিক ফ্রেম, প্লাস্টিক ব্যাক",
		"No":                      "না",
		"Black, Gold":             "ব্ল্যাক, গোল্ড",
		"180 g":                   "180 গ্রাম",
		"4G":                      "4G",
		"5,000 mAh":               "5000 mAh",
		"October 2023":            "অক্টোবর 2023",
		"6.7 inches":              "6.7 ইঞ্চি",
		"120Hz":                   "120Hz",
		"50 MP + 2 MP":            "50+2 মেগা",
		"32 MP":                   "32 মেগা",
		"Android 13":              "অ্যান্ড্রয়েড 13",
		"Available":               "উপলব্ধ",
		"Octa-core":               "অক্টা-কোর",
		"Mali-G57 MC2":            "মালি-G57 MC2",
		"AMOLED, 120Hz":           "AMOLED 120Hz",
		"Corning Gorilla Glass 5": "কর্নিং গরিলা গ্লাস 5",
		"Fingerprint (under display, optical), accelerometer, gyro, proximity, compass": "ফিঙ্গারপ্রিন্ট (আন্ডার ডিসপ্লে, অপটিকাল), অ্যাক্সেলেরোমিটার, জাইরো, প্রক্সিমিটি, কম্পাস",
		"GPS, GLONASS, BDS":                  "GPS, গ্লোনাস, বিডিএস",
		"5.3":                                "5.3",
		"Wi-Fi 802.11 a/b/g/n/ac/6":          "ওয়াই-ফাই 6",
		"USB-C 2.0":                          "USB-C",
		"Dual SIM (Nano-SIM, dual stand-by)": "ডুয়াল সিম",
		"Yes, 3.5 mm":                        "হ্যাঁ, 3.5 mm",
		"Single":                             "সিঙ্গেল",
		"Mono":                               "মনো",
		"Li-Ion (non-removable)":             "লি-আয়ন",
		"18 W wired":                         "18W",
		"microSDXC (dedicated slot)":         "মাইক্রো এসডিএক্সসি",
		"2.2 GHz":                            "2.2 GHz",
		"LED flash, HDR, panorama":           "LED ফ্ল্যাশ, HDR, প্যানোরামা",
		"4K @ 30fps; 1080p @ 30fps":          "4K 30fps",
		"2x":                                 "2x",
		"1080p @ 30fps":                      "1080p 30fps",
		"0.43 W/kg (head) 1.29 W/kg (body)":  "0.43 W/kg (হেড) 1.29 W/kg (বডি)",
		"0.32 W/kg (head) 1.50 W/kg (body)":  "0.32 W/kg (হেড) 1.50 W/kg (বডি)",
		"XANON X21":                          "XANON X21",
		"Fast charging":                      "ফাস্ট চার্জিং",
	}
	for k, v := range s.GetCommonBanglaTranslations() {
		translations[k] = v
	}
	return translations
}

// Seed inserts specification records for the product identified by slug 'walton-xanon-x21'
func (s *SpecificationSeederMobileWaltonXANONX21) Seed(db *gorm.DB) error {
	productSlug := "walton-xanon-x21"
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
		"dimensions":                    "163 x 76 x 8 mm",
		"processor":                     "Helio G99",
		"chipset":                       "Mediatek Helio G99 (6 nm)",
		"ram":                           "8 GB",
		"storage":                       "256 GB",
		"resolution":                    "1080 x 2400 pixels",
		"build_material":                "Glass front, plastic frame, plastic back",
		"water_resistance":              "No",
		"available_colors":              "Black, Gold",
		"weight":                        "180 g",
		"network_technology":            "4G",
		"battery":                       "5,000 mAh",
		"announcement_date":             "October 2023",
		"display_size":                  "6.7 inches",
		"refresh_rate":                  "120Hz",
		"rear_camera":                   "50 MP + 2 MP",
		"front_camera":                  "32 MP",
		"operating_system":              "Android 13",
		"device_status":                 "Available",
		"cpu_type":                      "Octa-core",
		"gpu_type":                      "Mali-G57 MC2",
		"display_type":                  "AMOLED, 120Hz",
		"screen_protection":             "Corning Gorilla Glass 5",
		"sensors":                       "Fingerprint (under display, optical), accelerometer, gyro, proximity, compass",
		"positioning_system":            "GPS, GLONASS, BDS",
		"bluetooth_version":             "5.3",
		"wifi_support":                  "Wi-Fi 802.11 a/b/g/n/ac/6",
		"usb_type":                      "USB-C 2.0",
		"sim_card_type":                 "Dual SIM (Nano-SIM, dual stand-by)",
		"audio_jack":                    "Yes, 3.5 mm",
		"loudspeaker_quality":           "Single",
		"audio_quality":                 "Mono",
		"battery_type":                  "Li-Ion (non-removable)",
		"charging_speed":                "18 W wired",
		"fast_charging":                 "Yes",
		"wireless_charging":             "No",
		"card_slot_type":                "microSDXC (dedicated slot)",
		"internal_memory_capacity":      "256 GB",
		"processor_speed":               "2.2 GHz",
		"camera_features":               "LED flash, HDR, panorama",
		"quad_camera_setup":             "50 MP + 2 MP",
		"camera_video_resolution":       "4K @ 30fps; 1080p @ 30fps",
		"optical_zoom":                  "2x",
		"front_camera_video_resolution": "1080p @ 30fps",
		"nfc_support":                   "No",
		"sar_rating":                    "0.43 W/kg (head) 1.29 W/kg (body)",
		"sar_rating_eu":                 "0.32 W/kg (head) 1.50 W/kg (body)",
		"model_variants":                "XANON X21",
		"special_features":              "Fast charging",
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
