package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederMobileSamsungGalaxyM565G seeds specifications/options for product 'samsung-galaxy-m56-5g'
type SpecificationSeederMobileSamsungGalaxyM565G struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSamsungGalaxyM565G creates a new seeder instance
func NewSpecificationSeederMobileSamsungGalaxyM565G() *SpecificationSeederMobileSamsungGalaxyM565G {
	return &SpecificationSeederMobileSamsungGalaxyM565G{BaseSeeder: BaseSeeder{name: "Specifications for Samsung Galaxy M56 5G"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSamsungGalaxyM565G) getBanglaTranslations() map[string]string {
	common := s.GetCommonBanglaTranslations()
	specific := map[string]string{
		"Exynos 1480":                        "এক্সিনস ১৪৮০",
		"Xclipse 530":                        "এক্সক্লিপস ৫৩০",
		"Octa-core (4×2.75 GHz + 4×2.0 GHz)": "অক্টা-কোর (৪×২.৭৫ গিগাহার্টজ + ৪×২.০ গিগাহার্টজ)",
		"Exynos 1480 (4 nm)":                 "এক্সিনস ১৪৮০ (৪ ন্যানোমিটার)",
		"Super AMOLED+ 120 Hz":               "সুপার অ্যামোলেড+ ১২০ হার্জ",
		"1080 × 2340 px":                     "১০৮০ × ২৩৪০ পিক্সেল",
		"6.74 inches":                        "৬.৭৪ ইঞ্চি",
		"50 MP + 8 MP + 2 MP":                "৫০ এমপি + ৮ এমপি + ২ এমপি",
		"12 MP":                              "১২ এমপি",
		"4K @ 30fps; 1080p @ 30/60fps":       "৪কে @ ৩০এফপিএস; ১০৮০পি @ ৩০/৬০এফপিএস",
		"4K @ 30fps; 1080p @ 30fps":          "৪কে @ ৩০এফপিএস; ১০৮০পি @ ৩০এফপিএস",
		"Black, Light Green":                 "কালো, হালকা সবুজ",
		"162 × 77.3 × 7.2 mm":                "১৬২ × ৭৭.৩ × ৭.২ মিলিমিটার",
		"180 g":                              "১৮০ গ্রাম",
		"5000 mAh":                           "৫০০০ মিলিঅ্যাম্পিয়ার আওয়ার",
		"45 W wired":                         "৪৫ ওয়াট তারযুক্ত",
		"Android 15, One UI 7":               "অ্যান্ড্রয়েড ১৫, ওয়ান ইউআই ৭",
		"GSM / HSPA / LTE / 5G":              "জিএসএম / এইচএসপিএ / এলটিই / ৫জি",
		"Nano-SIM + Nano-SIM":                "ন্যানো-সিম + ন্যানো-সিম",
		"Wi-Fi 802.11 a/b/g/n/ac/6":          "ওয়াই-ফাই ৮০২.১১ এ/বি/জি/এন/এসি/৬",
		"5.3":                                "৫.৩",
		"USB-C":                              "ইউএসবি-সি",
		"OIS, LED flash":                     "ওআইএস, এলইডি ফ্ল্যাশ",
		"April 2025":                         "এপ্রিল ২০২৫",
		"Glass front (Gorilla Glass Victus+), plastic frame, glass back": "কাচের সামনের অংশ (গরিলা গ্লাস ভিক্টাস+), প্লাস্টিকের ফ্রেম, কাচের পিছন",
		"128 / 256 GB + microSDXC":                                       "১২৮ / ২৫৬ জিবি + মাইক্রোএসডিএক্সসি",
		"120 Hz":                                                         "১২০ হার্জ",
		"Li-Po (non-removable)":                                          "লি-পো (অপসারণযোগ্য নয়)",
		"8 GB":                                                           "৮ জিবি",
		"Accelerometer, Gyro, Proximity, Compass, fingerprint (side?)": "অ্যাকসেলেরোমিটার, জাইরো, প্রক্সিমিটি, কম্পাস, ফিঙ্গারপ্রিন্ট (পাশ?)",
		"Stereo":                      "স্টেরিও",
		"GSM 850 / 900 / 1800 / 1900": "জিএসএম ৮৫০ / ৯০০ / ১৮০০ / ১৯০০",
		"HSDPA 850 / 900 / 1700(AWS) / 1900 / 2100":                                   "এইচএসডিপিএ ৮৫০ / ৯০০ / ১৭০০(এডব্লিউএস) / ১৯০০ / ২১০০",
		"LTE bands 1, 3, 5, 7, 8, 20, 28, 38, 40, 41":                                 "এলটিই ব্যান্ড ১, ৩, ৫, ৭, ৮, ২০, ২৮, ৩৮, ৪০, ৪১",
		"SA/NSA/Sub6/mmWave - 5G bands 1, 3, 5, 7, 8, 20, 28, 38, 40, 41, 66, 77, 78": "এসএ/এনএসএ/সাব৬/এমএমওয়েভ - ৫জি ব্যান্ড ১, ৩, ৫, ৭, ৮, ২০, ২৮, ৩৮, ৪০, ৪১, ৬৬, ৭৭, ৭৮",
		"45W":                          "৪৫ওয়াট",
		"Up to 2.75 GHz":               "২.৭৫ গিগাহার্টজ পর্যন্ত",
		"128GB 8GB RAM, 256GB 8GB RAM": "১২৮জিবি ৮জিবি র্যাম, ২৫৬জিবি ৮জিবি র্যাম",
		"microSDXC (dedicated slot)":   "মাইক্রোএসডিএক্সসি (ডেডিকেটেড স্লট)",
		"50 MP, f/1.8, (wide), 1/1.56\", 1.0µm, PDAF, OIS\n8 MP, f/2.2, (ultrawide), 1/4.0\", 1.12µm\n2 MP, f/2.4, (macro)": "৫০ এমপি, এফ/১.৮, (ওয়াইড), ১/১.৫৬\", ১.০মাইক্রোমিটার, পিডিএএফ, ওআইএস\n৮ এমপি, এফ/২.২, (আল্ট্রাওয়াইড), ১/৪.০\", ১.১২মাইক্রোমিটার\n২ এমপি, এফ/২.৪, (ম্যাক্রো)",
	}
	for k, v := range specific {
		common[k] = v
	}
	return common
}

// Seed inserts specification records for the product identified by slug 'samsung-galaxy-m56-5g'
func (s *SpecificationSeederMobileSamsungGalaxyM565G) Seed(db *gorm.DB) error {
	productSlug := "samsung-galaxy-m56-5g"
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
		"gpu_type":                      "Xclipse 530",
		"network_technology":            "GSM / HSPA / LTE / 5G",
		"camera_features":               "OIS, LED flash",
		"camera_video_resolution":       "4K @ 30fps; 1080p @ 30/60fps",
		"positioning_system":            "",
		"special_features":              "",
		"dimensions":                    "162 × 77.3 × 7.2 mm",
		"bluetooth_version":             "5.3",
		"front_camera":                  "12 MP",
		"device_status":                 "Available",
		"cpu_type":                      "Octa-core (4×2.75 GHz + 4×2.0 GHz)",
		"ram":                           "8 GB",
		"rear_camera":                   "50 MP + 8 MP + 2 MP",
		"audio_jack":                    "No",
		"announcement_date":             "April 2025",
		"display_type":                  "Super AMOLED+ 120 Hz",
		"nfc_support":                   "",
		"fast_charging":                 "45 W wired",
		"5g_support":                    "Yes",
		"storage":                       "128 / 256 GB + microSDXC",
		"sensors":                       "Accelerometer, Gyro, Proximity, Compass, fingerprint (side?)",
		"resolution":                    "1080 × 2340 px",
		"operating_system":              "Android 15, One UI 7",
		"battery":                       "5000 mAh",
		"battery_type":                  "Li-Po (non-removable)",
		"sim_card_type":                 "Nano-SIM + Nano-SIM",
		"audio_quality":                 "",
		"display_size":                  "6.74 inches",
		"processor":                     "Exynos 1480",
		"chipset":                       "Exynos 1480 (4 nm)",
		"refresh_rate":                  "120 Hz",
		"build_material":                "Glass front (Gorilla Glass Victus+), plastic frame, glass back",
		"weight":                        "180 g",
		"wifi_support":                  "Wi-Fi 802.11 a/b/g/n/ac/6",
		"optical_zoom":                  "None",
		"water_resistance":              "",
		"usb_type":                      "USB-C",
		"front_camera_video_resolution": "4K @ 30fps; 1080p @ 30fps",
		"wireless_charging":             "No",
		"loudspeaker_quality":           "Stereo",
		"available_colors":              "Black, Light Green",
		"2g_bands":                      "GSM 850 / 900 / 1800 / 1900",
		"3g_bands":                      "HSDPA 850 / 900 / 1700(AWS) / 1900 / 2100",
		"4g_bands":                      "LTE bands 1, 3, 5, 7, 8, 20, 28, 38, 40, 41",
		"5g_bands":                      "SA/NSA/Sub6/mmWave - 5G bands 1, 3, 5, 7, 8, 20, 28, 38, 40, 41, 66, 77, 78",
		"charging_speed":                "45W",
		"processor_speed":               "Up to 2.75 GHz",
		"internal_memory_capacity":      "128GB 8GB RAM, 256GB 8GB RAM",
		"card_slot_type":                "microSDXC (dedicated slot)",
		"quad_camera_setup":             "50 MP, f/1.8, (wide), 1/1.56\", 1.0µm, PDAF, OIS\n8 MP, f/2.2, (ultrawide), 1/4.0\", 1.12µm\n2 MP, f/2.4, (macro)",
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
