package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederMobileSamsungGalaxyA255G seeds specifications/options for product 'samsung-galaxy-a25-5g'
type SpecificationSeederMobileSamsungGalaxyA255G struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSamsungGalaxyA255G creates a new seeder instance
func NewSpecificationSeederMobileSamsungGalaxyA255G() *SpecificationSeederMobileSamsungGalaxyA255G {
	return &SpecificationSeederMobileSamsungGalaxyA255G{BaseSeeder: BaseSeeder{name: "Specifications for Samsung Galaxy A25 5G"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSamsungGalaxyA255G) getBanglaTranslations() map[string]string {
	common := s.GetCommonBanglaTranslations()
	specific := map[string]string{
		"6.5 inches":                    "৬.৫ ইঞ্চি",
		"Super AMOLED, 120 Hz":          "সুপার অ্যামোলেড, ১২০ হার্জ",
		"2340 × 1080 pixels (~396 ppi)": "২৩৪০ × ১০৮০ পিক্সেল (~৩৯৬ পিপিআই)",
		"120 Hz":                        "১২০ Hz",
		"Exynos 1280 (5 nm)":            "এক্সিনস ১২৮০ (৫ ন্যানোমিটার)",
		"Exynos 1280":                   "এক্সিনস ১২৮০",
		"Octa-core (2×2.4 GHz Cortex-A78 + 6×2.0 GHz Cortex-A55)": "অক্টা-কোর (২×২.৪ GHz কর্টেক্স-A78 + ৬×২.০ GHz কর্টেক্স-A55)",
		"Mali-G68":                            "মালি-G68",
		"6 / 8 GB":                            "৬ / ৮ জিবি",
		"128 / 256 GB + microSD (up to 1 TB)": "১২৮ / ২৫৬ জিবি + মাইক্রোএসডি (১ টেরাবাইট পর্যন্ত)",
		"128 / 256 GB":                        "১২৮ / ২৫৬ জিবি",
		"microSDXC (up to 1 TB)":              "মাইক্রোএসডি (১ টেরাবাইট পর্যন্ত)",
		"50 MP + 8 MP + 2 MP":                 "৫০ মেগাপিক্সেল + ৮ মেগাপিক্সেল + ২ মেগাপিক্সেল",
		"OIS (main), LED flash":               "OIS (মেইন), LED ফ্ল্যাশ",
		"50 MP (wide) + 8 MP (ultrawide) + 2 MP (macro)": "৫০ MP (ওয়াইড) + ৮ MP (আল্ট্রাওয়াইড) + ২ MP (ম্যাক্রো)",
		"4K @ 30fps; 1080p @ 30/60fps":                   "৪K @ ৩০fps; ১০৮০p @ ৩০/৬০fps",
		"None":                                           "কোনটি নয়",
		"13 MP":                                          "১৩ MP",
		"1080p @ 30fps":                                  "১০৮০p @ ৩০fps",
		"5000 mAh":                                       "৫০০০ মিলিঅ্যাম্পিয়ার আওয়ার",
		"Li-Ion (non-removable)":                         "লি-আয়ন (অপসারণযোগ্য নয়)",
		"25 W wired":                                     "২৫ ওয়াট তারযুক্ত",
		"25W wired":                                      "২৫ ওয়াট তারযুক্ত",
		"No":                                             "না",
		"Android 14, One UI 6":                           "অ্যান্ড্রয়েড ১৪, ওয়ান UI ৬",
		"Stereo":                                         "স্টেরিও",
		"Glass front, plastic frame / back":              "গ্লাস ফ্রন্ট, প্লাস্টিক ফ্রেম / ব্যাক",
		"197 g":                                          "১৯৭ গ্রাম",
		"161 × 76.5 × 8.3 mm":                            "১৬১ × ৭৬.৫ × ৮.৩ মিলিমিটার",
		"GSM / HSPA / LTE / 5G":                          "জিএসএম / এইচএসপিএ / এলটিই / ৫জি",
		"Wi-Fi 802.11 a/b/g/n/ac":                        "ওয়াই-ফাই ৮০২.১১ a/b/g/n/ac",
		"5.3":                                            "৫.৩",
		"Yes":                                            "হ্যাঁ",
		"USB-C 2.0":                                      "ইউএসবি-সি ২.০",
		"GPS, GLONASS, Galileo, BDS":                     "জিপিএস, গ্লোনাস, গ্যালিলিও, বিডিএস",
		"Accelerometer, Gyro, Proximity, Compass": "অ্যাকসেলেরোমিটার, জাইরো, প্রক্সিমিটি, কম্পাস",
		"Nano-SIM (or hybrid)":                    "ন্যানো-SIM (বা হাইব্রিড)",
		"Blue Black, Blue, Light Blue, Yellow":    "নীল-কালো, নীল, হালকা নীল, হলুদ",
		"December 2023":                           "ডিসেম্বর ২০২৩",
		"Available":                               "উপলব্ধ",
		"Not specified":                           "উল্লেখ করা হয়নি",
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

// Seed inserts specification records for the product identified by slug 'samsung-galaxy-a25-5g'
func (s *SpecificationSeederMobileSamsungGalaxyA255G) Seed(db *gorm.DB) error {
	productSlug := "samsung-galaxy-a25-5g"
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
		"display_type":                  "Super AMOLED, 120 Hz",
		"resolution":                    "2340 × 1080 pixels (~396 ppi)",
		"refresh_rate":                  "120 Hz",
		"screen_protection":             "Corning Gorilla Glass 5",
		"processor":                     "Exynos 1280",
		"chipset":                       "Exynos 1280 (5 nm)",
		"cpu_type":                      "Octa-core (2×2.4 GHz Cortex-A78 + 6×2.0 GHz Cortex-A55)",
		"gpu_type":                      "Mali-G68",
		"processor_speed":               "2.4 GHz",
		"ram":                           "6 / 8 GB",
		"storage":                       "128 / 256 GB + microSD (up to 1 TB)",
		"internal_memory_capacity":      "128 / 256 GB",
		"card_slot_type":                "microSDXC (up to 1 TB)",
		"rear_camera":                   "50 MP + 8 MP + 2 MP",
		"camera_features":               "OIS (main), LED flash",
		"quad_camera_setup":             "50 MP (wide) + 8 MP (ultrawide) + 2 MP (macro)",
		"camera_video_resolution":       "4K @ 30fps; 1080p @ 30/60fps",
		"optical_zoom":                  "None",
		"front_camera":                  "13 MP",
		"front_camera_video_resolution": "1080p @ 30fps",
		"battery":                       "5000 mAh",
		"battery_type":                  "Li-Ion (non-removable)",
		"fast_charging":                 "25 W wired",
		"charging_speed":                "25W wired",
		"wireless_charging":             "No",
		"operating_system":              "Android 14, One UI 6",
		"audio_quality":                 "Not specified",
		"loudspeaker_quality":           "Stereo",
		"audio_jack":                    "No",
		"build_material":                "Glass front, plastic frame / back",
		"weight":                        "197 g",
		"dimensions":                    "161 × 76.5 × 8.3 mm",
		"water_resistance":              "None",
		"network_technology":            "GSM / HSPA / LTE / 5G",
		"2g_bands":                      "GSM 850 / 900 / 1800 / 1900",
		"3g_bands":                      "HSDPA 850 / 900 / 1700(AWS) / 1900 / 2100",
		"4g_bands":                      "LTE band 1(2100), 2(1900), 3(1800), 4(1700/2100), 5(850), 7(2600), 8(900), 12(700), 13(700), 17(700), 20(800), 26(850), 28(700), 32(1500), 38(2600), 40(2300), 41(2500), 66(1700/2100)",
		"5g_bands":                      "SA/NSA/Sub6",
		"5g_support":                    "Yes",
		"wifi_support":                  "Wi-Fi 802.11 a/b/g/n/ac",
		"bluetooth_version":             "5.3",
		"nfc_support":                   "Yes",
		"usb_type":                      "USB-C 2.0",
		"positioning_system":            "GPS, GLONASS, Galileo, BDS",
		"sensors":                       "Accelerometer, Gyro, Proximity, Compass",
		"special_features":              "None",
		"sim_card_type":                 "Nano-SIM (or hybrid)",
		"sar_rating":                    "0.41 W/kg (head) 1.35 W/kg (body)",
		"sar_rating_eu":                 "0.41 W/kg (head) 1.35 W/kg (body)",
		"available_colors":              "Blue Black, Blue, Light Blue, Yellow",
		"model_variants":                "SM-A256E, SM-A256B, SM-A256N",
		"announcement_date":             "December 2023",
		"device_status":                 "Available",
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
