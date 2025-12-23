package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederMobileSamsungGalaxyM535G seeds specifications/options for product 'samsung-galaxy-m53-5g'
type SpecificationSeederMobileSamsungGalaxyM535G struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSamsungGalaxyM535G creates a new seeder instance
func NewSpecificationSeederMobileSamsungGalaxyM535G() *SpecificationSeederMobileSamsungGalaxyM535G {
	return &SpecificationSeederMobileSamsungGalaxyM535G{BaseSeeder: BaseSeeder{name: "Specifications for Samsung Galaxy M53 5G"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSamsungGalaxyM535G) getBanglaTranslations() map[string]string {
	common := s.GetCommonBanglaTranslations()
	specific := map[string]string{
		"MediaTek Dimensity 900": "মিডিয়াটেক ডাইমেনসিটি ৯০০",
		"Mali-G68 MC4":           "মালি-জি৬৮ এমসি৪",
		"Octa-core (2×2.4 GHz Cortex-A78 + 6×2.0 GHz Cortex-A55)": "অক্টা-কোর (২×২.৪ গিগাহার্টজ কর্টেক্স-এ৭৮ + ৬×২.০ গিগাহার্টজ কর্টেক্স-এ৫৫)",
		"Dimensity 900 (6 nm)":               "ডাইমেনসিটি ৯০০ (৬ ন্যানোমিটার)",
		"Super AMOLED Plus, 120 Hz":          "সুপার অ্যামোলেড প্লাস, ১২০ হার্জ",
		"2408 × 1080 px":                     "২৪০৮ × ১০৮০ পিক্সেল",
		"6.7 inches":                         "৬.৭ ইঞ্চি",
		"108 MP + 8 MP + 2 MP + 2 MP":        "১০৮ এমপি + ৮ এমপি + ২ এমপি + ২ এমপি",
		"32 MP":                              "৩২ এমপি",
		"4K @ 30fps; 1080p @ 30/60fps":       "৪কে @ ৩০ এফপিএস; ১০৮০পি @ ৩০/৬০ এফপিএস",
		"4K @ 30fps; 1080p @ 30fps":          "৪কে @ ৩০ এফপিএস; ১০৮০পি @ ৩০ এফপিএস",
		"Green, Blue, Emerald Brown":         "সবুজ, নীল, পান্না বাদামী",
		"164.7 × 77.0 × 7.4 mm":              "১৬৪.৭ × ৭৭.০ × ৭.৪ মিলিমিটার",
		"176 g":                              "১৭৬ গ্রাম",
		"5000 mAh":                           "৫০০০ মিলিঅ্যাম্পিয়ার আওয়ার",
		"25 W wired":                         "২৫ ওয়াট তারযুক্ত",
		"Android 12, One UI 4.1 (or newer)":  "অ্যান্ড্রয়েড ১২, ওয়ান ইউআই ৪.১ (বা নতুন)",
		"GSM / HSPA / LTE / 5G":              "জিএসএম / এইচএসপিএ / এলটিই / ৫জি",
		"Dual SIM (Nano + Nano)":             "ডুয়াল সিম (ন্যানো + ন্যানো)",
		"Wi-Fi 802.11 a/b/g/n/ac, dual-band": "ওয়াই-ফাই ৮০২.১১ এ/বি/জি/এন/এসি, ডুয়াল-ব্যান্ড",
		"GPS, GLONASS, GALILEO, BDS":         "জিপিএস, গ্লোনাস, গ্যালিলিও, বিডিএস",
		"5.2":                                "৫.২",
		"USB-C 2.0":                          "ইউএসবি-সি ২.০",
		"PDAF, LED flash, HDR":               "পিডিএএফ, এলইডি ফ্ল্যাশ, এইচডিআর",
		"April 2022":                         "এপ্রিল ২০২২",
		"Glass front, plastic frame / back":  "কাচের সামনের অংশ, প্লাস্টিকের ফ্রেম / পিছন",
		"128 / 256 GB + microSD":             "১২৮ / ২৫৬ জিবি + মাইক্রোএসডি",
		"120 Hz":                             "১২০ হার্জ",
		"Li-Ion (non-removable)":             "লি-আয়ন (অপসারণযোগ্য নয়)",
		"6 / 8 GB":                           "৬ / ৮ জিবি",
		"Side fingerprint, Accelerometer, Gyro, Proximity, Compass": "পাশের ফিঙ্গারপ্রিন্ট, অ্যাকসেলেরোমিটার, জাইরো, প্রক্সিমিটি, কম্পাস",
		"Mono":                        "মোনো",
		"GSM 850 / 900 / 1800 / 1900": "জিএসএম ৮৫০ / ৯০০ / ১৮০০ / ১৯০০",
		"HSDPA 850 / 900 / 1700(AWS) / 1900 / 2100":                                   "এইচএসডিপিএ ৮৫০ / ৯০০ / ১৭০০(এডব্লিউএস) / ১৯০০ / ২১০০",
		"LTE bands 1, 3, 5, 7, 8, 20, 28, 38, 40, 41":                                 "এলটিই ব্যান্ড ১, ৩, ৫, ৭, ৮, ২০, ২৮, ৩৮, ৪০, ৪১",
		"SA/NSA/Sub6/mmWave - 5G bands 1, 3, 5, 7, 8, 20, 28, 38, 40, 41, 66, 77, 78": "এসএ/এনএসএ/সাব৬/এমএমওয়েভ - ৫জি ব্যান্ড ১, ৩, ৫, ৭, ৮, ২০, ২৮, ৩৮, ৪০, ৪১, ৬৬, ৭৭, ৭৮",
		"25W":                          "২৫ওয়াট",
		"Up to 2.4 GHz":                "২.৪ গিগাহার্টজ পর্যন্ত",
		"128GB 8GB RAM, 256GB 8GB RAM": "১২৮জিবি ৮জিবি র্যাম, ২৫৬জিবি ৮জিবি র্যাম",
		"microSDXC (dedicated slot)":   "মাইক্রোএসডিএক্সসি (ডেডিকেটেড স্লট)",
		"108 MP, f/1.8, (wide), 1/1.52\", 0.7µm, PDAF, OIS\n8 MP, f/2.2, (ultrawide), 1/4.0\", 1.12µm\n2 MP, f/2.4, (macro)\n2 MP, f/2.4, (depth)": "১০৮ এমপি, এফ/১.৮, (ওয়াইড), ১/১.৫২\", ০.৭মাইক্রোমিটার, পিডিএএফ, ওআইএস\n৮ এমপি, এফ/২.২, (আল্ট্রাওয়াইড), ১/৪.০\", ১.১২মাইক্রোমিটার\n২ এমপি, এফ/২.৪, (ম্যাক্রো)\n২ এমপি, এফ/২.৪, (ডেপথ)",
		"No":        "না",
		"None":      "কোনোটিই নয়",
		"Available": "উপলব্ধ",
	}
	for k, v := range specific {
		common[k] = v
	}
	return common
}

// Seed inserts specification records for the product identified by slug 'samsung-galaxy-m53-5g'
func (s *SpecificationSeederMobileSamsungGalaxyM535G) Seed(db *gorm.DB) error {
	productSlug := "samsung-galaxy-m53-5g"
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
		"sensors":                       "Side fingerprint, Accelerometer, Gyro, Proximity, Compass",
		"gpu_type":                      "Mali-G68 MC4",
		"ram":                           "6 / 8 GB",
		"display_type":                  "Super AMOLED Plus, 120 Hz",
		"network_technology":            "GSM / HSPA / LTE / 5G",
		"sim_card_type":                 "Dual SIM (Nano + Nano)",
		"loudspeaker_quality":           "Mono",
		"resolution":                    "2408 × 1080 px",
		"wifi_support":                  "Wi-Fi 802.11 a/b/g/n/ac, dual-band",
		"positioning_system":            "GPS, GLONASS, GALILEO, BDS",
		"special_features":              "",
		"device_status":                 "Available",
		"weight":                        "176 g",
		"water_resistance":              "",
		"bluetooth_version":             "5.2",
		"usb_type":                      "USB-C 2.0",
		"camera_features":               "PDAF, LED flash, HDR",
		"announcement_date":             "April 2022",
		"cpu_type":                      "Octa-core (2×2.4 GHz Cortex-A78 + 6×2.0 GHz Cortex-A55)",
		"build_material":                "Glass front, plastic frame / back",
		"rear_camera":                   "108 MP + 8 MP + 2 MP + 2 MP",
		"optical_zoom":                  "None",
		"audio_quality":                 "",
		"available_colors":              "Green, Blue, Emerald Brown",
		"storage":                       "128 / 256 GB + microSD",
		"refresh_rate":                  "120 Hz",
		"battery":                       "5000 mAh",
		"battery_type":                  "Li-Ion (non-removable)",
		"wireless_charging":             "No",
		"5g_support":                    "Yes",
		"display_size":                  "6.7 inches",
		"front_camera":                  "32 MP",
		"front_camera_video_resolution": "4K @ 30fps; 1080p @ 30fps",
		"audio_jack":                    "No",
		"processor":                     "MediaTek Dimensity 900",
		"dimensions":                    "164.7 × 77.0 × 7.4 mm",
		"nfc_support":                   "",
		"fast_charging":                 "25 W wired",
		"chipset":                       "Dimensity 900 (6 nm)",
		"camera_video_resolution":       "4K @ 30fps; 1080p @ 30/60fps",
		"operating_system":              "Android 12, One UI 4.1 (or newer)",
		"2g_bands":                      "GSM 850 / 900 / 1800 / 1900",
		"3g_bands":                      "HSDPA 850 / 900 / 1700(AWS) / 1900 / 2100",
		"4g_bands":                      "LTE bands 1, 3, 5, 7, 8, 20, 28, 38, 40, 41",
		"5g_bands":                      "SA/NSA/Sub6/mmWave - 5G bands 1, 3, 5, 7, 8, 20, 28, 38, 40, 41, 66, 77, 78",
		"charging_speed":                "25W",
		"processor_speed":               "Up to 2.4 GHz",
		"internal_memory_capacity":      "128GB 8GB RAM, 256GB 8GB RAM",
		"card_slot_type":                "microSDXC (dedicated slot)",
		"quad_camera_setup":             "108 MP, f/1.8, (wide), 1/1.52\", 0.7µm, PDAF, OIS\n8 MP, f/2.2, (ultrawide), 1/4.0\", 1.12µm\n2 MP, f/2.4, (macro)\n2 MP, f/2.4, (depth)",
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
