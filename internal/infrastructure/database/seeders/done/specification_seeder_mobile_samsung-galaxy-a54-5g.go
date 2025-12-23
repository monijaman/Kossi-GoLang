package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederMobileSamsungGalaxyA545G seeds specifications/options for product 'samsung-galaxy-a54-5g'
type SpecificationSeederMobileSamsungGalaxyA545G struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSamsungGalaxyA545G creates a new seeder instance
func NewSpecificationSeederMobileSamsungGalaxyA545G() *SpecificationSeederMobileSamsungGalaxyA545G {
	return &SpecificationSeederMobileSamsungGalaxyA545G{BaseSeeder: BaseSeeder{name: "Specifications for Samsung Galaxy A54 5G"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSamsungGalaxyA545G) getBanglaTranslations() map[string]string {
	common := s.GetCommonBanglaTranslations()
	specific := map[string]string{
		"6.4 inches":              "৬.৪ ইঞ্চি",
		"Super AMOLED, 120 Hz":    "সুপার অ্যামোলেড, ১২০ হার্জ",
		"2340 × 1080 px (19.5:9)": "২৩৪০ × ১০৮০ পিক্সেল (১৯.৫:৯)",
		"120 Hz":                  "১২০ হার্জ",
		"Gorilla Glass 5":         "গরিলা গ্লাস ৫",
		"Exynos 1380 (5 nm)":      "এক্সিনস ১৩৮০ (৫ ন্যানোমিটার)",
		"Exynos 1380":             "এক্সিনস ১৩৮০",
		"Octa-core (4×2.4 GHz Cortex-A78 + 4×2.0 GHz Cortex-A55)": "অক্টা-কোর (৪×২.৪ GHz কর্টেক্স-A78 + ৪×২.০ GHz কর্টেক্স-A55)",
		"Mali-G68 (MP5)":         "মালি-G68 (MP5)",
		"6 / 8 GB":               "৬ / ৮ জিবি",
		"128 / 256 GB + microSD": "১২৮ / ২৫৬ জিবি + মাইক্রোএসডি",
		"128 / 256 GB":           "১২৮ / ২৫৬ জিবি",
		"Nano-SIM or Hybrid Dual SIM (market dependent)": "ন্যানো-SIM অথবা হাইব্রিড ডুয়াল SIM (মার্কেট নির্ভর)",
		"50 MP + 8 MP + 5 MP":                            "৫০ মেগাপিক্সেল + ৮ মেগাপিক্সেল + ৫ মেগাপিক্সেল",
		"OIS (main), LED flash":                          "ওআইএস (প্রধান), এলইডি ফ্ল্যাশ",
		"4K @ 30fps; 1080p @ 30/60fps":                   "৪কে @ ৩০ এফপিএস; ১০৮০পি @ ৩০/৬০ এফপিএস",
		"None":                                           "কোনটি নয়",
		"32 MP":                                          "৩২ MP",
		"4K @ 30fps; 1080p @ 30fps":                      "৪K @ ৩০fps; ১০৮০p @ ৩০fps",
		"5000 mAh":                                       "৫০০০ মিলিঅ্যাম্পিয়ার আওয়ার",
		"Li-Ion (non-removable)":                         "লি-আয়ন (অপসারণযোগ্য নয়)",
		"25 W wired":                                     "২৫ ওয়াট তারযুক্ত",
		"No":                                             "না",
		"Android 13, One UI 5.1":                         "অ্যান্ড্রয়েড ১৩, ওয়ান UI ৫.১",
		"Not specified":                                  "উল্লেখ করা হয়নি",
		"Stereo":                                         "স্টেরিও",
		"Glass front (Gorilla Glass 5), plastic frame / back": "গ্লাস ফ্রন্ট (গরিলা গ্লাস ৫), প্লাস্টিক ফ্রেম / ব্যাক",
		"202 g":                              "২০২ গ্রাম",
		"158.2 × 76.7 × 8.2 mm":              "১৫৮.২ × ৭৬.৭ × ৮.২ মিলিমিটার",
		"IP67 (up to 1 m for 30 min)":        "আইপি৬৭ (১ মিটার পর্যন্ত ৩০ মিনিট)",
		"GSM / HSPA / LTE / 5G":              "জিএসএম / এইচএসপিএ / এলটিই / ৫জি",
		"Lime, Graphite, Violet, White":      "লাইম, গ্রাফাইট, ভায়োলেট, হোয়াইট",
		"March 2023":                         "মার্চ ২০২৩",
		"Available":                          "উপলব্ধ",
		"Wi-Fi 802.11 a/b/g/n/ac, dual-band": "ওয়াই-ফাই ৮০২.১১ এ/বি/জি/এন/এসি, ডুয়াল-ব্যান্ড",
		"5.2":                                "৫.২",
		"Yes (region dependent)":             "হ্যাঁ (অঞ্চল নির্ভর)",
		"USB-C 2.0":                          "ইউএসবি-সি ২.০",
		"GPS, GLONASS, GALILEO, BDS":         "জিপিএস, গ্লোনাস, গ্যালিলিও, বিডিএস",
		"Fingerprint (in-display, optical), Accelerometer, Gyro, Compass, Barometer": "ফিঙ্গারপ্রিন্ট (ইন-ডিসপ্লে, অপটিক্যাল), অ্যাক্সেলেরোমিটার, জাইরো, কম্পাস, ব্যারোমিটার",
		"4 OS upgrades, 5 years of security updates":                                 "৪ OS আপগ্রেড, ৫ বছরের নিরাপত্তা আপডেট",
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

// Seed inserts specification records for the product identified by slug 'samsung-galaxy-a54-5g'
func (s *SpecificationSeederMobileSamsungGalaxyA545G) Seed(db *gorm.DB) error {
	productSlug := "samsung-galaxy-a54-5g"
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
		"display_size":                  "6.4 inches",
		"display_type":                  "Super AMOLED, 120 Hz",
		"resolution":                    "2340 × 1080 px (19.5:9)",
		"refresh_rate":                  "120 Hz",
		"screen_protection":             "Gorilla Glass 5",
		"chipset":                       "Exynos 1380 (5 nm)",
		"processor":                     "Exynos 1380",
		"cpu_type":                      "Octa-core (4×2.4 GHz Cortex-A78 + 4×2.0 GHz Cortex-A55)",
		"gpu_type":                      "Mali-G68 (MP5)",
		"ram":                           "6 / 8 GB",
		"storage":                       "128 / 256 GB + microSD",
		"internal_memory_capacity":      "128 / 256 GB",
		"card_slot_type":                "microSD",
		"rear_camera":                   "50 MP + 8 MP + 5 MP",
		"quad_camera_setup":             "50 MP (wide) + 8 MP (ultrawide) + 5 MP (macro)",
		"camera_features":               "OIS (main), LED flash",
		"camera_video_resolution":       "4K @ 30fps; 1080p @ 30/60fps",
		"optical_zoom":                  "None",
		"front_camera":                  "32 MP",
		"front_camera_video_resolution": "4K @ 30fps; 1080p @ 30fps",
		"battery":                       "5000 mAh",
		"battery_type":                  "Li-Ion (non-removable)",
		"fast_charging":                 "25 W wired",
		"charging_speed":                "25 W wired",
		"wireless_charging":             "No",
		"operating_system":              "Android 13, One UI 5.1",
		"audio_quality":                 "Not specified",
		"loudspeaker_quality":           "Stereo",
		"audio_jack":                    "No",
		"build_material":                "Glass front (Gorilla Glass 5), plastic frame / back",
		"weight":                        "202 g",
		"dimensions":                    "158.2 × 76.7 × 8.2 mm",
		"water_resistance":              "IP67 (up to 1 m for 30 min)",
		"network_technology":            "GSM / HSPA / LTE / 5G",
		"2g_bands":                      "Not specified",
		"3g_bands":                      "Not specified",
		"4g_bands":                      "Not specified",
		"5g_bands":                      "Not specified",
		"5g_support":                    "Yes",
		"wifi_support":                  "Wi-Fi 802.11 a/b/g/n/ac, dual-band",
		"bluetooth_version":             "5.2",
		"nfc_support":                   "Yes (region dependent)",
		"usb_type":                      "USB-C 2.0",
		"positioning_system":            "GPS, GLONASS, GALILEO, BDS",
		"sensors":                       "Fingerprint (in-display, optical), Accelerometer, Gyro, Compass, Barometer",
		"special_features":              "4 OS upgrades, 5 years of security updates",
		"sim_card_type":                 "Nano-SIM or Hybrid Dual SIM (market dependent)",
		"sar_rating":                    "Not specified",
		"sar_rating_eu":                 "Not specified",
		"available_colors":              "Lime, Graphite, Violet, White",
		"model_variants":                "Not specified",
		"announcement_date":             "March 2023",
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
