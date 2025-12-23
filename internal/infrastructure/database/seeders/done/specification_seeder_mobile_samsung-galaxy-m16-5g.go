package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederMobileSamsungGalaxyM165G seeds specifications/options for product 'samsung-galaxy-m16-5g'
type SpecificationSeederMobileSamsungGalaxyM165G struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSamsungGalaxyM165G creates a new seeder instance
func NewSpecificationSeederMobileSamsungGalaxyM165G() *SpecificationSeederMobileSamsungGalaxyM165G {
	return &SpecificationSeederMobileSamsungGalaxyM165G{BaseSeeder: BaseSeeder{name: "Specifications for Samsung Galaxy M16 5G"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSamsungGalaxyM165G) getBanglaTranslations() map[string]string {
	common := s.GetCommonBanglaTranslations()
	specific := map[string]string{
		"Wi-Fi 802.11 a/b/g/n/ac": "ওয়াই-ফাই ৮০২.১১ a/b/g/n/ac",
		"None":                    "কোনটি না",
		"Side fingerprint, Accelerometer, Gyro, Proximity, Compass": "সাইড ফিঙ্গারপ্রিন্ট, অ্যাক্সেলেরোমিটার, গাইরো, প্রক্সিমিটি, কম্পাস",
		"Li-Ion (non-removable)":                                    "লি-আয়ন (অপসারণযোগ্য নয়)",
		"Yes":                                                       "হ্যাঁ",
		"6 years of updates (claimed)":                              "৬ বছরের আপডেট (দাবি করা)",
		"6 / 8 GB":                                                  "৬ / ৮ GB",
		"1080 × 2340 px":                                            "১০৮০ × ২৩৪০ px",
		"90 Hz":                                                     "৯০ Hz",
		"Thunder Black, Mint Green, Blush Pink":                     "থান্ডার ব্ল্যাক, মিন্ট গ্রীন, ব্লাশ পিঙ্ক",
		"Mali-G57 MC2":                                              "মালি-G৫৭ MC২",
		"5.3":                                                       "৫.৩",
		"Android 15, One UI 7":                                      "অ্যান্ড্রয়েড ১৫, ওয়ান UI ৭",
		"5000 mAh":                                                  "৫০০০ mAh",
		"No":                                                        "না",
		"6.7 inches":                                                "৬.৭ ইঞ্চি",
		"Octa-core (2×2.4 GHz Cortex-A76 + 6×2.0 GHz Cortex-A55)": "অক্টা-কোর (২×২.৪ GHz Cortex-A76 + ৬×২.০ GHz Cortex-A55)",
		"Glass front, plastic back / frame":                       "সামনে গ্লাস, পিছনে / ফ্রেম প্লাস্টিক",
		"GSM / HSPA / LTE / 5G":                                   "জিএসএম / এইচএসপিএ / এলটিই / ৫জি",
		"13 MP":                                                   "১৩ MP",
		"25 W wired":                                              "২৫ W তারযুক্ত",
		"GPS, GLONASS, GALILEO, BDS":                              "জিপিএস, গ্লোনাস, গ্যালিলিও, বিডিএস",
		"Dimensity 6300 (6 nm)":                                   "ডাইমেনসিটি ৬৩০০ (৬ nm)",
		"IP54 (splash)":                                           "IP54 (স্প্ল্যাশ)",
		"50 MP + 5 MP + 2 MP":                                     "৫০ MP + ৫ MP + ২ MP",
		"Nano-SIM + Nano-SIM (hybrid)":                            "ন্যানো-সিম + ন্যানো-সিম (হাইব্রিড)",
		"USB-C (no 3.5mm)":                                        "ইউএসবি-সি (৩.৫mm নেই)",
		"Super AMOLED":                                            "সুপার AMOLED",
		"LED flash, Auto-focus, Depth / Macro":                    "এলইডি ফ্ল্যাশ, অটো-ফোকাস, ডেপথ / ম্যাক্রো",
		"Stereo":                                                  "স্টেরিও",
		"128 GB (maybe more)":                                     "১২৮ GB (হয়তো আরও)",
		"164.4 × 77.9 × 7.9 mm":                                   "১৬৪.৪ × ৭৭.৯ × ৭.৯ mm",
		"February 2025":                                           "ফেব্রুয়ারি ২০২৫",
		"Available":                                               "উপলব্ধ",
		"MediaTek Dimensity 6300":                                 "মিডিয়াটেক ডাইমেনসিটি ৬৩০০",
		"191 g (according to 8 GB variant)":                       "১৯১ g (৮ GB ভেরিয়েন্ট অনুসারে)",
		"GSM 850 / 900 / 1800 / 1900":                             "জিএসএম ৮৫০ / ৯০০ / ১৮০০ / ১৯০০",
		"HSDPA 850 / 900 / 1900 / 2100":                           "এইচএসডিপিএ ৮৫০ / ৯০০ / ১৯০০ / ২১০০",
		"LTE":                                                     "এলটিই",
		"SA/NSA/Sub6":                                             "এসএ/এনএসএ/সাব৬",
		"128 GB":                                                  "১২৮ GB",
		"microSDXC (hybrid)":                                      "মাইক্রো এসডিএক্সসি (হাইব্রিড)",
		"2.4 GHz":                                                 "২.৪ GHz",
		"25 W":                                                    "২৫ W",
	}
	for k, v := range specific {
		common[k] = v
	}
	return common
}

// Seed inserts specification records for the product identified by slug 'samsung-galaxy-m16-5g'
func (s *SpecificationSeederMobileSamsungGalaxyM165G) Seed(db *gorm.DB) error {
	productSlug := "samsung-galaxy-m16-5g"
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
		"wifi_support":                  "Wi-Fi 802.11 a/b/g/n/ac",
		"optical_zoom":                  "None",
		"front_camera_video_resolution": "",
		"sensors":                       "Side fingerprint, Accelerometer, Gyro, Proximity, Compass",
		"battery_type":                  "Li-Ion (non-removable)",
		"5g_support":                    "Yes",
		"special_features":              "6 years of updates (claimed)",
		"ram":                           "6 / 8 GB",
		"resolution":                    "1080 × 2340 px",
		"refresh_rate":                  "90 Hz",
		"camera_video_resolution":       "",
		"available_colors":              "Thunder Black, Mint Green, Blush Pink",
		"gpu_type":                      "Mali-G57 MC2",
		"bluetooth_version":             "5.3",
		"nfc_support":                   "Yes",
		"usb_type":                      "USB-C",
		"operating_system":              "Android 15, One UI 7",
		"battery":                       "5000 mAh",
		"wireless_charging":             "No",
		"audio_quality":                 "",
		"display_size":                  "6.7 inches",
		"cpu_type":                      "Octa-core (2×2.4 GHz Cortex-A76 + 6×2.0 GHz Cortex-A55)",
		"build_material":                "Glass front, plastic back / frame",
		"network_technology":            "GSM / HSPA / LTE / 5G",
		"2g_bands":                      "GSM 850 / 900 / 1800 / 1900",
		"3g_bands":                      "HSDPA 850 / 900 / 1900 / 2100",
		"4g_bands":                      "LTE",
		"5g_bands":                      "SA/NSA/Sub6",
		"front_camera":                  "13 MP",
		"fast_charging":                 "25 W wired",
		"positioning_system":            "GPS, GLONASS, GALILEO, BDS",
		"chipset":                       "Dimensity 6300 (6 nm)",
		"water_resistance":              "IP54 (splash)",
		"rear_camera":                   "50 MP + 5 MP + 2 MP",
		"sim_card_type":                 "Nano-SIM + Nano-SIM (hybrid)",
		"audio_jack":                    "USB-C (no 3.5mm)",
		"display_type":                  "Super AMOLED",
		"camera_features":               "LED flash, Auto-focus, Depth / Macro",
		"loudspeaker_quality":           "Stereo",
		"storage":                       "128 GB (maybe more)",
		"dimensions":                    "164.4 × 77.9 × 7.9 mm",
		"announcement_date":             "February 2025",
		"device_status":                 "Available",
		"processor":                     "MediaTek Dimensity 6300",
		"weight":                        "191 g (according to 8 GB variant)",
		"internal_memory_capacity":      "128 GB",
		"card_slot_type":                "microSDXC (hybrid)",
		"processor_speed":               "2.4 GHz",
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
