package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederMobileVivoX200Pro seeds specifications/options for product 'vivo-x200-pro'
type SpecificationSeederMobileVivoX200Pro struct {
	BaseSeeder
}

// NewSpecificationSeederMobileVivoX200Pro creates a new seeder instance
func NewSpecificationSeederMobileVivoX200Pro() *SpecificationSeederMobileVivoX200Pro {
	return &SpecificationSeederMobileVivoX200Pro{BaseSeeder: BaseSeeder{name: "Specifications for Vivo X200 Pro"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileVivoX200Pro) getBanglaTranslations() map[string]string {
	translations := map[string]string{
		"Octa-core":                  "অক্টা-কোর",
		"Mali-G77 MC9":               "Mali-G77",
		"Li-Ion (non-removable)":     "লি-আয়ন",
		"33 W wired":                 "33W",
		"GPS, GLONASS, GALILEO, BDS": "GPS",
		"Stereo":                     "স্টেরিও",
		"Available":                  "উপলব্ধ",
		"196 g":                      "196 গ্রাম",
		"Wi-Fi 802.11 a/b/g/n/ac/6":  "ওয়াই-ফাই",
		"5.2":                        "5.2",
		"32 MP":                      "32 মেগা",
		"Hi-Res Audio":               "হাই-রেজ",
		"No":                         "না",
		"Fingerprint (under display), Accelerometer, Gyro, Proximity, Compass": "ফিঙ্গারপ্রিন্ট",
		"Glass front/back, plastic frame":                                      "গ্লাস ফ্রন্ট",
		"64 MP + 8 MP + 2 MP":                                                  "64+8+2 মেগা",
		"4200 mAh":                                                             "4200 mAh",
		"December 2022":                                                        "ডিসেম্বর 2022",
		"8 / 12 GB":                                                            "8/12 জিবি",
		"2376 × 1080 px":                                                       "2376x1080",
		"1080p @ 30fps":                                                        "1080p",
		"Android 11, Funtouch OS 11":                                           "অ্যান্ড্রয়েড 11",
		"AMOLED HDR10+":                                                        "AMOLED HDR10+",
		"Dual SIM (Nano-SIM)":                                                  "ডুয়াল সিম",
		"Dimensity 1200 (6 nm)":                                                "ডাইমেনসিটি 1200",
		"GSM / HSPA / LTE / 5G":                                                "GSM/LTE/5G",
		"Yes":                                                                  "হ্যাঁ",
		"OIS, LED flash":                                                       "LED ফ্ল্যাশ",
		"Black, Silver":                                                        "কালো, সিলভার",
		"MediaTek Dimensity 1200":                                              "মিডিয়াটেক 1200",
		"120 Hz":                                                               "120 হার্জ",
		"USB-C 3.1":                                                            "USB-C",
		"6.56 inches":                                                          "6.56 ইঞ্চি",
		"128 / 256 GB":                                                         "128/256 জিবি",
		"AMOLED, 120 Hz":                                                       "AMOLED 120Hz",
		"4K @ 30/60fps; 1080p @ 30/120fps":                                     "4K 30fps",
		"2x":                                                                   "2x",
	}
	for k, v := range s.GetCommonBanglaTranslations() {
		if _, exists := translations[k]; !exists {
			translations[k] = v
		}
	}
	return translations
}

// Seed inserts specification records for the product identified by slug 'vivo-x200-pro'
func (s *SpecificationSeederMobileVivoX200Pro) Seed(db *gorm.DB) error {
	productSlug := "vivo-x200-pro"
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
		"gpu_type":                      "Mali-G77 MC9",
		"battery_type":                  "Li-Ion (non-removable)",
		"fast_charging":                 "33 W wired",
		"positioning_system":            "GPS, GLONASS, GALILEO, BDS",
		"loudspeaker_quality":           "Stereo",
		"device_status":                 "Available",
		"weight":                        "196 g",
		"wifi_support":                  "Wi-Fi 802.11 a/b/g/n/ac/6",
		"bluetooth_version":             "5.2",
		"front_camera":                  "32 MP",
		"audio_quality":                 "Hi-Res Audio",
		"wireless_charging":             "No",
		"sensors":                       "Fingerprint (under display), Accelerometer, Gyro, Proximity, Compass",
		"audio_jack":                    "No",
		"build_material":                "Glass front/back, plastic frame",
		"rear_camera":                   "64 MP + 8 MP + 2 MP",
		"battery":                       "4200 mAh",
		"announcement_date":             "December 2022",
		"ram":                           "8 / 12 GB",
		"resolution":                    "2376 × 1080 px",
		"front_camera_video_resolution": "1080p @ 30fps",
		"operating_system":              "Android 11, Funtouch OS 11",
		"special_features":              "AMOLED HDR10+",
		"sim_card_type":                 "Dual SIM (Nano-SIM)",
		"chipset":                       "Dimensity 1200 (6 nm)",
		"network_technology":            "GSM / HSPA / LTE / 5G",
		"nfc_support":                   "Yes",
		"camera_features":               "OIS, LED flash",
		"5g_support":                    "Yes",
		"available_colors":              "Black, Silver",
		"processor":                     "MediaTek Dimensity 1200",
		"refresh_rate":                  "120 Hz",
		"usb_type":                      "USB-C 3.1",
		"display_size":                  "6.56 inches",
		"storage":                       "128 / 256 GB",
		"display_type":                  "AMOLED, 120 Hz",
		"camera_video_resolution":       "4K @ 30/60fps; 1080p @ 30/120fps",
		"optical_zoom":                  "2x",
	}
	banglaTranslations := s.getBanglaTranslations()
	for key, value := range specs {
		specKeyID, exists := existingkeyMapping[key]
		if !exists {
			log.Printf("⚠️  Key not found in existingkeyMapping: '%s' (used in product: %s)", key, productSlug)
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
