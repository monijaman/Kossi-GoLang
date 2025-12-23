package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederMobileSamsungGalaxyF545G seeds specifications/options for product 'samsung-galaxy-f54-5g'
type SpecificationSeederMobileSamsungGalaxyF545G struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSamsungGalaxyF545G creates a new seeder instance
func NewSpecificationSeederMobileSamsungGalaxyF545G() *SpecificationSeederMobileSamsungGalaxyF545G {
	return &SpecificationSeederMobileSamsungGalaxyF545G{BaseSeeder: BaseSeeder{name: "Specifications for Samsung Galaxy F54 5G"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSamsungGalaxyF545G) getBanglaTranslations() map[string]string {
	common := s.GetCommonBanglaTranslations()
	specific := map[string]string{
		// Add specific translations here if needed
		"Exynos 1380":             "এক্সিনস ১৩৮০",
		"Mali-G68 MP5":            "मालि-जी६८ एमपी५",
		"Octa-core":               "অক্টা-কোর",
		"6 / 8 GB":                "६ / ८ जिबि",
		"128 / 256 GB":            "१२८ / २५६ जिबि",
		"Glass front & back":      "সামনে এবং পিছনে গ্লাস",
		"Wi-Fi 802.11 a/b/g/n/ac": "ওয়াই-ফাই ৮০২.১১ a/b/g/n/ac",
		"1080p @ 30fps":           "১০৮০p @ ৩০fps",
		"GSM / HSPA / LTE / 5G":   "জিএসএম / এইচএসপিএ / এলটিই / ৫জি",
		"Side fingerprint, Accelerometer, Gyro, Compass": "পাশের ফিঙ্গারপ্রিন্ট, অ্যাক্সেলেরোমিটার, জাইরো, কম্পাস",
		"Yes":                          "হ্যাঁ",
		"2025":                         "২০২৫",
		"Available":                    "উপলব্ধ",
		"5000 mAh":                     "৫০০০ mAh",
		"1080 × 2400 px":               "১০৮০ × ২৪০০ px",
		"202 g":                        "২০২ g",
		"108 MP + 12 MP + 5 MP":        "১০৮ MP + ১২ MP + ৫ MP",
		"Android 14, One UI 6":         "অ্যান্ড্রয়েড ১৪, ওয়ান UI ৬",
		"25 W wired":                   "২৫ W তারযুক্ত",
		"Graphite, Blue, Violet":       "গ্রাফাইট, ব্লু, ভায়োলেট",
		"120 Hz":                       "১২০ Hz",
		"2x":                           "২x",
		"32 MP":                        "৩২ MP",
		"No":                           "না",
		"5.3":                          "৫.৩",
		"OIS, LED flash":               "OIS, LED ফ্ল্যাশ",
		"4K @ 30fps; 1080p @ 30/60fps": "4K @ 30fps; 1080p @ 30/60fps",
		"Nano-SIM + Nano-SIM":          "ন্যানো-SIM + ন্যানো-SIM",
		"165 × 77 × 8.2 mm":            "১৬৫ × ৭৭ × ৮.২ mm",
		"USB-C":                        "USB-C",
		"GPS, GLONASS, GALILEO, BDS":   "GPS, GLONASS, GALILEO, BDS",
	}
	for k, v := range specific {
		common[k] = v
	}
	return common
}

// Seed inserts specification records for the product identified by slug 'samsung-galaxy-f54-5g'
func (s *SpecificationSeederMobileSamsungGalaxyF545G) Seed(db *gorm.DB) error {
	productSlug := "samsung-galaxy-f54-5g"
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
		"nfc_support":                   "Yes",
		"rear_camera":                   "108 MP + 12 MP + 5 MP",
		"battery":                       "5000 mAh",
		"fast_charging":                 "25 W wired",
		"sensors":                       "Side fingerprint, Accelerometer, Gyro, Compass",
		"audio_jack":                    "No",
		"chipset":                       "Exynos 1380",
		"ram":                           "6 / 8 GB",
		"weight":                        "202 g",
		"dimensions":                    "165 × 77 × 8.2 mm",
		"camera_video_resolution":       "4K @ 30fps; 1080p @ 30/60fps",
		"battery_type":                  "Li-Ion (non-removable)",
		"wireless_charging":             "No",
		"loudspeaker_quality":           "Yes",
		"display_size":                  "6.7 inches",
		"cpu_type":                      "Octa-core",
		"gpu_type":                      "Mali-G68 MP5",
		"storage":                       "128 / 256 GB",
		"wifi_support":                  "Wi-Fi 802.11 a/b/g/n/ac",
		"usb_type":                      "USB-C",
		"operating_system":              "Android 14, One UI 6",
		"positioning_system":            "GPS, GLONASS, GALILEO, BDS",
		"front_camera_video_resolution": "1080p @ 30fps",
		"audio_quality":                 "",
		"available_colors":              "Graphite, Blue, Violet",
		"announcement_date":             "2025",
		"processor":                     "Exynos 1380",
		"display_type":                  "Super AMOLED+, 120 Hz",
		"resolution":                    "1080 × 2400 px",
		"build_material":                "Glass front & back",
		"water_resistance":              "IP67",
		"network_technology":            "GSM / HSPA / LTE / 5G",
		"optical_zoom":                  "2x",
		"refresh_rate":                  "120 Hz",
		"bluetooth_version":             "5.3",
		"sim_card_type":                 "Nano-SIM + Nano-SIM",
		"device_status":                 "Available",
		"camera_features":               "OIS, LED flash",
		"5g_support":                    "Yes",
		"front_camera":                  "32 MP",
		"special_features":              "",
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
