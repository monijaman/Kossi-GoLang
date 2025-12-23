package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederMobileSamsungGalaxyM155G seeds specifications/options for product 'samsung-galaxy-m15-5g'
type SpecificationSeederMobileSamsungGalaxyM155G struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSamsungGalaxyM155G creates a new seeder instance
func NewSpecificationSeederMobileSamsungGalaxyM155G() *SpecificationSeederMobileSamsungGalaxyM155G {
	return &SpecificationSeederMobileSamsungGalaxyM155G{BaseSeeder: BaseSeeder{name: "Specifications for Samsung Galaxy M15 5G"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSamsungGalaxyM155G) getBanglaTranslations() map[string]string {
	common := s.GetCommonBanglaTranslations()
	specific := map[string]string{
		"160.1 × 76.8 × 9.3 mm":   "১৬০.১ × ৭৬.৮ × ৯.৩ mm",
		"Nano-SIM + Nano-SIM":     "ন্যানো-সিম + ন্যানো-সিম",
		"Super AMOLED":            "সুপার AMOLED",
		"2340 × 1080 px":          "২৩৪০ × ১০৮০ px",
		"217 g":                   "২১৭ g",
		"6000 mAh":                "৬০০০ mAh",
		"Available / Unofficial":  "উপলব্ধ / অননুমোদিত",
		"Octa-core":               "অক্টা-কোর",
		"GSM / HSPA / LTE / 5G":   "জিএসএম / এইচএসপিএ / এলটিই / ৫জি",
		"Wi-Fi 802.11 a/b/g/n/ac": "ওয়াই-ফাই ৮০২.১১ a/b/g/n/ac",
		"Li-Ion (non-removable)":  "লি-আয়ন (অপসারণযোগ্য নয়)",
		"No":                      "না",
		"MediaTek Dimensity 810":  "মিডিয়াটেক ডাইমেনসিটি ৮১০",
		"Mali-G57 MC2":            "মালি-G৫৭ MC২",
		"6 / 8 GB":                "৬ / ৮ GB",
		"5.3":                     "৫.৩",
		"LED flash":               "এলইডি ফ্ল্যাশ",
		"13 MP":                   "১৩ MP",
		"Accelerometer, Gyro, Proximity, Compass, Fingerprint": "অ্যাক্সেলেরোমিটার, গাইরো, প্রক্সিমিটি, কম্পাস, ফিঙ্গারপ্রিন্ট",
		"90 Hz":                             "৯০ Hz",
		"Glass front, plastic frame / back": "সামনে গ্লাস, ফ্রেম / পিছনে প্লাস্টিক",
		"None":                              "কোনটি না",
		"Yes":                               "হ্যাঁ",
		"3.5 mm":                            "৩.৫ mm",
		"6.5 inches":                        "৬.৫ ইঞ্চি",
		"50 MP + 5 MP + 2 MP":               "৫০ MP + ৫ MP + ২ MP",
		"1080p @ 30fps":                     "১০৮০p @ ৩০fps",
		"Dimensity 810 (7 nm)":              "ডাইমেনসিটি ৮১০ (৭ nm)",
		"128 / 256 GB + microSD":            "১২৮ / ২৫৬ GB + microSD",
		"Mono":                              "মনো",
		"GSM 850 / 900 / 1800 / 1900":       "জিএসএম ৮৫০ / ৯০০ / ১৮০০ / ১৯০০",
		"HSDPA 850 / 900 / 1900 / 2100":     "এইচএসডিপিএ ৮৫০ / ৯০০ / ১৯০০ / ২১০০",
		"LTE":                               "এলটিই",
		"SA/NSA/Sub6":                       "এসএ/এনএসএ/সাব৬",
		"128 / 256 GB":                      "১২৮ / ২৫৬ GB",
		"microSDXC (dedicated slot)":        "মাইক্রো এসডিএক্সসি (ডেডিকেটেড স্লট)",
		"2.0 GHz":                           "২.০ GHz",
		"15 W":                              "১৫ W",
	}
	for k, v := range specific {
		common[k] = v
	}
	return common
}

// Seed inserts specification records for the product identified by slug 'samsung-galaxy-m15-5g'
func (s *SpecificationSeederMobileSamsungGalaxyM155G) Seed(db *gorm.DB) error {
	productSlug := "samsung-galaxy-m15-5g"
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
		"dimensions":                    "160.1 × 76.8 × 9.3 mm",
		"sim_card_type":                 "Nano-SIM + Nano-SIM",
		"display_type":                  "Super AMOLED",
		"resolution":                    "2340 × 1080 px",
		"weight":                        "217 g",
		"battery":                       "6000 mAh",
		"device_status":                 "Available / Unofficial",
		"cpu_type":                      "Octa-core",
		"network_technology":            "GSM / HSPA / LTE / 5G",
		"2g_bands":                      "GSM 850 / 900 / 1800 / 1900",
		"3g_bands":                      "HSDPA 850 / 900 / 1900 / 2100",
		"4g_bands":                      "LTE",
		"5g_bands":                      "SA/NSA/Sub6",
		"wifi_support":                  "Wi-Fi 802.11 a/b/g/n/ac",
		"battery_type":                  "Li-Ion (non-removable)",
		"wireless_charging":             "No",
		"processor":                     "MediaTek Dimensity 810",
		"gpu_type":                      "Mali-G57 MC2",
		"ram":                           "6 / 8 GB",
		"bluetooth_version":             "5.3",
		"camera_features":               "LED flash",
		"front_camera":                  "13 MP",
		"sensors":                       "Accelerometer, Gyro, Proximity, Compass, Fingerprint",
		"available_colors":              "",
		"refresh_rate":                  "90 Hz",
		"audio_quality":                 "",
		"build_material":                "Glass front, plastic frame / back",
		"usb_type":                      "USB-C",
		"optical_zoom":                  "None",
		"operating_system":              "",
		"5g_support":                    "Yes",
		"audio_jack":                    "3.5 mm",
		"announcement_date":             "",
		"display_size":                  "6.5 inches",
		"water_resistance":              "",
		"rear_camera":                   "50 MP + 5 MP + 2 MP",
		"camera_video_resolution":       "1080p @ 30fps",
		"special_features":              "",
		"chipset":                       "Dimensity 810 (7 nm)",
		"storage":                       "128 / 256 GB + microSD",
		"nfc_support":                   "Yes",
		"front_camera_video_resolution": "1080p @ 30fps",
		"fast_charging":                 "",
		"positioning_system":            "",
		"loudspeaker_quality":           "Mono",
		"internal_memory_capacity":      "128 / 256 GB",
		"card_slot_type":                "microSDXC (dedicated slot)",
		"processor_speed":               "2.0 GHz",
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
