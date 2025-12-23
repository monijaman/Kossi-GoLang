package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederMobileSamsungGalaxyM065G seeds specifications/options for product 'samsung-galaxy-m06-5g'
type SpecificationSeederMobileSamsungGalaxyM065G struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSamsungGalaxyM065G creates a new seeder instance
func NewSpecificationSeederMobileSamsungGalaxyM065G() *SpecificationSeederMobileSamsungGalaxyM065G {
	return &SpecificationSeederMobileSamsungGalaxyM065G{BaseSeeder: BaseSeeder{name: "Specifications for Samsung Galaxy M06 5G"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSamsungGalaxyM065G) getBanglaTranslations() map[string]string {
	common := s.GetCommonBanglaTranslations()
	specific := map[string]string{
		"Li-Ion (non-removable)":        "লি-আয়ন (অপসারণযোগ্য নয়)",
		"25 W wired (likely)":           "২৫ W তারযুক্ত (সম্ভবত)",
		"Yes":                           "হ্যাঁ",
		"Nano-SIM + Nano-SIM or hybrid": "ন্যানো-সিম + ন্যানো-সিম বা হাইব্রিড",
		"Mono":                          "মনো",
		"64 / 128 GB":                   "৬৪ / ১২৮ GB",
		"No":                            "না",
		"2025":                          "২০২৫",
		"4 / 6 GB":                      "৪ / ৬ GB",
		"Available":                     "উপলব্ধ",
		"6.74 inches":                   "৬.৭৪ ইঞ্চি",
		"720 × 1600 px (reportedly)":    "৭২০ × ১৬০০ px (রিপোর্ট অনুসারে)",
		"Glass front, plastic back":     "সামনে গ্লাস, পিছনে প্লাস্টিক",
		"GSM / HSPA / LTE / 5G":         "জিএসএম / এইচএসপিএ / এলটিই / ৫জি",
		"Wi-Fi 802.11 a/b/g/n/ac":       "ওয়াই-ফাই ৮০২.১১ a/b/g/n/ac",
		"5.3":                           "৫.৩",
		"LED flash, Depth":              "এলইডি ফ্ল্যাশ, ডেপথ",
		"8 MP":                          "৮ MP",
		"Accelerometer, Gyro, Proximity, Compass, Fingerprint (side?)": "অ্যাক্সেলেরোমিটার, গাইরো, প্রক্সিমিটি, কম্পাস, ফিঙ্গারপ্রিন্ট (সাইড?)",
		"MediaTek Dimensity 6300":                                      "মিডিয়াটেক ডাইমেনসিটি ৬৩০০",
		"3.5 mm (some sources)":                                        "৩.৫ mm (কিছু সূত্র)",
		"Dimensity 6300 (6 nm)":                                        "ডাইমেনসিটি ৬৩০০ (৬ nm)",
		"90 Hz":                                                        "৯০ Hz",
		"50 MP + 2 MP":                                                 "৫০ MP + ২ MP",
		"Android (version not clearly confirmed)":                      "অ্যান্ড্রয়েড (ভার্সন স্পষ্টভাবে নিশ্চিত নয়)",
		"Octa-core":                                                    "অক্টা-কোর",
		"PLS LCD":                                                      "পিএলএস LCD",
		"5000 mAh":                                                     "৫০০০ mAh",
		"167.4 × 77.4 × 8 mm":                                          "১৬৭.৪ × ৭৭.৪ × ৮ mm",
		"Blazing Black, Sage Green":                                    "ব্লেজিং ব্ল্যাক, সেজ গ্রীন",
		"GSM 850 / 900 / 1800 / 1900":                                  "GSM ৮৫০ / ৯০০ / ১৮০০ / ১৯০০",
		"HSDPA 850 / 900 / 1900 / 2100":                                "HSDPA ৮৫০ / ৯০০ / ১৯০০ / ২১০০",
		"LTE":                                                          "এলটিই",
		"SA/NSA/Sub6":                                                  "এসএ/এনএসএ/সাব৬",
	}
	for k, v := range specific {
		common[k] = v
	}
	return common
}

// Seed inserts specification records for the product identified by slug 'samsung-galaxy-m06-5g'
func (s *SpecificationSeederMobileSamsungGalaxyM065G) Seed(db *gorm.DB) error {
	productSlug := "samsung-galaxy-m06-5g"
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
		"gpu_type":                      "",
		"dimensions":                    "167.4 × 77.4 × 8 mm",
		"optical_zoom":                  "None",
		"front_camera_video_resolution": "",
		"sim_card_type":                 "Nano-SIM + Nano-SIM or hybrid",
		"loudspeaker_quality":           "Mono",
		"storage":                       "64 / 128 GB",
		"weight":                        "",
		"wireless_charging":             "No",
		"announcement_date":             "2025",
		"ram":                           "4 / 6 GB",
		"nfc_support":                   "",
		"device_status":                 "Available",
		"display_size":                  "6.74 inches",
		"camera_video_resolution":       "",
		"battery_type":                  "Li-Ion (non-removable)",
		"available_colors":              "Blazing Black, Sage Green",
		"resolution":                    "720 × 1600 px (reportedly)",
		"build_material":                "Glass front, plastic back",
		"network_technology":            "GSM / HSPA / LTE / 5G",
		"2g_bands":                      "GSM 850 / 900 / 1800 / 1900",
		"3g_bands":                      "HSDPA 850 / 900 / 1900 / 2100",
		"4g_bands":                      "LTE",
		"5g_bands":                      "SA/NSA/Sub6",
		"wifi_support":                  "Wi-Fi 802.11 a/b/g/n/ac",
		"bluetooth_version":             "5.3",
		"camera_features":               "LED flash, Depth",
		"front_camera":                  "8 MP",
		"sensors":                       "Accelerometer, Gyro, Proximity, Compass, Fingerprint (side?)",
		"processor":                     "MediaTek Dimensity 6300",
		"water_resistance":              "",
		"special_features":              "",
		"audio_quality":                 "",
		"audio_jack":                    "3.5 mm (some sources)",
		"chipset":                       "Dimensity 6300 (6 nm)",
		"refresh_rate":                  "90 Hz",
		"usb_type":                      "USB-C",
		"rear_camera":                   "50 MP + 2 MP",
		"operating_system":              "Android (version not clearly confirmed)",
		"fast_charging":                 "25 W wired (likely)",
		"cpu_type":                      "Octa-core",
		"display_type":                  "PLS LCD",
		"battery":                       "5000 mAh",
		"5g_support":                    "Yes",
		"positioning_system":            "",
		"internal_memory_capacity":      "64 / 128 GB",
		"card_slot_type":                "microSDXC (hybrid)",
		"processor_speed":               "2.2 GHz",
		"charging_speed":                "25 W",
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
