package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederMobileSamsungGalaxyA265G seeds specifications/options for product 'samsung-galaxy-a26-5g'
type SpecificationSeederMobileSamsungGalaxyA265G struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSamsungGalaxyA265G creates a new seeder instance
func NewSpecificationSeederMobileSamsungGalaxyA265G() *SpecificationSeederMobileSamsungGalaxyA265G {
	return &SpecificationSeederMobileSamsungGalaxyA265G{BaseSeeder: BaseSeeder{name: "Specifications for Samsung Galaxy A26 5G"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSamsungGalaxyA265G) getBanglaTranslations() map[string]string {
	common := s.GetCommonBanglaTranslations()
	specific := map[string]string{
		"6.5 inches":                            "৬.৫ ইঞ্চি",
		"PLS LCD, 120Hz":                        "পিএলএস এলসিডি, ১২০ হার্জ",
		"1080 x 2400 pixels (~405 ppi density)": "১০৮০ × ২৪০০ পিক্সেল (~৪০৫ পিপিআই)",
		"120Hz":                                 "১২০ হার্জ",
		"Corning Gorilla Glass 3":               "কর্নিং গরিলা গ্লাস ৩",
		"MediaTek MT6833 Dimensity 6020 (7 nm)": "মিডিয়াটেক MT6833 ডাইমেনসিটি ৬০২০ (৭ ন্যানোমিটার)",
		"Dimensity 6020":                        "ডাইমেনসিটি ৬০২০",
		"Octa-core (2x2.2 GHz Cortex-A76 & 6x2.0 GHz Cortex-A55)": "অক্টা-কোর (২×২.২ গিগাহার্জ কোর্টেক্স-A76 ও ৬×২.০ গিগাহার্জ কোর্টেক্স-A55)",
		"Mali-G57 MC2":                "মালি-G57 MC2",
		"6 GB / 8 GB":                 "৬ জিবি / ৮ জিবি",
		"128 GB":                      "১২৮ জিবি",
		"50 MP + 2 MP":                "৫০ MP + ২ MP",
		"50 MP (wide) + 2 MP (depth)": "৫০ MP (ওয়াইড) + ২ MP (ডেপ্থ)",
		"None":                        "কোনটি নয়",
		"13 MP":                       "১৩ MP",
		"4500 mAh":                    "৪৫০০ মিলিঅ্যাম্পিয়ার আওয়ার",
		"Android 14, One UI 6":        "অ্যান্ড্রয়েড ১৪, ওয়ান UI ৬",
		"Plastic frame, plastic back": "প্লাস্টিক ফ্রেম, প্লাস্টিক ব্যাক",
		"189 g (6.67 oz)":             "১৮৯ গ্রাম (৬.৬৭ আউন্স)",
		"161.3 x 73.9 x 8.4 mm":       "১৬১.৩ × ৭৩.৯ × ৮.৪ মিলিমিটার",
		"GSM / HSPA / LTE / 5G":       "জিএসএম / এইচএসপিএ / এলটিই / ৫G",
		"Black, Blue":                 "ব্ল্যাক, ব্লু",
		"August 2025":                 "আগস্ট ২০২৫",
		"Upcoming":                    "আসন্ন",
		"Not specified":               "উল্লেখ করা হয়নি",
	}
	result := make(map[string]string)
	for k, v := range common {
		result[k] = v
	}
	for k, v := range specific {
		result[k] = v
	}
	return result
}

// Seed inserts specification records for the product identified by slug 'samsung-galaxy-a26-5g'
func (s *SpecificationSeederMobileSamsungGalaxyA265G) Seed(db *gorm.DB) error {
	productSlug := "samsung-galaxy-a26-5g"
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
		"display_size":                  "6.5 inches",
		"display_type":                  "PLS LCD, 120Hz",
		"resolution":                    "1080 x 2400 pixels (~405 ppi density)",
		"refresh_rate":                  "120Hz",
		"screen_protection":             "Corning Gorilla Glass 3",
		"processor":                     "Dimensity 6020",
		"chipset":                       "MediaTek MT6833 Dimensity 6020 (7 nm)",
		"cpu_type":                      "Octa-core (2x2.2 GHz Cortex-A76 & 6x2.0 GHz Cortex-A55)",
		"gpu_type":                      "Mali-G57 MC2",
		"processor_speed":               "2.2 GHz",
		"ram":                           "6 GB / 8 GB",
		"storage":                       "128 GB",
		"internal_memory_capacity":      "128 GB",
		"card_slot_type":                "microSDXC (dedicated slot)",
		"rear_camera":                   "50 MP + 2 MP",
		"camera_features":               "LED flash",
		"quad_camera_setup":             "50 MP (wide) + 2 MP (depth)",
		"camera_video_resolution":       "1080p @ 30fps",
		"optical_zoom":                  "None",
		"front_camera":                  "13 MP",
		"front_camera_video_resolution": "1080p @ 30fps",
		"battery":                       "4500 mAh",
		"battery_type":                  "Li-Ion (non-removable)",
		"fast_charging":                 "25 W wired",
		"charging_speed":                "25W wired",
		"wireless_charging":             "No",
		"operating_system":              "Android 14, One UI 6",
		"audio_quality":                 "Not specified",
		"loudspeaker_quality":           "Mono",
		"audio_jack":                    "Yes, 3.5 mm",
		"build_material":                "Plastic frame, plastic back",
		"weight":                        "189 g (6.67 oz)",
		"dimensions":                    "161.3 x 73.9 x 8.4 mm",
		"water_resistance":              "None",
		"network_technology":            "GSM / HSPA / LTE / 5G",
		"2g_bands":                      "GSM 850 / 900 / 1800 / 1900",
		"3g_bands":                      "HSDPA 850 / 900 / 1700(AWS) / 1900 / 2100",
		"4g_bands":                      "LTE band 1(2100), 2(1900), 3(1800), 4(1700/2100), 5(850), 7(2600), 8(900), 12(700), 13(700), 17(700), 20(800), 26(850), 28(700), 32(1500), 38(2600), 40(2300), 41(2500), 66(1700/2100)",
		"5g_bands":                      "SA/NSA/Sub6",
		"5g_support":                    "Yes",
		"wifi_support":                  "Wi-Fi 802.11 a/b/g/n/ac",
		"bluetooth_version":             "5.2",
		"nfc_support":                   "Yes (market dependent)",
		"usb_type":                      "USB-C 2.0",
		"positioning_system":            "GPS, GLONASS, GALILEO, BDS",
		"sensors":                       "Fingerprint (side), Accelerometer, Gyro, Proximity, Compass",
		"special_features":              "None",
		"sim_card_type":                 "Nano-SIM or Dual SIM",
		"sar_rating":                    "0.41 W/kg (head) 1.35 W/kg (body)",
		"sar_rating_eu":                 "0.41 W/kg (head) 1.35 W/kg (body)",
		"available_colors":              "Black, Blue",
		"model_variants":                "SM-A265F, SM-A265M, SM-A265N",
		"announcement_date":             "August 2025",
		"device_status":                 "Upcoming",
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
