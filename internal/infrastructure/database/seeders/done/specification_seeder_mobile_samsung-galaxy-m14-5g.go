package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederMobileSamsungGalaxyM145G seeds specifications/options for product 'samsung-galaxy-m14-5g'
type SpecificationSeederMobileSamsungGalaxyM145G struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSamsungGalaxyM145G creates a new seeder instance
func NewSpecificationSeederMobileSamsungGalaxyM145G() *SpecificationSeederMobileSamsungGalaxyM145G {
	return &SpecificationSeederMobileSamsungGalaxyM145G{BaseSeeder: BaseSeeder{name: "Specifications for Samsung Galaxy M14 5G"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSamsungGalaxyM145G) getBanglaTranslations() map[string]string {
	common := s.GetCommonBanglaTranslations()
	specific := map[string]string{
		"PLS LCD":                         "পিএলএস LCD",
		"50 MP + 2 MP + 2 MP":             "৫০ MP + ২ MP + ২ MP",
		"Yes":                             "হ্যাঁ",
		"Exynos 1330 (5 nm)":              "এক্সিনস ১৩৩০ (৫ nm)",
		"13 MP":                           "১৩ MP",
		"Li-Ion (non-removable)":          "লি-আয়ন (অপসারণযোগ্য নয়)",
		"No":                              "না",
		"Android 13, One UI 5.1":          "অ্যান্ড্রয়েড ১৩, ওয়ান UI ৫.১",
		"6000 mAh":                        "৬০০০ mAh",
		"6.6 inches":                      "৬.৬ ইঞ্চি",
		"64 / 128 GB + microSD":           "৬৪ / ১২৮ GB + microSD",
		"206 g":                           "২০৬ g",
		"166.8 × 77.2 × 9.4 mm":           "১৬৬.৮ × ৭৭.২ × ৯.৪ mm",
		"1080p @ 30fps":                   "১০৮০p @ ৩০fps",
		"March 2023":                      "মার্চ ২০২৩",
		"4 / 6 GB":                        "৪ / ৬ GB",
		"Glass front, plastic back/frame": "সামনে গ্লাস, পিছনে/ফ্রেম প্লাস্টিক",
		"GSM / HSPA / LTE / 5G":           "জিএসএম / এইচএসপিএ / এলটিই / ৫জি",
		"Nano-SIM or Dual SIM (hybrid)":   "ন্যানো-সিম বা ডুয়াল সিম (হাইব্রিড)",
		"Mono":                            "মনো",
		"3.5 mm":                          "৩.৫ mm",
		"Exynos 1330":                     "এক্সিনস ১৩৩০",
		"Mali-G68 MP2":                    "মালি-G৬৮ MP২",
		"USB-C 2.0":                       "ইউএসবি-সি ২.০",
		"15 W / 25 W (region dependent)":  "১৫ W / ২৫ W (অঞ্চল নির্ভর)",
		"Navy Blue, Light Blue, Silver":   "নেভি ব্লু, লাইট ব্লু, সিলভার",
		"Octa-core (2×2.4 GHz Cortex-A78 + 6×2.0 GHz Cortex-A55)": "অক্টা-কোর (২×২.৪ GHz Cortex-A78 + ৬×২.০ GHz Cortex-A55)",
		"2408 × 1080 px": "২৪০৮ × ১০৮০ px",
		"90 Hz":          "৯০ Hz",
		"None":           "কোনটি না",
		"Side fingerprint, Accelerometer, Gyro, Proximity, Compass": "সাইড ফিঙ্গারপ্রিন্ট, অ্যাক্সেলেরোমিটার, গাইরো, প্রক্সিমিটি, কম্পাস",
		"Wi-Fi 802.11 a/b/g/n/ac":                                   "ওয়াই-ফাই ৮০২.১১ a/b/g/n/ac",
		"LED flash, Depth":                                          "এলইডি ফ্ল্যাশ, ডেপথ",
		"Available":                                                 "উপলব্ধ",
		"GSM 850 / 900 / 1800 / 1900":                               "জিএসএম ৮৫০ / ৯০০ / ১৮০০ / ১৯০০",
		"HSDPA 850 / 900 / 1900 / 2100":                             "এইচএসডিপিএ ৮৫০ / ৯০০ / ১৯০০ / ২১০০",
		"LTE":                                                       "এলটিই",
		"SA/NSA/Sub6":                                               "এসএ/এনএসএ/সাব৬",
		"microSDXC (hybrid)":                                        "মাইক্রো এসডিএক্সসি (হাইব্রিড)",
		"2.4 GHz":                                                   "২.৪ GHz",
		"15 W":                                                      "১৫ W",
	}
	for k, v := range specific {
		common[k] = v
	}
	return common
}

// Seed inserts specification records for the product identified by slug 'samsung-galaxy-m14-5g'
func (s *SpecificationSeederMobileSamsungGalaxyM145G) Seed(db *gorm.DB) error {
	productSlug := "samsung-galaxy-m14-5g"
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
		"display_type":                  "PLS LCD",
		"bluetooth_version":             "",
		"rear_camera":                   "50 MP + 2 MP + 2 MP",
		"5g_support":                    "Yes",
		"chipset":                       "Exynos 1330 (5 nm)",
		"front_camera":                  "13 MP",
		"battery_type":                  "Li-Ion (non-removable)",
		"special_features":              "",
		"water_resistance":              "",
		"nfc_support":                   "",
		"operating_system":              "Android 13, One UI 5.1",
		"battery":                       "6000 mAh",
		"display_size":                  "6.6 inches",
		"storage":                       "64 / 128 GB + microSD",
		"weight":                        "206 g",
		"dimensions":                    "166.8 × 77.2 × 9.4 mm",
		"front_camera_video_resolution": "1080p @ 30fps",
		"positioning_system":            "",
		"announcement_date":             "March 2023",
		"ram":                           "4 / 6 GB",
		"build_material":                "Glass front, plastic back/frame",
		"network_technology":            "GSM / HSPA / LTE / 5G",
		"2g_bands":                      "GSM 850 / 900 / 1800 / 1900",
		"3g_bands":                      "HSDPA 850 / 900 / 1900 / 2100",
		"4g_bands":                      "LTE",
		"5g_bands":                      "SA/NSA/Sub6",
		"sim_card_type":                 "Nano-SIM or Dual SIM (hybrid)",
		"loudspeaker_quality":           "Mono",
		"audio_quality":                 "",
		"audio_jack":                    "3.5 mm",
		"processor":                     "Exynos 1330",
		"gpu_type":                      "Mali-G68 MP2",
		"usb_type":                      "USB-C 2.0",
		"camera_video_resolution":       "1080p @ 30fps",
		"fast_charging":                 "15 W / 25 W (region dependent)",
		"wireless_charging":             "No",
		"available_colors":              "Navy Blue, Light Blue, Silver",
		"cpu_type":                      "Octa-core (2×2.4 GHz Cortex-A78 + 6×2.0 GHz Cortex-A55)",
		"resolution":                    "2408 × 1080 px",
		"refresh_rate":                  "90 Hz",
		"optical_zoom":                  "None",
		"sensors":                       "Side fingerprint, Accelerometer, Gyro, Proximity, Compass",
		"wifi_support":                  "Wi-Fi 802.11 a/b/g/n/ac",
		"camera_features":               "LED flash, Depth",
		"device_status":                 "Available",
		"internal_memory_capacity":      "64 / 128 GB",
		"card_slot_type":                "microSDXC (hybrid)",
		"processor_speed":               "2.4 GHz",
		"charging_speed":                "15 W",
		"screen_protection":             "",
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
